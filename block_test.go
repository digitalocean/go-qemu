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
	"testing"
)

func TestBlockDevices(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	expectedLen := 2
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

func TestBlockDevicesMonitorFail(t *testing.T) {
	m := &mockMonitor{}
	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	m.alwaysFail = true
	_, err = d.BlockDevices()
	if err == nil {
		t.Errorf("expected monitor failure")
	}
}

func TestBlockDevicesInvalidJSON(t *testing.T) {
	m := &mockMonitor{invalidJSON: true}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	_, err = d.BlockDevices()
	if err == nil {
		t.Errorf("expected invalid json to cause failure")
	}
}

func TestCancelJob(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	err = disks[0].CancelJob(d, defaultTestTimeout)
	if err != nil {
		t.Error(err)
	}
}

func TestCancelJobMonitorFailure(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	m.alwaysFail = true
	err = disks[0].CancelJob(d, defaultTestTimeout)
	if err == nil {
		t.Error("expected monitor failure")
	}
}

func TestCommit(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	err = disks[0].Commit(d, "/tmp/foo", defaultTestTimeout)
	if err != nil {
		t.Error(err)
	}
}

func TestCommitActiveBlockJob(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	m.activeJobs = true
	err = disks[0].Commit(d, "/tmp/foo", defaultTestTimeout)
	if err == nil {
		t.Errorf("expected blockcommit with active blockjob to fail")
	}
}

func TestCommitBlockJobError(t *testing.T) {
	m := &mockMonitor{eventErrors: true}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	err = disks[0].Commit(d, "/tmp/foo", defaultTestTimeout)
	if err == nil {
		t.Error("expected block job error to cause failure")
	}
}

func TestCommitTimeout(t *testing.T) {
	m := &mockMonitor{eventTimeout: true}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	err = disks[0].Commit(d, "/tmp/foo", 0)
	if err == nil {
		t.Error("expected timeout")
	}
}

func TestJobComplete(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	err = disks[0].CompleteJob(d, defaultTestTimeout)
	if err != nil {
		t.Error(err)
	}
}

func TestJobCompleteMonitorFail(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	m.alwaysFail = true
	err = disks[0].CompleteJob(d, defaultTestTimeout)
	if err == nil {
		t.Error("expected monitor failure")
	}
}

func TestJobCompleteEventError(t *testing.T) {
	m := &mockMonitor{eventErrors: true}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	err = disks[0].CompleteJob(d, defaultTestTimeout)
	if err == nil {
		t.Error("expected block job error to cause failure")
	}
}

func TestJobCompleteTimeout(t *testing.T) {
	m := &mockMonitor{eventTimeout: true}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	err = disks[0].CompleteJob(d, 0)
	if err == nil {
		t.Error("expected timeout")
	}
}

func TestMirror(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	err = disks[0].Mirror(d, "/tmp/foo.img", defaultTestTimeout)
	if err != nil {
		t.Error(err)
	}
}

func TestMirrorRelativePath(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	dest := "foo.img"
	err = disks[0].Mirror(d, dest, defaultTestTimeout)
	if err == nil {
		t.Errorf("expected relative path %q to fail", dest)
	}
}

func TestMirrorMonitorFailure(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	m.alwaysFail = true
	err = disks[0].Mirror(d, "/tmp/foo.img", defaultTestTimeout)
	if err == nil {
		t.Error("expected monitor failure")
	}
}

func TestSnapshot(t *testing.T) {
	m := &mockMonitor{}

	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	err = disks[0].Snapshot(d, "/tmp/foo.img")
	if err != nil {
		t.Error(err)
	}
}

func TestSnapshotMonitorFail(t *testing.T) {
	m := &mockMonitor{}
	d, err := NewDomain(m, "foo")
	if err != nil {
		t.Error(err)
	}

	disks, err := d.BlockDevices()
	if err != nil {
		t.Error(err)
	}

	m.alwaysFail = true
	err = disks[0].Snapshot(d, "/tmp/foo.img")
	if err == nil {
		t.Error("expected monitor failure")
	}
}
