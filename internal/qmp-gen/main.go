package gen

import (
	"bytes"
	"encoding/json"
	"flag"
	"go/format"
	"io/ioutil"
)

const specURL = `https://raw.githubusercontent.com/danderson/qemu/master/qapi-schema.json`

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
	if err = ioutil.WriteFile(*outputGo, formatted, 0640); err != nil {
		return err
	}

	return nil
}
