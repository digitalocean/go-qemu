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
	virConn *libvirt.VirConnection
}

func (mon libvirtGoMonitorLinux) connect() error {
	virConn, err := libvirt.NewVirConnection(mon.uri)
	if err == nil {
		mon.virConn = &virConn
	}
	return err
}

func (mon *libvirtGoMonitorLinux) disconnect() error {
	var err error
	if mon.virConn != nil {
		_, err = mon.virConn.CloseConnection()
		mon.virConn = nil
	}
	return err
}

func (mon libvirtGoMonitorLinux) run(cmd []byte) ([]byte, error) {
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

func (mon *libvirtGoMonitorLinux) events() (<-chan Event, error) {
	return nil, nil
}

func newLibvirtGoMonitor(uri, domain string) Monitor {
	return &libvirtGoMonitorLinux{
		uri:    uri,
		domain: domain,
	}
}
