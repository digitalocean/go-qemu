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
	"fmt"
	"time"

	"github.com/digitalocean/go-qemu/qmp"
)

var _ Driver = &SocketDriver{}

// A SocketDriver is a QEMU QMP monitor driver which communicates directly
// with a QEMU monitor socket.
type SocketDriver struct {
	addrs []SocketAddress
}

// A SocketAddress is a QEMU QMP monitor socket address used to configure
// a SocketDriver.
type SocketAddress struct {
	// Network should be one of "unix", "tcp", etc.
	Network string

	// Address should be a host:port address or UNIX socket filepath,
	// such "8.8.8.8:4444" or "/var/lib/qemu/example.socket".
	Address string

	// Timeout specifies how long the monitor socket will attempt to be
	// reached before timing out.
	Timeout time.Duration
}

// NewMonitor creates a new qmp.Monitor using a QEMU monitor socket at
// the specified address.
func (d *SocketDriver) NewMonitor(addr string) (qmp.Monitor, error) {
	for _, a := range d.addrs {
		if a.Address == addr {
			return qmp.NewSocketMonitor(
				a.Network,
				a.Address,
				a.Timeout,
			)
		}
	}

	return nil, fmt.Errorf("address not known to SocketDriver: %q", addr)
}

// DomainNames retrieves all hypervisor domain names known to the SocketDriver.
func (d *SocketDriver) DomainNames() ([]string, error) {
	names := make([]string, 0, len(d.addrs))
	for _, a := range d.addrs {
		names = append(names, a.Address)
	}

	return names, nil
}

// NewSocketDriver configures a SocketDriver using one or more SocketAddress
// structures for configuration.
func NewSocketDriver(addrs []SocketAddress) *SocketDriver {
	return &SocketDriver{
		addrs: addrs,
	}
}
