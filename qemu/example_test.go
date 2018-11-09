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

package qemu_test

import (
	"fmt"
	"log"

	"github.com/digitalocean/go-qemu/qemu"
	"github.com/digitalocean/go-qemu/qmp"
	"github.com/digitalocean/go-qemu/qmp/qmptest"
)

// This example demonstrates how to use qemu.NewDomain with a qmp.Monitor to
// perform actions on a Domain.
//
// Typically, these actions would be performed using an actual monitor of type
// qmp.LibvirtRPCMonitor or qmp.SocketMonitor, instead of using
// qmptest.NewMonitor.
func ExampleNewDomain() {
	// Use qmptest.NewMonitor to create an "example" qmp.Monitor, that returns
	// mock data from another function.
	//
	// Normally, qmp.NewLibvirtRPCMonitor or qmp.NewSocketMonitor would be used
	// here instead.
	mon := qmptest.NewMonitor(func(cmd qmp.Command) (interface{}, error) {
		return exampleRun(cmd)
	})

	// Monitor must be connected before it can be used.
	if err := mon.Connect(); err != nil {
		log.Fatalf("failed to connect monitor: %v", err)
	}

	// Wrap monitor in a qemu.Domain to perform higher-level management actions.
	dom, err := qemu.NewDomain(mon, "example")
	if err != nil {
		log.Fatalf("failed to create domain: %v", err)
	}

	// Return the domain's QEMU version.
	version, err := dom.Version()
	if err != nil {
		log.Fatalf("failed to check domain QEMU version: %v", err)
	}

	// List the available QMP commands supported by the domain.
	commands, err := dom.Commands()
	if err != nil {
		log.Fatalf("failed to list domain commands: %v", err)
	}

	// Print some information about the domain.
	fmt.Printf("%s - QEMU %s\n", dom.Name, version)
	fmt.Println("  - commands:")
	for _, c := range commands {
		fmt.Printf("    - %s\n", c)
	}

	// Close the domain to clean up its resources and underlying monitor.
	if err := dom.Close(); err != nil {
		log.Fatalf("failed to close domain: %v", err)
	}

	// Output:
	// example - QEMU 2.0.0
	//   - commands:
	//     - query-block
	//     - query-commands
	//     - query-version
}

func exampleRun(cmd qmp.Command) (interface{}, error) {
	switch cmd.Execute {
	case "query-commands":
		return runQueryCommands(), nil
	case "query-version":
		return runQueryVersion(), nil
	}

	return nil, fmt.Errorf("unknown command: %q", cmd.Execute)
}

func runQueryCommands() interface{} {
	var response struct {
		ID     string        `json:"id"`
		Return []nameWrapper `json:"return"`
	}

	commands := []string{
		"query-block",
		"query-commands",
		"query-version",
	}

	response.Return = make([]nameWrapper, 0, len(commands))
	for _, c := range commands {
		response.Return = append(response.Return, nameWrapper{
			Name: c,
		})
	}

	return response
}

func runQueryVersion() interface{} {
	var response struct {
		ID     string      `json:"id"`
		Return qmp.Version `json:"return"`
	}
	response.Return.QEMU.Major = 2

	return response
}

type nameWrapper struct {
	Name string `json:"name"`
}
