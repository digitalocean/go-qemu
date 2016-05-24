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

// Package libvirtrpc is a pure Go implementation of the libvirt RPC protocol.
// For more information on the protocol, see https://libvirt.org/internals/rpc.html
package libvirtrpc

import (
	"encoding/json"
	"net"

	"github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-qemu/qmp"
)

var _ qmp.Monitor = &Monitor{}

// Monitor implements LibVirt's remote procedure call protocol.
type Monitor struct {
	l *libvirt.Libvirt
	// Domain name as seen by libvirt, e.g., stage-lb-1
	Domain string
}

// New configures a new RPC Monitor connection.
// The provided domain should be the name of the domain as seen
// by libvirt, e.g., stage-lb-1.
func New(domain string, conn net.Conn) *Monitor {
	l := libvirt.New(conn)

	return &Monitor{
		l:      l,
		Domain: domain,
	}
}

// Connect establishes communication with the libvirt server.
// The underlying libvirt socket connection must be previously established.
func (rpc *Monitor) Connect() error {
	return rpc.l.Connect()
}

// Disconnect shuts down communication with the libvirt server
// and closes the underlying net.Conn.
func (rpc *Monitor) Disconnect() error {
	return rpc.l.Disconnect()
}

// Events streams QEMU QMP Events.
// If a problem is encountered setting up the event monitor connection
// an error will be returned. Errors encountered during streaming will
// cause the returned event channel to be closed.
func (rpc *Monitor) Events() (<-chan qmp.Event, error) {
	events, err := rpc.l.Events(rpc.Domain)
	if err != nil {
		return nil, err
	}

	c := make(chan qmp.Event)
	go func() {
		// process events
		for e := range events {
			qe, err := qmpEvent(&e)
			if err != nil {
				close(c)
				break
			}

			c <- *qe
		}
	}()

	return c, nil
}

// Run executes the given QAPI command against a domain's QEMU instance.
// For a list of available QAPI commands, see:
//	http://git.qemu.org/?p=qemu.git;a=blob;f=qapi-schema.json;hb=HEAD
func (rpc *Monitor) Run(cmd []byte) ([]byte, error) {
	return rpc.l.Run(rpc.Domain, cmd)
}

// qmpEvent takes a libvirt DomainEvent and returns the QMP equivalent.
func qmpEvent(e *libvirt.DomainEvent) (*qmp.Event, error) {
	var qe qmp.Event

	if e.Details != nil {
		if err := json.Unmarshal(e.Details, &qe.Data); err != nil {
			return nil, err
		}
	}

	qe.Event = e.Event
	qe.Timestamp.Seconds = int64(e.Seconds)
	qe.Timestamp.Microseconds = int64(e.Microseconds)

	return &qe, nil
}
