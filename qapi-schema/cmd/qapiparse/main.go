package main

import (
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/repr"

	qapischema "github.com/digitalocean/go-qemu/qapi-schema"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
	}
}

func run() error {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	schema, err := qapischema.Parse(string(input))
	if err != nil {
		return err
	}

	repr.Println(schema)
	return nil
}
