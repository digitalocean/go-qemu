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

	"github.com/digitalocean/go-qemu/qmp"
	"github.com/digitalocean/go-qemu/qmp/libvirtrpc"
)

var _ Driver = &RPCDriver{}

// RPCDriver is a QEMU QMP monitor driver which
// communicates via libvirt's RPC interface.
type RPCDriver struct {
	h       *libvirtrpc.Hypervisor
	newConn func() (net.Conn, error)
}

// NewMonitor creates a new qmp.Monitor using libvirt RPC.
func (d *RPCDriver) NewMonitor(domain string) (qmp.Monitor, error) {
	conn, err := d.newConn()
	return libvirtrpc.New(domain, conn), err
}

// DomainNames retrieves all hypervisor domain names using libvirt RPC.
func (d *RPCDriver) DomainNames() ([]string, error) {
	return d.h.DomainNames()
}

// NewRPCDriver configures a RPCDriver.
// The provided newConn function should return an established
// network connection with the target libvirt daemon.
func NewRPCDriver(newConn func() (net.Conn, error)) *RPCDriver {
	return &RPCDriver{
		h:       libvirtrpc.NewHypervisor(newConn),
		newConn: newConn,
	}
}
