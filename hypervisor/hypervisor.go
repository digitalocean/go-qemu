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

// Package hypervisor provides management facilities for one or more QEMU
// virtual machines on a hypervisor.
package hypervisor

import (
	"fmt"

	"github.com/digitalocean/go-qemu/qemu"
	"github.com/digitalocean/go-qemu/qmp"
)

// A Hypervisor enables access to all virtual machines on a QEMU hypervisor.
type Hypervisor struct {
	// Driver used to communicate with domains
	driver Driver
}

// New creates a new Hypervisor using the input Driver.
func New(driver Driver) *Hypervisor {
	return &Hypervisor{
		driver: driver,
	}
}

// A Driver is a QEMU QMP monitor driver that a Hypervisor can use to perform
// actions on groups of virtual machines.
type Driver interface {
	NewMonitor(domain string) (qmp.Monitor, error)
	DomainNames() ([]string, error)
}

// A Versioner is a Driver that is able to report its version on the
// Hypervisor.
type Versioner interface {
	Version() (string, error)
}

// Domains retrieves all virtual machines which reside on a given hypervisor,
// connecting to the monitor socket of each machine.
//
// Each Domain's Close method must be called to clean up its resources when
// they are no longer needed.
func (hv *Hypervisor) Domains() ([]*qemu.Domain, error) {
	domains, err := hv.driver.DomainNames()
	if err != nil {
		return nil, err
	}

	ds := make([]*qemu.Domain, 0, len(domains))
	for _, d := range domains {
		domain, err := hv.connectDomain(d)
		if err != nil {
			return nil, err
		}

		ds = append(ds, domain)
	}

	return ds, nil
}

// Domain retrieves a single virtual machine from a hypervisor, connecting to
// its monitor socket.  If no virtual machine exists with the given name, an
// error is returned.
//
// The Domain's Close method must be called to clean up its resources when
// they are no longer needed.
func (hv *Hypervisor) Domain(name string) (*qemu.Domain, error) {
	domains, err := hv.driver.DomainNames()
	if err != nil {
		return nil, err
	}

	for _, d := range domains {
		if d == name {
			return hv.connectDomain(name)
		}
	}

	return nil, fmt.Errorf("domain %q not found", name)
}

// DomainNames retrieves the names of all virtual machines which reside on
// a given hypervisor, so that individual connections can be opened using
// the Domain method, as needed.
func (hv *Hypervisor) DomainNames() ([]string, error) {
	return hv.driver.DomainNames()
}

// connectDomain opens a monitor socket connection and creates a virtual
// machine for the machine with the specified name.
func (hv *Hypervisor) connectDomain(name string) (*qemu.Domain, error) {
	mon, err := hv.driver.NewMonitor(name)
	if err != nil {
		return nil, err
	}

	if err := mon.Connect(); err != nil {
		return nil, err
	}

	return qemu.NewDomain(mon, name)
}
