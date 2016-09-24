package gen

import (
	"bytes"
	"errors"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

// neededTypes returns the subset of API that should be rendered to file.
//
// The subset to render is defined as the transitive closure of all types
// reachable from commands and events.
func neededTypes(api map[name]interface{}) map[name]interface{} {
	need := map[name]interface{}{}

	for n, i := range api {
		switch i.(type) {
		case command:
			need[n] = i
		case event:
			need[n] = i
		}
	}

	for {
		nextNeed := expandTypes(api, need)
		if len(need) == len(nextNeed) {
			break
		}
		need = nextNeed
	}

	return need
}

// expandTypes returns the union of need and all types it directly references.
func expandTypes(api, need map[name]interface{}) map[name]interface{} {
	ret := map[name]interface{}{}

	for n, i := range need {
		ret[n] = i
		switch v := i.(type) {
		case command:
			addName(api, ret, v.Name)
			if v.BoxedInput {
				addName(api, ret, v.Inputs.Ref)
			} else {
				for _, f := range v.Inputs.Fields(api) {
					addName(api, ret, f.Type)
				}
			}
			if v.Output.Type != "" {
				addName(api, ret, v.Output.Type)
			}
		case event:
			addName(api, ret, v.Name)
			if v.BoxedData {
				addName(api, ret, v.Data.Ref)
			} else {
				for _, f := range v.Data.Fields(api) {
					addName(api, ret, f.Type)
				}
			}
		case struct_:
			addName(api, ret, v.Name)
			for _, f := range v.Fields {
				addName(api, ret, f.Type)
			}
			if v.Base != "" {
				for _, f := range api[v.Base].(struct_).Fields {
					addName(api, ret, f.Type)
				}
			}
		case simpleUnion:
			addName(api, ret, v.Name)
			for _, typ := range v.Options {
				addName(api, ret, typ)
			}
		case flatUnion:
			addName(api, ret, v.Name)
			for _, f := range v.Base.Fields(api) {
				addName(api, ret, f.Type)
			}
			for _, typ := range v.Options {
				for _, f := range api[typ].(struct_).Fields {
					addName(api, ret, f.Type)
				}
			}
		case alternate:
			addName(api, ret, v.Name)
			for _, typ := range v.Options {
				addName(api, ret, typ)
			}
		case enum:
			addName(api, ret, v.Name)
		}
	}
	return ret
}

func addName(api, need map[name]interface{}, name name) {
	if name.SimpleType() {
		return
	}
	v, ok := api[name]
	if ok {
		need[name] = v
	}
}

// renderAPI generates the Go implementation for the types listed in need.
func renderAPI(templateDir string, api, need map[name]interface{}) ([]byte, error) {
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{
		"API": func() map[name]interface{} { return api },
		"typeOf": func(v interface{}) string {
			if v == nil {
				return "nil"
			}
			return strings.ToLower(reflect.TypeOf(v).Name())
		},
		"unexport": func(s string) string {
			ret := strings.ToLower(s[:1]) + s[1:]
			if ret == "type" {
				return "typ"
			}
			return ret
		},
		"render": func(v interface{}) (string, error) {
			n := strings.Replace(strings.ToLower(reflect.TypeOf(v).Name()), "_", "", -1)
			var b bytes.Buffer
			err := tmpl.ExecuteTemplate(&b, n, v)
			return b.String(), err
		},
		"abort": func(s string) (string, error) { return "", errors.New(s) },
	})
	tmpl, err := tmpl.ParseGlob(filepath.Join(templateDir, "*"))
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if err := tmpl.ExecuteTemplate(&out, "main", need); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
