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

package hypervisor

import (
	"net"

	"github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-qemu/qmp"
)

var _ Driver = &RPCDriver{}
var _ Versioner = &RPCDriver{}

// RPCDriver is a QEMU QMP monitor driver which
// communicates via libvirt's RPC interface.
type RPCDriver struct {
	newConn func() (net.Conn, error)
}

// DomainNames retrieves all hypervisor domain names using libvirt RPC.
func (d *RPCDriver) DomainNames() ([]string, error) {
	var names []string
	err := d.withLibvirt(func(l *libvirt.Libvirt) error {
		domains, err := l.Domains()
		if err != nil {
			return err
		}

		names = make([]string, 0, len(domains))
		for _, d := range domains {
			names = append(names, d.Name)
		}

		return nil
	})

	return names, err
}

// NewMonitor creates a new qmp.Monitor using libvirt RPC.
func (d *RPCDriver) NewMonitor(domain string) (qmp.Monitor, error) {
	conn, err := d.newConn()
	return qmp.NewLibvirtRPCMonitor(domain, conn), err
}

// Version returns the version string for the libvirt daemon.
func (d *RPCDriver) Version() (string, error) {
	var version string
	err := d.withLibvirt(func(l *libvirt.Libvirt) error {
		v, err := l.Version()
		version = v
		return err
	})

	return version, err
}

// NewRPCDriver configures a hypervisor driver using Libvirt RPC.
// The provided newConn function should return an established
// network connection with the target libvirt daemon.
func NewRPCDriver(newConn func() (net.Conn, error)) *RPCDriver {
	return &RPCDriver{
		newConn: newConn,
	}
}

// withLibvirt opens a new RPC connection to Libvirt and invokes the input
// closure using the RPC connection.  Connections are automatically cleaned
// up when the close returns.
func (d *RPCDriver) withLibvirt(fn func(l *libvirt.Libvirt) error) error {
	conn, err := d.newConn()
	if err != nil {
		return err
	}

	l := libvirt.New(conn)
	if err := l.Connect(); err != nil {
		_ = conn.Close()
		return err
	}
	defer func() {
		_ = l.Disconnect()
	}()

	return fn(l)
}
