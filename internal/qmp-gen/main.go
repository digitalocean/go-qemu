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
	"flag"
	"go/format"
	"io/ioutil"
)

const specURL = `https://raw.githubusercontent.com/qemu/qemu/stable-2.11/qapi-schema.json`

var (
	inputSpec  = flag.String("input", specURL, "Input spec")
	outputGo   = flag.String("output", "", "Generated code output")
	outputSpec = flag.String("spec", "", "Generated spec output")
	templates  = flag.String("templates", "", "Template directory")
)

// Generate parses a QMP API spec and generates a Go implementation.
func Generate() error {
	flag.Parse()

	defs, err := readDefinitions(*inputSpec)
	if err != nil {
		return err
	}

	// Write out the concatenated spec.
	var spec bytes.Buffer

	// First add a comment with the best guess of version
	spec.WriteString("\n##QEMU SPECIFICATION VERSION:  ")
	spec.WriteString(tryGetVersionFromSpecPath(*inputSpec))
	spec.WriteString("\n\n")

	for _, def := range defs {
		spec.WriteString(def.Docstring)
		spec.WriteByte('\n')
		json.Indent(&spec, def.JSON, "", "  ")
		spec.WriteString("\n\n")
	}
	if err = ioutil.WriteFile(*outputSpec, spec.Bytes(), 0640); err != nil {
		return err
	}
	// Generate and write out the code.
	symbols, err := parse(defs)
	if err != nil {
		return err
	}
	need := neededTypes(symbols)
	bs, err := renderAPI(*templates, symbols, need)
	if err != nil {
		return err
	}
	formatted, err := format.Source(bs)
	if err != nil {
		// Ignore the error, we're just trying to write out the
		// unformatted code to help debugging.
		ioutil.WriteFile(*outputGo, bs, 0640)
		return err
	}

	return ioutil.WriteFile(*outputGo, formatted, 0640)
}
