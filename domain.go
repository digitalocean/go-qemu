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

// Package qemu provides an interface for interacting with running QEMU instances.
package qemu

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/digitalocean/go-qemu/qmp"
)

var (
	// ErrBlockDeviceNotFound is returned given when a block device is not found
	ErrBlockDeviceNotFound = errors.New("block device not found")
)

// Domain represents a QEMU instance.
type Domain struct {
	Name       string
	m          qmp.Monitor
	done       chan struct{}
	connect    chan chan qmp.Event
	disconnect chan chan qmp.Event
	listeners  []chan qmp.Event
}

// Close cleans up internal resources of a Domain.  Close must be called
// when done with a Domain to avoid leaking resources.
//
// Close does not disconnect the underlying qmp.Monitor.
func (d *Domain) Close() error {
	close(d.done)
	return nil
}

// Commands returns all QMP commands supported by the domain.
func (d *Domain) Commands() ([]string, error) {
	raw, err := d.Run(qmp.Cmd{Execute: "query-commands"})
	if err != nil {
		return nil, err
	}

	var response struct {
		ID     string `json:"id"`
		Return []struct {
			Name string `json:"name"`
		} `json:"return"`
	}

	if err = json.Unmarshal(raw, &response); err != nil {
		return nil, err
	}

	// flatten response
	cmds := make([]string, 0, len(response.Return))
	for _, c := range response.Return {
		cmds = append(cmds, c.Name)
	}

	return cmds, nil
}

// BlockDevice searches a domain for the given block device.
// If found, a BlockDevice is returned. If the device is not found,
// the returned error will be ErrBlockDeviceNotFound.
func (d *Domain) BlockDevice(name string) (BlockDevice, error) {
	devs, err := d.BlockDevices()
	if err != nil {
		return BlockDevice{}, err
	}

	for _, d := range devs {
		if d.Device == name {
			return d, nil
		}
	}

	return BlockDevice{}, ErrBlockDeviceNotFound
}

// BlockDevices returns a domain's block devices.
func (d *Domain) BlockDevices() ([]BlockDevice, error) {
	raw, err := d.Run(qmp.Cmd{Execute: "query-block"})
	if err != nil {
		return []BlockDevice{}, err
	}

	var response struct {
		ID     string        `json:"id"`
		Return []BlockDevice `json:"return"`
	}

	if err = json.Unmarshal(raw, &response); err != nil {
		return []BlockDevice{}, err
	}

	return response.Return, nil
}

// BlockJobs returns active block job operations.
func (d *Domain) BlockJobs() ([]BlockJob, error) {
	var jobs []BlockJob
	raw, err := d.Run(qmp.Cmd{Execute: "query-block-jobs"})
	if err != nil {
		return jobs, err
	}

	var response struct {
		ID     string     `json:"id"`
		Return []BlockJob `json:"return"`
	}

	if err = json.Unmarshal(raw, &response); err != nil {
		return jobs, err
	}

	return response.Return, nil
}

// BlockStats returns block device statistics for a domain.
func (d *Domain) BlockStats() ([]BlockStats, error) {
	raw, err := d.Run(qmp.Cmd{Execute: "query-blockstats"})
	if err != nil {
		return nil, err
	}

	var response struct {
		Return []struct {
			Device string `json:"device"`
			// TODO(mdlayher): figure out what Parent is.
			Parent struct {
				Stats BlockStats `json:"stats"`
			} `json:"parent,omitempty"`
			Stats BlockStats `json:"stats"`
		} `json:"return"`
	}

	if err = json.Unmarshal(raw, &response); err != nil {
		return nil, err
	}

	stats := make([]BlockStats, 0, len(response.Return))
	for _, s := range response.Return {
		// Add device to the stat structure, even though QEMU does
		// not place it there
		s.Stats.Device = s.Device
		stats = append(stats, s.Stats)
	}

	return stats, nil
}

// PCIDevices returns a domain's PCI devices.
func (d *Domain) PCIDevices() ([]PCIDevice, error) {
	raw, err := d.Run(qmp.Cmd{Execute: "query-pci"})
	if err != nil {
		return nil, err
	}

	var response struct {
		Return []struct {
			Bus     int         `json:"bus"`
			Devices []PCIDevice `json:"devices"`
		} `json:"return"`
	}

	if err = json.Unmarshal(raw, &response); err != nil {
		return nil, err
	}

	// Merge devices from each bus into slice after counting them up
	// so only a single allocation is needed
	var count int
	for i := range response.Return {
		count += len(response.Return[i].Devices)
	}

	devices := make([]PCIDevice, 0, count)
	for _, bus := range response.Return {
		devices = append(devices, bus.Devices...)
	}

	return devices, nil
}

// Run executes the given QMP command against the domain.
// The returned []byte is the raw output from the QMP monitor.
//
// Run should be used with caution, as it allows the execution of
// arbitrary QMP commands against the domain.
func (d *Domain) Run(c qmp.Cmd) ([]byte, error) {
	cmd, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return d.m.Run(cmd)
}

// Status represents the current status of the domain.
type Status string

// Status constants which indicate the status of a domain.
const (
	StatusDebug         Status = "debug"
	StatusFinishMigrate Status = "finish-migrate"
	StatusIOError       Status = "io-error"
	StatusInMigrate     Status = "inmigrate"
	StatusInternalError Status = "internal-error"
	StatusPaused        Status = "paused"
	StatusPostMigrate   Status = "postmigrate"
	StatusPreLaunch     Status = "prelaunch"
	StatusRestoreVM     Status = "restore-vm"
	StatusRunning       Status = "running"
	StatusSaveVM        Status = "save-vm"
	StatusShutdown      Status = "shutdown"
	StatusWatchdog      Status = "watchdog"
)

// Status returns the current status of the domain.
func (d *Domain) Status() (Status, error) {
	raw, err := d.Run(qmp.Cmd{Execute: "query-status"})
	if err != nil {
		// libvirt returns an error if the domain is not running
		if strings.Contains(err.Error(), "not running") {
			return StatusShutdown, nil
		}

		return "", err
	}

	var response struct {
		ID     string `json:"id"`
		Return struct {
			Running    bool   `json:"running"`
			Singlestep bool   `json:"singlestep"`
			Status     string `json:"status"`
		} `json:"return"`
	}

	if err = json.Unmarshal(raw, &response); err != nil {
		return "", err
	}

	return Status(response.Return.Status), nil
}

// Supported returns true if the provided command is supported by the domain.
func (d *Domain) Supported(cmd string) (bool, error) {
	cmds, err := d.Commands()
	if err != nil {
		return false, err
	}

	for _, c := range cmds {
		if c == cmd {
			return true, nil
		}
	}

	return false, nil
}

// Version returns the domain's QEMU version.
func (d *Domain) Version() (string, error) {
	raw, err := d.Run(qmp.Cmd{Execute: "query-version"})
	if err != nil {
		return "", err
	}

	var response struct {
		ID     string      `json:"id"`
		Return qmp.Version `json:"return"`
	}

	if err = json.Unmarshal(raw, &response); err != nil {
		return "", err
	}

	return response.Return.String(), nil
}

// Events streams QEMU QMP events.
// Two channels are returned, the first contains events emitted by the domain.
// The second is used to signal completion of event processing.
// It is the responsibility of the caller to always signal when finished.
func (d *Domain) Events() (chan qmp.Event, chan struct{}) {
	stream := make(chan qmp.Event)
	done := make(chan struct{})

	// handle disconnection
	go func() {
		<-done
		d.disconnect <- stream
		close(stream)
		close(done)
	}()

	// add stream to broadcast
	d.connect <- stream

	return stream, done
}

// listenAndServe handles a domain's event broadcast service.
func (d *Domain) listenAndServe() error {
	stream, err := d.m.Events()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-d.done:
				return
			case client := <-d.connect:
				d.addListener(client)
			case client := <-d.disconnect:
				d.removeListener(client)
			case event := <-stream:
				d.broadcast(event)
			}
		}
	}()

	return nil
}

// addListener adds the given stream to the domain's event broadcast.
func (d *Domain) addListener(stream chan qmp.Event) {
	d.listeners = append(d.listeners, stream)
}

// removeListener removes the given stream from the domain's event broadcast.
func (d *Domain) removeListener(stream chan qmp.Event) {
	for i, client := range d.listeners {
		if client == stream {
			d.listeners = append(d.listeners[:i], d.listeners[i+1:]...)
		}
	}
}

// broadcast sends the provided event to all event listeners.
func (d *Domain) broadcast(event qmp.Event) {
	for _, stream := range d.listeners {
		stream <- event
	}
}

// NewDomain returns a new QEMU domain identified by the given name.
// QMP communication is handled by the provided monitor socket.
func NewDomain(m qmp.Monitor, name string) (*Domain, error) {
	d := &Domain{
		Name:       name,
		m:          m,
		done:       make(chan struct{}),
		connect:    make(chan chan qmp.Event),
		disconnect: make(chan chan qmp.Event),
		listeners:  []chan qmp.Event{},
	}

	// start event broadcast
	err := d.listenAndServe()

	return d, err
}
