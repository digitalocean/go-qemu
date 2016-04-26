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
	"bytes"
	"testing"
)

const (
	TestLibvirtURI = "qemu:///system"
	TestVM         = "domain-00000"
)

func TestNewLibvirtMonitor(t *testing.T) {
	monitor, err := NewLibvirtMonitor(TestLibvirtURI, TestVM)
	if err != nil {
		t.Error(err)
	}

	if monitor == nil {
		t.Error("failed to create monitor")
	}
}

func TestDisconnect(t *testing.T) {
	monitor, err := NewLibvirtMonitor(TestLibvirtURI, TestVM)
	if err != nil {
		t.Error(err)
	}

	err = monitor.Disconnect()
	if err != nil {
		t.Error("failed to disconnect monitor")
	}
}

func TestConnectInvalidURI(t *testing.T) {
	_, err := NewLibvirtMonitor("%%%%", TestVM)
	if err == nil {
		t.Errorf("expected invalid libvirt uri: %q to fail", err)
	}
}

func TestNormalize(t *testing.T) {
	raw := []byte(`event BLOCK_JOB_READY at 1457536213.106725 for domain domain-00000: {"device":"drive-virtio-disk0","len":0,"offset":0,"speed":0,"type":"commit"}`)
	expected := []byte(`{ "event": "BLOCK_JOB_READY", "data": {"device":"drive-virtio-disk0","len":0,"offset":0,"speed":0,"type":"commit"}, "timestamp": { "seconds": 1457536213, "microseconds": 106725 } }`)

	actual := normalize([]byte(raw))

	if !bytes.Equal(actual, expected) {
		t.Errorf("expected normalized event to be '%s' instead got '%s'", string(expected), string(actual))
	}
}

func TestNormalizeNull(t *testing.T) {
	raw := []byte(`event BLOCK_JOB_READY at 1457536213.106725 for domain domain-00000: <null>`)
	expected := []byte(`{ "event": "BLOCK_JOB_READY", "data": {}, "timestamp": { "seconds": 1457536213, "microseconds": 106725 } }`)

	actual := normalize([]byte(raw))

	if !bytes.Equal(actual, expected) {
		t.Errorf("expected normalized event to be %q instead got %q", string(expected), string(actual))
	}
}

func TestNormalizeInvalidEvent(t *testing.T) {
	expected := []byte(`foo`)
	actual := normalize(expected)

	if !bytes.Equal(actual, expected) {
		t.Errorf("expected invalid event to be %q instead got %q", string(expected), string(actual))
	}
}

func TestEventStream(t *testing.T) {
	b := bytes.NewBufferString(`event BLOCK_JOB_READY at 1457536213.106725 for domain domain-00000: {"device":"drive-virtio-disk0","len":0,"offset":0,"speed":0,"type":"commit"}`)
	stream := make(chan Event, 1)

	streamEvents(b, stream)

	select {
	case e := <-stream:
		expectedE := "BLOCK_JOB_READY"
		actualE := e.Event
		if expectedE != actualE {
			t.Errorf("expected event %q instead got %q", expectedE, actualE)
		}

		expectedS := int64(1457536213)
		actualS := e.Timestamp.Seconds
		if expectedS != actualS {
			t.Errorf("expected event timestamp seconds %d instead got %d", expectedS, actualS)
		}
	default:
		t.Error("expected event from event stream")
	}
}
