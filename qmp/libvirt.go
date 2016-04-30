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

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/digitalocean/go-qemu/internal/shellexec"
)

// eventFormat represents QEMU event JSON formatting.
// This is used to normalize mangled events passed through libvirt.
const eventFormat = `{ "event": %q, "data": %s, "timestamp": { "seconds": %s, "microseconds": %s } }`

// used to parse libvirt's managled QMP event structure
var extractor = regexp.MustCompile(`event (\w+) at (\d+)\.(\d+) for domain .*: (.*)$`)

type response struct {
	ID     string      `json:"id"`
	Return interface{} `json:"return,omitempty"`
	Error  struct {
		Class string `json:"class"`
		Desc  string `json:"desc"`
	} `json:"error,omitempty"`
}

func (r *response) Err() error {
	if r.Error.Desc == "" {
		return nil
	}

	return errors.New(r.Error.Desc)
}

// Libvirt represents a QEMU Machine Protocol (QMP) socket.
// Communication is proxied via the libvirtd daemon. Multiple
// connections to the same hypervisor and domain are permitted.
type Libvirt struct {
	domain     string
	url        *url.URL
	disconnect chan error

	cmd shellexec.Command
}

// Connect sets up a QEMU QMP connection via libvirt's QEMU monitor socket.
// An error is returned if the libvirt daemon is unreachable.
func (mon Libvirt) Connect() error {
	_, err := Virsh(mon.cmd, mon.url.String(), "connect")
	return err
}

// Disconnect tears down open QMP socket connections.
func (mon *Libvirt) Disconnect() error {
	mon.disconnect <- nil

	select {
	case err := <-mon.disconnect:
		return err
	case <-time.After(5 * time.Second):
		return errors.New("timeout disconnecting from monitor")
	}
}

// Run executes the given QAPI command against a domain's QEMU instance.
// For a list of available QAPI commands, see:
//	http://git.qemu.org/?p=qemu.git;a=blob;f=qapi-schema.json;hb=HEAD
func (mon Libvirt) Run(cmd []byte) ([]byte, error) {
	raw, err := Virsh(
		mon.cmd,
		mon.url.String(),
		"qemu-monitor-command",
		mon.domain,
		string(cmd),
	)
	if err != nil {
		return raw, err
	}

	// check for QEMU errors
	var r response
	if err = json.Unmarshal(raw, &r); err != nil {
		return raw, err
	}
	if err := r.Err(); err != nil {
		return raw, err
	}

	return raw, err
}

// Events streams QEMU QMP Events.
// If a problem is encountered setting up the event monitor connection
// an error will be returned. Errors encountered during streaming will
// cause the returned event channel to be closed.
func (mon *Libvirt) Events() (<-chan Event, error) {
	stream := make(chan Event)
	cmd := mon.cmd.Prepare(
		"virsh",
		"-c",
		mon.url.String(),
		"qemu-monitor-event",
		"--loop",
		mon.domain,
	)

	output, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	go func() {
		<-mon.disconnect

		if err := cmd.Kill(); err != nil {
			mon.disconnect <- err
			return
		}

		// Must wait after killing process to avoid creating zombie processes
		mon.disconnect <- cmd.Wait()
	}()

	go streamEvents(output, stream)

	return stream, nil
}

// NewLibvirtMonitor configures a connection to the provided hypervisor and domain.
// An error is returned in the provided libvirt connection URI is invalid.
//
// Hypervisor URIs may be local or remote, e.g.,
//	qemu:///system
//	qemu+ssh://libvirt@example.com/system
func NewLibvirtMonitor(uri, domain string) (*Libvirt, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	// Shell out to virsh to perform management actions
	cmd := &shellexec.SystemCommand{}

	monitor := &Libvirt{
		url:        u,
		domain:     domain,
		disconnect: make(chan error, 1),
		cmd:        cmd,
	}

	return monitor, nil
}

// TODO(mdlayher): move Virsh function to internal/libvirt or similar.

// Virsh is a wrapper for shelling out to the `virsh` executable.
func Virsh(command shellexec.Command, url string, args ...string) ([]byte, error) {
	args = append([]string{"-c", url}, args...)
	cmd := command.Prepare("virsh", args...)

	stdout, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			msg := strings.TrimPrefix(string(ee.Stderr), "error: ")
			return stdout, errors.New(strings.ToLower(msg))
		}

		return stdout, err
	}

	return stdout, nil
}

// streamEvents monitors the provided io.Reader, passing Events
// down stream. Any invalid event data encountered is discarded.
func streamEvents(data io.Reader, stream chan<- Event) {
	defer close(stream)

	var e Event
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		data := normalize(scanner.Bytes())
		if err := json.Unmarshal(data, &e); err == nil {
			stream <- e
		}
	}
}

// Libvirt does unspeakable things to the JSON emitted from QEMU.
// normalize attempts to rewrite events passed through libvirt
// to match the proper QEMU event JSON. This is best-effort, if
// the event formatting is invalid, the original value is returned.
func normalize(raw []byte) []byte {
	match := extractor.FindAllSubmatch(raw, -1)
	if match == nil {
		return raw
	}
	event := match[0]

	// handle libvirt's "<null>" "data" values.
	data := strings.Replace(string(event[4]), "<null>", "{}", 1)

	normalized := fmt.Sprintf(eventFormat, event[1], data, event[2], event[3])

	return []byte(normalized)
}
