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
	"encoding/json"
	"testing"
	"time"

	"github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-libvirt/libvirttest"
)

var testEvent = []byte{
	0x00, 0x00, 0x00, 0xb0, // length
	0x20, 0x00, 0x80, 0x87, // program
	0x00, 0x00, 0x00, 0x01, // version
	0x00, 0x00, 0x00, 0x06, // procedure
	0x00, 0x00, 0x00, 0x01, // type
	0x00, 0x00, 0x00, 0x00, // serial
	0x00, 0x00, 0x00, 0x00, // status
	0x00, 0x00, 0x00, 0x01, // callback id

	// domain name ("test")
	0x00, 0x00, 0x00, 0x04, 0x74, 0x65, 0x73, 0x74,

	// uuid (dc229f87d4de47198cfd2e21c6105b01)
	0xdc, 0x22, 0x9f, 0x87, 0xd4, 0xde, 0x47, 0x19,
	0x8c, 0xfd, 0x2e, 0x21, 0xc6, 0x10, 0x5b, 0x01,

	// domain id (14)
	0x00, 0x00, 0x00, 0x0e,

	// event name (BLOCK_JOB_COMPLETED)
	0x00, 0x00, 0x00, 0x13, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x00,

	// seconds (1462211891)
	0x00, 0x00, 0x00, 0x00, 0x57, 0x27, 0x95, 0x33,

	// microseconds (931791)
	0x00, 0x0e, 0x37, 0xcf,

	// event json data
	// ({"device":"drive-ide0-0-0","len":0,"offset":0,"speed":0,"type":"commit"})
	0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x48,
	0x7b, 0x22, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x3a, 0x22, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2d, 0x69, 0x64, 0x65, 0x30, 0x2d, 0x30, 0x2d,
	0x30, 0x22, 0x2c, 0x22, 0x6c, 0x65, 0x6e, 0x22,
	0x3a, 0x30, 0x2c, 0x22, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x3a, 0x30, 0x2c, 0x22, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x22, 0x3a, 0x30, 0x2c,
	0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x7d,
}

func TestNewLibvirtRPCMonitor(t *testing.T) {
	conn := libvirttest.New()

	domain := "test-1"
	rpc := NewLibvirtRPCMonitor(domain, conn)

	if rpc.Domain != domain {
		t.Errorf("expected domain %q, got %q", domain, rpc.Domain)
	}
}

func TestQMPEvent(t *testing.T) {
	e := &libvirt.DomainEvent{
		CallbackID:   1,
		Domain:       libvirt.Domain{Name: "test"},
		Event:        "test",
		Seconds:      1,
		Microseconds: 1,
		Details:      []byte(`{"device":"drive-virtio-disk0","len":0,"offset":0,"speed":0,"type":"commit"}`),
	}

	qe, err := qmpEvent(e)
	if err != nil {
		t.Error(err)
	}

	if qe.Event != e.Event {
		t.Errorf("expected event %q, got %q", e.Event, qe.Event)
	}

	if qe.Timestamp.Seconds != int64(e.Seconds) {
		t.Errorf("expected seconds to be %q, got %q", e.Seconds, qe.Timestamp.Seconds)
	}

	if qe.Timestamp.Microseconds != int64(e.Microseconds) {
		t.Errorf("expected microseconds to be %q, got %q", e.Microseconds, qe.Timestamp.Microseconds)
	}

	expected := "drive-virtio-disk0"
	actual, ok := qe.Data["device"]
	if !ok {
		t.Error("expected event data 'device'")
	}

	if actual != expected {
		t.Errorf("expected device %q, got %q", expected, actual)
	}
}

func TestLibvirtRPCMonitorConnect(t *testing.T) {
	conn := libvirttest.New()
	mon := NewLibvirtRPCMonitor("test", conn)

	err := mon.Connect()
	if err != nil {
		t.Error(err)
	}
}

func TestLibvirtRPCMonitorDisconnect(t *testing.T) {
	conn := libvirttest.New()
	mon := NewLibvirtRPCMonitor("test", conn)

	err := mon.Disconnect()
	if err != nil {
		t.Error(err)
	}
}

func TestLibvirtRPCMonitorRun(t *testing.T) {
	conn := libvirttest.New()
	mon := NewLibvirtRPCMonitor("test", conn)

	res, err := mon.Run([]byte(`{"query-version"}`))
	if err != nil {
		t.Error(err)
	}

	type version struct {
		Return struct {
			Package string `json:"package"`
			QEMU    struct {
				Major int `json:"major"`
				Micro int `json:"micro"`
				Minor int `json:"minor"`
			} `json:"qemu"`
		} `json:"return"`
	}

	var v version
	err = json.Unmarshal(res, &v)
	if err != nil {
		t.Error(err)
	}

	expected := 2
	if v.Return.QEMU.Major != expected {
		t.Errorf("expected qemu major version %d, got %d", expected, v.Return.QEMU.Major)
	}
}

func TestLibvirtRPCMonitorEvents(t *testing.T) {
	conn := libvirttest.New()
	mon := NewLibvirtRPCMonitor("test", conn)
	done := make(chan struct{})

	stream, err := mon.Events()
	if err != nil {
		t.Error(err)
	}

	go func() {
		var e Event
		select {
		case e = <-stream:
		case <-time.After(time.Second * 1):
			t.Error("expected event, received timeout")
		}

		expected := "drive-ide0-0-0"
		if e.Data["device"] != expected {
			t.Errorf("expected device %q, got %q", expected, e.Data["device"])
		}

		done <- struct{}{}
	}()

	// send an event to the listener goroutine
	_, _ = conn.Test.Write(testEvent)

	// wait for completion
	<-done
}
