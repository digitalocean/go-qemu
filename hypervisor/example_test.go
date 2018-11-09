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

package hypervisor_test

import (
	"fmt"
	"log"

	"github.com/digitalocean/go-qemu/hypervisor"
	"github.com/digitalocean/go-qemu/qmp"
	"github.com/digitalocean/go-qemu/qmp/qmptest"
)

func ExampleNew() {
	// Create a hypervisor.Hypervisor using an in-memory hypervisor.Driver,
	// that returns a fixed list of three domains.
	hv := hypervisor.New(newDriver([]string{
		"foo",
		"bar",
		"baz",
	}))

	// Connect to the monitor socket of each domain using the in-memory driver.
	domains, err := hv.Domains()
	if err != nil {
		log.Fatalf("failed to connect to domains: %v", err)
	}

	fmt.Println("hypervisor:")
	for _, d := range domains {
		// Return the domain's QEMU version.
		version, err := d.Version()
		if err != nil {
			log.Fatalf("failed to check domain QEMU version: %v", err)
		}

		// Print some information about the domain.
		fmt.Printf("  - %s - QEMU %s\n", d.Name, version)

		// Close the domain to clean up its resources and underlying monitor.
		if err := d.Close(); err != nil {
			log.Fatalf("failed to close domain: %v", err)
		}
	}

	// Output:
	// hypervisor:
	//   - foo - QEMU 2.0.0
	//   - bar - QEMU 2.0.1
	//   - baz - QEMU 2.0.2
}

var _ hypervisor.Driver = &driver{}

func newDriver(domains []string) *driver {
	return &driver{
		domains: domains,
	}
}

type driver struct {
	domains []string
}

func (d *driver) NewMonitor(domain string) (qmp.Monitor, error) {
	m := qmptest.NewMonitor(func(cmd qmp.Command) (interface{}, error) {
		return run(domain, cmd)
	})

	return m, nil
}

func (d *driver) DomainNames() ([]string, error) {
	return d.domains, nil
}

func run(domain string, cmd qmp.Command) (interface{}, error) {
	switch cmd.Execute {
	case "query-version":
		return runQueryVersion(domain), nil
	}

	return nil, fmt.Errorf("unknown command: %q", cmd.Execute)
}

func runQueryVersion(domain string) interface{} {
	var response struct {
		ID     string      `json:"id"`
		Return qmp.Version `json:"return"`
	}
	response.Return.QEMU.Major = 2

	switch domain {
	case "bar":
		response.Return.QEMU.Micro = 1
	case "baz":
		response.Return.QEMU.Micro = 2
	}

	return response
}
