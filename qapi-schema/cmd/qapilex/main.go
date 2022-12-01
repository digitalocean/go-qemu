package main

import (
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/repr"

	"github.com/digitalocean/go-qemu/qapi-schema/internal/lex"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func run() error {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	tokens := lex.Lex(string(input))
	repr.Println(tokens)

	return nil
}
