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

// IsNullable is implemented by any
// JSON null type
type IsNullable interface {
	isNull() bool
}

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
	BlkdebugEventReadAIO
	BlkdebugEventReadBackingAIO
	BlkdebugEventReadCompressed
	BlkdebugEventWriteAIO
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
	BlkdebugEventL1ShrinkWriteTable
	BlkdebugEventL1ShrinkFreeL2Clusters
	BlkdebugEventCorWrite
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
	case BlkdebugEventReadAIO:
		return "read_aio"
	case BlkdebugEventReadBackingAIO:
		return "read_backing_aio"
	case BlkdebugEventReadCompressed:
		return "read_compressed"
	case BlkdebugEventWriteAIO:
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
	case BlkdebugEventL1ShrinkWriteTable:
		return "l1_shrink_write_table"
	case BlkdebugEventL1ShrinkFreeL2Clusters:
		return "l1_shrink_free_l2_clusters"
	case BlkdebugEventCorWrite:
		return "cor_write"
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
	case BlkdebugEventReadAIO:
		return json.Marshal("read_aio")
	case BlkdebugEventReadBackingAIO:
		return json.Marshal("read_backing_aio")
	case BlkdebugEventReadCompressed:
		return json.Marshal("read_compressed")
	case BlkdebugEventWriteAIO:
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
	case BlkdebugEventL1ShrinkWriteTable:
		return json.Marshal("l1_shrink_write_table")
	case BlkdebugEventL1ShrinkFreeL2Clusters:
		return json.Marshal("l1_shrink_free_l2_clusters")
	case BlkdebugEventCorWrite:
		return json.Marshal("cor_write")
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
		*e = BlkdebugEventReadAIO
	case "read_backing_aio":
		*e = BlkdebugEventReadBackingAIO
	case "read_compressed":
		*e = BlkdebugEventReadCompressed
	case "write_aio":
		*e = BlkdebugEventWriteAIO
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
	case "l1_shrink_write_table":
		*e = BlkdebugEventL1ShrinkWriteTable
	case "l1_shrink_free_l2_clusters":
		*e = BlkdebugEventL1ShrinkFreeL2Clusters
	case "cor_write":
		*e = BlkdebugEventCorWrite
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
	Persistent  *bool   `json:"persistent,omitempty"`
	Autoload    *bool   `json:"autoload,omitempty"`
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
	Qdev         *string              `json:"qdev,omitempty"`
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

// BlockdevAioOptions -> BlockdevAIOOptions (enum)

// BlockdevAIOOptions implements the "BlockdevAioOptions" QMP API type.
type BlockdevAIOOptions int

// Known values of BlockdevAIOOptions.
const (
	BlockdevAIOOptionsThreads BlockdevAIOOptions = iota
	BlockdevAIOOptionsNative
)

// String implements fmt.Stringer.
func (e BlockdevAIOOptions) String() string {
	switch e {
	case BlockdevAIOOptionsThreads:
		return "threads"
	case BlockdevAIOOptionsNative:
		return "native"
	default:
		return fmt.Sprintf("BlockdevAIOOptions(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevAIOOptions) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevAIOOptionsThreads:
		return json.Marshal("threads")
	case BlockdevAIOOptionsNative:
		return json.Marshal("native")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevAIOOptions", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevAIOOptions) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "threads":
		*e = BlockdevAIOOptionsThreads
	case "native":
		*e = BlockdevAIOOptionsNative
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevAIOOptions", s)
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
	BlockdevDriverBlkdebug BlockdevDriver = iota
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
	BlockdevDriverIscsi
	BlockdevDriverLUKS
	BlockdevDriverNBD
	BlockdevDriverNfs
	BlockdevDriverNullAIO
	BlockdevDriverNullCo
	BlockdevDriverParallels
	BlockdevDriverQcow
	BlockdevDriverQcow2
	BlockdevDriverQed
	BlockdevDriverQuorum
	BlockdevDriverRaw
	BlockdevDriverRbd
	BlockdevDriverReplication
	BlockdevDriverSheepdog
	BlockdevDriverSSH
	BlockdevDriverThrottle
	BlockdevDriverVdi
	BlockdevDriverVhdx
	BlockdevDriverVMDK
	BlockdevDriverVpc
	BlockdevDriverVvfat
	BlockdevDriverVxhs
)

// String implements fmt.Stringer.
func (e BlockdevDriver) String() string {
	switch e {
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
	case BlockdevDriverIscsi:
		return "iscsi"
	case BlockdevDriverLUKS:
		return "luks"
	case BlockdevDriverNBD:
		return "nbd"
	case BlockdevDriverNfs:
		return "nfs"
	case BlockdevDriverNullAIO:
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
	case BlockdevDriverRbd:
		return "rbd"
	case BlockdevDriverReplication:
		return "replication"
	case BlockdevDriverSheepdog:
		return "sheepdog"
	case BlockdevDriverSSH:
		return "ssh"
	case BlockdevDriverThrottle:
		return "throttle"
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
	case BlockdevDriverVxhs:
		return "vxhs"
	default:
		return fmt.Sprintf("BlockdevDriver(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevDriver) MarshalJSON() ([]byte, error) {
	switch e {
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
	case BlockdevDriverIscsi:
		return json.Marshal("iscsi")
	case BlockdevDriverLUKS:
		return json.Marshal("luks")
	case BlockdevDriverNBD:
		return json.Marshal("nbd")
	case BlockdevDriverNfs:
		return json.Marshal("nfs")
	case BlockdevDriverNullAIO:
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
	case BlockdevDriverRbd:
		return json.Marshal("rbd")
	case BlockdevDriverReplication:
		return json.Marshal("replication")
	case BlockdevDriverSheepdog:
		return json.Marshal("sheepdog")
	case BlockdevDriverSSH:
		return json.Marshal("ssh")
	case BlockdevDriverThrottle:
		return json.Marshal("throttle")
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
	case BlockdevDriverVxhs:
		return json.Marshal("vxhs")
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
	case "iscsi":
		*e = BlockdevDriverIscsi
	case "luks":
		*e = BlockdevDriverLUKS
	case "nbd":
		*e = BlockdevDriverNBD
	case "nfs":
		*e = BlockdevDriverNfs
	case "null-aio":
		*e = BlockdevDriverNullAIO
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
	case "rbd":
		*e = BlockdevDriverRbd
	case "replication":
		*e = BlockdevDriverReplication
	case "sheepdog":
		*e = BlockdevDriverSheepdog
	case "ssh":
		*e = BlockdevDriverSSH
	case "throttle":
		*e = BlockdevDriverThrottle
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
	case "vxhs":
		*e = BlockdevDriverVxhs
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
//   - BlockdevOptionsIscsi
//   - BlockdevOptionsLUKS
//   - BlockdevOptionsNBD
//   - BlockdevOptionsNfs
//   - BlockdevOptionsNullAIO
//   - BlockdevOptionsNullCo
//   - BlockdevOptionsParallels
//   - BlockdevOptionsQcow
//   - BlockdevOptionsQcow2
//   - BlockdevOptionsQed
//   - BlockdevOptionsQuorum
//   - BlockdevOptionsRaw
//   - BlockdevOptionsRbd
//   - BlockdevOptionsReplication
//   - BlockdevOptionsSheepdog
//   - BlockdevOptionsSSH
//   - BlockdevOptionsThrottle
//   - BlockdevOptionsVdi
//   - BlockdevOptionsVhdx
//   - BlockdevOptionsVMDK
//   - BlockdevOptionsVpc
//   - BlockdevOptionsVvfat
//   - BlockdevOptionsVxhs
type BlockdevOptions interface {
	isBlockdevOptions()
}

// BlockdevOptionsBlkdebug is an implementation of BlockdevOptions.
type BlockdevOptionsBlkdebug struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Image        BlockdevRef                  `json:"image"`
	Config       *string                      `json:"config,omitempty"`
	Align        *int64                       `json:"align,omitempty"`
	MaxTransfer  *int32                       `json:"max-transfer,omitempty"`
	OptWriteZero *int32                       `json:"opt-write-zero,omitempty"`
	MaxWriteZero *int32                       `json:"max-write-zero,omitempty"`
	OptDiscard   *int32                       `json:"opt-discard,omitempty"`
	MaxDiscard   *int32                       `json:"max-discard,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
	PrManager    *string                      `json:"pr-manager,omitempty"`
	Locking      *OnOffAuto                   `json:"locking,omitempty"`
	AIO          *BlockdevAIOOptions          `json:"aio,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Sslverify    *bool                        `json:"sslverify,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Volume       string                       `json:"volume"`
	Path         string                       `json:"path"`
	Server       []SocketAddress              `json:"server"`
	Debug        *int64                       `json:"debug,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
	PrManager    *string                      `json:"pr-manager,omitempty"`
	Locking      *OnOffAuto                   `json:"locking,omitempty"`
	AIO          *BlockdevAIOOptions          `json:"aio,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename     string                       `json:"filename"`
	PrManager    *string                      `json:"pr-manager,omitempty"`
	Locking      *OnOffAuto                   `json:"locking,omitempty"`
	AIO          *BlockdevAIOOptions          `json:"aio,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Cookie       *string                      `json:"cookie,omitempty"`
	CookieSecret *string                      `json:"cookie-secret,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Cookie       *string                      `json:"cookie,omitempty"`
	Sslverify    *bool                        `json:"sslverify,omitempty"`
	CookieSecret *string                      `json:"cookie-secret,omitempty"`
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

// BlockdevOptionsIscsi is an implementation of BlockdevOptions.
type BlockdevOptionsIscsi struct {
	NodeName       *string                      `json:"node-name,omitempty"`
	Discard        *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache          *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly       *bool                        `json:"read-only,omitempty"`
	ForceShare     *bool                        `json:"force-share,omitempty"`
	DetectZeroes   *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Transport      IscsiTransport               `json:"transport"`
	Portal         string                       `json:"portal"`
	Target         string                       `json:"target"`
	Lun            *int64                       `json:"lun,omitempty"`
	User           *string                      `json:"user,omitempty"`
	PasswordSecret *string                      `json:"password-secret,omitempty"`
	InitiatorName  *string                      `json:"initiator-name,omitempty"`
	HeaderDigest   *IscsiHeaderDigest           `json:"header-digest,omitempty"`
	Timeout        *int64                       `json:"timeout,omitempty"`
}

func (BlockdevOptionsIscsi) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsIscsi) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsIscsi
	}{
		BlockdevDriverIscsi,
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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

// BlockdevOptionsNBD is an implementation of BlockdevOptions.
type BlockdevOptionsNBD struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Server       SocketAddress                `json:"server"`
	Export       *string                      `json:"export,omitempty"`
	TLSCreds     *string                      `json:"tls-creds,omitempty"`
}

func (BlockdevOptionsNBD) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsNBD) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsNBD
	}{
		BlockdevDriverNBD,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsNfs is an implementation of BlockdevOptions.
type BlockdevOptionsNfs struct {
	NodeName      *string                      `json:"node-name,omitempty"`
	Discard       *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache         *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly      *bool                        `json:"read-only,omitempty"`
	ForceShare    *bool                        `json:"force-share,omitempty"`
	DetectZeroes  *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Server        NfsServer                    `json:"server"`
	Path          string                       `json:"path"`
	User          *int64                       `json:"user,omitempty"`
	Group         *int64                       `json:"group,omitempty"`
	TCPSynCount   *int64                       `json:"tcp-syn-count,omitempty"`
	ReadaheadSize *int64                       `json:"readahead-size,omitempty"`
	PageCacheSize *int64                       `json:"page-cache-size,omitempty"`
	Debug         *int64                       `json:"debug,omitempty"`
}

func (BlockdevOptionsNfs) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsNfs) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsNfs
	}{
		BlockdevDriverNfs,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsNullAIO is an implementation of BlockdevOptions.
type BlockdevOptionsNullAIO struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Size         *int64                       `json:"size,omitempty"`
	LatencyNs    *uint64                      `json:"latency-ns,omitempty"`
}

func (BlockdevOptionsNullAIO) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsNullAIO) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsNullAIO
	}{
		BlockdevDriverNullAIO,
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Encrypt      BlockdevQcowEncryption       `json:"encrypt,omitempty"`
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
	ForceShare          *bool                        `json:"force-share,omitempty"`
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
	Encrypt             BlockdevQcow2Encryption      `json:"encrypt,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Backing      BlockdevRefOrNull            `json:"backing,omitempty"`
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
	ForceShare       *bool                        `json:"force-share,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Offset       *int64                       `json:"offset,omitempty"`
	Size         *int64                       `json:"size,omitempty"`
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

// BlockdevOptionsRbd is an implementation of BlockdevOptions.
type BlockdevOptionsRbd struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Pool         string                       `json:"pool"`
	Image        string                       `json:"image"`
	Conf         *string                      `json:"conf,omitempty"`
	Snapshot     *string                      `json:"snapshot,omitempty"`
	User         *string                      `json:"user,omitempty"`
	Server       []InetSocketAddressBase      `json:"server,omitempty"`
}

func (BlockdevOptionsRbd) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsRbd) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsRbd
	}{
		BlockdevDriverRbd,
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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

// BlockdevOptionsSheepdog is an implementation of BlockdevOptions.
type BlockdevOptionsSheepdog struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Server       SocketAddress                `json:"server"`
	Vdi          string                       `json:"vdi"`
	SnapID       *uint32                      `json:"snap-id,omitempty"`
	Tag          *string                      `json:"tag,omitempty"`
}

func (BlockdevOptionsSheepdog) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsSheepdog) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsSheepdog
	}{
		BlockdevDriverSheepdog,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsSSH is an implementation of BlockdevOptions.
type BlockdevOptionsSSH struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Server       InetSocketAddress            `json:"server"`
	Path         string                       `json:"path"`
	User         *string                      `json:"user,omitempty"`
}

func (BlockdevOptionsSSH) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsSSH) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsSSH
	}{
		BlockdevDriverSSH,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsThrottle is an implementation of BlockdevOptions.
type BlockdevOptionsThrottle struct {
	NodeName      *string                      `json:"node-name,omitempty"`
	Discard       *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache         *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly      *bool                        `json:"read-only,omitempty"`
	ForceShare    *bool                        `json:"force-share,omitempty"`
	DetectZeroes  *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	ThrottleGroup string                       `json:"throttle-group"`
	File          BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsThrottle) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsThrottle) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsThrottle
	}{
		BlockdevDriverThrottle,
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Backing      BlockdevRefOrNull            `json:"backing,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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
	ForceShare   *bool                        `json:"force-share,omitempty"`
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

// BlockdevOptionsVxhs is an implementation of BlockdevOptions.
type BlockdevOptionsVxhs struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	VdiskID      string                       `json:"vdisk-id"`
	Server       InetSocketAddressBase        `json:"server"`
	TLSCreds     *string                      `json:"tls-creds,omitempty"`
}

func (BlockdevOptionsVxhs) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVxhs) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVxhs
	}{
		BlockdevDriverVxhs,
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
	case BlockdevDriverIscsi:
		var ret BlockdevOptionsIscsi
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverLUKS:
		var ret BlockdevOptionsLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNBD:
		var ret BlockdevOptionsNBD
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNfs:
		var ret BlockdevOptionsNfs
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNullAIO:
		var ret BlockdevOptionsNullAIO
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
	case BlockdevDriverRbd:
		var ret BlockdevOptionsRbd
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverReplication:
		var ret BlockdevOptionsReplication
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverSheepdog:
		var ret BlockdevOptionsSheepdog
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverSSH:
		var ret BlockdevOptionsSSH
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverThrottle:
		var ret BlockdevOptionsThrottle
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
	case BlockdevDriverVxhs:
		var ret BlockdevOptionsVxhs
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union BlockdevOptions", v.Driver)
	}
}

// BlockdevQcow2Encryption -> BlockdevQcow2Encryption (flat union)

// BlockdevQcow2Encryption implements the "BlockdevQcow2Encryption" QMP API type.
//
// Can be one of:
//   - BlockdevQcow2EncryptionAes
//   - BlockdevQcow2EncryptionLUKS
type BlockdevQcow2Encryption interface {
	isBlockdevQcow2Encryption()
}

// BlockdevQcow2EncryptionAes is an implementation of BlockdevQcow2Encryption.
type BlockdevQcow2EncryptionAes struct {
	KeySecret *string `json:"key-secret,omitempty"`
}

func (BlockdevQcow2EncryptionAes) isBlockdevQcow2Encryption() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevQcow2EncryptionAes) MarshalJSON() ([]byte, error) {
	v := struct {
		Format BlockdevQcow2EncryptionFormat `json:"format"`
		BlockdevQcow2EncryptionAes
	}{
		BlockdevQcow2EncryptionFormatAes,
		s,
	}
	return json.Marshal(v)
}

// BlockdevQcow2EncryptionLUKS is an implementation of BlockdevQcow2Encryption.
type BlockdevQcow2EncryptionLUKS struct {
	KeySecret *string `json:"key-secret,omitempty"`
}

func (BlockdevQcow2EncryptionLUKS) isBlockdevQcow2Encryption() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevQcow2EncryptionLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Format BlockdevQcow2EncryptionFormat `json:"format"`
		BlockdevQcow2EncryptionLUKS
	}{
		BlockdevQcow2EncryptionFormatLUKS,
		s,
	}
	return json.Marshal(v)
}

func decodeBlockdevQcow2Encryption(bs json.RawMessage) (BlockdevQcow2Encryption, error) {
	v := struct {
		Format BlockdevQcow2EncryptionFormat `json:"format"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Format {
	case BlockdevQcow2EncryptionFormatAes:
		var ret BlockdevQcow2EncryptionAes
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevQcow2EncryptionFormatLUKS:
		var ret BlockdevQcow2EncryptionLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union BlockdevQcow2Encryption", v.Format)
	}
}

// BlockdevQcow2EncryptionFormat -> BlockdevQcow2EncryptionFormat (enum)

// BlockdevQcow2EncryptionFormat implements the "BlockdevQcow2EncryptionFormat" QMP API type.
type BlockdevQcow2EncryptionFormat int

// Known values of BlockdevQcow2EncryptionFormat.
const (
	BlockdevQcow2EncryptionFormatAes BlockdevQcow2EncryptionFormat = iota
	BlockdevQcow2EncryptionFormatLUKS
)

// String implements fmt.Stringer.
func (e BlockdevQcow2EncryptionFormat) String() string {
	switch e {
	case BlockdevQcow2EncryptionFormatAes:
		return "aes"
	case BlockdevQcow2EncryptionFormatLUKS:
		return "luks"
	default:
		return fmt.Sprintf("BlockdevQcow2EncryptionFormat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevQcow2EncryptionFormat) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevQcow2EncryptionFormatAes:
		return json.Marshal("aes")
	case BlockdevQcow2EncryptionFormatLUKS:
		return json.Marshal("luks")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevQcow2EncryptionFormat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevQcow2EncryptionFormat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "aes":
		*e = BlockdevQcow2EncryptionFormatAes
	case "luks":
		*e = BlockdevQcow2EncryptionFormatLUKS
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevQcow2EncryptionFormat", s)
	}
	return nil
}

// BlockdevQcowEncryption -> BlockdevQcowEncryption (flat union)

// BlockdevQcowEncryption implements the "BlockdevQcowEncryption" QMP API type.
//
// Can be one of:
//   - BlockdevQcowEncryptionAes
type BlockdevQcowEncryption interface {
	isBlockdevQcowEncryption()
}

// BlockdevQcowEncryptionAes is an implementation of BlockdevQcowEncryption.
type BlockdevQcowEncryptionAes struct {
	KeySecret *string `json:"key-secret,omitempty"`
}

func (BlockdevQcowEncryptionAes) isBlockdevQcowEncryption() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevQcowEncryptionAes) MarshalJSON() ([]byte, error) {
	v := struct {
		Format BlockdevQcowEncryptionFormat `json:"format"`
		BlockdevQcowEncryptionAes
	}{
		BlockdevQcowEncryptionFormatAes,
		s,
	}
	return json.Marshal(v)
}

func decodeBlockdevQcowEncryption(bs json.RawMessage) (BlockdevQcowEncryption, error) {
	v := struct {
		Format BlockdevQcowEncryptionFormat `json:"format"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Format {
	case BlockdevQcowEncryptionFormatAes:
		var ret BlockdevQcowEncryptionAes
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union BlockdevQcowEncryption", v.Format)
	}
}

// BlockdevQcowEncryptionFormat -> BlockdevQcowEncryptionFormat (enum)

// BlockdevQcowEncryptionFormat implements the "BlockdevQcowEncryptionFormat" QMP API type.
type BlockdevQcowEncryptionFormat int

// Known values of BlockdevQcowEncryptionFormat.
const (
	BlockdevQcowEncryptionFormatAes BlockdevQcowEncryptionFormat = iota
)

// String implements fmt.Stringer.
func (e BlockdevQcowEncryptionFormat) String() string {
	switch e {
	case BlockdevQcowEncryptionFormatAes:
		return "aes"
	default:
		return fmt.Sprintf("BlockdevQcowEncryptionFormat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevQcowEncryptionFormat) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevQcowEncryptionFormatAes:
		return json.Marshal("aes")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevQcowEncryptionFormat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevQcowEncryptionFormat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "aes":
		*e = BlockdevQcowEncryptionFormatAes
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevQcowEncryptionFormat", s)
	}
	return nil
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
func (BlockdevOptionsIscsi) isBlockdevRef()       {}
func (BlockdevOptionsLUKS) isBlockdevRef()        {}
func (BlockdevOptionsNBD) isBlockdevRef()         {}
func (BlockdevOptionsNfs) isBlockdevRef()         {}
func (BlockdevOptionsNullAIO) isBlockdevRef()     {}
func (BlockdevOptionsNullCo) isBlockdevRef()      {}
func (BlockdevOptionsParallels) isBlockdevRef()   {}
func (BlockdevOptionsQcow) isBlockdevRef()        {}
func (BlockdevOptionsQcow2) isBlockdevRef()       {}
func (BlockdevOptionsQed) isBlockdevRef()         {}
func (BlockdevOptionsQuorum) isBlockdevRef()      {}
func (BlockdevOptionsRaw) isBlockdevRef()         {}
func (BlockdevOptionsRbd) isBlockdevRef()         {}
func (BlockdevOptionsReplication) isBlockdevRef() {}
func (BlockdevOptionsSheepdog) isBlockdevRef()    {}
func (BlockdevOptionsSSH) isBlockdevRef()         {}
func (BlockdevOptionsThrottle) isBlockdevRef()    {}
func (BlockdevOptionsVdi) isBlockdevRef()         {}
func (BlockdevOptionsVhdx) isBlockdevRef()        {}
func (BlockdevOptionsVMDK) isBlockdevRef()        {}
func (BlockdevOptionsVpc) isBlockdevRef()         {}
func (BlockdevOptionsVvfat) isBlockdevRef()       {}
func (BlockdevOptionsVxhs) isBlockdevRef()        {}

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
		case BlockdevOptionsIscsi:
			return impl, nil
		case BlockdevOptionsLUKS:
			return impl, nil
		case BlockdevOptionsNBD:
			return impl, nil
		case BlockdevOptionsNfs:
			return impl, nil
		case BlockdevOptionsNullAIO:
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
		case BlockdevOptionsRbd:
			return impl, nil
		case BlockdevOptionsReplication:
			return impl, nil
		case BlockdevOptionsSheepdog:
			return impl, nil
		case BlockdevOptionsSSH:
			return impl, nil
		case BlockdevOptionsThrottle:
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
		case BlockdevOptionsVxhs:
			return impl, nil
		}
	}
	return nil, fmt.Errorf("unable to decode %q as a BlockdevRef", string(bs))
}

// BlockdevRefOrNull -> BlockdevRefOrNull (alternate)

// BlockdevRefOrNull implements the "BlockdevRefOrNull" QMP API type.
//
// Can be one of:
//   - BlockdevRefOrNullDefinition
//   - BlockdevRefOrNullNull
//   - BlockdevRefOrNullReference
type BlockdevRefOrNull interface {
	isBlockdevRefOrNull()
}

// BlockdevRefOrNullDefinition is an implementation of BlockdevRefOrNull
type BlockdevRefOrNullDefinition BlockdevOptions

func (BlockdevOptionsBlkdebug) isBlockdevRefOrNull()    {}
func (BlockdevOptionsBlkverify) isBlockdevRefOrNull()   {}
func (BlockdevOptionsBochs) isBlockdevRefOrNull()       {}
func (BlockdevOptionsCloop) isBlockdevRefOrNull()       {}
func (BlockdevOptionsDmg) isBlockdevRefOrNull()         {}
func (BlockdevOptionsFile) isBlockdevRefOrNull()        {}
func (BlockdevOptionsFTP) isBlockdevRefOrNull()         {}
func (BlockdevOptionsFTPS) isBlockdevRefOrNull()        {}
func (BlockdevOptionsGluster) isBlockdevRefOrNull()     {}
func (BlockdevOptionsHostCdrom) isBlockdevRefOrNull()   {}
func (BlockdevOptionsHostDevice) isBlockdevRefOrNull()  {}
func (BlockdevOptionsHTTP) isBlockdevRefOrNull()        {}
func (BlockdevOptionsHTTPS) isBlockdevRefOrNull()       {}
func (BlockdevOptionsIscsi) isBlockdevRefOrNull()       {}
func (BlockdevOptionsLUKS) isBlockdevRefOrNull()        {}
func (BlockdevOptionsNBD) isBlockdevRefOrNull()         {}
func (BlockdevOptionsNfs) isBlockdevRefOrNull()         {}
func (BlockdevOptionsNullAIO) isBlockdevRefOrNull()     {}
func (BlockdevOptionsNullCo) isBlockdevRefOrNull()      {}
func (BlockdevOptionsParallels) isBlockdevRefOrNull()   {}
func (BlockdevOptionsQcow) isBlockdevRefOrNull()        {}
func (BlockdevOptionsQcow2) isBlockdevRefOrNull()       {}
func (BlockdevOptionsQed) isBlockdevRefOrNull()         {}
func (BlockdevOptionsQuorum) isBlockdevRefOrNull()      {}
func (BlockdevOptionsRaw) isBlockdevRefOrNull()         {}
func (BlockdevOptionsRbd) isBlockdevRefOrNull()         {}
func (BlockdevOptionsReplication) isBlockdevRefOrNull() {}
func (BlockdevOptionsSheepdog) isBlockdevRefOrNull()    {}
func (BlockdevOptionsSSH) isBlockdevRefOrNull()         {}
func (BlockdevOptionsThrottle) isBlockdevRefOrNull()    {}
func (BlockdevOptionsVdi) isBlockdevRefOrNull()         {}
func (BlockdevOptionsVhdx) isBlockdevRefOrNull()        {}
func (BlockdevOptionsVMDK) isBlockdevRefOrNull()        {}
func (BlockdevOptionsVpc) isBlockdevRefOrNull()         {}
func (BlockdevOptionsVvfat) isBlockdevRefOrNull()       {}
func (BlockdevOptionsVxhs) isBlockdevRefOrNull()        {}

// BlockdevRefOrNullNull is a JSON null type, so it must
// also implement the isNullable interface.
type BlockdevRefOrNullNull struct{}

func (BlockdevRefOrNullNull) isNull() bool         { return true }
func (BlockdevRefOrNullNull) isBlockdevRefOrNull() {}

// BlockdevRefOrNullReference is an implementation of BlockdevRefOrNull
type BlockdevRefOrNullReference string

func (BlockdevRefOrNullReference) isBlockdevRefOrNull() {}

func decodeBlockdevRefOrNull(bs json.RawMessage) (BlockdevRefOrNull, error) {

	// Always try unmarshalling for nil first if it's an option
	// because other types could unmarshal successfully in the case
	// where a Null json type was provided.
	var null *int
	if err := json.Unmarshal([]byte(bs), &null); err == nil {
		if null == nil {
			return BlockdevRefOrNullNull{}, nil
		}
	}
	var reference BlockdevRefOrNullReference
	if err := json.Unmarshal([]byte(bs), &reference); err == nil {
		return reference, nil
	}

	if definition, err := decodeBlockdevOptions([]byte(bs)); err == nil {
		switch impl := definition.(type) {
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
		case BlockdevOptionsIscsi:
			return impl, nil
		case BlockdevOptionsLUKS:
			return impl, nil
		case BlockdevOptionsNBD:
			return impl, nil
		case BlockdevOptionsNfs:
			return impl, nil
		case BlockdevOptionsNullAIO:
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
		case BlockdevOptionsRbd:
			return impl, nil
		case BlockdevOptionsReplication:
			return impl, nil
		case BlockdevOptionsSheepdog:
			return impl, nil
		case BlockdevOptionsSSH:
			return impl, nil
		case BlockdevOptionsThrottle:
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
		case BlockdevOptionsVxhs:
			return impl, nil
		}
	}
	return nil, fmt.Errorf("unable to decode %q as a BlockdevRefOrNull", string(bs))
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
//   - ChardevBackendWctablet
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

// ChardevBackendWctablet is an implementation of ChardevBackend.
type ChardevBackendWctablet ChardevCommon

func (ChardevBackendWctablet) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendWctablet) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "wctablet",
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
	case "wctablet":
		var ret ChardevBackendWctablet
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
	Logfile   *string             `json:"logfile,omitempty"`
	Logappend *bool               `json:"logappend,omitempty"`
	Addr      SocketAddressLegacy `json:"addr"`
	TLSCreds  *string             `json:"tls-creds,omitempty"`
	Server    *bool               `json:"server,omitempty"`
	Wait      *bool               `json:"wait,omitempty"`
	Nodelay   *bool               `json:"nodelay,omitempty"`
	Telnet    *bool               `json:"telnet,omitempty"`
	Tn3270    *bool               `json:"tn3270,omitempty"`
	Reconnect *int64              `json:"reconnect,omitempty"`
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
		Tn3270    *bool           `json:"tn3270,omitempty"`
		Reconnect *int64          `json:"reconnect,omitempty"`
	}{}
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return err
	}
	s.Logfile = v.Logfile
	s.Logappend = v.Logappend
	s.Addr, err = decodeSocketAddressLegacy(v.Addr)
	if err != nil {
		return err
	}
	s.TLSCreds = v.TLSCreds
	s.Server = v.Server
	s.Wait = v.Wait
	s.Nodelay = v.Nodelay
	s.Telnet = v.Telnet
	s.Tn3270 = v.Tn3270
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
	Logfile   *string             `json:"logfile,omitempty"`
	Logappend *bool               `json:"logappend,omitempty"`
	Remote    SocketAddressLegacy `json:"remote"`
	Local     SocketAddressLegacy `json:"local,omitempty"`
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
	s.Remote, err = decodeSocketAddressLegacy(v.Remote)
	if err != nil {
		return err
	}
	if len(v.Local) > 0 {
		s.Local, err = decodeSocketAddressLegacy(v.Local)
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
	Name                string   `json:"name"`
	MigrationSafe       *bool    `json:"migration-safe,omitempty"`
	Static              bool     `json:"static"`
	UnavailableFeatures []string `json:"unavailable-features,omitempty"`
	Typename            string   `json:"typename"`
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
	CPU      int64                  `json:"CPU"`
	Current  bool                   `json:"current"`
	Halted   bool                   `json:"halted"`
	QomPath  string                 `json:"qom_path"`
	ThreadID int64                  `json:"thread_id"`
	Props    *CPUInstanceProperties `json:"props,omitempty"`
	Pc       int64                  `json:"PC"`
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
	CPU      int64                  `json:"CPU"`
	Current  bool                   `json:"current"`
	Halted   bool                   `json:"halted"`
	QomPath  string                 `json:"qom_path"`
	ThreadID int64                  `json:"thread_id"`
	Props    *CPUInstanceProperties `json:"props,omitempty"`
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
	CPU      int64                  `json:"CPU"`
	Current  bool                   `json:"current"`
	Halted   bool                   `json:"halted"`
	QomPath  string                 `json:"qom_path"`
	ThreadID int64                  `json:"thread_id"`
	Props    *CPUInstanceProperties `json:"props,omitempty"`
	Nip      int64                  `json:"nip"`
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
	CPU      int64                  `json:"CPU"`
	Current  bool                   `json:"current"`
	Halted   bool                   `json:"halted"`
	QomPath  string                 `json:"qom_path"`
	ThreadID int64                  `json:"thread_id"`
	Props    *CPUInstanceProperties `json:"props,omitempty"`
	Pc       int64                  `json:"pc"`
	Npc      int64                  `json:"npc"`
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
	CPU      int64                  `json:"CPU"`
	Current  bool                   `json:"current"`
	Halted   bool                   `json:"halted"`
	QomPath  string                 `json:"qom_path"`
	ThreadID int64                  `json:"thread_id"`
	Props    *CPUInstanceProperties `json:"props,omitempty"`
	Pc       int64                  `json:"PC"`
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
	CPU      int64                  `json:"CPU"`
	Current  bool                   `json:"current"`
	Halted   bool                   `json:"halted"`
	QomPath  string                 `json:"qom_path"`
	ThreadID int64                  `json:"thread_id"`
	Props    *CPUInstanceProperties `json:"props,omitempty"`
	Pc       int64                  `json:"pc"`
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

// GuestPanicAction -> GuestPanicAction (enum)

// GuestPanicAction implements the "GuestPanicAction" QMP API type.
type GuestPanicAction int

// Known values of GuestPanicAction.
const (
	GuestPanicActionPause GuestPanicAction = iota
	GuestPanicActionPoweroff
)

// String implements fmt.Stringer.
func (e GuestPanicAction) String() string {
	switch e {
	case GuestPanicActionPause:
		return "pause"
	case GuestPanicActionPoweroff:
		return "poweroff"
	default:
		return fmt.Sprintf("GuestPanicAction(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e GuestPanicAction) MarshalJSON() ([]byte, error) {
	switch e {
	case GuestPanicActionPause:
		return json.Marshal("pause")
	case GuestPanicActionPoweroff:
		return json.Marshal("poweroff")
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
	case "poweroff":
		*e = GuestPanicActionPoweroff
	default:
		return fmt.Errorf("unknown enum value %q for GuestPanicAction", s)
	}
	return nil
}

// GuestPanicInformation -> GuestPanicInformation (flat union)

// GuestPanicInformation implements the "GuestPanicInformation" QMP API type.
//
// Can be one of:
//   - GuestPanicInformationHyperV
type GuestPanicInformation interface {
	isGuestPanicInformation()
}

// GuestPanicInformationHyperV is an implementation of GuestPanicInformation.
type GuestPanicInformationHyperV struct {
	Arg1 uint64 `json:"arg1"`
	Arg2 uint64 `json:"arg2"`
	Arg3 uint64 `json:"arg3"`
	Arg4 uint64 `json:"arg4"`
	Arg5 uint64 `json:"arg5"`
}

func (GuestPanicInformationHyperV) isGuestPanicInformation() {}

// MarshalJSON implements json.Marshaler.
func (s GuestPanicInformationHyperV) MarshalJSON() ([]byte, error) {
	v := struct {
		Type GuestPanicInformationType `json:"type"`
		GuestPanicInformationHyperV
	}{
		GuestPanicInformationTypeHyperV,
		s,
	}
	return json.Marshal(v)
}

func decodeGuestPanicInformation(bs json.RawMessage) (GuestPanicInformation, error) {
	v := struct {
		Type GuestPanicInformationType `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case GuestPanicInformationTypeHyperV:
		var ret GuestPanicInformationHyperV
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union GuestPanicInformation", v.Type)
	}
}

// GuestPanicInformationType -> GuestPanicInformationType (enum)

// GuestPanicInformationType implements the "GuestPanicInformationType" QMP API type.
type GuestPanicInformationType int

// Known values of GuestPanicInformationType.
const (
	GuestPanicInformationTypeHyperV GuestPanicInformationType = iota
)

// String implements fmt.Stringer.
func (e GuestPanicInformationType) String() string {
	switch e {
	case GuestPanicInformationTypeHyperV:
		return "hyper-v"
	default:
		return fmt.Sprintf("GuestPanicInformationType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e GuestPanicInformationType) MarshalJSON() ([]byte, error) {
	switch e {
	case GuestPanicInformationTypeHyperV:
		return json.Marshal("hyper-v")
	default:
		return nil, fmt.Errorf("unknown enum value %q for GuestPanicInformationType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *GuestPanicInformationType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "hyper-v":
		*e = GuestPanicInformationTypeHyperV
	default:
		return fmt.Errorf("unknown enum value %q for GuestPanicInformationType", s)
	}
	return nil
}

// GuidInfo -> GUIDInfo (struct)

// GUIDInfo implements the "GuidInfo" QMP API type.
type GUIDInfo struct {
	GUID string `json:"guid"`
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
	ID         string `json:"id"`
	ThreadID   int64  `json:"thread-id"`
	PollMaxNs  int64  `json:"poll-max-ns"`
	PollGrow   int64  `json:"poll-grow"`
	PollShrink int64  `json:"poll-shrink"`
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
	Compat        string                           `json:"compat"`
	LazyRefcounts *bool                            `json:"lazy-refcounts,omitempty"`
	Corrupt       *bool                            `json:"corrupt,omitempty"`
	RefcountBits  int64                            `json:"refcount-bits"`
	Encrypt       ImageInfoSpecificQCow2Encryption `json:"encrypt,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *ImageInfoSpecificQCow2) UnmarshalJSON(bs []byte) error {
	v := struct {
		Compat        string          `json:"compat"`
		LazyRefcounts *bool           `json:"lazy-refcounts,omitempty"`
		Corrupt       *bool           `json:"corrupt,omitempty"`
		RefcountBits  int64           `json:"refcount-bits"`
		Encrypt       json.RawMessage `json:"encrypt"`
	}{}
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return err
	}
	s.Compat = v.Compat
	s.LazyRefcounts = v.LazyRefcounts
	s.Corrupt = v.Corrupt
	s.RefcountBits = v.RefcountBits
	if len(v.Encrypt) > 0 {
		s.Encrypt, err = decodeImageInfoSpecificQCow2Encryption(v.Encrypt)
		if err != nil {
			return err
		}
	} else {
		s.Encrypt = nil
	}

	return nil
}

// ImageInfoSpecificQCow2Encryption -> ImageInfoSpecificQCow2Encryption (flat union)

// ImageInfoSpecificQCow2Encryption implements the "ImageInfoSpecificQCow2Encryption" QMP API type.
//
// Can be one of:
//   - ImageInfoSpecificQCow2EncryptionAes
//   - ImageInfoSpecificQCow2EncryptionLUKS
type ImageInfoSpecificQCow2Encryption interface {
	isImageInfoSpecificQCow2Encryption()
}

// ImageInfoSpecificQCow2EncryptionAes is an implementation of ImageInfoSpecificQCow2Encryption.
type ImageInfoSpecificQCow2EncryptionAes struct {
}

func (ImageInfoSpecificQCow2EncryptionAes) isImageInfoSpecificQCow2Encryption() {}

// MarshalJSON implements json.Marshaler.
func (s ImageInfoSpecificQCow2EncryptionAes) MarshalJSON() ([]byte, error) {
	v := struct {
		Format BlockdevQcow2EncryptionFormat `json:"format"`
		ImageInfoSpecificQCow2EncryptionAes
	}{
		BlockdevQcow2EncryptionFormatAes,
		s,
	}
	return json.Marshal(v)
}

// ImageInfoSpecificQCow2EncryptionLUKS is an implementation of ImageInfoSpecificQCow2Encryption.
type ImageInfoSpecificQCow2EncryptionLUKS struct {
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

func (ImageInfoSpecificQCow2EncryptionLUKS) isImageInfoSpecificQCow2Encryption() {}

// MarshalJSON implements json.Marshaler.
func (s ImageInfoSpecificQCow2EncryptionLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Format BlockdevQcow2EncryptionFormat `json:"format"`
		ImageInfoSpecificQCow2EncryptionLUKS
	}{
		BlockdevQcow2EncryptionFormatLUKS,
		s,
	}
	return json.Marshal(v)
}

func decodeImageInfoSpecificQCow2Encryption(bs json.RawMessage) (ImageInfoSpecificQCow2Encryption, error) {
	v := struct {
		Format BlockdevQcow2EncryptionFormat `json:"format"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Format {
	case BlockdevQcow2EncryptionFormatAes:
		var ret ImageInfoSpecificQCow2EncryptionAes
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevQcow2EncryptionFormatLUKS:
		var ret ImageInfoSpecificQCow2EncryptionLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union ImageInfoSpecificQCow2Encryption", v.Format)
	}
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
	Host    string  `json:"host"`
	Port    string  `json:"port"`
	Numeric *bool   `json:"numeric,omitempty"`
	To      *uint16 `json:"to,omitempty"`
	Ipv4    *bool   `json:"ipv4,omitempty"`
	Ipv6    *bool   `json:"ipv6,omitempty"`
}

// InetSocketAddressBase -> InetSocketAddressBase (struct)

// InetSocketAddressBase implements the "InetSocketAddressBase" QMP API type.
type InetSocketAddressBase struct {
	Host string `json:"host"`
	Port string `json:"port"`
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
	InputButtonSide
	InputButtonExtra
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
	case InputButtonSide:
		return "side"
	case InputButtonExtra:
		return "extra"
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
	case InputButtonSide:
		return json.Marshal("side")
	case InputButtonExtra:
		return json.Marshal("extra")
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
	case "side":
		*e = InputButtonSide
	case "extra":
		*e = InputButtonExtra
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

// IscsiHeaderDigest -> IscsiHeaderDigest (enum)

// IscsiHeaderDigest implements the "IscsiHeaderDigest" QMP API type.
type IscsiHeaderDigest int

// Known values of IscsiHeaderDigest.
const (
	IscsiHeaderDigestCrc32C IscsiHeaderDigest = iota
	IscsiHeaderDigestNone
	IscsiHeaderDigestCrc32CNone
	IscsiHeaderDigestNoneCrc32C
)

// String implements fmt.Stringer.
func (e IscsiHeaderDigest) String() string {
	switch e {
	case IscsiHeaderDigestCrc32C:
		return "crc32c"
	case IscsiHeaderDigestNone:
		return "none"
	case IscsiHeaderDigestCrc32CNone:
		return "crc32c-none"
	case IscsiHeaderDigestNoneCrc32C:
		return "none-crc32c"
	default:
		return fmt.Sprintf("IscsiHeaderDigest(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e IscsiHeaderDigest) MarshalJSON() ([]byte, error) {
	switch e {
	case IscsiHeaderDigestCrc32C:
		return json.Marshal("crc32c")
	case IscsiHeaderDigestNone:
		return json.Marshal("none")
	case IscsiHeaderDigestCrc32CNone:
		return json.Marshal("crc32c-none")
	case IscsiHeaderDigestNoneCrc32C:
		return json.Marshal("none-crc32c")
	default:
		return nil, fmt.Errorf("unknown enum value %q for IscsiHeaderDigest", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *IscsiHeaderDigest) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "crc32c":
		*e = IscsiHeaderDigestCrc32C
	case "none":
		*e = IscsiHeaderDigestNone
	case "crc32c-none":
		*e = IscsiHeaderDigestCrc32CNone
	case "none-crc32c":
		*e = IscsiHeaderDigestNoneCrc32C
	default:
		return fmt.Errorf("unknown enum value %q for IscsiHeaderDigest", s)
	}
	return nil
}

// IscsiTransport -> IscsiTransport (enum)

// IscsiTransport implements the "IscsiTransport" QMP API type.
type IscsiTransport int

// Known values of IscsiTransport.
const (
	IscsiTransportTCP IscsiTransport = iota
	IscsiTransportIser
)

// String implements fmt.Stringer.
func (e IscsiTransport) String() string {
	switch e {
	case IscsiTransportTCP:
		return "tcp"
	case IscsiTransportIser:
		return "iser"
	default:
		return fmt.Sprintf("IscsiTransport(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e IscsiTransport) MarshalJSON() ([]byte, error) {
	switch e {
	case IscsiTransportTCP:
		return json.Marshal("tcp")
	case IscsiTransportIser:
		return json.Marshal("iser")
	default:
		return nil, fmt.Errorf("unknown enum value %q for IscsiTransport", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *IscsiTransport) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "tcp":
		*e = IscsiTransportTCP
	case "iser":
		*e = IscsiTransportIser
	default:
		return fmt.Errorf("unknown enum value %q for IscsiTransport", s)
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
	ID        *string       `json:"id,omitempty"`
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

// MemoryInfo -> MemoryInfo (struct)

// MemoryInfo implements the "MemoryInfo" QMP API type.
type MemoryInfo struct {
	BaseMemory    uint64  `json:"base-memory"`
	PluggedMemory *uint64 `json:"plugged-memory,omitempty"`
}

// MigrateSetParameters -> MigrateSetParameters (struct)

// MigrateSetParameters implements the "MigrateSetParameters" QMP API type.
type MigrateSetParameters struct {
	CompressLevel        *int64    `json:"compress-level,omitempty"`
	CompressThreads      *int64    `json:"compress-threads,omitempty"`
	DecompressThreads    *int64    `json:"decompress-threads,omitempty"`
	CPUThrottleInitial   *int64    `json:"cpu-throttle-initial,omitempty"`
	CPUThrottleIncrement *int64    `json:"cpu-throttle-increment,omitempty"`
	TLSCreds             StrOrNull `json:"tls-creds,omitempty"`
	TLSHostname          StrOrNull `json:"tls-hostname,omitempty"`
	MaxBandwidth         *int64    `json:"max-bandwidth,omitempty"`
	DowntimeLimit        *int64    `json:"downtime-limit,omitempty"`
	XCheckpointDelay     *int64    `json:"x-checkpoint-delay,omitempty"`
	BlockIncremental     *bool     `json:"block-incremental,omitempty"`
	XMultifdChannels     *int64    `json:"x-multifd-channels,omitempty"`
	XMultifdPageCount    *int64    `json:"x-multifd-page-count,omitempty"`
	XbzrleCacheSize      *uint64   `json:"xbzrle-cache-size,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *MigrateSetParameters) UnmarshalJSON(bs []byte) error {
	v := struct {
		CompressLevel        *int64          `json:"compress-level,omitempty"`
		CompressThreads      *int64          `json:"compress-threads,omitempty"`
		DecompressThreads    *int64          `json:"decompress-threads,omitempty"`
		CPUThrottleInitial   *int64          `json:"cpu-throttle-initial,omitempty"`
		CPUThrottleIncrement *int64          `json:"cpu-throttle-increment,omitempty"`
		TLSCreds             json.RawMessage `json:"tls-creds"`
		TLSHostname          json.RawMessage `json:"tls-hostname"`
		MaxBandwidth         *int64          `json:"max-bandwidth,omitempty"`
		DowntimeLimit        *int64          `json:"downtime-limit,omitempty"`
		XCheckpointDelay     *int64          `json:"x-checkpoint-delay,omitempty"`
		BlockIncremental     *bool           `json:"block-incremental,omitempty"`
		XMultifdChannels     *int64          `json:"x-multifd-channels,omitempty"`
		XMultifdPageCount    *int64          `json:"x-multifd-page-count,omitempty"`
		XbzrleCacheSize      *uint64         `json:"xbzrle-cache-size,omitempty"`
	}{}
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return err
	}
	s.CompressLevel = v.CompressLevel
	s.CompressThreads = v.CompressThreads
	s.DecompressThreads = v.DecompressThreads
	s.CPUThrottleInitial = v.CPUThrottleInitial
	s.CPUThrottleIncrement = v.CPUThrottleIncrement
	if len(v.TLSCreds) > 0 {
		s.TLSCreds, err = decodeStrOrNull(v.TLSCreds)
		if err != nil {
			return err
		}
	} else {
		s.TLSCreds = nil
	}
	if len(v.TLSHostname) > 0 {
		s.TLSHostname, err = decodeStrOrNull(v.TLSHostname)
		if err != nil {
			return err
		}
	} else {
		s.TLSHostname = nil
	}
	s.MaxBandwidth = v.MaxBandwidth
	s.DowntimeLimit = v.DowntimeLimit
	s.XCheckpointDelay = v.XCheckpointDelay
	s.BlockIncremental = v.BlockIncremental
	s.XMultifdChannels = v.XMultifdChannels
	s.XMultifdPageCount = v.XMultifdPageCount
	s.XbzrleCacheSize = v.XbzrleCacheSize

	return nil
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
	MigrationCapabilityXColo
	MigrationCapabilityReleaseRAM
	MigrationCapabilityBlock
	MigrationCapabilityReturnPath
	MigrationCapabilityPauseBeforeSwitchover
	MigrationCapabilityXMultifd
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
	case MigrationCapabilityXColo:
		return "x-colo"
	case MigrationCapabilityReleaseRAM:
		return "release-ram"
	case MigrationCapabilityBlock:
		return "block"
	case MigrationCapabilityReturnPath:
		return "return-path"
	case MigrationCapabilityPauseBeforeSwitchover:
		return "pause-before-switchover"
	case MigrationCapabilityXMultifd:
		return "x-multifd"
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
	case MigrationCapabilityXColo:
		return json.Marshal("x-colo")
	case MigrationCapabilityReleaseRAM:
		return json.Marshal("release-ram")
	case MigrationCapabilityBlock:
		return json.Marshal("block")
	case MigrationCapabilityReturnPath:
		return json.Marshal("return-path")
	case MigrationCapabilityPauseBeforeSwitchover:
		return json.Marshal("pause-before-switchover")
	case MigrationCapabilityXMultifd:
		return json.Marshal("x-multifd")
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
	case "x-colo":
		*e = MigrationCapabilityXColo
	case "release-ram":
		*e = MigrationCapabilityReleaseRAM
	case "block":
		*e = MigrationCapabilityBlock
	case "return-path":
		*e = MigrationCapabilityReturnPath
	case "pause-before-switchover":
		*e = MigrationCapabilityPauseBeforeSwitchover
	case "x-multifd":
		*e = MigrationCapabilityXMultifd
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
	XCheckpointDelay     *int64  `json:"x-checkpoint-delay,omitempty"`
	BlockIncremental     *bool   `json:"block-incremental,omitempty"`
	XMultifdChannels     *int64  `json:"x-multifd-channels,omitempty"`
	XMultifdPageCount    *int64  `json:"x-multifd-page-count,omitempty"`
	XbzrleCacheSize      *uint64 `json:"xbzrle-cache-size,omitempty"`
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
	PageSize         int64   `json:"page-size"`
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
	MigrationStatusColo
	MigrationStatusPreSwitchover
	MigrationStatusDevice
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
	case MigrationStatusColo:
		return "colo"
	case MigrationStatusPreSwitchover:
		return "pre-switchover"
	case MigrationStatusDevice:
		return "device"
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
	case MigrationStatusColo:
		return json.Marshal("colo")
	case MigrationStatusPreSwitchover:
		return json.Marshal("pre-switchover")
	case MigrationStatusDevice:
		return json.Marshal("device")
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
	case "colo":
		*e = MigrationStatusColo
	case "pre-switchover":
		*e = MigrationStatusPreSwitchover
	case "device":
		*e = MigrationStatusDevice
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

// NFSServer -> NfsServer (struct)

// NfsServer implements the "NFSServer" QMP API type.
type NfsServer struct {
	Type NfsTransport `json:"type"`
	Host string       `json:"host"`
}

// NFSTransport -> NfsTransport (enum)

// NfsTransport implements the "NFSTransport" QMP API type.
type NfsTransport int

// Known values of NfsTransport.
const (
	NfsTransportInet NfsTransport = iota
)

// String implements fmt.Stringer.
func (e NfsTransport) String() string {
	switch e {
	case NfsTransportInet:
		return "inet"
	default:
		return fmt.Sprintf("NfsTransport(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e NfsTransport) MarshalJSON() ([]byte, error) {
	switch e {
	case NfsTransportInet:
		return json.Marshal("inet")
	default:
		return nil, fmt.Errorf("unknown enum value %q for NfsTransport", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NfsTransport) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "inet":
		*e = NfsTransportInet
	default:
		return fmt.Errorf("unknown enum value %q for NfsTransport", s)
	}
	return nil
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
	NetworkAddressFamilyVsock
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
	case NetworkAddressFamilyVsock:
		return "vsock"
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
	case NetworkAddressFamilyVsock:
		return json.Marshal("vsock")
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
	case "vsock":
		*e = NetworkAddressFamilyVsock
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
	Name     string  `json:"name"`
	Abstract *bool   `json:"abstract,omitempty"`
	Parent   *string `json:"parent,omitempty"`
}

// OnOffAuto -> OnOffAuto (enum)

// OnOffAuto implements the "OnOffAuto" QMP API type.
type OnOffAuto int

// Known values of OnOffAuto.
const (
	OnOffAutoAuto OnOffAuto = iota
	OnOffAutoOn
	OnOffAutoOff
)

// String implements fmt.Stringer.
func (e OnOffAuto) String() string {
	switch e {
	case OnOffAutoAuto:
		return "auto"
	case OnOffAutoOn:
		return "on"
	case OnOffAutoOff:
		return "off"
	default:
		return fmt.Sprintf("OnOffAuto(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e OnOffAuto) MarshalJSON() ([]byte, error) {
	switch e {
	case OnOffAutoAuto:
		return json.Marshal("auto")
	case OnOffAutoOn:
		return json.Marshal("on")
	case OnOffAutoOff:
		return json.Marshal("off")
	default:
		return nil, fmt.Errorf("unknown enum value %q for OnOffAuto", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *OnOffAuto) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "auto":
		*e = OnOffAutoAuto
	case "on":
		*e = OnOffAutoOn
	case "off":
		*e = OnOffAutoOff
	default:
		return fmt.Errorf("unknown enum value %q for OnOffAuto", s)
	}
	return nil
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
	QCryptoCipherAlgorithm3Des
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
	case QCryptoCipherAlgorithm3Des:
		return "3des"
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
	case QCryptoCipherAlgorithm3Des:
		return json.Marshal("3des")
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
	case "3des":
		*e = QCryptoCipherAlgorithm3Des
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
	QCryptoCipherModeCtr
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
	case QCryptoCipherModeCtr:
		return "ctr"
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
	case QCryptoCipherModeCtr:
		return json.Marshal("ctr")
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
	case "ctr":
		*e = QCryptoCipherModeCtr
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
	QKeyCodeHiragana
	QKeyCodeHenkan
	QKeyCodeYen
	QKeyCodeKpComma
	QKeyCodeKpEquals
	QKeyCodePower
	QKeyCodeSleep
	QKeyCodeWake
	QKeyCodeAudionext
	QKeyCodeAudioprev
	QKeyCodeAudiostop
	QKeyCodeAudioplay
	QKeyCodeAudiomute
	QKeyCodeVolumeup
	QKeyCodeVolumedown
	QKeyCodeMediaselect
	QKeyCodeMail
	QKeyCodeCalculator
	QKeyCodeComputer
	QKeyCodeAcHome
	QKeyCodeAcBack
	QKeyCodeAcForward
	QKeyCodeAcRefresh
	QKeyCodeAcBookmarks
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
	case QKeyCodeHiragana:
		return "hiragana"
	case QKeyCodeHenkan:
		return "henkan"
	case QKeyCodeYen:
		return "yen"
	case QKeyCodeKpComma:
		return "kp_comma"
	case QKeyCodeKpEquals:
		return "kp_equals"
	case QKeyCodePower:
		return "power"
	case QKeyCodeSleep:
		return "sleep"
	case QKeyCodeWake:
		return "wake"
	case QKeyCodeAudionext:
		return "audionext"
	case QKeyCodeAudioprev:
		return "audioprev"
	case QKeyCodeAudiostop:
		return "audiostop"
	case QKeyCodeAudioplay:
		return "audioplay"
	case QKeyCodeAudiomute:
		return "audiomute"
	case QKeyCodeVolumeup:
		return "volumeup"
	case QKeyCodeVolumedown:
		return "volumedown"
	case QKeyCodeMediaselect:
		return "mediaselect"
	case QKeyCodeMail:
		return "mail"
	case QKeyCodeCalculator:
		return "calculator"
	case QKeyCodeComputer:
		return "computer"
	case QKeyCodeAcHome:
		return "ac_home"
	case QKeyCodeAcBack:
		return "ac_back"
	case QKeyCodeAcForward:
		return "ac_forward"
	case QKeyCodeAcRefresh:
		return "ac_refresh"
	case QKeyCodeAcBookmarks:
		return "ac_bookmarks"
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
	case QKeyCodeHiragana:
		return json.Marshal("hiragana")
	case QKeyCodeHenkan:
		return json.Marshal("henkan")
	case QKeyCodeYen:
		return json.Marshal("yen")
	case QKeyCodeKpComma:
		return json.Marshal("kp_comma")
	case QKeyCodeKpEquals:
		return json.Marshal("kp_equals")
	case QKeyCodePower:
		return json.Marshal("power")
	case QKeyCodeSleep:
		return json.Marshal("sleep")
	case QKeyCodeWake:
		return json.Marshal("wake")
	case QKeyCodeAudionext:
		return json.Marshal("audionext")
	case QKeyCodeAudioprev:
		return json.Marshal("audioprev")
	case QKeyCodeAudiostop:
		return json.Marshal("audiostop")
	case QKeyCodeAudioplay:
		return json.Marshal("audioplay")
	case QKeyCodeAudiomute:
		return json.Marshal("audiomute")
	case QKeyCodeVolumeup:
		return json.Marshal("volumeup")
	case QKeyCodeVolumedown:
		return json.Marshal("volumedown")
	case QKeyCodeMediaselect:
		return json.Marshal("mediaselect")
	case QKeyCodeMail:
		return json.Marshal("mail")
	case QKeyCodeCalculator:
		return json.Marshal("calculator")
	case QKeyCodeComputer:
		return json.Marshal("computer")
	case QKeyCodeAcHome:
		return json.Marshal("ac_home")
	case QKeyCodeAcBack:
		return json.Marshal("ac_back")
	case QKeyCodeAcForward:
		return json.Marshal("ac_forward")
	case QKeyCodeAcRefresh:
		return json.Marshal("ac_refresh")
	case QKeyCodeAcBookmarks:
		return json.Marshal("ac_bookmarks")
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
	case "hiragana":
		*e = QKeyCodeHiragana
	case "henkan":
		*e = QKeyCodeHenkan
	case "yen":
		*e = QKeyCodeYen
	case "kp_comma":
		*e = QKeyCodeKpComma
	case "kp_equals":
		*e = QKeyCodeKpEquals
	case "power":
		*e = QKeyCodePower
	case "sleep":
		*e = QKeyCodeSleep
	case "wake":
		*e = QKeyCodeWake
	case "audionext":
		*e = QKeyCodeAudionext
	case "audioprev":
		*e = QKeyCodeAudioprev
	case "audiostop":
		*e = QKeyCodeAudiostop
	case "audioplay":
		*e = QKeyCodeAudioplay
	case "audiomute":
		*e = QKeyCodeAudiomute
	case "volumeup":
		*e = QKeyCodeVolumeup
	case "volumedown":
		*e = QKeyCodeVolumedown
	case "mediaselect":
		*e = QKeyCodeMediaselect
	case "mail":
		*e = QKeyCodeMail
	case "calculator":
		*e = QKeyCodeCalculator
	case "computer":
		*e = QKeyCodeComputer
	case "ac_home":
		*e = QKeyCodeAcHome
	case "ac_back":
		*e = QKeyCodeAcBack
	case "ac_forward":
		*e = QKeyCodeAcForward
	case "ac_refresh":
		*e = QKeyCodeAcRefresh
	case "ac_bookmarks":
		*e = QKeyCodeAcBookmarks
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

// ReplicationStatus -> ReplicationStatus (struct)

// ReplicationStatus implements the "ReplicationStatus" QMP API type.
type ReplicationStatus struct {
	Error bool    `json:"error"`
	Desc  *string `json:"desc,omitempty"`
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
	RunStateColo
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
	case RunStateColo:
		return "colo"
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
	case RunStateColo:
		return json.Marshal("colo")
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
	case "colo":
		*e = RunStateColo
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

// SocketAddress -> SocketAddress (flat union)

// SocketAddress implements the "SocketAddress" QMP API type.
//
// Can be one of:
//   - SocketAddressFD
//   - SocketAddressInet
//   - SocketAddressUnix
//   - SocketAddressVsock
type SocketAddress interface {
	isSocketAddress()
}

// SocketAddressFD is an implementation of SocketAddress.
type SocketAddressFD struct {
	Str string `json:"str"`
}

func (SocketAddressFD) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressFD) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressFD
	}{
		SocketAddressTypeFD,
		s,
	}
	return json.Marshal(v)
}

// SocketAddressInet is an implementation of SocketAddress.
type SocketAddressInet struct {
	Numeric *bool   `json:"numeric,omitempty"`
	To      *uint16 `json:"to,omitempty"`
	Ipv4    *bool   `json:"ipv4,omitempty"`
	Ipv6    *bool   `json:"ipv6,omitempty"`
}

func (SocketAddressInet) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressInet) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressInet
	}{
		SocketAddressTypeInet,
		s,
	}
	return json.Marshal(v)
}

// SocketAddressUnix is an implementation of SocketAddress.
type SocketAddressUnix struct {
	Path string `json:"path"`
}

func (SocketAddressUnix) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressUnix) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressUnix
	}{
		SocketAddressTypeUnix,
		s,
	}
	return json.Marshal(v)
}

// SocketAddressVsock is an implementation of SocketAddress.
type SocketAddressVsock struct {
	Cid  string `json:"cid"`
	Port string `json:"port"`
}

func (SocketAddressVsock) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressVsock) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressVsock
	}{
		SocketAddressTypeVsock,
		s,
	}
	return json.Marshal(v)
}

func decodeSocketAddress(bs json.RawMessage) (SocketAddress, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case SocketAddressTypeFD:
		var ret SocketAddressFD
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SocketAddressTypeInet:
		var ret SocketAddressInet
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SocketAddressTypeUnix:
		var ret SocketAddressUnix
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SocketAddressTypeVsock:
		var ret SocketAddressVsock
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union SocketAddress", v.Type)
	}
}

// SocketAddressLegacy -> SocketAddressLegacy (simple union)

// SocketAddressLegacy implements the "SocketAddressLegacy" QMP API type.
//
// Can be one of:
//   - SocketAddressLegacyFD
//   - SocketAddressLegacyInet
//   - SocketAddressLegacyUnix
//   - SocketAddressLegacyVsock
type SocketAddressLegacy interface {
	isSocketAddressLegacy()
}

// SocketAddressLegacyFD is an implementation of SocketAddressLegacy.
type SocketAddressLegacyFD String

func (SocketAddressLegacyFD) isSocketAddressLegacy() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressLegacyFD) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "fd",
		"data": s,
	}
	return json.Marshal(v)
}

// SocketAddressLegacyInet is an implementation of SocketAddressLegacy.
type SocketAddressLegacyInet InetSocketAddress

func (SocketAddressLegacyInet) isSocketAddressLegacy() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressLegacyInet) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "inet",
		"data": s,
	}
	return json.Marshal(v)
}

// SocketAddressLegacyUnix is an implementation of SocketAddressLegacy.
type SocketAddressLegacyUnix UnixSocketAddress

func (SocketAddressLegacyUnix) isSocketAddressLegacy() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressLegacyUnix) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "unix",
		"data": s,
	}
	return json.Marshal(v)
}

// SocketAddressLegacyVsock is an implementation of SocketAddressLegacy.
type SocketAddressLegacyVsock VsockSocketAddress

func (SocketAddressLegacyVsock) isSocketAddressLegacy() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressLegacyVsock) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "vsock",
		"data": s,
	}
	return json.Marshal(v)
}

func decodeSocketAddressLegacy(bs json.RawMessage) (SocketAddressLegacy, error) {
	v := struct {
		T string          `json:"type"`
		V json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.T {
	case "fd":
		var ret SocketAddressLegacyFD
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "inet":
		var ret SocketAddressLegacyInet
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "unix":
		var ret SocketAddressLegacyUnix
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	case "vsock":
		var ret SocketAddressLegacyVsock
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("unknown subtype %q for union SocketAddressLegacy", v.T)
	}
}

// SocketAddressType -> SocketAddressType (enum)

// SocketAddressType implements the "SocketAddressType" QMP API type.
type SocketAddressType int

// Known values of SocketAddressType.
const (
	SocketAddressTypeInet SocketAddressType = iota
	SocketAddressTypeUnix
	SocketAddressTypeVsock
	SocketAddressTypeFD
)

// String implements fmt.Stringer.
func (e SocketAddressType) String() string {
	switch e {
	case SocketAddressTypeInet:
		return "inet"
	case SocketAddressTypeUnix:
		return "unix"
	case SocketAddressTypeVsock:
		return "vsock"
	case SocketAddressTypeFD:
		return "fd"
	default:
		return fmt.Sprintf("SocketAddressType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e SocketAddressType) MarshalJSON() ([]byte, error) {
	switch e {
	case SocketAddressTypeInet:
		return json.Marshal("inet")
	case SocketAddressTypeUnix:
		return json.Marshal("unix")
	case SocketAddressTypeVsock:
		return json.Marshal("vsock")
	case SocketAddressTypeFD:
		return json.Marshal("fd")
	default:
		return nil, fmt.Errorf("unknown enum value %q for SocketAddressType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *SocketAddressType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "inet":
		*e = SocketAddressTypeInet
	case "unix":
		*e = SocketAddressTypeUnix
	case "vsock":
		*e = SocketAddressTypeVsock
	case "fd":
		*e = SocketAddressTypeFD
	default:
		return fmt.Errorf("unknown enum value %q for SocketAddressType", s)
	}
	return nil
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

// StrOrNull -> StrOrNull (alternate)

// StrOrNull implements the "StrOrNull" QMP API type.
//
// Can be one of:
//   - StrOrNullN
//   - StrOrNullS
type StrOrNull interface {
	isStrOrNull()
}

// StrOrNullN is a JSON null type, so it must
// also implement the isNullable interface.
type StrOrNullN struct{}

func (StrOrNullN) isNull() bool { return true }
func (StrOrNullN) isStrOrNull() {}

// StrOrNullS is an implementation of StrOrNull
type StrOrNullS string

func (StrOrNullS) isStrOrNull() {}

func decodeStrOrNull(bs json.RawMessage) (StrOrNull, error) {

	// Always try unmarshalling for nil first if it's an option
	// because other types could unmarshal successfully in the case
	// where a Null json type was provided.
	var n *int
	if err := json.Unmarshal([]byte(bs), &n); err == nil {
		if n == nil {
			return StrOrNullN{}, nil
		}
	}
	var s StrOrNullS
	if err := json.Unmarshal([]byte(bs), &s); err == nil {
		return s, nil
	}
	return nil, fmt.Errorf("unable to decode %q as a StrOrNull", string(bs))
}

// String -> String (struct)

// String implements the "String" QMP API type.
type String struct {
	Str string `json:"str"`
}

// TPMEmulatorOptions -> TPMEmulatorOptions (struct)

// TPMEmulatorOptions implements the "TPMEmulatorOptions" QMP API type.
type TPMEmulatorOptions struct {
	Chardev string `json:"chardev"`
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
	TPMTypeEmulator
)

// String implements fmt.Stringer.
func (e TPMType) String() string {
	switch e {
	case TPMTypePassthrough:
		return "passthrough"
	case TPMTypeEmulator:
		return "emulator"
	default:
		return fmt.Sprintf("TPMType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e TPMType) MarshalJSON() ([]byte, error) {
	switch e {
	case TPMTypePassthrough:
		return json.Marshal("passthrough")
	case TPMTypeEmulator:
		return json.Marshal("emulator")
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
	case "emulator":
		*e = TPMTypeEmulator
	default:
		return fmt.Errorf("unknown enum value %q for TPMType", s)
	}
	return nil
}

// TpmTypeOptions -> TPMTypeOptions (simple union)

// TPMTypeOptions implements the "TpmTypeOptions" QMP API type.
//
// Can be one of:
//   - TPMTypeOptionsEmulator
//   - TPMTypeOptionsPassthrough
type TPMTypeOptions interface {
	isTPMTypeOptions()
}

// TPMTypeOptionsEmulator is an implementation of TPMTypeOptions.
type TPMTypeOptionsEmulator TPMEmulatorOptions

func (TPMTypeOptionsEmulator) isTPMTypeOptions() {}

// MarshalJSON implements json.Marshaler.
func (s TPMTypeOptionsEmulator) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"type": "emulator",
		"data": s,
	}
	return json.Marshal(v)
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
	case "emulator":
		var ret TPMTypeOptionsEmulator
		if err := json.Unmarshal([]byte(v.V), &ret); err != nil {
			return nil, err
		}
		return ret, nil
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
	Server   []VNCServerInfo2    `json:"server"`
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

// VncServerInfo2 -> VNCServerInfo2 (struct)

// VNCServerInfo2 implements the "VncServerInfo2" QMP API type.
type VNCServerInfo2 struct {
	Host      string               `json:"host"`
	Service   string               `json:"service"`
	Family    NetworkAddressFamily `json:"family"`
	Websocket bool                 `json:"websocket"`
	Auth      VNCPrimaryAuth       `json:"auth"`
	Vencrypt  *VNCVencryptSubAuth  `json:"vencrypt,omitempty"`
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

// VsockSocketAddress -> VsockSocketAddress (struct)

// VsockSocketAddress implements the "VsockSocketAddress" QMP API type.
type VsockSocketAddress struct {
	Cid  string `json:"cid"`
	Port string `json:"port"`
}

// EVENT WAKEUP

// EVENT WATCHDOG

// WatchdogAction -> WatchdogAction (enum)

// WatchdogAction implements the "WatchdogAction" QMP API type.
type WatchdogAction int

// Known values of WatchdogAction.
const (
	WatchdogActionReset WatchdogAction = iota
	WatchdogActionShutdown
	WatchdogActionPoweroff
	WatchdogActionPause
	WatchdogActionDebug
	WatchdogActionNone
	WatchdogActionInjectNmi
)

// String implements fmt.Stringer.
func (e WatchdogAction) String() string {
	switch e {
	case WatchdogActionReset:
		return "reset"
	case WatchdogActionShutdown:
		return "shutdown"
	case WatchdogActionPoweroff:
		return "poweroff"
	case WatchdogActionPause:
		return "pause"
	case WatchdogActionDebug:
		return "debug"
	case WatchdogActionNone:
		return "none"
	case WatchdogActionInjectNmi:
		return "inject-nmi"
	default:
		return fmt.Sprintf("WatchdogAction(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e WatchdogAction) MarshalJSON() ([]byte, error) {
	switch e {
	case WatchdogActionReset:
		return json.Marshal("reset")
	case WatchdogActionShutdown:
		return json.Marshal("shutdown")
	case WatchdogActionPoweroff:
		return json.Marshal("poweroff")
	case WatchdogActionPause:
		return json.Marshal("pause")
	case WatchdogActionDebug:
		return json.Marshal("debug")
	case WatchdogActionNone:
		return json.Marshal("none")
	case WatchdogActionInjectNmi:
		return json.Marshal("inject-nmi")
	default:
		return nil, fmt.Errorf("unknown enum value %q for WatchdogAction", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *WatchdogAction) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "reset":
		*e = WatchdogActionReset
	case "shutdown":
		*e = WatchdogActionShutdown
	case "poweroff":
		*e = WatchdogActionPoweroff
	case "pause":
		*e = WatchdogActionPause
	case "debug":
		*e = WatchdogActionDebug
	case "none":
		*e = WatchdogActionNone
	case "inject-nmi":
		*e = WatchdogActionInjectNmi
	default:
		return fmt.Errorf("unknown enum value %q for WatchdogAction", s)
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

// add-fd -> AddFD (command)

// AddFD implements the "add-fd" QMP API call.
func (m *Monitor) AddFD(fdsetID *int64, opaque *string) (ret AddfdInfo, err error) {
	cmd := struct {
		FdsetID *int64  `json:"fdset-id,omitempty"`
		Opaque  *string `json:"opaque,omitempty"`
	}{
		fdsetID,
		opaque,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "add-fd",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// add_client -> AddClient (command)

// AddClient implements the "add_client" QMP API call.
func (m *Monitor) AddClient(protocol string, fdname string, skipauth *bool, tls *bool) (err error) {
	cmd := struct {
		Protocol string `json:"protocol"`
		Fdname   string `json:"fdname"`
		Skipauth *bool  `json:"skipauth,omitempty"`
		TLS      *bool  `json:"tls,omitempty"`
	}{
		protocol,
		fdname,
		skipauth,
		tls,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "add_client",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// balloon -> Balloon (command)

// Balloon implements the "balloon" QMP API call.
func (m *Monitor) Balloon(value int64) (err error) {
	cmd := struct {
		Value int64 `json:"value"`
	}{
		value,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "balloon",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-commit -> BlockCommit (command)

// BlockCommit implements the "block-commit" QMP API call.
func (m *Monitor) BlockCommit(jobID *string, device string, base *string, top *string, backingFile *string, speed *int64, filterNodeName *string) (err error) {
	cmd := struct {
		JobID          *string `json:"job-id,omitempty"`
		Device         string  `json:"device"`
		Base           *string `json:"base,omitempty"`
		Top            *string `json:"top,omitempty"`
		BackingFile    *string `json:"backing-file,omitempty"`
		Speed          *int64  `json:"speed,omitempty"`
		FilterNodeName *string `json:"filter-node-name,omitempty"`
	}{
		jobID,
		device,
		base,
		top,
		backingFile,
		speed,
		filterNodeName,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-commit",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-dirty-bitmap-add -> BlockDirtyBitmapAdd (command)

// BlockDirtyBitmapAdd implements the "block-dirty-bitmap-add" QMP API call.
func (m *Monitor) BlockDirtyBitmapAdd(node string, name string, granularity *uint32, persistent *bool, autoload *bool) (err error) {
	cmd := struct {
		Node        string  `json:"node"`
		Name        string  `json:"name"`
		Granularity *uint32 `json:"granularity,omitempty"`
		Persistent  *bool   `json:"persistent,omitempty"`
		Autoload    *bool   `json:"autoload,omitempty"`
	}{
		node,
		name,
		granularity,
		persistent,
		autoload,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-dirty-bitmap-add",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-dirty-bitmap-clear -> BlockDirtyBitmapClear (command)

// BlockDirtyBitmapClear implements the "block-dirty-bitmap-clear" QMP API call.
func (m *Monitor) BlockDirtyBitmapClear(node string, name string) (err error) {
	cmd := struct {
		Node string `json:"node"`
		Name string `json:"name"`
	}{
		node,
		name,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-dirty-bitmap-clear",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-dirty-bitmap-remove -> BlockDirtyBitmapRemove (command)

// BlockDirtyBitmapRemove implements the "block-dirty-bitmap-remove" QMP API call.
func (m *Monitor) BlockDirtyBitmapRemove(node string, name string) (err error) {
	cmd := struct {
		Node string `json:"node"`
		Name string `json:"name"`
	}{
		node,
		name,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-dirty-bitmap-remove",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-job-cancel -> BlockJobCancel (command)

// BlockJobCancel implements the "block-job-cancel" QMP API call.
func (m *Monitor) BlockJobCancel(device string, force *bool) (err error) {
	cmd := struct {
		Device string `json:"device"`
		Force  *bool  `json:"force,omitempty"`
	}{
		device,
		force,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-job-cancel",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-job-complete -> BlockJobComplete (command)

// BlockJobComplete implements the "block-job-complete" QMP API call.
func (m *Monitor) BlockJobComplete(device string) (err error) {
	cmd := struct {
		Device string `json:"device"`
	}{
		device,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-job-complete",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-job-pause -> BlockJobPause (command)

// BlockJobPause implements the "block-job-pause" QMP API call.
func (m *Monitor) BlockJobPause(device string) (err error) {
	cmd := struct {
		Device string `json:"device"`
	}{
		device,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-job-pause",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-job-resume -> BlockJobResume (command)

// BlockJobResume implements the "block-job-resume" QMP API call.
func (m *Monitor) BlockJobResume(device string) (err error) {
	cmd := struct {
		Device string `json:"device"`
	}{
		device,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-job-resume",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-job-set-speed -> BlockJobSetSpeed (command)

// BlockJobSetSpeed implements the "block-job-set-speed" QMP API call.
func (m *Monitor) BlockJobSetSpeed(device string, speed int64) (err error) {
	cmd := struct {
		Device string `json:"device"`
		Speed  int64  `json:"speed"`
	}{
		device,
		speed,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-job-set-speed",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-set-write-threshold -> BlockSetWriteThreshold (command)

// BlockSetWriteThreshold implements the "block-set-write-threshold" QMP API call.
func (m *Monitor) BlockSetWriteThreshold(nodeName string, writeThreshold uint64) (err error) {
	cmd := struct {
		NodeName       string `json:"node-name"`
		WriteThreshold uint64 `json:"write-threshold"`
	}{
		nodeName,
		writeThreshold,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-set-write-threshold",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block-stream -> BlockStream (command)

// BlockStream implements the "block-stream" QMP API call.
func (m *Monitor) BlockStream(jobID *string, device string, base *string, baseNode *string, backingFile *string, speed *int64, onError *BlockdevOnError) (err error) {
	cmd := struct {
		JobID       *string          `json:"job-id,omitempty"`
		Device      string           `json:"device"`
		Base        *string          `json:"base,omitempty"`
		BaseNode    *string          `json:"base-node,omitempty"`
		BackingFile *string          `json:"backing-file,omitempty"`
		Speed       *int64           `json:"speed,omitempty"`
		OnError     *BlockdevOnError `json:"on-error,omitempty"`
	}{
		jobID,
		device,
		base,
		baseNode,
		backingFile,
		speed,
		onError,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-stream",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block_passwd -> BlockPasswd (command)

// BlockPasswd implements the "block_passwd" QMP API call.
func (m *Monitor) BlockPasswd(device *string, nodeName *string, password string) (err error) {
	cmd := struct {
		Device   *string `json:"device,omitempty"`
		NodeName *string `json:"node-name,omitempty"`
		Password string  `json:"password"`
	}{
		device,
		nodeName,
		password,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block_passwd",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block_resize -> BlockResize (command)

// BlockResize implements the "block_resize" QMP API call.
func (m *Monitor) BlockResize(device *string, nodeName *string, size int64) (err error) {
	cmd := struct {
		Device   *string `json:"device,omitempty"`
		NodeName *string `json:"node-name,omitempty"`
		Size     int64   `json:"size"`
	}{
		device,
		nodeName,
		size,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block_resize",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// block_set_io_throttle -> BlockSetIOThrottle (command)

// BlockSetIOThrottle implements the "block_set_io_throttle" QMP API call.
func (m *Monitor) BlockSetIOThrottle(cmd *BlockIOThrottle) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block_set_io_throttle",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// blockdev-add -> BlockdevAdd (command)

// BlockdevAdd implements the "blockdev-add" QMP API call.
func (m *Monitor) BlockdevAdd(cmd *BlockdevOptions) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-add",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// blockdev-backup -> BlockdevBackup (command)

// BlockdevBackup implements the "blockdev-backup" QMP API call.
func (m *Monitor) BlockdevBackup(cmd *BlockdevBackup) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-backup",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// blockdev-change-medium -> BlockdevChangeMedium (command)

// BlockdevChangeMedium implements the "blockdev-change-medium" QMP API call.
func (m *Monitor) BlockdevChangeMedium(device *string, id *string, filename string, format *string, readOnlyMode *BlockdevChangeReadOnlyMode) (err error) {
	cmd := struct {
		Device       *string                     `json:"device,omitempty"`
		ID           *string                     `json:"id,omitempty"`
		Filename     string                      `json:"filename"`
		Format       *string                     `json:"format,omitempty"`
		ReadOnlyMode *BlockdevChangeReadOnlyMode `json:"read-only-mode,omitempty"`
	}{
		device,
		id,
		filename,
		format,
		readOnlyMode,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-change-medium",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// blockdev-close-tray -> BlockdevCloseTray (command)

// BlockdevCloseTray implements the "blockdev-close-tray" QMP API call.
func (m *Monitor) BlockdevCloseTray(device *string, id *string) (err error) {
	cmd := struct {
		Device *string `json:"device,omitempty"`
		ID     *string `json:"id,omitempty"`
	}{
		device,
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-close-tray",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// blockdev-del -> BlockdevDel (command)

// BlockdevDel implements the "blockdev-del" QMP API call.
func (m *Monitor) BlockdevDel(nodeName string) (err error) {
	cmd := struct {
		NodeName string `json:"node-name"`
	}{
		nodeName,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-del",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// blockdev-mirror -> BlockdevMirror (command)

// BlockdevMirror implements the "blockdev-mirror" QMP API call.
func (m *Monitor) BlockdevMirror(jobID *string, device string, target string, replaces *string, sync MirrorSyncMode, speed *int64, granularity *uint32, bufSize *int64, onSourceError *BlockdevOnError, onTargetError *BlockdevOnError, filterNodeName *string) (err error) {
	cmd := struct {
		JobID          *string          `json:"job-id,omitempty"`
		Device         string           `json:"device"`
		Target         string           `json:"target"`
		Replaces       *string          `json:"replaces,omitempty"`
		Sync           MirrorSyncMode   `json:"sync"`
		Speed          *int64           `json:"speed,omitempty"`
		Granularity    *uint32          `json:"granularity,omitempty"`
		BufSize        *int64           `json:"buf-size,omitempty"`
		OnSourceError  *BlockdevOnError `json:"on-source-error,omitempty"`
		OnTargetError  *BlockdevOnError `json:"on-target-error,omitempty"`
		FilterNodeName *string          `json:"filter-node-name,omitempty"`
	}{
		jobID,
		device,
		target,
		replaces,
		sync,
		speed,
		granularity,
		bufSize,
		onSourceError,
		onTargetError,
		filterNodeName,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-mirror",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// blockdev-open-tray -> BlockdevOpenTray (command)

// BlockdevOpenTray implements the "blockdev-open-tray" QMP API call.
func (m *Monitor) BlockdevOpenTray(device *string, id *string, force *bool) (err error) {
	cmd := struct {
		Device *string `json:"device,omitempty"`
		ID     *string `json:"id,omitempty"`
		Force  *bool   `json:"force,omitempty"`
	}{
		device,
		id,
		force,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-open-tray",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// blockdev-snapshot -> BlockdevSnapshot (command)

// BlockdevSnapshot implements the "blockdev-snapshot" QMP API call.
func (m *Monitor) BlockdevSnapshot(node string, overlay string) (err error) {
	cmd := struct {
		Node    string `json:"node"`
		Overlay string `json:"overlay"`
	}{
		node,
		overlay,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-snapshot",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// blockdev-snapshot-delete-internal-sync -> BlockdevSnapshotDeleteInternalSync (command)

// BlockdevSnapshotDeleteInternalSync implements the "blockdev-snapshot-delete-internal-sync" QMP API call.
func (m *Monitor) BlockdevSnapshotDeleteInternalSync(device string, id *string, name *string) (ret SnapshotInfo, err error) {
	cmd := struct {
		Device string  `json:"device"`
		ID     *string `json:"id,omitempty"`
		Name   *string `json:"name,omitempty"`
	}{
		device,
		id,
		name,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-snapshot-delete-internal-sync",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// blockdev-snapshot-internal-sync -> BlockdevSnapshotInternalSync (command)

// BlockdevSnapshotInternalSync implements the "blockdev-snapshot-internal-sync" QMP API call.
func (m *Monitor) BlockdevSnapshotInternalSync(device string, name string) (err error) {
	cmd := struct {
		Device string `json:"device"`
		Name   string `json:"name"`
	}{
		device,
		name,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-snapshot-internal-sync",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// blockdev-snapshot-sync -> BlockdevSnapshotSync (command)

// BlockdevSnapshotSync implements the "blockdev-snapshot-sync" QMP API call.
func (m *Monitor) BlockdevSnapshotSync(device *string, nodeName *string, snapshotFile string, snapshotNodeName *string, format *string, mode *NewImageMode) (err error) {
	cmd := struct {
		Device           *string       `json:"device,omitempty"`
		NodeName         *string       `json:"node-name,omitempty"`
		SnapshotFile     string        `json:"snapshot-file"`
		SnapshotNodeName *string       `json:"snapshot-node-name,omitempty"`
		Format           *string       `json:"format,omitempty"`
		Mode             *NewImageMode `json:"mode,omitempty"`
	}{
		device,
		nodeName,
		snapshotFile,
		snapshotNodeName,
		format,
		mode,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-snapshot-sync",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// change -> Change (command)

// Change implements the "change" QMP API call.
func (m *Monitor) Change(device string, target string, arg *string) (err error) {
	cmd := struct {
		Device string  `json:"device"`
		Target string  `json:"target"`
		Arg    *string `json:"arg,omitempty"`
	}{
		device,
		target,
		arg,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "change",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// change-backing-file -> ChangeBackingFile (command)

// ChangeBackingFile implements the "change-backing-file" QMP API call.
func (m *Monitor) ChangeBackingFile(device string, imageNodeName string, backingFile string) (err error) {
	cmd := struct {
		Device        string `json:"device"`
		ImageNodeName string `json:"image-node-name"`
		BackingFile   string `json:"backing-file"`
	}{
		device,
		imageNodeName,
		backingFile,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "change-backing-file",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// change-vnc-password -> ChangeVNCPassword (command)

// ChangeVNCPassword implements the "change-vnc-password" QMP API call.
func (m *Monitor) ChangeVNCPassword(password string) (err error) {
	cmd := struct {
		Password string `json:"password"`
	}{
		password,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "change-vnc-password",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// chardev-add -> ChardevAdd (command)

// ChardevAdd implements the "chardev-add" QMP API call.
func (m *Monitor) ChardevAdd(id string, backend ChardevBackend) (ret ChardevReturn, err error) {
	cmd := struct {
		ID      string         `json:"id"`
		Backend ChardevBackend `json:"backend"`
	}{
		id,
		backend,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "chardev-add",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// chardev-change -> ChardevChange (command)

// ChardevChange implements the "chardev-change" QMP API call.
func (m *Monitor) ChardevChange(id string, backend ChardevBackend) (ret ChardevReturn, err error) {
	cmd := struct {
		ID      string         `json:"id"`
		Backend ChardevBackend `json:"backend"`
	}{
		id,
		backend,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "chardev-change",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// chardev-remove -> ChardevRemove (command)

// ChardevRemove implements the "chardev-remove" QMP API call.
func (m *Monitor) ChardevRemove(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "chardev-remove",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// chardev-send-break -> ChardevSendBreak (command)

// ChardevSendBreak implements the "chardev-send-break" QMP API call.
func (m *Monitor) ChardevSendBreak(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "chardev-send-break",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// client_migrate_info -> ClientMigrateInfo (command)

// ClientMigrateInfo implements the "client_migrate_info" QMP API call.
func (m *Monitor) ClientMigrateInfo(protocol string, hostname string, port *int64, tlsPort *int64, certSubject *string) (err error) {
	cmd := struct {
		Protocol    string  `json:"protocol"`
		Hostname    string  `json:"hostname"`
		Port        *int64  `json:"port,omitempty"`
		TLSPort     *int64  `json:"tls-port,omitempty"`
		CertSubject *string `json:"cert-subject,omitempty"`
	}{
		protocol,
		hostname,
		port,
		tlsPort,
		certSubject,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "client_migrate_info",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// closefd -> Closefd (command)

// Closefd implements the "closefd" QMP API call.
func (m *Monitor) Closefd(fdname string) (err error) {
	cmd := struct {
		Fdname string `json:"fdname"`
	}{
		fdname,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "closefd",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// cont -> Cont (command)

// Cont implements the "cont" QMP API call.
func (m *Monitor) Cont() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "cont",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// cpu -> CPU (command)

// CPU implements the "cpu" QMP API call.
func (m *Monitor) CPU(index int64) (err error) {
	cmd := struct {
		Index int64 `json:"index"`
	}{
		index,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "cpu",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// cpu-add -> CPUAdd (command)

// CPUAdd implements the "cpu-add" QMP API call.
func (m *Monitor) CPUAdd(id int64) (err error) {
	cmd := struct {
		ID int64 `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "cpu-add",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// device-list-properties -> DeviceListProperties (command)

// DeviceListProperties implements the "device-list-properties" QMP API call.
func (m *Monitor) DeviceListProperties(typename string) (ret []DevicePropertyInfo, err error) {
	cmd := struct {
		Typename string `json:"typename"`
	}{
		typename,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "device-list-properties",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// device_add -> DeviceAdd (command)

// DeviceAdd implements the "device_add" QMP API call.
func (m *Monitor) DeviceAdd(driver string, bus *string, id *string) (err error) {
	cmd := struct {
		Driver string  `json:"driver"`
		Bus    *string `json:"bus,omitempty"`
		ID     *string `json:"id,omitempty"`
	}{
		driver,
		bus,
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "device_add",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// device_del -> DeviceDel (command)

// DeviceDel implements the "device_del" QMP API call.
func (m *Monitor) DeviceDel(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "device_del",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// drive-backup -> DriveBackup (command)

// DriveBackup implements the "drive-backup" QMP API call.
func (m *Monitor) DriveBackup(cmd *DriveBackup) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "drive-backup",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// drive-mirror -> DriveMirror (command)

// DriveMirror implements the "drive-mirror" QMP API call.
func (m *Monitor) DriveMirror(cmd *DriveMirror) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "drive-mirror",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// dump-guest-memory -> DumpGuestMemory (command)

// DumpGuestMemory implements the "dump-guest-memory" QMP API call.
func (m *Monitor) DumpGuestMemory(paging bool, protocol string, detach *bool, begin *int64, length *int64, format *DumpGuestMemoryFormat) (err error) {
	cmd := struct {
		Paging   bool                   `json:"paging"`
		Protocol string                 `json:"protocol"`
		Detach   *bool                  `json:"detach,omitempty"`
		Begin    *int64                 `json:"begin,omitempty"`
		Length   *int64                 `json:"length,omitempty"`
		Format   *DumpGuestMemoryFormat `json:"format,omitempty"`
	}{
		paging,
		protocol,
		detach,
		begin,
		length,
		format,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "dump-guest-memory",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// dump-skeys -> DumpSkeys (command)

// DumpSkeys implements the "dump-skeys" QMP API call.
func (m *Monitor) DumpSkeys(filename string) (err error) {
	cmd := struct {
		Filename string `json:"filename"`
	}{
		filename,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "dump-skeys",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// eject -> Eject (command)

// Eject implements the "eject" QMP API call.
func (m *Monitor) Eject(device *string, id *string, force *bool) (err error) {
	cmd := struct {
		Device *string `json:"device,omitempty"`
		ID     *string `json:"id,omitempty"`
		Force  *bool   `json:"force,omitempty"`
	}{
		device,
		id,
		force,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "eject",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// expire_password -> ExpirePassword (command)

// ExpirePassword implements the "expire_password" QMP API call.
func (m *Monitor) ExpirePassword(protocol string, time string) (err error) {
	cmd := struct {
		Protocol string `json:"protocol"`
		Time     string `json:"time"`
	}{
		protocol,
		time,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "expire_password",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// getfd -> Getfd (command)

// Getfd implements the "getfd" QMP API call.
func (m *Monitor) Getfd(fdname string) (err error) {
	cmd := struct {
		Fdname string `json:"fdname"`
	}{
		fdname,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "getfd",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// human-monitor-command -> HumanMonitorCommand (command)

// HumanMonitorCommand implements the "human-monitor-command" QMP API call.
func (m *Monitor) HumanMonitorCommand(commandLine string, cpuIndex *int64) (ret string, err error) {
	cmd := struct {
		CommandLine string `json:"command-line"`
		CPUIndex    *int64 `json:"cpu-index,omitempty"`
	}{
		commandLine,
		cpuIndex,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "human-monitor-command",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// inject-nmi -> InjectNmi (command)

// InjectNmi implements the "inject-nmi" QMP API call.
func (m *Monitor) InjectNmi() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "inject-nmi",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// input-send-event -> InputSendEvent (command)

// InputSendEvent implements the "input-send-event" QMP API call.
func (m *Monitor) InputSendEvent(device *string, head *int64, events []InputEvent) (err error) {
	cmd := struct {
		Device *string      `json:"device,omitempty"`
		Head   *int64       `json:"head,omitempty"`
		Events []InputEvent `json:"events"`
	}{
		device,
		head,
		events,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "input-send-event",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// memsave -> Memsave (command)

// Memsave implements the "memsave" QMP API call.
func (m *Monitor) Memsave(val int64, size int64, filename string, cpuIndex *int64) (err error) {
	cmd := struct {
		Val      int64  `json:"val"`
		Size     int64  `json:"size"`
		Filename string `json:"filename"`
		CPUIndex *int64 `json:"cpu-index,omitempty"`
	}{
		val,
		size,
		filename,
		cpuIndex,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "memsave",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// migrate -> Migrate (command)

// Migrate implements the "migrate" QMP API call.
func (m *Monitor) Migrate(uri string, blk *bool, inc *bool, detach *bool) (err error) {
	cmd := struct {
		URI    string `json:"uri"`
		Blk    *bool  `json:"blk,omitempty"`
		Inc    *bool  `json:"inc,omitempty"`
		Detach *bool  `json:"detach,omitempty"`
	}{
		uri,
		blk,
		inc,
		detach,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// migrate-continue -> MigrateContinue (command)

// MigrateContinue implements the "migrate-continue" QMP API call.
func (m *Monitor) MigrateContinue(state MigrationStatus) (err error) {
	cmd := struct {
		State MigrationStatus `json:"state"`
	}{
		state,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate-continue",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// migrate-incoming -> MigrateIncoming (command)

// MigrateIncoming implements the "migrate-incoming" QMP API call.
func (m *Monitor) MigrateIncoming(uri string) (err error) {
	cmd := struct {
		URI string `json:"uri"`
	}{
		uri,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate-incoming",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// migrate-set-cache-size -> MigrateSetCacheSize (command)

// MigrateSetCacheSize implements the "migrate-set-cache-size" QMP API call.
func (m *Monitor) MigrateSetCacheSize(value int64) (err error) {
	cmd := struct {
		Value int64 `json:"value"`
	}{
		value,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate-set-cache-size",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// migrate-set-capabilities -> MigrateSetCapabilities (command)

// MigrateSetCapabilities implements the "migrate-set-capabilities" QMP API call.
func (m *Monitor) MigrateSetCapabilities(capabilities []MigrationCapabilityStatus) (err error) {
	cmd := struct {
		Capabilities []MigrationCapabilityStatus `json:"capabilities"`
	}{
		capabilities,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate-set-capabilities",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// migrate-set-parameters -> MigrateSetParameters (command)

// MigrateSetParameters implements the "migrate-set-parameters" QMP API call.
func (m *Monitor) MigrateSetParameters(cmd *MigrateSetParameters) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate-set-parameters",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// migrate-start-postcopy -> MigrateStartPostcopy (command)

// MigrateStartPostcopy implements the "migrate-start-postcopy" QMP API call.
func (m *Monitor) MigrateStartPostcopy() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate-start-postcopy",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// migrate_cancel -> MigrateCancel (command)

// MigrateCancel implements the "migrate_cancel" QMP API call.
func (m *Monitor) MigrateCancel() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate_cancel",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// migrate_set_downtime -> MigrateSetDowntime (command)

// MigrateSetDowntime implements the "migrate_set_downtime" QMP API call.
func (m *Monitor) MigrateSetDowntime(value float64) (err error) {
	cmd := struct {
		Value float64 `json:"value"`
	}{
		value,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate_set_downtime",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// migrate_set_speed -> MigrateSetSpeed (command)

// MigrateSetSpeed implements the "migrate_set_speed" QMP API call.
func (m *Monitor) MigrateSetSpeed(value int64) (err error) {
	cmd := struct {
		Value int64 `json:"value"`
	}{
		value,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate_set_speed",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// nbd-server-add -> NBDServerAdd (command)

// NBDServerAdd implements the "nbd-server-add" QMP API call.
func (m *Monitor) NBDServerAdd(device string, writable *bool) (err error) {
	cmd := struct {
		Device   string `json:"device"`
		Writable *bool  `json:"writable,omitempty"`
	}{
		device,
		writable,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "nbd-server-add",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// nbd-server-start -> NBDServerStart (command)

// NBDServerStart implements the "nbd-server-start" QMP API call.
func (m *Monitor) NBDServerStart(addr SocketAddressLegacy, tlsCreds *string) (err error) {
	cmd := struct {
		Addr     SocketAddressLegacy `json:"addr"`
		TLSCreds *string             `json:"tls-creds,omitempty"`
	}{
		addr,
		tlsCreds,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "nbd-server-start",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// nbd-server-stop -> NBDServerStop (command)

// NBDServerStop implements the "nbd-server-stop" QMP API call.
func (m *Monitor) NBDServerStop() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "nbd-server-stop",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// netdev_add -> NetdevAdd (command)

// NetdevAdd implements the "netdev_add" QMP API call.
func (m *Monitor) NetdevAdd(typ string, id string) (err error) {
	cmd := struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}{
		typ,
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "netdev_add",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// netdev_del -> NetdevDel (command)

// NetdevDel implements the "netdev_del" QMP API call.
func (m *Monitor) NetdevDel(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "netdev_del",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// object-add -> ObjectAdd (command)

// ObjectAdd implements the "object-add" QMP API call.
func (m *Monitor) ObjectAdd(qomType string, id string, props *interface{}) (err error) {
	cmd := struct {
		QomType string       `json:"qom-type"`
		ID      string       `json:"id"`
		Props   *interface{} `json:"props,omitempty"`
	}{
		qomType,
		id,
		props,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "object-add",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// object-del -> ObjectDel (command)

// ObjectDel implements the "object-del" QMP API call.
func (m *Monitor) ObjectDel(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "object-del",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// pmemsave -> Pmemsave (command)

// Pmemsave implements the "pmemsave" QMP API call.
func (m *Monitor) Pmemsave(val int64, size int64, filename string) (err error) {
	cmd := struct {
		Val      int64  `json:"val"`
		Size     int64  `json:"size"`
		Filename string `json:"filename"`
	}{
		val,
		size,
		filename,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "pmemsave",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// qmp_capabilities -> QMPCapabilities (command)

// QMPCapabilities implements the "qmp_capabilities" QMP API call.
func (m *Monitor) QMPCapabilities() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "qmp_capabilities",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// qom-get -> QomGet (command)

// QomGet implements the "qom-get" QMP API call.
func (m *Monitor) QomGet(path string, property string) (ret interface{}, err error) {
	cmd := struct {
		Path     string `json:"path"`
		Property string `json:"property"`
	}{
		path,
		property,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "qom-get",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// qom-list -> QomList (command)

// QomList implements the "qom-list" QMP API call.
func (m *Monitor) QomList(path string) (ret []ObjectPropertyInfo, err error) {
	cmd := struct {
		Path string `json:"path"`
	}{
		path,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "qom-list",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// qom-list-types -> QomListTypes (command)

// QomListTypes implements the "qom-list-types" QMP API call.
func (m *Monitor) QomListTypes(implements *string, abstract *bool) (ret []ObjectTypeInfo, err error) {
	cmd := struct {
		Implements *string `json:"implements,omitempty"`
		Abstract   *bool   `json:"abstract,omitempty"`
	}{
		implements,
		abstract,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "qom-list-types",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// qom-set -> QomSet (command)

// QomSet implements the "qom-set" QMP API call.
func (m *Monitor) QomSet(path string, property string, value interface{}) (err error) {
	cmd := struct {
		Path     string      `json:"path"`
		Property string      `json:"property"`
		Value    interface{} `json:"value"`
	}{
		path,
		property,
		value,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "qom-set",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// query-acpi-ospm-status -> QueryACPIOspmStatus (command)

// QueryACPIOspmStatus implements the "query-acpi-ospm-status" QMP API call.
func (m *Monitor) QueryACPIOspmStatus() (ret []ACPIOSTInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-acpi-ospm-status",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-balloon -> QueryBalloon (command)

// QueryBalloon implements the "query-balloon" QMP API call.
func (m *Monitor) QueryBalloon() (ret BalloonInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-balloon",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-block -> QueryBlock (command)

// QueryBlock implements the "query-block" QMP API call.
func (m *Monitor) QueryBlock() (ret []BlockInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-block",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-block-jobs -> QueryBlockJobs (command)

// QueryBlockJobs implements the "query-block-jobs" QMP API call.
func (m *Monitor) QueryBlockJobs() (ret []BlockJobInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-block-jobs",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-blockstats -> QueryBlockstats (command)

// QueryBlockstats implements the "query-blockstats" QMP API call.
func (m *Monitor) QueryBlockstats(queryNodes *bool) (ret []BlockStats, err error) {
	cmd := struct {
		QueryNodes *bool `json:"query-nodes,omitempty"`
	}{
		queryNodes,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-blockstats",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-chardev -> QueryChardev (command)

// QueryChardev implements the "query-chardev" QMP API call.
func (m *Monitor) QueryChardev() (ret []ChardevInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-chardev",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-chardev-backends -> QueryChardevBackends (command)

// QueryChardevBackends implements the "query-chardev-backends" QMP API call.
func (m *Monitor) QueryChardevBackends() (ret []ChardevBackendInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-chardev-backends",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-command-line-options -> QueryCommandLineOptions (command)

// QueryCommandLineOptions implements the "query-command-line-options" QMP API call.
func (m *Monitor) QueryCommandLineOptions(option *string) (ret []CommandLineOptionInfo, err error) {
	cmd := struct {
		Option *string `json:"option,omitempty"`
	}{
		option,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-command-line-options",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-commands -> QueryCommands (command)

// QueryCommands implements the "query-commands" QMP API call.
func (m *Monitor) QueryCommands() (ret []CommandInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-commands",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-cpu-definitions -> QueryCPUDefinitions (command)

// QueryCPUDefinitions implements the "query-cpu-definitions" QMP API call.
func (m *Monitor) QueryCPUDefinitions() (ret []CPUDefinitionInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-cpu-definitions",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-cpu-model-baseline -> QueryCPUModelBaseline (command)

// QueryCPUModelBaseline implements the "query-cpu-model-baseline" QMP API call.
func (m *Monitor) QueryCPUModelBaseline(modela CPUModelInfo, modelb CPUModelInfo) (ret CPUModelBaselineInfo, err error) {
	cmd := struct {
		Modela CPUModelInfo `json:"modela"`
		Modelb CPUModelInfo `json:"modelb"`
	}{
		modela,
		modelb,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-cpu-model-baseline",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-cpu-model-comparison -> QueryCPUModelComparison (command)

// QueryCPUModelComparison implements the "query-cpu-model-comparison" QMP API call.
func (m *Monitor) QueryCPUModelComparison(modela CPUModelInfo, modelb CPUModelInfo) (ret CPUModelCompareInfo, err error) {
	cmd := struct {
		Modela CPUModelInfo `json:"modela"`
		Modelb CPUModelInfo `json:"modelb"`
	}{
		modela,
		modelb,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-cpu-model-comparison",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-cpu-model-expansion -> QueryCPUModelExpansion (command)

// QueryCPUModelExpansion implements the "query-cpu-model-expansion" QMP API call.
func (m *Monitor) QueryCPUModelExpansion(typ CPUModelExpansionType, model CPUModelInfo) (ret CPUModelExpansionInfo, err error) {
	cmd := struct {
		Type  CPUModelExpansionType `json:"type"`
		Model CPUModelInfo          `json:"model"`
	}{
		typ,
		model,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-cpu-model-expansion",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-cpus -> QueryCpus (command)

// QueryCpus implements the "query-cpus" QMP API call.
func (m *Monitor) QueryCpus() (ret []CPUInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-cpus",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	var reslist []json.RawMessage
	if err = json.Unmarshal([]byte(res.Res), &reslist); err != nil {
		return
	}
	for _, r := range reslist {
		v, err := decodeCPUInfo(r)
		if err != nil {
			return nil, err
		}
		ret = append(ret, v)
	}
	return
}

// query-dump -> QueryDump (command)

// QueryDump implements the "query-dump" QMP API call.
func (m *Monitor) QueryDump() (ret DumpQueryResult, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-dump",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-dump-guest-memory-capability -> QueryDumpGuestMemoryCapability (command)

// QueryDumpGuestMemoryCapability implements the "query-dump-guest-memory-capability" QMP API call.
func (m *Monitor) QueryDumpGuestMemoryCapability() (ret DumpGuestMemoryCapability, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-dump-guest-memory-capability",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-events -> QueryEvents (command)

// QueryEvents implements the "query-events" QMP API call.
func (m *Monitor) QueryEvents() (ret []EventInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-events",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-fdsets -> QueryFdsets (command)

// QueryFdsets implements the "query-fdsets" QMP API call.
func (m *Monitor) QueryFdsets() (ret []FdsetInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-fdsets",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-gic-capabilities -> QueryGicCapabilities (command)

// QueryGicCapabilities implements the "query-gic-capabilities" QMP API call.
func (m *Monitor) QueryGicCapabilities() (ret []GicCapability, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-gic-capabilities",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-hotpluggable-cpus -> QueryHotpluggableCpus (command)

// QueryHotpluggableCpus implements the "query-hotpluggable-cpus" QMP API call.
func (m *Monitor) QueryHotpluggableCpus() (ret []HotpluggableCPU, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-hotpluggable-cpus",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-iothreads -> QueryIothreads (command)

// QueryIothreads implements the "query-iothreads" QMP API call.
func (m *Monitor) QueryIothreads() (ret []IOThreadInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-iothreads",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-kvm -> QueryKVM (command)

// QueryKVM implements the "query-kvm" QMP API call.
func (m *Monitor) QueryKVM() (ret KVMInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-kvm",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-machines -> QueryMachines (command)

// QueryMachines implements the "query-machines" QMP API call.
func (m *Monitor) QueryMachines() (ret []MachineInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-machines",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-memdev -> QueryMemdev (command)

// QueryMemdev implements the "query-memdev" QMP API call.
func (m *Monitor) QueryMemdev() (ret []Memdev, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-memdev",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-memory-devices -> QueryMemoryDevices (command)

// QueryMemoryDevices implements the "query-memory-devices" QMP API call.
func (m *Monitor) QueryMemoryDevices() (ret []MemoryDeviceInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-memory-devices",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	var reslist []json.RawMessage
	if err = json.Unmarshal([]byte(res.Res), &reslist); err != nil {
		return
	}
	for _, r := range reslist {
		v, err := decodeMemoryDeviceInfo(r)
		if err != nil {
			return nil, err
		}
		ret = append(ret, v)
	}
	return
}

// query-memory-size-summary -> QueryMemorySizeSummary (command)

// QueryMemorySizeSummary implements the "query-memory-size-summary" QMP API call.
func (m *Monitor) QueryMemorySizeSummary() (ret MemoryInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-memory-size-summary",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-mice -> QueryMice (command)

// QueryMice implements the "query-mice" QMP API call.
func (m *Monitor) QueryMice() (ret []MouseInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-mice",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-migrate -> QueryMigrate (command)

// QueryMigrate implements the "query-migrate" QMP API call.
func (m *Monitor) QueryMigrate() (ret MigrationInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-migrate",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-migrate-cache-size -> QueryMigrateCacheSize (command)

// QueryMigrateCacheSize implements the "query-migrate-cache-size" QMP API call.
func (m *Monitor) QueryMigrateCacheSize() (ret int64, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-migrate-cache-size",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-migrate-capabilities -> QueryMigrateCapabilities (command)

// QueryMigrateCapabilities implements the "query-migrate-capabilities" QMP API call.
func (m *Monitor) QueryMigrateCapabilities() (ret []MigrationCapabilityStatus, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-migrate-capabilities",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-migrate-parameters -> QueryMigrateParameters (command)

// QueryMigrateParameters implements the "query-migrate-parameters" QMP API call.
func (m *Monitor) QueryMigrateParameters() (ret MigrationParameters, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-migrate-parameters",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-name -> QueryName (command)

// QueryName implements the "query-name" QMP API call.
func (m *Monitor) QueryName() (ret NameInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-name",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-named-block-nodes -> QueryNamedBlockNodes (command)

// QueryNamedBlockNodes implements the "query-named-block-nodes" QMP API call.
func (m *Monitor) QueryNamedBlockNodes() (ret []BlockDeviceInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-named-block-nodes",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-pci -> QueryPCI (command)

// QueryPCI implements the "query-pci" QMP API call.
func (m *Monitor) QueryPCI() (ret []PCIInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-pci",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-qmp-schema -> QueryQMPSchema (command)

// QueryQMPSchema implements the "query-qmp-schema" QMP API call.
func (m *Monitor) QueryQMPSchema() (ret []SchemaInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-qmp-schema",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	var reslist []json.RawMessage
	if err = json.Unmarshal([]byte(res.Res), &reslist); err != nil {
		return
	}
	for _, r := range reslist {
		v, err := decodeSchemaInfo(r)
		if err != nil {
			return nil, err
		}
		ret = append(ret, v)
	}
	return
}

// query-rocker -> QueryRocker (command)

// QueryRocker implements the "query-rocker" QMP API call.
func (m *Monitor) QueryRocker(name string) (ret RockerSwitch, err error) {
	cmd := struct {
		Name string `json:"name"`
	}{
		name,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-rocker",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-rocker-of-dpa-flows -> QueryRockerOfDpaFlows (command)

// QueryRockerOfDpaFlows implements the "query-rocker-of-dpa-flows" QMP API call.
func (m *Monitor) QueryRockerOfDpaFlows(name string, tblID *uint32) (ret []RockerOfDpaFlow, err error) {
	cmd := struct {
		Name  string  `json:"name"`
		TblID *uint32 `json:"tbl-id,omitempty"`
	}{
		name,
		tblID,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-rocker-of-dpa-flows",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-rocker-of-dpa-groups -> QueryRockerOfDpaGroups (command)

// QueryRockerOfDpaGroups implements the "query-rocker-of-dpa-groups" QMP API call.
func (m *Monitor) QueryRockerOfDpaGroups(name string, typ *uint8) (ret []RockerOfDpaGroup, err error) {
	cmd := struct {
		Name string `json:"name"`
		Type *uint8 `json:"type,omitempty"`
	}{
		name,
		typ,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-rocker-of-dpa-groups",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-rocker-ports -> QueryRockerPorts (command)

// QueryRockerPorts implements the "query-rocker-ports" QMP API call.
func (m *Monitor) QueryRockerPorts(name string) (ret []RockerPort, err error) {
	cmd := struct {
		Name string `json:"name"`
	}{
		name,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-rocker-ports",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-rx-filter -> QueryRxFilter (command)

// QueryRxFilter implements the "query-rx-filter" QMP API call.
func (m *Monitor) QueryRxFilter(name *string) (ret []RxFilterInfo, err error) {
	cmd := struct {
		Name *string `json:"name,omitempty"`
	}{
		name,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-rx-filter",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-spice -> QuerySpice (command)

// QuerySpice implements the "query-spice" QMP API call.
func (m *Monitor) QuerySpice() (ret SpiceInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-spice",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-status -> QueryStatus (command)

// QueryStatus implements the "query-status" QMP API call.
func (m *Monitor) QueryStatus() (ret StatusInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-status",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-target -> QueryTarget (command)

// QueryTarget implements the "query-target" QMP API call.
func (m *Monitor) QueryTarget() (ret TargetInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-target",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-tpm -> QueryTPM (command)

// QueryTPM implements the "query-tpm" QMP API call.
func (m *Monitor) QueryTPM() (ret []TPMInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-tpm",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-tpm-models -> QueryTPMModels (command)

// QueryTPMModels implements the "query-tpm-models" QMP API call.
func (m *Monitor) QueryTPMModels() (ret []TPMModel, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-tpm-models",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-tpm-types -> QueryTPMTypes (command)

// QueryTPMTypes implements the "query-tpm-types" QMP API call.
func (m *Monitor) QueryTPMTypes() (ret []TPMType, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-tpm-types",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-uuid -> QueryUUID (command)

// QueryUUID implements the "query-uuid" QMP API call.
func (m *Monitor) QueryUUID() (ret UUIDInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-uuid",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-version -> QueryVersion (command)

// QueryVersion implements the "query-version" QMP API call.
func (m *Monitor) QueryVersion() (ret VersionInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-version",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-vm-generation-id -> QueryVMGenerationID (command)

// QueryVMGenerationID implements the "query-vm-generation-id" QMP API call.
func (m *Monitor) QueryVMGenerationID() (ret GUIDInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-vm-generation-id",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-vnc -> QueryVNC (command)

// QueryVNC implements the "query-vnc" QMP API call.
func (m *Monitor) QueryVNC() (ret VNCInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-vnc",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-vnc-servers -> QueryVNCServers (command)

// QueryVNCServers implements the "query-vnc-servers" QMP API call.
func (m *Monitor) QueryVNCServers() (ret []VNCInfo2, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-vnc-servers",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-xen-replication-status -> QueryXenReplicationStatus (command)

// QueryXenReplicationStatus implements the "query-xen-replication-status" QMP API call.
func (m *Monitor) QueryXenReplicationStatus() (ret ReplicationStatus, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-xen-replication-status",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// quit -> Quit (command)

// Quit implements the "quit" QMP API call.
func (m *Monitor) Quit() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "quit",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// remove-fd -> RemoveFD (command)

// RemoveFD implements the "remove-fd" QMP API call.
func (m *Monitor) RemoveFD(fdsetID int64, fd *int64) (err error) {
	cmd := struct {
		FdsetID int64  `json:"fdset-id"`
		FD      *int64 `json:"fd,omitempty"`
	}{
		fdsetID,
		fd,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "remove-fd",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// ringbuf-read -> RingbufRead (command)

// RingbufRead implements the "ringbuf-read" QMP API call.
func (m *Monitor) RingbufRead(device string, size int64, format *DataFormat) (ret string, err error) {
	cmd := struct {
		Device string      `json:"device"`
		Size   int64       `json:"size"`
		Format *DataFormat `json:"format,omitempty"`
	}{
		device,
		size,
		format,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "ringbuf-read",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// ringbuf-write -> RingbufWrite (command)

// RingbufWrite implements the "ringbuf-write" QMP API call.
func (m *Monitor) RingbufWrite(device string, data string, format *DataFormat) (err error) {
	cmd := struct {
		Device string      `json:"device"`
		Data   string      `json:"data"`
		Format *DataFormat `json:"format,omitempty"`
	}{
		device,
		data,
		format,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "ringbuf-write",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// rtc-reset-reinjection -> RtcResetReinjection (command)

// RtcResetReinjection implements the "rtc-reset-reinjection" QMP API call.
func (m *Monitor) RtcResetReinjection() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "rtc-reset-reinjection",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// screendump -> Screendump (command)

// Screendump implements the "screendump" QMP API call.
func (m *Monitor) Screendump(filename string) (err error) {
	cmd := struct {
		Filename string `json:"filename"`
	}{
		filename,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "screendump",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// send-key -> SendKey (command)

// SendKey implements the "send-key" QMP API call.
func (m *Monitor) SendKey(keys []KeyValue, holdTime *int64) (err error) {
	cmd := struct {
		Keys     []KeyValue `json:"keys"`
		HoldTime *int64     `json:"hold-time,omitempty"`
	}{
		keys,
		holdTime,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "send-key",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// set_link -> SetLink (command)

// SetLink implements the "set_link" QMP API call.
func (m *Monitor) SetLink(name string, up bool) (err error) {
	cmd := struct {
		Name string `json:"name"`
		Up   bool   `json:"up"`
	}{
		name,
		up,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "set_link",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// set_password -> SetPassword (command)

// SetPassword implements the "set_password" QMP API call.
func (m *Monitor) SetPassword(protocol string, password string, connected *string) (err error) {
	cmd := struct {
		Protocol  string  `json:"protocol"`
		Password  string  `json:"password"`
		Connected *string `json:"connected,omitempty"`
	}{
		protocol,
		password,
		connected,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "set_password",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// stop -> Stop (command)

// Stop implements the "stop" QMP API call.
func (m *Monitor) Stop() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "stop",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// system_powerdown -> SystemPowerdown (command)

// SystemPowerdown implements the "system_powerdown" QMP API call.
func (m *Monitor) SystemPowerdown() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "system_powerdown",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// system_reset -> SystemReset (command)

// SystemReset implements the "system_reset" QMP API call.
func (m *Monitor) SystemReset() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "system_reset",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// system_wakeup -> SystemWakeup (command)

// SystemWakeup implements the "system_wakeup" QMP API call.
func (m *Monitor) SystemWakeup() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "system_wakeup",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// trace-event-get-state -> TraceEventGetState (command)

// TraceEventGetState implements the "trace-event-get-state" QMP API call.
func (m *Monitor) TraceEventGetState(name string, vcpu *int64) (ret []TraceEventInfo, err error) {
	cmd := struct {
		Name string `json:"name"`
		Vcpu *int64 `json:"vcpu,omitempty"`
	}{
		name,
		vcpu,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "trace-event-get-state",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	res := struct {
		Res json.RawMessage `json:"return"`
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// trace-event-set-state -> TraceEventSetState (command)

// TraceEventSetState implements the "trace-event-set-state" QMP API call.
func (m *Monitor) TraceEventSetState(name string, enable bool, ignoreUnavailable *bool, vcpu *int64) (err error) {
	cmd := struct {
		Name              string `json:"name"`
		Enable            bool   `json:"enable"`
		IgnoreUnavailable *bool  `json:"ignore-unavailable,omitempty"`
		Vcpu              *int64 `json:"vcpu,omitempty"`
	}{
		name,
		enable,
		ignoreUnavailable,
		vcpu,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "trace-event-set-state",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// transaction -> Transaction (command)

// Transaction implements the "transaction" QMP API call.
func (m *Monitor) Transaction(actions []TransactionAction, properties *TransactionProperties) (err error) {
	cmd := struct {
		Actions    []TransactionAction    `json:"actions"`
		Properties *TransactionProperties `json:"properties,omitempty"`
	}{
		actions,
		properties,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "transaction",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// watchdog-set-action -> WatchdogSetAction (command)

// WatchdogSetAction implements the "watchdog-set-action" QMP API call.
func (m *Monitor) WatchdogSetAction(action WatchdogAction) (err error) {
	cmd := struct {
		Action WatchdogAction `json:"action"`
	}{
		action,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "watchdog-set-action",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// x-blockdev-change -> XBlockdevChange (command)

// XBlockdevChange implements the "x-blockdev-change" QMP API call.
func (m *Monitor) XBlockdevChange(parent string, child *string, node *string) (err error) {
	cmd := struct {
		Parent string  `json:"parent"`
		Child  *string `json:"child,omitempty"`
		Node   *string `json:"node,omitempty"`
	}{
		parent,
		child,
		node,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-blockdev-change",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// x-blockdev-insert-medium -> XBlockdevInsertMedium (command)

// XBlockdevInsertMedium implements the "x-blockdev-insert-medium" QMP API call.
func (m *Monitor) XBlockdevInsertMedium(device *string, id *string, nodeName string) (err error) {
	cmd := struct {
		Device   *string `json:"device,omitempty"`
		ID       *string `json:"id,omitempty"`
		NodeName string  `json:"node-name"`
	}{
		device,
		id,
		nodeName,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-blockdev-insert-medium",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// x-blockdev-remove-medium -> XBlockdevRemoveMedium (command)

// XBlockdevRemoveMedium implements the "x-blockdev-remove-medium" QMP API call.
func (m *Monitor) XBlockdevRemoveMedium(device *string, id *string) (err error) {
	cmd := struct {
		Device *string `json:"device,omitempty"`
		ID     *string `json:"id,omitempty"`
	}{
		device,
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-blockdev-remove-medium",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// x-colo-lost-heartbeat -> XColoLostHeartbeat (command)

// XColoLostHeartbeat implements the "x-colo-lost-heartbeat" QMP API call.
func (m *Monitor) XColoLostHeartbeat() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-colo-lost-heartbeat",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// xen-colo-do-checkpoint -> XenColoDoCheckpoint (command)

// XenColoDoCheckpoint implements the "xen-colo-do-checkpoint" QMP API call.
func (m *Monitor) XenColoDoCheckpoint() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "xen-colo-do-checkpoint",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// xen-load-devices-state -> XenLoadDevicesState (command)

// XenLoadDevicesState implements the "xen-load-devices-state" QMP API call.
func (m *Monitor) XenLoadDevicesState(filename string) (err error) {
	cmd := struct {
		Filename string `json:"filename"`
	}{
		filename,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "xen-load-devices-state",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// xen-save-devices-state -> XenSaveDevicesState (command)

// XenSaveDevicesState implements the "xen-save-devices-state" QMP API call.
func (m *Monitor) XenSaveDevicesState(filename string, live *bool) (err error) {
	cmd := struct {
		Filename string `json:"filename"`
		Live     *bool  `json:"live,omitempty"`
	}{
		filename,
		live,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "xen-save-devices-state",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// xen-set-global-dirty-log -> XenSetGlobalDirtyLog (command)

// XenSetGlobalDirtyLog implements the "xen-set-global-dirty-log" QMP API call.
func (m *Monitor) XenSetGlobalDirtyLog(enable bool) (err error) {
	cmd := struct {
		Enable bool `json:"enable"`
	}{
		enable,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "xen-set-global-dirty-log",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}

// xen-set-replication -> XenSetReplication (command)

// XenSetReplication implements the "xen-set-replication" QMP API call.
func (m *Monitor) XenSetReplication(enable bool, primary bool, failover *bool) (err error) {
	cmd := struct {
		Enable   bool  `json:"enable"`
		Primary  bool  `json:"primary"`
		Failover *bool `json:"failover,omitempty"`
	}{
		enable,
		primary,
		failover,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "xen-set-replication",
		"arguments": cmd,
	})
	if err != nil {
		return
	}
	bs, err = m.mon.Run(bs)
	if err != nil {
		return
	}
	return
}
