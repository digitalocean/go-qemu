//+build ignore

package main

import (
	"fmt"
	"os"

	gen "github.com/digitalocean/go-qemu/internal/qmp-gen"
)

func main() {
	if err := gen.Generate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
