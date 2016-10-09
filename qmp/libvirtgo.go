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

package qmp

import "net/url"

// LibvirtGoMonitor is a monitor that wraps the libvirt-go api to
// communicate with a QEMU Machine Protocol (QMP) socket.
// Communication is proxied via the libvirtd daemon. Multiple
// connections to the same hypervisor and domain are permitted.
type LibvirtGoMonitor struct {
	domain     string
	url        *url.URL
}

// Connect - Needs better comments
func (mon LibvirtGoMonitor) Connect() error {
	return nil
}

// Disconnect tears down open QMP socket connections.
func (mon *LibvirtGoMonitor) Disconnect() error {
	
}

// Run executes the given QAPI command against a domain's QEMU instance.
// For a list of available QAPI commands, see:
//	http://git.qemu.org/?p=qemu.git;a=blob;f=qapi-schema.json;hb=HEAD
func (mon LibvirtGoMonitor) Run(cmd []byte) ([]byte, error) {
    return []byte[], nil
}

// Events streams QEMU QMP Events.
// If a problem is encountered setting up the event monitor connection
// an error will be returned. Errors encountered during streaming will
// cause the returned event channel to be closed.
func (mon *LibvirtGoMonitor) Events() (<-chan Event, error) {
    return nil, nil
}

// NewLibvirtGoMonitor configures a connection to the provided hypervisor and domain.
// An error is returned if the provided libvirt connection URI is invalid.
//
// Hypervisor URIs may be local or remote, e.g.,
//	qemu:///system
//	qemu+ssh://libvirt@example.com/system
func NewLibvirtGoMonitor(uri, domain string) (*LibvirtGoMonitor, error) {
    return nil, nil
}