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

package libvirtrpc

import (
	"net"

	"github.com/digitalocean/go-libvirt"
)

// Hypervisor represents a running libvirt daemon.
type Hypervisor struct {
	newConn func() (net.Conn, error)
}

// NewHypervisor configures a new Hypervisor used for libvirt daemon
// level tasks such as retrieving a list of domains.
// The provided newConn function should return an established
// network connection with the target libvirt daemon.
func NewHypervisor(newConn func() (net.Conn, error)) *Hypervisor {
	return &Hypervisor{newConn}
}

// DomainNames returns a list of all domains on the hypervisor.
func (h *Hypervisor) DomainNames() ([]string, error) {
	conn, err := h.newConn()
	if err != nil {
		return nil, err
	}

	l := libvirt.New(conn)
	if err := l.Connect(); err != nil {
		conn.Close()
		return nil, err
	}
	defer func() {
		l.Disconnect()
	}()

	domains, err := l.Domains()
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(domains))
	for _, d := range domains {
		names = append(names, d.Name)
	}

	return names, nil
}

// Version returns the version of the libvirt daemon.
func (h *Hypervisor) Version() (string, error) {
	conn, err := h.newConn()
	if err != nil {
		return "", err
	}

	l := libvirt.New(conn)
	if err := l.Connect(); err != nil {
		conn.Close()
		return "", err
	}
	defer func() {
		l.Disconnect()
	}()

	return l.Version()
}
