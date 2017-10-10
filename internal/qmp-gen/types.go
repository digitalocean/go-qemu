// Copyright 2016 The go-qemu Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/fatih/camelcase"
)

// a definition is the definition of one QMP API type and its docstring.
type definition struct {
	Docstring string
	JSON      []byte
}

// definitions that are safe to ignore when a docstring is missing.
func ignoreWithoutDocstring(buf []byte) bool {
	switch {
	case bytes.HasPrefix(buf, []byte("{ 'pragma'")):
		return true
	case bytes.HasPrefix(buf, []byte("{ 'include'")):
		return true
	}

	return false
}

func importDefinitions(path string, jsonBuf []byte) ([]definition, error) {
	v := struct {
		I string `json:"include"`
	}{}

	if err := json.Unmarshal(jsonBuf, &v); err != nil {
		return nil, fmt.Errorf("failed to unmarshal include %q json: %s", path, err)
	}

	incPath, err := resolvePath(path, v.I)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve include %q relative to %q: %s", v.I, path, err)
	}

	subdefs, err := readDefinitions(incPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse included file %q: %s", incPath, err)
	}

	return subdefs, nil
}

// readDefinitions reads the definitions from a QAPI spec file.
//
// Includes are processed, so the returned definitions is the full API
// spec.
func readDefinitions(path string) ([]definition, error) {
	var ret []definition

	bs, err := getQAPI(path)
	if err != nil {
		return nil, err
	}

	// Most of the spec has a regular form with "\n\n" between each
	// definition. However, there are a few places with a single
	// newline between definitions. This replacement regularizes the
	// input so that we can process it with less gymnastics.
	bs = bytes.Replace(bs, []byte("}\n##"), []byte("}\n\n##"), -1)

	// As mentioned directly above, most of the spec includes two newlines
	// between definitions. This normalizes stacked include statements.
	bs = bytes.Replace(bs, []byte("}\n{ 'include'"), []byte("}\n\n{ 'include'"), -1)

	for _, part := range bytes.Split(bs, []byte("\n\n")) {
		if len(part) < 1 {
			continue
		}

		fs := bytes.SplitN(part, []byte("\n{"), 2)
		switch len(fs) {
		default:
			return nil, fmt.Errorf("unexpected part of spec file %q: %s", path, string(part))
		case 1:
			if len(fs) == 1 && part[0] == '{' && !ignoreWithoutDocstring(part) {
				return nil, fmt.Errorf("found type definition without a docstring in %q: %s", path, string(part))
			}

			// handle 'include' when no docstring is present
			if bytes.HasPrefix(fs[0], []byte("{ 'include'")) {
				js := pyToJSON(fs[0])

				subdefs, err := importDefinitions(path, js)
				if err != nil {
					return nil, err

				}

				ret = append(ret, subdefs...)
			}

			// This part looks like a non-docstring comment, just skip it.
		case 2:
			docstring := string(fs[0])
			js := pyToJSON(append([]byte{'{'}, fs[1]...))

			if bytes.HasPrefix(fs[0], []byte("{ 'include'")) {
				subdefs, err := importDefinitions(path, js)
				if err != nil {
					return nil, err

				}

				ret = append(ret, subdefs...)
			} else {
				ret = append(ret, definition{docstring, js})
			}
		}
	}

	return ret, nil
}

// parse converts definitions into a set of Go types suitable for rendering.
func parse(defs []definition) (map[name]interface{}, error) {
	ret := map[name]interface{}{}

	for _, def := range defs {
		m := map[string]interface{}{}
		if err := json.Unmarshal(def.JSON, &m); err != nil {
			return nil, err
		}
		switch {
		case m["struct"] != nil:
			v := structType{}
			if err := json.Unmarshal(def.JSON, &v); err != nil {
				return nil, err
			}
			ret[v.Name] = v
		case m["enum"] != nil:
			v := enum{}
			if err := json.Unmarshal(def.JSON, &v); err != nil {
				return nil, err
			}
			ret[v.Name] = v
		case m["union"] != nil:
			v := struct {
				N name          `json:"union"`
				O map[name]name `json:"data"`
				B fieldsOrRef   `json:"base"`
				D name          `json:"discriminator"`
			}{}
			if err := json.Unmarshal(def.JSON, &v); err != nil {
				return nil, err
			}
			if v.D == "" {
				u := simpleUnion{
					Name:    v.N,
					Options: v.O,
				}
				ret[u.Name] = u
			} else {
				u := flatUnion{
					Name:          v.N,
					Options:       v.O,
					Base:          v.B,
					Discriminator: v.D,
				}
				ret[u.Name] = u
			}
		case m["alternate"] != nil:
			v := alternate{}
			if err := json.Unmarshal(def.JSON, &v); err != nil {
				return nil, err
			}
			ret[v.Name] = v
		case m["command"] != nil:
			v := struct {
				N name            `json:"command"`
				I fieldsOrRef     `json:"data"`
				O json.RawMessage `json:"returns"`
				B bool            `json:"boxed"`
			}{}
			if err := json.Unmarshal(def.JSON, &v); err != nil {
				return nil, err
			}
			c := command{
				Name:       v.N,
				Inputs:     v.I,
				BoxedInput: v.B,
			}
			if v.O != nil {
				d := json.NewDecoder(bytes.NewBuffer([]byte(v.O)))
				t, err := d.Token()
				if err != nil {
					return nil, err
				}
				switch val := t.(type) {
				case string:
					c.Output.Type = name(val)
				case json.Delim:
					if val != '[' {
						return nil, fmt.Errorf("unexpected delimiter %q, wanted a dict value", val)
					}
					c.Output.List = true
					t, err := d.Token()
					if err != nil {
						return nil, err
					}
					s, ok := t.(string)
					if !ok {
						return nil, fmt.Errorf("unexpected token %q (%T), wanted string", t, t)
					}
					c.Output.Type = name(s)
				}
			}
			ret[c.Name] = c
		case m["event"] != nil:
			v := event{}
			if err := json.Unmarshal(def.JSON, &v); err != nil {
				return nil, err
			}
			ret[v.Name] = v
		case m["pragma"] != nil:
			// We ignore pragmas, they're there for the benefit of
			// qemu's own self-validation.
		default:
			return nil, fmt.Errorf("unknown definition kind: %q", string(def.JSON))
		}
	}

	return ret, nil
}

type event struct {
	Name      name        `json:"event"`
	Data      fieldsOrRef `json:"data"`
	BoxedData bool        `json:"boxed"`
}

type command struct {
	Name       name
	Inputs     fieldsOrRef
	Output     field
	BoxedInput bool
}

type alternate struct {
	Name    name          `json:"alternate"`
	Options map[name]name `json:"data"`
}

type flatUnion struct {
	Name          name
	Base          fieldsOrRef
	Discriminator name
	Options       map[name]name
}

func (fu flatUnion) DiscriminatorField(api map[name]interface{}) field {
	for _, f := range fu.Base.Fields(api) {
		if f.Name == fu.Discriminator {
			return f
		}
	}
	panic("no discriminator field found")
}

type simpleUnion struct {
	Name    name
	Options map[name]name
}

type structType struct {
	Name   name   `json:"struct"`
	Fields fields `json:"data"`
	Base   name   `json:"base"`
}

func (s structType) AllFields(api map[name]interface{}) fields {
	var ret fields
	if s.Base != "" {
		ret = append(ret, api[s.Base].(structType).Fields...)
	}
	ret = append(ret, s.Fields...)
	return ret
}

func (s structType) HasInterfaceField(api map[name]interface{}) bool {
	if s.Fields.HasInterfaceField(api) {
		return true
	}
	if s.Base != "" && api[s.Base].(structType).Fields.HasInterfaceField(api) {
		return true
	}
	return false
}

type enum struct {
	Name   name   `json:"enum"`
	Values []name `json:"data"`
}

type field struct {
	Name     name
	List     bool
	Optional bool
	Type     name
}

type fields []field

func (fs fields) HasInterfaceField(api map[name]interface{}) bool {
	for _, f := range fs {
		if f.Type.InterfaceType(api) {
			return true
		}
	}
	return false
}

type fieldsOrRef struct {
	AnonFields fields
	Ref        name
}

func (f fieldsOrRef) Fields(api map[name]interface{}) fields {
	if f.AnonFields != nil {
		return f.AnonFields
	}
	if f.Ref != "" {
		return api[f.Ref].(structType).Fields
	}
	return nil
}

func (f *fieldsOrRef) UnmarshalJSON(bs []byte) error {
	var s name
	if err := json.Unmarshal(bs, &s); err == nil {
		f.AnonFields = nil
		f.Ref = s
		return nil
	}
	return json.Unmarshal(bs, &f.AnonFields)
}

func (fs *fields) UnmarshalJSON(bs []byte) error {
	d := json.NewDecoder(bytes.NewBuffer(bs))

	// Dictionary starts with {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if delim, ok := t.(json.Delim); !ok || delim != '{' {
		return fmt.Errorf("unexpected token %q (%T)", t, t)
	}

	// Then a series of key/value pairs until }.
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}

		// End of the dict?
		if delim, ok := t.(json.Delim); ok && delim == '}' {
			break
		}

		var field field

		// Dictionary key
		s, ok := t.(string)
		if !ok {
			return fmt.Errorf("expected dictionary key (string), got %q (%T)", t, t)
		}
		if s[0] == '*' {
			field.Name = name(s[1:])
			field.Optional = true
		} else {
			field.Name = name(s)
		}

		// The value, either a string or a list of one string.
		t, err = d.Token()
		if err != nil {
			return err
		}

		switch v := t.(type) {
		case string:
			field.Type = name(v)
		case json.Delim:
			if v != '[' {
				return fmt.Errorf("unexpected delimiter %q, wanted a dict value", v)
			}

			// The actual type name
			t, err = d.Token()
			if err != nil {
				return err
			}
			s, ok := t.(string)
			if !ok {
				return fmt.Errorf("unexpected token %q (%T), wanted string", t, t)
			}
			// Closing square bracket
			t, err := d.Token()
			if err != nil {
				return err
			}
			if delim, ok := t.(json.Delim); !ok || delim != ']' {
				return fmt.Errorf("unexpected token %q (%T), wanted ']'", t, t)
			}

			field.Type = name(s)
			field.List = true
		default:
			return fmt.Errorf("unexpected token %q (%T), wanted string or '['", t, t)
		}

		*fs = append(*fs, field)
	}
	return nil
}

type name string

func (n name) Go() string {
	if v, ok := builtinTypes[string(n)]; ok {
		return v
	}

	return n.FieldName()
}

func (n name) FieldName() string {
	out := []string{}
	for _, p := range camelcase.Split(string(n)) {
		if p == "-" || p == "_" {
			continue
		}
		if upperWords[strings.ToLower(p)] {
			p = strings.ToUpper(p)
		} else {
			p = strings.Title(strings.ToLower(p))
		}
		out = append(out, p)

	}

	return strings.Join(out, "")
}

func (n name) SimpleType() bool {
	_, ok := builtinTypes[string(n)]
	return ok
}

func (n name) NullType() bool {
	if n.SimpleType() {
		return false
	}
	return strings.EqualFold(string(n), "Null")
}

func (n name) InterfaceType(api map[name]interface{}) bool {
	if n.SimpleType() {
		return false
	}
	switch api[n].(type) {
	case simpleUnion, flatUnion, alternate:
		return true
	}
	return false
}

var builtinTypes = map[string]string{
	"str":    "string",
	"number": "float64",
	"int":    "int64",
	"int8":   "int8",
	"int16":  "int16",
	"int32":  "int32",
	"int64":  "int64",
	"uint8":  "uint8",
	"uint16": "uint16",
	"uint32": "uint32",
	"uint64": "uint64",
	"size":   "uint64",
	"bool":   "bool",
	"any":    "interface{}",
}

var upperWords = map[string]bool{
	"acpi":    true,
	"acpiost": true,
	"aio":     true,
	"cpu":     true,
	"fd":      true,
	"ftp":     true,
	"ftps":    true,
	"guid":    true,
	"http":    true,
	"https":   true,
	"id":      true,
	"io":      true,
	"ip":      true,
	"json":    true,
	"kvm":     true,
	"luks":    true,
	"mips":    true,
	"nbd":     true,
	"ok":      true,
	"pci":     true,
	"ppc":     true,
	"qmp":     true,
	"ram":     true,
	"sparc":   true,
	"ssh":     true,
	"tcp":     true,
	"tls":     true,
	"tpm":     true,
	"ttl":     true,
	"udp":     true,
	"uri":     true,
	"uuid":    true,
	"vm":      true,
	"vmdk":    true,
	"vnc":     true,
}

// pyToJSON converts one QAPI type definition from its source
// "json-ish python dict" representation to a string parseable by
// encoding/json.
func pyToJSON(py []byte) []byte {
	var (
		advanceTo byte
		drop      bool
	)
	ret := make([]byte, 0, len(py))
	for _, b := range py {
		switch {
		case advanceTo != 0:
			// Inside a quoted string or a comment.
			if b == advanceTo {
				advanceTo = 0
				drop = false
			}
		default:
			switch b {
			case '\'', '"':
				// Start of a quoted string, advance to the end of it
				// without looking for comments.
				advanceTo = b
			case '#':
				// Start of a line comment, skip to EOL
				advanceTo = '\n'
				drop = true
			}
		}
		if !drop {
			if b == '\'' {
				// In JSON, strings are always dquoted.
				b = '"'
			}
			ret = append(ret, b)
		}
	}
	return ret
}

func tryGetVersionFromSpecPath(specPath string) string {
	retVersion := "UNKNOWN"
	verPath, err := resolvePath(specPath, "VERSION")
	if err == nil {
		verBuf, err := getQAPI(verPath)
		if err == nil {
			return string(verBuf)
		}
	}
	return retVersion
}

func resolvePath(orig, new string) (string, error) {
	u, err := url.Parse(orig)
	if err != nil {
		return "", err
	}
	u.Path = filepath.Join(filepath.Dir(u.Path), new)
	return u.String(), nil
}

// getQAPI reads a QMP API spec file, from local disk or over HTTP(S).
func getQAPI(path string) ([]byte, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		return ioutil.ReadFile(path)
	}
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
