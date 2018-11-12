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
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/digitalocean/go-qemu/qmp"
)

const (
	blockJobCompleted = "BLOCK_JOB_COMPLETED"
	blockJobError     = "BLOCK_JOB_ERROR"
	blockJobReady     = "BLOCK_JOB_READY"
)

// BlockDevice represents a QEMU block device.
type BlockDevice struct {
	Device   string `json:"device"`
	Inserted struct {
		BackingFile      string `json:"backing_file"`
		BackingFileDepth int    `json:"backing_file_depth"`
		BPS              int    `json:"bps"`
		BPSRead          int    `json:"bps_rd"`
		BPSWrite         int    `json:"bps_wr"`
		Cache            struct {
			Direct    bool `json:"direct"`
			NoFlush   bool `json:"no-flush"`
			Writeback bool `json:"writeback"`
		} `json:"cache"`
		DetectZeroes         string `json:"detect_zeroes"`
		Driver               string `json:"drv"`
		Encrypted            bool   `json:"encrypted"`
		EncryptionKeyMissing bool   `json:"encryption_key_missing"`
		File                 string `json:"file"`
		Image                Image  `json:"image"`
		IOPs                 int    `json:"iops"`
		IOPsRead             int    `json:"iops_rd"`
		IOPsWrite            int    `json:"iops_wr"`
		NodeName             string `json:"node-name"`
		ReadOnly             bool   `json:"ro"`
		WriteThreshold       int    `json:"write_threshold"`
	} `json:"inserted"`
	IOStatus  string `json:"io-status"`
	Locked    bool   `json:"locked"`
	Removable bool   `json:"removable"`
	Type      string `json:"type"`
}

// BlockJob represents a QEMU blockjob.
type BlockJob struct {
	Busy     bool   `json:"busy"`
	Device   string `json:"device"`
	IOStatus string `json:"io-status"`
	Len      int    `json:"len"`
	Offset   int    `json:"offset"`
	Paused   bool   `json:"paused"`
	Ready    bool   `json:"ready"`
	Speed    int    `json:"speed"`
	Type     string `json:"type"`
}

// Image represents a BlockDevice backing image.
type Image struct {
	ActualSize            uint64 `json:"actual-size"`
	BackingFilename       string `json:"backing-filename"`
	BackingFilenameFormat string `json:"backing-filename-format"`
	BackingImage          struct {
		ActualSize  uint64 `json:"actual-size"`
		Dirty       bool   `json:"dirty-flag"`
		Filename    string `json:"filename"`
		Format      string `json:"format"`
		VirtualSize uint64 `json:"virtual-size"`
	} `json:"backing-image"`
	ClusterSize    int    `json:"cluster-size"`
	Dirty          bool   `json:"dirty-flag"`
	Filename       string `json:"filename"`
	Format         string `json:"format"`
	FormatSpecific struct {
		Data struct {
			Compat        string `json:"compat"`
			Corrupt       bool   `json:"corrupt"`
			LazyRefcounts bool   `json:"lazy-refcounts"`
			RefcountBits  int    `json:"refcount-bits"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"format-specific"`
	VirtualSize uint64 `json:"virtual-size"`
}

// BlockStats represents QEMU block device statistics.
type BlockStats struct {
	// Device does not actually appear QEMU's output, but is added
	// by this package.
	Device string `json:"-"`

	AccountFailed             bool   `json:"account_failed"`
	AccountInvalid            bool   `json:"account_invalid"`
	FailedFlushOperations     int    `json:"failed_flush_operations"`
	FailedReadOperations      int    `json:"failed_rd_operations"`
	FailedWriteOperations     int    `json:"failed_wr_operations"`
	FlushOperations           int    `json:"flush_operations"`
	FlushTotalTimeNanoseconds int64  `json:"flush_total_time_ns"`
	IdleTimeNanoseconds       int64  `json:"idle_time_ns"`
	InvalidFlushOperations    int    `json:"invalid_flush_operations"`
	InvalidReadOperations     int    `json:"invalid_rd_operations"`
	InvalidWriteOperations    int    `json:"invalid_wr_operations"`
	ReadBytes                 uint64 `json:"rd_bytes"`
	ReadMerged                int    `json:"rd_merged"`
	ReadOperations            int    `json:"rd_operations"`
	ReadTotalTimeNanoseconds  int    `json:"rd_total_time_ns"`
	WriteBytes                uint64 `json:"wr_bytes"`
	WriteHighestOffset        int64  `json:"wr_highest_offset"`
	WriteMerged               int    `json:"wr_merged"`
	WriteOperations           int    `json:"wr_operations"`
	WriteTotalTimeNanoseconds int64  `json:"wr_total_time_ns"`
}

// Mirror copies a block device to the given destination.
// The destination path specified by dest must be absolute.
func (bd BlockDevice) Mirror(d *Domain, dest string, timeout time.Duration) error {
	if !filepath.IsAbs(dest) {
		return errors.New("destination path must be absolute")
	}

	return waitForSignal(d, blockJobReady, timeout, func() error {
		cmd := qmp.Command{
			Execute: "drive-mirror",
			Args: map[string]string{
				"device": bd.Device,
				"target": dest,
				"sync":   "full",
				"mode":   "absolute-paths",
				"format": "raw",
			},
		}

		_, err := d.Run(cmd)
		return err
	})
}

// Commit synchronously merges an overlay image onto a block device's
// root backing image. Once the operation is complete, CompleteJob()
// must be called to pivot the domain back to the original backing image.
func (bd BlockDevice) Commit(d *Domain, overlay string, timeout time.Duration) error {
	return waitForSignal(d, blockJobReady, timeout, func() error {
		cmd := qmp.Command{
			Execute: "block-commit",
			Args: map[string]string{
				"device": bd.Device,
				"top":    overlay,
			},
		}

		_, err := d.Run(cmd)
		return err
	})
}

// CancelJob cancels any active block jobs on a block device.
// For block-mirror operations, this completes the block job.
func (bd BlockDevice) CancelJob(d *Domain, timeout time.Duration) error {
	return waitForSignal(d, blockJobCompleted, timeout, func() error {
		cmd := qmp.Command{
			Execute: "block-job-cancel",
			Args: map[string]string{
				"device": bd.Device,
			},
		}

		_, err := d.Run(cmd)
		return err
	})
}

// CompleteJob finalizes any running block jobs on the provided block device.
// For blockcommit backups, this performs the "pivot" back to the original
// backing image.
func (bd BlockDevice) CompleteJob(d *Domain, timeout time.Duration) error {
	return waitForSignal(d, blockJobCompleted, timeout, func() error {
		cmd := qmp.Command{
			Execute: "block-job-complete",
			Args: map[string]string{
				"device": bd.Device,
			},
		}

		_, err := d.Run(cmd)
		return err
	})
}

// Snapshot creates a point in time snapshot.
// The disk's image is given a new QCOW2 overlay, leaving the underlying image
// in a state that is considered safe for copying.
func (bd BlockDevice) Snapshot(d *Domain, overlay string) error {
	cmd := qmp.Command{
		Execute: "blockdev-snapshot-sync",
		Args: map[string]string{
			"device":        bd.Device,
			"snapshot-file": overlay,
		},
	}

	_, err := d.Run(cmd)
	return err
}

// waitForSignal opens a domain's QMP event stream and invokes an input
// closure to send commands which would create results on that event stream.
func waitForSignal(d *Domain, signal string, timeout time.Duration, fn func() error) error {
	// "done" signal must always be sent to avoid leaking goroutines.
	events, done, err := d.Events()
	if err != nil {
		return err
	}
	defer func() { done <- struct{}{} }()

	// start listening for events prior to command execution. QMP events
	// may emit before the command returns.
	jobErr := make(chan error)
	go func() {
		jobErr <- waitForJob(events, signal, timeout)
	}()

	if err := fn(); err != nil {
		return err
	}

	return <-jobErr
}

// waitForJob monitors the domain's QMP event stream, waiting
// for the provided signal, timeout, or BLOCK_JOB_ERROR.
// An error is returned should either BLOCK_JOB_ERROR or timeout occur.
func waitForJob(events chan qmp.Event, signal string, timeout time.Duration) error {
	// Consider events stalled after timeout for X amount of time total,
	// rather than X amount of time without an incoming event
	stalled := time.After(timeout)

	for {
		select {
		case e := <-events:
			switch e.Event {
			case signal:
				return nil
			case blockJobError:
				return fmt.Errorf("block job error: %v", e.Data)
			}
		case <-stalled:
			return errors.New("block job timeout")
		}
	}
}
