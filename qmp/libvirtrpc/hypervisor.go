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
	"bytes"
	"fmt"
	"net"

	"github.com/davecgh/go-xdr/xdr2"
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

	rpc := setup(conn)
	if err := rpc.Connect(); err != nil {
		conn.Close()
		return nil, err
	}
	defer func() {
		rpc.Disconnect()
		conn.Close()
	}()

	// these are the flags as passed by `virsh`, defined in:
	// src/remote/remote_protocol.x # remote_connect_list_all_domains_args
	req := struct {
		NeedResults uint32
		Flags       uint32
	}{
		NeedResults: 1,
		Flags:       3,
	}

	buf, err := encode(&req)
	if err != nil {
		return nil, err
	}

	resp, err := rpc.request(procConnectListAllDomains, programRemote, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Status != StatusOK {
		return nil, decodeError(r.Payload)
	}

	result := struct {
		Domains []domain
		Count   uint32
	}{}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}

	domains := make([]string, 0, result.Count)
	for _, d := range result.Domains {
		domains = append(domains, d.Name)
	}

	return domains, nil
}

// Version returns the version of the libvirt daemon.
func (h *Hypervisor) Version() (string, error) {
	conn, err := h.newConn()
	if err != nil {
		return "", err
	}

	rpc := setup(conn)
	if err := rpc.Connect(); err != nil {
		conn.Close()
		return "", err
	}
	defer func() {
		rpc.Disconnect()
		conn.Close()
	}()

	resp, err := rpc.request(procConnectGetLibVersion, programRemote, nil)
	if err != nil {
		return "", err
	}

	r := <-resp
	if r.Status != StatusOK {
		return "", decodeError(r.Payload)
	}

	result := struct {
		Version uint64
	}{}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))
	_, err = dec.Decode(&result)
	if err != nil {
		return "", err
	}

	// The version is provided as an int following this formula:
	// version * 1,000,000 + minor * 1000 + micro
	// See src/libvirt-host.c # virConnectGetLibVersion
	version := result.Version
	major := version / 1000000
	version %= 1000000
	minor := version / 1000
	version %= 1000
	micro := version

	versionString := fmt.Sprintf("%d.%d.%d", major, minor, micro)
	return versionString, nil
}
