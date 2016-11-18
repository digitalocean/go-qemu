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

package qemu

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/digitalocean/go-qemu/qmp"
	"github.com/digitalocean/go-qemu/qmp/raw"
)

const defaultTestTimeout = 50 * time.Millisecond

func TestBlockDevice(t *testing.T) {
	const device = "drive-virtio-disk0"

	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-block", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return queryBlockResponse{
			Return: []BlockDevice{{
				Device: device,
			}},
		}, nil
	})
	defer done()

	bd, err := d.BlockDevice(device)
	if err != nil {
		t.Error(err)
	}

	if bd.Device != device {
		t.Errorf("expected device %q, got %q", device, bd.Device)
	}
}

func TestBlockDeviceNotFound(t *testing.T) {
	const device = "drive-virtio-disk0"

	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-block", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return queryBlockResponse{
			Return: []BlockDevice{{
				Device: device,
			}},
		}, nil
	})
	defer done()

	_, err := d.BlockDevice("foo")
	if err == nil {
		t.Errorf("expected block device %q to not exist", device)
	}

	if err != ErrBlockDeviceNotFound {
		t.Errorf("expected ErrBlockDeviceNotFound")
	}
}

func TestBlockDevices(t *testing.T) {
	const device = "drive-virtio-disk0"

	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-block", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return queryBlockResponse{
			Return: []BlockDevice{{
				Device: device,
			}},
		}, nil
	})
	defer done()

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	expectedLen := 1
	actualLen := len(disks)
	if actualLen != expectedLen {
		t.Errorf("expected %d disks, got %d", expectedLen, actualLen)
	}

	expected := "drive-virtio-disk0"
	actual := disks[0].Device
	if expected != actual {
		t.Errorf("expected device %q, got %q", expected, actual)
	}
}

func TestBlockJobs(t *testing.T) {
	const device = "drive-virtio-disk0"

	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-block-jobs", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return success{
			Return: []BlockJob{{
				Device:   device,
				IOStatus: "ok",
			}},
		}, nil
	})
	defer done()

	jobs, err := d.BlockJobs()
	if err != nil {
		t.Error(err)
	}

	if len(jobs) != 1 {
		t.Error("expected running backup job")
	}

	expected := "ok"
	if jobs[0].IOStatus != expected {
		t.Errorf("expected i/o status %q, got %q", expected, jobs[0].IOStatus)
	}

	expected = device
	if jobs[0].Device != expected {
		t.Errorf("expected device %q, got %q", expected, jobs[0].Device)
	}
}

func TestBlockStats(t *testing.T) {
	const device = "drive-virtio-disk0"
	const bytes = 9786368

	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-blockstats", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		type response struct {
			Device string
			Stats  BlockStats `json:"stats"`
		}

		return success{
			Return: []response{{
				Device: device,
				Stats: BlockStats{
					WriteBytes: bytes,
				},
			}},
		}, nil
	})
	defer done()

	stats, err := d.BlockStats()
	if err != nil {
		t.Error(err)
	}

	if stats[0].Device != device {
		t.Errorf("expected device %q, got %q", device, stats[0].Device)
	}

	expectedBytes := uint64(9786368)
	if stats[0].WriteBytes != expectedBytes {
		t.Errorf("expected %d write bytes, got %d", expectedBytes, stats[0].WriteBytes)
	}
}

func TestClose(t *testing.T) {
	m := &testMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	if err := d.Close(); err != nil {
		t.Error(err)
	}

	if _, ok := <-d.done; ok {
		t.Error("domain should be closed")
	}
}

func TestCommands(t *testing.T) {
	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-commands", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		type command struct {
			Name string
		}

		return success{
			Return: []command{
				{Name: "query-block"},
				{Name: "qeury-foo"},
			},
		}, nil
	})
	defer done()

	cmds, err := d.Commands()
	if err != nil {
		t.Error(err)
	}

	expected := 2
	actual := len(cmds)
	if actual != expected {
		t.Errorf("expected number of supported commands to be %d, got %d", expected, actual)
	}

	found := false
	search := "query-block"
	for _, c := range cmds {
		if c == search {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected command %q to be returned", search)
	}
}

func TestDomainScreenDump(t *testing.T) {
	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "screendump", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return success{}, nil
	})
	defer done()

	// Use a fixed file name generation function
	name := filepath.Join(os.TempDir(), "test-screendump")
	d.tempFileName = func(_ string, _ string) string {
		return name
	}

	want := []byte("hello world")
	if err := ioutil.WriteFile(name, want, 0666); err != nil {
		t.Error(err)
	}

	rc, err := d.ScreenDump()
	if err != nil {
		t.Error(err)
	}

	got, err := ioutil.ReadAll(rc)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(want, got) {
		t.Errorf("unexpected bytes:\n- want: %v\n-  got: %v", want, got)
	}

	if err := rc.Close(); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(name); !os.IsNotExist(err) {
		t.Errorf("file should no longer exist, but got: %v", err)
	}
}

func TestPCIDevices(t *testing.T) {
	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-pci", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		type classInfo struct {
			Desc string `json:"desc"`
		}

		type pciDevice struct {
			Bus       int       `json:"bus"`
			ClassInfo classInfo `json:"class_info"`
		}

		type wrapReturn struct {
			Bus     int         `json:"bus"`
			Devices []pciDevice `json:"devices"`
		}

		response := struct {
			Return []wrapReturn `json:"return"`
		}{
			Return: []wrapReturn{
				{
					Bus: 0,
					Devices: []pciDevice{{
						Bus: 0,
					}},
				},
				{
					Bus: 1,
					Devices: []pciDevice{{
						Bus: 1,
						ClassInfo: classInfo{
							Desc: "Intel Ethernet controller",
						},
					}},
				},
			},
		}

		return response, nil
	})
	defer done()

	devices, err := d.PCIDevices()
	if err != nil {
		t.Error(err)
	}

	if len(devices) != 2 {
		t.Error("expected two PCI devices")
	}

	expected := 0
	if devices[0].Bus != expected {
		t.Errorf("expected device bus %d, got %q", expected, devices[0].Bus)
	}

	expectedDesc := "Intel Ethernet controller"
	if devices[1].ClassInfo.Desc != expectedDesc {
		t.Errorf("expected device %q, got %q", expectedDesc, devices[1].ClassInfo.Desc)
	}
}

func TestCPUHalted(t *testing.T) {
	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-cpus", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		response := struct {
			Return []CPU `json:"return"`
		}{
			Return: []CPU{
				{
					CPU:     0,
					Current: true,
					Halted:  true,
					PC:      -2130295322,
				},
			},
		}
		return response, nil
	})
	defer done()

	cpu, err := d.CPUs()
	if err != nil {
		t.Error(err)
	}

	if len(cpu) != 1 {
		t.Error("expected one CPU")
	}

	if !cpu[0].Halted {
		t.Error("expected CPU to be halted")
	}

}

func TestStatusRunning(t *testing.T) {
	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-status", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return success{
			Return: raw.StatusInfo{
				Status: raw.RunStateRunning,
			},
		}, nil
	})
	defer done()

	status, err := d.Status()
	if err != nil {
		t.Error(err)
	}

	if status != StatusRunning {
		t.Error("expected domain to be running")
	}
}

func TestStatusShutdown(t *testing.T) {
	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-status", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return success{
			Return: raw.StatusInfo{
				Status: raw.RunStateShutdown,
			},
		}, nil
	})
	defer done()

	status, err := d.Status()
	if err != nil {
		t.Error(err)
	}

	if status != StatusShutdown {
		t.Error("expected domain to be powered off")
	}
}

func TestSupported(t *testing.T) {
	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-commands", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		type command struct {
			Name string
		}

		return success{
			Return: []command{
				{"query-block"},
			},
		}, nil
	})
	defer done()

	cmd := "query-block"
	supported, err := d.Supported(cmd)
	if err != nil {
		t.Error(err)
	}

	if !supported {
		t.Errorf("expected command %q to be supported", cmd)
	}
}

func TestSupportedFalse(t *testing.T) {
	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "query-commands", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		type command struct {
			Name string
		}

		return success{
			Return: []command{
				{"query-bar"},
				{"query-baz"},
			},
		}, nil
	})
	defer done()

	cmd := "query-foo"
	supported, err := d.Supported(cmd)
	if err != nil {
		t.Error(err)
	}

	if supported {
		t.Errorf("expected command %q to be unsupported", cmd)
	}
}

func TestSystemPowerdown(t *testing.T) {
	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "system_powerdown", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return success{}, nil
	})
	defer done()

	if err := d.SystemPowerdown(); err != nil {
		t.Errorf("error powering down domain: %v", err)
	}
}

type success struct {
	Return interface{} `json:"return"`
}

type failure struct {
	Error interface{} `json:"error"`
}

func TestSystemReset(t *testing.T) {
	d, done := testDomain(t, func(cmd qmp.Command) (interface{}, error) {
		if want, got := "system_reset", cmd.Execute; want != got {
			t.Fatalf("unexpected QMP command:\n- want: %q\n-  got: %q",
				want, got)
		}

		return success{}, nil
	})
	defer done()

	if err := d.SystemReset(); err != nil {
		t.Errorf("error resetting domain: %v", err)
	}
}

func TestEvents(t *testing.T) {
	d, done := testDomain(t, nil)
	defer done()

	events, stop, err := d.Events()
	if err != nil {
		t.Error(err)
	}

	select {
	case <-events:
		stop <- struct{}{}
	case <-time.After(time.Millisecond * 20):
		t.Error("expected event")
	}
}

func TestEventsUnsupported(t *testing.T) {
	d, done := testDomain(t, nil)
	defer done()
	d.eventsUnsupported = true

	_, _, err := d.Events()
	if err != qmp.ErrEventsNotSupported {
		t.Errorf("expected qmp.ErrEventsNotSupported, got %s", err.Error())
	}
}

func testDomain(t *testing.T, fn func(qmp.Command) (interface{}, error)) (*Domain, func()) {
	mon := &testMonitor{fn: fn}
	d, err := NewDomain(mon, "test")
	if err != nil {
		t.Fatalf("failed to create test domain: %v", err)
	}

	return d, func() {
		_ = d.Close()
	}
}

var _ qmp.Monitor = &testMonitor{}

type testMonitor struct {
	eventTimeout bool
	eventErrors  bool
	fn           func(qmp.Command) (interface{}, error)
	noopMonitor
}

func (t *testMonitor) Run(raw []byte) ([]byte, error) {
	var cmd qmp.Command
	if err := json.Unmarshal(raw, &cmd); err != nil {
		return nil, err
	}

	result, err := t.fn(cmd)
	if err != nil {
		return nil, err
	}

	return json.Marshal(result)
}

func (t *testMonitor) Events() (<-chan qmp.Event, error) {
	c := make(chan qmp.Event)
	go func() {
		events := []string{blockJobReady, blockJobCompleted}

		i := 0
		for {
			if t.eventTimeout {
				<-time.After(10 * time.Millisecond)
				continue
			}

			if t.eventErrors {
				c <- qmp.Event{Event: blockJobError}
				<-time.After(10 * time.Millisecond)
				continue
			}

			c <- qmp.Event{Event: events[i]}
			<-time.After(10 * time.Millisecond)

			i = (i + 1) % len(events)
		}
	}()

	return c, nil
}

var _ qmp.Monitor = &noopMonitor{}

type noopMonitor struct{}

func (noopMonitor) Connect() error                    { return nil }
func (noopMonitor) Disconnect() error                 { return nil }
func (noopMonitor) Run(_ []byte) ([]byte, error)      { return nil, nil }
func (noopMonitor) Events() (<-chan qmp.Event, error) { return nil, nil }
