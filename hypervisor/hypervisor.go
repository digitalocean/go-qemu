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
	"net/url"
	"strings"
	"sync"

	"github.com/digitalocean/go-qemu"
	"github.com/digitalocean/go-qemu/qmp"
)

// A Hypervisor enables access to all virtual machines on a QEMU hypervisor.
type Hypervisor interface {
	Domains() ([]*qemu.Domain, error)
	Domain(name string) (*qemu.Domain, error)
	DomainNames() ([]string, error)
	Disconnect() error
}

var _ Hypervisor = &Libvirt{}

// Libvirt represents a libvirt hypervisor.
type Libvirt struct {
	// URI of hypervisor
	url *url.URL

	// Currently connected monitor sockets
	mu        sync.Mutex
	connected []qmp.Monitor

	// Functions which can be swapped for easy testing
	newMonitor  monitorFunc
	domainNames domainNamesFunc
}

// A monitorFunc is a function which can create a qmp.Monitor.
type monitorFunc func(uri string, domain string) (qmp.Monitor, error)

// A domainNamesFunc is a function which can return a list of
// domain names.
type domainNamesFunc func() ([]string, error)

// NewLibvirt configures a connection to the provided hypervisor.
//
// Hypervisor URIs may be local or remote, e.g.,
//	qemu:///system
//	qemu+ssh://libvirt@example.com/system
func NewLibvirt(uri string) (*Libvirt, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	hv := &Libvirt{
		url:       u,
		connected: make([]qmp.Monitor, 0),

		// Default to creating libvirt monitors and using virsh to
		// retrieve domain names.
		newMonitor: func(uri string, domain string) (qmp.Monitor, error) {
			return qmp.NewLibvirtMonitor(uri, domain)
		},
		domainNames: func() ([]string, error) {
			return virshList(uri)
		},
	}

	return hv, nil
}

// Domains retrieves all virtual machines which reside on a given hypervisor,
// connecting to the monitor socket of each machine.
//
// Each Domain's Close method must be called to clean up its resources when
// they are no longer needed.
//
// The Disconnect method must be called to clean up monitor socket connections
// once the virtual machines are no longer needed.
func (hv *Libvirt) Domains() ([]*qemu.Domain, error) {
	domains, err := hv.domainNames()
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
//
// The Disconnect method must be called to clean up the monitor socket
// connection once the virtual machine is no longer needed.
func (hv *Libvirt) Domain(name string) (*qemu.Domain, error) {
	domains, err := hv.domainNames()
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
func (hv *Libvirt) DomainNames() ([]string, error) {
	return hv.domainNames()
}

// Disconnect cleans up monitor socket connections for all virtual
// machines.
func (hv *Libvirt) Disconnect() error {
	hv.mu.Lock()
	defer hv.mu.Unlock()

	for _, mon := range hv.connected {
		if err := mon.Disconnect(); err != nil {
			return err
		}
	}

	hv.connected = make([]qmp.Monitor, 0)
	return nil
}

// connectDomain opens a monitor socket connection and creates a virtual
// machine for the machine with the specified name.
func (hv *Libvirt) connectDomain(name string) (*qemu.Domain, error) {
	mon, err := hv.newMonitor(hv.url.String(), name)
	if err != nil {
		return nil, err
	}

	dom, err := qemu.NewDomain(mon, name)
	if err != nil {
		return nil, err
	}

	hv.mu.Lock()
	hv.connected = append(hv.connected, mon)
	hv.mu.Unlock()

	return dom, nil
}

// virshList shells out to 'virsh list --all --name' to produce a list of domain names.
func virshList(uri string) ([]string, error) {
	out, err := qmp.Virsh(uri, "list", "--all", "--name")
	if err != nil {
		return nil, err
	}

	// Remove blank newline at end of domains list
	domains := strings.Split(string(out), "\n")
	domains = domains[:len(domains)-2]
	return domains, nil
}
