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

//go:generate stringer -type=Status -output=string.gen.go
//go:generate ../scripts/prependlicense.sh string.gen.go

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/digitalocean/go-qemu/qmp"
	"github.com/digitalocean/go-qemu/qmp/raw"
)

var (
	// ErrBlockDeviceNotFound is returned given when a block device is not found
	ErrBlockDeviceNotFound = errors.New("block device not found")
)

// Domain represents a QEMU instance.
type Domain struct {
	Name       string
	m          qmp.Monitor
	rm         *raw.Monitor
	done       chan struct{}
	connect    chan chan qmp.Event
	disconnect chan chan qmp.Event
	listeners  []chan qmp.Event

	eventsUnsupported bool

	tempFileName func(domainName string, method string) string
}

// Close cleans up internal resources of a Domain and disconnects the underlying
// qmp.Monitor.  Close must be called when done with a Domain to avoid leaking
// resources.
func (d *Domain) Close() error {
	close(d.done)
	return d.m.Disconnect()
}

// Commands returns all QMP commands supported by the domain.
func (d *Domain) Commands() ([]string, error) {
	commands, err := d.rm.QueryCommands()
	if err != nil {
		return nil, err
	}

	// flatten response
	cmds := make([]string, 0, len(commands))
	for _, c := range commands {
		cmds = append(cmds, c.Name)
	}

	return cmds, nil
}

// queryBlockResponse is the structure returned by QEMU in response to
// a query-block QMP command.
type queryBlockResponse struct {
	ID     string        `json:"id"`
	Return []BlockDevice `json:"return"`
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
	raw, err := d.Run(qmp.Command{Execute: "query-block"})
	if err != nil {
		return []BlockDevice{}, err
	}

	response := queryBlockResponse{}
	if err = json.Unmarshal(raw, &response); err != nil {
		return []BlockDevice{}, err
	}

	return response.Return, nil
}

// BlockJobs returns active block job operations.
func (d *Domain) BlockJobs() ([]BlockJob, error) {
	var jobs []BlockJob
	raw, err := d.Run(qmp.Command{Execute: "query-block-jobs"})
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
	raw, err := d.Run(qmp.Command{Execute: "query-blockstats"})
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
	raw, err := d.Run(qmp.Command{Execute: "query-pci"})
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

// CPUs returns a domain's CPUs.
func (d *Domain) CPUs() ([]CPU, error) {
	raw, err := d.Run(qmp.Command{Execute: "query-cpus"})
	if err != nil {
		return nil, err
	}

	var response struct {
		Return []CPU `json:"return"`
	}

	if err = json.Unmarshal(raw, &response); err != nil {
		return nil, err
	}

	return response.Return, nil
}

// Run executes the given QMP command against the domain.
// The returned []byte is the raw output from the QMP monitor.
//
// Run should be used with caution, as it allows the execution of
// arbitrary QMP commands against the domain.
func (d *Domain) Run(c qmp.Command) ([]byte, error) {
	cmd, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return d.m.Run(cmd)
}

// ScreenDump captures the domain's screen and creates an output PPM image
// stream.  ScreenDump will only work if the Domain resides on a local
// hypervisor; not a remote one connected over SSH or TCP socket.
//
// If needed, a PPM image can be decoded using Go's package image and a
// PPM decoder, such as https://godoc.org/github.com/jbuchbinder/gopnm.
func (d *Domain) ScreenDump() (io.ReadCloser, error) {
	// Since QEMU only allows capturing output to a file, we create a
	// temporary file and use it for the screendump, providing a stream
	// to it on return which can be used with anything that accepts
	// io.Reader.
	name := d.tempFileName(d.Name, "screendump")

	cmd := qmp.Command{
		Execute: "screendump",
		Args: map[string]string{
			"filename": name,
		},
	}
	if _, err := d.Run(cmd); err != nil {
		return nil, err
	}

	// Automatically remove temporary file when the Close method
	// is called.
	return newRemoveFileReadCloser(name)
}

// Status represents the current status of the domain.
type Status int

// Status constants which indicate the status of a domain.
const (
	StatusDebug         Status = Status(raw.RunStateDebug)
	StatusFinishMigrate Status = Status(raw.RunStateFinishMigrate)
	StatusGuestPanicked Status = Status(raw.RunStateGuestPanicked)
	StatusIOError       Status = Status(raw.RunStateIOError)
	StatusInMigrate     Status = Status(raw.RunStateInmigrate)
	StatusInternalError Status = Status(raw.RunStateInternalError)
	StatusPaused        Status = Status(raw.RunStatePaused)
	StatusPostMigrate   Status = Status(raw.RunStatePostmigrate)
	StatusPreLaunch     Status = Status(raw.RunStatePrelaunch)
	StatusRestoreVM     Status = Status(raw.RunStateRestoreVM)
	StatusRunning       Status = Status(raw.RunStateRunning)
	StatusSaveVM        Status = Status(raw.RunStateSaveVM)
	StatusShutdown      Status = Status(raw.RunStateShutdown)
	StatusSuspended     Status = Status(raw.RunStateSuspended)
	StatusWatchdog      Status = Status(raw.RunStateWatchdog)
)

// Status returns the current status of the domain.
func (d *Domain) Status() (Status, error) {
	status, err := d.rm.QueryStatus()
	if err != nil {
		// libvirt returns an error if the domain is not running
		if strings.Contains(err.Error(), "not running") {
			return StatusShutdown, nil
		}

		return 0, err
	}

	return Status(status.Status), nil
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

// SystemPowerdown sends a system power down event to the domain.
func (d *Domain) SystemPowerdown() error {
	_, err := d.Run(qmp.Command{Execute: "system_powerdown"})
	return err
}

// SystemReset sends a system reset event to the domain.
func (d *Domain) SystemReset() error {
	_, err := d.Run(qmp.Command{Execute: "system_reset"})
	return err
}

// Version returns the domain's QEMU version.
func (d *Domain) Version() (string, error) {
	raw, err := d.Run(qmp.Command{Execute: "query-version"})
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

// PackageVersion returns the domain's QEMU package version, the full build
// information for QEMU.
func (d *Domain) PackageVersion() (string, error) {
	vers, err := d.rm.QueryVersion()
	if err != nil {
		return "", err
	}

	return vers.Package, nil
}

// Events streams QEMU QMP events.
// Two channels are returned, the first contains events emitted by the domain.
// The second is used to signal completion of event processing.
// It is the responsibility of the caller to always signal when finished.
func (d *Domain) Events() (chan qmp.Event, chan struct{}, error) {
	if d.eventsUnsupported {
		return nil, nil, qmp.ErrEventsNotSupported
	}

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

	return stream, done, nil
}

// listenAndServe handles a domain's event broadcast service.
func (d *Domain) listenAndServe() error {
	stream, err := d.m.Events()
	if err != nil {
		// let Event() inform the user events are not supported
		if err == qmp.ErrEventsNotSupported {
			d.eventsUnsupported = true
			return nil
		}

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
		rm:         raw.NewMonitor(m),
		done:       make(chan struct{}),
		connect:    make(chan chan qmp.Event),
		disconnect: make(chan chan qmp.Event),
		listeners:  []chan qmp.Event{},

		// By default, try to generate decently random file names
		// for temporary files.
		tempFileName: func(domainName string, method string) string {
			return filepath.Join(
				os.TempDir(),
				fmt.Sprintf("go-qemu-%s-%s-%d",
					domainName,
					method,
					time.Now().UnixNano(),
				),
			)
		},
	}

	// start event broadcast
	err := d.listenAndServe()

	return d, err
}
