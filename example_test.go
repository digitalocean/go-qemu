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
	"encoding/json"
	"fmt"
	"log"

	"github.com/digitalocean/go-qemu"
	"github.com/digitalocean/go-qemu/qmp"
)

// This example demonstrates how to use qemu.NewDomain with a qmp.Monitor to
// perform actions on a Domain.
//
// Typically, these actions would be performed using an actual monitor of type
// qmp.Libvirt or qmp.Socket, instead of using newExampleMonitor.
func ExampleNewDomain() {
	// Create a new "example" qmp.Monitor.  Normally, qmp.NewLibvirtMonitor
	// or qmp.NewSocketMonitor would be used here instead.
	mon := newExampleMonitor()

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

	// Close the domain to clean up its resources, and disconnect the
	// underlying monitor.
	if err := dom.Close(); err != nil {
		log.Fatalf("failed to close domain: %v", err)
	}
	if err := mon.Disconnect(); err != nil {
		log.Fatalf("failed to disconnect monitor: %v", err)
	}

	// Output:
	// example - QEMU 2.0.0
	//   - commands:
	//     - query-block
	//     - query-commands
	//     - query-version
}

func newExampleMonitor() qmp.Monitor {
	return &exampleMonitor{}
}

var _ qmp.Monitor = &exampleMonitor{}

type exampleMonitor struct{}

func (mon *exampleMonitor) Connect() error                    { return nil }
func (mon *exampleMonitor) Disconnect() error                 { return nil }
func (mon *exampleMonitor) Events() (<-chan qmp.Event, error) { return nil, nil }

func (mon *exampleMonitor) Run(command []byte) ([]byte, error) {
	var cmd qmp.Cmd
	if err := json.Unmarshal(command, &cmd); err != nil {
		return nil, err
	}

	switch cmd.Execute {
	case "query-commands":
		return mon.runQueryCommands()
	case "query-version":
		return mon.runQueryVersion()
	}

	return nil, fmt.Errorf("unknown command: %q", cmd.Execute)
}

func (mon *exampleMonitor) runQueryCommands() ([]byte, error) {
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

	return json.Marshal(response)
}

func (mon *exampleMonitor) runQueryVersion() ([]byte, error) {

	var response struct {
		ID     string      `json:"id"`
		Return qmp.Version `json:"return"`
	}
	response.Return.QEMU.Major = 2

	return json.Marshal(response)
}

type nameWrapper struct {
	Name string `json:"name"`
}
