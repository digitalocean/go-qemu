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

import libvirt "github.com/rgbkrk/libvirt-go"

type libvirtGoMonitorLinux struct {
	LibvirtGoMonitor
	domain  string
	uri     string
	virConn *libvirt.VirConnection
}

// Connect  sets up QEMU QMP connection via libvirt using
// the libvirt-go package.
// An error is returned if the libvirtd daemon is unreachable.
func (mon libvirtGoMonitorLinux) Connect() error {
	virConn, err := libvirt.NewVirConnection(mon.uri)
	if err == nil {
		mon.virConn = &virConn
	}
	return err
}

// Disconnect tears down open QMP socket connections.
func (mon *libvirtGoMonitorLinux) Disconnect() error {
	var err error
	if mon.virConn != nil {
		_, err = mon.virConn.CloseConnection()
		mon.virConn = nil
	}
	return err
}

// Run executes the given QAPI command against a domain's QEMU instance.
// For a list of available QAPI commands, see:
//	http://git.qemu.org/?p=qemu.git;a=blob;f=qapi-schema.json;hb=HEAD
func (mon libvirtGoMonitorLinux) Run(cmd []byte) ([]byte, error) {
	domain, err := mon.virConn.LookupDomainByName(mon.domain)
	if err != nil {
		return nil, err
	}

	result, err := domain.QemuMonitorCommand(
		libvirt.VIR_DOMAIN_QEMU_MONITOR_COMMAND_DEFAULT,
		string(cmd))
	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}

// Events streams QEMU QMP Events.
// If a problem is encountered setting up the event monitor connection
// an error will be returned. Errors encountered during streaming will
// cause the returned event channel to be closed.
func (mon *libvirtGoMonitorLinux) Events() (<-chan Event, error) {
	return nil, nil
}

// NewLibvirtGoMonitor configures a connection to the provided hypervisor
// and domain.
// An error is returned if the provided libvirt connection URI is invalid.
//
// Hypervisor URIs may be local or remote, e.g.,
//	qemu:///system
//	qemu+ssh://libvirt@example.com/system
func NewLibvirtGoMonitor(uri, domain string) Monitor {
	return &libvirtGoMonitorLinux{
		uri:    uri,
		domain: domain,
	}
}
