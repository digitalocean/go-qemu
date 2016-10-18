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

package raw

// Generated using `go generate`, do not edit by hand!

import (
	"encoding/json"
	"fmt"
)

// ACPIOSTInfo -> ACPIOSTInfo (struct)

// ACPIOSTInfo implements the "ACPIOSTInfo" QMP API type.
type ACPIOSTInfo struct {
	Device   *string      `json:"device,omitempty"`
	Slot     string       `json:"slot"`
	SlotType ACPISlotType `json:"slot-type"`
	Source   int64        `json:"source"`
	Status   int64        `json:"status"`
}

// ACPISlotType -> ACPISlotType (enum)

// ACPISlotType implements the "ACPISlotType" QMP API type.
type ACPISlotType int

// Known values of ACPISlotType.
const (
	ACPISlotTypeDimm ACPISlotType = iota
	ACPISlotTypeCPU
)

// String implements fmt.Stringer.
func (e ACPISlotType) String() string {
	switch e {
	case ACPISlotTypeDimm:
		return "DIMM"
	case ACPISlotTypeCPU:
		return "CPU"
	default:
		return fmt.Sprintf("ACPISlotType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ACPISlotType) MarshalJSON() ([]byte, error) {
	switch e {
	case ACPISlotTypeDimm:
		return json.Marshal("DIMM")
	case ACPISlotTypeCPU:
		return json.Marshal("CPU")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ACPISlotType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ACPISlotType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "DIMM":
		*e = ACPISlotTypeDimm
	case "CPU":
		*e = ACPISlotTypeCPU
	default:
		return fmt.Errorf("unknown enum value %q for ACPISlotType", s)
	}
	return nil
}

// EVENT ACPI_DEVICE_OST

// Abort -> Abort (struct)

// Abort implements the "Abort" QMP API type.
type Abort struct {
}

// ActionCompletionMode -> ActionCompletionMode (enum)

// ActionCompletionMode implements the "ActionCompletionMode" QMP API type.
type ActionCompletionMode int

// Known values of ActionCompletionMode.
const (
	ActionCompletionModeIndividual ActionCompletionMode = iota
	ActionCompletionModeGrouped
)

// String implements fmt.Stringer.
func (e ActionCompletionMode) String() string {
	switch e {
	case ActionCompletionModeIndividual:
		return "individual"
	case ActionCompletionModeGrouped:
		return "grouped"
	default:
		return fmt.Sprintf("ActionCompletionMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ActionCompletionMode) MarshalJSON() ([]byte, error) {
	switch e {
	case ActionCompletionModeIndividual:
		return json.Marshal("individual")
	case ActionCompletionModeGrouped:
		return json.Marshal("grouped")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ActionCompletionMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ActionCompletionMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "individual":
		*e = ActionCompletionModeIndividual
	case "grouped":
		*e = ActionCompletionModeGrouped
	default:
		return fmt.Errorf("unknown enum value %q for ActionCompletionMode", s)
	}
	return nil
}

// AddfdInfo -> AddfdInfo (struct)

// AddfdInfo implements the "AddfdInfo" QMP API type.
type AddfdInfo struct {
	FdsetID int64 `json:"fdset-id"`
	FD      int64 `json:"fd"`
}

// EVENT BALLOON_CHANGE

// EVENT BLOCK_IMAGE_CORRUPTED

// EVENT BLOCK_IO_ERROR

// EVENT BLOCK_JOB_CANCELLED

// EVENT BLOCK_JOB_COMPLETED

// EVENT BLOCK_JOB_ERROR

// EVENT BLOCK_JOB_READY

// EVENT BLOCK_WRITE_THRESHOLD

// BalloonInfo -> BalloonInfo (struct)

// BalloonInfo implements the "BalloonInfo" QMP API type.
type BalloonInfo struct {
	Actual int64 `json:"actual"`
}

// BlkdebugEvent -> BlkdebugEvent (enum)

// BlkdebugEvent implements the "BlkdebugEvent" QMP API type.
type BlkdebugEvent int

// Known values of BlkdebugEvent.
const (
	BlkdebugEventL1Update BlkdebugEvent = iota
	BlkdebugEventL1GrowAllocTable
	BlkdebugEventL1GrowWriteTable
	BlkdebugEventL1GrowActivateTable
	BlkdebugEventL2Load
	BlkdebugEventL2Update
	BlkdebugEventL2UpdateCompressed
	BlkdebugEventL2AllocCowRead
	BlkdebugEventL2AllocWrite
	BlkdebugEventReadAio
	BlkdebugEventReadBackingAio
	BlkdebugEventReadCompressed
	BlkdebugEventWriteAio
	BlkdebugEventWriteCompressed
	BlkdebugEventVmstateLoad
	BlkdebugEventVmstateSave
	BlkdebugEventCowRead
	BlkdebugEventCowWrite
	BlkdebugEventReftableLoad
	BlkdebugEventReftableGrow
	BlkdebugEventReftableUpdate
	BlkdebugEventRefblockLoad
	BlkdebugEventRefblockUpdate
	BlkdebugEventRefblockUpdatePart
	BlkdebugEventRefblockAlloc
	BlkdebugEventRefblockAllocHookup
	BlkdebugEventRefblockAllocWrite
	BlkdebugEventRefblockAllocWriteBlocks
	BlkdebugEventRefblockAllocWriteTable
	BlkdebugEventRefblockAllocSwitchTable
	BlkdebugEventClusterAlloc
	BlkdebugEventClusterAllocBytes
	BlkdebugEventClusterFree
	BlkdebugEventFlushToOs
	BlkdebugEventFlushToDisk
	BlkdebugEventPwritevRmwHead
	BlkdebugEventPwritevRmwAfterHead
	BlkdebugEventPwritevRmwTail
	BlkdebugEventPwritevRmwAfterTail
	BlkdebugEventPwritev
	BlkdebugEventPwritevZero
	BlkdebugEventPwritevDone
	BlkdebugEventEmptyImagePrepare
)

// String implements fmt.Stringer.
func (e BlkdebugEvent) String() string {
	switch e {
	case BlkdebugEventL1Update:
		return "l1_update"
	case BlkdebugEventL1GrowAllocTable:
		return "l1_grow_alloc_table"
	case BlkdebugEventL1GrowWriteTable:
		return "l1_grow_write_table"
	case BlkdebugEventL1GrowActivateTable:
		return "l1_grow_activate_table"
	case BlkdebugEventL2Load:
		return "l2_load"
	case BlkdebugEventL2Update:
		return "l2_update"
	case BlkdebugEventL2UpdateCompressed:
		return "l2_update_compressed"
	case BlkdebugEventL2AllocCowRead:
		return "l2_alloc_cow_read"
	case BlkdebugEventL2AllocWrite:
		return "l2_alloc_write"
	case BlkdebugEventReadAio:
		return "read_aio"
	case BlkdebugEventReadBackingAio:
		return "read_backing_aio"
	case BlkdebugEventReadCompressed:
		return "read_compressed"
	case BlkdebugEventWriteAio:
		return "write_aio"
	case BlkdebugEventWriteCompressed:
		return "write_compressed"
	case BlkdebugEventVmstateLoad:
		return "vmstate_load"
	case BlkdebugEventVmstateSave:
		return "vmstate_save"
	case BlkdebugEventCowRead:
		return "cow_read"
	case BlkdebugEventCowWrite:
		return "cow_write"
	case BlkdebugEventReftableLoad:
		return "reftable_load"
	case BlkdebugEventReftableGrow:
		return "reftable_grow"
	case BlkdebugEventReftableUpdate:
		return "reftable_update"
	case BlkdebugEventRefblockLoad:
		return "refblock_load"
	case BlkdebugEventRefblockUpdate:
		return "refblock_update"
	case BlkdebugEventRefblockUpdatePart:
		return "refblock_update_part"
	case BlkdebugEventRefblockAlloc:
		return "refblock_alloc"
	case BlkdebugEventRefblockAllocHookup:
		return "refblock_alloc_hookup"
	case BlkdebugEventRefblockAllocWrite:
		return "refblock_alloc_write"
	case BlkdebugEventRefblockAllocWriteBlocks:
		return "refblock_alloc_write_blocks"
	case BlkdebugEventRefblockAllocWriteTable:
		return "refblock_alloc_write_table"
	case BlkdebugEventRefblockAllocSwitchTable:
		return "refblock_alloc_switch_table"
	case BlkdebugEventClusterAlloc:
		return "cluster_alloc"
	case BlkdebugEventClusterAllocBytes:
		return "cluster_alloc_bytes"
	case BlkdebugEventClusterFree:
		return "cluster_free"
	case BlkdebugEventFlushToOs:
		return "flush_to_os"
	case BlkdebugEventFlushToDisk:
		return "flush_to_disk"
	case BlkdebugEventPwritevRmwHead:
		return "pwritev_rmw_head"
	case BlkdebugEventPwritevRmwAfterHead:
		return "pwritev_rmw_after_head"
	case BlkdebugEventPwritevRmwTail:
		return "pwritev_rmw_tail"
	case BlkdebugEventPwritevRmwAfterTail:
		return "pwritev_rmw_after_tail"
	case BlkdebugEventPwritev:
		return "pwritev"
	case BlkdebugEventPwritevZero:
		return "pwritev_zero"
	case BlkdebugEventPwritevDone:
		return "pwritev_done"
	case BlkdebugEventEmptyImagePrepare:
		return "empty_image_prepare"
	default:
		return fmt.Sprintf("BlkdebugEvent(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlkdebugEvent) MarshalJSON() ([]byte, error) {
	switch e {
	case BlkdebugEventL1Update:
		return json.Marshal("l1_update")
	case BlkdebugEventL1GrowAllocTable:
		return json.Marshal("l1_grow_alloc_table")
	case BlkdebugEventL1GrowWriteTable:
		return json.Marshal("l1_grow_write_table")
	case BlkdebugEventL1GrowActivateTable:
		return json.Marshal("l1_grow_activate_table")
	case BlkdebugEventL2Load:
		return json.Marshal("l2_load")
	case BlkdebugEventL2Update:
		return json.Marshal("l2_update")
	case BlkdebugEventL2UpdateCompressed:
		return json.Marshal("l2_update_compressed")
	case BlkdebugEventL2AllocCowRead:
		return json.Marshal("l2_alloc_cow_read")
	case BlkdebugEventL2AllocWrite:
		return json.Marshal("l2_alloc_write")
	case BlkdebugEventReadAio:
		return json.Marshal("read_aio")
	case BlkdebugEventReadBackingAio:
		return json.Marshal("read_backing_aio")
	case BlkdebugEventReadCompressed:
		return json.Marshal("read_compressed")
	case BlkdebugEventWriteAio:
		return json.Marshal("write_aio")
	case BlkdebugEventWriteCompressed:
		return json.Marshal("write_compressed")
	case BlkdebugEventVmstateLoad:
		return json.Marshal("vmstate_load")
	case BlkdebugEventVmstateSave:
		return json.Marshal("vmstate_save")
	case BlkdebugEventCowRead:
		return json.Marshal("cow_read")
	case BlkdebugEventCowWrite:
		return json.Marshal("cow_write")
	case BlkdebugEventReftableLoad:
		return json.Marshal("reftable_load")
	case BlkdebugEventReftableGrow:
		return json.Marshal("reftable_grow")
	case BlkdebugEventReftableUpdate:
		return json.Marshal("reftable_update")
	case BlkdebugEventRefblockLoad:
		return json.Marshal("refblock_load")
	case BlkdebugEventRefblockUpdate:
		return json.Marshal("refblock_update")
	case BlkdebugEventRefblockUpdatePart:
		return json.Marshal("refblock_update_part")
	case BlkdebugEventRefblockAlloc:
		return json.Marshal("refblock_alloc")
	case BlkdebugEventRefblockAllocHookup:
		return json.Marshal("refblock_alloc_hookup")
	case BlkdebugEventRefblockAllocWrite:
		return json.Marshal("refblock_alloc_write")
	case BlkdebugEventRefblockAllocWriteBlocks:
		return json.Marshal("refblock_alloc_write_blocks")
	case BlkdebugEventRefblockAllocWriteTable:
		return json.Marshal("refblock_alloc_write_table")
	case BlkdebugEventRefblockAllocSwitchTable:
		return json.Marshal("refblock_alloc_switch_table")
	case BlkdebugEventClusterAlloc:
		return json.Marshal("cluster_alloc")
	case BlkdebugEventClusterAllocBytes:
		return json.Marshal("cluster_alloc_bytes")
	case BlkdebugEventClusterFree:
		return json.Marshal("cluster_free")
	case BlkdebugEventFlushToOs:
		return json.Marshal("flush_to_os")
	case BlkdebugEventFlushToDisk:
		return json.Marshal("flush_to_disk")
	case BlkdebugEventPwritevRmwHead:
		return json.Marshal("pwritev_rmw_head")
	case BlkdebugEventPwritevRmwAfterHead:
		return json.Marshal("pwritev_rmw_after_head")
	case BlkdebugEventPwritevRmwTail:
		return json.Marshal("pwritev_rmw_tail")
	case BlkdebugEventPwritevRmwAfterTail:
		return json.Marshal("pwritev_rmw_after_tail")
	case BlkdebugEventPwritev:
		return json.Marshal("pwritev")
	case BlkdebugEventPwritevZero:
		return json.Marshal("pwritev_zero")
	case BlkdebugEventPwritevDone:
		return json.Marshal("pwritev_done")
	case BlkdebugEventEmptyImagePrepare:
		return json.Marshal("empty_image_prepare")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlkdebugEvent", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlkdebugEvent) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "l1_update":
		*e = BlkdebugEventL1Update
	case "l1_grow_alloc_table":
		*e = BlkdebugEventL1GrowAllocTable
	case "l1_grow_write_table":
		*e = BlkdebugEventL1GrowWriteTable
	case "l1_grow_activate_table":
		*e = BlkdebugEventL1GrowActivateTable
	case "l2_load":
		*e = BlkdebugEventL2Load
	case "l2_update":
		*e = BlkdebugEventL2Update
	case "l2_update_compressed":
		*e = BlkdebugEventL2UpdateCompressed
	case "l2_alloc_cow_read":
		*e = BlkdebugEventL2AllocCowRead
	case "l2_alloc_write":
		*e = BlkdebugEventL2AllocWrite
	case "read_aio":
		*e = BlkdebugEventReadAio
	case "read_backing_aio":
		*e = BlkdebugEventReadBackingAio
	case "read_compressed":
		*e = BlkdebugEventReadCompressed
	case "write_aio":
		*e = BlkdebugEventWriteAio
	case "write_compressed":
		*e = BlkdebugEventWriteCompressed
	case "vmstate_load":
		*e = BlkdebugEventVmstateLoad
	case "vmstate_save":
		*e = BlkdebugEventVmstateSave
	case "cow_read":
		*e = BlkdebugEventCowRead
	case "cow_write":
		*e = BlkdebugEventCowWrite
	case "reftable_load":
		*e = BlkdebugEventReftableLoad
	case "reftable_grow":
		*e = BlkdebugEventReftableGrow
	case "reftable_update":
		*e = BlkdebugEventReftableUpdate
	case "refblock_load":
		*e = BlkdebugEventRefblockLoad
	case "refblock_update":
		*e = BlkdebugEventRefblockUpdate
	case "refblock_update_part":
		*e = BlkdebugEventRefblockUpdatePart
	case "refblock_alloc":
		*e = BlkdebugEventRefblockAlloc
	case "refblock_alloc_hookup":
		*e = BlkdebugEventRefblockAllocHookup
	case "refblock_alloc_write":
		*e = BlkdebugEventRefblockAllocWrite
	case "refblock_alloc_write_blocks":
		*e = BlkdebugEventRefblockAllocWriteBlocks
	case "refblock_alloc_write_table":
		*e = BlkdebugEventRefblockAllocWriteTable
	case "refblock_alloc_switch_table":
		*e = BlkdebugEventRefblockAllocSwitchTable
	case "cluster_alloc":
		*e = BlkdebugEventClusterAlloc
	case "cluster_alloc_bytes":
		*e = BlkdebugEventClusterAllocBytes
	case "cluster_free":
		*e = BlkdebugEventClusterFree
	case "flush_to_os":
		*e = BlkdebugEventFlushToOs
	case "flush_to_disk":
		*e = BlkdebugEventFlushToDisk
	case "pwritev_rmw_head":
		*e = BlkdebugEventPwritevRmwHead
	case "pwritev_rmw_after_head":
		*e = BlkdebugEventPwritevRmwAfterHead
	case "pwritev_rmw_tail":
		*e = BlkdebugEventPwritevRmwTail
	case "pwritev_rmw_after_tail":
		*e = BlkdebugEventPwritevRmwAfterTail
	case "pwritev":
		*e = BlkdebugEventPwritev
	case "pwritev_zero":
		*e = BlkdebugEventPwritevZero
	case "pwritev_done":
		*e = BlkdebugEventPwritevDone
	case "empty_image_prepare":
		*e = BlkdebugEventEmptyImagePrepare
	default:
		return fmt.Errorf("unknown enum value %q for BlkdebugEvent", s)
	}
	return nil
}

// BlkdebugInjectErrorOptions -> BlkdebugInjectErrorOptions (struct)

// BlkdebugInjectErrorOptions implements the "BlkdebugInjectErrorOptions" QMP API type.
type BlkdebugInjectErrorOptions struct {
	Event       BlkdebugEvent `json:"event"`
	State       *int64        `json:"state,omitempty"`
	Errno       *int64        `json:"errno,omitempty"`
	Sector      *int64        `json:"sector,omitempty"`
	Once        *bool         `json:"once,omitempty"`
	Immediately *bool         `json:"immediately,omitempty"`
}

// BlkdebugSetStateOptions -> BlkdebugSetStateOptions (struct)

// BlkdebugSetStateOptions implements the "BlkdebugSetStateOptions" QMP API type.
type BlkdebugSetStateOptions struct {
	Event    BlkdebugEvent `json:"event"`
	State    *int64        `json:"state,omitempty"`
	NewState int64         `json:"new_state"`
}

// BlockDeviceInfo -> BlockDeviceInfo (struct)

// BlockDeviceInfo implements the "BlockDeviceInfo" QMP API type.
type BlockDeviceInfo struct {
	File                 string                      `json:"file"`
	NodeName             *string                     `json:"node-name,omitempty"`
	Ro                   bool                        `json:"ro"`
	Drv                  string                      `json:"drv"`
	BackingFile          *string                     `json:"backing_file,omitempty"`
	BackingFileDepth     int64                       `json:"backing_file_depth"`
	Encrypted            bool                        `json:"encrypted"`
	EncryptionKeyMissing bool                        `json:"encryption_key_missing"`
	DetectZeroes         BlockdevDetectZeroesOptions `json:"detect_zeroes"`
	Bps                  int64                       `json:"bps"`
	BpsRd                int64                       `json:"bps_rd"`
	BpsWr                int64                       `json:"bps_wr"`
	Iops                 int64                       `json:"iops"`
	IopsRd               int64                       `json:"iops_rd"`
	IopsWr               int64                       `json:"iops_wr"`
	Image                ImageInfo                   `json:"image"`
	BpsMax               *int64                      `json:"bps_max,omitempty"`
	BpsRdMax             *int64                      `json:"bps_rd_max,omitempty"`
	BpsWrMax             *int64                      `json:"bps_wr_max,omitempty"`
	IopsMax              *int64                      `json:"iops_max,omitempty"`
	IopsRdMax            *int64                      `json:"iops_rd_max,omitempty"`
	IopsWrMax            *int64                      `json:"iops_wr_max,omitempty"`
	BpsMaxLength         *int64                      `json:"bps_max_length,omitempty"`
	BpsRdMaxLength       *int64                      `json:"bps_rd_max_length,omitempty"`
	BpsWrMaxLength       *int64                      `json:"bps_wr_max_length,omitempty"`
	IopsMaxLength        *int64                      `json:"iops_max_length,omitempty"`
	IopsRdMaxLength      *int64                      `json:"iops_rd_max_length,omitempty"`
	IopsWrMaxLength      *int64                      `json:"iops_wr_max_length,omitempty"`
	IopsSize             *int64                      `json:"iops_size,omitempty"`
	Group                *string                     `json:"group,omitempty"`
	Cache                BlockdevCacheInfo           `json:"cache"`
	WriteThreshold       int64                       `json:"write_threshold"`
}

// BlockDeviceIoStatus -> BlockDeviceIOStatus (enum)

// BlockDeviceIOStatus implements the "BlockDeviceIoStatus" QMP API type.
type BlockDeviceIOStatus int

// Known values of BlockDeviceIOStatus.
const (
	BlockDeviceIOStatusOK BlockDeviceIOStatus = iota
	BlockDeviceIOStatusFailed
	BlockDeviceIOStatusNospace
)

// String implements fmt.Stringer.
func (e BlockDeviceIOStatus) String() string {
	switch e {
	case BlockDeviceIOStatusOK:
		return "ok"
	case BlockDeviceIOStatusFailed:
		return "failed"
	case BlockDeviceIOStatusNospace:
		return "nospace"
	default:
		return fmt.Sprintf("BlockDeviceIOStatus(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockDeviceIOStatus) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockDeviceIOStatusOK:
		return json.Marshal("ok")
	case BlockDeviceIOStatusFailed:
		return json.Marshal("failed")
	case BlockDeviceIOStatusNospace:
		return json.Marshal("nospace")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockDeviceIOStatus", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockDeviceIOStatus) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "ok":
		*e = BlockDeviceIOStatusOK
	case "failed":
		*e = BlockDeviceIOStatusFailed
	case "nospace":
		*e = BlockDeviceIOStatusNospace
	default:
		return fmt.Errorf("unknown enum value %q for BlockDeviceIOStatus", s)
	}
	return nil
}

// BlockDeviceStats -> BlockDeviceStats (struct)

// BlockDeviceStats implements the "BlockDeviceStats" QMP API type.
type BlockDeviceStats struct {
	RdBytes                int64                   `json:"rd_bytes"`
	WrBytes                int64                   `json:"wr_bytes"`
	RdOperations           int64                   `json:"rd_operations"`
	WrOperations           int64                   `json:"wr_operations"`
	FlushOperations        int64                   `json:"flush_operations"`
	FlushTotalTimeNs       int64                   `json:"flush_total_time_ns"`
	WrTotalTimeNs          int64                   `json:"wr_total_time_ns"`
	RdTotalTimeNs          int64                   `json:"rd_total_time_ns"`
	WrHighestOffset        int64                   `json:"wr_highest_offset"`
	RdMerged               int64                   `json:"rd_merged"`
	WrMerged               int64                   `json:"wr_merged"`
	IdleTimeNs             *int64                  `json:"idle_time_ns,omitempty"`
	FailedRdOperations     int64                   `json:"failed_rd_operations"`
	FailedWrOperations     int64                   `json:"failed_wr_operations"`
	FailedFlushOperations  int64                   `json:"failed_flush_operations"`
	InvalidRdOperations    int64                   `json:"invalid_rd_operations"`
	InvalidWrOperations    int64                   `json:"invalid_wr_operations"`
	InvalidFlushOperations int64                   `json:"invalid_flush_operations"`
	AccountInvalid         bool                    `json:"account_invalid"`
	AccountFailed          bool                    `json:"account_failed"`
	TimedStats             []BlockDeviceTimedStats `json:"timed_stats"`
}

// BlockDeviceTimedStats -> BlockDeviceTimedStats (struct)

// BlockDeviceTimedStats implements the "BlockDeviceTimedStats" QMP API type.
type BlockDeviceTimedStats struct {
	IntervalLength    int64   `json:"interval_length"`
	MinRdLatencyNs    int64   `json:"min_rd_latency_ns"`
	MaxRdLatencyNs    int64   `json:"max_rd_latency_ns"`
	AvgRdLatencyNs    int64   `json:"avg_rd_latency_ns"`
	MinWrLatencyNs    int64   `json:"min_wr_latency_ns"`
	MaxWrLatencyNs    int64   `json:"max_wr_latency_ns"`
	AvgWrLatencyNs    int64   `json:"avg_wr_latency_ns"`
	MinFlushLatencyNs int64   `json:"min_flush_latency_ns"`
	MaxFlushLatencyNs int64   `json:"max_flush_latency_ns"`
	AvgFlushLatencyNs int64   `json:"avg_flush_latency_ns"`
	AvgRdQueueDepth   float64 `json:"avg_rd_queue_depth"`
	AvgWrQueueDepth   float64 `json:"avg_wr_queue_depth"`
}

// BlockDirtyBitmap -> BlockDirtyBitmap (struct)

// BlockDirtyBitmap implements the "BlockDirtyBitmap" QMP API type.
type BlockDirtyBitmap struct {
	Node string `json:"node"`
	Name string `json:"name"`
}

// BlockDirtyBitmapAdd -> BlockDirtyBitmapAdd (struct)

// BlockDirtyBitmapAdd implements the "BlockDirtyBitmapAdd" QMP API type.
type BlockDirtyBitmapAdd struct {
	Node        string  `json:"node"`
	Name        string  `json:"name"`
	Granularity *uint32 `json:"granularity,omitempty"`
}

// BlockDirtyInfo -> BlockDirtyInfo (struct)

// BlockDirtyInfo implements the "BlockDirtyInfo" QMP API type.
type BlockDirtyInfo struct {
	Name        *string           `json:"name,omitempty"`
	Count       int64             `json:"count"`
	Granularity uint32            `json:"granularity"`
	Status      DirtyBitmapStatus `json:"status"`
}

// BlockErrorAction -> BlockErrorAction (enum)

// BlockErrorAction implements the "BlockErrorAction" QMP API type.
type BlockErrorAction int

// Known values of BlockErrorAction.
const (
	BlockErrorActionIgnore BlockErrorAction = iota
	BlockErrorActionReport
	BlockErrorActionStop
)

// String implements fmt.Stringer.
func (e BlockErrorAction) String() string {
	switch e {
	case BlockErrorActionIgnore:
		return "ignore"
	case BlockErrorActionReport:
		return "report"
	case BlockErrorActionStop:
		return "stop"
	default:
		return fmt.Sprintf("BlockErrorAction(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockErrorAction) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockErrorActionIgnore:
		return json.Marshal("ignore")
	case BlockErrorActionReport:
		return json.Marshal("report")
	case BlockErrorActionStop:
		return json.Marshal("stop")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockErrorAction", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockErrorAction) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "ignore":
		*e = BlockErrorActionIgnore
	case "report":
		*e = BlockErrorActionReport
	case "stop":
		*e = BlockErrorActionStop
	default:
		return fmt.Errorf("unknown enum value %q for BlockErrorAction", s)
	}
	return nil
}

// BlockIOThrottle -> BlockIOThrottle (struct)

// BlockIOThrottle implements the "BlockIOThrottle" QMP API type.
type BlockIOThrottle struct {
	Device          *string `json:"device,omitempty"`
	ID              *string `json:"id,omitempty"`
	Bps             int64   `json:"bps"`
	BpsRd           int64   `json:"bps_rd"`
	BpsWr           int64   `json:"bps_wr"`
	Iops            int64   `json:"iops"`
	IopsRd          int64   `json:"iops_rd"`
	IopsWr          int64   `json:"iops_wr"`
	BpsMax          *int64  `json:"bps_max,omitempty"`
	BpsRdMax        *int64  `json:"bps_rd_max,omitempty"`
	BpsWrMax        *int64  `json:"bps_wr_max,omitempty"`
	IopsMax         *int64  `json:"iops_max,omitempty"`
	IopsRdMax       *int64  `json:"iops_rd_max,omitempty"`
	IopsWrMax       *int64  `json:"iops_wr_max,omitempty"`
	BpsMaxLength    *int64  `json:"bps_max_length,omitempty"`
	BpsRdMaxLength  *int64  `json:"bps_rd_max_length,omitempty"`
	BpsWrMaxLength  *int64  `json:"bps_wr_max_length,omitempty"`
	IopsMaxLength   *int64  `json:"iops_max_length,omitempty"`
	IopsRdMaxLength *int64  `json:"iops_rd_max_length,omitempty"`
	IopsWrMaxLength *int64  `json:"iops_wr_max_length,omitempty"`
	IopsSize        *int64  `json:"iops_size,omitempty"`
	Group           *string `json:"group,omitempty"`
}

// BlockInfo -> BlockInfo (struct)

// BlockInfo implements the "BlockInfo" QMP API type.
type BlockInfo struct {
	Device       string               `json:"device"`
	Type         string               `json:"type"`
	Removable    bool                 `json:"removable"`
	Locked       bool                 `json:"locked"`
	Inserted     *BlockDeviceInfo     `json:"inserted,omitempty"`
	TrayOpen     *bool                `json:"tray_open,omitempty"`
	IOStatus     *BlockDeviceIOStatus `json:"io-status,omitempty"`
	DirtyBitmaps []BlockDirtyInfo     `json:"dirty-bitmaps,omitempty"`
}

// BlockJobInfo -> BlockJobInfo (struct)

// BlockJobInfo implements the "BlockJobInfo" QMP API type.
type BlockJobInfo struct {
	Type     string              `json:"type"`
	Device   string              `json:"device"`
	Len      int64               `json:"len"`
	Offset   int64               `json:"offset"`
	Busy     bool                `json:"busy"`
	Paused   bool                `json:"paused"`
	Speed    int64               `json:"speed"`
	IOStatus BlockDeviceIOStatus `json:"io-status"`
	Ready    bool                `json:"ready"`
}

// BlockJobType -> BlockJobType (enum)

// BlockJobType implements the "BlockJobType" QMP API type.
type BlockJobType int

// Known values of BlockJobType.
const (
	BlockJobTypeCommit BlockJobType = iota
	BlockJobTypeStream
	BlockJobTypeMirror
	BlockJobTypeBackup
)

// String implements fmt.Stringer.
func (e BlockJobType) String() string {
	switch e {
	case BlockJobTypeCommit:
		return "commit"
	case BlockJobTypeStream:
		return "stream"
	case BlockJobTypeMirror:
		return "mirror"
	case BlockJobTypeBackup:
		return "backup"
	default:
		return fmt.Sprintf("BlockJobType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockJobType) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockJobTypeCommit:
		return json.Marshal("commit")
	case BlockJobTypeStream:
		return json.Marshal("stream")
	case BlockJobTypeMirror:
		return json.Marshal("mirror")
	case BlockJobTypeBackup:
		return json.Marshal("backup")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockJobType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockJobType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "commit":
		*e = BlockJobTypeCommit
	case "stream":
		*e = BlockJobTypeStream
	case "mirror":
		*e = BlockJobTypeMirror
	case "backup":
		*e = BlockJobTypeBackup
	default:
		return fmt.Errorf("unknown enum value %q for BlockJobType", s)
	}
	return nil
}

// BlockStats -> BlockStats (struct)

// BlockStats implements the "BlockStats" QMP API type.
type BlockStats struct {
	Device   *string          `json:"device,omitempty"`
	NodeName *string          `json:"node-name,omitempty"`
	Stats    BlockDeviceStats `json:"stats"`
	Parent   *BlockStats      `json:"parent,omitempty"`
	Backing  *BlockStats      `json:"backing,omitempty"`
}

// BlockdevAioOptions -> BlockdevAioOptions (enum)

// BlockdevAioOptions implements the "BlockdevAioOptions" QMP API type.
type BlockdevAioOptions int

// Known values of BlockdevAioOptions.
const (
	BlockdevAioOptionsThreads BlockdevAioOptions = iota
	BlockdevAioOptionsNative
)

// String implements fmt.Stringer.
func (e BlockdevAioOptions) String() string {
	switch e {
	case BlockdevAioOptionsThreads:
		return "threads"
	case BlockdevAioOptionsNative:
		return "native"
	default:
		return fmt.Sprintf("BlockdevAioOptions(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevAioOptions) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevAioOptionsThreads:
		return json.Marshal("threads")
	case BlockdevAioOptionsNative:
		return json.Marshal("native")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevAioOptions", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevAioOptions) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "threads":
		*e = BlockdevAioOptionsThreads
	case "native":
		*e = BlockdevAioOptionsNative
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevAioOptions", s)
	}
	return nil
}

// BlockdevBackup -> BlockdevBackup (struct)

// BlockdevBackup implements the "BlockdevBackup" QMP API type.
type BlockdevBackup struct {
	JobID         *string          `json:"job-id,omitempty"`
	Device        string           `json:"device"`
	Target        string           `json:"target"`
	Sync          MirrorSyncMode   `json:"sync"`
	Speed         *int64           `json:"speed,omitempty"`
	Compress      *bool            `json:"compress,omitempty"`
	OnSourceError *BlockdevOnError `json:"on-source-error,omitempty"`
	OnTargetError *BlockdevOnError `json:"on-target-error,omitempty"`
}

// BlockdevCacheInfo -> BlockdevCacheInfo (struct)

// BlockdevCacheInfo implements the "BlockdevCacheInfo" QMP API type.
type BlockdevCacheInfo struct {
	Writeback bool `json:"writeback"`
	Direct    bool `json:"direct"`
	NoFlush   bool `json:"no-flush"`
}

// BlockdevCacheOptions -> BlockdevCacheOptions (struct)

// BlockdevCacheOptions implements the "BlockdevCacheOptions" QMP API type.
type BlockdevCacheOptions struct {
	Direct  *bool `json:"direct,omitempty"`
	NoFlush *bool `json:"no-flush,omitempty"`
}

// BlockdevChangeReadOnlyMode -> BlockdevChangeReadOnlyMode (enum)

// BlockdevChangeReadOnlyMode implements the "BlockdevChangeReadOnlyMode" QMP API type.
type BlockdevChangeReadOnlyMode int

// Known values of BlockdevChangeReadOnlyMode.
const (
	BlockdevChangeReadOnlyModeRetain BlockdevChangeReadOnlyMode = iota
	BlockdevChangeReadOnlyModeReadOnly
	BlockdevChangeReadOnlyModeReadWrite
)

// String implements fmt.Stringer.
func (e BlockdevChangeReadOnlyMode) String() string {
	switch e {
	case BlockdevChangeReadOnlyModeRetain:
		return "retain"
	case BlockdevChangeReadOnlyModeReadOnly:
		return "read-only"
	case BlockdevChangeReadOnlyModeReadWrite:
		return "read-write"
	default:
		return fmt.Sprintf("BlockdevChangeReadOnlyMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevChangeReadOnlyMode) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevChangeReadOnlyModeRetain:
		return json.Marshal("retain")
	case BlockdevChangeReadOnlyModeReadOnly:
		return json.Marshal("read-only")
	case BlockdevChangeReadOnlyModeReadWrite:
		return json.Marshal("read-write")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevChangeReadOnlyMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevChangeReadOnlyMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "retain":
		*e = BlockdevChangeReadOnlyModeRetain
	case "read-only":
		*e = BlockdevChangeReadOnlyModeReadOnly
	case "read-write":
		*e = BlockdevChangeReadOnlyModeReadWrite
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevChangeReadOnlyMode", s)
	}
	return nil
}

// BlockdevDetectZeroesOptions -> BlockdevDetectZeroesOptions (enum)

// BlockdevDetectZeroesOptions implements the "BlockdevDetectZeroesOptions" QMP API type.
type BlockdevDetectZeroesOptions int

// Known values of BlockdevDetectZeroesOptions.
const (
	BlockdevDetectZeroesOptionsOff BlockdevDetectZeroesOptions = iota
	BlockdevDetectZeroesOptionsOn
	BlockdevDetectZeroesOptionsUnmap
)

// String implements fmt.Stringer.
func (e BlockdevDetectZeroesOptions) String() string {
	switch e {
	case BlockdevDetectZeroesOptionsOff:
		return "off"
	case BlockdevDetectZeroesOptionsOn:
		return "on"
	case BlockdevDetectZeroesOptionsUnmap:
		return "unmap"
	default:
		return fmt.Sprintf("BlockdevDetectZeroesOptions(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevDetectZeroesOptions) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevDetectZeroesOptionsOff:
		return json.Marshal("off")
	case BlockdevDetectZeroesOptionsOn:
		return json.Marshal("on")
	case BlockdevDetectZeroesOptionsUnmap:
		return json.Marshal("unmap")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevDetectZeroesOptions", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevDetectZeroesOptions) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "off":
		*e = BlockdevDetectZeroesOptionsOff
	case "on":
		*e = BlockdevDetectZeroesOptionsOn
	case "unmap":
		*e = BlockdevDetectZeroesOptionsUnmap
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevDetectZeroesOptions", s)
	}
	return nil
}

// BlockdevDiscardOptions -> BlockdevDiscardOptions (enum)

// BlockdevDiscardOptions implements the "BlockdevDiscardOptions" QMP API type.
type BlockdevDiscardOptions int

// Known values of BlockdevDiscardOptions.
const (
	BlockdevDiscardOptionsIgnore BlockdevDiscardOptions = iota
	BlockdevDiscardOptionsUnmap
)

// String implements fmt.Stringer.
func (e BlockdevDiscardOptions) String() string {
	switch e {
	case BlockdevDiscardOptionsIgnore:
		return "ignore"
	case BlockdevDiscardOptionsUnmap:
		return "unmap"
	default:
		return fmt.Sprintf("BlockdevDiscardOptions(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevDiscardOptions) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevDiscardOptionsIgnore:
		return json.Marshal("ignore")
	case BlockdevDiscardOptionsUnmap:
		return json.Marshal("unmap")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevDiscardOptions", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevDiscardOptions) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "ignore":
		*e = BlockdevDiscardOptionsIgnore
	case "unmap":
		*e = BlockdevDiscardOptionsUnmap
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevDiscardOptions", s)
	}
	return nil
}

// BlockdevDriver -> BlockdevDriver (enum)

// BlockdevDriver implements the "BlockdevDriver" QMP API type.
type BlockdevDriver int

// Known values of BlockdevDriver.
const (
	BlockdevDriverArchipelago BlockdevDriver = iota
	BlockdevDriverBlkdebug
	BlockdevDriverBlkverify
	BlockdevDriverBochs
	BlockdevDriverCloop
	BlockdevDriverDmg
	BlockdevDriverFile
	BlockdevDriverFTP
	BlockdevDriverFTPS
	BlockdevDriverGluster
	BlockdevDriverHostCdrom
	BlockdevDriverHostDevice
	BlockdevDriverHTTP
	BlockdevDriverHTTPS
	BlockdevDriverLUKS
	BlockdevDriverNullAio
	BlockdevDriverNullCo
	BlockdevDriverParallels
	BlockdevDriverQcow
	BlockdevDriverQcow2
	BlockdevDriverQed
	BlockdevDriverQuorum
	BlockdevDriverRaw
	BlockdevDriverReplication
	BlockdevDriverTftp
	BlockdevDriverVdi
	BlockdevDriverVhdx
	BlockdevDriverVMDK
	BlockdevDriverVpc
	BlockdevDriverVvfat
)

// String implements fmt.Stringer.
func (e BlockdevDriver) String() string {
	switch e {
	case BlockdevDriverArchipelago:
		return "archipelago"
	case BlockdevDriverBlkdebug:
		return "blkdebug"
	case BlockdevDriverBlkverify:
		return "blkverify"
	case BlockdevDriverBochs:
		return "bochs"
	case BlockdevDriverCloop:
		return "cloop"
	case BlockdevDriverDmg:
		return "dmg"
	case BlockdevDriverFile:
		return "file"
	case BlockdevDriverFTP:
		return "ftp"
	case BlockdevDriverFTPS:
		return "ftps"
	case BlockdevDriverGluster:
		return "gluster"
	case BlockdevDriverHostCdrom:
		return "host_cdrom"
	case BlockdevDriverHostDevice:
		return "host_device"
	case BlockdevDriverHTTP:
		return "http"
	case BlockdevDriverHTTPS:
		return "https"
	case BlockdevDriverLUKS:
		return "luks"
	case BlockdevDriverNullAio:
		return "null-aio"
	case BlockdevDriverNullCo:
		return "null-co"
	case BlockdevDriverParallels:
		return "parallels"
	case BlockdevDriverQcow:
		return "qcow"
	case BlockdevDriverQcow2:
		return "qcow2"
	case BlockdevDriverQed:
		return "qed"
	case BlockdevDriverQuorum:
		return "quorum"
	case BlockdevDriverRaw:
		return "raw"
	case BlockdevDriverReplication:
		return "replication"
	case BlockdevDriverTftp:
		return "tftp"
	case BlockdevDriverVdi:
		return "vdi"
	case BlockdevDriverVhdx:
		return "vhdx"
	case BlockdevDriverVMDK:
		return "vmdk"
	case BlockdevDriverVpc:
		return "vpc"
	case BlockdevDriverVvfat:
		return "vvfat"
	default:
		return fmt.Sprintf("BlockdevDriver(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevDriver) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevDriverArchipelago:
		return json.Marshal("archipelago")
	case BlockdevDriverBlkdebug:
		return json.Marshal("blkdebug")
	case BlockdevDriverBlkverify:
		return json.Marshal("blkverify")
	case BlockdevDriverBochs:
		return json.Marshal("bochs")
	case BlockdevDriverCloop:
		return json.Marshal("cloop")
	case BlockdevDriverDmg:
		return json.Marshal("dmg")
	case BlockdevDriverFile:
		return json.Marshal("file")
	case BlockdevDriverFTP:
		return json.Marshal("ftp")
	case BlockdevDriverFTPS:
		return json.Marshal("ftps")
	case BlockdevDriverGluster:
		return json.Marshal("gluster")
	case BlockdevDriverHostCdrom:
		return json.Marshal("host_cdrom")
	case BlockdevDriverHostDevice:
		return json.Marshal("host_device")
	case BlockdevDriverHTTP:
		return json.Marshal("http")
	case BlockdevDriverHTTPS:
		return json.Marshal("https")
	case BlockdevDriverLUKS:
		return json.Marshal("luks")
	case BlockdevDriverNullAio:
		return json.Marshal("null-aio")
	case BlockdevDriverNullCo:
		return json.Marshal("null-co")
	case BlockdevDriverParallels:
		return json.Marshal("parallels")
	case BlockdevDriverQcow:
		return json.Marshal("qcow")
	case BlockdevDriverQcow2:
		return json.Marshal("qcow2")
	case BlockdevDriverQed:
		return json.Marshal("qed")
	case BlockdevDriverQuorum:
		return json.Marshal("quorum")
	case BlockdevDriverRaw:
		return json.Marshal("raw")
	case BlockdevDriverReplication:
		return json.Marshal("replication")
	case BlockdevDriverTftp:
		return json.Marshal("tftp")
	case BlockdevDriverVdi:
		return json.Marshal("vdi")
	case BlockdevDriverVhdx:
		return json.Marshal("vhdx")
	case BlockdevDriverVMDK:
		return json.Marshal("vmdk")
	case BlockdevDriverVpc:
		return json.Marshal("vpc")
	case BlockdevDriverVvfat:
		return json.Marshal("vvfat")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevDriver", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevDriver) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "archipelago":
		*e = BlockdevDriverArchipelago
	case "blkdebug":
		*e = BlockdevDriverBlkdebug
	case "blkverify":
		*e = BlockdevDriverBlkverify
	case "bochs":
		*e = BlockdevDriverBochs
	case "cloop":
		*e = BlockdevDriverCloop
	case "dmg":
		*e = BlockdevDriverDmg
	case "file":
		*e = BlockdevDriverFile
	case "ftp":
		*e = BlockdevDriverFTP
	case "ftps":
		*e = BlockdevDriverFTPS
	case "gluster":
		*e = BlockdevDriverGluster
	case "host_cdrom":
		*e = BlockdevDriverHostCdrom
	case "host_device":
		*e = BlockdevDriverHostDevice
	case "http":
		*e = BlockdevDriverHTTP
	case "https":
		*e = BlockdevDriverHTTPS
	case "luks":
		*e = BlockdevDriverLUKS
	case "null-aio":
		*e = BlockdevDriverNullAio
	case "null-co":
		*e = BlockdevDriverNullCo
	case "parallels":
		*e = BlockdevDriverParallels
	case "qcow":
		*e = BlockdevDriverQcow
	case "qcow2":
		*e = BlockdevDriverQcow2
	case "qed":
		*e = BlockdevDriverQed
	case "quorum":
		*e = BlockdevDriverQuorum
	case "raw":
		*e = BlockdevDriverRaw
	case "replication":
		*e = BlockdevDriverReplication
	case "tftp":
		*e = BlockdevDriverTftp
	case "vdi":
		*e = BlockdevDriverVdi
	case "vhdx":
		*e = BlockdevDriverVhdx
	case "vmdk":
		*e = BlockdevDriverVMDK
	case "vpc":
		*e = BlockdevDriverVpc
	case "vvfat":
		*e = BlockdevDriverVvfat
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevDriver", s)
	}
	return nil
}

// BlockdevOnError -> BlockdevOnError (enum)

// BlockdevOnError implements the "BlockdevOnError" QMP API type.
type BlockdevOnError int

// Known values of BlockdevOnError.
const (
	BlockdevOnErrorReport BlockdevOnError = iota
	BlockdevOnErrorIgnore
	BlockdevOnErrorEnospc
	BlockdevOnErrorStop
	BlockdevOnErrorAuto
)

// String implements fmt.Stringer.
func (e BlockdevOnError) String() string {
	switch e {
	case BlockdevOnErrorReport:
		return "report"
	case BlockdevOnErrorIgnore:
		return "ignore"
	case BlockdevOnErrorEnospc:
		return "enospc"
	case BlockdevOnErrorStop:
		return "stop"
	case BlockdevOnErrorAuto:
		return "auto"
	default:
		return fmt.Sprintf("BlockdevOnError(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevOnError) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevOnErrorReport:
		return json.Marshal("report")
	case BlockdevOnErrorIgnore:
		return json.Marshal("ignore")
	case BlockdevOnErrorEnospc:
		return json.Marshal("enospc")
	case BlockdevOnErrorStop:
		return json.Marshal("stop")
	case BlockdevOnErrorAuto:
		return json.Marshal("auto")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevOnError", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevOnError) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "report":
		*e = BlockdevOnErrorReport
	case "ignore":
		*e = BlockdevOnErrorIgnore
	case "enospc":
		*e = BlockdevOnErrorEnospc
	case "stop":
		*e = BlockdevOnErrorStop
	case "auto":
		*e = BlockdevOnErrorAuto
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevOnError", s)
	}
	return nil
}

// BlockdevOptions -> BlockdevOptions (flat union)

// BlockdevOptions implements the "BlockdevOptions" QMP API type.
//
// Can be one of:
//   - BlockdevOptionsArchipelago
//   - BlockdevOptionsBlkdebug
//   - BlockdevOptionsBlkverify
//   - BlockdevOptionsBochs
//   - BlockdevOptionsCloop
//   - BlockdevOptionsDmg
//   - BlockdevOptionsFile
//   - BlockdevOptionsFTP
//   - BlockdevOptionsFTPS
//   - BlockdevOptionsGluster
//   - BlockdevOptionsHostCdrom
//   - BlockdevOptionsHostDevice
//   - BlockdevOptionsHTTP
//   - BlockdevOptionsHTTPS
//   - BlockdevOptionsLUKS
//   - BlockdevOptionsNullAio
//   - BlockdevOptionsNullCo
//   - BlockdevOptionsParallels
//   - BlockdevOptionsQcow
//   - BlockdevOptionsQcow2
//   - BlockdevOptionsQed
//   - BlockdevOptionsQuorum
//   - BlockdevOptionsRaw
//   - BlockdevOptionsReplication
//   - BlockdevOptionsTftp
//   - BlockdevOptionsVdi
//   - BlockdevOptionsVhdx
//   - BlockdevOptionsVMDK
//   - BlockdevOptionsVpc
//   - BlockdevOptionsVvfat
type BlockdevOptions interface {
	isBlockdevOptions()
}

// BlockdevOptionsArchipelago is an implementation of BlockdevOptions.
type BlockdevOptionsArchipelago struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Volume       string                       `json:"volume"`
	Mport        *int64                       `json:"mport,omitempty"`
	Vport        *int64                       `json:"vport,omitempty"`
	Segment      *string                      `json:"segment,omitempty"`
}

func (BlockdevOptionsArchipelago) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsArchipelago) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsArchipelago
	}{
		BlockdevDriverArchipelago,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsBlkdebug is an implementation of BlockdevOptions.
type BlockdevOptionsBlkdebug struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Image        BlockdevRef                  `json:"image"`
	Config       *string                      `json:"config,omitempty"`
	Align        *int64                       `json:"align,omitempty"`
	InjectError  []BlkdebugInjectErrorOptions `json:"inject-error,omitempty"`
	SetState     []BlkdebugSetStateOptions    `json:"set-state,omitempty"`
}

func (BlockdevOptionsBlkdebug) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsBlkdebug) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsBlkdebug
	}{
		BlockdevDriverBlkdebug,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsBlkverify is an implementation of BlockdevOptions.
type BlockdevOptionsBlkverify struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Test         BlockdevRef                  `json:"test"`
	Raw          BlockdevRef                  `json:"raw"`
}

func (BlockdevOptionsBlkverify) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsBlkverify) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsBlkverify
	}{
		BlockdevDriverBlkverify,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsBochs is an implementation of BlockdevOptions.
type BlockdevOptionsBochs struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsBochs) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsBochs) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsBochs
	}{
		BlockdevDriverBochs,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsCloop is an implementation of BlockdevOptions.
type BlockdevOptionsCloop struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsCloop) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsCloop) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsCloop
	}{
		BlockdevDriverCloop,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsDmg is an implementation of BlockdevOptions.
type BlockdevOptionsDmg struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsDmg) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsDmg) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsDmg
	}{
		BlockdevDriverDmg,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsFile is an implementation of BlockdevOptions.
type BlockdevOptionsFile struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
	Aio          *BlockdevAioOptions          `json:"aio,omitempty"`
}

func (BlockdevOptionsFile) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsFile) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsFile
	}{
		BlockdevDriverFile,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsFTP is an implementation of BlockdevOptions.
type BlockdevOptionsFTP struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
}

func (BlockdevOptionsFTP) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsFTP) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsFTP
	}{
		BlockdevDriverFTP,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsFTPS is an implementation of BlockdevOptions.
type BlockdevOptionsFTPS struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
}

func (BlockdevOptionsFTPS) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsFTPS) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsFTPS
	}{
		BlockdevDriverFTPS,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsGluster is an implementation of BlockdevOptions.
type BlockdevOptionsGluster struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Volume       string                       `json:"volume"`
	Path         string                       `json:"path"`
	Server       []GlusterServer              `json:"server"`
	DebugLevel   *int64                       `json:"debug-level,omitempty"`
	Logfile      *string                      `json:"logfile,omitempty"`
}

func (BlockdevOptionsGluster) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsGluster) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsGluster
	}{
		BlockdevDriverGluster,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsHostCdrom is an implementation of BlockdevOptions.
type BlockdevOptionsHostCdrom struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
	Aio          *BlockdevAioOptions          `json:"aio,omitempty"`
}

func (BlockdevOptionsHostCdrom) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsHostCdrom) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsHostCdrom
	}{
		BlockdevDriverHostCdrom,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsHostDevice is an implementation of BlockdevOptions.
type BlockdevOptionsHostDevice struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
	Aio          *BlockdevAioOptions          `json:"aio,omitempty"`
}

func (BlockdevOptionsHostDevice) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsHostDevice) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsHostDevice
	}{
		BlockdevDriverHostDevice,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsHTTP is an implementation of BlockdevOptions.
type BlockdevOptionsHTTP struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
}

func (BlockdevOptionsHTTP) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsHTTP) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsHTTP
	}{
		BlockdevDriverHTTP,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsHTTPS is an implementation of BlockdevOptions.
type BlockdevOptionsHTTPS struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
}

func (BlockdevOptionsHTTPS) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsHTTPS) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsHTTPS
	}{
		BlockdevDriverHTTPS,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsLUKS is an implementation of BlockdevOptions.
type BlockdevOptionsLUKS struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	KeySecret    *string                      `json:"key-secret,omitempty"`
}

func (BlockdevOptionsLUKS) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsLUKS
	}{
		BlockdevDriverLUKS,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsNullAio is an implementation of BlockdevOptions.
type BlockdevOptionsNullAio struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Size         *int64                       `json:"size,omitempty"`
	LatencyNs    *uint64                      `json:"latency-ns,omitempty"`
}

func (BlockdevOptionsNullAio) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsNullAio) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsNullAio
	}{
		BlockdevDriverNullAio,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsNullCo is an implementation of BlockdevOptions.
type BlockdevOptionsNullCo struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Size         *int64                       `json:"size,omitempty"`
	LatencyNs    *uint64                      `json:"latency-ns,omitempty"`
}

func (BlockdevOptionsNullCo) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsNullCo) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsNullCo
	}{
		BlockdevDriverNullCo,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsParallels is an implementation of BlockdevOptions.
type BlockdevOptionsParallels struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsParallels) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsParallels) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsParallels
	}{
		BlockdevDriverParallels,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsQcow is an implementation of BlockdevOptions.
type BlockdevOptionsQcow struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Backing      BlockdevRef                  `json:"backing,omitempty"`
}

func (BlockdevOptionsQcow) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsQcow) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsQcow
	}{
		BlockdevDriverQcow,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsQcow2 is an implementation of BlockdevOptions.
type BlockdevOptionsQcow2 struct {
	NodeName            *string                      `json:"node-name,omitempty"`
	Discard             *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache               *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly            *bool                        `json:"read-only,omitempty"`
	DetectZeroes        *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	LazyRefcounts       *bool                        `json:"lazy-refcounts,omitempty"`
	PassDiscardRequest  *bool                        `json:"pass-discard-request,omitempty"`
	PassDiscardSnapshot *bool                        `json:"pass-discard-snapshot,omitempty"`
	PassDiscardOther    *bool                        `json:"pass-discard-other,omitempty"`
	OverlapCheck        Qcow2OverlapChecks           `json:"overlap-check,omitempty"`
	CacheSize           *int64                       `json:"cache-size,omitempty"`
	L2CacheSize         *int64                       `json:"l2-cache-size,omitempty"`
	RefcountCacheSize   *int64                       `json:"refcount-cache-size,omitempty"`
	CacheCleanInterval  *int64                       `json:"cache-clean-interval,omitempty"`
}

func (BlockdevOptionsQcow2) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsQcow2) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsQcow2
	}{
		BlockdevDriverQcow2,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsQed is an implementation of BlockdevOptions.
type BlockdevOptionsQed struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Backing      BlockdevRef                  `json:"backing,omitempty"`
}

func (BlockdevOptionsQed) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsQed) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsQed
	}{
		BlockdevDriverQed,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsQuorum is an implementation of BlockdevOptions.
type BlockdevOptionsQuorum struct {
	NodeName         *string                      `json:"node-name,omitempty"`
	Discard          *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache            *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly         *bool                        `json:"read-only,omitempty"`
	DetectZeroes     *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Blkverify        *bool                        `json:"blkverify,omitempty"`
	Children         []BlockdevRef                `json:"children"`
	VoteThreshold    int64                        `json:"vote-threshold"`
	RewriteCorrupted *bool                        `json:"rewrite-corrupted,omitempty"`
	ReadPattern      *QuorumReadPattern           `json:"read-pattern,omitempty"`
}

func (BlockdevOptionsQuorum) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsQuorum) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsQuorum
	}{
		BlockdevDriverQuorum,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsRaw is an implementation of BlockdevOptions.
type BlockdevOptionsRaw struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsRaw) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsRaw) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsRaw
	}{
		BlockdevDriverRaw,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsReplication is an implementation of BlockdevOptions.
type BlockdevOptionsReplication struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Mode         ReplicationMode              `json:"mode"`
	TopID        *string                      `json:"top-id,omitempty"`
}

func (BlockdevOptionsReplication) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsReplication) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsReplication
	}{
		BlockdevDriverReplication,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsTftp is an implementation of BlockdevOptions.
type BlockdevOptionsTftp struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
}

func (BlockdevOptionsTftp) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsTftp) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsTftp
	}{
		BlockdevDriverTftp,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVdi is an implementation of BlockdevOptions.
type BlockdevOptionsVdi struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVdi) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVdi) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVdi
	}{
		BlockdevDriverVdi,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVhdx is an implementation of BlockdevOptions.
type BlockdevOptionsVhdx struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVhdx) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVhdx) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVhdx
	}{
		BlockdevDriverVhdx,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVMDK is an implementation of BlockdevOptions.
type BlockdevOptionsVMDK struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Backing      BlockdevRef                  `json:"backing,omitempty"`
}

func (BlockdevOptionsVMDK) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVMDK) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVMDK
	}{
		BlockdevDriverVMDK,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVpc is an implementation of BlockdevOptions.
type BlockdevOptionsVpc struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVpc) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVpc) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVpc
	}{
		BlockdevDriverVpc,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVvfat is an implementation of BlockdevOptions.
type BlockdevOptionsVvfat struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Dir          string                       `json:"dir"`
	FatType      *int64                       `json:"fat-type,omitempty"`
	Floppy       *bool                        `json:"floppy,omitempty"`
	Label        *string                      `json:"label,omitempty"`
	Rw           *bool                        `json:"rw,omitempty"`
}

func (BlockdevOptionsVvfat) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVvfat) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVvfat
	}{
		BlockdevDriverVvfat,
		s,
	}
	return json.Marshal(v)
}

func decodeBlockdevOptions(bs json.RawMessage) (BlockdevOptions, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Driver {
	case BlockdevDriverArchipelago:
		var ret BlockdevOptionsArchipelago
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverBlkdebug:
		var ret BlockdevOptionsBlkdebug
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverBlkverify:
		var ret BlockdevOptionsBlkverify
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverBochs:
		var ret BlockdevOptionsBochs
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverCloop:
		var ret BlockdevOptionsCloop
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverDmg:
		var ret BlockdevOptionsDmg
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverFile:
		var ret BlockdevOptionsFile
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverFTP:
		var ret BlockdevOptionsFTP
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverFTPS:
		var ret BlockdevOptionsFTPS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverGluster:
		var ret BlockdevOptionsGluster
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverHostCdrom:
		var ret BlockdevOptionsHostCdrom
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverHostDevice:
		var ret BlockdevOptionsHostDevice
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverHTTP:
		var ret BlockdevOptionsHTTP
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverHTTPS:
		var ret BlockdevOptionsHTTPS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverLUKS:
		var ret BlockdevOptionsLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNullAio:
		var ret BlockdevOptionsNullAio
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNullCo:
		var ret BlockdevOptionsNullCo
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverParallels:
		var ret BlockdevOptionsParallels
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQcow:
		var ret BlockdevOptionsQcow
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQcow2:
		var ret BlockdevOptionsQcow2
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQed:
		var ret BlockdevOptionsQed
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQuorum:
		var ret BlockdevOptionsQuorum
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverRaw:
		var ret BlockdevOptionsRaw
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverReplication:
		var ret BlockdevOptionsReplication
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverTftp:
		var ret BlockdevOptionsTftp
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVdi:
		var ret BlockdevOptionsVdi
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVhdx:
		var ret BlockdevOptionsVhdx
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVMDK:
		var ret BlockdevOptionsVMDK
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVpc:
		var ret BlockdevOptionsVpc
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVvfat:
		var ret BlockdevOptionsVvfat
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union BlockdevOptions", v.Driver)
	}
}

// BlockdevRef -> BlockdevRef (alternate)

// BlockdevRef implements the "BlockdevRef" QMP API type.
//
// Can be one of:
//   - BlockdevRefDefinition
//   - BlockdevRefReference
type BlockdevRef interface {
	isBlockdevRef()
}

// BlockdevRefDefinition is an implementation of BlockdevRef
type BlockdevRefDefinition BlockdevOptions

func (BlockdevOptionsArchipelago) isBlockdevRef() {}
func (BlockdevOptionsBlkdebug) isBlockdevRef()    {}
func (BlockdevOptionsBlkverify) isBlockdevRef()   {}
func (BlockdevOptionsBochs) isBlockdevRef()       {}
func (BlockdevOptionsCloop) isBlockdevRef()       {}
func (BlockdevOptionsDmg) isBlockdevRef()         {}
func (BlockdevOptionsFile) isBlockdevRef()        {}
func (BlockdevOptionsFTP) isBlockdevRef()         {}
func (BlockdevOptionsFTPS) isBlockdevRef()        {}
func (BlockdevOptionsGluster) isBlockdevRef()     {}
func (BlockdevOptionsHostCdrom) isBlockdevRef()   {}
func (BlockdevOptionsHostDevice) isBlockdevRef()  {}
func (BlockdevOptionsHTTP) isBlockdevRef()        {}
func (BlockdevOptionsHTTPS) isBlockdevRef()       {}
func (BlockdevOptionsLUKS) isBlockdevRef()        {}
func (BlockdevOptionsNullAio) isBlockdevRef()     {}
func (BlockdevOptionsNullCo) isBlockdevRef()      {}
func (BlockdevOptionsParallels) isBlockdevRef()   {}
func (BlockdevOptionsQcow) isBlockdevRef()        {}
func (BlockdevOptionsQcow2) isBlockdevRef()       {}
func (BlockdevOptionsQed) isBlockdevRef()         {}
func (BlockdevOptionsQuorum) isBlockdevRef()      {}
func (BlockdevOptionsRaw) isBlockdevRef()         {}
func (BlockdevOptionsReplication) isBlockdevRef() {}
func (BlockdevOptionsTftp) isBlockdevRef()        {}
func (BlockdevOptionsVdi) isBlockdevRef()         {}
func (BlockdevOptionsVhdx) isBlockdevRef()        {}
func (BlockdevOptionsVMDK) isBlockdevRef()        {}
func (BlockdevOptionsVpc) isBlockdevRef()         {}
func (BlockdevOptionsVvfat) isBlockdevRef()       {}

// BlockdevRefReference is an implementation of BlockdevRef
type BlockdevRefReference string

func (BlockdevRefReference) isBlockdevRef() {}

func decodeBlockdevRef(bs json.RawMessage) (BlockdevRef, error) {
	var reference BlockdevRefReference
	if err := json.Unmarshal([]byte(bs), &reference); err == nil {
		return reference, nil
	}

	if definition, err := decodeBlockdevOptions([]byte(bs)); err == nil {
		switch impl := definition.(type) {
		case BlockdevOptionsArchipelago:
			return impl, nil
		case BlockdevOptionsBlkdebug:
			return impl, nil
		case BlockdevOptionsBlkverify:
			return impl, nil
		case BlockdevOptionsBochs:
			return impl, nil
		case BlockdevOptionsCloop:
			return impl, nil
		case BlockdevOptionsDmg:
			return impl, nil
		case BlockdevOptionsFile:
			return impl, nil
		case BlockdevOptionsFTP:
			return impl, nil
		case BlockdevOptionsFTPS:
			return impl, nil
		case BlockdevOptionsGluster:
			return impl, nil
		case BlockdevOptionsHostCdrom:
			return impl, nil
		case BlockdevOptionsHostDevice:
			return impl, nil
		case BlockdevOptionsHTTP:
			return impl, nil
		case BlockdevOptionsHTTPS:
			return impl, nil
		case BlockdevOptionsLUKS:
			return impl, nil
		case BlockdevOptionsNullAio:
			return impl, nil
		case BlockdevOptionsNullCo:
			return impl, nil
		case BlockdevOptionsParallels:
			return impl, nil
		case BlockdevOptionsQcow:
			return impl, nil
		case BlockdevOptionsQcow2:
			return impl, nil
		case BlockdevOptionsQed:
			return impl, nil
		case BlockdevOptionsQuorum:
			return impl, nil
		case BlockdevOptionsRaw:
			return impl, nil
		case BlockdevOptionsReplication:
			return impl, nil
		case BlockdevOptionsTftp:
			return impl, nil
		case BlockdevOptionsVdi:
			return impl, nil
		case BlockdevOptionsVhdx:
			return impl, nil
		case BlockdevOptionsVMDK:
			return impl, nil
		case BlockdevOptionsVpc:
			return impl, nil
		case BlockdevOptionsVvfat:
			return impl, nil
		}
	}
	return nil, fmt.Errorf("unable to decode %q as a BlockdevRef", string(bs))
}

// BlockdevSnapshot -> BlockdevSnapshot (struct)

// BlockdevSnapshot implements the "BlockdevSnapshot" QMP API type.
type BlockdevSnapshot struct {
	Node    string `json:"node"`
	Overlay string `json:"overlay"`
}

// BlockdevSnapshotInternal -> BlockdevSnapshotInternal (struct)

// BlockdevSnapshotInternal implements the "BlockdevSnapshotInternal" QMP API type.
type BlockdevSnapshotInternal struct {
	Device string `json:"device"`
	Name   string `json:"name"`
}

// BlockdevSnapshotSync -> BlockdevSnapshotSync (struct)

// BlockdevSnapshotSync implements the "BlockdevSnapshotSync" QMP API type.
type BlockdevSnapshotSync struct {
	Device           *string       `json:"device,omitempty"`
	NodeName         *string       `json:"node-name,omitempty"`
	SnapshotFile     string        `json:"snapshot-file"`
	SnapshotNodeName *string       `json:"snapshot-node-name,omitempty"`
	Format           *string       `json:"format,omitempty"`
	Mode             *NewImageMode `json:"mode,omitempty"`
}

// ChardevBackend -> ChardevBackend (simple union)

// ChardevBackend implements the "ChardevBackend" QMP API type.
//
// Can be one of:
//   - ChardevBackendBraille
//   - ChardevBackendConsole
//   - ChardevBackendFile
//   - ChardevBackendMemory
//   - ChardevBackendMsmouse
//   - ChardevBackendMux
//   - ChardevBackendNull
//   - ChardevBackendParallel
//   - ChardevBackendPipe
//   - ChardevBackendPty
//   - ChardevBackendRingbuf
//   - ChardevBackendSerial
//   - ChardevBackendSocket
//   - ChardevBackendSpiceport
//   - ChardevBackendSpicevmc
//   - ChardevBackendStdio
//   - ChardevBackendTestdev
//   - ChardevBackendUDP
//   - ChardevBackendVc
type ChardevBackend interface {
	isChardevBackend()
}

// ChardevBackendBraille is an implementation of ChardevBackend.
type ChardevBackendBraille ChardevCommon

func (ChardevBackendBraille) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendBraille) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "braille",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendConsole is an implementation of ChardevBackend.
type ChardevBackendConsole ChardevCommon

func (ChardevBackendConsole) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendConsole) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "console",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendFile is an implementation of ChardevBackend.
type ChardevBackendFile ChardevFile

func (ChardevBackendFile) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendFile) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "file",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendMemory is an implementation of ChardevBackend.
type ChardevBackendMemory ChardevRingbuf

func (ChardevBackendMemory) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendMemory) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "memory",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendMsmouse is an implementation of ChardevBackend.
type ChardevBackendMsmouse ChardevCommon

func (ChardevBackendMsmouse) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendMsmouse) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "msmouse",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendMux is an implementation of ChardevBackend.
type ChardevBackendMux ChardevMux

func (ChardevBackendMux) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendMux) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "mux",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendNull is an implementation of ChardevBackend.
type ChardevBackendNull ChardevCommon

func (ChardevBackendNull) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendNull) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "null",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendParallel is an implementation of ChardevBackend.
type ChardevBackendParallel ChardevHostdev

func (ChardevBackendParallel) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendParallel) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "parallel",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendPipe is an implementation of ChardevBackend.
type ChardevBackendPipe ChardevHostdev

func (ChardevBackendPipe) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendPipe) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "pipe",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendPty is an implementation of ChardevBackend.
type ChardevBackendPty ChardevCommon

func (ChardevBackendPty) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendPty) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "pty",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendRingbuf is an implementation of ChardevBackend.
type ChardevBackendRingbuf ChardevRingbuf

func (ChardevBackendRingbuf) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendRingbuf) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "ringbuf",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendSerial is an implementation of ChardevBackend.
type ChardevBackendSerial ChardevHostdev

func (ChardevBackendSerial) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendSerial) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "serial",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendSocket is an implementation of ChardevBackend.
type ChardevBackendSocket ChardevSocket

func (ChardevBackendSocket) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendSocket) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "socket",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendSpiceport is an implementation of ChardevBackend.
type ChardevBackendSpiceport ChardevSpicePort

func (ChardevBackendSpiceport) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendSpiceport) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "spiceport",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendSpicevmc is an implementation of ChardevBackend.
type ChardevBackendSpicevmc ChardevSpiceChannel

func (ChardevBackendSpicevmc) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendSpicevmc) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "spicevmc",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendStdio is an implementation of ChardevBackend.
type ChardevBackendStdio ChardevStdio

func (ChardevBackendStdio) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendStdio) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "stdio",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendTestdev is an implementation of ChardevBackend.
type ChardevBackendTestdev ChardevCommon

func (ChardevBackendTestdev) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendTestdev) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "testdev",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendUDP is an implementation of ChardevBackend.
type ChardevBackendUDP ChardevUDP

func (ChardevBackendUDP) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendUDP) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "udp",
		"data": s,
	}
	return json.Marshal(v)
}

// ChardevBackendVc is an implementation of ChardevBackend.
type ChardevBackendVc ChardevVc

func (ChardevBackendVc) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVc) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "vc",
		"data": s,
	}
	return json.Marshal(v)
}

func decodeChardevBackend(bs json.RawMessage) (ChardevBackend, error) {
	v := struct {
		T string          `json:"type"`
		V json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.T {
	case "braille":
		var ret ChardevBackendBraille
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "console":
		var ret ChardevBackendConsole
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "file":
		var ret ChardevBackendFile
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "memory":
		var ret ChardevBackendMemory
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "msmouse":
		var ret ChardevBackendMsmouse
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "mux":
		var ret ChardevBackendMux
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "null":
		var ret ChardevBackendNull
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "parallel":
		var ret ChardevBackendParallel
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "pipe":
		var ret ChardevBackendPipe
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "pty":
		var ret ChardevBackendPty
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "ringbuf":
		var ret ChardevBackendRingbuf
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "serial":
		var ret ChardevBackendSerial
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "socket":
		var ret ChardevBackendSocket
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "spiceport":
		var ret ChardevBackendSpiceport
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "spicevmc":
		var ret ChardevBackendSpicevmc
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "stdio":
		var ret ChardevBackendStdio
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "testdev":
		var ret ChardevBackendTestdev
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "udp":
		var ret ChardevBackendUDP
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "vc":
		var ret ChardevBackendVc
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("unknown subtype %q for union ChardevBackend", v.T)
	}
}

// ChardevBackendInfo -> ChardevBackendInfo (struct)

// ChardevBackendInfo implements the "ChardevBackendInfo" QMP API type.
type ChardevBackendInfo struct {
	Name string `json:"name"`
}

// ChardevCommon -> ChardevCommon (struct)

// ChardevCommon implements the "ChardevCommon" QMP API type.
type ChardevCommon struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
}

// ChardevFile -> ChardevFile (struct)

// ChardevFile implements the "ChardevFile" QMP API type.
type ChardevFile struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
	In        *string `json:"in,omitempty"`
	Out       string  `json:"out"`
	Append    *bool   `json:"append,omitempty"`
}

// ChardevHostdev -> ChardevHostdev (struct)

// ChardevHostdev implements the "ChardevHostdev" QMP API type.
type ChardevHostdev struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
	Device    string  `json:"device"`
}

// ChardevInfo -> ChardevInfo (struct)

// ChardevInfo implements the "ChardevInfo" QMP API type.
type ChardevInfo struct {
	Label        string `json:"label"`
	Filename     string `json:"filename"`
	FrontendOpen bool   `json:"frontend-open"`
}

// ChardevMux -> ChardevMux (struct)

// ChardevMux implements the "ChardevMux" QMP API type.
type ChardevMux struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
	Chardev   string  `json:"chardev"`
}

// ChardevReturn -> ChardevReturn (struct)

// ChardevReturn implements the "ChardevReturn" QMP API type.
type ChardevReturn struct {
	Pty *string `json:"pty,omitempty"`
}

// ChardevRingbuf -> ChardevRingbuf (struct)

// ChardevRingbuf implements the "ChardevRingbuf" QMP API type.
type ChardevRingbuf struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
	Size      *int64  `json:"size,omitempty"`
}

// ChardevSocket -> ChardevSocket (struct)

// ChardevSocket implements the "ChardevSocket" QMP API type.
type ChardevSocket struct {
	Logfile   *string       `json:"logfile,omitempty"`
	Logappend *bool         `json:"logappend,omitempty"`
	Addr      SocketAddress `json:"addr"`
	TLSCreds  *string       `json:"tls-creds,omitempty"`
	Server    *bool         `json:"server,omitempty"`
	Wait      *bool         `json:"wait,omitempty"`
	Nodelay   *bool         `json:"nodelay,omitempty"`
	Telnet    *bool         `json:"telnet,omitempty"`
	Reconnect *int64        `json:"reconnect,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *ChardevSocket) UnmarshalJSON(bs []byte) error {
	v := struct {
		Logfile   *string         `json:"logfile,omitempty"`
		Logappend *bool           `json:"logappend,omitempty"`
		Addr      json.RawMessage `json:"addr"`
		TLSCreds  *string         `json:"tls-creds,omitempty"`
		Server    *bool           `json:"server,omitempty"`
		Wait      *bool           `json:"wait,omitempty"`
		Nodelay   *bool           `json:"nodelay,omitempty"`
		Telnet    *bool           `json:"telnet,omitempty"`
		Reconnect *int64          `json:"reconnect,omitempty"`
	}{}
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return err
	}
	s.Logfile = v.Logfile
	s.Logappend = v.Logappend
	s.Addr, err = decodeSocketAddress(v.Addr)
	if err != nil {
		return err
	}
	s.TLSCreds = v.TLSCreds
	s.Server = v.Server
	s.Wait = v.Wait
	s.Nodelay = v.Nodelay
	s.Telnet = v.Telnet
	s.Reconnect = v.Reconnect

	return nil
}

// ChardevSpiceChannel -> ChardevSpiceChannel (struct)

// ChardevSpiceChannel implements the "ChardevSpiceChannel" QMP API type.
type ChardevSpiceChannel struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
	Type      string  `json:"type"`
}

// ChardevSpicePort -> ChardevSpicePort (struct)

// ChardevSpicePort implements the "ChardevSpicePort" QMP API type.
type ChardevSpicePort struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
	Fqdn      string  `json:"fqdn"`
}

// ChardevStdio -> ChardevStdio (struct)

// ChardevStdio implements the "ChardevStdio" QMP API type.
type ChardevStdio struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
	Signal    *bool   `json:"signal,omitempty"`
}

// ChardevUdp -> ChardevUDP (struct)

// ChardevUDP implements the "ChardevUdp" QMP API type.
type ChardevUDP struct {
	Logfile   *string       `json:"logfile,omitempty"`
	Logappend *bool         `json:"logappend,omitempty"`
	Remote    SocketAddress `json:"remote"`
	Local     SocketAddress `json:"local,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *ChardevUDP) UnmarshalJSON(bs []byte) error {
	v := struct {
		Logfile   *string         `json:"logfile,omitempty"`
		Logappend *bool           `json:"logappend,omitempty"`
		Remote    json.RawMessage `json:"remote"`
		Local     json.RawMessage `json:"local"`
	}{}
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return err
	}
	s.Logfile = v.Logfile
	s.Logappend = v.Logappend
	s.Remote, err = decodeSocketAddress(v.Remote)
	if err != nil {
		return err
	}
	if len(v.Local) > 0 {
		s.Local, err = decodeSocketAddress(v.Local)
		if err != nil {
			return err
		}
	} else {
		s.Local = nil
	}

	return nil
}

// ChardevVC -> ChardevVc (struct)

// ChardevVc implements the "ChardevVC" QMP API type.
type ChardevVc struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
	Width     *int64  `json:"width,omitempty"`
	Height    *int64  `json:"height,omitempty"`
	Cols      *int64  `json:"cols,omitempty"`
	Rows      *int64  `json:"rows,omitempty"`
}

// CommandInfo -> CommandInfo (struct)

// CommandInfo implements the "CommandInfo" QMP API type.
type CommandInfo struct {
	Name string `json:"name"`
}

// CommandLineOptionInfo -> CommandLineOptionInfo (struct)

// CommandLineOptionInfo implements the "CommandLineOptionInfo" QMP API type.
type CommandLineOptionInfo struct {
	Option     string                     `json:"option"`
	Parameters []CommandLineParameterInfo `json:"parameters"`
}

// CommandLineParameterInfo -> CommandLineParameterInfo (struct)

// CommandLineParameterInfo implements the "CommandLineParameterInfo" QMP API type.
type CommandLineParameterInfo struct {
	Name    string                   `json:"name"`
	Type    CommandLineParameterType `json:"type"`
	Help    *string                  `json:"help,omitempty"`
	Default *string                  `json:"default,omitempty"`
}

// CommandLineParameterType -> CommandLineParameterType (enum)

// CommandLineParameterType implements the "CommandLineParameterType" QMP API type.
type CommandLineParameterType int

// Known values of CommandLineParameterType.
const (
	CommandLineParameterTypeString CommandLineParameterType = iota
	CommandLineParameterTypeBoolean
	CommandLineParameterTypefloat64
	CommandLineParameterTypeuint64
)

// String implements fmt.Stringer.
func (e CommandLineParameterType) String() string {
	switch e {
	case CommandLineParameterTypeString:
		return "string"
	case CommandLineParameterTypeBoolean:
		return "boolean"
	case CommandLineParameterTypefloat64:
		return "number"
	case CommandLineParameterTypeuint64:
		return "size"
	default:
		return fmt.Sprintf("CommandLineParameterType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e CommandLineParameterType) MarshalJSON() ([]byte, error) {
	switch e {
	case CommandLineParameterTypeString:
		return json.Marshal("string")
	case CommandLineParameterTypeBoolean:
		return json.Marshal("boolean")
	case CommandLineParameterTypefloat64:
		return json.Marshal("number")
	case CommandLineParameterTypeuint64:
		return json.Marshal("size")
	default:
		return nil, fmt.Errorf("unknown enum value %q for CommandLineParameterType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *CommandLineParameterType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "string":
		*e = CommandLineParameterTypeString
	case "boolean":
		*e = CommandLineParameterTypeBoolean
	case "number":
		*e = CommandLineParameterTypefloat64
	case "size":
		*e = CommandLineParameterTypeuint64
	default:
		return fmt.Errorf("unknown enum value %q for CommandLineParameterType", s)
	}
	return nil
}

// CpuDefinitionInfo -> CPUDefinitionInfo (struct)

// CPUDefinitionInfo implements the "CpuDefinitionInfo" QMP API type.
type CPUDefinitionInfo struct {
	Name          string `json:"name"`
	MigrationSafe *bool  `json:"migration-safe,omitempty"`
	Static        bool   `json:"static"`
}

// CpuInfo -> CPUInfo (flat union)

// CPUInfo implements the "CpuInfo" QMP API type.
//
// Can be one of:
//   - CPUInfoMIPS
//   - CPUInfoOther
//   - CPUInfoPPC
//   - CPUInfoSPARC
//   - CPUInfoTricore
//   - CPUInfoX86
type CPUInfo interface {
	isCPUInfo()
}

// CPUInfoMIPS is an implementation of CPUInfo.
type CPUInfoMIPS struct {
	CPU      int64  `json:"CPU"`
	Current  bool   `json:"current"`
	Halted   bool   `json:"halted"`
	QomPath  string `json:"qom_path"`
	ThreadID int64  `json:"thread_id"`
	Pc       int64  `json:"PC"`
}

func (CPUInfoMIPS) isCPUInfo() {}

// MarshalJSON implements json.Marshaler.
func (s CPUInfoMIPS) MarshalJSON() ([]byte, error) {
	v := struct {
		Arch CPUInfoArch `json:"arch"`
		CPUInfoMIPS
	}{
		CPUInfoArchMIPS,
		s,
	}
	return json.Marshal(v)
}

// CPUInfoOther is an implementation of CPUInfo.
type CPUInfoOther struct {
	CPU      int64  `json:"CPU"`
	Current  bool   `json:"current"`
	Halted   bool   `json:"halted"`
	QomPath  string `json:"qom_path"`
	ThreadID int64  `json:"thread_id"`
}

func (CPUInfoOther) isCPUInfo() {}

// MarshalJSON implements json.Marshaler.
func (s CPUInfoOther) MarshalJSON() ([]byte, error) {
	v := struct {
		Arch CPUInfoArch `json:"arch"`
		CPUInfoOther
	}{
		CPUInfoArchOther,
		s,
	}
	return json.Marshal(v)
}

// CPUInfoPPC is an implementation of CPUInfo.
type CPUInfoPPC struct {
	CPU      int64  `json:"CPU"`
	Current  bool   `json:"current"`
	Halted   bool   `json:"halted"`
	QomPath  string `json:"qom_path"`
	ThreadID int64  `json:"thread_id"`
	Nip      int64  `json:"nip"`
}

func (CPUInfoPPC) isCPUInfo() {}

// MarshalJSON implements json.Marshaler.
func (s CPUInfoPPC) MarshalJSON() ([]byte, error) {
	v := struct {
		Arch CPUInfoArch `json:"arch"`
		CPUInfoPPC
	}{
		CPUInfoArchPPC,
		s,
	}
	return json.Marshal(v)
}

// CPUInfoSPARC is an implementation of CPUInfo.
type CPUInfoSPARC struct {
	CPU      int64  `json:"CPU"`
	Current  bool   `json:"current"`
	Halted   bool   `json:"halted"`
	QomPath  string `json:"qom_path"`
	ThreadID int64  `json:"thread_id"`
	Pc       int64  `json:"pc"`
	Npc      int64  `json:"npc"`
}

func (CPUInfoSPARC) isCPUInfo() {}

// MarshalJSON implements json.Marshaler.
func (s CPUInfoSPARC) MarshalJSON() ([]byte, error) {
	v := struct {
		Arch CPUInfoArch `json:"arch"`
		CPUInfoSPARC
	}{
		CPUInfoArchSPARC,
		s,
	}
	return json.Marshal(v)
}

// CPUInfoTricore is an implementation of CPUInfo.
type CPUInfoTricore struct {
	CPU      int64  `json:"CPU"`
	Current  bool   `json:"current"`
	Halted   bool   `json:"halted"`
	QomPath  string `json:"qom_path"`
	ThreadID int64  `json:"thread_id"`
	Pc       int64  `json:"PC"`
}

func (CPUInfoTricore) isCPUInfo() {}

// MarshalJSON implements json.Marshaler.
func (s CPUInfoTricore) MarshalJSON() ([]byte, error) {
	v := struct {
		Arch CPUInfoArch `json:"arch"`
		CPUInfoTricore
	}{
		CPUInfoArchTricore,
		s,
	}
	return json.Marshal(v)
}

// CPUInfoX86 is an implementation of CPUInfo.
type CPUInfoX86 struct {
	CPU      int64  `json:"CPU"`
	Current  bool   `json:"current"`
	Halted   bool   `json:"halted"`
	QomPath  string `json:"qom_path"`
	ThreadID int64  `json:"thread_id"`
	Pc       int64  `json:"pc"`
}

func (CPUInfoX86) isCPUInfo() {}

// MarshalJSON implements json.Marshaler.
func (s CPUInfoX86) MarshalJSON() ([]byte, error) {
	v := struct {
		Arch CPUInfoArch `json:"arch"`
		CPUInfoX86
	}{
		CPUInfoArchX86,
		s,
	}
	return json.Marshal(v)
}

func decodeCPUInfo(bs json.RawMessage) (CPUInfo, error) {
	v := struct {
		Arch CPUInfoArch `json:"arch"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Arch {
	case CPUInfoArchMIPS:
		var ret CPUInfoMIPS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case CPUInfoArchOther:
		var ret CPUInfoOther
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case CPUInfoArchPPC:
		var ret CPUInfoPPC
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case CPUInfoArchSPARC:
		var ret CPUInfoSPARC
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case CPUInfoArchTricore:
		var ret CPUInfoTricore
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case CPUInfoArchX86:
		var ret CPUInfoX86
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union CPUInfo", v.Arch)
	}
}

// CpuInfoArch -> CPUInfoArch (enum)

// CPUInfoArch implements the "CpuInfoArch" QMP API type.
type CPUInfoArch int

// Known values of CPUInfoArch.
const (
	CPUInfoArchX86 CPUInfoArch = iota
	CPUInfoArchSPARC
	CPUInfoArchPPC
	CPUInfoArchMIPS
	CPUInfoArchTricore
	CPUInfoArchOther
)

// String implements fmt.Stringer.
func (e CPUInfoArch) String() string {
	switch e {
	case CPUInfoArchX86:
		return "x86"
	case CPUInfoArchSPARC:
		return "sparc"
	case CPUInfoArchPPC:
		return "ppc"
	case CPUInfoArchMIPS:
		return "mips"
	case CPUInfoArchTricore:
		return "tricore"
	case CPUInfoArchOther:
		return "other"
	default:
		return fmt.Sprintf("CPUInfoArch(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e CPUInfoArch) MarshalJSON() ([]byte, error) {
	switch e {
	case CPUInfoArchX86:
		return json.Marshal("x86")
	case CPUInfoArchSPARC:
		return json.Marshal("sparc")
	case CPUInfoArchPPC:
		return json.Marshal("ppc")
	case CPUInfoArchMIPS:
		return json.Marshal("mips")
	case CPUInfoArchTricore:
		return json.Marshal("tricore")
	case CPUInfoArchOther:
		return json.Marshal("other")
	default:
		return nil, fmt.Errorf("unknown enum value %q for CPUInfoArch", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *CPUInfoArch) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "x86":
		*e = CPUInfoArchX86
	case "sparc":
		*e = CPUInfoArchSPARC
	case "ppc":
		*e = CPUInfoArchPPC
	case "mips":
		*e = CPUInfoArchMIPS
	case "tricore":
		*e = CPUInfoArchTricore
	case "other":
		*e = CPUInfoArchOther
	default:
		return fmt.Errorf("unknown enum value %q for CPUInfoArch", s)
	}
	return nil
}

// CpuInstanceProperties -> CPUInstanceProperties (struct)

// CPUInstanceProperties implements the "CpuInstanceProperties" QMP API type.
type CPUInstanceProperties struct {
	NodeID   *int64 `json:"node-id,omitempty"`
	SocketID *int64 `json:"socket-id,omitempty"`
	CoreID   *int64 `json:"core-id,omitempty"`
	ThreadID *int64 `json:"thread-id,omitempty"`
}

// CpuModelBaselineInfo -> CPUModelBaselineInfo (struct)

// CPUModelBaselineInfo implements the "CpuModelBaselineInfo" QMP API type.
type CPUModelBaselineInfo struct {
	Model CPUModelInfo `json:"model"`
}

// CpuModelCompareInfo -> CPUModelCompareInfo (struct)

// CPUModelCompareInfo implements the "CpuModelCompareInfo" QMP API type.
type CPUModelCompareInfo struct {
	Result                CPUModelCompareResult `json:"result"`
	ResponsibleProperties []string              `json:"responsible-properties"`
}

// CpuModelCompareResult -> CPUModelCompareResult (enum)

// CPUModelCompareResult implements the "CpuModelCompareResult" QMP API type.
type CPUModelCompareResult int

// Known values of CPUModelCompareResult.
const (
	CPUModelCompareResultIncompatible CPUModelCompareResult = iota
	CPUModelCompareResultIdentical
	CPUModelCompareResultSuperset
	CPUModelCompareResultSubset
)

// String implements fmt.Stringer.
func (e CPUModelCompareResult) String() string {
	switch e {
	case CPUModelCompareResultIncompatible:
		return "incompatible"
	case CPUModelCompareResultIdentical:
		return "identical"
	case CPUModelCompareResultSuperset:
		return "superset"
	case CPUModelCompareResultSubset:
		return "subset"
	default:
		return fmt.Sprintf("CPUModelCompareResult(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e CPUModelCompareResult) MarshalJSON() ([]byte, error) {
	switch e {
	case CPUModelCompareResultIncompatible:
		return json.Marshal("incompatible")
	case CPUModelCompareResultIdentical:
		return json.Marshal("identical")
	case CPUModelCompareResultSuperset:
		return json.Marshal("superset")
	case CPUModelCompareResultSubset:
		return json.Marshal("subset")
	default:
		return nil, fmt.Errorf("unknown enum value %q for CPUModelCompareResult", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *CPUModelCompareResult) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "incompatible":
		*e = CPUModelCompareResultIncompatible
	case "identical":
		*e = CPUModelCompareResultIdentical
	case "superset":
		*e = CPUModelCompareResultSuperset
	case "subset":
		*e = CPUModelCompareResultSubset
	default:
		return fmt.Errorf("unknown enum value %q for CPUModelCompareResult", s)
	}
	return nil
}

// CpuModelExpansionInfo -> CPUModelExpansionInfo (struct)

// CPUModelExpansionInfo implements the "CpuModelExpansionInfo" QMP API type.
type CPUModelExpansionInfo struct {
	Model CPUModelInfo `json:"model"`
}

// CpuModelExpansionType -> CPUModelExpansionType (enum)

// CPUModelExpansionType implements the "CpuModelExpansionType" QMP API type.
type CPUModelExpansionType int

// Known values of CPUModelExpansionType.
const (
	CPUModelExpansionTypeStatic CPUModelExpansionType = iota
	CPUModelExpansionTypeFull
)

// String implements fmt.Stringer.
func (e CPUModelExpansionType) String() string {
	switch e {
	case CPUModelExpansionTypeStatic:
		return "static"
	case CPUModelExpansionTypeFull:
		return "full"
	default:
		return fmt.Sprintf("CPUModelExpansionType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e CPUModelExpansionType) MarshalJSON() ([]byte, error) {
	switch e {
	case CPUModelExpansionTypeStatic:
		return json.Marshal("static")
	case CPUModelExpansionTypeFull:
		return json.Marshal("full")
	default:
		return nil, fmt.Errorf("unknown enum value %q for CPUModelExpansionType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *CPUModelExpansionType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "static":
		*e = CPUModelExpansionTypeStatic
	case "full":
		*e = CPUModelExpansionTypeFull
	default:
		return fmt.Errorf("unknown enum value %q for CPUModelExpansionType", s)
	}
	return nil
}

// CpuModelInfo -> CPUModelInfo (struct)

// CPUModelInfo implements the "CpuModelInfo" QMP API type.
type CPUModelInfo struct {
	Name  string       `json:"name"`
	Props *interface{} `json:"props,omitempty"`
}

// EVENT DEVICE_DELETED

// EVENT DEVICE_TRAY_MOVED

// EVENT DUMP_COMPLETED

// DataFormat -> DataFormat (enum)

// DataFormat implements the "DataFormat" QMP API type.
type DataFormat int

// Known values of DataFormat.
const (
	DataFormatUtf8 DataFormat = iota
	DataFormatBase64
)

// String implements fmt.Stringer.
func (e DataFormat) String() string {
	switch e {
	case DataFormatUtf8:
		return "utf8"
	case DataFormatBase64:
		return "base64"
	default:
		return fmt.Sprintf("DataFormat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DataFormat) MarshalJSON() ([]byte, error) {
	switch e {
	case DataFormatUtf8:
		return json.Marshal("utf8")
	case DataFormatBase64:
		return json.Marshal("base64")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DataFormat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DataFormat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "utf8":
		*e = DataFormatUtf8
	case "base64":
		*e = DataFormatBase64
	default:
		return fmt.Errorf("unknown enum value %q for DataFormat", s)
	}
	return nil
}

// DevicePropertyInfo -> DevicePropertyInfo (struct)

// DevicePropertyInfo implements the "DevicePropertyInfo" QMP API type.
type DevicePropertyInfo struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description *string `json:"description,omitempty"`
}

// DirtyBitmapStatus -> DirtyBitmapStatus (enum)

// DirtyBitmapStatus implements the "DirtyBitmapStatus" QMP API type.
type DirtyBitmapStatus int

// Known values of DirtyBitmapStatus.
const (
	DirtyBitmapStatusActive DirtyBitmapStatus = iota
	DirtyBitmapStatusDisabled
	DirtyBitmapStatusFrozen
)

// String implements fmt.Stringer.
func (e DirtyBitmapStatus) String() string {
	switch e {
	case DirtyBitmapStatusActive:
		return "active"
	case DirtyBitmapStatusDisabled:
		return "disabled"
	case DirtyBitmapStatusFrozen:
		return "frozen"
	default:
		return fmt.Sprintf("DirtyBitmapStatus(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DirtyBitmapStatus) MarshalJSON() ([]byte, error) {
	switch e {
	case DirtyBitmapStatusActive:
		return json.Marshal("active")
	case DirtyBitmapStatusDisabled:
		return json.Marshal("disabled")
	case DirtyBitmapStatusFrozen:
		return json.Marshal("frozen")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DirtyBitmapStatus", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DirtyBitmapStatus) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "active":
		*e = DirtyBitmapStatusActive
	case "disabled":
		*e = DirtyBitmapStatusDisabled
	case "frozen":
		*e = DirtyBitmapStatusFrozen
	default:
		return fmt.Errorf("unknown enum value %q for DirtyBitmapStatus", s)
	}
	return nil
}

// DriveBackup -> DriveBackup (struct)

// DriveBackup implements the "DriveBackup" QMP API type.
type DriveBackup struct {
	JobID         *string          `json:"job-id,omitempty"`
	Device        string           `json:"device"`
	Target        string           `json:"target"`
	Format        *string          `json:"format,omitempty"`
	Sync          MirrorSyncMode   `json:"sync"`
	Mode          *NewImageMode    `json:"mode,omitempty"`
	Speed         *int64           `json:"speed,omitempty"`
	Bitmap        *string          `json:"bitmap,omitempty"`
	Compress      *bool            `json:"compress,omitempty"`
	OnSourceError *BlockdevOnError `json:"on-source-error,omitempty"`
	OnTargetError *BlockdevOnError `json:"on-target-error,omitempty"`
}

// DriveMirror -> DriveMirror (struct)

// DriveMirror implements the "DriveMirror" QMP API type.
type DriveMirror struct {
	JobID         *string          `json:"job-id,omitempty"`
	Device        string           `json:"device"`
	Target        string           `json:"target"`
	Format        *string          `json:"format,omitempty"`
	NodeName      *string          `json:"node-name,omitempty"`
	Replaces      *string          `json:"replaces,omitempty"`
	Sync          MirrorSyncMode   `json:"sync"`
	Mode          *NewImageMode    `json:"mode,omitempty"`
	Speed         *int64           `json:"speed,omitempty"`
	Granularity   *uint32          `json:"granularity,omitempty"`
	BufSize       *int64           `json:"buf-size,omitempty"`
	OnSourceError *BlockdevOnError `json:"on-source-error,omitempty"`
	OnTargetError *BlockdevOnError `json:"on-target-error,omitempty"`
	Unmap         *bool            `json:"unmap,omitempty"`
}

// DumpGuestMemoryCapability -> DumpGuestMemoryCapability (struct)

// DumpGuestMemoryCapability implements the "DumpGuestMemoryCapability" QMP API type.
type DumpGuestMemoryCapability struct {
	Formats []DumpGuestMemoryFormat `json:"formats"`
}

// DumpGuestMemoryFormat -> DumpGuestMemoryFormat (enum)

// DumpGuestMemoryFormat implements the "DumpGuestMemoryFormat" QMP API type.
type DumpGuestMemoryFormat int

// Known values of DumpGuestMemoryFormat.
const (
	DumpGuestMemoryFormatElf DumpGuestMemoryFormat = iota
	DumpGuestMemoryFormatKdumpZlib
	DumpGuestMemoryFormatKdumpLzo
	DumpGuestMemoryFormatKdumpSnappy
)

// String implements fmt.Stringer.
func (e DumpGuestMemoryFormat) String() string {
	switch e {
	case DumpGuestMemoryFormatElf:
		return "elf"
	case DumpGuestMemoryFormatKdumpZlib:
		return "kdump-zlib"
	case DumpGuestMemoryFormatKdumpLzo:
		return "kdump-lzo"
	case DumpGuestMemoryFormatKdumpSnappy:
		return "kdump-snappy"
	default:
		return fmt.Sprintf("DumpGuestMemoryFormat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DumpGuestMemoryFormat) MarshalJSON() ([]byte, error) {
	switch e {
	case DumpGuestMemoryFormatElf:
		return json.Marshal("elf")
	case DumpGuestMemoryFormatKdumpZlib:
		return json.Marshal("kdump-zlib")
	case DumpGuestMemoryFormatKdumpLzo:
		return json.Marshal("kdump-lzo")
	case DumpGuestMemoryFormatKdumpSnappy:
		return json.Marshal("kdump-snappy")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DumpGuestMemoryFormat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DumpGuestMemoryFormat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "elf":
		*e = DumpGuestMemoryFormatElf
	case "kdump-zlib":
		*e = DumpGuestMemoryFormatKdumpZlib
	case "kdump-lzo":
		*e = DumpGuestMemoryFormatKdumpLzo
	case "kdump-snappy":
		*e = DumpGuestMemoryFormatKdumpSnappy
	default:
		return fmt.Errorf("unknown enum value %q for DumpGuestMemoryFormat", s)
	}
	return nil
}

// DumpQueryResult -> DumpQueryResult (struct)

// DumpQueryResult implements the "DumpQueryResult" QMP API type.
type DumpQueryResult struct {
	Status    DumpStatus `json:"status"`
	Completed int64      `json:"completed"`
	Total     int64      `json:"total"`
}

// DumpStatus -> DumpStatus (enum)

// DumpStatus implements the "DumpStatus" QMP API type.
type DumpStatus int

// Known values of DumpStatus.
const (
	DumpStatusNone DumpStatus = iota
	DumpStatusActive
	DumpStatusCompleted
	DumpStatusFailed
)

// String implements fmt.Stringer.
func (e DumpStatus) String() string {
	switch e {
	case DumpStatusNone:
		return "none"
	case DumpStatusActive:
		return "active"
	case DumpStatusCompleted:
		return "completed"
	case DumpStatusFailed:
		return "failed"
	default:
		return fmt.Sprintf("DumpStatus(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DumpStatus) MarshalJSON() ([]byte, error) {
	switch e {
	case DumpStatusNone:
		return json.Marshal("none")
	case DumpStatusActive:
		return json.Marshal("active")
	case DumpStatusCompleted:
		return json.Marshal("completed")
	case DumpStatusFailed:
		return json.Marshal("failed")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DumpStatus", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DumpStatus) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = DumpStatusNone
	case "active":
		*e = DumpStatusActive
	case "completed":
		*e = DumpStatusCompleted
	case "failed":
		*e = DumpStatusFailed
	default:
		return fmt.Errorf("unknown enum value %q for DumpStatus", s)
	}
	return nil
}

// EventInfo -> EventInfo (struct)

// EventInfo implements the "EventInfo" QMP API type.
type EventInfo struct {
	Name string `json:"name"`
}

// FdsetFdInfo -> FdsetFDInfo (struct)

// FdsetFDInfo implements the "FdsetFdInfo" QMP API type.
type FdsetFDInfo struct {
	FD     int64   `json:"fd"`
	Opaque *string `json:"opaque,omitempty"`
}

// FdsetInfo -> FdsetInfo (struct)

// FdsetInfo implements the "FdsetInfo" QMP API type.
type FdsetInfo struct {
	FdsetID int64         `json:"fdset-id"`
	Fds     []FdsetFDInfo `json:"fds"`
}

// GICCapability -> GicCapability (struct)

// GicCapability implements the "GICCapability" QMP API type.
type GicCapability struct {
	Version  int64 `json:"version"`
	Emulated bool  `json:"emulated"`
	Kernel   bool  `json:"kernel"`
}

// EVENT GUEST_PANICKED

// GlusterServer -> GlusterServer (flat union)

// GlusterServer implements the "GlusterServer" QMP API type.
//
// Can be one of:
//   - GlusterServerTCP
//   - GlusterServerUnix
type GlusterServer interface {
	isGlusterServer()
}

// GlusterServerTCP is an implementation of GlusterServer.
type GlusterServerTCP struct {
	Host string  `json:"host"`
	Port string  `json:"port"`
	To   *uint16 `json:"to,omitempty"`
	Ipv4 *bool   `json:"ipv4,omitempty"`
	Ipv6 *bool   `json:"ipv6,omitempty"`
}

func (GlusterServerTCP) isGlusterServer() {}

// MarshalJSON implements json.Marshaler.
func (s GlusterServerTCP) MarshalJSON() ([]byte, error) {
	v := struct {
		Type GlusterTransport `json:"type"`
		GlusterServerTCP
	}{
		GlusterTransportTCP,
		s,
	}
	return json.Marshal(v)
}

// GlusterServerUnix is an implementation of GlusterServer.
type GlusterServerUnix struct {
	Path string `json:"path"`
}

func (GlusterServerUnix) isGlusterServer() {}

// MarshalJSON implements json.Marshaler.
func (s GlusterServerUnix) MarshalJSON() ([]byte, error) {
	v := struct {
		Type GlusterTransport `json:"type"`
		GlusterServerUnix
	}{
		GlusterTransportUnix,
		s,
	}
	return json.Marshal(v)
}

func decodeGlusterServer(bs json.RawMessage) (GlusterServer, error) {
	v := struct {
		Type GlusterTransport `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case GlusterTransportTCP:
		var ret GlusterServerTCP
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case GlusterTransportUnix:
		var ret GlusterServerUnix
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union GlusterServer", v.Type)
	}
}

// GlusterTransport -> GlusterTransport (enum)

// GlusterTransport implements the "GlusterTransport" QMP API type.
type GlusterTransport int

// Known values of GlusterTransport.
const (
	GlusterTransportUnix GlusterTransport = iota
	GlusterTransportTCP
)

// String implements fmt.Stringer.
func (e GlusterTransport) String() string {
	switch e {
	case GlusterTransportUnix:
		return "unix"
	case GlusterTransportTCP:
		return "tcp"
	default:
		return fmt.Sprintf("GlusterTransport(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e GlusterTransport) MarshalJSON() ([]byte, error) {
	switch e {
	case GlusterTransportUnix:
		return json.Marshal("unix")
	case GlusterTransportTCP:
		return json.Marshal("tcp")
	default:
		return nil, fmt.Errorf("unknown enum value %q for GlusterTransport", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *GlusterTransport) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "unix":
		*e = GlusterTransportUnix
	case "tcp":
		*e = GlusterTransportTCP
	default:
		return fmt.Errorf("unknown enum value %q for GlusterTransport", s)
	}
	return nil
}

// GuestPanicAction -> GuestPanicAction (enum)

// GuestPanicAction implements the "GuestPanicAction" QMP API type.
type GuestPanicAction int

// Known values of GuestPanicAction.
const (
	GuestPanicActionPause GuestPanicAction = iota
)

// String implements fmt.Stringer.
func (e GuestPanicAction) String() string {
	switch e {
	case GuestPanicActionPause:
		return "pause"
	default:
		return fmt.Sprintf("GuestPanicAction(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e GuestPanicAction) MarshalJSON() ([]byte, error) {
	switch e {
	case GuestPanicActionPause:
		return json.Marshal("pause")
	default:
		return nil, fmt.Errorf("unknown enum value %q for GuestPanicAction", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *GuestPanicAction) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "pause":
		*e = GuestPanicActionPause
	default:
		return fmt.Errorf("unknown enum value %q for GuestPanicAction", s)
	}
	return nil
}

// HostMemPolicy -> HostMemPolicy (enum)

// HostMemPolicy implements the "HostMemPolicy" QMP API type.
type HostMemPolicy int

// Known values of HostMemPolicy.
const (
	HostMemPolicyDefault HostMemPolicy = iota
	HostMemPolicyPreferred
	HostMemPolicyBind
	HostMemPolicyInterleave
)

// String implements fmt.Stringer.
func (e HostMemPolicy) String() string {
	switch e {
	case HostMemPolicyDefault:
		return "default"
	case HostMemPolicyPreferred:
		return "preferred"
	case HostMemPolicyBind:
		return "bind"
	case HostMemPolicyInterleave:
		return "interleave"
	default:
		return fmt.Sprintf("HostMemPolicy(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e HostMemPolicy) MarshalJSON() ([]byte, error) {
	switch e {
	case HostMemPolicyDefault:
		return json.Marshal("default")
	case HostMemPolicyPreferred:
		return json.Marshal("preferred")
	case HostMemPolicyBind:
		return json.Marshal("bind")
	case HostMemPolicyInterleave:
		return json.Marshal("interleave")
	default:
		return nil, fmt.Errorf("unknown enum value %q for HostMemPolicy", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *HostMemPolicy) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "default":
		*e = HostMemPolicyDefault
	case "preferred":
		*e = HostMemPolicyPreferred
	case "bind":
		*e = HostMemPolicyBind
	case "interleave":
		*e = HostMemPolicyInterleave
	default:
		return fmt.Errorf("unknown enum value %q for HostMemPolicy", s)
	}
	return nil
}

// HotpluggableCPU -> HotpluggableCPU (struct)

// HotpluggableCPU implements the "HotpluggableCPU" QMP API type.
type HotpluggableCPU struct {
	Type       string                `json:"type"`
	VcpusCount int64                 `json:"vcpus-count"`
	Props      CPUInstanceProperties `json:"props"`
	QomPath    *string               `json:"qom-path,omitempty"`
}

// IOThreadInfo -> IOThreadInfo (struct)

// IOThreadInfo implements the "IOThreadInfo" QMP API type.
type IOThreadInfo struct {
	ID       string `json:"id"`
	ThreadID int64  `json:"thread-id"`
}

// ImageInfo -> ImageInfo (struct)

// ImageInfo implements the "ImageInfo" QMP API type.
type ImageInfo struct {
	Filename              string            `json:"filename"`
	Format                string            `json:"format"`
	DirtyFlag             *bool             `json:"dirty-flag,omitempty"`
	ActualSize            *int64            `json:"actual-size,omitempty"`
	VirtualSize           int64             `json:"virtual-size"`
	ClusterSize           *int64            `json:"cluster-size,omitempty"`
	Encrypted             *bool             `json:"encrypted,omitempty"`
	Compressed            *bool             `json:"compressed,omitempty"`
	BackingFilename       *string           `json:"backing-filename,omitempty"`
	FullBackingFilename   *string           `json:"full-backing-filename,omitempty"`
	BackingFilenameFormat *string           `json:"backing-filename-format,omitempty"`
	Snapshots             []SnapshotInfo    `json:"snapshots,omitempty"`
	BackingImage          *ImageInfo        `json:"backing-image,omitempty"`
	FormatSpecific        ImageInfoSpecific `json:"format-specific,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *ImageInfo) UnmarshalJSON(bs []byte) error {
	v := struct {
		Filename              string          `json:"filename"`
		Format                string          `json:"format"`
		DirtyFlag             *bool           `json:"dirty-flag,omitempty"`
		ActualSize            *int64          `json:"actual-size,omitempty"`
		VirtualSize           int64           `json:"virtual-size"`
		ClusterSize           *int64          `json:"cluster-size,omitempty"`
		Encrypted             *bool           `json:"encrypted,omitempty"`
		Compressed            *bool           `json:"compressed,omitempty"`
		BackingFilename       *string         `json:"backing-filename,omitempty"`
		FullBackingFilename   *string         `json:"full-backing-filename,omitempty"`
		BackingFilenameFormat *string         `json:"backing-filename-format,omitempty"`
		Snapshots             []SnapshotInfo  `json:"snapshots,omitempty"`
		BackingImage          *ImageInfo      `json:"backing-image,omitempty"`
		FormatSpecific        json.RawMessage `json:"format-specific"`
	}{}
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return err
	}
	s.Filename = v.Filename
	s.Format = v.Format
	s.DirtyFlag = v.DirtyFlag
	s.ActualSize = v.ActualSize
	s.VirtualSize = v.VirtualSize
	s.ClusterSize = v.ClusterSize
	s.Encrypted = v.Encrypted
	s.Compressed = v.Compressed
	s.BackingFilename = v.BackingFilename
	s.FullBackingFilename = v.FullBackingFilename
	s.BackingFilenameFormat = v.BackingFilenameFormat
	s.Snapshots = v.Snapshots
	s.BackingImage = v.BackingImage
	if len(v.FormatSpecific) > 0 {
		s.FormatSpecific, err = decodeImageInfoSpecific(v.FormatSpecific)
		if err != nil {
			return err
		}
	} else {
		s.FormatSpecific = nil
	}

	return nil
}

// ImageInfoSpecific -> ImageInfoSpecific (simple union)

// ImageInfoSpecific implements the "ImageInfoSpecific" QMP API type.
//
// Can be one of:
//   - ImageInfoSpecificLUKS
//   - ImageInfoSpecificQcow2
//   - ImageInfoSpecificVMDK
type ImageInfoSpecific interface {
	isImageInfoSpecific()
}

// ImageInfoSpecificLUKS is an implementation of ImageInfoSpecific.
type ImageInfoSpecificLUKS QCryptoBlockInfoLUKS

func (ImageInfoSpecificLUKS) isImageInfoSpecific() {}

// MarshalJSON implements json.Marshaler.
func (s ImageInfoSpecificLUKS) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "luks",
		"data": s,
	}
	return json.Marshal(v)
}

// ImageInfoSpecificQcow2 is an implementation of ImageInfoSpecific.
type ImageInfoSpecificQcow2 ImageInfoSpecificQCow2

func (ImageInfoSpecificQcow2) isImageInfoSpecific() {}

// MarshalJSON implements json.Marshaler.
func (s ImageInfoSpecificQcow2) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "qcow2",
		"data": s,
	}
	return json.Marshal(v)
}

func (ImageInfoSpecificVMDK) isImageInfoSpecific() {}

// MarshalJSON implements json.Marshaler.
func (s ImageInfoSpecificVMDK) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "vmdk",
		"data": s,
	}
	return json.Marshal(v)
}

func decodeImageInfoSpecific(bs json.RawMessage) (ImageInfoSpecific, error) {
	v := struct {
		T string          `json:"type"`
		V json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.T {
	case "luks":
		var ret ImageInfoSpecificLUKS
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "qcow2":
		var ret ImageInfoSpecificQcow2
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "vmdk":
		var ret ImageInfoSpecificVMDK
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("unknown subtype %q for union ImageInfoSpecific", v.T)
	}
}

// ImageInfoSpecificQCow2 -> ImageInfoSpecificQCow2 (struct)

// ImageInfoSpecificQCow2 implements the "ImageInfoSpecificQCow2" QMP API type.
type ImageInfoSpecificQCow2 struct {
	Compat        string `json:"compat"`
	LazyRefcounts *bool  `json:"lazy-refcounts,omitempty"`
	Corrupt       *bool  `json:"corrupt,omitempty"`
	RefcountBits  int64  `json:"refcount-bits"`
}

// ImageInfoSpecificVmdk -> ImageInfoSpecificVMDK (struct)

// ImageInfoSpecificVMDK implements the "ImageInfoSpecificVmdk" QMP API type.
type ImageInfoSpecificVMDK struct {
	CreateType string      `json:"create-type"`
	Cid        int64       `json:"cid"`
	ParentCid  int64       `json:"parent-cid"`
	Extents    []ImageInfo `json:"extents"`
}

// InetSocketAddress -> InetSocketAddress (struct)

// InetSocketAddress implements the "InetSocketAddress" QMP API type.
type InetSocketAddress struct {
	Host string  `json:"host"`
	Port string  `json:"port"`
	To   *uint16 `json:"to,omitempty"`
	Ipv4 *bool   `json:"ipv4,omitempty"`
	Ipv6 *bool   `json:"ipv6,omitempty"`
}

// InputAxis -> InputAxis (enum)

// InputAxis implements the "InputAxis" QMP API type.
type InputAxis int

// Known values of InputAxis.
const (
	InputAxisX InputAxis = iota
	InputAxisY
)

// String implements fmt.Stringer.
func (e InputAxis) String() string {
	switch e {
	case InputAxisX:
		return "x"
	case InputAxisY:
		return "y"
	default:
		return fmt.Sprintf("InputAxis(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e InputAxis) MarshalJSON() ([]byte, error) {
	switch e {
	case InputAxisX:
		return json.Marshal("x")
	case InputAxisY:
		return json.Marshal("y")
	default:
		return nil, fmt.Errorf("unknown enum value %q for InputAxis", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *InputAxis) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "x":
		*e = InputAxisX
	case "y":
		*e = InputAxisY
	default:
		return fmt.Errorf("unknown enum value %q for InputAxis", s)
	}
	return nil
}

// InputBtnEvent -> InputBtnEvent (struct)

// InputBtnEvent implements the "InputBtnEvent" QMP API type.
type InputBtnEvent struct {
	Button InputButton `json:"button"`
	Down   bool        `json:"down"`
}

// InputButton -> InputButton (enum)

// InputButton implements the "InputButton" QMP API type.
type InputButton int

// Known values of InputButton.
const (
	InputButtonLeft InputButton = iota
	InputButtonMiddle
	InputButtonRight
	InputButtonWheelUp
	InputButtonWheelDown
)

// String implements fmt.Stringer.
func (e InputButton) String() string {
	switch e {
	case InputButtonLeft:
		return "left"
	case InputButtonMiddle:
		return "middle"
	case InputButtonRight:
		return "right"
	case InputButtonWheelUp:
		return "wheel-up"
	case InputButtonWheelDown:
		return "wheel-down"
	default:
		return fmt.Sprintf("InputButton(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e InputButton) MarshalJSON() ([]byte, error) {
	switch e {
	case InputButtonLeft:
		return json.Marshal("left")
	case InputButtonMiddle:
		return json.Marshal("middle")
	case InputButtonRight:
		return json.Marshal("right")
	case InputButtonWheelUp:
		return json.Marshal("wheel-up")
	case InputButtonWheelDown:
		return json.Marshal("wheel-down")
	default:
		return nil, fmt.Errorf("unknown enum value %q for InputButton", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *InputButton) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "left":
		*e = InputButtonLeft
	case "middle":
		*e = InputButtonMiddle
	case "right":
		*e = InputButtonRight
	case "wheel-up":
		*e = InputButtonWheelUp
	case "wheel-down":
		*e = InputButtonWheelDown
	default:
		return fmt.Errorf("unknown enum value %q for InputButton", s)
	}
	return nil
}

// InputEvent -> InputEvent (simple union)

// InputEvent implements the "InputEvent" QMP API type.
//
// Can be one of:
//   - InputEventAbs
//   - InputEventBtn
//   - InputEventKey
//   - InputEventRel
type InputEvent interface {
	isInputEvent()
}

// InputEventAbs is an implementation of InputEvent.
type InputEventAbs InputMoveEvent

func (InputEventAbs) isInputEvent() {}

// MarshalJSON implements json.Marshaler.
func (s InputEventAbs) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "abs",
		"data": s,
	}
	return json.Marshal(v)
}

// InputEventBtn is an implementation of InputEvent.
type InputEventBtn InputBtnEvent

func (InputEventBtn) isInputEvent() {}

// MarshalJSON implements json.Marshaler.
func (s InputEventBtn) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "btn",
		"data": s,
	}
	return json.Marshal(v)
}

// InputEventKey is an implementation of InputEvent.
type InputEventKey InputKeyEvent

func (InputEventKey) isInputEvent() {}

// MarshalJSON implements json.Marshaler.
func (s InputEventKey) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "key",
		"data": s,
	}
	return json.Marshal(v)
}

// InputEventRel is an implementation of InputEvent.
type InputEventRel InputMoveEvent

func (InputEventRel) isInputEvent() {}

// MarshalJSON implements json.Marshaler.
func (s InputEventRel) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "rel",
		"data": s,
	}
	return json.Marshal(v)
}

func decodeInputEvent(bs json.RawMessage) (InputEvent, error) {
	v := struct {
		T string          `json:"type"`
		V json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.T {
	case "abs":
		var ret InputEventAbs
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "btn":
		var ret InputEventBtn
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "key":
		var ret InputEventKey
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "rel":
		var ret InputEventRel
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("unknown subtype %q for union InputEvent", v.T)
	}
}

// InputKeyEvent -> InputKeyEvent (struct)

// InputKeyEvent implements the "InputKeyEvent" QMP API type.
type InputKeyEvent struct {
	Key  KeyValue `json:"key"`
	Down bool     `json:"down"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *InputKeyEvent) UnmarshalJSON(bs []byte) error {
	v := struct {
		Key  json.RawMessage `json:"key"`
		Down bool            `json:"down"`
	}{}
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return err
	}
	s.Key, err = decodeKeyValue(v.Key)
	if err != nil {
		return err
	}
	s.Down = v.Down

	return nil
}

// InputMoveEvent -> InputMoveEvent (struct)

// InputMoveEvent implements the "InputMoveEvent" QMP API type.
type InputMoveEvent struct {
	Axis  InputAxis `json:"axis"`
	Value int64     `json:"value"`
}

// IoOperationType -> IOOperationType (enum)

// IOOperationType implements the "IoOperationType" QMP API type.
type IOOperationType int

// Known values of IOOperationType.
const (
	IOOperationTypeRead IOOperationType = iota
	IOOperationTypeWrite
)

// String implements fmt.Stringer.
func (e IOOperationType) String() string {
	switch e {
	case IOOperationTypeRead:
		return "read"
	case IOOperationTypeWrite:
		return "write"
	default:
		return fmt.Sprintf("IOOperationType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e IOOperationType) MarshalJSON() ([]byte, error) {
	switch e {
	case IOOperationTypeRead:
		return json.Marshal("read")
	case IOOperationTypeWrite:
		return json.Marshal("write")
	default:
		return nil, fmt.Errorf("unknown enum value %q for IOOperationType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *IOOperationType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "read":
		*e = IOOperationTypeRead
	case "write":
		*e = IOOperationTypeWrite
	default:
		return fmt.Errorf("unknown enum value %q for IOOperationType", s)
	}
	return nil
}

// JSONType -> JSONType (enum)

// JSONType implements the "JSONType" QMP API type.
type JSONType int

// Known values of JSONType.
const (
	JSONTypeString JSONType = iota
	JSONTypefloat64
	JSONTypeint64
	JSONTypeBoolean
	JSONTypeNull
	JSONTypeObject
	JSONTypeArray
	JSONTypeValue
)

// String implements fmt.Stringer.
func (e JSONType) String() string {
	switch e {
	case JSONTypeString:
		return "string"
	case JSONTypefloat64:
		return "number"
	case JSONTypeint64:
		return "int"
	case JSONTypeBoolean:
		return "boolean"
	case JSONTypeNull:
		return "null"
	case JSONTypeObject:
		return "object"
	case JSONTypeArray:
		return "array"
	case JSONTypeValue:
		return "value"
	default:
		return fmt.Sprintf("JSONType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e JSONType) MarshalJSON() ([]byte, error) {
	switch e {
	case JSONTypeString:
		return json.Marshal("string")
	case JSONTypefloat64:
		return json.Marshal("number")
	case JSONTypeint64:
		return json.Marshal("int")
	case JSONTypeBoolean:
		return json.Marshal("boolean")
	case JSONTypeNull:
		return json.Marshal("null")
	case JSONTypeObject:
		return json.Marshal("object")
	case JSONTypeArray:
		return json.Marshal("array")
	case JSONTypeValue:
		return json.Marshal("value")
	default:
		return nil, fmt.Errorf("unknown enum value %q for JSONType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *JSONType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "string":
		*e = JSONTypeString
	case "number":
		*e = JSONTypefloat64
	case "int":
		*e = JSONTypeint64
	case "boolean":
		*e = JSONTypeBoolean
	case "null":
		*e = JSONTypeNull
	case "object":
		*e = JSONTypeObject
	case "array":
		*e = JSONTypeArray
	case "value":
		*e = JSONTypeValue
	default:
		return fmt.Errorf("unknown enum value %q for JSONType", s)
	}
	return nil
}

// KeyValue -> KeyValue (simple union)

// KeyValue implements the "KeyValue" QMP API type.
//
// Can be one of:
//   - KeyValuefloat64
//   - KeyValueQcode
type KeyValue interface {
	isKeyValue()
}

// KeyValuefloat64 is an implementation of KeyValue.
type KeyValuefloat64 int64

func (KeyValuefloat64) isKeyValue() {}

// MarshalJSON implements json.Marshaler.
func (s KeyValuefloat64) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "number",
		"data": s,
	}
	return json.Marshal(v)
}

// KeyValueQcode is an implementation of KeyValue.
type KeyValueQcode QKeyCode

func (KeyValueQcode) isKeyValue() {}

// MarshalJSON implements json.Marshaler.
func (s KeyValueQcode) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "qcode",
		"data": s,
	}
	return json.Marshal(v)
}

func decodeKeyValue(bs json.RawMessage) (KeyValue, error) {
	v := struct {
		T string          `json:"type"`
		V json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.T {
	case "number":
		var ret KeyValuefloat64
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "qcode":
		var ret KeyValueQcode
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("unknown subtype %q for union KeyValue", v.T)
	}
}

// KvmInfo -> KVMInfo (struct)

// KVMInfo implements the "KvmInfo" QMP API type.
type KVMInfo struct {
	Enabled bool `json:"enabled"`
	Present bool `json:"present"`
}

// EVENT MEM_UNPLUG_ERROR

// EVENT MIGRATION

// EVENT MIGRATION_PASS

// MachineInfo -> MachineInfo (struct)

// MachineInfo implements the "MachineInfo" QMP API type.
type MachineInfo struct {
	Name             string  `json:"name"`
	Alias            *string `json:"alias,omitempty"`
	IsDefault        *bool   `json:"is-default,omitempty"`
	CPUMax           int64   `json:"cpu-max"`
	HotpluggableCpus bool    `json:"hotpluggable-cpus"`
}

// Memdev -> Memdev (struct)

// Memdev implements the "Memdev" QMP API type.
type Memdev struct {
	Size      uint64        `json:"size"`
	Merge     bool          `json:"merge"`
	Dump      bool          `json:"dump"`
	Prealloc  bool          `json:"prealloc"`
	HostNodes []uint16      `json:"host-nodes"`
	Policy    HostMemPolicy `json:"policy"`
}

// MemoryDeviceInfo -> MemoryDeviceInfo (simple union)

// MemoryDeviceInfo implements the "MemoryDeviceInfo" QMP API type.
//
// Can be one of:
//   - MemoryDeviceInfoDimm
type MemoryDeviceInfo interface {
	isMemoryDeviceInfo()
}

// MemoryDeviceInfoDimm is an implementation of MemoryDeviceInfo.
type MemoryDeviceInfoDimm PcdimmDeviceInfo

func (MemoryDeviceInfoDimm) isMemoryDeviceInfo() {}

// MarshalJSON implements json.Marshaler.
func (s MemoryDeviceInfoDimm) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "dimm",
		"data": s,
	}
	return json.Marshal(v)
}

func decodeMemoryDeviceInfo(bs json.RawMessage) (MemoryDeviceInfo, error) {
	v := struct {
		T string          `json:"type"`
		V json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.T {
	case "dimm":
		var ret MemoryDeviceInfoDimm
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("unknown subtype %q for union MemoryDeviceInfo", v.T)
	}
}

// MigrationCapability -> MigrationCapability (enum)

// MigrationCapability implements the "MigrationCapability" QMP API type.
type MigrationCapability int

// Known values of MigrationCapability.
const (
	MigrationCapabilityXbzrle MigrationCapability = iota
	MigrationCapabilityRdmaPinAll
	MigrationCapabilityAutoConverge
	MigrationCapabilityZeroBlocks
	MigrationCapabilityCompress
	MigrationCapabilityEvents
	MigrationCapabilityPostcopyRAM
)

// String implements fmt.Stringer.
func (e MigrationCapability) String() string {
	switch e {
	case MigrationCapabilityXbzrle:
		return "xbzrle"
	case MigrationCapabilityRdmaPinAll:
		return "rdma-pin-all"
	case MigrationCapabilityAutoConverge:
		return "auto-converge"
	case MigrationCapabilityZeroBlocks:
		return "zero-blocks"
	case MigrationCapabilityCompress:
		return "compress"
	case MigrationCapabilityEvents:
		return "events"
	case MigrationCapabilityPostcopyRAM:
		return "postcopy-ram"
	default:
		return fmt.Sprintf("MigrationCapability(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e MigrationCapability) MarshalJSON() ([]byte, error) {
	switch e {
	case MigrationCapabilityXbzrle:
		return json.Marshal("xbzrle")
	case MigrationCapabilityRdmaPinAll:
		return json.Marshal("rdma-pin-all")
	case MigrationCapabilityAutoConverge:
		return json.Marshal("auto-converge")
	case MigrationCapabilityZeroBlocks:
		return json.Marshal("zero-blocks")
	case MigrationCapabilityCompress:
		return json.Marshal("compress")
	case MigrationCapabilityEvents:
		return json.Marshal("events")
	case MigrationCapabilityPostcopyRAM:
		return json.Marshal("postcopy-ram")
	default:
		return nil, fmt.Errorf("unknown enum value %q for MigrationCapability", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *MigrationCapability) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "xbzrle":
		*e = MigrationCapabilityXbzrle
	case "rdma-pin-all":
		*e = MigrationCapabilityRdmaPinAll
	case "auto-converge":
		*e = MigrationCapabilityAutoConverge
	case "zero-blocks":
		*e = MigrationCapabilityZeroBlocks
	case "compress":
		*e = MigrationCapabilityCompress
	case "events":
		*e = MigrationCapabilityEvents
	case "postcopy-ram":
		*e = MigrationCapabilityPostcopyRAM
	default:
		return fmt.Errorf("unknown enum value %q for MigrationCapability", s)
	}
	return nil
}

// MigrationCapabilityStatus -> MigrationCapabilityStatus (struct)

// MigrationCapabilityStatus implements the "MigrationCapabilityStatus" QMP API type.
type MigrationCapabilityStatus struct {
	Capability MigrationCapability `json:"capability"`
	State      bool                `json:"state"`
}

// MigrationInfo -> MigrationInfo (struct)

// MigrationInfo implements the "MigrationInfo" QMP API type.
type MigrationInfo struct {
	Status                *MigrationStatus  `json:"status,omitempty"`
	RAM                   *MigrationStats   `json:"ram,omitempty"`
	Disk                  *MigrationStats   `json:"disk,omitempty"`
	XbzrleCache           *XbzrleCacheStats `json:"xbzrle-cache,omitempty"`
	TotalTime             *int64            `json:"total-time,omitempty"`
	ExpectedDowntime      *int64            `json:"expected-downtime,omitempty"`
	Downtime              *int64            `json:"downtime,omitempty"`
	SetupTime             *int64            `json:"setup-time,omitempty"`
	CPUThrottlePercentage *int64            `json:"cpu-throttle-percentage,omitempty"`
	ErrorDesc             *string           `json:"error-desc,omitempty"`
}

// MigrationParameters -> MigrationParameters (struct)

// MigrationParameters implements the "MigrationParameters" QMP API type.
type MigrationParameters struct {
	CompressLevel        *int64  `json:"compress-level,omitempty"`
	CompressThreads      *int64  `json:"compress-threads,omitempty"`
	DecompressThreads    *int64  `json:"decompress-threads,omitempty"`
	CPUThrottleInitial   *int64  `json:"cpu-throttle-initial,omitempty"`
	CPUThrottleIncrement *int64  `json:"cpu-throttle-increment,omitempty"`
	TLSCreds             *string `json:"tls-creds,omitempty"`
	TLSHostname          *string `json:"tls-hostname,omitempty"`
	MaxBandwidth         *int64  `json:"max-bandwidth,omitempty"`
	DowntimeLimit        *int64  `json:"downtime-limit,omitempty"`
}

// MigrationStats -> MigrationStats (struct)

// MigrationStats implements the "MigrationStats" QMP API type.
type MigrationStats struct {
	Transferred      int64   `json:"transferred"`
	Remaining        int64   `json:"remaining"`
	Total            int64   `json:"total"`
	Duplicate        int64   `json:"duplicate"`
	Skipped          int64   `json:"skipped"`
	Normal           int64   `json:"normal"`
	NormalBytes      int64   `json:"normal-bytes"`
	DirtyPagesRate   int64   `json:"dirty-pages-rate"`
	Mbps             float64 `json:"mbps"`
	DirtySyncCount   int64   `json:"dirty-sync-count"`
	PostcopyRequests int64   `json:"postcopy-requests"`
}

// MigrationStatus -> MigrationStatus (enum)

// MigrationStatus implements the "MigrationStatus" QMP API type.
type MigrationStatus int

// Known values of MigrationStatus.
const (
	MigrationStatusNone MigrationStatus = iota
	MigrationStatusSetup
	MigrationStatusCancelling
	MigrationStatusCancelled
	MigrationStatusActive
	MigrationStatusPostcopyActive
	MigrationStatusCompleted
	MigrationStatusFailed
)

// String implements fmt.Stringer.
func (e MigrationStatus) String() string {
	switch e {
	case MigrationStatusNone:
		return "none"
	case MigrationStatusSetup:
		return "setup"
	case MigrationStatusCancelling:
		return "cancelling"
	case MigrationStatusCancelled:
		return "cancelled"
	case MigrationStatusActive:
		return "active"
	case MigrationStatusPostcopyActive:
		return "postcopy-active"
	case MigrationStatusCompleted:
		return "completed"
	case MigrationStatusFailed:
		return "failed"
	default:
		return fmt.Sprintf("MigrationStatus(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e MigrationStatus) MarshalJSON() ([]byte, error) {
	switch e {
	case MigrationStatusNone:
		return json.Marshal("none")
	case MigrationStatusSetup:
		return json.Marshal("setup")
	case MigrationStatusCancelling:
		return json.Marshal("cancelling")
	case MigrationStatusCancelled:
		return json.Marshal("cancelled")
	case MigrationStatusActive:
		return json.Marshal("active")
	case MigrationStatusPostcopyActive:
		return json.Marshal("postcopy-active")
	case MigrationStatusCompleted:
		return json.Marshal("completed")
	case MigrationStatusFailed:
		return json.Marshal("failed")
	default:
		return nil, fmt.Errorf("unknown enum value %q for MigrationStatus", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *MigrationStatus) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = MigrationStatusNone
	case "setup":
		*e = MigrationStatusSetup
	case "cancelling":
		*e = MigrationStatusCancelling
	case "cancelled":
		*e = MigrationStatusCancelled
	case "active":
		*e = MigrationStatusActive
	case "postcopy-active":
		*e = MigrationStatusPostcopyActive
	case "completed":
		*e = MigrationStatusCompleted
	case "failed":
		*e = MigrationStatusFailed
	default:
		return fmt.Errorf("unknown enum value %q for MigrationStatus", s)
	}
	return nil
}

// MirrorSyncMode -> MirrorSyncMode (enum)

// MirrorSyncMode implements the "MirrorSyncMode" QMP API type.
type MirrorSyncMode int

// Known values of MirrorSyncMode.
const (
	MirrorSyncModeTop MirrorSyncMode = iota
	MirrorSyncModeFull
	MirrorSyncModeNone
	MirrorSyncModeIncremental
)

// String implements fmt.Stringer.
func (e MirrorSyncMode) String() string {
	switch e {
	case MirrorSyncModeTop:
		return "top"
	case MirrorSyncModeFull:
		return "full"
	case MirrorSyncModeNone:
		return "none"
	case MirrorSyncModeIncremental:
		return "incremental"
	default:
		return fmt.Sprintf("MirrorSyncMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e MirrorSyncMode) MarshalJSON() ([]byte, error) {
	switch e {
	case MirrorSyncModeTop:
		return json.Marshal("top")
	case MirrorSyncModeFull:
		return json.Marshal("full")
	case MirrorSyncModeNone:
		return json.Marshal("none")
	case MirrorSyncModeIncremental:
		return json.Marshal("incremental")
	default:
		return nil, fmt.Errorf("unknown enum value %q for MirrorSyncMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *MirrorSyncMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "top":
		*e = MirrorSyncModeTop
	case "full":
		*e = MirrorSyncModeFull
	case "none":
		*e = MirrorSyncModeNone
	case "incremental":
		*e = MirrorSyncModeIncremental
	default:
		return fmt.Errorf("unknown enum value %q for MirrorSyncMode", s)
	}
	return nil
}

// MouseInfo -> MouseInfo (struct)

// MouseInfo implements the "MouseInfo" QMP API type.
type MouseInfo struct {
	Name     string `json:"name"`
	Index    int64  `json:"index"`
	Current  bool   `json:"current"`
	Absolute bool   `json:"absolute"`
}

// EVENT NIC_RX_FILTER_CHANGED

// NameInfo -> NameInfo (struct)

// NameInfo implements the "NameInfo" QMP API type.
type NameInfo struct {
	Name *string `json:"name,omitempty"`
}

// NetworkAddressFamily -> NetworkAddressFamily (enum)

// NetworkAddressFamily implements the "NetworkAddressFamily" QMP API type.
type NetworkAddressFamily int

// Known values of NetworkAddressFamily.
const (
	NetworkAddressFamilyIpv4 NetworkAddressFamily = iota
	NetworkAddressFamilyIpv6
	NetworkAddressFamilyUnix
	NetworkAddressFamilyUnknown
)

// String implements fmt.Stringer.
func (e NetworkAddressFamily) String() string {
	switch e {
	case NetworkAddressFamilyIpv4:
		return "ipv4"
	case NetworkAddressFamilyIpv6:
		return "ipv6"
	case NetworkAddressFamilyUnix:
		return "unix"
	case NetworkAddressFamilyUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("NetworkAddressFamily(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e NetworkAddressFamily) MarshalJSON() ([]byte, error) {
	switch e {
	case NetworkAddressFamilyIpv4:
		return json.Marshal("ipv4")
	case NetworkAddressFamilyIpv6:
		return json.Marshal("ipv6")
	case NetworkAddressFamilyUnix:
		return json.Marshal("unix")
	case NetworkAddressFamilyUnknown:
		return json.Marshal("unknown")
	default:
		return nil, fmt.Errorf("unknown enum value %q for NetworkAddressFamily", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NetworkAddressFamily) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "ipv4":
		*e = NetworkAddressFamilyIpv4
	case "ipv6":
		*e = NetworkAddressFamilyIpv6
	case "unix":
		*e = NetworkAddressFamilyUnix
	case "unknown":
		*e = NetworkAddressFamilyUnknown
	default:
		return fmt.Errorf("unknown enum value %q for NetworkAddressFamily", s)
	}
	return nil
}

// NewImageMode -> NewImageMode (enum)

// NewImageMode implements the "NewImageMode" QMP API type.
type NewImageMode int

// Known values of NewImageMode.
const (
	NewImageModeExisting NewImageMode = iota
	NewImageModeAbsolutePaths
)

// String implements fmt.Stringer.
func (e NewImageMode) String() string {
	switch e {
	case NewImageModeExisting:
		return "existing"
	case NewImageModeAbsolutePaths:
		return "absolute-paths"
	default:
		return fmt.Sprintf("NewImageMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e NewImageMode) MarshalJSON() ([]byte, error) {
	switch e {
	case NewImageModeExisting:
		return json.Marshal("existing")
	case NewImageModeAbsolutePaths:
		return json.Marshal("absolute-paths")
	default:
		return nil, fmt.Errorf("unknown enum value %q for NewImageMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NewImageMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "existing":
		*e = NewImageModeExisting
	case "absolute-paths":
		*e = NewImageModeAbsolutePaths
	default:
		return fmt.Errorf("unknown enum value %q for NewImageMode", s)
	}
	return nil
}

// ObjectPropertyInfo -> ObjectPropertyInfo (struct)

// ObjectPropertyInfo implements the "ObjectPropertyInfo" QMP API type.
type ObjectPropertyInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// ObjectTypeInfo -> ObjectTypeInfo (struct)

// ObjectTypeInfo implements the "ObjectTypeInfo" QMP API type.
type ObjectTypeInfo struct {
	Name string `json:"name"`
}

// PCDIMMDeviceInfo -> PcdimmDeviceInfo (struct)

// PcdimmDeviceInfo implements the "PCDIMMDeviceInfo" QMP API type.
type PcdimmDeviceInfo struct {
	ID           *string `json:"id,omitempty"`
	Addr         int64   `json:"addr"`
	Size         int64   `json:"size"`
	Slot         int64   `json:"slot"`
	Node         int64   `json:"node"`
	Memdev       string  `json:"memdev"`
	Hotplugged   bool    `json:"hotplugged"`
	Hotpluggable bool    `json:"hotpluggable"`
}

// EVENT POWERDOWN

// PciBridgeInfo -> PCIBridgeInfo (struct)

// PCIBridgeInfo implements the "PciBridgeInfo" QMP API type.
type PCIBridgeInfo struct {
	Bus     PCIBusInfo      `json:"bus"`
	Devices []PCIDeviceInfo `json:"devices,omitempty"`
}

// PciBusInfo -> PCIBusInfo (struct)

// PCIBusInfo implements the "PciBusInfo" QMP API type.
type PCIBusInfo struct {
	Number            int64          `json:"number"`
	Secondary         int64          `json:"secondary"`
	Subordinate       int64          `json:"subordinate"`
	IORange           PCIMemoryRange `json:"io_range"`
	MemoryRange       PCIMemoryRange `json:"memory_range"`
	PrefetchableRange PCIMemoryRange `json:"prefetchable_range"`
}

// PciDeviceClass -> PCIDeviceClass (struct)

// PCIDeviceClass implements the "PciDeviceClass" QMP API type.
type PCIDeviceClass struct {
	Desc  *string `json:"desc,omitempty"`
	Class int64   `json:"class"`
}

// PciDeviceId -> PCIDeviceID (struct)

// PCIDeviceID implements the "PciDeviceId" QMP API type.
type PCIDeviceID struct {
	Device int64 `json:"device"`
	Vendor int64 `json:"vendor"`
}

// PciDeviceInfo -> PCIDeviceInfo (struct)

// PCIDeviceInfo implements the "PciDeviceInfo" QMP API type.
type PCIDeviceInfo struct {
	Bus       int64             `json:"bus"`
	Slot      int64             `json:"slot"`
	Function  int64             `json:"function"`
	ClassInfo PCIDeviceClass    `json:"class_info"`
	ID        PCIDeviceID       `json:"id"`
	Irq       *int64            `json:"irq,omitempty"`
	QdevID    string            `json:"qdev_id"`
	PCIBridge *PCIBridgeInfo    `json:"pci_bridge,omitempty"`
	Regions   []PCIMemoryRegion `json:"regions"`
}

// PciInfo -> PCIInfo (struct)

// PCIInfo implements the "PciInfo" QMP API type.
type PCIInfo struct {
	Bus     int64           `json:"bus"`
	Devices []PCIDeviceInfo `json:"devices"`
}

// PciMemoryRange -> PCIMemoryRange (struct)

// PCIMemoryRange implements the "PciMemoryRange" QMP API type.
type PCIMemoryRange struct {
	Base  int64 `json:"base"`
	Limit int64 `json:"limit"`
}

// PciMemoryRegion -> PCIMemoryRegion (struct)

// PCIMemoryRegion implements the "PciMemoryRegion" QMP API type.
type PCIMemoryRegion struct {
	Bar       int64  `json:"bar"`
	Type      string `json:"type"`
	Address   int64  `json:"address"`
	Size      int64  `json:"size"`
	Prefetch  *bool  `json:"prefetch,omitempty"`
	MemType64 *bool  `json:"mem_type_64,omitempty"`
}

// QCryptoBlockInfoLUKS -> QCryptoBlockInfoLUKS (struct)

// QCryptoBlockInfoLUKS implements the "QCryptoBlockInfoLUKS" QMP API type.
type QCryptoBlockInfoLUKS struct {
	CipherAlg      QCryptoCipherAlgorithm     `json:"cipher-alg"`
	CipherMode     QCryptoCipherMode          `json:"cipher-mode"`
	IvgenAlg       QCryptoIvGenAlgorithm      `json:"ivgen-alg"`
	IvgenHashAlg   *QCryptoHashAlgorithm      `json:"ivgen-hash-alg,omitempty"`
	HashAlg        QCryptoHashAlgorithm       `json:"hash-alg"`
	PayloadOffset  int64                      `json:"payload-offset"`
	MasterKeyIters int64                      `json:"master-key-iters"`
	UUID           string                     `json:"uuid"`
	Slots          []QCryptoBlockInfoLUKSSlot `json:"slots"`
}

// QCryptoBlockInfoLUKSSlot -> QCryptoBlockInfoLUKSSlot (struct)

// QCryptoBlockInfoLUKSSlot implements the "QCryptoBlockInfoLUKSSlot" QMP API type.
type QCryptoBlockInfoLUKSSlot struct {
	Active    bool   `json:"active"`
	Iters     *int64 `json:"iters,omitempty"`
	Stripes   *int64 `json:"stripes,omitempty"`
	KeyOffset int64  `json:"key-offset"`
}

// QCryptoCipherAlgorithm -> QCryptoCipherAlgorithm (enum)

// QCryptoCipherAlgorithm implements the "QCryptoCipherAlgorithm" QMP API type.
type QCryptoCipherAlgorithm int

// Known values of QCryptoCipherAlgorithm.
const (
	QCryptoCipherAlgorithmAes128 QCryptoCipherAlgorithm = iota
	QCryptoCipherAlgorithmAes192
	QCryptoCipherAlgorithmAes256
	QCryptoCipherAlgorithmDesRfb
	QCryptoCipherAlgorithmCast5128
	QCryptoCipherAlgorithmSerpent128
	QCryptoCipherAlgorithmSerpent192
	QCryptoCipherAlgorithmSerpent256
	QCryptoCipherAlgorithmTwofish128
	QCryptoCipherAlgorithmTwofish192
	QCryptoCipherAlgorithmTwofish256
)

// String implements fmt.Stringer.
func (e QCryptoCipherAlgorithm) String() string {
	switch e {
	case QCryptoCipherAlgorithmAes128:
		return "aes-128"
	case QCryptoCipherAlgorithmAes192:
		return "aes-192"
	case QCryptoCipherAlgorithmAes256:
		return "aes-256"
	case QCryptoCipherAlgorithmDesRfb:
		return "des-rfb"
	case QCryptoCipherAlgorithmCast5128:
		return "cast5-128"
	case QCryptoCipherAlgorithmSerpent128:
		return "serpent-128"
	case QCryptoCipherAlgorithmSerpent192:
		return "serpent-192"
	case QCryptoCipherAlgorithmSerpent256:
		return "serpent-256"
	case QCryptoCipherAlgorithmTwofish128:
		return "twofish-128"
	case QCryptoCipherAlgorithmTwofish192:
		return "twofish-192"
	case QCryptoCipherAlgorithmTwofish256:
		return "twofish-256"
	default:
		return fmt.Sprintf("QCryptoCipherAlgorithm(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QCryptoCipherAlgorithm) MarshalJSON() ([]byte, error) {
	switch e {
	case QCryptoCipherAlgorithmAes128:
		return json.Marshal("aes-128")
	case QCryptoCipherAlgorithmAes192:
		return json.Marshal("aes-192")
	case QCryptoCipherAlgorithmAes256:
		return json.Marshal("aes-256")
	case QCryptoCipherAlgorithmDesRfb:
		return json.Marshal("des-rfb")
	case QCryptoCipherAlgorithmCast5128:
		return json.Marshal("cast5-128")
	case QCryptoCipherAlgorithmSerpent128:
		return json.Marshal("serpent-128")
	case QCryptoCipherAlgorithmSerpent192:
		return json.Marshal("serpent-192")
	case QCryptoCipherAlgorithmSerpent256:
		return json.Marshal("serpent-256")
	case QCryptoCipherAlgorithmTwofish128:
		return json.Marshal("twofish-128")
	case QCryptoCipherAlgorithmTwofish192:
		return json.Marshal("twofish-192")
	case QCryptoCipherAlgorithmTwofish256:
		return json.Marshal("twofish-256")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QCryptoCipherAlgorithm", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QCryptoCipherAlgorithm) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "aes-128":
		*e = QCryptoCipherAlgorithmAes128
	case "aes-192":
		*e = QCryptoCipherAlgorithmAes192
	case "aes-256":
		*e = QCryptoCipherAlgorithmAes256
	case "des-rfb":
		*e = QCryptoCipherAlgorithmDesRfb
	case "cast5-128":
		*e = QCryptoCipherAlgorithmCast5128
	case "serpent-128":
		*e = QCryptoCipherAlgorithmSerpent128
	case "serpent-192":
		*e = QCryptoCipherAlgorithmSerpent192
	case "serpent-256":
		*e = QCryptoCipherAlgorithmSerpent256
	case "twofish-128":
		*e = QCryptoCipherAlgorithmTwofish128
	case "twofish-192":
		*e = QCryptoCipherAlgorithmTwofish192
	case "twofish-256":
		*e = QCryptoCipherAlgorithmTwofish256
	default:
		return fmt.Errorf("unknown enum value %q for QCryptoCipherAlgorithm", s)
	}
	return nil
}

// QCryptoCipherMode -> QCryptoCipherMode (enum)

// QCryptoCipherMode implements the "QCryptoCipherMode" QMP API type.
type QCryptoCipherMode int

// Known values of QCryptoCipherMode.
const (
	QCryptoCipherModeEcb QCryptoCipherMode = iota
	QCryptoCipherModeCbc
	QCryptoCipherModeXts
)

// String implements fmt.Stringer.
func (e QCryptoCipherMode) String() string {
	switch e {
	case QCryptoCipherModeEcb:
		return "ecb"
	case QCryptoCipherModeCbc:
		return "cbc"
	case QCryptoCipherModeXts:
		return "xts"
	default:
		return fmt.Sprintf("QCryptoCipherMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QCryptoCipherMode) MarshalJSON() ([]byte, error) {
	switch e {
	case QCryptoCipherModeEcb:
		return json.Marshal("ecb")
	case QCryptoCipherModeCbc:
		return json.Marshal("cbc")
	case QCryptoCipherModeXts:
		return json.Marshal("xts")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QCryptoCipherMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QCryptoCipherMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "ecb":
		*e = QCryptoCipherModeEcb
	case "cbc":
		*e = QCryptoCipherModeCbc
	case "xts":
		*e = QCryptoCipherModeXts
	default:
		return fmt.Errorf("unknown enum value %q for QCryptoCipherMode", s)
	}
	return nil
}

// QCryptoHashAlgorithm -> QCryptoHashAlgorithm (enum)

// QCryptoHashAlgorithm implements the "QCryptoHashAlgorithm" QMP API type.
type QCryptoHashAlgorithm int

// Known values of QCryptoHashAlgorithm.
const (
	QCryptoHashAlgorithmMd5 QCryptoHashAlgorithm = iota
	QCryptoHashAlgorithmSha1
	QCryptoHashAlgorithmSha224
	QCryptoHashAlgorithmSha256
	QCryptoHashAlgorithmSha384
	QCryptoHashAlgorithmSha512
	QCryptoHashAlgorithmRipemd160
)

// String implements fmt.Stringer.
func (e QCryptoHashAlgorithm) String() string {
	switch e {
	case QCryptoHashAlgorithmMd5:
		return "md5"
	case QCryptoHashAlgorithmSha1:
		return "sha1"
	case QCryptoHashAlgorithmSha224:
		return "sha224"
	case QCryptoHashAlgorithmSha256:
		return "sha256"
	case QCryptoHashAlgorithmSha384:
		return "sha384"
	case QCryptoHashAlgorithmSha512:
		return "sha512"
	case QCryptoHashAlgorithmRipemd160:
		return "ripemd160"
	default:
		return fmt.Sprintf("QCryptoHashAlgorithm(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QCryptoHashAlgorithm) MarshalJSON() ([]byte, error) {
	switch e {
	case QCryptoHashAlgorithmMd5:
		return json.Marshal("md5")
	case QCryptoHashAlgorithmSha1:
		return json.Marshal("sha1")
	case QCryptoHashAlgorithmSha224:
		return json.Marshal("sha224")
	case QCryptoHashAlgorithmSha256:
		return json.Marshal("sha256")
	case QCryptoHashAlgorithmSha384:
		return json.Marshal("sha384")
	case QCryptoHashAlgorithmSha512:
		return json.Marshal("sha512")
	case QCryptoHashAlgorithmRipemd160:
		return json.Marshal("ripemd160")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QCryptoHashAlgorithm", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QCryptoHashAlgorithm) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "md5":
		*e = QCryptoHashAlgorithmMd5
	case "sha1":
		*e = QCryptoHashAlgorithmSha1
	case "sha224":
		*e = QCryptoHashAlgorithmSha224
	case "sha256":
		*e = QCryptoHashAlgorithmSha256
	case "sha384":
		*e = QCryptoHashAlgorithmSha384
	case "sha512":
		*e = QCryptoHashAlgorithmSha512
	case "ripemd160":
		*e = QCryptoHashAlgorithmRipemd160
	default:
		return fmt.Errorf("unknown enum value %q for QCryptoHashAlgorithm", s)
	}
	return nil
}

// QCryptoIVGenAlgorithm -> QCryptoIvGenAlgorithm (enum)

// QCryptoIvGenAlgorithm implements the "QCryptoIVGenAlgorithm" QMP API type.
type QCryptoIvGenAlgorithm int

// Known values of QCryptoIvGenAlgorithm.
const (
	QCryptoIvGenAlgorithmPlain QCryptoIvGenAlgorithm = iota
	QCryptoIvGenAlgorithmPlain64
	QCryptoIvGenAlgorithmEssiv
)

// String implements fmt.Stringer.
func (e QCryptoIvGenAlgorithm) String() string {
	switch e {
	case QCryptoIvGenAlgorithmPlain:
		return "plain"
	case QCryptoIvGenAlgorithmPlain64:
		return "plain64"
	case QCryptoIvGenAlgorithmEssiv:
		return "essiv"
	default:
		return fmt.Sprintf("QCryptoIvGenAlgorithm(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QCryptoIvGenAlgorithm) MarshalJSON() ([]byte, error) {
	switch e {
	case QCryptoIvGenAlgorithmPlain:
		return json.Marshal("plain")
	case QCryptoIvGenAlgorithmPlain64:
		return json.Marshal("plain64")
	case QCryptoIvGenAlgorithmEssiv:
		return json.Marshal("essiv")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QCryptoIvGenAlgorithm", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QCryptoIvGenAlgorithm) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "plain":
		*e = QCryptoIvGenAlgorithmPlain
	case "plain64":
		*e = QCryptoIvGenAlgorithmPlain64
	case "essiv":
		*e = QCryptoIvGenAlgorithmEssiv
	default:
		return fmt.Errorf("unknown enum value %q for QCryptoIvGenAlgorithm", s)
	}
	return nil
}

// QKeyCode -> QKeyCode (enum)

// QKeyCode implements the "QKeyCode" QMP API type.
type QKeyCode int

// Known values of QKeyCode.
const (
	QKeyCodeUnmapped QKeyCode = iota
	QKeyCodeShift
	QKeyCodeShiftR
	QKeyCodeAlt
	QKeyCodeAltR
	QKeyCodeAltgr
	QKeyCodeAltgrR
	QKeyCodeCtrl
	QKeyCodeCtrlR
	QKeyCodeMenu
	QKeyCodeEsc
	QKeyCode1
	QKeyCode2
	QKeyCode3
	QKeyCode4
	QKeyCode5
	QKeyCode6
	QKeyCode7
	QKeyCode8
	QKeyCode9
	QKeyCode0
	QKeyCodeMinus
	QKeyCodeEqual
	QKeyCodeBackspace
	QKeyCodeTab
	QKeyCodeQ
	QKeyCodeW
	QKeyCodeE
	QKeyCodeR
	QKeyCodeT
	QKeyCodeY
	QKeyCodeU
	QKeyCodeI
	QKeyCodeO
	QKeyCodeP
	QKeyCodeBracketLeft
	QKeyCodeBracketRight
	QKeyCodeRet
	QKeyCodeA
	QKeyCodeS
	QKeyCodeD
	QKeyCodeF
	QKeyCodeG
	QKeyCodeH
	QKeyCodeJ
	QKeyCodeK
	QKeyCodeL
	QKeyCodeSemicolon
	QKeyCodeApostrophe
	QKeyCodeGraveAccent
	QKeyCodeBackslash
	QKeyCodeZ
	QKeyCodeX
	QKeyCodeC
	QKeyCodeV
	QKeyCodeB
	QKeyCodeN
	QKeyCodeM
	QKeyCodeComma
	QKeyCodeDot
	QKeyCodeSlash
	QKeyCodeAsterisk
	QKeyCodeSpc
	QKeyCodeCapsLock
	QKeyCodeF1
	QKeyCodeF2
	QKeyCodeF3
	QKeyCodeF4
	QKeyCodeF5
	QKeyCodeF6
	QKeyCodeF7
	QKeyCodeF8
	QKeyCodeF9
	QKeyCodeF10
	QKeyCodeNumLock
	QKeyCodeScrollLock
	QKeyCodeKpDivide
	QKeyCodeKpMultiply
	QKeyCodeKpSubtract
	QKeyCodeKpAdd
	QKeyCodeKpEnter
	QKeyCodeKpDecimal
	QKeyCodeSysrq
	QKeyCodeKp0
	QKeyCodeKp1
	QKeyCodeKp2
	QKeyCodeKp3
	QKeyCodeKp4
	QKeyCodeKp5
	QKeyCodeKp6
	QKeyCodeKp7
	QKeyCodeKp8
	QKeyCodeKp9
	QKeyCodeLess
	QKeyCodeF11
	QKeyCodeF12
	QKeyCodePrint
	QKeyCodeHome
	QKeyCodePgup
	QKeyCodePgdn
	QKeyCodeEnd
	QKeyCodeLeft
	QKeyCodeUp
	QKeyCodeDown
	QKeyCodeRight
	QKeyCodeInsert
	QKeyCodeDelete
	QKeyCodeStop
	QKeyCodeAgain
	QKeyCodeProps
	QKeyCodeUndo
	QKeyCodeFront
	QKeyCodeCopy
	QKeyCodeOpen
	QKeyCodePaste
	QKeyCodeFind
	QKeyCodeCut
	QKeyCodeLf
	QKeyCodeHelp
	QKeyCodeMetaL
	QKeyCodeMetaR
	QKeyCodeCompose
	QKeyCodePause
	QKeyCodeRo
	QKeyCodeKpComma
	QKeyCodeKpEquals
	QKeyCodePower
)

// String implements fmt.Stringer.
func (e QKeyCode) String() string {
	switch e {
	case QKeyCodeUnmapped:
		return "unmapped"
	case QKeyCodeShift:
		return "shift"
	case QKeyCodeShiftR:
		return "shift_r"
	case QKeyCodeAlt:
		return "alt"
	case QKeyCodeAltR:
		return "alt_r"
	case QKeyCodeAltgr:
		return "altgr"
	case QKeyCodeAltgrR:
		return "altgr_r"
	case QKeyCodeCtrl:
		return "ctrl"
	case QKeyCodeCtrlR:
		return "ctrl_r"
	case QKeyCodeMenu:
		return "menu"
	case QKeyCodeEsc:
		return "esc"
	case QKeyCode1:
		return "1"
	case QKeyCode2:
		return "2"
	case QKeyCode3:
		return "3"
	case QKeyCode4:
		return "4"
	case QKeyCode5:
		return "5"
	case QKeyCode6:
		return "6"
	case QKeyCode7:
		return "7"
	case QKeyCode8:
		return "8"
	case QKeyCode9:
		return "9"
	case QKeyCode0:
		return "0"
	case QKeyCodeMinus:
		return "minus"
	case QKeyCodeEqual:
		return "equal"
	case QKeyCodeBackspace:
		return "backspace"
	case QKeyCodeTab:
		return "tab"
	case QKeyCodeQ:
		return "q"
	case QKeyCodeW:
		return "w"
	case QKeyCodeE:
		return "e"
	case QKeyCodeR:
		return "r"
	case QKeyCodeT:
		return "t"
	case QKeyCodeY:
		return "y"
	case QKeyCodeU:
		return "u"
	case QKeyCodeI:
		return "i"
	case QKeyCodeO:
		return "o"
	case QKeyCodeP:
		return "p"
	case QKeyCodeBracketLeft:
		return "bracket_left"
	case QKeyCodeBracketRight:
		return "bracket_right"
	case QKeyCodeRet:
		return "ret"
	case QKeyCodeA:
		return "a"
	case QKeyCodeS:
		return "s"
	case QKeyCodeD:
		return "d"
	case QKeyCodeF:
		return "f"
	case QKeyCodeG:
		return "g"
	case QKeyCodeH:
		return "h"
	case QKeyCodeJ:
		return "j"
	case QKeyCodeK:
		return "k"
	case QKeyCodeL:
		return "l"
	case QKeyCodeSemicolon:
		return "semicolon"
	case QKeyCodeApostrophe:
		return "apostrophe"
	case QKeyCodeGraveAccent:
		return "grave_accent"
	case QKeyCodeBackslash:
		return "backslash"
	case QKeyCodeZ:
		return "z"
	case QKeyCodeX:
		return "x"
	case QKeyCodeC:
		return "c"
	case QKeyCodeV:
		return "v"
	case QKeyCodeB:
		return "b"
	case QKeyCodeN:
		return "n"
	case QKeyCodeM:
		return "m"
	case QKeyCodeComma:
		return "comma"
	case QKeyCodeDot:
		return "dot"
	case QKeyCodeSlash:
		return "slash"
	case QKeyCodeAsterisk:
		return "asterisk"
	case QKeyCodeSpc:
		return "spc"
	case QKeyCodeCapsLock:
		return "caps_lock"
	case QKeyCodeF1:
		return "f1"
	case QKeyCodeF2:
		return "f2"
	case QKeyCodeF3:
		return "f3"
	case QKeyCodeF4:
		return "f4"
	case QKeyCodeF5:
		return "f5"
	case QKeyCodeF6:
		return "f6"
	case QKeyCodeF7:
		return "f7"
	case QKeyCodeF8:
		return "f8"
	case QKeyCodeF9:
		return "f9"
	case QKeyCodeF10:
		return "f10"
	case QKeyCodeNumLock:
		return "num_lock"
	case QKeyCodeScrollLock:
		return "scroll_lock"
	case QKeyCodeKpDivide:
		return "kp_divide"
	case QKeyCodeKpMultiply:
		return "kp_multiply"
	case QKeyCodeKpSubtract:
		return "kp_subtract"
	case QKeyCodeKpAdd:
		return "kp_add"
	case QKeyCodeKpEnter:
		return "kp_enter"
	case QKeyCodeKpDecimal:
		return "kp_decimal"
	case QKeyCodeSysrq:
		return "sysrq"
	case QKeyCodeKp0:
		return "kp_0"
	case QKeyCodeKp1:
		return "kp_1"
	case QKeyCodeKp2:
		return "kp_2"
	case QKeyCodeKp3:
		return "kp_3"
	case QKeyCodeKp4:
		return "kp_4"
	case QKeyCodeKp5:
		return "kp_5"
	case QKeyCodeKp6:
		return "kp_6"
	case QKeyCodeKp7:
		return "kp_7"
	case QKeyCodeKp8:
		return "kp_8"
	case QKeyCodeKp9:
		return "kp_9"
	case QKeyCodeLess:
		return "less"
	case QKeyCodeF11:
		return "f11"
	case QKeyCodeF12:
		return "f12"
	case QKeyCodePrint:
		return "print"
	case QKeyCodeHome:
		return "home"
	case QKeyCodePgup:
		return "pgup"
	case QKeyCodePgdn:
		return "pgdn"
	case QKeyCodeEnd:
		return "end"
	case QKeyCodeLeft:
		return "left"
	case QKeyCodeUp:
		return "up"
	case QKeyCodeDown:
		return "down"
	case QKeyCodeRight:
		return "right"
	case QKeyCodeInsert:
		return "insert"
	case QKeyCodeDelete:
		return "delete"
	case QKeyCodeStop:
		return "stop"
	case QKeyCodeAgain:
		return "again"
	case QKeyCodeProps:
		return "props"
	case QKeyCodeUndo:
		return "undo"
	case QKeyCodeFront:
		return "front"
	case QKeyCodeCopy:
		return "copy"
	case QKeyCodeOpen:
		return "open"
	case QKeyCodePaste:
		return "paste"
	case QKeyCodeFind:
		return "find"
	case QKeyCodeCut:
		return "cut"
	case QKeyCodeLf:
		return "lf"
	case QKeyCodeHelp:
		return "help"
	case QKeyCodeMetaL:
		return "meta_l"
	case QKeyCodeMetaR:
		return "meta_r"
	case QKeyCodeCompose:
		return "compose"
	case QKeyCodePause:
		return "pause"
	case QKeyCodeRo:
		return "ro"
	case QKeyCodeKpComma:
		return "kp_comma"
	case QKeyCodeKpEquals:
		return "kp_equals"
	case QKeyCodePower:
		return "power"
	default:
		return fmt.Sprintf("QKeyCode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QKeyCode) MarshalJSON() ([]byte, error) {
	switch e {
	case QKeyCodeUnmapped:
		return json.Marshal("unmapped")
	case QKeyCodeShift:
		return json.Marshal("shift")
	case QKeyCodeShiftR:
		return json.Marshal("shift_r")
	case QKeyCodeAlt:
		return json.Marshal("alt")
	case QKeyCodeAltR:
		return json.Marshal("alt_r")
	case QKeyCodeAltgr:
		return json.Marshal("altgr")
	case QKeyCodeAltgrR:
		return json.Marshal("altgr_r")
	case QKeyCodeCtrl:
		return json.Marshal("ctrl")
	case QKeyCodeCtrlR:
		return json.Marshal("ctrl_r")
	case QKeyCodeMenu:
		return json.Marshal("menu")
	case QKeyCodeEsc:
		return json.Marshal("esc")
	case QKeyCode1:
		return json.Marshal("1")
	case QKeyCode2:
		return json.Marshal("2")
	case QKeyCode3:
		return json.Marshal("3")
	case QKeyCode4:
		return json.Marshal("4")
	case QKeyCode5:
		return json.Marshal("5")
	case QKeyCode6:
		return json.Marshal("6")
	case QKeyCode7:
		return json.Marshal("7")
	case QKeyCode8:
		return json.Marshal("8")
	case QKeyCode9:
		return json.Marshal("9")
	case QKeyCode0:
		return json.Marshal("0")
	case QKeyCodeMinus:
		return json.Marshal("minus")
	case QKeyCodeEqual:
		return json.Marshal("equal")
	case QKeyCodeBackspace:
		return json.Marshal("backspace")
	case QKeyCodeTab:
		return json.Marshal("tab")
	case QKeyCodeQ:
		return json.Marshal("q")
	case QKeyCodeW:
		return json.Marshal("w")
	case QKeyCodeE:
		return json.Marshal("e")
	case QKeyCodeR:
		return json.Marshal("r")
	case QKeyCodeT:
		return json.Marshal("t")
	case QKeyCodeY:
		return json.Marshal("y")
	case QKeyCodeU:
		return json.Marshal("u")
	case QKeyCodeI:
		return json.Marshal("i")
	case QKeyCodeO:
		return json.Marshal("o")
	case QKeyCodeP:
		return json.Marshal("p")
	case QKeyCodeBracketLeft:
		return json.Marshal("bracket_left")
	case QKeyCodeBracketRight:
		return json.Marshal("bracket_right")
	case QKeyCodeRet:
		return json.Marshal("ret")
	case QKeyCodeA:
		return json.Marshal("a")
	case QKeyCodeS:
		return json.Marshal("s")
	case QKeyCodeD:
		return json.Marshal("d")
	case QKeyCodeF:
		return json.Marshal("f")
	case QKeyCodeG:
		return json.Marshal("g")
	case QKeyCodeH:
		return json.Marshal("h")
	case QKeyCodeJ:
		return json.Marshal("j")
	case QKeyCodeK:
		return json.Marshal("k")
	case QKeyCodeL:
		return json.Marshal("l")
	case QKeyCodeSemicolon:
		return json.Marshal("semicolon")
	case QKeyCodeApostrophe:
		return json.Marshal("apostrophe")
	case QKeyCodeGraveAccent:
		return json.Marshal("grave_accent")
	case QKeyCodeBackslash:
		return json.Marshal("backslash")
	case QKeyCodeZ:
		return json.Marshal("z")
	case QKeyCodeX:
		return json.Marshal("x")
	case QKeyCodeC:
		return json.Marshal("c")
	case QKeyCodeV:
		return json.Marshal("v")
	case QKeyCodeB:
		return json.Marshal("b")
	case QKeyCodeN:
		return json.Marshal("n")
	case QKeyCodeM:
		return json.Marshal("m")
	case QKeyCodeComma:
		return json.Marshal("comma")
	case QKeyCodeDot:
		return json.Marshal("dot")
	case QKeyCodeSlash:
		return json.Marshal("slash")
	case QKeyCodeAsterisk:
		return json.Marshal("asterisk")
	case QKeyCodeSpc:
		return json.Marshal("spc")
	case QKeyCodeCapsLock:
		return json.Marshal("caps_lock")
	case QKeyCodeF1:
		return json.Marshal("f1")
	case QKeyCodeF2:
		return json.Marshal("f2")
	case QKeyCodeF3:
		return json.Marshal("f3")
	case QKeyCodeF4:
		return json.Marshal("f4")
	case QKeyCodeF5:
		return json.Marshal("f5")
	case QKeyCodeF6:
		return json.Marshal("f6")
	case QKeyCodeF7:
		return json.Marshal("f7")
	case QKeyCodeF8:
		return json.Marshal("f8")
	case QKeyCodeF9:
		return json.Marshal("f9")
	case QKeyCodeF10:
		return json.Marshal("f10")
	case QKeyCodeNumLock:
		return json.Marshal("num_lock")
	case QKeyCodeScrollLock:
		return json.Marshal("scroll_lock")
	case QKeyCodeKpDivide:
		return json.Marshal("kp_divide")
	case QKeyCodeKpMultiply:
		return json.Marshal("kp_multiply")
	case QKeyCodeKpSubtract:
		return json.Marshal("kp_subtract")
	case QKeyCodeKpAdd:
		return json.Marshal("kp_add")
	case QKeyCodeKpEnter:
		return json.Marshal("kp_enter")
	case QKeyCodeKpDecimal:
		return json.Marshal("kp_decimal")
	case QKeyCodeSysrq:
		return json.Marshal("sysrq")
	case QKeyCodeKp0:
		return json.Marshal("kp_0")
	case QKeyCodeKp1:
		return json.Marshal("kp_1")
	case QKeyCodeKp2:
		return json.Marshal("kp_2")
	case QKeyCodeKp3:
		return json.Marshal("kp_3")
	case QKeyCodeKp4:
		return json.Marshal("kp_4")
	case QKeyCodeKp5:
		return json.Marshal("kp_5")
	case QKeyCodeKp6:
		return json.Marshal("kp_6")
	case QKeyCodeKp7:
		return json.Marshal("kp_7")
	case QKeyCodeKp8:
		return json.Marshal("kp_8")
	case QKeyCodeKp9:
		return json.Marshal("kp_9")
	case QKeyCodeLess:
		return json.Marshal("less")
	case QKeyCodeF11:
		return json.Marshal("f11")
	case QKeyCodeF12:
		return json.Marshal("f12")
	case QKeyCodePrint:
		return json.Marshal("print")
	case QKeyCodeHome:
		return json.Marshal("home")
	case QKeyCodePgup:
		return json.Marshal("pgup")
	case QKeyCodePgdn:
		return json.Marshal("pgdn")
	case QKeyCodeEnd:
		return json.Marshal("end")
	case QKeyCodeLeft:
		return json.Marshal("left")
	case QKeyCodeUp:
		return json.Marshal("up")
	case QKeyCodeDown:
		return json.Marshal("down")
	case QKeyCodeRight:
		return json.Marshal("right")
	case QKeyCodeInsert:
		return json.Marshal("insert")
	case QKeyCodeDelete:
		return json.Marshal("delete")
	case QKeyCodeStop:
		return json.Marshal("stop")
	case QKeyCodeAgain:
		return json.Marshal("again")
	case QKeyCodeProps:
		return json.Marshal("props")
	case QKeyCodeUndo:
		return json.Marshal("undo")
	case QKeyCodeFront:
		return json.Marshal("front")
	case QKeyCodeCopy:
		return json.Marshal("copy")
	case QKeyCodeOpen:
		return json.Marshal("open")
	case QKeyCodePaste:
		return json.Marshal("paste")
	case QKeyCodeFind:
		return json.Marshal("find")
	case QKeyCodeCut:
		return json.Marshal("cut")
	case QKeyCodeLf:
		return json.Marshal("lf")
	case QKeyCodeHelp:
		return json.Marshal("help")
	case QKeyCodeMetaL:
		return json.Marshal("meta_l")
	case QKeyCodeMetaR:
		return json.Marshal("meta_r")
	case QKeyCodeCompose:
		return json.Marshal("compose")
	case QKeyCodePause:
		return json.Marshal("pause")
	case QKeyCodeRo:
		return json.Marshal("ro")
	case QKeyCodeKpComma:
		return json.Marshal("kp_comma")
	case QKeyCodeKpEquals:
		return json.Marshal("kp_equals")
	case QKeyCodePower:
		return json.Marshal("power")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QKeyCode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QKeyCode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "unmapped":
		*e = QKeyCodeUnmapped
	case "shift":
		*e = QKeyCodeShift
	case "shift_r":
		*e = QKeyCodeShiftR
	case "alt":
		*e = QKeyCodeAlt
	case "alt_r":
		*e = QKeyCodeAltR
	case "altgr":
		*e = QKeyCodeAltgr
	case "altgr_r":
		*e = QKeyCodeAltgrR
	case "ctrl":
		*e = QKeyCodeCtrl
	case "ctrl_r":
		*e = QKeyCodeCtrlR
	case "menu":
		*e = QKeyCodeMenu
	case "esc":
		*e = QKeyCodeEsc
	case "1":
		*e = QKeyCode1
	case "2":
		*e = QKeyCode2
	case "3":
		*e = QKeyCode3
	case "4":
		*e = QKeyCode4
	case "5":
		*e = QKeyCode5
	case "6":
		*e = QKeyCode6
	case "7":
		*e = QKeyCode7
	case "8":
		*e = QKeyCode8
	case "9":
		*e = QKeyCode9
	case "0":
		*e = QKeyCode0
	case "minus":
		*e = QKeyCodeMinus
	case "equal":
		*e = QKeyCodeEqual
	case "backspace":
		*e = QKeyCodeBackspace
	case "tab":
		*e = QKeyCodeTab
	case "q":
		*e = QKeyCodeQ
	case "w":
		*e = QKeyCodeW
	case "e":
		*e = QKeyCodeE
	case "r":
		*e = QKeyCodeR
	case "t":
		*e = QKeyCodeT
	case "y":
		*e = QKeyCodeY
	case "u":
		*e = QKeyCodeU
	case "i":
		*e = QKeyCodeI
	case "o":
		*e = QKeyCodeO
	case "p":
		*e = QKeyCodeP
	case "bracket_left":
		*e = QKeyCodeBracketLeft
	case "bracket_right":
		*e = QKeyCodeBracketRight
	case "ret":
		*e = QKeyCodeRet
	case "a":
		*e = QKeyCodeA
	case "s":
		*e = QKeyCodeS
	case "d":
		*e = QKeyCodeD
	case "f":
		*e = QKeyCodeF
	case "g":
		*e = QKeyCodeG
	case "h":
		*e = QKeyCodeH
	case "j":
		*e = QKeyCodeJ
	case "k":
		*e = QKeyCodeK
	case "l":
		*e = QKeyCodeL
	case "semicolon":
		*e = QKeyCodeSemicolon
	case "apostrophe":
		*e = QKeyCodeApostrophe
	case "grave_accent":
		*e = QKeyCodeGraveAccent
	case "backslash":
		*e = QKeyCodeBackslash
	case "z":
		*e = QKeyCodeZ
	case "x":
		*e = QKeyCodeX
	case "c":
		*e = QKeyCodeC
	case "v":
		*e = QKeyCodeV
	case "b":
		*e = QKeyCodeB
	case "n":
		*e = QKeyCodeN
	case "m":
		*e = QKeyCodeM
	case "comma":
		*e = QKeyCodeComma
	case "dot":
		*e = QKeyCodeDot
	case "slash":
		*e = QKeyCodeSlash
	case "asterisk":
		*e = QKeyCodeAsterisk
	case "spc":
		*e = QKeyCodeSpc
	case "caps_lock":
		*e = QKeyCodeCapsLock
	case "f1":
		*e = QKeyCodeF1
	case "f2":
		*e = QKeyCodeF2
	case "f3":
		*e = QKeyCodeF3
	case "f4":
		*e = QKeyCodeF4
	case "f5":
		*e = QKeyCodeF5
	case "f6":
		*e = QKeyCodeF6
	case "f7":
		*e = QKeyCodeF7
	case "f8":
		*e = QKeyCodeF8
	case "f9":
		*e = QKeyCodeF9
	case "f10":
		*e = QKeyCodeF10
	case "num_lock":
		*e = QKeyCodeNumLock
	case "scroll_lock":
		*e = QKeyCodeScrollLock
	case "kp_divide":
		*e = QKeyCodeKpDivide
	case "kp_multiply":
		*e = QKeyCodeKpMultiply
	case "kp_subtract":
		*e = QKeyCodeKpSubtract
	case "kp_add":
		*e = QKeyCodeKpAdd
	case "kp_enter":
		*e = QKeyCodeKpEnter
	case "kp_decimal":
		*e = QKeyCodeKpDecimal
	case "sysrq":
		*e = QKeyCodeSysrq
	case "kp_0":
		*e = QKeyCodeKp0
	case "kp_1":
		*e = QKeyCodeKp1
	case "kp_2":
		*e = QKeyCodeKp2
	case "kp_3":
		*e = QKeyCodeKp3
	case "kp_4":
		*e = QKeyCodeKp4
	case "kp_5":
		*e = QKeyCodeKp5
	case "kp_6":
		*e = QKeyCodeKp6
	case "kp_7":
		*e = QKeyCodeKp7
	case "kp_8":
		*e = QKeyCodeKp8
	case "kp_9":
		*e = QKeyCodeKp9
	case "less":
		*e = QKeyCodeLess
	case "f11":
		*e = QKeyCodeF11
	case "f12":
		*e = QKeyCodeF12
	case "print":
		*e = QKeyCodePrint
	case "home":
		*e = QKeyCodeHome
	case "pgup":
		*e = QKeyCodePgup
	case "pgdn":
		*e = QKeyCodePgdn
	case "end":
		*e = QKeyCodeEnd
	case "left":
		*e = QKeyCodeLeft
	case "up":
		*e = QKeyCodeUp
	case "down":
		*e = QKeyCodeDown
	case "right":
		*e = QKeyCodeRight
	case "insert":
		*e = QKeyCodeInsert
	case "delete":
		*e = QKeyCodeDelete
	case "stop":
		*e = QKeyCodeStop
	case "again":
		*e = QKeyCodeAgain
	case "props":
		*e = QKeyCodeProps
	case "undo":
		*e = QKeyCodeUndo
	case "front":
		*e = QKeyCodeFront
	case "copy":
		*e = QKeyCodeCopy
	case "open":
		*e = QKeyCodeOpen
	case "paste":
		*e = QKeyCodePaste
	case "find":
		*e = QKeyCodeFind
	case "cut":
		*e = QKeyCodeCut
	case "lf":
		*e = QKeyCodeLf
	case "help":
		*e = QKeyCodeHelp
	case "meta_l":
		*e = QKeyCodeMetaL
	case "meta_r":
		*e = QKeyCodeMetaR
	case "compose":
		*e = QKeyCodeCompose
	case "pause":
		*e = QKeyCodePause
	case "ro":
		*e = QKeyCodeRo
	case "kp_comma":
		*e = QKeyCodeKpComma
	case "kp_equals":
		*e = QKeyCodeKpEquals
	case "power":
		*e = QKeyCodePower
	default:
		return fmt.Errorf("unknown enum value %q for QKeyCode", s)
	}
	return nil
}

// EVENT QUORUM_FAILURE

// EVENT QUORUM_REPORT_BAD

// Qcow2OverlapCheckFlags -> Qcow2OverlapCheckFlags (struct)

// Qcow2OverlapCheckFlags implements the "Qcow2OverlapCheckFlags" QMP API type.
type Qcow2OverlapCheckFlags struct {
	Template      *Qcow2OverlapCheckMode `json:"template,omitempty"`
	MainHeader    *bool                  `json:"main-header,omitempty"`
	ActiveL1      *bool                  `json:"active-l1,omitempty"`
	ActiveL2      *bool                  `json:"active-l2,omitempty"`
	RefcountTable *bool                  `json:"refcount-table,omitempty"`
	RefcountBlock *bool                  `json:"refcount-block,omitempty"`
	SnapshotTable *bool                  `json:"snapshot-table,omitempty"`
	InactiveL1    *bool                  `json:"inactive-l1,omitempty"`
	InactiveL2    *bool                  `json:"inactive-l2,omitempty"`
}

// Qcow2OverlapCheckMode -> Qcow2OverlapCheckMode (enum)

// Qcow2OverlapCheckMode implements the "Qcow2OverlapCheckMode" QMP API type.
type Qcow2OverlapCheckMode int

// Known values of Qcow2OverlapCheckMode.
const (
	Qcow2OverlapCheckModeNone Qcow2OverlapCheckMode = iota
	Qcow2OverlapCheckModeConstant
	Qcow2OverlapCheckModeCached
	Qcow2OverlapCheckModeAll
)

// String implements fmt.Stringer.
func (e Qcow2OverlapCheckMode) String() string {
	switch e {
	case Qcow2OverlapCheckModeNone:
		return "none"
	case Qcow2OverlapCheckModeConstant:
		return "constant"
	case Qcow2OverlapCheckModeCached:
		return "cached"
	case Qcow2OverlapCheckModeAll:
		return "all"
	default:
		return fmt.Sprintf("Qcow2OverlapCheckMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e Qcow2OverlapCheckMode) MarshalJSON() ([]byte, error) {
	switch e {
	case Qcow2OverlapCheckModeNone:
		return json.Marshal("none")
	case Qcow2OverlapCheckModeConstant:
		return json.Marshal("constant")
	case Qcow2OverlapCheckModeCached:
		return json.Marshal("cached")
	case Qcow2OverlapCheckModeAll:
		return json.Marshal("all")
	default:
		return nil, fmt.Errorf("unknown enum value %q for Qcow2OverlapCheckMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Qcow2OverlapCheckMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = Qcow2OverlapCheckModeNone
	case "constant":
		*e = Qcow2OverlapCheckModeConstant
	case "cached":
		*e = Qcow2OverlapCheckModeCached
	case "all":
		*e = Qcow2OverlapCheckModeAll
	default:
		return fmt.Errorf("unknown enum value %q for Qcow2OverlapCheckMode", s)
	}
	return nil
}

// Qcow2OverlapChecks -> Qcow2OverlapChecks (alternate)

// Qcow2OverlapChecks implements the "Qcow2OverlapChecks" QMP API type.
//
// Can be one of:
//   - Qcow2OverlapChecksFlags
//   - Qcow2OverlapChecksMode
type Qcow2OverlapChecks interface {
	isQcow2OverlapChecks()
}

// Qcow2OverlapChecksFlags is an implementation of Qcow2OverlapChecks
type Qcow2OverlapChecksFlags Qcow2OverlapCheckFlags

func (Qcow2OverlapChecksFlags) isQcow2OverlapChecks() {}

// Qcow2OverlapChecksMode is an implementation of Qcow2OverlapChecks
type Qcow2OverlapChecksMode Qcow2OverlapCheckMode

func (Qcow2OverlapChecksMode) isQcow2OverlapChecks() {}

func decodeQcow2OverlapChecks(bs json.RawMessage) (Qcow2OverlapChecks, error) {

	var flags Qcow2OverlapChecksFlags
	if err := json.Unmarshal([]byte(bs), &flags); err == nil {
		return flags, nil
	}

	var mode Qcow2OverlapChecksMode
	if err := json.Unmarshal([]byte(bs), &mode); err == nil {
		return mode, nil
	}
	return nil, fmt.Errorf("unable to decode %q as a Qcow2OverlapChecks", string(bs))
}

// QuorumOpType -> QuorumOpType (enum)

// QuorumOpType implements the "QuorumOpType" QMP API type.
type QuorumOpType int

// Known values of QuorumOpType.
const (
	QuorumOpTypeRead QuorumOpType = iota
	QuorumOpTypeWrite
	QuorumOpTypeFlush
)

// String implements fmt.Stringer.
func (e QuorumOpType) String() string {
	switch e {
	case QuorumOpTypeRead:
		return "read"
	case QuorumOpTypeWrite:
		return "write"
	case QuorumOpTypeFlush:
		return "flush"
	default:
		return fmt.Sprintf("QuorumOpType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QuorumOpType) MarshalJSON() ([]byte, error) {
	switch e {
	case QuorumOpTypeRead:
		return json.Marshal("read")
	case QuorumOpTypeWrite:
		return json.Marshal("write")
	case QuorumOpTypeFlush:
		return json.Marshal("flush")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QuorumOpType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QuorumOpType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "read":
		*e = QuorumOpTypeRead
	case "write":
		*e = QuorumOpTypeWrite
	case "flush":
		*e = QuorumOpTypeFlush
	default:
		return fmt.Errorf("unknown enum value %q for QuorumOpType", s)
	}
	return nil
}

// QuorumReadPattern -> QuorumReadPattern (enum)

// QuorumReadPattern implements the "QuorumReadPattern" QMP API type.
type QuorumReadPattern int

// Known values of QuorumReadPattern.
const (
	QuorumReadPatternQuorum QuorumReadPattern = iota
	QuorumReadPatternFifo
)

// String implements fmt.Stringer.
func (e QuorumReadPattern) String() string {
	switch e {
	case QuorumReadPatternQuorum:
		return "quorum"
	case QuorumReadPatternFifo:
		return "fifo"
	default:
		return fmt.Sprintf("QuorumReadPattern(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QuorumReadPattern) MarshalJSON() ([]byte, error) {
	switch e {
	case QuorumReadPatternQuorum:
		return json.Marshal("quorum")
	case QuorumReadPatternFifo:
		return json.Marshal("fifo")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QuorumReadPattern", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QuorumReadPattern) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "quorum":
		*e = QuorumReadPatternQuorum
	case "fifo":
		*e = QuorumReadPatternFifo
	default:
		return fmt.Errorf("unknown enum value %q for QuorumReadPattern", s)
	}
	return nil
}

// EVENT RESET

// EVENT RESUME

// EVENT RTC_CHANGE

// ReplicationMode -> ReplicationMode (enum)

// ReplicationMode implements the "ReplicationMode" QMP API type.
type ReplicationMode int

// Known values of ReplicationMode.
const (
	ReplicationModePrimary ReplicationMode = iota
	ReplicationModeSecondary
)

// String implements fmt.Stringer.
func (e ReplicationMode) String() string {
	switch e {
	case ReplicationModePrimary:
		return "primary"
	case ReplicationModeSecondary:
		return "secondary"
	default:
		return fmt.Sprintf("ReplicationMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ReplicationMode) MarshalJSON() ([]byte, error) {
	switch e {
	case ReplicationModePrimary:
		return json.Marshal("primary")
	case ReplicationModeSecondary:
		return json.Marshal("secondary")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ReplicationMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ReplicationMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "primary":
		*e = ReplicationModePrimary
	case "secondary":
		*e = ReplicationModeSecondary
	default:
		return fmt.Errorf("unknown enum value %q for ReplicationMode", s)
	}
	return nil
}

// RockerOfDpaFlow -> RockerOfDpaFlow (struct)

// RockerOfDpaFlow implements the "RockerOfDpaFlow" QMP API type.
type RockerOfDpaFlow struct {
	Cookie uint64                `json:"cookie"`
	Hits   uint64                `json:"hits"`
	Key    RockerOfDpaFlowKey    `json:"key"`
	Mask   RockerOfDpaFlowMask   `json:"mask"`
	Action RockerOfDpaFlowAction `json:"action"`
}

// RockerOfDpaFlowAction -> RockerOfDpaFlowAction (struct)

// RockerOfDpaFlowAction implements the "RockerOfDpaFlowAction" QMP API type.
type RockerOfDpaFlowAction struct {
	GotoTbl     *uint32 `json:"goto-tbl,omitempty"`
	GroupID     *uint32 `json:"group-id,omitempty"`
	TunnelLport *uint32 `json:"tunnel-lport,omitempty"`
	VlanID      *uint16 `json:"vlan-id,omitempty"`
	NewVlanID   *uint16 `json:"new-vlan-id,omitempty"`
	OutPport    *uint32 `json:"out-pport,omitempty"`
}

// RockerOfDpaFlowKey -> RockerOfDpaFlowKey (struct)

// RockerOfDpaFlowKey implements the "RockerOfDpaFlowKey" QMP API type.
type RockerOfDpaFlowKey struct {
	Priority uint32  `json:"priority"`
	TblID    uint32  `json:"tbl-id"`
	InPport  *uint32 `json:"in-pport,omitempty"`
	TunnelID *uint32 `json:"tunnel-id,omitempty"`
	VlanID   *uint16 `json:"vlan-id,omitempty"`
	EthType  *uint16 `json:"eth-type,omitempty"`
	EthSrc   *string `json:"eth-src,omitempty"`
	EthDst   *string `json:"eth-dst,omitempty"`
	IPProto  *uint8  `json:"ip-proto,omitempty"`
	IPTos    *uint8  `json:"ip-tos,omitempty"`
	IPDst    *string `json:"ip-dst,omitempty"`
}

// RockerOfDpaFlowMask -> RockerOfDpaFlowMask (struct)

// RockerOfDpaFlowMask implements the "RockerOfDpaFlowMask" QMP API type.
type RockerOfDpaFlowMask struct {
	InPport  *uint32 `json:"in-pport,omitempty"`
	TunnelID *uint32 `json:"tunnel-id,omitempty"`
	VlanID   *uint16 `json:"vlan-id,omitempty"`
	EthSrc   *string `json:"eth-src,omitempty"`
	EthDst   *string `json:"eth-dst,omitempty"`
	IPProto  *uint8  `json:"ip-proto,omitempty"`
	IPTos    *uint8  `json:"ip-tos,omitempty"`
}

// RockerOfDpaGroup -> RockerOfDpaGroup (struct)

// RockerOfDpaGroup implements the "RockerOfDpaGroup" QMP API type.
type RockerOfDpaGroup struct {
	ID        uint32   `json:"id"`
	Type      uint8    `json:"type"`
	VlanID    *uint16  `json:"vlan-id,omitempty"`
	Pport     *uint32  `json:"pport,omitempty"`
	Index     *uint32  `json:"index,omitempty"`
	OutPport  *uint32  `json:"out-pport,omitempty"`
	GroupID   *uint32  `json:"group-id,omitempty"`
	SetVlanID *uint16  `json:"set-vlan-id,omitempty"`
	PopVlan   *uint8   `json:"pop-vlan,omitempty"`
	GroupIds  []uint32 `json:"group-ids,omitempty"`
	SetEthSrc *string  `json:"set-eth-src,omitempty"`
	SetEthDst *string  `json:"set-eth-dst,omitempty"`
	TTLCheck  *uint8   `json:"ttl-check,omitempty"`
}

// RockerPort -> RockerPort (struct)

// RockerPort implements the "RockerPort" QMP API type.
type RockerPort struct {
	Name    string            `json:"name"`
	Enabled bool              `json:"enabled"`
	LinkUp  bool              `json:"link-up"`
	Speed   uint32            `json:"speed"`
	Duplex  RockerPortDuplex  `json:"duplex"`
	Autoneg RockerPortAutoneg `json:"autoneg"`
}

// RockerPortAutoneg -> RockerPortAutoneg (enum)

// RockerPortAutoneg implements the "RockerPortAutoneg" QMP API type.
type RockerPortAutoneg int

// Known values of RockerPortAutoneg.
const (
	RockerPortAutonegOff RockerPortAutoneg = iota
	RockerPortAutonegOn
)

// String implements fmt.Stringer.
func (e RockerPortAutoneg) String() string {
	switch e {
	case RockerPortAutonegOff:
		return "off"
	case RockerPortAutonegOn:
		return "on"
	default:
		return fmt.Sprintf("RockerPortAutoneg(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e RockerPortAutoneg) MarshalJSON() ([]byte, error) {
	switch e {
	case RockerPortAutonegOff:
		return json.Marshal("off")
	case RockerPortAutonegOn:
		return json.Marshal("on")
	default:
		return nil, fmt.Errorf("unknown enum value %q for RockerPortAutoneg", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *RockerPortAutoneg) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "off":
		*e = RockerPortAutonegOff
	case "on":
		*e = RockerPortAutonegOn
	default:
		return fmt.Errorf("unknown enum value %q for RockerPortAutoneg", s)
	}
	return nil
}

// RockerPortDuplex -> RockerPortDuplex (enum)

// RockerPortDuplex implements the "RockerPortDuplex" QMP API type.
type RockerPortDuplex int

// Known values of RockerPortDuplex.
const (
	RockerPortDuplexHalf RockerPortDuplex = iota
	RockerPortDuplexFull
)

// String implements fmt.Stringer.
func (e RockerPortDuplex) String() string {
	switch e {
	case RockerPortDuplexHalf:
		return "half"
	case RockerPortDuplexFull:
		return "full"
	default:
		return fmt.Sprintf("RockerPortDuplex(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e RockerPortDuplex) MarshalJSON() ([]byte, error) {
	switch e {
	case RockerPortDuplexHalf:
		return json.Marshal("half")
	case RockerPortDuplexFull:
		return json.Marshal("full")
	default:
		return nil, fmt.Errorf("unknown enum value %q for RockerPortDuplex", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *RockerPortDuplex) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "half":
		*e = RockerPortDuplexHalf
	case "full":
		*e = RockerPortDuplexFull
	default:
		return fmt.Errorf("unknown enum value %q for RockerPortDuplex", s)
	}
	return nil
}

// RockerSwitch -> RockerSwitch (struct)

// RockerSwitch implements the "RockerSwitch" QMP API type.
type RockerSwitch struct {
	Name  string `json:"name"`
	ID    uint64 `json:"id"`
	Ports uint32 `json:"ports"`
}

// RunState -> RunState (enum)

// RunState implements the "RunState" QMP API type.
type RunState int

// Known values of RunState.
const (
	RunStateDebug RunState = iota
	RunStateInmigrate
	RunStateInternalError
	RunStateIOError
	RunStatePaused
	RunStatePostmigrate
	RunStatePrelaunch
	RunStateFinishMigrate
	RunStateRestoreVM
	RunStateRunning
	RunStateSaveVM
	RunStateShutdown
	RunStateSuspended
	RunStateWatchdog
	RunStateGuestPanicked
)

// String implements fmt.Stringer.
func (e RunState) String() string {
	switch e {
	case RunStateDebug:
		return "debug"
	case RunStateInmigrate:
		return "inmigrate"
	case RunStateInternalError:
		return "internal-error"
	case RunStateIOError:
		return "io-error"
	case RunStatePaused:
		return "paused"
	case RunStatePostmigrate:
		return "postmigrate"
	case RunStatePrelaunch:
		return "prelaunch"
	case RunStateFinishMigrate:
		return "finish-migrate"
	case RunStateRestoreVM:
		return "restore-vm"
	case RunStateRunning:
		return "running"
	case RunStateSaveVM:
		return "save-vm"
	case RunStateShutdown:
		return "shutdown"
	case RunStateSuspended:
		return "suspended"
	case RunStateWatchdog:
		return "watchdog"
	case RunStateGuestPanicked:
		return "guest-panicked"
	default:
		return fmt.Sprintf("RunState(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e RunState) MarshalJSON() ([]byte, error) {
	switch e {
	case RunStateDebug:
		return json.Marshal("debug")
	case RunStateInmigrate:
		return json.Marshal("inmigrate")
	case RunStateInternalError:
		return json.Marshal("internal-error")
	case RunStateIOError:
		return json.Marshal("io-error")
	case RunStatePaused:
		return json.Marshal("paused")
	case RunStatePostmigrate:
		return json.Marshal("postmigrate")
	case RunStatePrelaunch:
		return json.Marshal("prelaunch")
	case RunStateFinishMigrate:
		return json.Marshal("finish-migrate")
	case RunStateRestoreVM:
		return json.Marshal("restore-vm")
	case RunStateRunning:
		return json.Marshal("running")
	case RunStateSaveVM:
		return json.Marshal("save-vm")
	case RunStateShutdown:
		return json.Marshal("shutdown")
	case RunStateSuspended:
		return json.Marshal("suspended")
	case RunStateWatchdog:
		return json.Marshal("watchdog")
	case RunStateGuestPanicked:
		return json.Marshal("guest-panicked")
	default:
		return nil, fmt.Errorf("unknown enum value %q for RunState", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *RunState) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "debug":
		*e = RunStateDebug
	case "inmigrate":
		*e = RunStateInmigrate
	case "internal-error":
		*e = RunStateInternalError
	case "io-error":
		*e = RunStateIOError
	case "paused":
		*e = RunStatePaused
	case "postmigrate":
		*e = RunStatePostmigrate
	case "prelaunch":
		*e = RunStatePrelaunch
	case "finish-migrate":
		*e = RunStateFinishMigrate
	case "restore-vm":
		*e = RunStateRestoreVM
	case "running":
		*e = RunStateRunning
	case "save-vm":
		*e = RunStateSaveVM
	case "shutdown":
		*e = RunStateShutdown
	case "suspended":
		*e = RunStateSuspended
	case "watchdog":
		*e = RunStateWatchdog
	case "guest-panicked":
		*e = RunStateGuestPanicked
	default:
		return fmt.Errorf("unknown enum value %q for RunState", s)
	}
	return nil
}

// RxFilterInfo -> RxFilterInfo (struct)

// RxFilterInfo implements the "RxFilterInfo" QMP API type.
type RxFilterInfo struct {
	Name              string   `json:"name"`
	Promiscuous       bool     `json:"promiscuous"`
	Multicast         RxState  `json:"multicast"`
	Unicast           RxState  `json:"unicast"`
	Vlan              RxState  `json:"vlan"`
	BroadcastAllowed  bool     `json:"broadcast-allowed"`
	MulticastOverflow bool     `json:"multicast-overflow"`
	UnicastOverflow   bool     `json:"unicast-overflow"`
	MainMac           string   `json:"main-mac"`
	VlanTable         []int64  `json:"vlan-table"`
	UnicastTable      []string `json:"unicast-table"`
	MulticastTable    []string `json:"multicast-table"`
}

// RxState -> RxState (enum)

// RxState implements the "RxState" QMP API type.
type RxState int

// Known values of RxState.
const (
	RxStateNormal RxState = iota
	RxStateNone
	RxStateAll
)

// String implements fmt.Stringer.
func (e RxState) String() string {
	switch e {
	case RxStateNormal:
		return "normal"
	case RxStateNone:
		return "none"
	case RxStateAll:
		return "all"
	default:
		return fmt.Sprintf("RxState(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e RxState) MarshalJSON() ([]byte, error) {
	switch e {
	case RxStateNormal:
		return json.Marshal("normal")
	case RxStateNone:
		return json.Marshal("none")
	case RxStateAll:
		return json.Marshal("all")
	default:
		return nil, fmt.Errorf("unknown enum value %q for RxState", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *RxState) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "normal":
		*e = RxStateNormal
	case "none":
		*e = RxStateNone
	case "all":
		*e = RxStateAll
	default:
		return fmt.Errorf("unknown enum value %q for RxState", s)
	}
	return nil
}

// EVENT SHUTDOWN

// EVENT SPICE_CONNECTED

// EVENT SPICE_DISCONNECTED

// EVENT SPICE_INITIALIZED

// EVENT SPICE_MIGRATE_COMPLETED

// EVENT STOP

// EVENT SUSPEND

// EVENT SUSPEND_DISK

// SchemaInfo -> SchemaInfo (flat union)

// SchemaInfo implements the "SchemaInfo" QMP API type.
//
// Can be one of:
//   - SchemaInfoAlternate
//   - SchemaInfoArray
//   - SchemaInfoBuiltin
//   - SchemaInfoCommand
//   - SchemaInfoEnum
//   - SchemaInfoEvent
//   - SchemaInfoObject
type SchemaInfo interface {
	isSchemaInfo()
}

// SchemaInfoAlternate is an implementation of SchemaInfo.
type SchemaInfoAlternate struct {
	Name    string                      `json:"name"`
	Members []SchemaInfoAlternateMember `json:"members"`
}

func (SchemaInfoAlternate) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoAlternate) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoAlternate
	}{
		SchemaMetaTypeAlternate,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoArray is an implementation of SchemaInfo.
type SchemaInfoArray struct {
	Name        string `json:"name"`
	ElementType string `json:"element-type"`
}

func (SchemaInfoArray) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoArray) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoArray
	}{
		SchemaMetaTypeArray,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoBuiltin is an implementation of SchemaInfo.
type SchemaInfoBuiltin struct {
	Name     string   `json:"name"`
	JSONType JSONType `json:"json-type"`
}

func (SchemaInfoBuiltin) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoBuiltin) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoBuiltin
	}{
		SchemaMetaTypeBuiltin,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoCommand is an implementation of SchemaInfo.
type SchemaInfoCommand struct {
	Name    string `json:"name"`
	ArgType string `json:"arg-type"`
	RetType string `json:"ret-type"`
}

func (SchemaInfoCommand) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoCommand) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoCommand
	}{
		SchemaMetaTypeCommand,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoEnum is an implementation of SchemaInfo.
type SchemaInfoEnum struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

func (SchemaInfoEnum) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoEnum) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoEnum
	}{
		SchemaMetaTypeEnum,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoEvent is an implementation of SchemaInfo.
type SchemaInfoEvent struct {
	Name    string `json:"name"`
	ArgType string `json:"arg-type"`
}

func (SchemaInfoEvent) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoEvent) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoEvent
	}{
		SchemaMetaTypeEvent,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoObject is an implementation of SchemaInfo.
type SchemaInfoObject struct {
	Name     string                    `json:"name"`
	Members  []SchemaInfoObjectMember  `json:"members"`
	Tag      *string                   `json:"tag,omitempty"`
	Variants []SchemaInfoObjectVariant `json:"variants,omitempty"`
}

func (SchemaInfoObject) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoObject) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoObject
	}{
		SchemaMetaTypeObject,
		s,
	}
	return json.Marshal(v)
}

func decodeSchemaInfo(bs json.RawMessage) (SchemaInfo, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.MetaType {
	case SchemaMetaTypeAlternate:
		var ret SchemaInfoAlternate
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeArray:
		var ret SchemaInfoArray
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeBuiltin:
		var ret SchemaInfoBuiltin
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeCommand:
		var ret SchemaInfoCommand
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeEnum:
		var ret SchemaInfoEnum
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeEvent:
		var ret SchemaInfoEvent
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeObject:
		var ret SchemaInfoObject
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union SchemaInfo", v.MetaType)
	}
}

// SchemaInfoAlternateMember -> SchemaInfoAlternateMember (struct)

// SchemaInfoAlternateMember implements the "SchemaInfoAlternateMember" QMP API type.
type SchemaInfoAlternateMember struct {
	Type string `json:"type"`
}

// SchemaInfoObjectMember -> SchemaInfoObjectMember (struct)

// SchemaInfoObjectMember implements the "SchemaInfoObjectMember" QMP API type.
type SchemaInfoObjectMember struct {
	Name    string       `json:"name"`
	Type    string       `json:"type"`
	Default *interface{} `json:"default,omitempty"`
}

// SchemaInfoObjectVariant -> SchemaInfoObjectVariant (struct)

// SchemaInfoObjectVariant implements the "SchemaInfoObjectVariant" QMP API type.
type SchemaInfoObjectVariant struct {
	Case string `json:"case"`
	Type string `json:"type"`
}

// SchemaMetaType -> SchemaMetaType (enum)

// SchemaMetaType implements the "SchemaMetaType" QMP API type.
type SchemaMetaType int

// Known values of SchemaMetaType.
const (
	SchemaMetaTypeBuiltin SchemaMetaType = iota
	SchemaMetaTypeEnum
	SchemaMetaTypeArray
	SchemaMetaTypeObject
	SchemaMetaTypeAlternate
	SchemaMetaTypeCommand
	SchemaMetaTypeEvent
)

// String implements fmt.Stringer.
func (e SchemaMetaType) String() string {
	switch e {
	case SchemaMetaTypeBuiltin:
		return "builtin"
	case SchemaMetaTypeEnum:
		return "enum"
	case SchemaMetaTypeArray:
		return "array"
	case SchemaMetaTypeObject:
		return "object"
	case SchemaMetaTypeAlternate:
		return "alternate"
	case SchemaMetaTypeCommand:
		return "command"
	case SchemaMetaTypeEvent:
		return "event"
	default:
		return fmt.Sprintf("SchemaMetaType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e SchemaMetaType) MarshalJSON() ([]byte, error) {
	switch e {
	case SchemaMetaTypeBuiltin:
		return json.Marshal("builtin")
	case SchemaMetaTypeEnum:
		return json.Marshal("enum")
	case SchemaMetaTypeArray:
		return json.Marshal("array")
	case SchemaMetaTypeObject:
		return json.Marshal("object")
	case SchemaMetaTypeAlternate:
		return json.Marshal("alternate")
	case SchemaMetaTypeCommand:
		return json.Marshal("command")
	case SchemaMetaTypeEvent:
		return json.Marshal("event")
	default:
		return nil, fmt.Errorf("unknown enum value %q for SchemaMetaType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *SchemaMetaType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "builtin":
		*e = SchemaMetaTypeBuiltin
	case "enum":
		*e = SchemaMetaTypeEnum
	case "array":
		*e = SchemaMetaTypeArray
	case "object":
		*e = SchemaMetaTypeObject
	case "alternate":
		*e = SchemaMetaTypeAlternate
	case "command":
		*e = SchemaMetaTypeCommand
	case "event":
		*e = SchemaMetaTypeEvent
	default:
		return fmt.Errorf("unknown enum value %q for SchemaMetaType", s)
	}
	return nil
}

// SnapshotInfo -> SnapshotInfo (struct)

// SnapshotInfo implements the "SnapshotInfo" QMP API type.
type SnapshotInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	VMStateSize int64  `json:"vm-state-size"`
	DateSec     int64  `json:"date-sec"`
	DateNsec    int64  `json:"date-nsec"`
	VMClockSec  int64  `json:"vm-clock-sec"`
	VMClockNsec int64  `json:"vm-clock-nsec"`
}

// SocketAddress -> SocketAddress (simple union)

// SocketAddress implements the "SocketAddress" QMP API type.
//
// Can be one of:
//   - SocketAddressFD
//   - SocketAddressInet
//   - SocketAddressUnix
type SocketAddress interface {
	isSocketAddress()
}

// SocketAddressFD is an implementation of SocketAddress.
type SocketAddressFD String

func (SocketAddressFD) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressFD) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "fd",
		"data": s,
	}
	return json.Marshal(v)
}

// SocketAddressInet is an implementation of SocketAddress.
type SocketAddressInet InetSocketAddress

func (SocketAddressInet) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressInet) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "inet",
		"data": s,
	}
	return json.Marshal(v)
}

// SocketAddressUnix is an implementation of SocketAddress.
type SocketAddressUnix UnixSocketAddress

func (SocketAddressUnix) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressUnix) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "unix",
		"data": s,
	}
	return json.Marshal(v)
}

func decodeSocketAddress(bs json.RawMessage) (SocketAddress, error) {
	v := struct {
		T string          `json:"type"`
		V json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.T {
	case "fd":
		var ret SocketAddressFD
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "inet":
		var ret SocketAddressInet
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "unix":
		var ret SocketAddressUnix
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("unknown subtype %q for union SocketAddress", v.T)
	}
}

// SpiceBasicInfo -> SpiceBasicInfo (struct)

// SpiceBasicInfo implements the "SpiceBasicInfo" QMP API type.
type SpiceBasicInfo struct {
	Host   string               `json:"host"`
	Port   string               `json:"port"`
	Family NetworkAddressFamily `json:"family"`
}

// SpiceChannel -> SpiceChannel (struct)

// SpiceChannel implements the "SpiceChannel" QMP API type.
type SpiceChannel struct {
	Host         string               `json:"host"`
	Port         string               `json:"port"`
	Family       NetworkAddressFamily `json:"family"`
	ConnectionID int64                `json:"connection-id"`
	ChannelType  int64                `json:"channel-type"`
	ChannelID    int64                `json:"channel-id"`
	TLS          bool                 `json:"tls"`
}

// SpiceInfo -> SpiceInfo (struct)

// SpiceInfo implements the "SpiceInfo" QMP API type.
type SpiceInfo struct {
	Enabled         bool                `json:"enabled"`
	Migrated        bool                `json:"migrated"`
	Host            *string             `json:"host,omitempty"`
	Port            *int64              `json:"port,omitempty"`
	TLSPort         *int64              `json:"tls-port,omitempty"`
	Auth            *string             `json:"auth,omitempty"`
	CompiledVersion *string             `json:"compiled-version,omitempty"`
	MouseMode       SpiceQueryMouseMode `json:"mouse-mode"`
	Channels        []SpiceChannel      `json:"channels,omitempty"`
}

// SpiceQueryMouseMode -> SpiceQueryMouseMode (enum)

// SpiceQueryMouseMode implements the "SpiceQueryMouseMode" QMP API type.
type SpiceQueryMouseMode int

// Known values of SpiceQueryMouseMode.
const (
	SpiceQueryMouseModeClient SpiceQueryMouseMode = iota
	SpiceQueryMouseModeServer
	SpiceQueryMouseModeUnknown
)

// String implements fmt.Stringer.
func (e SpiceQueryMouseMode) String() string {
	switch e {
	case SpiceQueryMouseModeClient:
		return "client"
	case SpiceQueryMouseModeServer:
		return "server"
	case SpiceQueryMouseModeUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("SpiceQueryMouseMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e SpiceQueryMouseMode) MarshalJSON() ([]byte, error) {
	switch e {
	case SpiceQueryMouseModeClient:
		return json.Marshal("client")
	case SpiceQueryMouseModeServer:
		return json.Marshal("server")
	case SpiceQueryMouseModeUnknown:
		return json.Marshal("unknown")
	default:
		return nil, fmt.Errorf("unknown enum value %q for SpiceQueryMouseMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *SpiceQueryMouseMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "client":
		*e = SpiceQueryMouseModeClient
	case "server":
		*e = SpiceQueryMouseModeServer
	case "unknown":
		*e = SpiceQueryMouseModeUnknown
	default:
		return fmt.Errorf("unknown enum value %q for SpiceQueryMouseMode", s)
	}
	return nil
}

// SpiceServerInfo -> SpiceServerInfo (struct)

// SpiceServerInfo implements the "SpiceServerInfo" QMP API type.
type SpiceServerInfo struct {
	Host   string               `json:"host"`
	Port   string               `json:"port"`
	Family NetworkAddressFamily `json:"family"`
	Auth   *string              `json:"auth,omitempty"`
}

// StatusInfo -> StatusInfo (struct)

// StatusInfo implements the "StatusInfo" QMP API type.
type StatusInfo struct {
	Running    bool     `json:"running"`
	Singlestep bool     `json:"singlestep"`
	Status     RunState `json:"status"`
}

// String -> String (struct)

// String implements the "String" QMP API type.
type String struct {
	Str string `json:"str"`
}

// TPMInfo -> TPMInfo (struct)

// TPMInfo implements the "TPMInfo" QMP API type.
type TPMInfo struct {
	ID      string         `json:"id"`
	Model   TPMModel       `json:"model"`
	Options TPMTypeOptions `json:"options"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *TPMInfo) UnmarshalJSON(bs []byte) error {
	v := struct {
		ID      string          `json:"id"`
		Model   TPMModel        `json:"model"`
		Options json.RawMessage `json:"options"`
	}{}
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return err
	}
	s.ID = v.ID
	s.Model = v.Model
	s.Options, err = decodeTPMTypeOptions(v.Options)
	if err != nil {
		return err
	}

	return nil
}

// TPMPassthroughOptions -> TPMPassthroughOptions (struct)

// TPMPassthroughOptions implements the "TPMPassthroughOptions" QMP API type.
type TPMPassthroughOptions struct {
	Path       *string `json:"path,omitempty"`
	CancelPath *string `json:"cancel-path,omitempty"`
}

// TargetInfo -> TargetInfo (struct)

// TargetInfo implements the "TargetInfo" QMP API type.
type TargetInfo struct {
	Arch string `json:"arch"`
}

// TpmModel -> TPMModel (enum)

// TPMModel implements the "TpmModel" QMP API type.
type TPMModel int

// Known values of TPMModel.
const (
	TPMModelTPMTis TPMModel = iota
)

// String implements fmt.Stringer.
func (e TPMModel) String() string {
	switch e {
	case TPMModelTPMTis:
		return "tpm-tis"
	default:
		return fmt.Sprintf("TPMModel(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e TPMModel) MarshalJSON() ([]byte, error) {
	switch e {
	case TPMModelTPMTis:
		return json.Marshal("tpm-tis")
	default:
		return nil, fmt.Errorf("unknown enum value %q for TPMModel", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *TPMModel) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "tpm-tis":
		*e = TPMModelTPMTis
	default:
		return fmt.Errorf("unknown enum value %q for TPMModel", s)
	}
	return nil
}

// TpmType -> TPMType (enum)

// TPMType implements the "TpmType" QMP API type.
type TPMType int

// Known values of TPMType.
const (
	TPMTypePassthrough TPMType = iota
)

// String implements fmt.Stringer.
func (e TPMType) String() string {
	switch e {
	case TPMTypePassthrough:
		return "passthrough"
	default:
		return fmt.Sprintf("TPMType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e TPMType) MarshalJSON() ([]byte, error) {
	switch e {
	case TPMTypePassthrough:
		return json.Marshal("passthrough")
	default:
		return nil, fmt.Errorf("unknown enum value %q for TPMType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *TPMType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "passthrough":
		*e = TPMTypePassthrough
	default:
		return fmt.Errorf("unknown enum value %q for TPMType", s)
	}
	return nil
}

// TpmTypeOptions -> TPMTypeOptions (simple union)

// TPMTypeOptions implements the "TpmTypeOptions" QMP API type.
//
// Can be one of:
//   - TPMTypeOptionsPassthrough
type TPMTypeOptions interface {
	isTPMTypeOptions()
}

// TPMTypeOptionsPassthrough is an implementation of TPMTypeOptions.
type TPMTypeOptionsPassthrough TPMPassthroughOptions

func (TPMTypeOptionsPassthrough) isTPMTypeOptions() {}

// MarshalJSON implements json.Marshaler.
func (s TPMTypeOptionsPassthrough) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "passthrough",
		"data": s,
	}
	return json.Marshal(v)
}

func decodeTPMTypeOptions(bs json.RawMessage) (TPMTypeOptions, error) {
	v := struct {
		T string          `json:"type"`
		V json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.T {
	case "passthrough":
		var ret TPMTypeOptionsPassthrough
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("unknown subtype %q for union TPMTypeOptions", v.T)
	}
}

// TraceEventInfo -> TraceEventInfo (struct)

// TraceEventInfo implements the "TraceEventInfo" QMP API type.
type TraceEventInfo struct {
	Name  string          `json:"name"`
	State TraceEventState `json:"state"`
	Vcpu  bool            `json:"vcpu"`
}

// TraceEventState -> TraceEventState (enum)

// TraceEventState implements the "TraceEventState" QMP API type.
type TraceEventState int

// Known values of TraceEventState.
const (
	TraceEventStateUnavailable TraceEventState = iota
	TraceEventStateDisabled
	TraceEventStateEnabled
)

// String implements fmt.Stringer.
func (e TraceEventState) String() string {
	switch e {
	case TraceEventStateUnavailable:
		return "unavailable"
	case TraceEventStateDisabled:
		return "disabled"
	case TraceEventStateEnabled:
		return "enabled"
	default:
		return fmt.Sprintf("TraceEventState(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e TraceEventState) MarshalJSON() ([]byte, error) {
	switch e {
	case TraceEventStateUnavailable:
		return json.Marshal("unavailable")
	case TraceEventStateDisabled:
		return json.Marshal("disabled")
	case TraceEventStateEnabled:
		return json.Marshal("enabled")
	default:
		return nil, fmt.Errorf("unknown enum value %q for TraceEventState", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *TraceEventState) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "unavailable":
		*e = TraceEventStateUnavailable
	case "disabled":
		*e = TraceEventStateDisabled
	case "enabled":
		*e = TraceEventStateEnabled
	default:
		return fmt.Errorf("unknown enum value %q for TraceEventState", s)
	}
	return nil
}

// TransactionAction -> TransactionAction (simple union)

// TransactionAction implements the "TransactionAction" QMP API type.
//
// Can be one of:
//   - TransactionActionAbort
//   - TransactionActionBlockDirtyBitmapAdd
//   - TransactionActionBlockDirtyBitmapClear
//   - TransactionActionBlockdevBackup
//   - TransactionActionBlockdevSnapshot
//   - TransactionActionBlockdevSnapshotInternalSync
//   - TransactionActionBlockdevSnapshotSync
//   - TransactionActionDriveBackup
type TransactionAction interface {
	isTransactionAction()
}

// TransactionActionAbort is an implementation of TransactionAction.
type TransactionActionAbort Abort

func (TransactionActionAbort) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionAbort) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "abort",
		"data": s,
	}
	return json.Marshal(v)
}

// TransactionActionBlockDirtyBitmapAdd is an implementation of TransactionAction.
type TransactionActionBlockDirtyBitmapAdd BlockDirtyBitmapAdd

func (TransactionActionBlockDirtyBitmapAdd) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionBlockDirtyBitmapAdd) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "block-dirty-bitmap-add",
		"data": s,
	}
	return json.Marshal(v)
}

// TransactionActionBlockDirtyBitmapClear is an implementation of TransactionAction.
type TransactionActionBlockDirtyBitmapClear BlockDirtyBitmap

func (TransactionActionBlockDirtyBitmapClear) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionBlockDirtyBitmapClear) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "block-dirty-bitmap-clear",
		"data": s,
	}
	return json.Marshal(v)
}

// TransactionActionBlockdevBackup is an implementation of TransactionAction.
type TransactionActionBlockdevBackup BlockdevBackup

func (TransactionActionBlockdevBackup) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionBlockdevBackup) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "blockdev-backup",
		"data": s,
	}
	return json.Marshal(v)
}

// TransactionActionBlockdevSnapshot is an implementation of TransactionAction.
type TransactionActionBlockdevSnapshot BlockdevSnapshot

func (TransactionActionBlockdevSnapshot) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionBlockdevSnapshot) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "blockdev-snapshot",
		"data": s,
	}
	return json.Marshal(v)
}

// TransactionActionBlockdevSnapshotInternalSync is an implementation of TransactionAction.
type TransactionActionBlockdevSnapshotInternalSync BlockdevSnapshotInternal

func (TransactionActionBlockdevSnapshotInternalSync) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionBlockdevSnapshotInternalSync) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "blockdev-snapshot-internal-sync",
		"data": s,
	}
	return json.Marshal(v)
}

// TransactionActionBlockdevSnapshotSync is an implementation of TransactionAction.
type TransactionActionBlockdevSnapshotSync BlockdevSnapshotSync

func (TransactionActionBlockdevSnapshotSync) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionBlockdevSnapshotSync) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "blockdev-snapshot-sync",
		"data": s,
	}
	return json.Marshal(v)
}

// TransactionActionDriveBackup is an implementation of TransactionAction.
type TransactionActionDriveBackup DriveBackup

func (TransactionActionDriveBackup) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionDriveBackup) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "drive-backup",
		"data": s,
	}
	return json.Marshal(v)
}

func decodeTransactionAction(bs json.RawMessage) (TransactionAction, error) {
	v := struct {
		T string          `json:"type"`
		V json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.T {
	case "abort":
		var ret TransactionActionAbort
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "block-dirty-bitmap-add":
		var ret TransactionActionBlockDirtyBitmapAdd
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "block-dirty-bitmap-clear":
		var ret TransactionActionBlockDirtyBitmapClear
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "blockdev-backup":
		var ret TransactionActionBlockdevBackup
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "blockdev-snapshot":
		var ret TransactionActionBlockdevSnapshot
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "blockdev-snapshot-internal-sync":
		var ret TransactionActionBlockdevSnapshotInternalSync
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "blockdev-snapshot-sync":
		var ret TransactionActionBlockdevSnapshotSync
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "drive-backup":
		var ret TransactionActionDriveBackup
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("unknown subtype %q for union TransactionAction", v.T)
	}
}

// TransactionProperties -> TransactionProperties (struct)

// TransactionProperties implements the "TransactionProperties" QMP API type.
type TransactionProperties struct {
	CompletionMode *ActionCompletionMode `json:"completion-mode,omitempty"`
}

// UnixSocketAddress -> UnixSocketAddress (struct)

// UnixSocketAddress implements the "UnixSocketAddress" QMP API type.
type UnixSocketAddress struct {
	Path string `json:"path"`
}

// UuidInfo -> UUIDInfo (struct)

// UUIDInfo implements the "UuidInfo" QMP API type.
type UUIDInfo struct {
	UUID string `json:"UUID"`
}

// EVENT VNC_CONNECTED

// EVENT VNC_DISCONNECTED

// EVENT VNC_INITIALIZED

// EVENT VSERPORT_CHANGE

// VersionInfo -> VersionInfo (struct)

// VersionInfo implements the "VersionInfo" QMP API type.
type VersionInfo struct {
	Qemu    VersionTriple `json:"qemu"`
	Package string        `json:"package"`
}

// VersionTriple -> VersionTriple (struct)

// VersionTriple implements the "VersionTriple" QMP API type.
type VersionTriple struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor"`
	Micro int64 `json:"micro"`
}

// VncBasicInfo -> VNCBasicInfo (struct)

// VNCBasicInfo implements the "VncBasicInfo" QMP API type.
type VNCBasicInfo struct {
	Host      string               `json:"host"`
	Service   string               `json:"service"`
	Family    NetworkAddressFamily `json:"family"`
	Websocket bool                 `json:"websocket"`
}

// VncClientInfo -> VNCClientInfo (struct)

// VNCClientInfo implements the "VncClientInfo" QMP API type.
type VNCClientInfo struct {
	Host         string               `json:"host"`
	Service      string               `json:"service"`
	Family       NetworkAddressFamily `json:"family"`
	Websocket    bool                 `json:"websocket"`
	X509Dname    *string              `json:"x509_dname,omitempty"`
	SaslUsername *string              `json:"sasl_username,omitempty"`
}

// VncInfo -> VNCInfo (struct)

// VNCInfo implements the "VncInfo" QMP API type.
type VNCInfo struct {
	Enabled bool                  `json:"enabled"`
	Host    *string               `json:"host,omitempty"`
	Family  *NetworkAddressFamily `json:"family,omitempty"`
	Service *string               `json:"service,omitempty"`
	Auth    *string               `json:"auth,omitempty"`
	Clients []VNCClientInfo       `json:"clients,omitempty"`
}

// VncInfo2 -> VNCInfo2 (struct)

// VNCInfo2 implements the "VncInfo2" QMP API type.
type VNCInfo2 struct {
	ID       string              `json:"id"`
	Server   []VNCBasicInfo      `json:"server"`
	Clients  []VNCClientInfo     `json:"clients"`
	Auth     VNCPrimaryAuth      `json:"auth"`
	Vencrypt *VNCVencryptSubAuth `json:"vencrypt,omitempty"`
	Display  *string             `json:"display,omitempty"`
}

// VncPrimaryAuth -> VNCPrimaryAuth (enum)

// VNCPrimaryAuth implements the "VncPrimaryAuth" QMP API type.
type VNCPrimaryAuth int

// Known values of VNCPrimaryAuth.
const (
	VNCPrimaryAuthNone VNCPrimaryAuth = iota
	VNCPrimaryAuthVNC
	VNCPrimaryAuthRa2
	VNCPrimaryAuthRa2Ne
	VNCPrimaryAuthTight
	VNCPrimaryAuthUltra
	VNCPrimaryAuthTLS
	VNCPrimaryAuthVencrypt
	VNCPrimaryAuthSasl
)

// String implements fmt.Stringer.
func (e VNCPrimaryAuth) String() string {
	switch e {
	case VNCPrimaryAuthNone:
		return "none"
	case VNCPrimaryAuthVNC:
		return "vnc"
	case VNCPrimaryAuthRa2:
		return "ra2"
	case VNCPrimaryAuthRa2Ne:
		return "ra2ne"
	case VNCPrimaryAuthTight:
		return "tight"
	case VNCPrimaryAuthUltra:
		return "ultra"
	case VNCPrimaryAuthTLS:
		return "tls"
	case VNCPrimaryAuthVencrypt:
		return "vencrypt"
	case VNCPrimaryAuthSasl:
		return "sasl"
	default:
		return fmt.Sprintf("VNCPrimaryAuth(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e VNCPrimaryAuth) MarshalJSON() ([]byte, error) {
	switch e {
	case VNCPrimaryAuthNone:
		return json.Marshal("none")
	case VNCPrimaryAuthVNC:
		return json.Marshal("vnc")
	case VNCPrimaryAuthRa2:
		return json.Marshal("ra2")
	case VNCPrimaryAuthRa2Ne:
		return json.Marshal("ra2ne")
	case VNCPrimaryAuthTight:
		return json.Marshal("tight")
	case VNCPrimaryAuthUltra:
		return json.Marshal("ultra")
	case VNCPrimaryAuthTLS:
		return json.Marshal("tls")
	case VNCPrimaryAuthVencrypt:
		return json.Marshal("vencrypt")
	case VNCPrimaryAuthSasl:
		return json.Marshal("sasl")
	default:
		return nil, fmt.Errorf("unknown enum value %q for VNCPrimaryAuth", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *VNCPrimaryAuth) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = VNCPrimaryAuthNone
	case "vnc":
		*e = VNCPrimaryAuthVNC
	case "ra2":
		*e = VNCPrimaryAuthRa2
	case "ra2ne":
		*e = VNCPrimaryAuthRa2Ne
	case "tight":
		*e = VNCPrimaryAuthTight
	case "ultra":
		*e = VNCPrimaryAuthUltra
	case "tls":
		*e = VNCPrimaryAuthTLS
	case "vencrypt":
		*e = VNCPrimaryAuthVencrypt
	case "sasl":
		*e = VNCPrimaryAuthSasl
	default:
		return fmt.Errorf("unknown enum value %q for VNCPrimaryAuth", s)
	}
	return nil
}

// VncServerInfo -> VNCServerInfo (struct)

// VNCServerInfo implements the "VncServerInfo" QMP API type.
type VNCServerInfo struct {
	Host      string               `json:"host"`
	Service   string               `json:"service"`
	Family    NetworkAddressFamily `json:"family"`
	Websocket bool                 `json:"websocket"`
	Auth      *string              `json:"auth,omitempty"`
}

// VncVencryptSubAuth -> VNCVencryptSubAuth (enum)

// VNCVencryptSubAuth implements the "VncVencryptSubAuth" QMP API type.
type VNCVencryptSubAuth int

// Known values of VNCVencryptSubAuth.
const (
	VNCVencryptSubAuthPlain VNCVencryptSubAuth = iota
	VNCVencryptSubAuthTLSNone
	VNCVencryptSubAuthX509None
	VNCVencryptSubAuthTLSVNC
	VNCVencryptSubAuthX509VNC
	VNCVencryptSubAuthTLSPlain
	VNCVencryptSubAuthX509Plain
	VNCVencryptSubAuthTLSSasl
	VNCVencryptSubAuthX509Sasl
)

// String implements fmt.Stringer.
func (e VNCVencryptSubAuth) String() string {
	switch e {
	case VNCVencryptSubAuthPlain:
		return "plain"
	case VNCVencryptSubAuthTLSNone:
		return "tls-none"
	case VNCVencryptSubAuthX509None:
		return "x509-none"
	case VNCVencryptSubAuthTLSVNC:
		return "tls-vnc"
	case VNCVencryptSubAuthX509VNC:
		return "x509-vnc"
	case VNCVencryptSubAuthTLSPlain:
		return "tls-plain"
	case VNCVencryptSubAuthX509Plain:
		return "x509-plain"
	case VNCVencryptSubAuthTLSSasl:
		return "tls-sasl"
	case VNCVencryptSubAuthX509Sasl:
		return "x509-sasl"
	default:
		return fmt.Sprintf("VNCVencryptSubAuth(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e VNCVencryptSubAuth) MarshalJSON() ([]byte, error) {
	switch e {
	case VNCVencryptSubAuthPlain:
		return json.Marshal("plain")
	case VNCVencryptSubAuthTLSNone:
		return json.Marshal("tls-none")
	case VNCVencryptSubAuthX509None:
		return json.Marshal("x509-none")
	case VNCVencryptSubAuthTLSVNC:
		return json.Marshal("tls-vnc")
	case VNCVencryptSubAuthX509VNC:
		return json.Marshal("x509-vnc")
	case VNCVencryptSubAuthTLSPlain:
		return json.Marshal("tls-plain")
	case VNCVencryptSubAuthX509Plain:
		return json.Marshal("x509-plain")
	case VNCVencryptSubAuthTLSSasl:
		return json.Marshal("tls-sasl")
	case VNCVencryptSubAuthX509Sasl:
		return json.Marshal("x509-sasl")
	default:
		return nil, fmt.Errorf("unknown enum value %q for VNCVencryptSubAuth", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *VNCVencryptSubAuth) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "plain":
		*e = VNCVencryptSubAuthPlain
	case "tls-none":
		*e = VNCVencryptSubAuthTLSNone
	case "x509-none":
		*e = VNCVencryptSubAuthX509None
	case "tls-vnc":
		*e = VNCVencryptSubAuthTLSVNC
	case "x509-vnc":
		*e = VNCVencryptSubAuthX509VNC
	case "tls-plain":
		*e = VNCVencryptSubAuthTLSPlain
	case "x509-plain":
		*e = VNCVencryptSubAuthX509Plain
	case "tls-sasl":
		*e = VNCVencryptSubAuthTLSSasl
	case "x509-sasl":
		*e = VNCVencryptSubAuthX509Sasl
	default:
		return fmt.Errorf("unknown enum value %q for VNCVencryptSubAuth", s)
	}
	return nil
}

// EVENT WAKEUP

// EVENT WATCHDOG

// WatchdogExpirationAction -> WatchdogExpirationAction (enum)

// WatchdogExpirationAction implements the "WatchdogExpirationAction" QMP API type.
type WatchdogExpirationAction int

// Known values of WatchdogExpirationAction.
const (
	WatchdogExpirationActionReset WatchdogExpirationAction = iota
	WatchdogExpirationActionShutdown
	WatchdogExpirationActionPoweroff
	WatchdogExpirationActionPause
	WatchdogExpirationActionDebug
	WatchdogExpirationActionNone
	WatchdogExpirationActionInjectNmi
)

// String implements fmt.Stringer.
func (e WatchdogExpirationAction) String() string {
	switch e {
	case WatchdogExpirationActionReset:
		return "reset"
	case WatchdogExpirationActionShutdown:
		return "shutdown"
	case WatchdogExpirationActionPoweroff:
		return "poweroff"
	case WatchdogExpirationActionPause:
		return "pause"
	case WatchdogExpirationActionDebug:
		return "debug"
	case WatchdogExpirationActionNone:
		return "none"
	case WatchdogExpirationActionInjectNmi:
		return "inject-nmi"
	default:
		return fmt.Sprintf("WatchdogExpirationAction(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e WatchdogExpirationAction) MarshalJSON() ([]byte, error) {
	switch e {
	case WatchdogExpirationActionReset:
		return json.Marshal("reset")
	case WatchdogExpirationActionShutdown:
		return json.Marshal("shutdown")
	case WatchdogExpirationActionPoweroff:
		return json.Marshal("poweroff")
	case WatchdogExpirationActionPause:
		return json.Marshal("pause")
	case WatchdogExpirationActionDebug:
		return json.Marshal("debug")
	case WatchdogExpirationActionNone:
		return json.Marshal("none")
	case WatchdogExpirationActionInjectNmi:
		return json.Marshal("inject-nmi")
	default:
		return nil, fmt.Errorf("unknown enum value %q for WatchdogExpirationAction", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *WatchdogExpirationAction) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "reset":
		*e = WatchdogExpirationActionReset
	case "shutdown":
		*e = WatchdogExpirationActionShutdown
	case "poweroff":
		*e = WatchdogExpirationActionPoweroff
	case "pause":
		*e = WatchdogExpirationActionPause
	case "debug":
		*e = WatchdogExpirationActionDebug
	case "none":
		*e = WatchdogExpirationActionNone
	case "inject-nmi":
		*e = WatchdogExpirationActionInjectNmi
	default:
		return fmt.Errorf("unknown enum value %q for WatchdogExpirationAction", s)
	}
	return nil
}

// XBZRLECacheStats -> XbzrleCacheStats (struct)

// XbzrleCacheStats implements the "XBZRLECacheStats" QMP API type.
type XbzrleCacheStats struct {
	CacheSize     int64   `json:"cache-size"`
	Bytes         int64   `json:"bytes"`
	Pages         int64   `json:"pages"`
	CacheMiss     int64   `json:"cache-miss"`
	CacheMissRate float64 `json:"cache-miss-rate"`
	Overflow      int64   `json:"overflow"`
}

// COMMAND add-fd

// COMMAND add_client

// COMMAND balloon

// COMMAND block-commit

// COMMAND block-dirty-bitmap-add

// COMMAND block-dirty-bitmap-clear

// COMMAND block-dirty-bitmap-remove

// COMMAND block-job-cancel

// COMMAND block-job-complete

// COMMAND block-job-pause

// COMMAND block-job-resume

// COMMAND block-job-set-speed

// COMMAND block-set-write-threshold

// COMMAND block-stream

// COMMAND block_passwd

// COMMAND block_resize

// COMMAND block_set_io_throttle

// COMMAND blockdev-add

// COMMAND blockdev-backup

// COMMAND blockdev-change-medium

// COMMAND blockdev-close-tray

// COMMAND blockdev-mirror

// COMMAND blockdev-open-tray

// COMMAND blockdev-snapshot

// COMMAND blockdev-snapshot-delete-internal-sync

// COMMAND blockdev-snapshot-internal-sync

// COMMAND blockdev-snapshot-sync

// COMMAND change

// COMMAND change-backing-file

// COMMAND change-vnc-password

// COMMAND chardev-add

// COMMAND chardev-remove

// COMMAND client_migrate_info

// COMMAND closefd

// COMMAND cont

// COMMAND cpu

// COMMAND cpu-add

// COMMAND device-list-properties

// COMMAND device_add

// COMMAND device_del

// COMMAND drive-backup

// COMMAND drive-mirror

// COMMAND dump-guest-memory

// COMMAND dump-skeys

// COMMAND eject

// COMMAND expire_password

// COMMAND getfd

// COMMAND human-monitor-command

// COMMAND inject-nmi

// COMMAND input-send-event

// COMMAND memsave

// COMMAND migrate

// COMMAND migrate-incoming

// COMMAND migrate-set-cache-size

// COMMAND migrate-set-capabilities

// COMMAND migrate-set-parameters

// COMMAND migrate-start-postcopy

// COMMAND migrate_cancel

// COMMAND migrate_set_downtime

// COMMAND migrate_set_speed

// COMMAND nbd-server-add

// COMMAND nbd-server-start

// COMMAND nbd-server-stop

// COMMAND netdev_add

// COMMAND netdev_del

// COMMAND object-add

// COMMAND object-del

// COMMAND pmemsave

// COMMAND qmp_capabilities

// COMMAND qom-get

// COMMAND qom-list

// COMMAND qom-list-types

// COMMAND qom-set

// COMMAND query-acpi-ospm-status

// COMMAND query-balloon

// COMMAND query-block

// COMMAND query-block-jobs

// COMMAND query-blockstats

// COMMAND query-chardev

// COMMAND query-chardev-backends

// COMMAND query-command-line-options

// COMMAND query-commands

// COMMAND query-cpu-definitions

// COMMAND query-cpu-model-baseline

// COMMAND query-cpu-model-comparison

// COMMAND query-cpu-model-expansion

// COMMAND query-cpus

// COMMAND query-dump

// COMMAND query-dump-guest-memory-capability

// COMMAND query-events

// COMMAND query-fdsets

// COMMAND query-gic-capabilities

// COMMAND query-hotpluggable-cpus

// COMMAND query-iothreads

// COMMAND query-kvm

// COMMAND query-machines

// COMMAND query-memdev

// COMMAND query-memory-devices

// COMMAND query-mice

// COMMAND query-migrate

// COMMAND query-migrate-cache-size

// COMMAND query-migrate-capabilities

// COMMAND query-migrate-parameters

// COMMAND query-name

// COMMAND query-named-block-nodes

// COMMAND query-pci

// COMMAND query-qmp-schema

// COMMAND query-rocker

// COMMAND query-rocker-of-dpa-flows

// COMMAND query-rocker-of-dpa-groups

// COMMAND query-rocker-ports

// COMMAND query-rx-filter

// COMMAND query-spice

// COMMAND query-status

// COMMAND query-target

// COMMAND query-tpm

// COMMAND query-tpm-models

// COMMAND query-tpm-types

// COMMAND query-uuid

// COMMAND query-version

// COMMAND query-vnc

// COMMAND query-vnc-servers

// COMMAND quit

// COMMAND remove-fd

// COMMAND ringbuf-read

// COMMAND ringbuf-write

// COMMAND rtc-reset-reinjection

// COMMAND screendump

// COMMAND send-key

// COMMAND set_link

// COMMAND set_password

// COMMAND stop

// COMMAND system_powerdown

// COMMAND system_reset

// COMMAND system_wakeup

// COMMAND trace-event-get-state

// COMMAND trace-event-set-state

// COMMAND transaction

// COMMAND x-blockdev-change

// COMMAND x-blockdev-del

// COMMAND x-blockdev-insert-medium

// COMMAND x-blockdev-remove-medium

// COMMAND xen-load-devices-state

// COMMAND xen-save-devices-state

// COMMAND xen-set-global-dirty-log
