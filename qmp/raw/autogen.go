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

// AnnounceParameters -> AnnounceParameters (struct)

// AnnounceParameters implements the "AnnounceParameters" QMP API type.
type AnnounceParameters struct {
	Initial    int64    `json:"initial"`
	Max        int64    `json:"max"`
	Rounds     int64    `json:"rounds"`
	Step       int64    `json:"step"`
	Interfaces []string `json:"interfaces,omitempty"`
	ID         *string  `json:"id,omitempty"`
}

// EVENT BALLOON_CHANGE

// EVENT BLOCK_EXPORT_DELETED

// EVENT BLOCK_IMAGE_CORRUPTED

// EVENT BLOCK_IO_ERROR

// EVENT BLOCK_JOB_CANCELLED

// EVENT BLOCK_JOB_COMPLETED

// EVENT BLOCK_JOB_ERROR

// EVENT BLOCK_JOB_PENDING

// EVENT BLOCK_JOB_READY

// EVENT BLOCK_WRITE_THRESHOLD

// BackupPerf -> BackupPerf (struct)

// BackupPerf implements the "BackupPerf" QMP API type.
type BackupPerf struct {
	UseCopyRange *bool  `json:"use-copy-range,omitempty"`
	MaxWorkers   *int64 `json:"max-workers,omitempty"`
	MaxChunk     *int64 `json:"max-chunk,omitempty"`
}

// BalloonInfo -> BalloonInfo (struct)

// BalloonInfo implements the "BalloonInfo" QMP API type.
type BalloonInfo struct {
	Actual int64 `json:"actual"`
}

// BitmapMigrationBitmapAlias -> BitmapMigrationBitmapAlias (struct)

// BitmapMigrationBitmapAlias implements the "BitmapMigrationBitmapAlias" QMP API type.
type BitmapMigrationBitmapAlias struct {
	Name      string                               `json:"name"`
	Alias     string                               `json:"alias"`
	Transform *BitmapMigrationBitmapAliasTransform `json:"transform,omitempty"`
}

// BitmapMigrationBitmapAliasTransform -> BitmapMigrationBitmapAliasTransform (struct)

// BitmapMigrationBitmapAliasTransform implements the "BitmapMigrationBitmapAliasTransform" QMP API type.
type BitmapMigrationBitmapAliasTransform struct {
	Persistent *bool `json:"persistent,omitempty"`
}

// BitmapMigrationNodeAlias -> BitmapMigrationNodeAlias (struct)

// BitmapMigrationNodeAlias implements the "BitmapMigrationNodeAlias" QMP API type.
type BitmapMigrationNodeAlias struct {
	NodeName string                       `json:"node-name"`
	Alias    string                       `json:"alias"`
	Bitmaps  []BitmapMigrationBitmapAlias `json:"bitmaps"`
}

// BitmapSyncMode -> BitmapSyncMode (enum)

// BitmapSyncMode implements the "BitmapSyncMode" QMP API type.
type BitmapSyncMode int

// Known values of BitmapSyncMode.
const (
	BitmapSyncModeOnSuccess BitmapSyncMode = iota
	BitmapSyncModeNever
	BitmapSyncModeAlways
)

// String implements fmt.Stringer.
func (e BitmapSyncMode) String() string {
	switch e {
	case BitmapSyncModeOnSuccess:
		return "on-success"
	case BitmapSyncModeNever:
		return "never"
	case BitmapSyncModeAlways:
		return "always"
	default:
		return fmt.Sprintf("BitmapSyncMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BitmapSyncMode) MarshalJSON() ([]byte, error) {
	switch e {
	case BitmapSyncModeOnSuccess:
		return json.Marshal("on-success")
	case BitmapSyncModeNever:
		return json.Marshal("never")
	case BitmapSyncModeAlways:
		return json.Marshal("always")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BitmapSyncMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BitmapSyncMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "on-success":
		*e = BitmapSyncModeOnSuccess
	case "never":
		*e = BitmapSyncModeNever
	case "always":
		*e = BitmapSyncModeAlways
	default:
		return fmt.Errorf("unknown enum value %q for BitmapSyncMode", s)
	}
	return nil
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
	BlkdebugEventClusterAllocSpace
	BlkdebugEventNone
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
	case BlkdebugEventClusterAllocSpace:
		return "cluster_alloc_space"
	case BlkdebugEventNone:
		return "none"
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
	case BlkdebugEventClusterAllocSpace:
		return json.Marshal("cluster_alloc_space")
	case BlkdebugEventNone:
		return json.Marshal("none")
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
	case "cluster_alloc_space":
		*e = BlkdebugEventClusterAllocSpace
	case "none":
		*e = BlkdebugEventNone
	default:
		return fmt.Errorf("unknown enum value %q for BlkdebugEvent", s)
	}
	return nil
}

// BlkdebugIOType -> BlkdebugIOType (enum)

// BlkdebugIOType implements the "BlkdebugIOType" QMP API type.
type BlkdebugIOType int

// Known values of BlkdebugIOType.
const (
	BlkdebugIOTypeRead BlkdebugIOType = iota
	BlkdebugIOTypeWrite
	BlkdebugIOTypeWriteZeroes
	BlkdebugIOTypeDiscard
	BlkdebugIOTypeFlush
	BlkdebugIOTypeBlockStatus
)

// String implements fmt.Stringer.
func (e BlkdebugIOType) String() string {
	switch e {
	case BlkdebugIOTypeRead:
		return "read"
	case BlkdebugIOTypeWrite:
		return "write"
	case BlkdebugIOTypeWriteZeroes:
		return "write-zeroes"
	case BlkdebugIOTypeDiscard:
		return "discard"
	case BlkdebugIOTypeFlush:
		return "flush"
	case BlkdebugIOTypeBlockStatus:
		return "block-status"
	default:
		return fmt.Sprintf("BlkdebugIOType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlkdebugIOType) MarshalJSON() ([]byte, error) {
	switch e {
	case BlkdebugIOTypeRead:
		return json.Marshal("read")
	case BlkdebugIOTypeWrite:
		return json.Marshal("write")
	case BlkdebugIOTypeWriteZeroes:
		return json.Marshal("write-zeroes")
	case BlkdebugIOTypeDiscard:
		return json.Marshal("discard")
	case BlkdebugIOTypeFlush:
		return json.Marshal("flush")
	case BlkdebugIOTypeBlockStatus:
		return json.Marshal("block-status")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlkdebugIOType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlkdebugIOType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "read":
		*e = BlkdebugIOTypeRead
	case "write":
		*e = BlkdebugIOTypeWrite
	case "write-zeroes":
		*e = BlkdebugIOTypeWriteZeroes
	case "discard":
		*e = BlkdebugIOTypeDiscard
	case "flush":
		*e = BlkdebugIOTypeFlush
	case "block-status":
		*e = BlkdebugIOTypeBlockStatus
	default:
		return fmt.Errorf("unknown enum value %q for BlkdebugIOType", s)
	}
	return nil
}

// BlkdebugInjectErrorOptions -> BlkdebugInjectErrorOptions (struct)

// BlkdebugInjectErrorOptions implements the "BlkdebugInjectErrorOptions" QMP API type.
type BlkdebugInjectErrorOptions struct {
	Event       BlkdebugEvent   `json:"event"`
	State       *int64          `json:"state,omitempty"`
	Iotype      *BlkdebugIOType `json:"iotype,omitempty"`
	Errno       *int64          `json:"errno,omitempty"`
	Sector      *int64          `json:"sector,omitempty"`
	Once        *bool           `json:"once,omitempty"`
	Immediately *bool           `json:"immediately,omitempty"`
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
	File             string                      `json:"file"`
	NodeName         *string                     `json:"node-name,omitempty"`
	Ro               bool                        `json:"ro"`
	Drv              string                      `json:"drv"`
	BackingFile      *string                     `json:"backing_file,omitempty"`
	BackingFileDepth int64                       `json:"backing_file_depth"`
	Encrypted        bool                        `json:"encrypted"`
	DetectZeroes     BlockdevDetectZeroesOptions `json:"detect_zeroes"`
	Bps              int64                       `json:"bps"`
	BpsRd            int64                       `json:"bps_rd"`
	BpsWr            int64                       `json:"bps_wr"`
	Iops             int64                       `json:"iops"`
	IopsRd           int64                       `json:"iops_rd"`
	IopsWr           int64                       `json:"iops_wr"`
	Image            ImageInfo                   `json:"image"`
	BpsMax           *int64                      `json:"bps_max,omitempty"`
	BpsRdMax         *int64                      `json:"bps_rd_max,omitempty"`
	BpsWrMax         *int64                      `json:"bps_wr_max,omitempty"`
	IopsMax          *int64                      `json:"iops_max,omitempty"`
	IopsRdMax        *int64                      `json:"iops_rd_max,omitempty"`
	IopsWrMax        *int64                      `json:"iops_wr_max,omitempty"`
	BpsMaxLength     *int64                      `json:"bps_max_length,omitempty"`
	BpsRdMaxLength   *int64                      `json:"bps_rd_max_length,omitempty"`
	BpsWrMaxLength   *int64                      `json:"bps_wr_max_length,omitempty"`
	IopsMaxLength    *int64                      `json:"iops_max_length,omitempty"`
	IopsRdMaxLength  *int64                      `json:"iops_rd_max_length,omitempty"`
	IopsWrMaxLength  *int64                      `json:"iops_wr_max_length,omitempty"`
	IopsSize         *int64                      `json:"iops_size,omitempty"`
	Group            *string                     `json:"group,omitempty"`
	Cache            BlockdevCacheInfo           `json:"cache"`
	WriteThreshold   int64                       `json:"write_threshold"`
	DirtyBitmaps     []BlockDirtyInfo            `json:"dirty-bitmaps,omitempty"`
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
	RdBytes                int64                      `json:"rd_bytes"`
	WrBytes                int64                      `json:"wr_bytes"`
	UnmapBytes             int64                      `json:"unmap_bytes"`
	RdOperations           int64                      `json:"rd_operations"`
	WrOperations           int64                      `json:"wr_operations"`
	FlushOperations        int64                      `json:"flush_operations"`
	UnmapOperations        int64                      `json:"unmap_operations"`
	RdTotalTimeNs          int64                      `json:"rd_total_time_ns"`
	WrTotalTimeNs          int64                      `json:"wr_total_time_ns"`
	FlushTotalTimeNs       int64                      `json:"flush_total_time_ns"`
	UnmapTotalTimeNs       int64                      `json:"unmap_total_time_ns"`
	WrHighestOffset        int64                      `json:"wr_highest_offset"`
	RdMerged               int64                      `json:"rd_merged"`
	WrMerged               int64                      `json:"wr_merged"`
	UnmapMerged            int64                      `json:"unmap_merged"`
	IdleTimeNs             *int64                     `json:"idle_time_ns,omitempty"`
	FailedRdOperations     int64                      `json:"failed_rd_operations"`
	FailedWrOperations     int64                      `json:"failed_wr_operations"`
	FailedFlushOperations  int64                      `json:"failed_flush_operations"`
	FailedUnmapOperations  int64                      `json:"failed_unmap_operations"`
	InvalidRdOperations    int64                      `json:"invalid_rd_operations"`
	InvalidWrOperations    int64                      `json:"invalid_wr_operations"`
	InvalidFlushOperations int64                      `json:"invalid_flush_operations"`
	InvalidUnmapOperations int64                      `json:"invalid_unmap_operations"`
	AccountInvalid         bool                       `json:"account_invalid"`
	AccountFailed          bool                       `json:"account_failed"`
	TimedStats             []BlockDeviceTimedStats    `json:"timed_stats"`
	RdLatencyHistogram     *BlockLatencyHistogramInfo `json:"rd_latency_histogram,omitempty"`
	WrLatencyHistogram     *BlockLatencyHistogramInfo `json:"wr_latency_histogram,omitempty"`
	FlushLatencyHistogram  *BlockLatencyHistogramInfo `json:"flush_latency_histogram,omitempty"`
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
	Disabled    *bool   `json:"disabled,omitempty"`
}

// BlockDirtyBitmapMerge -> BlockDirtyBitmapMerge (struct)

// BlockDirtyBitmapMerge implements the "BlockDirtyBitmapMerge" QMP API type.
type BlockDirtyBitmapMerge struct {
	Node    string                  `json:"node"`
	Target  string                  `json:"target"`
	Bitmaps []BlockDirtyBitmapOrStr `json:"bitmaps"`
}

// BlockDirtyBitmapOrStr -> BlockDirtyBitmapOrStr (alternate)

// BlockDirtyBitmapOrStr implements the "BlockDirtyBitmapOrStr" QMP API type.
//
// Can be one of:
//   - BlockDirtyBitmapOrStrVariantExternal
//   - BlockDirtyBitmapOrStrVariantLocal
type BlockDirtyBitmapOrStr interface {
	isBlockDirtyBitmapOrStr()
}

// BlockDirtyBitmapOrStrVariantExternal is an implementation of BlockDirtyBitmapOrStr
type BlockDirtyBitmapOrStrVariantExternal BlockDirtyBitmap

func (BlockDirtyBitmapOrStrVariantExternal) isBlockDirtyBitmapOrStr() {}

// BlockDirtyBitmapOrStrVariantLocal is an implementation of BlockDirtyBitmapOrStr
type BlockDirtyBitmapOrStrVariantLocal string

func (BlockDirtyBitmapOrStrVariantLocal) isBlockDirtyBitmapOrStr() {}

func decodeBlockDirtyBitmapOrStr(bs json.RawMessage) (BlockDirtyBitmapOrStr, error) {

	var local BlockDirtyBitmapOrStrVariantLocal
	if err := json.Unmarshal([]byte(bs), &local); err == nil {
		return local, nil
	}

	var external BlockDirtyBitmapOrStrVariantExternal
	if err := json.Unmarshal([]byte(bs), &external); err == nil {
		return external, nil
	}
	return nil, fmt.Errorf("unable to decode %q as a BlockDirtyBitmapOrStr", string(bs))
}

// BlockDirtyBitmapSha256 -> BlockDirtyBitmapSha256 (struct)

// BlockDirtyBitmapSha256 implements the "BlockDirtyBitmapSha256" QMP API type.
type BlockDirtyBitmapSha256 struct {
	Sha256 string `json:"sha256"`
}

// BlockDirtyInfo -> BlockDirtyInfo (struct)

// BlockDirtyInfo implements the "BlockDirtyInfo" QMP API type.
type BlockDirtyInfo struct {
	Name         *string `json:"name,omitempty"`
	Count        int64   `json:"count"`
	Granularity  uint32  `json:"granularity"`
	Recording    bool    `json:"recording"`
	Busy         bool    `json:"busy"`
	Persistent   bool    `json:"persistent"`
	Inconsistent *bool   `json:"inconsistent,omitempty"`
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

// BlockExportInfo -> BlockExportInfo (struct)

// BlockExportInfo implements the "BlockExportInfo" QMP API type.
type BlockExportInfo struct {
	ID           string          `json:"id"`
	Type         BlockExportType `json:"type"`
	NodeName     string          `json:"node-name"`
	ShuttingDown bool            `json:"shutting-down"`
}

// BlockExportOptions -> BlockExportOptions (flat union)

// BlockExportOptions implements the "BlockExportOptions" QMP API type.
//
// Can be one of:
//   - BlockExportOptionsVariantFuse
//   - BlockExportOptionsVariantNBD
//   - BlockExportOptionsVariantVduseBlk
//   - BlockExportOptionsVariantVhostUserBlk
type BlockExportOptions interface {
	isBlockExportOptions()
}

// BlockExportOptionsVariantFuse is an implementation of BlockExportOptions.
type BlockExportOptionsVariantFuse struct {
	ID            string                `json:"id"`
	FixedIothread *bool                 `json:"fixed-iothread,omitempty"`
	Iothread      *string               `json:"iothread,omitempty"`
	NodeName      string                `json:"node-name"`
	Writable      *bool                 `json:"writable,omitempty"`
	Writethrough  *bool                 `json:"writethrough,omitempty"`
	Mountpoint    string                `json:"mountpoint"`
	Growable      *bool                 `json:"growable,omitempty"`
	AllowOther    *FuseExportAllowOther `json:"allow-other,omitempty"`
}

func (BlockExportOptionsVariantFuse) isBlockExportOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockExportOptionsVariantFuse) MarshalJSON() ([]byte, error) {
	v := struct {
		Type BlockExportType `json:"type"`
		BlockExportOptionsVariantFuse
	}{
		BlockExportTypeFuse,
		s,
	}
	return json.Marshal(v)
}

// BlockExportOptionsVariantNBD is an implementation of BlockExportOptions.
type BlockExportOptionsVariantNBD struct {
	ID              string                  `json:"id"`
	FixedIothread   *bool                   `json:"fixed-iothread,omitempty"`
	Iothread        *string                 `json:"iothread,omitempty"`
	NodeName        string                  `json:"node-name"`
	Writable        *bool                   `json:"writable,omitempty"`
	Writethrough    *bool                   `json:"writethrough,omitempty"`
	Bitmaps         []BlockDirtyBitmapOrStr `json:"bitmaps,omitempty"`
	AllocationDepth *bool                   `json:"allocation-depth,omitempty"`
}

func (BlockExportOptionsVariantNBD) isBlockExportOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockExportOptionsVariantNBD) MarshalJSON() ([]byte, error) {
	v := struct {
		Type BlockExportType `json:"type"`
		BlockExportOptionsVariantNBD
	}{
		BlockExportTypeNBD,
		s,
	}
	return json.Marshal(v)
}

// BlockExportOptionsVariantVduseBlk is an implementation of BlockExportOptions.
type BlockExportOptionsVariantVduseBlk struct {
	ID               string  `json:"id"`
	FixedIothread    *bool   `json:"fixed-iothread,omitempty"`
	Iothread         *string `json:"iothread,omitempty"`
	NodeName         string  `json:"node-name"`
	Writable         *bool   `json:"writable,omitempty"`
	Writethrough     *bool   `json:"writethrough,omitempty"`
	Name             string  `json:"name"`
	NumQueues        *uint16 `json:"num-queues,omitempty"`
	QueueSize        *uint16 `json:"queue-size,omitempty"`
	LogicalBlockSize *uint64 `json:"logical-block-size,omitempty"`
	Serial           *string `json:"serial,omitempty"`
}

func (BlockExportOptionsVariantVduseBlk) isBlockExportOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockExportOptionsVariantVduseBlk) MarshalJSON() ([]byte, error) {
	v := struct {
		Type BlockExportType `json:"type"`
		BlockExportOptionsVariantVduseBlk
	}{
		BlockExportTypeVduseBlk,
		s,
	}
	return json.Marshal(v)
}

// BlockExportOptionsVariantVhostUserBlk is an implementation of BlockExportOptions.
type BlockExportOptionsVariantVhostUserBlk struct {
	ID               string        `json:"id"`
	FixedIothread    *bool         `json:"fixed-iothread,omitempty"`
	Iothread         *string       `json:"iothread,omitempty"`
	NodeName         string        `json:"node-name"`
	Writable         *bool         `json:"writable,omitempty"`
	Writethrough     *bool         `json:"writethrough,omitempty"`
	Addr             SocketAddress `json:"addr"`
	LogicalBlockSize *uint64       `json:"logical-block-size,omitempty"`
	NumQueues        *uint16       `json:"num-queues,omitempty"`
}

func (BlockExportOptionsVariantVhostUserBlk) isBlockExportOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockExportOptionsVariantVhostUserBlk) MarshalJSON() ([]byte, error) {
	v := struct {
		Type BlockExportType `json:"type"`
		BlockExportOptionsVariantVhostUserBlk
	}{
		BlockExportTypeVhostUserBlk,
		s,
	}
	return json.Marshal(v)
}

func decodeBlockExportOptions(bs json.RawMessage) (BlockExportOptions, error) {
	v := struct {
		Type BlockExportType `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case BlockExportTypeFuse:
		var ret BlockExportOptionsVariantFuse
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockExportTypeNBD:
		var ret BlockExportOptionsVariantNBD
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockExportTypeVduseBlk:
		var ret BlockExportOptionsVariantVduseBlk
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockExportTypeVhostUserBlk:
		var ret BlockExportOptionsVariantVhostUserBlk
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union BlockExportOptions", v.Type)
	}
}

// BlockExportRemoveMode -> BlockExportRemoveMode (enum)

// BlockExportRemoveMode implements the "BlockExportRemoveMode" QMP API type.
type BlockExportRemoveMode int

// Known values of BlockExportRemoveMode.
const (
	BlockExportRemoveModeSafe BlockExportRemoveMode = iota
	BlockExportRemoveModeHard
)

// String implements fmt.Stringer.
func (e BlockExportRemoveMode) String() string {
	switch e {
	case BlockExportRemoveModeSafe:
		return "safe"
	case BlockExportRemoveModeHard:
		return "hard"
	default:
		return fmt.Sprintf("BlockExportRemoveMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockExportRemoveMode) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockExportRemoveModeSafe:
		return json.Marshal("safe")
	case BlockExportRemoveModeHard:
		return json.Marshal("hard")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockExportRemoveMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockExportRemoveMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "safe":
		*e = BlockExportRemoveModeSafe
	case "hard":
		*e = BlockExportRemoveModeHard
	default:
		return fmt.Errorf("unknown enum value %q for BlockExportRemoveMode", s)
	}
	return nil
}

// BlockExportType -> BlockExportType (enum)

// BlockExportType implements the "BlockExportType" QMP API type.
type BlockExportType int

// Known values of BlockExportType.
const (
	BlockExportTypeNBD BlockExportType = iota
	BlockExportTypeVhostUserBlk
	BlockExportTypeFuse
	BlockExportTypeVduseBlk
)

// String implements fmt.Stringer.
func (e BlockExportType) String() string {
	switch e {
	case BlockExportTypeNBD:
		return "nbd"
	case BlockExportTypeVhostUserBlk:
		return "vhost-user-blk"
	case BlockExportTypeFuse:
		return "fuse"
	case BlockExportTypeVduseBlk:
		return "vduse-blk"
	default:
		return fmt.Sprintf("BlockExportType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockExportType) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockExportTypeNBD:
		return json.Marshal("nbd")
	case BlockExportTypeVhostUserBlk:
		return json.Marshal("vhost-user-blk")
	case BlockExportTypeFuse:
		return json.Marshal("fuse")
	case BlockExportTypeVduseBlk:
		return json.Marshal("vduse-blk")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockExportType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockExportType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "nbd":
		*e = BlockExportTypeNBD
	case "vhost-user-blk":
		*e = BlockExportTypeVhostUserBlk
	case "fuse":
		*e = BlockExportTypeFuse
	case "vduse-blk":
		*e = BlockExportTypeVduseBlk
	default:
		return fmt.Errorf("unknown enum value %q for BlockExportType", s)
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
	Device    string               `json:"device"`
	Qdev      *string              `json:"qdev,omitempty"`
	Type      string               `json:"type"`
	Removable bool                 `json:"removable"`
	Locked    bool                 `json:"locked"`
	Inserted  *BlockDeviceInfo     `json:"inserted,omitempty"`
	TrayOpen  *bool                `json:"tray_open,omitempty"`
	IOStatus  *BlockDeviceIOStatus `json:"io-status,omitempty"`
}

// BlockJobInfo -> BlockJobInfo (struct)

// BlockJobInfo implements the "BlockJobInfo" QMP API type.
type BlockJobInfo struct {
	Type         string              `json:"type"`
	Device       string              `json:"device"`
	Len          int64               `json:"len"`
	Offset       int64               `json:"offset"`
	Busy         bool                `json:"busy"`
	Paused       bool                `json:"paused"`
	Speed        int64               `json:"speed"`
	IOStatus     BlockDeviceIOStatus `json:"io-status"`
	Ready        bool                `json:"ready"`
	Status       JobStatus           `json:"status"`
	AutoFinalize bool                `json:"auto-finalize"`
	AutoDismiss  bool                `json:"auto-dismiss"`
	Error        *string             `json:"error,omitempty"`
}

// BlockLatencyHistogramInfo -> BlockLatencyHistogramInfo (struct)

// BlockLatencyHistogramInfo implements the "BlockLatencyHistogramInfo" QMP API type.
type BlockLatencyHistogramInfo struct {
	Boundaries []uint64 `json:"boundaries"`
	Bins       []uint64 `json:"bins"`
}

// BlockPermission -> BlockPermission (enum)

// BlockPermission implements the "BlockPermission" QMP API type.
type BlockPermission int

// Known values of BlockPermission.
const (
	BlockPermissionConsistentRead BlockPermission = iota
	BlockPermissionWrite
	BlockPermissionWriteUnchanged
	BlockPermissionResize
)

// String implements fmt.Stringer.
func (e BlockPermission) String() string {
	switch e {
	case BlockPermissionConsistentRead:
		return "consistent-read"
	case BlockPermissionWrite:
		return "write"
	case BlockPermissionWriteUnchanged:
		return "write-unchanged"
	case BlockPermissionResize:
		return "resize"
	default:
		return fmt.Sprintf("BlockPermission(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockPermission) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockPermissionConsistentRead:
		return json.Marshal("consistent-read")
	case BlockPermissionWrite:
		return json.Marshal("write")
	case BlockPermissionWriteUnchanged:
		return json.Marshal("write-unchanged")
	case BlockPermissionResize:
		return json.Marshal("resize")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockPermission", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockPermission) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "consistent-read":
		*e = BlockPermissionConsistentRead
	case "write":
		*e = BlockPermissionWrite
	case "write-unchanged":
		*e = BlockPermissionWriteUnchanged
	case "resize":
		*e = BlockPermissionResize
	default:
		return fmt.Errorf("unknown enum value %q for BlockPermission", s)
	}
	return nil
}

// BlockStats -> BlockStats (struct)

// BlockStats implements the "BlockStats" QMP API type.
type BlockStats struct {
	Device         *string             `json:"device,omitempty"`
	Qdev           *string             `json:"qdev,omitempty"`
	NodeName       *string             `json:"node-name,omitempty"`
	Stats          BlockDeviceStats    `json:"stats"`
	DriverSpecific *BlockStatsSpecific `json:"driver-specific,omitempty"`
	Parent         *BlockStats         `json:"parent,omitempty"`
	Backing        *BlockStats         `json:"backing,omitempty"`
}

// BlockStatsSpecific -> BlockStatsSpecific (flat union)

// BlockStatsSpecific implements the "BlockStatsSpecific" QMP API type.
//
// Can be one of:
//   - BlockStatsSpecificVariantFile
//   - BlockStatsSpecificVariantHostDevice
//   - BlockStatsSpecificVariantNvme
type BlockStatsSpecific interface {
	isBlockStatsSpecific()
}

// BlockStatsSpecificVariantFile is an implementation of BlockStatsSpecific.
type BlockStatsSpecificVariantFile struct {
	DiscardNbOK     uint64 `json:"discard-nb-ok"`
	DiscardNbFailed uint64 `json:"discard-nb-failed"`
	DiscardBytesOK  uint64 `json:"discard-bytes-ok"`
}

func (BlockStatsSpecificVariantFile) isBlockStatsSpecific() {}

// MarshalJSON implements json.Marshaler.
func (s BlockStatsSpecificVariantFile) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockStatsSpecificVariantFile
	}{
		BlockdevDriverFile,
		s,
	}
	return json.Marshal(v)
}

// BlockStatsSpecificVariantHostDevice is an implementation of BlockStatsSpecific.
type BlockStatsSpecificVariantHostDevice struct {
	DiscardNbOK     uint64 `json:"discard-nb-ok"`
	DiscardNbFailed uint64 `json:"discard-nb-failed"`
	DiscardBytesOK  uint64 `json:"discard-bytes-ok"`
}

func (BlockStatsSpecificVariantHostDevice) isBlockStatsSpecific() {}

// MarshalJSON implements json.Marshaler.
func (s BlockStatsSpecificVariantHostDevice) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockStatsSpecificVariantHostDevice
	}{
		BlockdevDriverHostDevice,
		s,
	}
	return json.Marshal(v)
}

// BlockStatsSpecificVariantNvme is an implementation of BlockStatsSpecific.
type BlockStatsSpecificVariantNvme struct {
	CompletionErrors  uint64 `json:"completion-errors"`
	AlignedAccesses   uint64 `json:"aligned-accesses"`
	UnalignedAccesses uint64 `json:"unaligned-accesses"`
}

func (BlockStatsSpecificVariantNvme) isBlockStatsSpecific() {}

// MarshalJSON implements json.Marshaler.
func (s BlockStatsSpecificVariantNvme) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockStatsSpecificVariantNvme
	}{
		BlockdevDriverNvme,
		s,
	}
	return json.Marshal(v)
}

func decodeBlockStatsSpecific(bs json.RawMessage) (BlockStatsSpecific, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Driver {
	case BlockdevDriverFile:
		var ret BlockStatsSpecificVariantFile
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverHostDevice:
		var ret BlockStatsSpecificVariantHostDevice
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNvme:
		var ret BlockStatsSpecificVariantNvme
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union BlockStatsSpecific", v.Driver)
	}
}

// BlockdevAioOptions -> BlockdevAIOOptions (enum)

// BlockdevAIOOptions implements the "BlockdevAioOptions" QMP API type.
type BlockdevAIOOptions int

// Known values of BlockdevAIOOptions.
const (
	BlockdevAIOOptionsThreads BlockdevAIOOptions = iota
	BlockdevAIOOptionsNative
	BlockdevAIOOptionsIOUring
)

// String implements fmt.Stringer.
func (e BlockdevAIOOptions) String() string {
	switch e {
	case BlockdevAIOOptionsThreads:
		return "threads"
	case BlockdevAIOOptionsNative:
		return "native"
	case BlockdevAIOOptionsIOUring:
		return "io_uring"
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
	case BlockdevAIOOptionsIOUring:
		return json.Marshal("io_uring")
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
	case "io_uring":
		*e = BlockdevAIOOptionsIOUring
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevAIOOptions", s)
	}
	return nil
}

// BlockdevAmendOptions -> BlockdevAmendOptions (flat union)

// BlockdevAmendOptions implements the "BlockdevAmendOptions" QMP API type.
//
// Can be one of:
//   - BlockdevAmendOptionsVariantLUKS
//   - BlockdevAmendOptionsVariantQcow2
type BlockdevAmendOptions interface {
	isBlockdevAmendOptions()
}

// BlockdevAmendOptionsVariantLUKS is an implementation of BlockdevAmendOptions.
type BlockdevAmendOptionsVariantLUKS struct {
}

func (BlockdevAmendOptionsVariantLUKS) isBlockdevAmendOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevAmendOptionsVariantLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevAmendOptionsVariantLUKS
	}{
		BlockdevDriverLUKS,
		s,
	}
	return json.Marshal(v)
}

// BlockdevAmendOptionsVariantQcow2 is an implementation of BlockdevAmendOptions.
type BlockdevAmendOptionsVariantQcow2 struct {
	Encrypt *QCryptoBlockAmendOptions `json:"encrypt,omitempty"`
}

func (BlockdevAmendOptionsVariantQcow2) isBlockdevAmendOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevAmendOptionsVariantQcow2) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevAmendOptionsVariantQcow2
	}{
		BlockdevDriverQcow2,
		s,
	}
	return json.Marshal(v)
}

func decodeBlockdevAmendOptions(bs json.RawMessage) (BlockdevAmendOptions, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Driver {
	case BlockdevDriverLUKS:
		var ret BlockdevAmendOptionsVariantLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQcow2:
		var ret BlockdevAmendOptionsVariantQcow2
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union BlockdevAmendOptions", v.Driver)
	}
}

// BlockdevBackup -> BlockdevBackup (struct)

// BlockdevBackup implements the "BlockdevBackup" QMP API type.
type BlockdevBackup struct {
	JobID          *string          `json:"job-id,omitempty"`
	Device         string           `json:"device"`
	Sync           MirrorSyncMode   `json:"sync"`
	Speed          *int64           `json:"speed,omitempty"`
	Bitmap         *string          `json:"bitmap,omitempty"`
	BitmapMode     *BitmapSyncMode  `json:"bitmap-mode,omitempty"`
	Compress       *bool            `json:"compress,omitempty"`
	OnSourceError  *BlockdevOnError `json:"on-source-error,omitempty"`
	OnTargetError  *BlockdevOnError `json:"on-target-error,omitempty"`
	AutoFinalize   *bool            `json:"auto-finalize,omitempty"`
	AutoDismiss    *bool            `json:"auto-dismiss,omitempty"`
	FilterNodeName *string          `json:"filter-node-name,omitempty"`
	XPerf          *BackupPerf      `json:"x-perf,omitempty"`
	Target         string           `json:"target"`
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

// BlockdevCreateOptions -> BlockdevCreateOptions (flat union)

// BlockdevCreateOptions implements the "BlockdevCreateOptions" QMP API type.
//
// Can be one of:
//   - BlockdevCreateOptionsVariantFile
//   - BlockdevCreateOptionsVariantGluster
//   - BlockdevCreateOptionsVariantLUKS
//   - BlockdevCreateOptionsVariantNfs
//   - BlockdevCreateOptionsVariantParallels
//   - BlockdevCreateOptionsVariantQcow
//   - BlockdevCreateOptionsVariantQcow2
//   - BlockdevCreateOptionsVariantQed
//   - BlockdevCreateOptionsVariantRbd
//   - BlockdevCreateOptionsVariantSSH
//   - BlockdevCreateOptionsVariantVdi
//   - BlockdevCreateOptionsVariantVhdx
//   - BlockdevCreateOptionsVariantVMDK
//   - BlockdevCreateOptionsVariantVpc
type BlockdevCreateOptions interface {
	isBlockdevCreateOptions()
}

// BlockdevCreateOptionsVariantFile is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantFile struct {
	Filename       string        `json:"filename"`
	Size           uint64        `json:"size"`
	Preallocation  *PreallocMode `json:"preallocation,omitempty"`
	Nocow          *bool         `json:"nocow,omitempty"`
	ExtentSizeHint *uint64       `json:"extent-size-hint,omitempty"`
}

func (BlockdevCreateOptionsVariantFile) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantFile) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantFile
	}{
		BlockdevDriverFile,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantGluster is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantGluster struct {
	Location      BlockdevOptionsGluster `json:"location"`
	Size          uint64                 `json:"size"`
	Preallocation *PreallocMode          `json:"preallocation,omitempty"`
}

func (BlockdevCreateOptionsVariantGluster) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantGluster) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantGluster
	}{
		BlockdevDriverGluster,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantLUKS is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantLUKS struct {
	File          BlockdevRef   `json:"file"`
	Size          uint64        `json:"size"`
	Preallocation *PreallocMode `json:"preallocation,omitempty"`
}

func (BlockdevCreateOptionsVariantLUKS) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantLUKS
	}{
		BlockdevDriverLUKS,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantNfs is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantNfs struct {
	Location BlockdevOptionsNfs `json:"location"`
	Size     uint64             `json:"size"`
}

func (BlockdevCreateOptionsVariantNfs) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantNfs) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantNfs
	}{
		BlockdevDriverNfs,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantParallels is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantParallels struct {
	File        BlockdevRef `json:"file"`
	Size        uint64      `json:"size"`
	ClusterSize *uint64     `json:"cluster-size,omitempty"`
}

func (BlockdevCreateOptionsVariantParallels) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantParallels) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantParallels
	}{
		BlockdevDriverParallels,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantQcow is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantQcow struct {
	File        BlockdevRef                `json:"file"`
	Size        uint64                     `json:"size"`
	BackingFile *string                    `json:"backing-file,omitempty"`
	Encrypt     *QCryptoBlockCreateOptions `json:"encrypt,omitempty"`
}

func (BlockdevCreateOptionsVariantQcow) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantQcow) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantQcow
	}{
		BlockdevDriverQcow,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantQcow2 is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantQcow2 struct {
	File            BlockdevRef                `json:"file"`
	DataFile        *BlockdevRef               `json:"data-file,omitempty"`
	DataFileRaw     *bool                      `json:"data-file-raw,omitempty"`
	ExtendedL2      *bool                      `json:"extended-l2,omitempty"`
	Size            uint64                     `json:"size"`
	Version         *BlockdevQcow2Version      `json:"version,omitempty"`
	BackingFile     *string                    `json:"backing-file,omitempty"`
	BackingFmt      *BlockdevDriver            `json:"backing-fmt,omitempty"`
	Encrypt         *QCryptoBlockCreateOptions `json:"encrypt,omitempty"`
	ClusterSize     *uint64                    `json:"cluster-size,omitempty"`
	Preallocation   *PreallocMode              `json:"preallocation,omitempty"`
	LazyRefcounts   *bool                      `json:"lazy-refcounts,omitempty"`
	RefcountBits    *int64                     `json:"refcount-bits,omitempty"`
	CompressionType *Qcow2CompressionType      `json:"compression-type,omitempty"`
}

func (BlockdevCreateOptionsVariantQcow2) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantQcow2) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantQcow2
	}{
		BlockdevDriverQcow2,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantQed is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantQed struct {
	File        BlockdevRef     `json:"file"`
	Size        uint64          `json:"size"`
	BackingFile *string         `json:"backing-file,omitempty"`
	BackingFmt  *BlockdevDriver `json:"backing-fmt,omitempty"`
	ClusterSize *uint64         `json:"cluster-size,omitempty"`
	TableSize   *int64          `json:"table-size,omitempty"`
}

func (BlockdevCreateOptionsVariantQed) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantQed) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantQed
	}{
		BlockdevDriverQed,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantRbd is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantRbd struct {
	Location    BlockdevOptionsRbd          `json:"location"`
	Size        uint64                      `json:"size"`
	ClusterSize *uint64                     `json:"cluster-size,omitempty"`
	Encrypt     *RbdEncryptionCreateOptions `json:"encrypt,omitempty"`
}

func (BlockdevCreateOptionsVariantRbd) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantRbd) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantRbd
	}{
		BlockdevDriverRbd,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantSSH is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantSSH struct {
	Location BlockdevOptionsSSH `json:"location"`
	Size     uint64             `json:"size"`
}

func (BlockdevCreateOptionsVariantSSH) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantSSH) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantSSH
	}{
		BlockdevDriverSSH,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantVdi is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantVdi struct {
	File          BlockdevRef   `json:"file"`
	Size          uint64        `json:"size"`
	Preallocation *PreallocMode `json:"preallocation,omitempty"`
}

func (BlockdevCreateOptionsVariantVdi) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantVdi) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantVdi
	}{
		BlockdevDriverVdi,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantVhdx is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantVhdx struct {
	File           BlockdevRef            `json:"file"`
	Size           uint64                 `json:"size"`
	LogSize        *uint64                `json:"log-size,omitempty"`
	BlockSize      *uint64                `json:"block-size,omitempty"`
	Subformat      *BlockdevVhdxSubformat `json:"subformat,omitempty"`
	BlockStateZero *bool                  `json:"block-state-zero,omitempty"`
}

func (BlockdevCreateOptionsVariantVhdx) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantVhdx) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantVhdx
	}{
		BlockdevDriverVhdx,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantVMDK is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantVMDK struct {
	File         BlockdevRef              `json:"file"`
	Size         uint64                   `json:"size"`
	Extents      []BlockdevRef            `json:"extents,omitempty"`
	Subformat    *BlockdevVMDKSubformat   `json:"subformat,omitempty"`
	BackingFile  *string                  `json:"backing-file,omitempty"`
	AdapterType  *BlockdevVMDKAdapterType `json:"adapter-type,omitempty"`
	Hwversion    *string                  `json:"hwversion,omitempty"`
	Toolsversion *string                  `json:"toolsversion,omitempty"`
	ZeroedGrain  *bool                    `json:"zeroed-grain,omitempty"`
}

func (BlockdevCreateOptionsVariantVMDK) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantVMDK) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantVMDK
	}{
		BlockdevDriverVMDK,
		s,
	}
	return json.Marshal(v)
}

// BlockdevCreateOptionsVariantVpc is an implementation of BlockdevCreateOptions.
type BlockdevCreateOptionsVariantVpc struct {
	File      BlockdevRef           `json:"file"`
	Size      uint64                `json:"size"`
	Subformat *BlockdevVpcSubformat `json:"subformat,omitempty"`
	ForceSize *bool                 `json:"force-size,omitempty"`
}

func (BlockdevCreateOptionsVariantVpc) isBlockdevCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevCreateOptionsVariantVpc) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevCreateOptionsVariantVpc
	}{
		BlockdevDriverVpc,
		s,
	}
	return json.Marshal(v)
}

func decodeBlockdevCreateOptions(bs json.RawMessage) (BlockdevCreateOptions, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Driver {
	case BlockdevDriverFile:
		var ret BlockdevCreateOptionsVariantFile
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverGluster:
		var ret BlockdevCreateOptionsVariantGluster
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverLUKS:
		var ret BlockdevCreateOptionsVariantLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNfs:
		var ret BlockdevCreateOptionsVariantNfs
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverParallels:
		var ret BlockdevCreateOptionsVariantParallels
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQcow:
		var ret BlockdevCreateOptionsVariantQcow
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQcow2:
		var ret BlockdevCreateOptionsVariantQcow2
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQed:
		var ret BlockdevCreateOptionsVariantQed
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverRbd:
		var ret BlockdevCreateOptionsVariantRbd
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverSSH:
		var ret BlockdevCreateOptionsVariantSSH
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVdi:
		var ret BlockdevCreateOptionsVariantVdi
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVhdx:
		var ret BlockdevCreateOptionsVariantVhdx
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVMDK:
		var ret BlockdevCreateOptionsVariantVMDK
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVpc:
		var ret BlockdevCreateOptionsVariantVpc
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union BlockdevCreateOptions", v.Driver)
	}
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
	BlockdevDriverBlklogwrites
	BlockdevDriverBlkreplay
	BlockdevDriverBlkverify
	BlockdevDriverBochs
	BlockdevDriverCloop
	BlockdevDriverCompress
	BlockdevDriverCopyBeforeWrite
	BlockdevDriverCopyOnRead
	BlockdevDriverDmg
	BlockdevDriverFile
	BlockdevDriverSnapshotAccess
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
	BlockdevDriverNvme
	BlockdevDriverParallels
	BlockdevDriverPreallocate
	BlockdevDriverQcow
	BlockdevDriverQcow2
	BlockdevDriverQed
	BlockdevDriverQuorum
	BlockdevDriverRaw
	BlockdevDriverRbd
	BlockdevDriverReplication
	BlockdevDriverSSH
	BlockdevDriverThrottle
	BlockdevDriverVdi
	BlockdevDriverVhdx
	BlockdevDriverVMDK
	BlockdevDriverVpc
	BlockdevDriverVvfat
)

// String implements fmt.Stringer.
func (e BlockdevDriver) String() string {
	switch e {
	case BlockdevDriverBlkdebug:
		return "blkdebug"
	case BlockdevDriverBlklogwrites:
		return "blklogwrites"
	case BlockdevDriverBlkreplay:
		return "blkreplay"
	case BlockdevDriverBlkverify:
		return "blkverify"
	case BlockdevDriverBochs:
		return "bochs"
	case BlockdevDriverCloop:
		return "cloop"
	case BlockdevDriverCompress:
		return "compress"
	case BlockdevDriverCopyBeforeWrite:
		return "copy-before-write"
	case BlockdevDriverCopyOnRead:
		return "copy-on-read"
	case BlockdevDriverDmg:
		return "dmg"
	case BlockdevDriverFile:
		return "file"
	case BlockdevDriverSnapshotAccess:
		return "snapshot-access"
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
	case BlockdevDriverNvme:
		return "nvme"
	case BlockdevDriverParallels:
		return "parallels"
	case BlockdevDriverPreallocate:
		return "preallocate"
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
	default:
		return fmt.Sprintf("BlockdevDriver(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevDriver) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevDriverBlkdebug:
		return json.Marshal("blkdebug")
	case BlockdevDriverBlklogwrites:
		return json.Marshal("blklogwrites")
	case BlockdevDriverBlkreplay:
		return json.Marshal("blkreplay")
	case BlockdevDriverBlkverify:
		return json.Marshal("blkverify")
	case BlockdevDriverBochs:
		return json.Marshal("bochs")
	case BlockdevDriverCloop:
		return json.Marshal("cloop")
	case BlockdevDriverCompress:
		return json.Marshal("compress")
	case BlockdevDriverCopyBeforeWrite:
		return json.Marshal("copy-before-write")
	case BlockdevDriverCopyOnRead:
		return json.Marshal("copy-on-read")
	case BlockdevDriverDmg:
		return json.Marshal("dmg")
	case BlockdevDriverFile:
		return json.Marshal("file")
	case BlockdevDriverSnapshotAccess:
		return json.Marshal("snapshot-access")
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
	case BlockdevDriverNvme:
		return json.Marshal("nvme")
	case BlockdevDriverParallels:
		return json.Marshal("parallels")
	case BlockdevDriverPreallocate:
		return json.Marshal("preallocate")
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
	case "blklogwrites":
		*e = BlockdevDriverBlklogwrites
	case "blkreplay":
		*e = BlockdevDriverBlkreplay
	case "blkverify":
		*e = BlockdevDriverBlkverify
	case "bochs":
		*e = BlockdevDriverBochs
	case "cloop":
		*e = BlockdevDriverCloop
	case "compress":
		*e = BlockdevDriverCompress
	case "copy-before-write":
		*e = BlockdevDriverCopyBeforeWrite
	case "copy-on-read":
		*e = BlockdevDriverCopyOnRead
	case "dmg":
		*e = BlockdevDriverDmg
	case "file":
		*e = BlockdevDriverFile
	case "snapshot-access":
		*e = BlockdevDriverSnapshotAccess
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
	case "nvme":
		*e = BlockdevDriverNvme
	case "parallels":
		*e = BlockdevDriverParallels
	case "preallocate":
		*e = BlockdevDriverPreallocate
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
//   - BlockdevOptionsVariantBlkdebug
//   - BlockdevOptionsVariantBlklogwrites
//   - BlockdevOptionsVariantBlkreplay
//   - BlockdevOptionsVariantBlkverify
//   - BlockdevOptionsVariantBochs
//   - BlockdevOptionsVariantCloop
//   - BlockdevOptionsVariantCompress
//   - BlockdevOptionsVariantCopyBeforeWrite
//   - BlockdevOptionsVariantCopyOnRead
//   - BlockdevOptionsVariantDmg
//   - BlockdevOptionsVariantFile
//   - BlockdevOptionsVariantFTP
//   - BlockdevOptionsVariantFTPS
//   - BlockdevOptionsVariantGluster
//   - BlockdevOptionsVariantHostCdrom
//   - BlockdevOptionsVariantHostDevice
//   - BlockdevOptionsVariantHTTP
//   - BlockdevOptionsVariantHTTPS
//   - BlockdevOptionsVariantIscsi
//   - BlockdevOptionsVariantLUKS
//   - BlockdevOptionsVariantNBD
//   - BlockdevOptionsVariantNfs
//   - BlockdevOptionsVariantNullAIO
//   - BlockdevOptionsVariantNullCo
//   - BlockdevOptionsVariantNvme
//   - BlockdevOptionsVariantParallels
//   - BlockdevOptionsVariantPreallocate
//   - BlockdevOptionsVariantQcow
//   - BlockdevOptionsVariantQcow2
//   - BlockdevOptionsVariantQed
//   - BlockdevOptionsVariantQuorum
//   - BlockdevOptionsVariantRaw
//   - BlockdevOptionsVariantRbd
//   - BlockdevOptionsVariantReplication
//   - BlockdevOptionsVariantSnapshotAccess
//   - BlockdevOptionsVariantSSH
//   - BlockdevOptionsVariantThrottle
//   - BlockdevOptionsVariantVdi
//   - BlockdevOptionsVariantVhdx
//   - BlockdevOptionsVariantVMDK
//   - BlockdevOptionsVariantVpc
//   - BlockdevOptionsVariantVvfat
type BlockdevOptions interface {
	isBlockdevOptions()
}

// BlockdevOptionsVariantBlkdebug is an implementation of BlockdevOptions.
type BlockdevOptionsVariantBlkdebug struct {
	NodeName          *string                      `json:"node-name,omitempty"`
	Discard           *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache             *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly          *bool                        `json:"read-only,omitempty"`
	AutoReadOnly      *bool                        `json:"auto-read-only,omitempty"`
	ForceShare        *bool                        `json:"force-share,omitempty"`
	DetectZeroes      *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Image             BlockdevRef                  `json:"image"`
	Config            *string                      `json:"config,omitempty"`
	Align             *int64                       `json:"align,omitempty"`
	MaxTransfer       *int32                       `json:"max-transfer,omitempty"`
	OptWriteZero      *int32                       `json:"opt-write-zero,omitempty"`
	MaxWriteZero      *int32                       `json:"max-write-zero,omitempty"`
	OptDiscard        *int32                       `json:"opt-discard,omitempty"`
	MaxDiscard        *int32                       `json:"max-discard,omitempty"`
	InjectError       []BlkdebugInjectErrorOptions `json:"inject-error,omitempty"`
	SetState          []BlkdebugSetStateOptions    `json:"set-state,omitempty"`
	TakeChildPerms    []BlockPermission            `json:"take-child-perms,omitempty"`
	UnshareChildPerms []BlockPermission            `json:"unshare-child-perms,omitempty"`
}

func (BlockdevOptionsVariantBlkdebug) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantBlkdebug) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantBlkdebug
	}{
		BlockdevDriverBlkdebug,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantBlklogwrites is an implementation of BlockdevOptions.
type BlockdevOptionsVariantBlklogwrites struct {
	NodeName               *string                      `json:"node-name,omitempty"`
	Discard                *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache                  *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly               *bool                        `json:"read-only,omitempty"`
	AutoReadOnly           *bool                        `json:"auto-read-only,omitempty"`
	ForceShare             *bool                        `json:"force-share,omitempty"`
	DetectZeroes           *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File                   BlockdevRef                  `json:"file"`
	Log                    BlockdevRef                  `json:"log"`
	LogSectorSize          *uint32                      `json:"log-sector-size,omitempty"`
	LogAppend              *bool                        `json:"log-append,omitempty"`
	LogSuperUpdateInterval *uint64                      `json:"log-super-update-interval,omitempty"`
}

func (BlockdevOptionsVariantBlklogwrites) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantBlklogwrites) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantBlklogwrites
	}{
		BlockdevDriverBlklogwrites,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantBlkreplay is an implementation of BlockdevOptions.
type BlockdevOptionsVariantBlkreplay struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Image        BlockdevRef                  `json:"image"`
}

func (BlockdevOptionsVariantBlkreplay) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantBlkreplay) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantBlkreplay
	}{
		BlockdevDriverBlkreplay,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantBlkverify is an implementation of BlockdevOptions.
type BlockdevOptionsVariantBlkverify struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Test         BlockdevRef                  `json:"test"`
	Raw          BlockdevRef                  `json:"raw"`
}

func (BlockdevOptionsVariantBlkverify) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantBlkverify) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantBlkverify
	}{
		BlockdevDriverBlkverify,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantBochs is an implementation of BlockdevOptions.
type BlockdevOptionsVariantBochs struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVariantBochs) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantBochs) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantBochs
	}{
		BlockdevDriverBochs,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantCloop is an implementation of BlockdevOptions.
type BlockdevOptionsVariantCloop struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVariantCloop) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantCloop) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantCloop
	}{
		BlockdevDriverCloop,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantCompress is an implementation of BlockdevOptions.
type BlockdevOptionsVariantCompress struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVariantCompress) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantCompress) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantCompress
	}{
		BlockdevDriverCompress,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantCopyBeforeWrite is an implementation of BlockdevOptions.
type BlockdevOptionsVariantCopyBeforeWrite struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Target       BlockdevRef                  `json:"target"`
	Bitmap       *BlockDirtyBitmap            `json:"bitmap,omitempty"`
	OnCbwError   *OnCbwError                  `json:"on-cbw-error,omitempty"`
	CbwTimeout   *uint32                      `json:"cbw-timeout,omitempty"`
}

func (BlockdevOptionsVariantCopyBeforeWrite) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantCopyBeforeWrite) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantCopyBeforeWrite
	}{
		BlockdevDriverCopyBeforeWrite,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantCopyOnRead is an implementation of BlockdevOptions.
type BlockdevOptionsVariantCopyOnRead struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Bottom       *string                      `json:"bottom,omitempty"`
}

func (BlockdevOptionsVariantCopyOnRead) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantCopyOnRead) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantCopyOnRead
	}{
		BlockdevDriverCopyOnRead,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantDmg is an implementation of BlockdevOptions.
type BlockdevOptionsVariantDmg struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVariantDmg) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantDmg) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantDmg
	}{
		BlockdevDriverDmg,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantFile is an implementation of BlockdevOptions.
type BlockdevOptionsVariantFile struct {
	NodeName           *string                      `json:"node-name,omitempty"`
	Discard            *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache              *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly           *bool                        `json:"read-only,omitempty"`
	AutoReadOnly       *bool                        `json:"auto-read-only,omitempty"`
	ForceShare         *bool                        `json:"force-share,omitempty"`
	DetectZeroes       *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename           string                       `json:"filename"`
	PrManager          *string                      `json:"pr-manager,omitempty"`
	Locking            *OnOffAuto                   `json:"locking,omitempty"`
	AIO                *BlockdevAIOOptions          `json:"aio,omitempty"`
	AIOMaxBatch        *int64                       `json:"aio-max-batch,omitempty"`
	DropCache          *bool                        `json:"drop-cache,omitempty"`
	XCheckCacheDropped *bool                        `json:"x-check-cache-dropped,omitempty"`
}

func (BlockdevOptionsVariantFile) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantFile) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantFile
	}{
		BlockdevDriverFile,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantFTP is an implementation of BlockdevOptions.
type BlockdevOptionsVariantFTP struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
}

func (BlockdevOptionsVariantFTP) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantFTP) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantFTP
	}{
		BlockdevDriverFTP,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantFTPS is an implementation of BlockdevOptions.
type BlockdevOptionsVariantFTPS struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Sslverify    *bool                        `json:"sslverify,omitempty"`
}

func (BlockdevOptionsVariantFTPS) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantFTPS) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantFTPS
	}{
		BlockdevDriverFTPS,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantGluster is an implementation of BlockdevOptions.
type BlockdevOptionsVariantGluster struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Volume       string                       `json:"volume"`
	Path         string                       `json:"path"`
	Server       []SocketAddress              `json:"server"`
	Debug        *int64                       `json:"debug,omitempty"`
	Logfile      *string                      `json:"logfile,omitempty"`
}

func (BlockdevOptionsVariantGluster) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantGluster) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantGluster
	}{
		BlockdevDriverGluster,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantHostCdrom is an implementation of BlockdevOptions.
type BlockdevOptionsVariantHostCdrom struct {
	NodeName           *string                      `json:"node-name,omitempty"`
	Discard            *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache              *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly           *bool                        `json:"read-only,omitempty"`
	AutoReadOnly       *bool                        `json:"auto-read-only,omitempty"`
	ForceShare         *bool                        `json:"force-share,omitempty"`
	DetectZeroes       *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename           string                       `json:"filename"`
	PrManager          *string                      `json:"pr-manager,omitempty"`
	Locking            *OnOffAuto                   `json:"locking,omitempty"`
	AIO                *BlockdevAIOOptions          `json:"aio,omitempty"`
	AIOMaxBatch        *int64                       `json:"aio-max-batch,omitempty"`
	DropCache          *bool                        `json:"drop-cache,omitempty"`
	XCheckCacheDropped *bool                        `json:"x-check-cache-dropped,omitempty"`
}

func (BlockdevOptionsVariantHostCdrom) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantHostCdrom) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantHostCdrom
	}{
		BlockdevDriverHostCdrom,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantHostDevice is an implementation of BlockdevOptions.
type BlockdevOptionsVariantHostDevice struct {
	NodeName           *string                      `json:"node-name,omitempty"`
	Discard            *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache              *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly           *bool                        `json:"read-only,omitempty"`
	AutoReadOnly       *bool                        `json:"auto-read-only,omitempty"`
	ForceShare         *bool                        `json:"force-share,omitempty"`
	DetectZeroes       *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Filename           string                       `json:"filename"`
	PrManager          *string                      `json:"pr-manager,omitempty"`
	Locking            *OnOffAuto                   `json:"locking,omitempty"`
	AIO                *BlockdevAIOOptions          `json:"aio,omitempty"`
	AIOMaxBatch        *int64                       `json:"aio-max-batch,omitempty"`
	DropCache          *bool                        `json:"drop-cache,omitempty"`
	XCheckCacheDropped *bool                        `json:"x-check-cache-dropped,omitempty"`
}

func (BlockdevOptionsVariantHostDevice) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantHostDevice) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantHostDevice
	}{
		BlockdevDriverHostDevice,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantHTTP is an implementation of BlockdevOptions.
type BlockdevOptionsVariantHTTP struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Cookie       *string                      `json:"cookie,omitempty"`
	CookieSecret *string                      `json:"cookie-secret,omitempty"`
}

func (BlockdevOptionsVariantHTTP) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantHTTP) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantHTTP
	}{
		BlockdevDriverHTTP,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantHTTPS is an implementation of BlockdevOptions.
type BlockdevOptionsVariantHTTPS struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Cookie       *string                      `json:"cookie,omitempty"`
	Sslverify    *bool                        `json:"sslverify,omitempty"`
	CookieSecret *string                      `json:"cookie-secret,omitempty"`
}

func (BlockdevOptionsVariantHTTPS) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantHTTPS) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantHTTPS
	}{
		BlockdevDriverHTTPS,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantIscsi is an implementation of BlockdevOptions.
type BlockdevOptionsVariantIscsi struct {
	NodeName       *string                      `json:"node-name,omitempty"`
	Discard        *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache          *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly       *bool                        `json:"read-only,omitempty"`
	AutoReadOnly   *bool                        `json:"auto-read-only,omitempty"`
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

func (BlockdevOptionsVariantIscsi) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantIscsi) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantIscsi
	}{
		BlockdevDriverIscsi,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantLUKS is an implementation of BlockdevOptions.
type BlockdevOptionsVariantLUKS struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	KeySecret    *string                      `json:"key-secret,omitempty"`
}

func (BlockdevOptionsVariantLUKS) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantLUKS
	}{
		BlockdevDriverLUKS,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantNBD is an implementation of BlockdevOptions.
type BlockdevOptionsVariantNBD struct {
	NodeName       *string                      `json:"node-name,omitempty"`
	Discard        *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache          *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly       *bool                        `json:"read-only,omitempty"`
	AutoReadOnly   *bool                        `json:"auto-read-only,omitempty"`
	ForceShare     *bool                        `json:"force-share,omitempty"`
	DetectZeroes   *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Server         SocketAddress                `json:"server"`
	Export         *string                      `json:"export,omitempty"`
	TLSCreds       *string                      `json:"tls-creds,omitempty"`
	TLSHostname    *string                      `json:"tls-hostname,omitempty"`
	XDirtyBitmap   *string                      `json:"x-dirty-bitmap,omitempty"`
	ReconnectDelay *uint32                      `json:"reconnect-delay,omitempty"`
	OpenTimeout    *uint32                      `json:"open-timeout,omitempty"`
}

func (BlockdevOptionsVariantNBD) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantNBD) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantNBD
	}{
		BlockdevDriverNBD,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantNfs is an implementation of BlockdevOptions.
type BlockdevOptionsVariantNfs struct {
	NodeName      *string                      `json:"node-name,omitempty"`
	Discard       *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache         *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly      *bool                        `json:"read-only,omitempty"`
	AutoReadOnly  *bool                        `json:"auto-read-only,omitempty"`
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

func (BlockdevOptionsVariantNfs) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantNfs) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantNfs
	}{
		BlockdevDriverNfs,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantNullAIO is an implementation of BlockdevOptions.
type BlockdevOptionsVariantNullAIO struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Size         *int64                       `json:"size,omitempty"`
	LatencyNs    *uint64                      `json:"latency-ns,omitempty"`
	ReadZeroes   *bool                        `json:"read-zeroes,omitempty"`
}

func (BlockdevOptionsVariantNullAIO) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantNullAIO) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantNullAIO
	}{
		BlockdevDriverNullAIO,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantNullCo is an implementation of BlockdevOptions.
type BlockdevOptionsVariantNullCo struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Size         *int64                       `json:"size,omitempty"`
	LatencyNs    *uint64                      `json:"latency-ns,omitempty"`
	ReadZeroes   *bool                        `json:"read-zeroes,omitempty"`
}

func (BlockdevOptionsVariantNullCo) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantNullCo) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantNullCo
	}{
		BlockdevDriverNullCo,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantNvme is an implementation of BlockdevOptions.
type BlockdevOptionsVariantNvme struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Device       string                       `json:"device"`
	Namespace    int64                        `json:"namespace"`
}

func (BlockdevOptionsVariantNvme) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantNvme) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantNvme
	}{
		BlockdevDriverNvme,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantParallels is an implementation of BlockdevOptions.
type BlockdevOptionsVariantParallels struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVariantParallels) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantParallels) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantParallels
	}{
		BlockdevDriverParallels,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantPreallocate is an implementation of BlockdevOptions.
type BlockdevOptionsVariantPreallocate struct {
	NodeName      *string                      `json:"node-name,omitempty"`
	Discard       *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache         *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly      *bool                        `json:"read-only,omitempty"`
	AutoReadOnly  *bool                        `json:"auto-read-only,omitempty"`
	ForceShare    *bool                        `json:"force-share,omitempty"`
	DetectZeroes  *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	PreallocAlign *int64                       `json:"prealloc-align,omitempty"`
	PreallocSize  *int64                       `json:"prealloc-size,omitempty"`
}

func (BlockdevOptionsVariantPreallocate) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantPreallocate) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantPreallocate
	}{
		BlockdevDriverPreallocate,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantQcow is an implementation of BlockdevOptions.
type BlockdevOptionsVariantQcow struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Encrypt      *BlockdevQcowEncryption      `json:"encrypt,omitempty"`
}

func (BlockdevOptionsVariantQcow) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantQcow) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantQcow
	}{
		BlockdevDriverQcow,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantQcow2 is an implementation of BlockdevOptions.
type BlockdevOptionsVariantQcow2 struct {
	NodeName            *string                      `json:"node-name,omitempty"`
	Discard             *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache               *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly            *bool                        `json:"read-only,omitempty"`
	AutoReadOnly        *bool                        `json:"auto-read-only,omitempty"`
	ForceShare          *bool                        `json:"force-share,omitempty"`
	DetectZeroes        *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	LazyRefcounts       *bool                        `json:"lazy-refcounts,omitempty"`
	PassDiscardRequest  *bool                        `json:"pass-discard-request,omitempty"`
	PassDiscardSnapshot *bool                        `json:"pass-discard-snapshot,omitempty"`
	PassDiscardOther    *bool                        `json:"pass-discard-other,omitempty"`
	OverlapCheck        *Qcow2OverlapChecks          `json:"overlap-check,omitempty"`
	CacheSize           *int64                       `json:"cache-size,omitempty"`
	L2CacheSize         *int64                       `json:"l2-cache-size,omitempty"`
	L2CacheEntrySize    *int64                       `json:"l2-cache-entry-size,omitempty"`
	RefcountCacheSize   *int64                       `json:"refcount-cache-size,omitempty"`
	CacheCleanInterval  *int64                       `json:"cache-clean-interval,omitempty"`
	Encrypt             *BlockdevQcow2Encryption     `json:"encrypt,omitempty"`
	DataFile            *BlockdevRef                 `json:"data-file,omitempty"`
}

func (BlockdevOptionsVariantQcow2) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantQcow2) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantQcow2
	}{
		BlockdevDriverQcow2,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantQed is an implementation of BlockdevOptions.
type BlockdevOptionsVariantQed struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Backing      *BlockdevRefOrNull           `json:"backing,omitempty"`
}

func (BlockdevOptionsVariantQed) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantQed) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantQed
	}{
		BlockdevDriverQed,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantQuorum is an implementation of BlockdevOptions.
type BlockdevOptionsVariantQuorum struct {
	NodeName         *string                      `json:"node-name,omitempty"`
	Discard          *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache            *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly         *bool                        `json:"read-only,omitempty"`
	AutoReadOnly     *bool                        `json:"auto-read-only,omitempty"`
	ForceShare       *bool                        `json:"force-share,omitempty"`
	DetectZeroes     *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Blkverify        *bool                        `json:"blkverify,omitempty"`
	Children         []BlockdevRef                `json:"children"`
	VoteThreshold    int64                        `json:"vote-threshold"`
	RewriteCorrupted *bool                        `json:"rewrite-corrupted,omitempty"`
	ReadPattern      *QuorumReadPattern           `json:"read-pattern,omitempty"`
}

func (BlockdevOptionsVariantQuorum) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantQuorum) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantQuorum
	}{
		BlockdevDriverQuorum,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantRaw is an implementation of BlockdevOptions.
type BlockdevOptionsVariantRaw struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Offset       *int64                       `json:"offset,omitempty"`
	Size         *int64                       `json:"size,omitempty"`
}

func (BlockdevOptionsVariantRaw) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantRaw) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantRaw
	}{
		BlockdevDriverRaw,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantRbd is an implementation of BlockdevOptions.
type BlockdevOptionsVariantRbd struct {
	NodeName           *string                      `json:"node-name,omitempty"`
	Discard            *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache              *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly           *bool                        `json:"read-only,omitempty"`
	AutoReadOnly       *bool                        `json:"auto-read-only,omitempty"`
	ForceShare         *bool                        `json:"force-share,omitempty"`
	DetectZeroes       *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Pool               string                       `json:"pool"`
	Namespace          *string                      `json:"namespace,omitempty"`
	Image              string                       `json:"image"`
	Conf               *string                      `json:"conf,omitempty"`
	Snapshot           *string                      `json:"snapshot,omitempty"`
	Encrypt            *RbdEncryptionOptions        `json:"encrypt,omitempty"`
	User               *string                      `json:"user,omitempty"`
	AuthClientRequired []RbdAuthMode                `json:"auth-client-required,omitempty"`
	KeySecret          *string                      `json:"key-secret,omitempty"`
	Server             []InetSocketAddressBase      `json:"server,omitempty"`
}

func (BlockdevOptionsVariantRbd) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantRbd) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantRbd
	}{
		BlockdevDriverRbd,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantReplication is an implementation of BlockdevOptions.
type BlockdevOptionsVariantReplication struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Mode         ReplicationMode              `json:"mode"`
	TopID        *string                      `json:"top-id,omitempty"`
}

func (BlockdevOptionsVariantReplication) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantReplication) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantReplication
	}{
		BlockdevDriverReplication,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantSnapshotAccess is an implementation of BlockdevOptions.
type BlockdevOptionsVariantSnapshotAccess struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVariantSnapshotAccess) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantSnapshotAccess) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantSnapshotAccess
	}{
		BlockdevDriverSnapshotAccess,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantSSH is an implementation of BlockdevOptions.
type BlockdevOptionsVariantSSH struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Server       InetSocketAddress            `json:"server"`
	Path         string                       `json:"path"`
	User         *string                      `json:"user,omitempty"`
	HostKeyCheck *SSHHostKeyCheck             `json:"host-key-check,omitempty"`
}

func (BlockdevOptionsVariantSSH) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantSSH) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantSSH
	}{
		BlockdevDriverSSH,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantThrottle is an implementation of BlockdevOptions.
type BlockdevOptionsVariantThrottle struct {
	NodeName      *string                      `json:"node-name,omitempty"`
	Discard       *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache         *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly      *bool                        `json:"read-only,omitempty"`
	AutoReadOnly  *bool                        `json:"auto-read-only,omitempty"`
	ForceShare    *bool                        `json:"force-share,omitempty"`
	DetectZeroes  *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	ThrottleGroup string                       `json:"throttle-group"`
	File          BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVariantThrottle) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantThrottle) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantThrottle
	}{
		BlockdevDriverThrottle,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantVdi is an implementation of BlockdevOptions.
type BlockdevOptionsVariantVdi struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVariantVdi) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantVdi) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantVdi
	}{
		BlockdevDriverVdi,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantVhdx is an implementation of BlockdevOptions.
type BlockdevOptionsVariantVhdx struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVariantVhdx) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantVhdx) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantVhdx
	}{
		BlockdevDriverVhdx,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantVMDK is an implementation of BlockdevOptions.
type BlockdevOptionsVariantVMDK struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Backing      *BlockdevRefOrNull           `json:"backing,omitempty"`
}

func (BlockdevOptionsVariantVMDK) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantVMDK) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantVMDK
	}{
		BlockdevDriverVMDK,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantVpc is an implementation of BlockdevOptions.
type BlockdevOptionsVariantVpc struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	File         BlockdevRef                  `json:"file"`
}

func (BlockdevOptionsVariantVpc) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantVpc) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantVpc
	}{
		BlockdevDriverVpc,
		s,
	}
	return json.Marshal(v)
}

// BlockdevOptionsVariantVvfat is an implementation of BlockdevOptions.
type BlockdevOptionsVariantVvfat struct {
	NodeName     *string                      `json:"node-name,omitempty"`
	Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
	Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
	ReadOnly     *bool                        `json:"read-only,omitempty"`
	AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
	ForceShare   *bool                        `json:"force-share,omitempty"`
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
	Dir          string                       `json:"dir"`
	FatType      *int64                       `json:"fat-type,omitempty"`
	Floppy       *bool                        `json:"floppy,omitempty"`
	Label        *string                      `json:"label,omitempty"`
	Rw           *bool                        `json:"rw,omitempty"`
}

func (BlockdevOptionsVariantVvfat) isBlockdevOptions() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevOptionsVariantVvfat) MarshalJSON() ([]byte, error) {
	v := struct {
		Driver BlockdevDriver `json:"driver"`
		BlockdevOptionsVariantVvfat
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
	case BlockdevDriverBlkdebug:
		var ret BlockdevOptionsVariantBlkdebug
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverBlklogwrites:
		var ret BlockdevOptionsVariantBlklogwrites
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverBlkreplay:
		var ret BlockdevOptionsVariantBlkreplay
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverBlkverify:
		var ret BlockdevOptionsVariantBlkverify
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverBochs:
		var ret BlockdevOptionsVariantBochs
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverCloop:
		var ret BlockdevOptionsVariantCloop
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverCompress:
		var ret BlockdevOptionsVariantCompress
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverCopyBeforeWrite:
		var ret BlockdevOptionsVariantCopyBeforeWrite
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverCopyOnRead:
		var ret BlockdevOptionsVariantCopyOnRead
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverDmg:
		var ret BlockdevOptionsVariantDmg
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverFile:
		var ret BlockdevOptionsVariantFile
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverFTP:
		var ret BlockdevOptionsVariantFTP
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverFTPS:
		var ret BlockdevOptionsVariantFTPS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverGluster:
		var ret BlockdevOptionsVariantGluster
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverHostCdrom:
		var ret BlockdevOptionsVariantHostCdrom
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverHostDevice:
		var ret BlockdevOptionsVariantHostDevice
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverHTTP:
		var ret BlockdevOptionsVariantHTTP
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverHTTPS:
		var ret BlockdevOptionsVariantHTTPS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverIscsi:
		var ret BlockdevOptionsVariantIscsi
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverLUKS:
		var ret BlockdevOptionsVariantLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNBD:
		var ret BlockdevOptionsVariantNBD
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNfs:
		var ret BlockdevOptionsVariantNfs
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNullAIO:
		var ret BlockdevOptionsVariantNullAIO
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNullCo:
		var ret BlockdevOptionsVariantNullCo
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverNvme:
		var ret BlockdevOptionsVariantNvme
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverParallels:
		var ret BlockdevOptionsVariantParallels
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverPreallocate:
		var ret BlockdevOptionsVariantPreallocate
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQcow:
		var ret BlockdevOptionsVariantQcow
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQcow2:
		var ret BlockdevOptionsVariantQcow2
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQed:
		var ret BlockdevOptionsVariantQed
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverQuorum:
		var ret BlockdevOptionsVariantQuorum
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverRaw:
		var ret BlockdevOptionsVariantRaw
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverRbd:
		var ret BlockdevOptionsVariantRbd
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverReplication:
		var ret BlockdevOptionsVariantReplication
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverSnapshotAccess:
		var ret BlockdevOptionsVariantSnapshotAccess
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverSSH:
		var ret BlockdevOptionsVariantSSH
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverThrottle:
		var ret BlockdevOptionsVariantThrottle
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVdi:
		var ret BlockdevOptionsVariantVdi
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVhdx:
		var ret BlockdevOptionsVariantVhdx
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVMDK:
		var ret BlockdevOptionsVariantVMDK
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVpc:
		var ret BlockdevOptionsVariantVpc
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevDriverVvfat:
		var ret BlockdevOptionsVariantVvfat
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union BlockdevOptions", v.Driver)
	}
}

// BlockdevOptionsGluster -> BlockdevOptionsGluster (struct)

// BlockdevOptionsGluster implements the "BlockdevOptionsGluster" QMP API type.
type BlockdevOptionsGluster struct {
	Volume  string          `json:"volume"`
	Path    string          `json:"path"`
	Server  []SocketAddress `json:"server"`
	Debug   *int64          `json:"debug,omitempty"`
	Logfile *string         `json:"logfile,omitempty"`
}

// BlockdevOptionsNfs -> BlockdevOptionsNfs (struct)

// BlockdevOptionsNfs implements the "BlockdevOptionsNfs" QMP API type.
type BlockdevOptionsNfs struct {
	Server        NfsServer `json:"server"`
	Path          string    `json:"path"`
	User          *int64    `json:"user,omitempty"`
	Group         *int64    `json:"group,omitempty"`
	TCPSynCount   *int64    `json:"tcp-syn-count,omitempty"`
	ReadaheadSize *int64    `json:"readahead-size,omitempty"`
	PageCacheSize *int64    `json:"page-cache-size,omitempty"`
	Debug         *int64    `json:"debug,omitempty"`
}

// BlockdevOptionsRbd -> BlockdevOptionsRbd (struct)

// BlockdevOptionsRbd implements the "BlockdevOptionsRbd" QMP API type.
type BlockdevOptionsRbd struct {
	Pool               string                  `json:"pool"`
	Namespace          *string                 `json:"namespace,omitempty"`
	Image              string                  `json:"image"`
	Conf               *string                 `json:"conf,omitempty"`
	Snapshot           *string                 `json:"snapshot,omitempty"`
	Encrypt            *RbdEncryptionOptions   `json:"encrypt,omitempty"`
	User               *string                 `json:"user,omitempty"`
	AuthClientRequired []RbdAuthMode           `json:"auth-client-required,omitempty"`
	KeySecret          *string                 `json:"key-secret,omitempty"`
	Server             []InetSocketAddressBase `json:"server,omitempty"`
}

// BlockdevOptionsSsh -> BlockdevOptionsSSH (struct)

// BlockdevOptionsSSH implements the "BlockdevOptionsSsh" QMP API type.
type BlockdevOptionsSSH struct {
	Server       InetSocketAddress `json:"server"`
	Path         string            `json:"path"`
	User         *string           `json:"user,omitempty"`
	HostKeyCheck *SSHHostKeyCheck  `json:"host-key-check,omitempty"`
}

// BlockdevQcow2Encryption -> BlockdevQcow2Encryption (flat union)

// BlockdevQcow2Encryption implements the "BlockdevQcow2Encryption" QMP API type.
//
// Can be one of:
//   - BlockdevQcow2EncryptionVariantAes
//   - BlockdevQcow2EncryptionVariantLUKS
type BlockdevQcow2Encryption interface {
	isBlockdevQcow2Encryption()
}

// BlockdevQcow2EncryptionVariantAes is an implementation of BlockdevQcow2Encryption.
type BlockdevQcow2EncryptionVariantAes struct {
	KeySecret *string `json:"key-secret,omitempty"`
}

func (BlockdevQcow2EncryptionVariantAes) isBlockdevQcow2Encryption() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevQcow2EncryptionVariantAes) MarshalJSON() ([]byte, error) {
	v := struct {
		Format BlockdevQcow2EncryptionFormat `json:"format"`
		BlockdevQcow2EncryptionVariantAes
	}{
		BlockdevQcow2EncryptionFormatAes,
		s,
	}
	return json.Marshal(v)
}

// BlockdevQcow2EncryptionVariantLUKS is an implementation of BlockdevQcow2Encryption.
type BlockdevQcow2EncryptionVariantLUKS struct {
	KeySecret *string `json:"key-secret,omitempty"`
}

func (BlockdevQcow2EncryptionVariantLUKS) isBlockdevQcow2Encryption() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevQcow2EncryptionVariantLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Format BlockdevQcow2EncryptionFormat `json:"format"`
		BlockdevQcow2EncryptionVariantLUKS
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
		var ret BlockdevQcow2EncryptionVariantAes
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case BlockdevQcow2EncryptionFormatLUKS:
		var ret BlockdevQcow2EncryptionVariantLUKS
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

// BlockdevQcow2Version -> BlockdevQcow2Version (enum)

// BlockdevQcow2Version implements the "BlockdevQcow2Version" QMP API type.
type BlockdevQcow2Version int

// Known values of BlockdevQcow2Version.
const (
	BlockdevQcow2VersionV2 BlockdevQcow2Version = iota
	BlockdevQcow2VersionV3
)

// String implements fmt.Stringer.
func (e BlockdevQcow2Version) String() string {
	switch e {
	case BlockdevQcow2VersionV2:
		return "v2"
	case BlockdevQcow2VersionV3:
		return "v3"
	default:
		return fmt.Sprintf("BlockdevQcow2Version(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevQcow2Version) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevQcow2VersionV2:
		return json.Marshal("v2")
	case BlockdevQcow2VersionV3:
		return json.Marshal("v3")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevQcow2Version", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevQcow2Version) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "v2":
		*e = BlockdevQcow2VersionV2
	case "v3":
		*e = BlockdevQcow2VersionV3
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevQcow2Version", s)
	}
	return nil
}

// BlockdevQcowEncryption -> BlockdevQcowEncryption (flat union)

// BlockdevQcowEncryption implements the "BlockdevQcowEncryption" QMP API type.
//
// Can be one of:
//   - BlockdevQcowEncryptionVariantAes
type BlockdevQcowEncryption interface {
	isBlockdevQcowEncryption()
}

// BlockdevQcowEncryptionVariantAes is an implementation of BlockdevQcowEncryption.
type BlockdevQcowEncryptionVariantAes struct {
	KeySecret *string `json:"key-secret,omitempty"`
}

func (BlockdevQcowEncryptionVariantAes) isBlockdevQcowEncryption() {}

// MarshalJSON implements json.Marshaler.
func (s BlockdevQcowEncryptionVariantAes) MarshalJSON() ([]byte, error) {
	v := struct {
		Format BlockdevQcowEncryptionFormat `json:"format"`
		BlockdevQcowEncryptionVariantAes
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
		var ret BlockdevQcowEncryptionVariantAes
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
//   - BlockdevRefVariantDefinition
//   - BlockdevRefVariantReference
type BlockdevRef interface {
	isBlockdevRef()
}

// BlockdevRefVariantDefinition is an implementation of BlockdevRef
type BlockdevRefVariantDefinition BlockdevOptions

func (BlockdevOptionsVariantBlkdebug) isBlockdevRef()        {}
func (BlockdevOptionsVariantBlklogwrites) isBlockdevRef()    {}
func (BlockdevOptionsVariantBlkreplay) isBlockdevRef()       {}
func (BlockdevOptionsVariantBlkverify) isBlockdevRef()       {}
func (BlockdevOptionsVariantBochs) isBlockdevRef()           {}
func (BlockdevOptionsVariantCloop) isBlockdevRef()           {}
func (BlockdevOptionsVariantCompress) isBlockdevRef()        {}
func (BlockdevOptionsVariantCopyBeforeWrite) isBlockdevRef() {}
func (BlockdevOptionsVariantCopyOnRead) isBlockdevRef()      {}
func (BlockdevOptionsVariantDmg) isBlockdevRef()             {}
func (BlockdevOptionsVariantFile) isBlockdevRef()            {}
func (BlockdevOptionsVariantFTP) isBlockdevRef()             {}
func (BlockdevOptionsVariantFTPS) isBlockdevRef()            {}
func (BlockdevOptionsVariantGluster) isBlockdevRef()         {}
func (BlockdevOptionsVariantHostCdrom) isBlockdevRef()       {}
func (BlockdevOptionsVariantHostDevice) isBlockdevRef()      {}
func (BlockdevOptionsVariantHTTP) isBlockdevRef()            {}
func (BlockdevOptionsVariantHTTPS) isBlockdevRef()           {}
func (BlockdevOptionsVariantIscsi) isBlockdevRef()           {}
func (BlockdevOptionsVariantLUKS) isBlockdevRef()            {}
func (BlockdevOptionsVariantNBD) isBlockdevRef()             {}
func (BlockdevOptionsVariantNfs) isBlockdevRef()             {}
func (BlockdevOptionsVariantNullAIO) isBlockdevRef()         {}
func (BlockdevOptionsVariantNullCo) isBlockdevRef()          {}
func (BlockdevOptionsVariantNvme) isBlockdevRef()            {}
func (BlockdevOptionsVariantParallels) isBlockdevRef()       {}
func (BlockdevOptionsVariantPreallocate) isBlockdevRef()     {}
func (BlockdevOptionsVariantQcow) isBlockdevRef()            {}
func (BlockdevOptionsVariantQcow2) isBlockdevRef()           {}
func (BlockdevOptionsVariantQed) isBlockdevRef()             {}
func (BlockdevOptionsVariantQuorum) isBlockdevRef()          {}
func (BlockdevOptionsVariantRaw) isBlockdevRef()             {}
func (BlockdevOptionsVariantRbd) isBlockdevRef()             {}
func (BlockdevOptionsVariantReplication) isBlockdevRef()     {}
func (BlockdevOptionsVariantSnapshotAccess) isBlockdevRef()  {}
func (BlockdevOptionsVariantSSH) isBlockdevRef()             {}
func (BlockdevOptionsVariantThrottle) isBlockdevRef()        {}
func (BlockdevOptionsVariantVdi) isBlockdevRef()             {}
func (BlockdevOptionsVariantVhdx) isBlockdevRef()            {}
func (BlockdevOptionsVariantVMDK) isBlockdevRef()            {}
func (BlockdevOptionsVariantVpc) isBlockdevRef()             {}
func (BlockdevOptionsVariantVvfat) isBlockdevRef()           {}

// BlockdevRefVariantReference is an implementation of BlockdevRef
type BlockdevRefVariantReference string

func (BlockdevRefVariantReference) isBlockdevRef() {}

func decodeBlockdevRef(bs json.RawMessage) (BlockdevRef, error) {

	var reference BlockdevRefVariantReference
	if err := json.Unmarshal([]byte(bs), &reference); err == nil {
		return reference, nil
	}

	if definition, err := decodeBlockdevOptions([]byte(bs)); err == nil {
		switch impl := definition.(type) {
		case BlockdevOptionsVariantBlkdebug:
			return impl, nil
		case BlockdevOptionsVariantBlklogwrites:
			return impl, nil
		case BlockdevOptionsVariantBlkreplay:
			return impl, nil
		case BlockdevOptionsVariantBlkverify:
			return impl, nil
		case BlockdevOptionsVariantBochs:
			return impl, nil
		case BlockdevOptionsVariantCloop:
			return impl, nil
		case BlockdevOptionsVariantCompress:
			return impl, nil
		case BlockdevOptionsVariantCopyBeforeWrite:
			return impl, nil
		case BlockdevOptionsVariantCopyOnRead:
			return impl, nil
		case BlockdevOptionsVariantDmg:
			return impl, nil
		case BlockdevOptionsVariantFile:
			return impl, nil
		case BlockdevOptionsVariantFTP:
			return impl, nil
		case BlockdevOptionsVariantFTPS:
			return impl, nil
		case BlockdevOptionsVariantGluster:
			return impl, nil
		case BlockdevOptionsVariantHostCdrom:
			return impl, nil
		case BlockdevOptionsVariantHostDevice:
			return impl, nil
		case BlockdevOptionsVariantHTTP:
			return impl, nil
		case BlockdevOptionsVariantHTTPS:
			return impl, nil
		case BlockdevOptionsVariantIscsi:
			return impl, nil
		case BlockdevOptionsVariantLUKS:
			return impl, nil
		case BlockdevOptionsVariantNBD:
			return impl, nil
		case BlockdevOptionsVariantNfs:
			return impl, nil
		case BlockdevOptionsVariantNullAIO:
			return impl, nil
		case BlockdevOptionsVariantNullCo:
			return impl, nil
		case BlockdevOptionsVariantNvme:
			return impl, nil
		case BlockdevOptionsVariantParallels:
			return impl, nil
		case BlockdevOptionsVariantPreallocate:
			return impl, nil
		case BlockdevOptionsVariantQcow:
			return impl, nil
		case BlockdevOptionsVariantQcow2:
			return impl, nil
		case BlockdevOptionsVariantQed:
			return impl, nil
		case BlockdevOptionsVariantQuorum:
			return impl, nil
		case BlockdevOptionsVariantRaw:
			return impl, nil
		case BlockdevOptionsVariantRbd:
			return impl, nil
		case BlockdevOptionsVariantReplication:
			return impl, nil
		case BlockdevOptionsVariantSnapshotAccess:
			return impl, nil
		case BlockdevOptionsVariantSSH:
			return impl, nil
		case BlockdevOptionsVariantThrottle:
			return impl, nil
		case BlockdevOptionsVariantVdi:
			return impl, nil
		case BlockdevOptionsVariantVhdx:
			return impl, nil
		case BlockdevOptionsVariantVMDK:
			return impl, nil
		case BlockdevOptionsVariantVpc:
			return impl, nil
		case BlockdevOptionsVariantVvfat:
			return impl, nil
		}
	}
	return nil, fmt.Errorf("unable to decode %q as a BlockdevRef", string(bs))
}

// BlockdevRefOrNull -> BlockdevRefOrNull (alternate)

// BlockdevRefOrNull implements the "BlockdevRefOrNull" QMP API type.
//
// Can be one of:
//   - BlockdevRefOrNullVariantDefinition
//   - BlockdevRefOrNullVariantNull
//   - BlockdevRefOrNullVariantReference
type BlockdevRefOrNull interface {
	isBlockdevRefOrNull()
}

// BlockdevRefOrNullVariantDefinition is an implementation of BlockdevRefOrNull
type BlockdevRefOrNullVariantDefinition BlockdevOptions

func (BlockdevOptionsVariantBlkdebug) isBlockdevRefOrNull()        {}
func (BlockdevOptionsVariantBlklogwrites) isBlockdevRefOrNull()    {}
func (BlockdevOptionsVariantBlkreplay) isBlockdevRefOrNull()       {}
func (BlockdevOptionsVariantBlkverify) isBlockdevRefOrNull()       {}
func (BlockdevOptionsVariantBochs) isBlockdevRefOrNull()           {}
func (BlockdevOptionsVariantCloop) isBlockdevRefOrNull()           {}
func (BlockdevOptionsVariantCompress) isBlockdevRefOrNull()        {}
func (BlockdevOptionsVariantCopyBeforeWrite) isBlockdevRefOrNull() {}
func (BlockdevOptionsVariantCopyOnRead) isBlockdevRefOrNull()      {}
func (BlockdevOptionsVariantDmg) isBlockdevRefOrNull()             {}
func (BlockdevOptionsVariantFile) isBlockdevRefOrNull()            {}
func (BlockdevOptionsVariantFTP) isBlockdevRefOrNull()             {}
func (BlockdevOptionsVariantFTPS) isBlockdevRefOrNull()            {}
func (BlockdevOptionsVariantGluster) isBlockdevRefOrNull()         {}
func (BlockdevOptionsVariantHostCdrom) isBlockdevRefOrNull()       {}
func (BlockdevOptionsVariantHostDevice) isBlockdevRefOrNull()      {}
func (BlockdevOptionsVariantHTTP) isBlockdevRefOrNull()            {}
func (BlockdevOptionsVariantHTTPS) isBlockdevRefOrNull()           {}
func (BlockdevOptionsVariantIscsi) isBlockdevRefOrNull()           {}
func (BlockdevOptionsVariantLUKS) isBlockdevRefOrNull()            {}
func (BlockdevOptionsVariantNBD) isBlockdevRefOrNull()             {}
func (BlockdevOptionsVariantNfs) isBlockdevRefOrNull()             {}
func (BlockdevOptionsVariantNullAIO) isBlockdevRefOrNull()         {}
func (BlockdevOptionsVariantNullCo) isBlockdevRefOrNull()          {}
func (BlockdevOptionsVariantNvme) isBlockdevRefOrNull()            {}
func (BlockdevOptionsVariantParallels) isBlockdevRefOrNull()       {}
func (BlockdevOptionsVariantPreallocate) isBlockdevRefOrNull()     {}
func (BlockdevOptionsVariantQcow) isBlockdevRefOrNull()            {}
func (BlockdevOptionsVariantQcow2) isBlockdevRefOrNull()           {}
func (BlockdevOptionsVariantQed) isBlockdevRefOrNull()             {}
func (BlockdevOptionsVariantQuorum) isBlockdevRefOrNull()          {}
func (BlockdevOptionsVariantRaw) isBlockdevRefOrNull()             {}
func (BlockdevOptionsVariantRbd) isBlockdevRefOrNull()             {}
func (BlockdevOptionsVariantReplication) isBlockdevRefOrNull()     {}
func (BlockdevOptionsVariantSnapshotAccess) isBlockdevRefOrNull()  {}
func (BlockdevOptionsVariantSSH) isBlockdevRefOrNull()             {}
func (BlockdevOptionsVariantThrottle) isBlockdevRefOrNull()        {}
func (BlockdevOptionsVariantVdi) isBlockdevRefOrNull()             {}
func (BlockdevOptionsVariantVhdx) isBlockdevRefOrNull()            {}
func (BlockdevOptionsVariantVMDK) isBlockdevRefOrNull()            {}
func (BlockdevOptionsVariantVpc) isBlockdevRefOrNull()             {}
func (BlockdevOptionsVariantVvfat) isBlockdevRefOrNull()           {}

// BlockdevRefOrNullVariantNull is a JSON null type, so it must
// also implement the isNullable interface.
type BlockdevRefOrNullVariantNull struct{}

func (BlockdevRefOrNullVariantNull) isNull() bool         { return true }
func (BlockdevRefOrNullVariantNull) isBlockdevRefOrNull() {}

// BlockdevRefOrNullVariantReference is an implementation of BlockdevRefOrNull
type BlockdevRefOrNullVariantReference string

func (BlockdevRefOrNullVariantReference) isBlockdevRefOrNull() {}

func decodeBlockdevRefOrNull(bs json.RawMessage) (BlockdevRefOrNull, error) {

	// Always try unmarshalling for nil first if it's an option
	// because other types could unmarshal successfully in the case
	// where a Null json type was provided.
	var null *int
	if err := json.Unmarshal([]byte(bs), &null); err == nil {
		if null == nil {
			return BlockdevRefOrNullVariantNull{}, nil
		}
	}
	var reference BlockdevRefOrNullVariantReference
	if err := json.Unmarshal([]byte(bs), &reference); err == nil {
		return reference, nil
	}

	if definition, err := decodeBlockdevOptions([]byte(bs)); err == nil {
		switch impl := definition.(type) {
		case BlockdevOptionsVariantBlkdebug:
			return impl, nil
		case BlockdevOptionsVariantBlklogwrites:
			return impl, nil
		case BlockdevOptionsVariantBlkreplay:
			return impl, nil
		case BlockdevOptionsVariantBlkverify:
			return impl, nil
		case BlockdevOptionsVariantBochs:
			return impl, nil
		case BlockdevOptionsVariantCloop:
			return impl, nil
		case BlockdevOptionsVariantCompress:
			return impl, nil
		case BlockdevOptionsVariantCopyBeforeWrite:
			return impl, nil
		case BlockdevOptionsVariantCopyOnRead:
			return impl, nil
		case BlockdevOptionsVariantDmg:
			return impl, nil
		case BlockdevOptionsVariantFile:
			return impl, nil
		case BlockdevOptionsVariantFTP:
			return impl, nil
		case BlockdevOptionsVariantFTPS:
			return impl, nil
		case BlockdevOptionsVariantGluster:
			return impl, nil
		case BlockdevOptionsVariantHostCdrom:
			return impl, nil
		case BlockdevOptionsVariantHostDevice:
			return impl, nil
		case BlockdevOptionsVariantHTTP:
			return impl, nil
		case BlockdevOptionsVariantHTTPS:
			return impl, nil
		case BlockdevOptionsVariantIscsi:
			return impl, nil
		case BlockdevOptionsVariantLUKS:
			return impl, nil
		case BlockdevOptionsVariantNBD:
			return impl, nil
		case BlockdevOptionsVariantNfs:
			return impl, nil
		case BlockdevOptionsVariantNullAIO:
			return impl, nil
		case BlockdevOptionsVariantNullCo:
			return impl, nil
		case BlockdevOptionsVariantNvme:
			return impl, nil
		case BlockdevOptionsVariantParallels:
			return impl, nil
		case BlockdevOptionsVariantPreallocate:
			return impl, nil
		case BlockdevOptionsVariantQcow:
			return impl, nil
		case BlockdevOptionsVariantQcow2:
			return impl, nil
		case BlockdevOptionsVariantQed:
			return impl, nil
		case BlockdevOptionsVariantQuorum:
			return impl, nil
		case BlockdevOptionsVariantRaw:
			return impl, nil
		case BlockdevOptionsVariantRbd:
			return impl, nil
		case BlockdevOptionsVariantReplication:
			return impl, nil
		case BlockdevOptionsVariantSnapshotAccess:
			return impl, nil
		case BlockdevOptionsVariantSSH:
			return impl, nil
		case BlockdevOptionsVariantThrottle:
			return impl, nil
		case BlockdevOptionsVariantVdi:
			return impl, nil
		case BlockdevOptionsVariantVhdx:
			return impl, nil
		case BlockdevOptionsVariantVMDK:
			return impl, nil
		case BlockdevOptionsVariantVpc:
			return impl, nil
		case BlockdevOptionsVariantVvfat:
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

// BlockdevVhdxSubformat -> BlockdevVhdxSubformat (enum)

// BlockdevVhdxSubformat implements the "BlockdevVhdxSubformat" QMP API type.
type BlockdevVhdxSubformat int

// Known values of BlockdevVhdxSubformat.
const (
	BlockdevVhdxSubformatDynamic BlockdevVhdxSubformat = iota
	BlockdevVhdxSubformatFixed
)

// String implements fmt.Stringer.
func (e BlockdevVhdxSubformat) String() string {
	switch e {
	case BlockdevVhdxSubformatDynamic:
		return "dynamic"
	case BlockdevVhdxSubformatFixed:
		return "fixed"
	default:
		return fmt.Sprintf("BlockdevVhdxSubformat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevVhdxSubformat) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevVhdxSubformatDynamic:
		return json.Marshal("dynamic")
	case BlockdevVhdxSubformatFixed:
		return json.Marshal("fixed")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevVhdxSubformat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevVhdxSubformat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "dynamic":
		*e = BlockdevVhdxSubformatDynamic
	case "fixed":
		*e = BlockdevVhdxSubformatFixed
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevVhdxSubformat", s)
	}
	return nil
}

// BlockdevVmdkAdapterType -> BlockdevVMDKAdapterType (enum)

// BlockdevVMDKAdapterType implements the "BlockdevVmdkAdapterType" QMP API type.
type BlockdevVMDKAdapterType int

// Known values of BlockdevVMDKAdapterType.
const (
	BlockdevVMDKAdapterTypeIde BlockdevVMDKAdapterType = iota
	BlockdevVMDKAdapterTypeBuslogic
	BlockdevVMDKAdapterTypeLsilogic
	BlockdevVMDKAdapterTypeLegacyEsx
)

// String implements fmt.Stringer.
func (e BlockdevVMDKAdapterType) String() string {
	switch e {
	case BlockdevVMDKAdapterTypeIde:
		return "ide"
	case BlockdevVMDKAdapterTypeBuslogic:
		return "buslogic"
	case BlockdevVMDKAdapterTypeLsilogic:
		return "lsilogic"
	case BlockdevVMDKAdapterTypeLegacyEsx:
		return "legacyESX"
	default:
		return fmt.Sprintf("BlockdevVMDKAdapterType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevVMDKAdapterType) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevVMDKAdapterTypeIde:
		return json.Marshal("ide")
	case BlockdevVMDKAdapterTypeBuslogic:
		return json.Marshal("buslogic")
	case BlockdevVMDKAdapterTypeLsilogic:
		return json.Marshal("lsilogic")
	case BlockdevVMDKAdapterTypeLegacyEsx:
		return json.Marshal("legacyESX")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevVMDKAdapterType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevVMDKAdapterType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "ide":
		*e = BlockdevVMDKAdapterTypeIde
	case "buslogic":
		*e = BlockdevVMDKAdapterTypeBuslogic
	case "lsilogic":
		*e = BlockdevVMDKAdapterTypeLsilogic
	case "legacyESX":
		*e = BlockdevVMDKAdapterTypeLegacyEsx
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevVMDKAdapterType", s)
	}
	return nil
}

// BlockdevVmdkSubformat -> BlockdevVMDKSubformat (enum)

// BlockdevVMDKSubformat implements the "BlockdevVmdkSubformat" QMP API type.
type BlockdevVMDKSubformat int

// Known values of BlockdevVMDKSubformat.
const (
	BlockdevVMDKSubformatMonolithicSparse BlockdevVMDKSubformat = iota
	BlockdevVMDKSubformatMonolithicFlat
	BlockdevVMDKSubformatTwoGbMaxExtentSparse
	BlockdevVMDKSubformatTwoGbMaxExtentFlat
	BlockdevVMDKSubformatStreamOptimized
)

// String implements fmt.Stringer.
func (e BlockdevVMDKSubformat) String() string {
	switch e {
	case BlockdevVMDKSubformatMonolithicSparse:
		return "monolithicSparse"
	case BlockdevVMDKSubformatMonolithicFlat:
		return "monolithicFlat"
	case BlockdevVMDKSubformatTwoGbMaxExtentSparse:
		return "twoGbMaxExtentSparse"
	case BlockdevVMDKSubformatTwoGbMaxExtentFlat:
		return "twoGbMaxExtentFlat"
	case BlockdevVMDKSubformatStreamOptimized:
		return "streamOptimized"
	default:
		return fmt.Sprintf("BlockdevVMDKSubformat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevVMDKSubformat) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevVMDKSubformatMonolithicSparse:
		return json.Marshal("monolithicSparse")
	case BlockdevVMDKSubformatMonolithicFlat:
		return json.Marshal("monolithicFlat")
	case BlockdevVMDKSubformatTwoGbMaxExtentSparse:
		return json.Marshal("twoGbMaxExtentSparse")
	case BlockdevVMDKSubformatTwoGbMaxExtentFlat:
		return json.Marshal("twoGbMaxExtentFlat")
	case BlockdevVMDKSubformatStreamOptimized:
		return json.Marshal("streamOptimized")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevVMDKSubformat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevVMDKSubformat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "monolithicSparse":
		*e = BlockdevVMDKSubformatMonolithicSparse
	case "monolithicFlat":
		*e = BlockdevVMDKSubformatMonolithicFlat
	case "twoGbMaxExtentSparse":
		*e = BlockdevVMDKSubformatTwoGbMaxExtentSparse
	case "twoGbMaxExtentFlat":
		*e = BlockdevVMDKSubformatTwoGbMaxExtentFlat
	case "streamOptimized":
		*e = BlockdevVMDKSubformatStreamOptimized
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevVMDKSubformat", s)
	}
	return nil
}

// BlockdevVpcSubformat -> BlockdevVpcSubformat (enum)

// BlockdevVpcSubformat implements the "BlockdevVpcSubformat" QMP API type.
type BlockdevVpcSubformat int

// Known values of BlockdevVpcSubformat.
const (
	BlockdevVpcSubformatDynamic BlockdevVpcSubformat = iota
	BlockdevVpcSubformatFixed
)

// String implements fmt.Stringer.
func (e BlockdevVpcSubformat) String() string {
	switch e {
	case BlockdevVpcSubformatDynamic:
		return "dynamic"
	case BlockdevVpcSubformatFixed:
		return "fixed"
	default:
		return fmt.Sprintf("BlockdevVpcSubformat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e BlockdevVpcSubformat) MarshalJSON() ([]byte, error) {
	switch e {
	case BlockdevVpcSubformatDynamic:
		return json.Marshal("dynamic")
	case BlockdevVpcSubformatFixed:
		return json.Marshal("fixed")
	default:
		return nil, fmt.Errorf("unknown enum value %q for BlockdevVpcSubformat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *BlockdevVpcSubformat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "dynamic":
		*e = BlockdevVpcSubformatDynamic
	case "fixed":
		*e = BlockdevVpcSubformatFixed
	default:
		return fmt.Errorf("unknown enum value %q for BlockdevVpcSubformat", s)
	}
	return nil
}

// COLOExitReason -> ColoExitReason (enum)

// ColoExitReason implements the "COLOExitReason" QMP API type.
type ColoExitReason int

// Known values of ColoExitReason.
const (
	ColoExitReasonNone ColoExitReason = iota
	ColoExitReasonRequest
	ColoExitReasonError
	ColoExitReasonProcessing
)

// String implements fmt.Stringer.
func (e ColoExitReason) String() string {
	switch e {
	case ColoExitReasonNone:
		return "none"
	case ColoExitReasonRequest:
		return "request"
	case ColoExitReasonError:
		return "error"
	case ColoExitReasonProcessing:
		return "processing"
	default:
		return fmt.Sprintf("ColoExitReason(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ColoExitReason) MarshalJSON() ([]byte, error) {
	switch e {
	case ColoExitReasonNone:
		return json.Marshal("none")
	case ColoExitReasonRequest:
		return json.Marshal("request")
	case ColoExitReasonError:
		return json.Marshal("error")
	case ColoExitReasonProcessing:
		return json.Marshal("processing")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ColoExitReason", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ColoExitReason) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = ColoExitReasonNone
	case "request":
		*e = ColoExitReasonRequest
	case "error":
		*e = ColoExitReasonError
	case "processing":
		*e = ColoExitReasonProcessing
	default:
		return fmt.Errorf("unknown enum value %q for ColoExitReason", s)
	}
	return nil
}

// COLOMode -> ColoMode (enum)

// ColoMode implements the "COLOMode" QMP API type.
type ColoMode int

// Known values of ColoMode.
const (
	ColoModeNone ColoMode = iota
	ColoModePrimary
	ColoModeSecondary
)

// String implements fmt.Stringer.
func (e ColoMode) String() string {
	switch e {
	case ColoModeNone:
		return "none"
	case ColoModePrimary:
		return "primary"
	case ColoModeSecondary:
		return "secondary"
	default:
		return fmt.Sprintf("ColoMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ColoMode) MarshalJSON() ([]byte, error) {
	switch e {
	case ColoModeNone:
		return json.Marshal("none")
	case ColoModePrimary:
		return json.Marshal("primary")
	case ColoModeSecondary:
		return json.Marshal("secondary")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ColoMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ColoMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = ColoModeNone
	case "primary":
		*e = ColoModePrimary
	case "secondary":
		*e = ColoModeSecondary
	default:
		return fmt.Errorf("unknown enum value %q for ColoMode", s)
	}
	return nil
}

// COLOStatus -> ColoStatus (struct)

// ColoStatus implements the "COLOStatus" QMP API type.
type ColoStatus struct {
	Mode     ColoMode       `json:"mode"`
	LastMode ColoMode       `json:"last-mode"`
	Reason   ColoExitReason `json:"reason"`
}

// EVENT COLO_EXIT

// ChardevBackend -> ChardevBackend (flat union)

// ChardevBackend implements the "ChardevBackend" QMP API type.
//
// Can be one of:
//   - ChardevBackendVariantBraille
//   - ChardevBackendVariantConsole
//   - ChardevBackendVariantDbus
//   - ChardevBackendVariantFile
//   - ChardevBackendVariantMemory
//   - ChardevBackendVariantMsmouse
//   - ChardevBackendVariantMux
//   - ChardevBackendVariantNull
//   - ChardevBackendVariantParallel
//   - ChardevBackendVariantPipe
//   - ChardevBackendVariantPty
//   - ChardevBackendVariantQemuVdagent
//   - ChardevBackendVariantRingbuf
//   - ChardevBackendVariantSerial
//   - ChardevBackendVariantSocket
//   - ChardevBackendVariantSpiceport
//   - ChardevBackendVariantSpicevmc
//   - ChardevBackendVariantStdio
//   - ChardevBackendVariantTestdev
//   - ChardevBackendVariantUDP
//   - ChardevBackendVariantVc
//   - ChardevBackendVariantWctablet
type ChardevBackend interface {
	isChardevBackend()
}

// ChardevBackendVariantBraille is an implementation of ChardevBackend.
type ChardevBackendVariantBraille struct {
	Data ChardevCommon `json:"data"`
}

func (ChardevBackendVariantBraille) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantBraille) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantBraille
	}{
		ChardevBackendKindBraille,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantConsole is an implementation of ChardevBackend.
type ChardevBackendVariantConsole struct {
	Data ChardevCommon `json:"data"`
}

func (ChardevBackendVariantConsole) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantConsole) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantConsole
	}{
		ChardevBackendKindConsole,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantDbus is an implementation of ChardevBackend.
type ChardevBackendVariantDbus struct {
	Data ChardevDBus `json:"data"`
}

func (ChardevBackendVariantDbus) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantDbus) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantDbus
	}{
		ChardevBackendKindDbus,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantFile is an implementation of ChardevBackend.
type ChardevBackendVariantFile struct {
	Data ChardevFile `json:"data"`
}

func (ChardevBackendVariantFile) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantFile) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantFile
	}{
		ChardevBackendKindFile,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantMemory is an implementation of ChardevBackend.
type ChardevBackendVariantMemory struct {
	Data ChardevRingbuf `json:"data"`
}

func (ChardevBackendVariantMemory) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantMemory) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantMemory
	}{
		ChardevBackendKindMemory,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantMsmouse is an implementation of ChardevBackend.
type ChardevBackendVariantMsmouse struct {
	Data ChardevCommon `json:"data"`
}

func (ChardevBackendVariantMsmouse) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantMsmouse) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantMsmouse
	}{
		ChardevBackendKindMsmouse,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantMux is an implementation of ChardevBackend.
type ChardevBackendVariantMux struct {
	Data ChardevMux `json:"data"`
}

func (ChardevBackendVariantMux) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantMux) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantMux
	}{
		ChardevBackendKindMux,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantNull is an implementation of ChardevBackend.
type ChardevBackendVariantNull struct {
	Data ChardevCommon `json:"data"`
}

func (ChardevBackendVariantNull) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantNull) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantNull
	}{
		ChardevBackendKindNull,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantParallel is an implementation of ChardevBackend.
type ChardevBackendVariantParallel struct {
	Data ChardevHostdev `json:"data"`
}

func (ChardevBackendVariantParallel) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantParallel) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantParallel
	}{
		ChardevBackendKindParallel,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantPipe is an implementation of ChardevBackend.
type ChardevBackendVariantPipe struct {
	Data ChardevHostdev `json:"data"`
}

func (ChardevBackendVariantPipe) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantPipe) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantPipe
	}{
		ChardevBackendKindPipe,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantPty is an implementation of ChardevBackend.
type ChardevBackendVariantPty struct {
	Data ChardevCommon `json:"data"`
}

func (ChardevBackendVariantPty) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantPty) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantPty
	}{
		ChardevBackendKindPty,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantQemuVdagent is an implementation of ChardevBackend.
type ChardevBackendVariantQemuVdagent struct {
	Data ChardevQemuVdAgent `json:"data"`
}

func (ChardevBackendVariantQemuVdagent) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantQemuVdagent) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantQemuVdagent
	}{
		ChardevBackendKindQemuVdagent,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantRingbuf is an implementation of ChardevBackend.
type ChardevBackendVariantRingbuf struct {
	Data ChardevRingbuf `json:"data"`
}

func (ChardevBackendVariantRingbuf) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantRingbuf) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantRingbuf
	}{
		ChardevBackendKindRingbuf,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantSerial is an implementation of ChardevBackend.
type ChardevBackendVariantSerial struct {
	Data ChardevHostdev `json:"data"`
}

func (ChardevBackendVariantSerial) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantSerial) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantSerial
	}{
		ChardevBackendKindSerial,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantSocket is an implementation of ChardevBackend.
type ChardevBackendVariantSocket struct {
	Data ChardevSocket `json:"data"`
}

func (ChardevBackendVariantSocket) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantSocket) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantSocket
	}{
		ChardevBackendKindSocket,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantSpiceport is an implementation of ChardevBackend.
type ChardevBackendVariantSpiceport struct {
	Data ChardevSpicePort `json:"data"`
}

func (ChardevBackendVariantSpiceport) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantSpiceport) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantSpiceport
	}{
		ChardevBackendKindSpiceport,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantSpicevmc is an implementation of ChardevBackend.
type ChardevBackendVariantSpicevmc struct {
	Data ChardevSpiceChannel `json:"data"`
}

func (ChardevBackendVariantSpicevmc) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantSpicevmc) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantSpicevmc
	}{
		ChardevBackendKindSpicevmc,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantStdio is an implementation of ChardevBackend.
type ChardevBackendVariantStdio struct {
	Data ChardevStdio `json:"data"`
}

func (ChardevBackendVariantStdio) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantStdio) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantStdio
	}{
		ChardevBackendKindStdio,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantTestdev is an implementation of ChardevBackend.
type ChardevBackendVariantTestdev struct {
	Data ChardevCommon `json:"data"`
}

func (ChardevBackendVariantTestdev) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantTestdev) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantTestdev
	}{
		ChardevBackendKindTestdev,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantUDP is an implementation of ChardevBackend.
type ChardevBackendVariantUDP struct {
	Data ChardevUDP `json:"data"`
}

func (ChardevBackendVariantUDP) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantUDP) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantUDP
	}{
		ChardevBackendKindUDP,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantVc is an implementation of ChardevBackend.
type ChardevBackendVariantVc struct {
	Data ChardevVc `json:"data"`
}

func (ChardevBackendVariantVc) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantVc) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantVc
	}{
		ChardevBackendKindVc,
		s,
	}
	return json.Marshal(v)
}

// ChardevBackendVariantWctablet is an implementation of ChardevBackend.
type ChardevBackendVariantWctablet struct {
	Data ChardevCommon `json:"data"`
}

func (ChardevBackendVariantWctablet) isChardevBackend() {}

// MarshalJSON implements json.Marshaler.
func (s ChardevBackendVariantWctablet) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
		ChardevBackendVariantWctablet
	}{
		ChardevBackendKindWctablet,
		s,
	}
	return json.Marshal(v)
}

func decodeChardevBackend(bs json.RawMessage) (ChardevBackend, error) {
	v := struct {
		Type ChardevBackendKind `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case ChardevBackendKindBraille:
		var ret ChardevBackendVariantBraille
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindConsole:
		var ret ChardevBackendVariantConsole
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindDbus:
		var ret ChardevBackendVariantDbus
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindFile:
		var ret ChardevBackendVariantFile
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindMemory:
		var ret ChardevBackendVariantMemory
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindMsmouse:
		var ret ChardevBackendVariantMsmouse
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindMux:
		var ret ChardevBackendVariantMux
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindNull:
		var ret ChardevBackendVariantNull
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindParallel:
		var ret ChardevBackendVariantParallel
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindPipe:
		var ret ChardevBackendVariantPipe
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindPty:
		var ret ChardevBackendVariantPty
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindQemuVdagent:
		var ret ChardevBackendVariantQemuVdagent
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindRingbuf:
		var ret ChardevBackendVariantRingbuf
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindSerial:
		var ret ChardevBackendVariantSerial
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindSocket:
		var ret ChardevBackendVariantSocket
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindSpiceport:
		var ret ChardevBackendVariantSpiceport
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindSpicevmc:
		var ret ChardevBackendVariantSpicevmc
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindStdio:
		var ret ChardevBackendVariantStdio
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindTestdev:
		var ret ChardevBackendVariantTestdev
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindUDP:
		var ret ChardevBackendVariantUDP
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindVc:
		var ret ChardevBackendVariantVc
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ChardevBackendKindWctablet:
		var ret ChardevBackendVariantWctablet
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union ChardevBackend", v.Type)
	}
}

// ChardevBackendInfo -> ChardevBackendInfo (struct)

// ChardevBackendInfo implements the "ChardevBackendInfo" QMP API type.
type ChardevBackendInfo struct {
	Name string `json:"name"`
}

// ChardevBackendKind -> ChardevBackendKind (enum)

// ChardevBackendKind implements the "ChardevBackendKind" QMP API type.
type ChardevBackendKind int

// Known values of ChardevBackendKind.
const (
	ChardevBackendKindFile ChardevBackendKind = iota
	ChardevBackendKindSerial
	ChardevBackendKindParallel
	ChardevBackendKindPipe
	ChardevBackendKindSocket
	ChardevBackendKindUDP
	ChardevBackendKindPty
	ChardevBackendKindNull
	ChardevBackendKindMux
	ChardevBackendKindMsmouse
	ChardevBackendKindWctablet
	ChardevBackendKindBraille
	ChardevBackendKindTestdev
	ChardevBackendKindStdio
	ChardevBackendKindConsole
	ChardevBackendKindSpicevmc
	ChardevBackendKindSpiceport
	ChardevBackendKindQemuVdagent
	ChardevBackendKindDbus
	ChardevBackendKindVc
	ChardevBackendKindRingbuf
	ChardevBackendKindMemory
)

// String implements fmt.Stringer.
func (e ChardevBackendKind) String() string {
	switch e {
	case ChardevBackendKindFile:
		return "file"
	case ChardevBackendKindSerial:
		return "serial"
	case ChardevBackendKindParallel:
		return "parallel"
	case ChardevBackendKindPipe:
		return "pipe"
	case ChardevBackendKindSocket:
		return "socket"
	case ChardevBackendKindUDP:
		return "udp"
	case ChardevBackendKindPty:
		return "pty"
	case ChardevBackendKindNull:
		return "null"
	case ChardevBackendKindMux:
		return "mux"
	case ChardevBackendKindMsmouse:
		return "msmouse"
	case ChardevBackendKindWctablet:
		return "wctablet"
	case ChardevBackendKindBraille:
		return "braille"
	case ChardevBackendKindTestdev:
		return "testdev"
	case ChardevBackendKindStdio:
		return "stdio"
	case ChardevBackendKindConsole:
		return "console"
	case ChardevBackendKindSpicevmc:
		return "spicevmc"
	case ChardevBackendKindSpiceport:
		return "spiceport"
	case ChardevBackendKindQemuVdagent:
		return "qemu-vdagent"
	case ChardevBackendKindDbus:
		return "dbus"
	case ChardevBackendKindVc:
		return "vc"
	case ChardevBackendKindRingbuf:
		return "ringbuf"
	case ChardevBackendKindMemory:
		return "memory"
	default:
		return fmt.Sprintf("ChardevBackendKind(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ChardevBackendKind) MarshalJSON() ([]byte, error) {
	switch e {
	case ChardevBackendKindFile:
		return json.Marshal("file")
	case ChardevBackendKindSerial:
		return json.Marshal("serial")
	case ChardevBackendKindParallel:
		return json.Marshal("parallel")
	case ChardevBackendKindPipe:
		return json.Marshal("pipe")
	case ChardevBackendKindSocket:
		return json.Marshal("socket")
	case ChardevBackendKindUDP:
		return json.Marshal("udp")
	case ChardevBackendKindPty:
		return json.Marshal("pty")
	case ChardevBackendKindNull:
		return json.Marshal("null")
	case ChardevBackendKindMux:
		return json.Marshal("mux")
	case ChardevBackendKindMsmouse:
		return json.Marshal("msmouse")
	case ChardevBackendKindWctablet:
		return json.Marshal("wctablet")
	case ChardevBackendKindBraille:
		return json.Marshal("braille")
	case ChardevBackendKindTestdev:
		return json.Marshal("testdev")
	case ChardevBackendKindStdio:
		return json.Marshal("stdio")
	case ChardevBackendKindConsole:
		return json.Marshal("console")
	case ChardevBackendKindSpicevmc:
		return json.Marshal("spicevmc")
	case ChardevBackendKindSpiceport:
		return json.Marshal("spiceport")
	case ChardevBackendKindQemuVdagent:
		return json.Marshal("qemu-vdagent")
	case ChardevBackendKindDbus:
		return json.Marshal("dbus")
	case ChardevBackendKindVc:
		return json.Marshal("vc")
	case ChardevBackendKindRingbuf:
		return json.Marshal("ringbuf")
	case ChardevBackendKindMemory:
		return json.Marshal("memory")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ChardevBackendKind", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ChardevBackendKind) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "file":
		*e = ChardevBackendKindFile
	case "serial":
		*e = ChardevBackendKindSerial
	case "parallel":
		*e = ChardevBackendKindParallel
	case "pipe":
		*e = ChardevBackendKindPipe
	case "socket":
		*e = ChardevBackendKindSocket
	case "udp":
		*e = ChardevBackendKindUDP
	case "pty":
		*e = ChardevBackendKindPty
	case "null":
		*e = ChardevBackendKindNull
	case "mux":
		*e = ChardevBackendKindMux
	case "msmouse":
		*e = ChardevBackendKindMsmouse
	case "wctablet":
		*e = ChardevBackendKindWctablet
	case "braille":
		*e = ChardevBackendKindBraille
	case "testdev":
		*e = ChardevBackendKindTestdev
	case "stdio":
		*e = ChardevBackendKindStdio
	case "console":
		*e = ChardevBackendKindConsole
	case "spicevmc":
		*e = ChardevBackendKindSpicevmc
	case "spiceport":
		*e = ChardevBackendKindSpiceport
	case "qemu-vdagent":
		*e = ChardevBackendKindQemuVdagent
	case "dbus":
		*e = ChardevBackendKindDbus
	case "vc":
		*e = ChardevBackendKindVc
	case "ringbuf":
		*e = ChardevBackendKindRingbuf
	case "memory":
		*e = ChardevBackendKindMemory
	default:
		return fmt.Errorf("unknown enum value %q for ChardevBackendKind", s)
	}
	return nil
}

// ChardevCommon -> ChardevCommon (struct)

// ChardevCommon implements the "ChardevCommon" QMP API type.
type ChardevCommon struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
}

// ChardevDBus -> ChardevDBus (struct)

// ChardevDBus implements the "ChardevDBus" QMP API type.
type ChardevDBus struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
	Name      string  `json:"name"`
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

// ChardevQemuVDAgent -> ChardevQemuVdAgent (struct)

// ChardevQemuVdAgent implements the "ChardevQemuVDAgent" QMP API type.
type ChardevQemuVdAgent struct {
	Logfile   *string `json:"logfile,omitempty"`
	Logappend *bool   `json:"logappend,omitempty"`
	Mouse     *bool   `json:"mouse,omitempty"`
	Clipboard *bool   `json:"clipboard,omitempty"`
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
	TLSAuthz  *string             `json:"tls-authz,omitempty"`
	Server    *bool               `json:"server,omitempty"`
	Wait      *bool               `json:"wait,omitempty"`
	Nodelay   *bool               `json:"nodelay,omitempty"`
	Telnet    *bool               `json:"telnet,omitempty"`
	Tn3270    *bool               `json:"tn3270,omitempty"`
	Websocket *bool               `json:"websocket,omitempty"`
	Reconnect *int64              `json:"reconnect,omitempty"`
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
	Logfile   *string              `json:"logfile,omitempty"`
	Logappend *bool                `json:"logappend,omitempty"`
	Remote    SocketAddressLegacy  `json:"remote"`
	Local     *SocketAddressLegacy `json:"local,omitempty"`
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

// CompressionStats -> CompressionStats (struct)

// CompressionStats implements the "CompressionStats" QMP API type.
type CompressionStats struct {
	Pages           int64   `json:"pages"`
	Busy            int64   `json:"busy"`
	BusyRate        float64 `json:"busy-rate"`
	CompressedSize  int64   `json:"compressed-size"`
	CompressionRate float64 `json:"compression-rate"`
}

// CpuDefinitionInfo -> CPUDefinitionInfo (struct)

// CPUDefinitionInfo implements the "CpuDefinitionInfo" QMP API type.
type CPUDefinitionInfo struct {
	Name                string   `json:"name"`
	MigrationSafe       *bool    `json:"migration-safe,omitempty"`
	Static              bool     `json:"static"`
	UnavailableFeatures []string `json:"unavailable-features,omitempty"`
	Typename            string   `json:"typename"`
	AliasOf             *string  `json:"alias-of,omitempty"`
	Deprecated          bool     `json:"deprecated"`
}

// CpuInfoFast -> CPUInfoFast (flat union)

// CPUInfoFast implements the "CpuInfoFast" QMP API type.
//
// Can be one of:
//   - CPUInfoFastVariantS390X
type CPUInfoFast interface {
	isCPUInfoFast()
}

// CPUInfoFastVariantS390X is an implementation of CPUInfoFast.
type CPUInfoFastVariantS390X struct {
	CPUIndex int64                  `json:"cpu-index"`
	QomPath  string                 `json:"qom-path"`
	ThreadID int64                  `json:"thread-id"`
	Props    *CPUInstanceProperties `json:"props,omitempty"`
	CPUState CPUS390State           `json:"cpu-state"`
}

func (CPUInfoFastVariantS390X) isCPUInfoFast() {}

// MarshalJSON implements json.Marshaler.
func (s CPUInfoFastVariantS390X) MarshalJSON() ([]byte, error) {
	v := struct {
		Target SysEmuTarget `json:"target"`
		CPUInfoFastVariantS390X
	}{
		SysEmuTargetS390X,
		s,
	}
	return json.Marshal(v)
}

func decodeCPUInfoFast(bs json.RawMessage) (CPUInfoFast, error) {
	v := struct {
		Target SysEmuTarget `json:"target"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Target {
	case SysEmuTargetS390X:
		var ret CPUInfoFastVariantS390X
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union CPUInfoFast", v.Target)
	}
}

// CpuInstanceProperties -> CPUInstanceProperties (struct)

// CPUInstanceProperties implements the "CpuInstanceProperties" QMP API type.
type CPUInstanceProperties struct {
	NodeID    *int64 `json:"node-id,omitempty"`
	SocketID  *int64 `json:"socket-id,omitempty"`
	DieID     *int64 `json:"die-id,omitempty"`
	ClusterID *int64 `json:"cluster-id,omitempty"`
	CoreID    *int64 `json:"core-id,omitempty"`
	ThreadID  *int64 `json:"thread-id,omitempty"`
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

// CpuS390State -> CPUS390State (enum)

// CPUS390State implements the "CpuS390State" QMP API type.
type CPUS390State int

// Known values of CPUS390State.
const (
	CPUS390StateUninitialized CPUS390State = iota
	CPUS390StateStopped
	CPUS390StateCheckStop
	CPUS390StateOperating
	CPUS390StateLoad
)

// String implements fmt.Stringer.
func (e CPUS390State) String() string {
	switch e {
	case CPUS390StateUninitialized:
		return "uninitialized"
	case CPUS390StateStopped:
		return "stopped"
	case CPUS390StateCheckStop:
		return "check-stop"
	case CPUS390StateOperating:
		return "operating"
	case CPUS390StateLoad:
		return "load"
	default:
		return fmt.Sprintf("CPUS390State(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e CPUS390State) MarshalJSON() ([]byte, error) {
	switch e {
	case CPUS390StateUninitialized:
		return json.Marshal("uninitialized")
	case CPUS390StateStopped:
		return json.Marshal("stopped")
	case CPUS390StateCheckStop:
		return json.Marshal("check-stop")
	case CPUS390StateOperating:
		return json.Marshal("operating")
	case CPUS390StateLoad:
		return json.Marshal("load")
	default:
		return nil, fmt.Errorf("unknown enum value %q for CPUS390State", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *CPUS390State) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "uninitialized":
		*e = CPUS390StateUninitialized
	case "stopped":
		*e = CPUS390StateStopped
	case "check-stop":
		*e = CPUS390StateCheckStop
	case "operating":
		*e = CPUS390StateOperating
	case "load":
		*e = CPUS390StateLoad
	default:
		return fmt.Errorf("unknown enum value %q for CPUS390State", s)
	}
	return nil
}

// CurrentMachineParams -> CurrentMachineParams (struct)

// CurrentMachineParams implements the "CurrentMachineParams" QMP API type.
type CurrentMachineParams struct {
	WakeupSuspendSupport bool `json:"wakeup-suspend-support"`
}

// EVENT DEVICE_DELETED

// EVENT DEVICE_TRAY_MOVED

// EVENT DEVICE_UNPLUG_GUEST_ERROR

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

// DirtyLimitInfo -> DirtyLimitInfo (struct)

// DirtyLimitInfo implements the "DirtyLimitInfo" QMP API type.
type DirtyLimitInfo struct {
	CPUIndex    int64  `json:"cpu-index"`
	LimitRate   uint64 `json:"limit-rate"`
	CurrentRate uint64 `json:"current-rate"`
}

// DirtyRateInfo -> DirtyRateInfo (struct)

// DirtyRateInfo implements the "DirtyRateInfo" QMP API type.
type DirtyRateInfo struct {
	DirtyRate     *int64               `json:"dirty-rate,omitempty"`
	Status        DirtyRateStatus      `json:"status"`
	StartTime     int64                `json:"start-time"`
	CalcTime      int64                `json:"calc-time"`
	SamplePages   uint64               `json:"sample-pages"`
	Mode          DirtyRateMeasureMode `json:"mode"`
	VcpuDirtyRate []DirtyRateVcpu      `json:"vcpu-dirty-rate,omitempty"`
}

// DirtyRateMeasureMode -> DirtyRateMeasureMode (enum)

// DirtyRateMeasureMode implements the "DirtyRateMeasureMode" QMP API type.
type DirtyRateMeasureMode int

// Known values of DirtyRateMeasureMode.
const (
	DirtyRateMeasureModePageSampling DirtyRateMeasureMode = iota
	DirtyRateMeasureModeDirtyRing
	DirtyRateMeasureModeDirtyBitmap
)

// String implements fmt.Stringer.
func (e DirtyRateMeasureMode) String() string {
	switch e {
	case DirtyRateMeasureModePageSampling:
		return "page-sampling"
	case DirtyRateMeasureModeDirtyRing:
		return "dirty-ring"
	case DirtyRateMeasureModeDirtyBitmap:
		return "dirty-bitmap"
	default:
		return fmt.Sprintf("DirtyRateMeasureMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DirtyRateMeasureMode) MarshalJSON() ([]byte, error) {
	switch e {
	case DirtyRateMeasureModePageSampling:
		return json.Marshal("page-sampling")
	case DirtyRateMeasureModeDirtyRing:
		return json.Marshal("dirty-ring")
	case DirtyRateMeasureModeDirtyBitmap:
		return json.Marshal("dirty-bitmap")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DirtyRateMeasureMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DirtyRateMeasureMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "page-sampling":
		*e = DirtyRateMeasureModePageSampling
	case "dirty-ring":
		*e = DirtyRateMeasureModeDirtyRing
	case "dirty-bitmap":
		*e = DirtyRateMeasureModeDirtyBitmap
	default:
		return fmt.Errorf("unknown enum value %q for DirtyRateMeasureMode", s)
	}
	return nil
}

// DirtyRateStatus -> DirtyRateStatus (enum)

// DirtyRateStatus implements the "DirtyRateStatus" QMP API type.
type DirtyRateStatus int

// Known values of DirtyRateStatus.
const (
	DirtyRateStatusUnstarted DirtyRateStatus = iota
	DirtyRateStatusMeasuring
	DirtyRateStatusMeasured
)

// String implements fmt.Stringer.
func (e DirtyRateStatus) String() string {
	switch e {
	case DirtyRateStatusUnstarted:
		return "unstarted"
	case DirtyRateStatusMeasuring:
		return "measuring"
	case DirtyRateStatusMeasured:
		return "measured"
	default:
		return fmt.Sprintf("DirtyRateStatus(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DirtyRateStatus) MarshalJSON() ([]byte, error) {
	switch e {
	case DirtyRateStatusUnstarted:
		return json.Marshal("unstarted")
	case DirtyRateStatusMeasuring:
		return json.Marshal("measuring")
	case DirtyRateStatusMeasured:
		return json.Marshal("measured")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DirtyRateStatus", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DirtyRateStatus) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "unstarted":
		*e = DirtyRateStatusUnstarted
	case "measuring":
		*e = DirtyRateStatusMeasuring
	case "measured":
		*e = DirtyRateStatusMeasured
	default:
		return fmt.Errorf("unknown enum value %q for DirtyRateStatus", s)
	}
	return nil
}

// DirtyRateVcpu -> DirtyRateVcpu (struct)

// DirtyRateVcpu implements the "DirtyRateVcpu" QMP API type.
type DirtyRateVcpu struct {
	ID        int64 `json:"id"`
	DirtyRate int64 `json:"dirty-rate"`
}

// DisplayGLMode -> DisplayGlMode (enum)

// DisplayGlMode implements the "DisplayGLMode" QMP API type.
type DisplayGlMode int

// Known values of DisplayGlMode.
const (
	DisplayGlModeOff DisplayGlMode = iota
	DisplayGlModeOn
	DisplayGlModeCore
	DisplayGlModeEs
)

// String implements fmt.Stringer.
func (e DisplayGlMode) String() string {
	switch e {
	case DisplayGlModeOff:
		return "off"
	case DisplayGlModeOn:
		return "on"
	case DisplayGlModeCore:
		return "core"
	case DisplayGlModeEs:
		return "es"
	default:
		return fmt.Sprintf("DisplayGlMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DisplayGlMode) MarshalJSON() ([]byte, error) {
	switch e {
	case DisplayGlModeOff:
		return json.Marshal("off")
	case DisplayGlModeOn:
		return json.Marshal("on")
	case DisplayGlModeCore:
		return json.Marshal("core")
	case DisplayGlModeEs:
		return json.Marshal("es")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DisplayGlMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DisplayGlMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "off":
		*e = DisplayGlModeOff
	case "on":
		*e = DisplayGlModeOn
	case "core":
		*e = DisplayGlModeCore
	case "es":
		*e = DisplayGlModeEs
	default:
		return fmt.Errorf("unknown enum value %q for DisplayGlMode", s)
	}
	return nil
}

// DisplayOptions -> DisplayOptions (flat union)

// DisplayOptions implements the "DisplayOptions" QMP API type.
//
// Can be one of:
//   - DisplayOptionsVariantCocoa
//   - DisplayOptionsVariantCurses
//   - DisplayOptionsVariantDbus
//   - DisplayOptionsVariantEglHeadless
//   - DisplayOptionsVariantGtk
//   - DisplayOptionsVariantSdl
type DisplayOptions interface {
	isDisplayOptions()
}

// DisplayOptionsVariantCocoa is an implementation of DisplayOptions.
type DisplayOptionsVariantCocoa struct {
	FullScreen     *bool          `json:"full-screen,omitempty"`
	WindowClose    *bool          `json:"window-close,omitempty"`
	ShowCursor     *bool          `json:"show-cursor,omitempty"`
	Gl             *DisplayGlMode `json:"gl,omitempty"`
	LeftCommandKey *bool          `json:"left-command-key,omitempty"`
	FullGrab       *bool          `json:"full-grab,omitempty"`
	SwapOptCmd     *bool          `json:"swap-opt-cmd,omitempty"`
}

func (DisplayOptionsVariantCocoa) isDisplayOptions() {}

// MarshalJSON implements json.Marshaler.
func (s DisplayOptionsVariantCocoa) MarshalJSON() ([]byte, error) {
	v := struct {
		Type DisplayType `json:"type"`
		DisplayOptionsVariantCocoa
	}{
		DisplayTypeCocoa,
		s,
	}
	return json.Marshal(v)
}

// DisplayOptionsVariantCurses is an implementation of DisplayOptions.
type DisplayOptionsVariantCurses struct {
	FullScreen  *bool          `json:"full-screen,omitempty"`
	WindowClose *bool          `json:"window-close,omitempty"`
	ShowCursor  *bool          `json:"show-cursor,omitempty"`
	Gl          *DisplayGlMode `json:"gl,omitempty"`
	Charset     *string        `json:"charset,omitempty"`
}

func (DisplayOptionsVariantCurses) isDisplayOptions() {}

// MarshalJSON implements json.Marshaler.
func (s DisplayOptionsVariantCurses) MarshalJSON() ([]byte, error) {
	v := struct {
		Type DisplayType `json:"type"`
		DisplayOptionsVariantCurses
	}{
		DisplayTypeCurses,
		s,
	}
	return json.Marshal(v)
}

// DisplayOptionsVariantDbus is an implementation of DisplayOptions.
type DisplayOptionsVariantDbus struct {
	FullScreen  *bool          `json:"full-screen,omitempty"`
	WindowClose *bool          `json:"window-close,omitempty"`
	ShowCursor  *bool          `json:"show-cursor,omitempty"`
	Gl          *DisplayGlMode `json:"gl,omitempty"`
	Rendernode  *string        `json:"rendernode,omitempty"`
	Addr        *string        `json:"addr,omitempty"`
	P2P         *bool          `json:"p2p,omitempty"`
	Audiodev    *string        `json:"audiodev,omitempty"`
}

func (DisplayOptionsVariantDbus) isDisplayOptions() {}

// MarshalJSON implements json.Marshaler.
func (s DisplayOptionsVariantDbus) MarshalJSON() ([]byte, error) {
	v := struct {
		Type DisplayType `json:"type"`
		DisplayOptionsVariantDbus
	}{
		DisplayTypeDbus,
		s,
	}
	return json.Marshal(v)
}

// DisplayOptionsVariantEglHeadless is an implementation of DisplayOptions.
type DisplayOptionsVariantEglHeadless struct {
	FullScreen  *bool          `json:"full-screen,omitempty"`
	WindowClose *bool          `json:"window-close,omitempty"`
	ShowCursor  *bool          `json:"show-cursor,omitempty"`
	Gl          *DisplayGlMode `json:"gl,omitempty"`
	Rendernode  *string        `json:"rendernode,omitempty"`
}

func (DisplayOptionsVariantEglHeadless) isDisplayOptions() {}

// MarshalJSON implements json.Marshaler.
func (s DisplayOptionsVariantEglHeadless) MarshalJSON() ([]byte, error) {
	v := struct {
		Type DisplayType `json:"type"`
		DisplayOptionsVariantEglHeadless
	}{
		DisplayTypeEglHeadless,
		s,
	}
	return json.Marshal(v)
}

// DisplayOptionsVariantGtk is an implementation of DisplayOptions.
type DisplayOptionsVariantGtk struct {
	FullScreen  *bool          `json:"full-screen,omitempty"`
	WindowClose *bool          `json:"window-close,omitempty"`
	ShowCursor  *bool          `json:"show-cursor,omitempty"`
	Gl          *DisplayGlMode `json:"gl,omitempty"`
	GrabOnHover *bool          `json:"grab-on-hover,omitempty"`
	ZoomToFit   *bool          `json:"zoom-to-fit,omitempty"`
	ShowTabs    *bool          `json:"show-tabs,omitempty"`
}

func (DisplayOptionsVariantGtk) isDisplayOptions() {}

// MarshalJSON implements json.Marshaler.
func (s DisplayOptionsVariantGtk) MarshalJSON() ([]byte, error) {
	v := struct {
		Type DisplayType `json:"type"`
		DisplayOptionsVariantGtk
	}{
		DisplayTypeGtk,
		s,
	}
	return json.Marshal(v)
}

// DisplayOptionsVariantSdl is an implementation of DisplayOptions.
type DisplayOptionsVariantSdl struct {
	FullScreen  *bool          `json:"full-screen,omitempty"`
	WindowClose *bool          `json:"window-close,omitempty"`
	ShowCursor  *bool          `json:"show-cursor,omitempty"`
	Gl          *DisplayGlMode `json:"gl,omitempty"`
	GrabMod     *HotKeyMod     `json:"grab-mod,omitempty"`
}

func (DisplayOptionsVariantSdl) isDisplayOptions() {}

// MarshalJSON implements json.Marshaler.
func (s DisplayOptionsVariantSdl) MarshalJSON() ([]byte, error) {
	v := struct {
		Type DisplayType `json:"type"`
		DisplayOptionsVariantSdl
	}{
		DisplayTypeSdl,
		s,
	}
	return json.Marshal(v)
}

func decodeDisplayOptions(bs json.RawMessage) (DisplayOptions, error) {
	v := struct {
		Type DisplayType `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case DisplayTypeCocoa:
		var ret DisplayOptionsVariantCocoa
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case DisplayTypeCurses:
		var ret DisplayOptionsVariantCurses
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case DisplayTypeDbus:
		var ret DisplayOptionsVariantDbus
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case DisplayTypeEglHeadless:
		var ret DisplayOptionsVariantEglHeadless
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case DisplayTypeGtk:
		var ret DisplayOptionsVariantGtk
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case DisplayTypeSdl:
		var ret DisplayOptionsVariantSdl
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union DisplayOptions", v.Type)
	}
}

// DisplayProtocol -> DisplayProtocol (enum)

// DisplayProtocol implements the "DisplayProtocol" QMP API type.
type DisplayProtocol int

// Known values of DisplayProtocol.
const (
	DisplayProtocolVNC DisplayProtocol = iota
	DisplayProtocolSpice
)

// String implements fmt.Stringer.
func (e DisplayProtocol) String() string {
	switch e {
	case DisplayProtocolVNC:
		return "vnc"
	case DisplayProtocolSpice:
		return "spice"
	default:
		return fmt.Sprintf("DisplayProtocol(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DisplayProtocol) MarshalJSON() ([]byte, error) {
	switch e {
	case DisplayProtocolVNC:
		return json.Marshal("vnc")
	case DisplayProtocolSpice:
		return json.Marshal("spice")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DisplayProtocol", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DisplayProtocol) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "vnc":
		*e = DisplayProtocolVNC
	case "spice":
		*e = DisplayProtocolSpice
	default:
		return fmt.Errorf("unknown enum value %q for DisplayProtocol", s)
	}
	return nil
}

// DisplayReloadOptions -> DisplayReloadOptions (flat union)

// DisplayReloadOptions implements the "DisplayReloadOptions" QMP API type.
//
// Can be one of:
//   - DisplayReloadOptionsVariantVNC
type DisplayReloadOptions interface {
	isDisplayReloadOptions()
}

// DisplayReloadOptionsVariantVNC is an implementation of DisplayReloadOptions.
type DisplayReloadOptionsVariantVNC struct {
	TLSCerts *bool `json:"tls-certs,omitempty"`
}

func (DisplayReloadOptionsVariantVNC) isDisplayReloadOptions() {}

// MarshalJSON implements json.Marshaler.
func (s DisplayReloadOptionsVariantVNC) MarshalJSON() ([]byte, error) {
	v := struct {
		Type DisplayReloadType `json:"type"`
		DisplayReloadOptionsVariantVNC
	}{
		DisplayReloadTypeVNC,
		s,
	}
	return json.Marshal(v)
}

func decodeDisplayReloadOptions(bs json.RawMessage) (DisplayReloadOptions, error) {
	v := struct {
		Type DisplayReloadType `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case DisplayReloadTypeVNC:
		var ret DisplayReloadOptionsVariantVNC
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union DisplayReloadOptions", v.Type)
	}
}

// DisplayReloadType -> DisplayReloadType (enum)

// DisplayReloadType implements the "DisplayReloadType" QMP API type.
type DisplayReloadType int

// Known values of DisplayReloadType.
const (
	DisplayReloadTypeVNC DisplayReloadType = iota
)

// String implements fmt.Stringer.
func (e DisplayReloadType) String() string {
	switch e {
	case DisplayReloadTypeVNC:
		return "vnc"
	default:
		return fmt.Sprintf("DisplayReloadType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DisplayReloadType) MarshalJSON() ([]byte, error) {
	switch e {
	case DisplayReloadTypeVNC:
		return json.Marshal("vnc")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DisplayReloadType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DisplayReloadType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "vnc":
		*e = DisplayReloadTypeVNC
	default:
		return fmt.Errorf("unknown enum value %q for DisplayReloadType", s)
	}
	return nil
}

// DisplayType -> DisplayType (enum)

// DisplayType implements the "DisplayType" QMP API type.
type DisplayType int

// Known values of DisplayType.
const (
	DisplayTypeDefault DisplayType = iota
	DisplayTypeNone
	DisplayTypeGtk
	DisplayTypeSdl
	DisplayTypeEglHeadless
	DisplayTypeCurses
	DisplayTypeCocoa
	DisplayTypeSpiceApp
	DisplayTypeDbus
)

// String implements fmt.Stringer.
func (e DisplayType) String() string {
	switch e {
	case DisplayTypeDefault:
		return "default"
	case DisplayTypeNone:
		return "none"
	case DisplayTypeGtk:
		return "gtk"
	case DisplayTypeSdl:
		return "sdl"
	case DisplayTypeEglHeadless:
		return "egl-headless"
	case DisplayTypeCurses:
		return "curses"
	case DisplayTypeCocoa:
		return "cocoa"
	case DisplayTypeSpiceApp:
		return "spice-app"
	case DisplayTypeDbus:
		return "dbus"
	default:
		return fmt.Sprintf("DisplayType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DisplayType) MarshalJSON() ([]byte, error) {
	switch e {
	case DisplayTypeDefault:
		return json.Marshal("default")
	case DisplayTypeNone:
		return json.Marshal("none")
	case DisplayTypeGtk:
		return json.Marshal("gtk")
	case DisplayTypeSdl:
		return json.Marshal("sdl")
	case DisplayTypeEglHeadless:
		return json.Marshal("egl-headless")
	case DisplayTypeCurses:
		return json.Marshal("curses")
	case DisplayTypeCocoa:
		return json.Marshal("cocoa")
	case DisplayTypeSpiceApp:
		return json.Marshal("spice-app")
	case DisplayTypeDbus:
		return json.Marshal("dbus")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DisplayType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DisplayType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "default":
		*e = DisplayTypeDefault
	case "none":
		*e = DisplayTypeNone
	case "gtk":
		*e = DisplayTypeGtk
	case "sdl":
		*e = DisplayTypeSdl
	case "egl-headless":
		*e = DisplayTypeEglHeadless
	case "curses":
		*e = DisplayTypeCurses
	case "cocoa":
		*e = DisplayTypeCocoa
	case "spice-app":
		*e = DisplayTypeSpiceApp
	case "dbus":
		*e = DisplayTypeDbus
	default:
		return fmt.Errorf("unknown enum value %q for DisplayType", s)
	}
	return nil
}

// DisplayUpdateOptions -> DisplayUpdateOptions (flat union)

// DisplayUpdateOptions implements the "DisplayUpdateOptions" QMP API type.
//
// Can be one of:
//   - DisplayUpdateOptionsVariantVNC
type DisplayUpdateOptions interface {
	isDisplayUpdateOptions()
}

// DisplayUpdateOptionsVariantVNC is an implementation of DisplayUpdateOptions.
type DisplayUpdateOptionsVariantVNC struct {
	Addresses []SocketAddress `json:"addresses,omitempty"`
}

func (DisplayUpdateOptionsVariantVNC) isDisplayUpdateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s DisplayUpdateOptionsVariantVNC) MarshalJSON() ([]byte, error) {
	v := struct {
		Type DisplayUpdateType `json:"type"`
		DisplayUpdateOptionsVariantVNC
	}{
		DisplayUpdateTypeVNC,
		s,
	}
	return json.Marshal(v)
}

func decodeDisplayUpdateOptions(bs json.RawMessage) (DisplayUpdateOptions, error) {
	v := struct {
		Type DisplayUpdateType `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case DisplayUpdateTypeVNC:
		var ret DisplayUpdateOptionsVariantVNC
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union DisplayUpdateOptions", v.Type)
	}
}

// DisplayUpdateType -> DisplayUpdateType (enum)

// DisplayUpdateType implements the "DisplayUpdateType" QMP API type.
type DisplayUpdateType int

// Known values of DisplayUpdateType.
const (
	DisplayUpdateTypeVNC DisplayUpdateType = iota
)

// String implements fmt.Stringer.
func (e DisplayUpdateType) String() string {
	switch e {
	case DisplayUpdateTypeVNC:
		return "vnc"
	default:
		return fmt.Sprintf("DisplayUpdateType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e DisplayUpdateType) MarshalJSON() ([]byte, error) {
	switch e {
	case DisplayUpdateTypeVNC:
		return json.Marshal("vnc")
	default:
		return nil, fmt.Errorf("unknown enum value %q for DisplayUpdateType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *DisplayUpdateType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "vnc":
		*e = DisplayUpdateTypeVNC
	default:
		return fmt.Errorf("unknown enum value %q for DisplayUpdateType", s)
	}
	return nil
}

// DriveBackup -> DriveBackup (struct)

// DriveBackup implements the "DriveBackup" QMP API type.
type DriveBackup struct {
	JobID          *string          `json:"job-id,omitempty"`
	Device         string           `json:"device"`
	Sync           MirrorSyncMode   `json:"sync"`
	Speed          *int64           `json:"speed,omitempty"`
	Bitmap         *string          `json:"bitmap,omitempty"`
	BitmapMode     *BitmapSyncMode  `json:"bitmap-mode,omitempty"`
	Compress       *bool            `json:"compress,omitempty"`
	OnSourceError  *BlockdevOnError `json:"on-source-error,omitempty"`
	OnTargetError  *BlockdevOnError `json:"on-target-error,omitempty"`
	AutoFinalize   *bool            `json:"auto-finalize,omitempty"`
	AutoDismiss    *bool            `json:"auto-dismiss,omitempty"`
	FilterNodeName *string          `json:"filter-node-name,omitempty"`
	XPerf          *BackupPerf      `json:"x-perf,omitempty"`
	Target         string           `json:"target"`
	Format         *string          `json:"format,omitempty"`
	Mode           *NewImageMode    `json:"mode,omitempty"`
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
	CopyMode      *MirrorCopyMode  `json:"copy-mode,omitempty"`
	AutoFinalize  *bool            `json:"auto-finalize,omitempty"`
	AutoDismiss   *bool            `json:"auto-dismiss,omitempty"`
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
	DumpGuestMemoryFormatWinDmp
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
	case DumpGuestMemoryFormatWinDmp:
		return "win-dmp"
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
	case DumpGuestMemoryFormatWinDmp:
		return json.Marshal("win-dmp")
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
	case "win-dmp":
		*e = DumpGuestMemoryFormatWinDmp
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

// ExpirePasswordOptions -> ExpirePasswordOptions (flat union)

// ExpirePasswordOptions implements the "ExpirePasswordOptions" QMP API type.
//
// Can be one of:
//   - ExpirePasswordOptionsVariantVNC
type ExpirePasswordOptions interface {
	isExpirePasswordOptions()
}

// ExpirePasswordOptionsVariantVNC is an implementation of ExpirePasswordOptions.
type ExpirePasswordOptionsVariantVNC struct {
	Time    string  `json:"time"`
	Display *string `json:"display,omitempty"`
}

func (ExpirePasswordOptionsVariantVNC) isExpirePasswordOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ExpirePasswordOptionsVariantVNC) MarshalJSON() ([]byte, error) {
	v := struct {
		Protocol DisplayProtocol `json:"protocol"`
		ExpirePasswordOptionsVariantVNC
	}{
		DisplayProtocolVNC,
		s,
	}
	return json.Marshal(v)
}

func decodeExpirePasswordOptions(bs json.RawMessage) (ExpirePasswordOptions, error) {
	v := struct {
		Protocol DisplayProtocol `json:"protocol"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Protocol {
	case DisplayProtocolVNC:
		var ret ExpirePasswordOptionsVariantVNC
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union ExpirePasswordOptions", v.Protocol)
	}
}

// EVENT FAILOVER_NEGOTIATED

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

// FuseExportAllowOther -> FuseExportAllowOther (enum)

// FuseExportAllowOther implements the "FuseExportAllowOther" QMP API type.
type FuseExportAllowOther int

// Known values of FuseExportAllowOther.
const (
	FuseExportAllowOtherOff FuseExportAllowOther = iota
	FuseExportAllowOtherOn
	FuseExportAllowOtherAuto
)

// String implements fmt.Stringer.
func (e FuseExportAllowOther) String() string {
	switch e {
	case FuseExportAllowOtherOff:
		return "off"
	case FuseExportAllowOtherOn:
		return "on"
	case FuseExportAllowOtherAuto:
		return "auto"
	default:
		return fmt.Sprintf("FuseExportAllowOther(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e FuseExportAllowOther) MarshalJSON() ([]byte, error) {
	switch e {
	case FuseExportAllowOtherOff:
		return json.Marshal("off")
	case FuseExportAllowOtherOn:
		return json.Marshal("on")
	case FuseExportAllowOtherAuto:
		return json.Marshal("auto")
	default:
		return nil, fmt.Errorf("unknown enum value %q for FuseExportAllowOther", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *FuseExportAllowOther) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "off":
		*e = FuseExportAllowOtherOff
	case "on":
		*e = FuseExportAllowOtherOn
	case "auto":
		*e = FuseExportAllowOtherAuto
	default:
		return fmt.Errorf("unknown enum value %q for FuseExportAllowOther", s)
	}
	return nil
}

// GICCapability -> GicCapability (struct)

// GicCapability implements the "GICCapability" QMP API type.
type GicCapability struct {
	Version  int64 `json:"version"`
	Emulated bool  `json:"emulated"`
	Kernel   bool  `json:"kernel"`
}

// EVENT GUEST_CRASHLOADED

// EVENT GUEST_PANICKED

// GrabToggleKeys -> GrabToggleKeys (enum)

// GrabToggleKeys implements the "GrabToggleKeys" QMP API type.
type GrabToggleKeys int

// Known values of GrabToggleKeys.
const (
	GrabToggleKeysCtrlCtrl GrabToggleKeys = iota
	GrabToggleKeysAltAlt
	GrabToggleKeysShiftShift
	GrabToggleKeysMetaMeta
	GrabToggleKeysScrolllock
	GrabToggleKeysCtrlScrolllock
)

// String implements fmt.Stringer.
func (e GrabToggleKeys) String() string {
	switch e {
	case GrabToggleKeysCtrlCtrl:
		return "ctrl-ctrl"
	case GrabToggleKeysAltAlt:
		return "alt-alt"
	case GrabToggleKeysShiftShift:
		return "shift-shift"
	case GrabToggleKeysMetaMeta:
		return "meta-meta"
	case GrabToggleKeysScrolllock:
		return "scrolllock"
	case GrabToggleKeysCtrlScrolllock:
		return "ctrl-scrolllock"
	default:
		return fmt.Sprintf("GrabToggleKeys(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e GrabToggleKeys) MarshalJSON() ([]byte, error) {
	switch e {
	case GrabToggleKeysCtrlCtrl:
		return json.Marshal("ctrl-ctrl")
	case GrabToggleKeysAltAlt:
		return json.Marshal("alt-alt")
	case GrabToggleKeysShiftShift:
		return json.Marshal("shift-shift")
	case GrabToggleKeysMetaMeta:
		return json.Marshal("meta-meta")
	case GrabToggleKeysScrolllock:
		return json.Marshal("scrolllock")
	case GrabToggleKeysCtrlScrolllock:
		return json.Marshal("ctrl-scrolllock")
	default:
		return nil, fmt.Errorf("unknown enum value %q for GrabToggleKeys", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *GrabToggleKeys) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "ctrl-ctrl":
		*e = GrabToggleKeysCtrlCtrl
	case "alt-alt":
		*e = GrabToggleKeysAltAlt
	case "shift-shift":
		*e = GrabToggleKeysShiftShift
	case "meta-meta":
		*e = GrabToggleKeysMetaMeta
	case "scrolllock":
		*e = GrabToggleKeysScrolllock
	case "ctrl-scrolllock":
		*e = GrabToggleKeysCtrlScrolllock
	default:
		return fmt.Errorf("unknown enum value %q for GrabToggleKeys", s)
	}
	return nil
}

// GuestPanicAction -> GuestPanicAction (enum)

// GuestPanicAction implements the "GuestPanicAction" QMP API type.
type GuestPanicAction int

// Known values of GuestPanicAction.
const (
	GuestPanicActionPause GuestPanicAction = iota
	GuestPanicActionPoweroff
	GuestPanicActionRun
)

// String implements fmt.Stringer.
func (e GuestPanicAction) String() string {
	switch e {
	case GuestPanicActionPause:
		return "pause"
	case GuestPanicActionPoweroff:
		return "poweroff"
	case GuestPanicActionRun:
		return "run"
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
	case GuestPanicActionRun:
		return json.Marshal("run")
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
	case "run":
		*e = GuestPanicActionRun
	default:
		return fmt.Errorf("unknown enum value %q for GuestPanicAction", s)
	}
	return nil
}

// GuestPanicInformation -> GuestPanicInformation (flat union)

// GuestPanicInformation implements the "GuestPanicInformation" QMP API type.
//
// Can be one of:
//   - GuestPanicInformationVariantHyperV
//   - GuestPanicInformationVariantS390
type GuestPanicInformation interface {
	isGuestPanicInformation()
}

// GuestPanicInformationVariantHyperV is an implementation of GuestPanicInformation.
type GuestPanicInformationVariantHyperV struct {
	Arg1 uint64 `json:"arg1"`
	Arg2 uint64 `json:"arg2"`
	Arg3 uint64 `json:"arg3"`
	Arg4 uint64 `json:"arg4"`
	Arg5 uint64 `json:"arg5"`
}

func (GuestPanicInformationVariantHyperV) isGuestPanicInformation() {}

// MarshalJSON implements json.Marshaler.
func (s GuestPanicInformationVariantHyperV) MarshalJSON() ([]byte, error) {
	v := struct {
		Type GuestPanicInformationType `json:"type"`
		GuestPanicInformationVariantHyperV
	}{
		GuestPanicInformationTypeHyperV,
		s,
	}
	return json.Marshal(v)
}

// GuestPanicInformationVariantS390 is an implementation of GuestPanicInformation.
type GuestPanicInformationVariantS390 struct {
	Core    uint32          `json:"core"`
	PswMask uint64          `json:"psw-mask"`
	PswAddr uint64          `json:"psw-addr"`
	Reason  S390CrashReason `json:"reason"`
}

func (GuestPanicInformationVariantS390) isGuestPanicInformation() {}

// MarshalJSON implements json.Marshaler.
func (s GuestPanicInformationVariantS390) MarshalJSON() ([]byte, error) {
	v := struct {
		Type GuestPanicInformationType `json:"type"`
		GuestPanicInformationVariantS390
	}{
		GuestPanicInformationTypeS390,
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
		var ret GuestPanicInformationVariantHyperV
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case GuestPanicInformationTypeS390:
		var ret GuestPanicInformationVariantS390
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
	GuestPanicInformationTypeS390
)

// String implements fmt.Stringer.
func (e GuestPanicInformationType) String() string {
	switch e {
	case GuestPanicInformationTypeHyperV:
		return "hyper-v"
	case GuestPanicInformationTypeS390:
		return "s390"
	default:
		return fmt.Sprintf("GuestPanicInformationType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e GuestPanicInformationType) MarshalJSON() ([]byte, error) {
	switch e {
	case GuestPanicInformationTypeHyperV:
		return json.Marshal("hyper-v")
	case GuestPanicInformationTypeS390:
		return json.Marshal("s390")
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
	case "s390":
		*e = GuestPanicInformationTypeS390
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

// HmatCacheAssociativity -> HmatCacheAssociativity (enum)

// HmatCacheAssociativity implements the "HmatCacheAssociativity" QMP API type.
type HmatCacheAssociativity int

// Known values of HmatCacheAssociativity.
const (
	HmatCacheAssociativityNone HmatCacheAssociativity = iota
	HmatCacheAssociativityDirect
	HmatCacheAssociativityComplex
)

// String implements fmt.Stringer.
func (e HmatCacheAssociativity) String() string {
	switch e {
	case HmatCacheAssociativityNone:
		return "none"
	case HmatCacheAssociativityDirect:
		return "direct"
	case HmatCacheAssociativityComplex:
		return "complex"
	default:
		return fmt.Sprintf("HmatCacheAssociativity(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e HmatCacheAssociativity) MarshalJSON() ([]byte, error) {
	switch e {
	case HmatCacheAssociativityNone:
		return json.Marshal("none")
	case HmatCacheAssociativityDirect:
		return json.Marshal("direct")
	case HmatCacheAssociativityComplex:
		return json.Marshal("complex")
	default:
		return nil, fmt.Errorf("unknown enum value %q for HmatCacheAssociativity", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *HmatCacheAssociativity) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = HmatCacheAssociativityNone
	case "direct":
		*e = HmatCacheAssociativityDirect
	case "complex":
		*e = HmatCacheAssociativityComplex
	default:
		return fmt.Errorf("unknown enum value %q for HmatCacheAssociativity", s)
	}
	return nil
}

// HmatCacheWritePolicy -> HmatCacheWritePolicy (enum)

// HmatCacheWritePolicy implements the "HmatCacheWritePolicy" QMP API type.
type HmatCacheWritePolicy int

// Known values of HmatCacheWritePolicy.
const (
	HmatCacheWritePolicyNone HmatCacheWritePolicy = iota
	HmatCacheWritePolicyWriteBack
	HmatCacheWritePolicyWriteThrough
)

// String implements fmt.Stringer.
func (e HmatCacheWritePolicy) String() string {
	switch e {
	case HmatCacheWritePolicyNone:
		return "none"
	case HmatCacheWritePolicyWriteBack:
		return "write-back"
	case HmatCacheWritePolicyWriteThrough:
		return "write-through"
	default:
		return fmt.Sprintf("HmatCacheWritePolicy(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e HmatCacheWritePolicy) MarshalJSON() ([]byte, error) {
	switch e {
	case HmatCacheWritePolicyNone:
		return json.Marshal("none")
	case HmatCacheWritePolicyWriteBack:
		return json.Marshal("write-back")
	case HmatCacheWritePolicyWriteThrough:
		return json.Marshal("write-through")
	default:
		return nil, fmt.Errorf("unknown enum value %q for HmatCacheWritePolicy", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *HmatCacheWritePolicy) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = HmatCacheWritePolicyNone
	case "write-back":
		*e = HmatCacheWritePolicyWriteBack
	case "write-through":
		*e = HmatCacheWritePolicyWriteThrough
	default:
		return fmt.Errorf("unknown enum value %q for HmatCacheWritePolicy", s)
	}
	return nil
}

// HmatLBDataType -> HmatLbDataType (enum)

// HmatLbDataType implements the "HmatLBDataType" QMP API type.
type HmatLbDataType int

// Known values of HmatLbDataType.
const (
	HmatLbDataTypeAccessLatency HmatLbDataType = iota
	HmatLbDataTypeReadLatency
	HmatLbDataTypeWriteLatency
	HmatLbDataTypeAccessBandwidth
	HmatLbDataTypeReadBandwidth
	HmatLbDataTypeWriteBandwidth
)

// String implements fmt.Stringer.
func (e HmatLbDataType) String() string {
	switch e {
	case HmatLbDataTypeAccessLatency:
		return "access-latency"
	case HmatLbDataTypeReadLatency:
		return "read-latency"
	case HmatLbDataTypeWriteLatency:
		return "write-latency"
	case HmatLbDataTypeAccessBandwidth:
		return "access-bandwidth"
	case HmatLbDataTypeReadBandwidth:
		return "read-bandwidth"
	case HmatLbDataTypeWriteBandwidth:
		return "write-bandwidth"
	default:
		return fmt.Sprintf("HmatLbDataType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e HmatLbDataType) MarshalJSON() ([]byte, error) {
	switch e {
	case HmatLbDataTypeAccessLatency:
		return json.Marshal("access-latency")
	case HmatLbDataTypeReadLatency:
		return json.Marshal("read-latency")
	case HmatLbDataTypeWriteLatency:
		return json.Marshal("write-latency")
	case HmatLbDataTypeAccessBandwidth:
		return json.Marshal("access-bandwidth")
	case HmatLbDataTypeReadBandwidth:
		return json.Marshal("read-bandwidth")
	case HmatLbDataTypeWriteBandwidth:
		return json.Marshal("write-bandwidth")
	default:
		return nil, fmt.Errorf("unknown enum value %q for HmatLbDataType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *HmatLbDataType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "access-latency":
		*e = HmatLbDataTypeAccessLatency
	case "read-latency":
		*e = HmatLbDataTypeReadLatency
	case "write-latency":
		*e = HmatLbDataTypeWriteLatency
	case "access-bandwidth":
		*e = HmatLbDataTypeAccessBandwidth
	case "read-bandwidth":
		*e = HmatLbDataTypeReadBandwidth
	case "write-bandwidth":
		*e = HmatLbDataTypeWriteBandwidth
	default:
		return fmt.Errorf("unknown enum value %q for HmatLbDataType", s)
	}
	return nil
}

// HmatLBMemoryHierarchy -> HmatLbMemoryHierarchy (enum)

// HmatLbMemoryHierarchy implements the "HmatLBMemoryHierarchy" QMP API type.
type HmatLbMemoryHierarchy int

// Known values of HmatLbMemoryHierarchy.
const (
	HmatLbMemoryHierarchyMemory HmatLbMemoryHierarchy = iota
	HmatLbMemoryHierarchyFirstLevel
	HmatLbMemoryHierarchySecondLevel
	HmatLbMemoryHierarchyThirdLevel
)

// String implements fmt.Stringer.
func (e HmatLbMemoryHierarchy) String() string {
	switch e {
	case HmatLbMemoryHierarchyMemory:
		return "memory"
	case HmatLbMemoryHierarchyFirstLevel:
		return "first-level"
	case HmatLbMemoryHierarchySecondLevel:
		return "second-level"
	case HmatLbMemoryHierarchyThirdLevel:
		return "third-level"
	default:
		return fmt.Sprintf("HmatLbMemoryHierarchy(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e HmatLbMemoryHierarchy) MarshalJSON() ([]byte, error) {
	switch e {
	case HmatLbMemoryHierarchyMemory:
		return json.Marshal("memory")
	case HmatLbMemoryHierarchyFirstLevel:
		return json.Marshal("first-level")
	case HmatLbMemoryHierarchySecondLevel:
		return json.Marshal("second-level")
	case HmatLbMemoryHierarchyThirdLevel:
		return json.Marshal("third-level")
	default:
		return nil, fmt.Errorf("unknown enum value %q for HmatLbMemoryHierarchy", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *HmatLbMemoryHierarchy) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "memory":
		*e = HmatLbMemoryHierarchyMemory
	case "first-level":
		*e = HmatLbMemoryHierarchyFirstLevel
	case "second-level":
		*e = HmatLbMemoryHierarchySecondLevel
	case "third-level":
		*e = HmatLbMemoryHierarchyThirdLevel
	default:
		return fmt.Errorf("unknown enum value %q for HmatLbMemoryHierarchy", s)
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

// HotKeyMod -> HotKeyMod (enum)

// HotKeyMod implements the "HotKeyMod" QMP API type.
type HotKeyMod int

// Known values of HotKeyMod.
const (
	HotKeyModLctrlLalt HotKeyMod = iota
	HotKeyModLshiftLctrlLalt
	HotKeyModRctrl
)

// String implements fmt.Stringer.
func (e HotKeyMod) String() string {
	switch e {
	case HotKeyModLctrlLalt:
		return "lctrl-lalt"
	case HotKeyModLshiftLctrlLalt:
		return "lshift-lctrl-lalt"
	case HotKeyModRctrl:
		return "rctrl"
	default:
		return fmt.Sprintf("HotKeyMod(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e HotKeyMod) MarshalJSON() ([]byte, error) {
	switch e {
	case HotKeyModLctrlLalt:
		return json.Marshal("lctrl-lalt")
	case HotKeyModLshiftLctrlLalt:
		return json.Marshal("lshift-lctrl-lalt")
	case HotKeyModRctrl:
		return json.Marshal("rctrl")
	default:
		return nil, fmt.Errorf("unknown enum value %q for HotKeyMod", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *HotKeyMod) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "lctrl-lalt":
		*e = HotKeyModLctrlLalt
	case "lshift-lctrl-lalt":
		*e = HotKeyModLshiftLctrlLalt
	case "rctrl":
		*e = HotKeyModRctrl
	default:
		return fmt.Errorf("unknown enum value %q for HotKeyMod", s)
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

// HumanReadableText -> HumanReadableText (struct)

// HumanReadableText implements the "HumanReadableText" QMP API type.
type HumanReadableText struct {
	HumanReadableText string `json:"human-readable-text"`
}

// IOThreadInfo -> IOThreadInfo (struct)

// IOThreadInfo implements the "IOThreadInfo" QMP API type.
type IOThreadInfo struct {
	ID          string `json:"id"`
	ThreadID    int64  `json:"thread-id"`
	PollMaxNs   int64  `json:"poll-max-ns"`
	PollGrow    int64  `json:"poll-grow"`
	PollShrink  int64  `json:"poll-shrink"`
	AIOMaxBatch int64  `json:"aio-max-batch"`
}

// ImageFormat -> ImageFormat (enum)

// ImageFormat implements the "ImageFormat" QMP API type.
type ImageFormat int

// Known values of ImageFormat.
const (
	ImageFormatPpm ImageFormat = iota
	ImageFormatPng
)

// String implements fmt.Stringer.
func (e ImageFormat) String() string {
	switch e {
	case ImageFormatPpm:
		return "ppm"
	case ImageFormatPng:
		return "png"
	default:
		return fmt.Sprintf("ImageFormat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ImageFormat) MarshalJSON() ([]byte, error) {
	switch e {
	case ImageFormatPpm:
		return json.Marshal("ppm")
	case ImageFormatPng:
		return json.Marshal("png")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ImageFormat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ImageFormat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "ppm":
		*e = ImageFormatPpm
	case "png":
		*e = ImageFormatPng
	default:
		return fmt.Errorf("unknown enum value %q for ImageFormat", s)
	}
	return nil
}

// ImageInfo -> ImageInfo (struct)

// ImageInfo implements the "ImageInfo" QMP API type.
type ImageInfo struct {
	Filename              string             `json:"filename"`
	Format                string             `json:"format"`
	DirtyFlag             *bool              `json:"dirty-flag,omitempty"`
	ActualSize            *int64             `json:"actual-size,omitempty"`
	VirtualSize           int64              `json:"virtual-size"`
	ClusterSize           *int64             `json:"cluster-size,omitempty"`
	Encrypted             *bool              `json:"encrypted,omitempty"`
	Compressed            *bool              `json:"compressed,omitempty"`
	BackingFilename       *string            `json:"backing-filename,omitempty"`
	FullBackingFilename   *string            `json:"full-backing-filename,omitempty"`
	BackingFilenameFormat *string            `json:"backing-filename-format,omitempty"`
	Snapshots             []SnapshotInfo     `json:"snapshots,omitempty"`
	BackingImage          *ImageInfo         `json:"backing-image,omitempty"`
	FormatSpecific        *ImageInfoSpecific `json:"format-specific,omitempty"`
}

// ImageInfoSpecific -> ImageInfoSpecific (flat union)

// ImageInfoSpecific implements the "ImageInfoSpecific" QMP API type.
//
// Can be one of:
//   - ImageInfoSpecificVariantLUKS
//   - ImageInfoSpecificVariantQcow2
//   - ImageInfoSpecificVariantRbd
//   - ImageInfoSpecificVariantVMDK
type ImageInfoSpecific interface {
	isImageInfoSpecific()
}

// ImageInfoSpecificVariantLUKS is an implementation of ImageInfoSpecific.
type ImageInfoSpecificVariantLUKS struct {
	Data QCryptoBlockInfoLUKS `json:"data"`
}

func (ImageInfoSpecificVariantLUKS) isImageInfoSpecific() {}

// MarshalJSON implements json.Marshaler.
func (s ImageInfoSpecificVariantLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ImageInfoSpecificKind `json:"type"`
		ImageInfoSpecificVariantLUKS
	}{
		ImageInfoSpecificKindLUKS,
		s,
	}
	return json.Marshal(v)
}

// ImageInfoSpecificVariantQcow2 is an implementation of ImageInfoSpecific.
type ImageInfoSpecificVariantQcow2 struct {
	Data ImageInfoSpecificQCow2 `json:"data"`
}

func (ImageInfoSpecificVariantQcow2) isImageInfoSpecific() {}

// MarshalJSON implements json.Marshaler.
func (s ImageInfoSpecificVariantQcow2) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ImageInfoSpecificKind `json:"type"`
		ImageInfoSpecificVariantQcow2
	}{
		ImageInfoSpecificKindQcow2,
		s,
	}
	return json.Marshal(v)
}

// ImageInfoSpecificVariantRbd is an implementation of ImageInfoSpecific.
type ImageInfoSpecificVariantRbd struct {
	Data ImageInfoSpecificRbd `json:"data"`
}

func (ImageInfoSpecificVariantRbd) isImageInfoSpecific() {}

// MarshalJSON implements json.Marshaler.
func (s ImageInfoSpecificVariantRbd) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ImageInfoSpecificKind `json:"type"`
		ImageInfoSpecificVariantRbd
	}{
		ImageInfoSpecificKindRbd,
		s,
	}
	return json.Marshal(v)
}

// ImageInfoSpecificVariantVMDK is an implementation of ImageInfoSpecific.
type ImageInfoSpecificVariantVMDK struct {
	Data ImageInfoSpecificVMDK `json:"data"`
}

func (ImageInfoSpecificVariantVMDK) isImageInfoSpecific() {}

// MarshalJSON implements json.Marshaler.
func (s ImageInfoSpecificVariantVMDK) MarshalJSON() ([]byte, error) {
	v := struct {
		Type ImageInfoSpecificKind `json:"type"`
		ImageInfoSpecificVariantVMDK
	}{
		ImageInfoSpecificKindVMDK,
		s,
	}
	return json.Marshal(v)
}

func decodeImageInfoSpecific(bs json.RawMessage) (ImageInfoSpecific, error) {
	v := struct {
		Type ImageInfoSpecificKind `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case ImageInfoSpecificKindLUKS:
		var ret ImageInfoSpecificVariantLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ImageInfoSpecificKindQcow2:
		var ret ImageInfoSpecificVariantQcow2
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ImageInfoSpecificKindRbd:
		var ret ImageInfoSpecificVariantRbd
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ImageInfoSpecificKindVMDK:
		var ret ImageInfoSpecificVariantVMDK
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union ImageInfoSpecific", v.Type)
	}
}

// ImageInfoSpecificKind -> ImageInfoSpecificKind (enum)

// ImageInfoSpecificKind implements the "ImageInfoSpecificKind" QMP API type.
type ImageInfoSpecificKind int

// Known values of ImageInfoSpecificKind.
const (
	ImageInfoSpecificKindQcow2 ImageInfoSpecificKind = iota
	ImageInfoSpecificKindVMDK
	ImageInfoSpecificKindLUKS
	ImageInfoSpecificKindRbd
)

// String implements fmt.Stringer.
func (e ImageInfoSpecificKind) String() string {
	switch e {
	case ImageInfoSpecificKindQcow2:
		return "qcow2"
	case ImageInfoSpecificKindVMDK:
		return "vmdk"
	case ImageInfoSpecificKindLUKS:
		return "luks"
	case ImageInfoSpecificKindRbd:
		return "rbd"
	default:
		return fmt.Sprintf("ImageInfoSpecificKind(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ImageInfoSpecificKind) MarshalJSON() ([]byte, error) {
	switch e {
	case ImageInfoSpecificKindQcow2:
		return json.Marshal("qcow2")
	case ImageInfoSpecificKindVMDK:
		return json.Marshal("vmdk")
	case ImageInfoSpecificKindLUKS:
		return json.Marshal("luks")
	case ImageInfoSpecificKindRbd:
		return json.Marshal("rbd")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ImageInfoSpecificKind", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ImageInfoSpecificKind) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "qcow2":
		*e = ImageInfoSpecificKindQcow2
	case "vmdk":
		*e = ImageInfoSpecificKindVMDK
	case "luks":
		*e = ImageInfoSpecificKindLUKS
	case "rbd":
		*e = ImageInfoSpecificKindRbd
	default:
		return fmt.Errorf("unknown enum value %q for ImageInfoSpecificKind", s)
	}
	return nil
}

// ImageInfoSpecificQCow2 -> ImageInfoSpecificQCow2 (struct)

// ImageInfoSpecificQCow2 implements the "ImageInfoSpecificQCow2" QMP API type.
type ImageInfoSpecificQCow2 struct {
	Compat          string                            `json:"compat"`
	DataFile        *string                           `json:"data-file,omitempty"`
	DataFileRaw     *bool                             `json:"data-file-raw,omitempty"`
	ExtendedL2      *bool                             `json:"extended-l2,omitempty"`
	LazyRefcounts   *bool                             `json:"lazy-refcounts,omitempty"`
	Corrupt         *bool                             `json:"corrupt,omitempty"`
	RefcountBits    int64                             `json:"refcount-bits"`
	Encrypt         *ImageInfoSpecificQCow2Encryption `json:"encrypt,omitempty"`
	Bitmaps         []Qcow2BitmapInfo                 `json:"bitmaps,omitempty"`
	CompressionType Qcow2CompressionType              `json:"compression-type"`
}

// ImageInfoSpecificQCow2Encryption -> ImageInfoSpecificQCow2Encryption (flat union)

// ImageInfoSpecificQCow2Encryption implements the "ImageInfoSpecificQCow2Encryption" QMP API type.
//
// Can be one of:
//   - ImageInfoSpecificQCow2EncryptionVariantLUKS
type ImageInfoSpecificQCow2Encryption interface {
	isImageInfoSpecificQCow2Encryption()
}

// ImageInfoSpecificQCow2EncryptionVariantLUKS is an implementation of ImageInfoSpecificQCow2Encryption.
type ImageInfoSpecificQCow2EncryptionVariantLUKS struct {
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

func (ImageInfoSpecificQCow2EncryptionVariantLUKS) isImageInfoSpecificQCow2Encryption() {}

// MarshalJSON implements json.Marshaler.
func (s ImageInfoSpecificQCow2EncryptionVariantLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Format BlockdevQcow2EncryptionFormat `json:"format"`
		ImageInfoSpecificQCow2EncryptionVariantLUKS
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
	case BlockdevQcow2EncryptionFormatLUKS:
		var ret ImageInfoSpecificQCow2EncryptionVariantLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union ImageInfoSpecificQCow2Encryption", v.Format)
	}
}

// ImageInfoSpecificRbd -> ImageInfoSpecificRbd (struct)

// ImageInfoSpecificRbd implements the "ImageInfoSpecificRbd" QMP API type.
type ImageInfoSpecificRbd struct {
	EncryptionFormat *RbdImageEncryptionFormat `json:"encryption-format,omitempty"`
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
	Host      string  `json:"host"`
	Port      string  `json:"port"`
	Numeric   *bool   `json:"numeric,omitempty"`
	To        *uint16 `json:"to,omitempty"`
	Ipv4      *bool   `json:"ipv4,omitempty"`
	Ipv6      *bool   `json:"ipv6,omitempty"`
	KeepAlive *bool   `json:"keep-alive,omitempty"`
	Mptcp     *bool   `json:"mptcp,omitempty"`
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
	InputButtonWheelLeft
	InputButtonWheelRight
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
	case InputButtonWheelLeft:
		return "wheel-left"
	case InputButtonWheelRight:
		return "wheel-right"
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
	case InputButtonWheelLeft:
		return json.Marshal("wheel-left")
	case InputButtonWheelRight:
		return json.Marshal("wheel-right")
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
	case "wheel-left":
		*e = InputButtonWheelLeft
	case "wheel-right":
		*e = InputButtonWheelRight
	default:
		return fmt.Errorf("unknown enum value %q for InputButton", s)
	}
	return nil
}

// InputEvent -> InputEvent (flat union)

// InputEvent implements the "InputEvent" QMP API type.
//
// Can be one of:
//   - InputEventVariantAbs
//   - InputEventVariantBtn
//   - InputEventVariantKey
//   - InputEventVariantRel
type InputEvent interface {
	isInputEvent()
}

// InputEventVariantAbs is an implementation of InputEvent.
type InputEventVariantAbs struct {
	Data InputMoveEvent `json:"data"`
}

func (InputEventVariantAbs) isInputEvent() {}

// MarshalJSON implements json.Marshaler.
func (s InputEventVariantAbs) MarshalJSON() ([]byte, error) {
	v := struct {
		Type InputEventKind `json:"type"`
		InputEventVariantAbs
	}{
		InputEventKindAbs,
		s,
	}
	return json.Marshal(v)
}

// InputEventVariantBtn is an implementation of InputEvent.
type InputEventVariantBtn struct {
	Data InputBtnEvent `json:"data"`
}

func (InputEventVariantBtn) isInputEvent() {}

// MarshalJSON implements json.Marshaler.
func (s InputEventVariantBtn) MarshalJSON() ([]byte, error) {
	v := struct {
		Type InputEventKind `json:"type"`
		InputEventVariantBtn
	}{
		InputEventKindBtn,
		s,
	}
	return json.Marshal(v)
}

// InputEventVariantKey is an implementation of InputEvent.
type InputEventVariantKey struct {
	Data InputKeyEvent `json:"data"`
}

func (InputEventVariantKey) isInputEvent() {}

// MarshalJSON implements json.Marshaler.
func (s InputEventVariantKey) MarshalJSON() ([]byte, error) {
	v := struct {
		Type InputEventKind `json:"type"`
		InputEventVariantKey
	}{
		InputEventKindKey,
		s,
	}
	return json.Marshal(v)
}

// InputEventVariantRel is an implementation of InputEvent.
type InputEventVariantRel struct {
	Data InputMoveEvent `json:"data"`
}

func (InputEventVariantRel) isInputEvent() {}

// MarshalJSON implements json.Marshaler.
func (s InputEventVariantRel) MarshalJSON() ([]byte, error) {
	v := struct {
		Type InputEventKind `json:"type"`
		InputEventVariantRel
	}{
		InputEventKindRel,
		s,
	}
	return json.Marshal(v)
}

func decodeInputEvent(bs json.RawMessage) (InputEvent, error) {
	v := struct {
		Type InputEventKind `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case InputEventKindAbs:
		var ret InputEventVariantAbs
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case InputEventKindBtn:
		var ret InputEventVariantBtn
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case InputEventKindKey:
		var ret InputEventVariantKey
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case InputEventKindRel:
		var ret InputEventVariantRel
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union InputEvent", v.Type)
	}
}

// InputEventKind -> InputEventKind (enum)

// InputEventKind implements the "InputEventKind" QMP API type.
type InputEventKind int

// Known values of InputEventKind.
const (
	InputEventKindKey InputEventKind = iota
	InputEventKindBtn
	InputEventKindRel
	InputEventKindAbs
)

// String implements fmt.Stringer.
func (e InputEventKind) String() string {
	switch e {
	case InputEventKindKey:
		return "key"
	case InputEventKindBtn:
		return "btn"
	case InputEventKindRel:
		return "rel"
	case InputEventKindAbs:
		return "abs"
	default:
		return fmt.Sprintf("InputEventKind(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e InputEventKind) MarshalJSON() ([]byte, error) {
	switch e {
	case InputEventKindKey:
		return json.Marshal("key")
	case InputEventKindBtn:
		return json.Marshal("btn")
	case InputEventKindRel:
		return json.Marshal("rel")
	case InputEventKindAbs:
		return json.Marshal("abs")
	default:
		return nil, fmt.Errorf("unknown enum value %q for InputEventKind", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *InputEventKind) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "key":
		*e = InputEventKindKey
	case "btn":
		*e = InputEventKindBtn
	case "rel":
		*e = InputEventKindRel
	case "abs":
		*e = InputEventKindAbs
	default:
		return fmt.Errorf("unknown enum value %q for InputEventKind", s)
	}
	return nil
}

// InputKeyEvent -> InputKeyEvent (struct)

// InputKeyEvent implements the "InputKeyEvent" QMP API type.
type InputKeyEvent struct {
	Key  KeyValue `json:"key"`
	Down bool     `json:"down"`
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

// EVENT JOB_STATUS_CHANGE

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

// JobInfo -> JobInfo (struct)

// JobInfo implements the "JobInfo" QMP API type.
type JobInfo struct {
	ID              string    `json:"id"`
	Type            JobType   `json:"type"`
	Status          JobStatus `json:"status"`
	CurrentProgress int64     `json:"current-progress"`
	TotalProgress   int64     `json:"total-progress"`
	Error           *string   `json:"error,omitempty"`
}

// JobStatus -> JobStatus (enum)

// JobStatus implements the "JobStatus" QMP API type.
type JobStatus int

// Known values of JobStatus.
const (
	JobStatusUndefined JobStatus = iota
	JobStatusCreated
	JobStatusRunning
	JobStatusPaused
	JobStatusReady
	JobStatusStandby
	JobStatusWaiting
	JobStatusPending
	JobStatusAborting
	JobStatusConcluded
	JobStatusNull
)

// String implements fmt.Stringer.
func (e JobStatus) String() string {
	switch e {
	case JobStatusUndefined:
		return "undefined"
	case JobStatusCreated:
		return "created"
	case JobStatusRunning:
		return "running"
	case JobStatusPaused:
		return "paused"
	case JobStatusReady:
		return "ready"
	case JobStatusStandby:
		return "standby"
	case JobStatusWaiting:
		return "waiting"
	case JobStatusPending:
		return "pending"
	case JobStatusAborting:
		return "aborting"
	case JobStatusConcluded:
		return "concluded"
	case JobStatusNull:
		return "null"
	default:
		return fmt.Sprintf("JobStatus(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e JobStatus) MarshalJSON() ([]byte, error) {
	switch e {
	case JobStatusUndefined:
		return json.Marshal("undefined")
	case JobStatusCreated:
		return json.Marshal("created")
	case JobStatusRunning:
		return json.Marshal("running")
	case JobStatusPaused:
		return json.Marshal("paused")
	case JobStatusReady:
		return json.Marshal("ready")
	case JobStatusStandby:
		return json.Marshal("standby")
	case JobStatusWaiting:
		return json.Marshal("waiting")
	case JobStatusPending:
		return json.Marshal("pending")
	case JobStatusAborting:
		return json.Marshal("aborting")
	case JobStatusConcluded:
		return json.Marshal("concluded")
	case JobStatusNull:
		return json.Marshal("null")
	default:
		return nil, fmt.Errorf("unknown enum value %q for JobStatus", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *JobStatus) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "undefined":
		*e = JobStatusUndefined
	case "created":
		*e = JobStatusCreated
	case "running":
		*e = JobStatusRunning
	case "paused":
		*e = JobStatusPaused
	case "ready":
		*e = JobStatusReady
	case "standby":
		*e = JobStatusStandby
	case "waiting":
		*e = JobStatusWaiting
	case "pending":
		*e = JobStatusPending
	case "aborting":
		*e = JobStatusAborting
	case "concluded":
		*e = JobStatusConcluded
	case "null":
		*e = JobStatusNull
	default:
		return fmt.Errorf("unknown enum value %q for JobStatus", s)
	}
	return nil
}

// JobType -> JobType (enum)

// JobType implements the "JobType" QMP API type.
type JobType int

// Known values of JobType.
const (
	JobTypeCommit JobType = iota
	JobTypeStream
	JobTypeMirror
	JobTypeBackup
	JobTypeCreate
	JobTypeAmend
	JobTypeSnapshotLoad
	JobTypeSnapshotSave
	JobTypeSnapshotDelete
)

// String implements fmt.Stringer.
func (e JobType) String() string {
	switch e {
	case JobTypeCommit:
		return "commit"
	case JobTypeStream:
		return "stream"
	case JobTypeMirror:
		return "mirror"
	case JobTypeBackup:
		return "backup"
	case JobTypeCreate:
		return "create"
	case JobTypeAmend:
		return "amend"
	case JobTypeSnapshotLoad:
		return "snapshot-load"
	case JobTypeSnapshotSave:
		return "snapshot-save"
	case JobTypeSnapshotDelete:
		return "snapshot-delete"
	default:
		return fmt.Sprintf("JobType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e JobType) MarshalJSON() ([]byte, error) {
	switch e {
	case JobTypeCommit:
		return json.Marshal("commit")
	case JobTypeStream:
		return json.Marshal("stream")
	case JobTypeMirror:
		return json.Marshal("mirror")
	case JobTypeBackup:
		return json.Marshal("backup")
	case JobTypeCreate:
		return json.Marshal("create")
	case JobTypeAmend:
		return json.Marshal("amend")
	case JobTypeSnapshotLoad:
		return json.Marshal("snapshot-load")
	case JobTypeSnapshotSave:
		return json.Marshal("snapshot-save")
	case JobTypeSnapshotDelete:
		return json.Marshal("snapshot-delete")
	default:
		return nil, fmt.Errorf("unknown enum value %q for JobType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *JobType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "commit":
		*e = JobTypeCommit
	case "stream":
		*e = JobTypeStream
	case "mirror":
		*e = JobTypeMirror
	case "backup":
		*e = JobTypeBackup
	case "create":
		*e = JobTypeCreate
	case "amend":
		*e = JobTypeAmend
	case "snapshot-load":
		*e = JobTypeSnapshotLoad
	case "snapshot-save":
		*e = JobTypeSnapshotSave
	case "snapshot-delete":
		*e = JobTypeSnapshotDelete
	default:
		return fmt.Errorf("unknown enum value %q for JobType", s)
	}
	return nil
}

// KeyValue -> KeyValue (flat union)

// KeyValue implements the "KeyValue" QMP API type.
//
// Can be one of:
//   - KeyValueVariantfloat64
//   - KeyValueVariantQcode
type KeyValue interface {
	isKeyValue()
}

// KeyValueVariantfloat64 is an implementation of KeyValue.
type KeyValueVariantfloat64 struct {
	Data int64 `json:"data"`
}

func (KeyValueVariantfloat64) isKeyValue() {}

// MarshalJSON implements json.Marshaler.
func (s KeyValueVariantfloat64) MarshalJSON() ([]byte, error) {
	v := struct {
		Type KeyValueKind `json:"type"`
		KeyValueVariantfloat64
	}{
		KeyValueKindfloat64,
		s,
	}
	return json.Marshal(v)
}

// KeyValueVariantQcode is an implementation of KeyValue.
type KeyValueVariantQcode struct {
	Data QKeyCode `json:"data"`
}

func (KeyValueVariantQcode) isKeyValue() {}

// MarshalJSON implements json.Marshaler.
func (s KeyValueVariantQcode) MarshalJSON() ([]byte, error) {
	v := struct {
		Type KeyValueKind `json:"type"`
		KeyValueVariantQcode
	}{
		KeyValueKindQcode,
		s,
	}
	return json.Marshal(v)
}

func decodeKeyValue(bs json.RawMessage) (KeyValue, error) {
	v := struct {
		Type KeyValueKind `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case KeyValueKindfloat64:
		var ret KeyValueVariantfloat64
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case KeyValueKindQcode:
		var ret KeyValueVariantQcode
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union KeyValue", v.Type)
	}
}

// KeyValueKind -> KeyValueKind (enum)

// KeyValueKind implements the "KeyValueKind" QMP API type.
type KeyValueKind int

// Known values of KeyValueKind.
const (
	KeyValueKindfloat64 KeyValueKind = iota
	KeyValueKindQcode
)

// String implements fmt.Stringer.
func (e KeyValueKind) String() string {
	switch e {
	case KeyValueKindfloat64:
		return "number"
	case KeyValueKindQcode:
		return "qcode"
	default:
		return fmt.Sprintf("KeyValueKind(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e KeyValueKind) MarshalJSON() ([]byte, error) {
	switch e {
	case KeyValueKindfloat64:
		return json.Marshal("number")
	case KeyValueKindQcode:
		return json.Marshal("qcode")
	default:
		return nil, fmt.Errorf("unknown enum value %q for KeyValueKind", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *KeyValueKind) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "number":
		*e = KeyValueKindfloat64
	case "qcode":
		*e = KeyValueKindQcode
	default:
		return fmt.Errorf("unknown enum value %q for KeyValueKind", s)
	}
	return nil
}

// KvmInfo -> KVMInfo (struct)

// KVMInfo implements the "KvmInfo" QMP API type.
type KVMInfo struct {
	Enabled bool `json:"enabled"`
	Present bool `json:"present"`
}

// EVENT MEMORY_DEVICE_SIZE_CHANGE

// EVENT MEMORY_FAILURE

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
	NumaMemSupported bool    `json:"numa-mem-supported"`
	Deprecated       bool    `json:"deprecated"`
	DefaultCPUType   *string `json:"default-cpu-type,omitempty"`
	DefaultRAMID     *string `json:"default-ram-id,omitempty"`
}

// Memdev -> Memdev (struct)

// Memdev implements the "Memdev" QMP API type.
type Memdev struct {
	ID        *string       `json:"id,omitempty"`
	Size      uint64        `json:"size"`
	Merge     bool          `json:"merge"`
	Dump      bool          `json:"dump"`
	Prealloc  bool          `json:"prealloc"`
	Share     bool          `json:"share"`
	Reserve   *bool         `json:"reserve,omitempty"`
	HostNodes []uint16      `json:"host-nodes"`
	Policy    HostMemPolicy `json:"policy"`
}

// MemoryDeviceInfo -> MemoryDeviceInfo (flat union)

// MemoryDeviceInfo implements the "MemoryDeviceInfo" QMP API type.
//
// Can be one of:
//   - MemoryDeviceInfoVariantDimm
//   - MemoryDeviceInfoVariantNvdimm
//   - MemoryDeviceInfoVariantSgxEpc
//   - MemoryDeviceInfoVariantVirtioMem
//   - MemoryDeviceInfoVariantVirtioPmem
type MemoryDeviceInfo interface {
	isMemoryDeviceInfo()
}

// MemoryDeviceInfoVariantDimm is an implementation of MemoryDeviceInfo.
type MemoryDeviceInfoVariantDimm struct {
	Data PcdimmDeviceInfo `json:"data"`
}

func (MemoryDeviceInfoVariantDimm) isMemoryDeviceInfo() {}

// MarshalJSON implements json.Marshaler.
func (s MemoryDeviceInfoVariantDimm) MarshalJSON() ([]byte, error) {
	v := struct {
		Type MemoryDeviceInfoKind `json:"type"`
		MemoryDeviceInfoVariantDimm
	}{
		MemoryDeviceInfoKindDimm,
		s,
	}
	return json.Marshal(v)
}

// MemoryDeviceInfoVariantNvdimm is an implementation of MemoryDeviceInfo.
type MemoryDeviceInfoVariantNvdimm struct {
	Data PcdimmDeviceInfo `json:"data"`
}

func (MemoryDeviceInfoVariantNvdimm) isMemoryDeviceInfo() {}

// MarshalJSON implements json.Marshaler.
func (s MemoryDeviceInfoVariantNvdimm) MarshalJSON() ([]byte, error) {
	v := struct {
		Type MemoryDeviceInfoKind `json:"type"`
		MemoryDeviceInfoVariantNvdimm
	}{
		MemoryDeviceInfoKindNvdimm,
		s,
	}
	return json.Marshal(v)
}

// MemoryDeviceInfoVariantSgxEpc is an implementation of MemoryDeviceInfo.
type MemoryDeviceInfoVariantSgxEpc struct {
	Data SgxEpcDeviceInfo `json:"data"`
}

func (MemoryDeviceInfoVariantSgxEpc) isMemoryDeviceInfo() {}

// MarshalJSON implements json.Marshaler.
func (s MemoryDeviceInfoVariantSgxEpc) MarshalJSON() ([]byte, error) {
	v := struct {
		Type MemoryDeviceInfoKind `json:"type"`
		MemoryDeviceInfoVariantSgxEpc
	}{
		MemoryDeviceInfoKindSgxEpc,
		s,
	}
	return json.Marshal(v)
}

// MemoryDeviceInfoVariantVirtioMem is an implementation of MemoryDeviceInfo.
type MemoryDeviceInfoVariantVirtioMem struct {
	Data VirtioMemDeviceInfo `json:"data"`
}

func (MemoryDeviceInfoVariantVirtioMem) isMemoryDeviceInfo() {}

// MarshalJSON implements json.Marshaler.
func (s MemoryDeviceInfoVariantVirtioMem) MarshalJSON() ([]byte, error) {
	v := struct {
		Type MemoryDeviceInfoKind `json:"type"`
		MemoryDeviceInfoVariantVirtioMem
	}{
		MemoryDeviceInfoKindVirtioMem,
		s,
	}
	return json.Marshal(v)
}

// MemoryDeviceInfoVariantVirtioPmem is an implementation of MemoryDeviceInfo.
type MemoryDeviceInfoVariantVirtioPmem struct {
	Data VirtioPmemDeviceInfo `json:"data"`
}

func (MemoryDeviceInfoVariantVirtioPmem) isMemoryDeviceInfo() {}

// MarshalJSON implements json.Marshaler.
func (s MemoryDeviceInfoVariantVirtioPmem) MarshalJSON() ([]byte, error) {
	v := struct {
		Type MemoryDeviceInfoKind `json:"type"`
		MemoryDeviceInfoVariantVirtioPmem
	}{
		MemoryDeviceInfoKindVirtioPmem,
		s,
	}
	return json.Marshal(v)
}

func decodeMemoryDeviceInfo(bs json.RawMessage) (MemoryDeviceInfo, error) {
	v := struct {
		Type MemoryDeviceInfoKind `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case MemoryDeviceInfoKindDimm:
		var ret MemoryDeviceInfoVariantDimm
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case MemoryDeviceInfoKindNvdimm:
		var ret MemoryDeviceInfoVariantNvdimm
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case MemoryDeviceInfoKindSgxEpc:
		var ret MemoryDeviceInfoVariantSgxEpc
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case MemoryDeviceInfoKindVirtioMem:
		var ret MemoryDeviceInfoVariantVirtioMem
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case MemoryDeviceInfoKindVirtioPmem:
		var ret MemoryDeviceInfoVariantVirtioPmem
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union MemoryDeviceInfo", v.Type)
	}
}

// MemoryDeviceInfoKind -> MemoryDeviceInfoKind (enum)

// MemoryDeviceInfoKind implements the "MemoryDeviceInfoKind" QMP API type.
type MemoryDeviceInfoKind int

// Known values of MemoryDeviceInfoKind.
const (
	MemoryDeviceInfoKindDimm MemoryDeviceInfoKind = iota
	MemoryDeviceInfoKindNvdimm
	MemoryDeviceInfoKindVirtioPmem
	MemoryDeviceInfoKindVirtioMem
	MemoryDeviceInfoKindSgxEpc
)

// String implements fmt.Stringer.
func (e MemoryDeviceInfoKind) String() string {
	switch e {
	case MemoryDeviceInfoKindDimm:
		return "dimm"
	case MemoryDeviceInfoKindNvdimm:
		return "nvdimm"
	case MemoryDeviceInfoKindVirtioPmem:
		return "virtio-pmem"
	case MemoryDeviceInfoKindVirtioMem:
		return "virtio-mem"
	case MemoryDeviceInfoKindSgxEpc:
		return "sgx-epc"
	default:
		return fmt.Sprintf("MemoryDeviceInfoKind(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e MemoryDeviceInfoKind) MarshalJSON() ([]byte, error) {
	switch e {
	case MemoryDeviceInfoKindDimm:
		return json.Marshal("dimm")
	case MemoryDeviceInfoKindNvdimm:
		return json.Marshal("nvdimm")
	case MemoryDeviceInfoKindVirtioPmem:
		return json.Marshal("virtio-pmem")
	case MemoryDeviceInfoKindVirtioMem:
		return json.Marshal("virtio-mem")
	case MemoryDeviceInfoKindSgxEpc:
		return json.Marshal("sgx-epc")
	default:
		return nil, fmt.Errorf("unknown enum value %q for MemoryDeviceInfoKind", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *MemoryDeviceInfoKind) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "dimm":
		*e = MemoryDeviceInfoKindDimm
	case "nvdimm":
		*e = MemoryDeviceInfoKindNvdimm
	case "virtio-pmem":
		*e = MemoryDeviceInfoKindVirtioPmem
	case "virtio-mem":
		*e = MemoryDeviceInfoKindVirtioMem
	case "sgx-epc":
		*e = MemoryDeviceInfoKindSgxEpc
	default:
		return fmt.Errorf("unknown enum value %q for MemoryDeviceInfoKind", s)
	}
	return nil
}

// MemoryFailureAction -> MemoryFailureAction (enum)

// MemoryFailureAction implements the "MemoryFailureAction" QMP API type.
type MemoryFailureAction int

// Known values of MemoryFailureAction.
const (
	MemoryFailureActionIgnore MemoryFailureAction = iota
	MemoryFailureActionInject
	MemoryFailureActionFatal
	MemoryFailureActionReset
)

// String implements fmt.Stringer.
func (e MemoryFailureAction) String() string {
	switch e {
	case MemoryFailureActionIgnore:
		return "ignore"
	case MemoryFailureActionInject:
		return "inject"
	case MemoryFailureActionFatal:
		return "fatal"
	case MemoryFailureActionReset:
		return "reset"
	default:
		return fmt.Sprintf("MemoryFailureAction(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e MemoryFailureAction) MarshalJSON() ([]byte, error) {
	switch e {
	case MemoryFailureActionIgnore:
		return json.Marshal("ignore")
	case MemoryFailureActionInject:
		return json.Marshal("inject")
	case MemoryFailureActionFatal:
		return json.Marshal("fatal")
	case MemoryFailureActionReset:
		return json.Marshal("reset")
	default:
		return nil, fmt.Errorf("unknown enum value %q for MemoryFailureAction", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *MemoryFailureAction) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "ignore":
		*e = MemoryFailureActionIgnore
	case "inject":
		*e = MemoryFailureActionInject
	case "fatal":
		*e = MemoryFailureActionFatal
	case "reset":
		*e = MemoryFailureActionReset
	default:
		return fmt.Errorf("unknown enum value %q for MemoryFailureAction", s)
	}
	return nil
}

// MemoryFailureFlags -> MemoryFailureFlags (struct)

// MemoryFailureFlags implements the "MemoryFailureFlags" QMP API type.
type MemoryFailureFlags struct {
	ActionRequired bool `json:"action-required"`
	Recursive      bool `json:"recursive"`
}

// MemoryFailureRecipient -> MemoryFailureRecipient (enum)

// MemoryFailureRecipient implements the "MemoryFailureRecipient" QMP API type.
type MemoryFailureRecipient int

// Known values of MemoryFailureRecipient.
const (
	MemoryFailureRecipientHypervisor MemoryFailureRecipient = iota
	MemoryFailureRecipientGuest
)

// String implements fmt.Stringer.
func (e MemoryFailureRecipient) String() string {
	switch e {
	case MemoryFailureRecipientHypervisor:
		return "hypervisor"
	case MemoryFailureRecipientGuest:
		return "guest"
	default:
		return fmt.Sprintf("MemoryFailureRecipient(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e MemoryFailureRecipient) MarshalJSON() ([]byte, error) {
	switch e {
	case MemoryFailureRecipientHypervisor:
		return json.Marshal("hypervisor")
	case MemoryFailureRecipientGuest:
		return json.Marshal("guest")
	default:
		return nil, fmt.Errorf("unknown enum value %q for MemoryFailureRecipient", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *MemoryFailureRecipient) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "hypervisor":
		*e = MemoryFailureRecipientHypervisor
	case "guest":
		*e = MemoryFailureRecipientGuest
	default:
		return fmt.Errorf("unknown enum value %q for MemoryFailureRecipient", s)
	}
	return nil
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
	AnnounceInitial          *uint64                    `json:"announce-initial,omitempty"`
	AnnounceMax              *uint64                    `json:"announce-max,omitempty"`
	AnnounceRounds           *uint64                    `json:"announce-rounds,omitempty"`
	AnnounceStep             *uint64                    `json:"announce-step,omitempty"`
	CompressLevel            *uint8                     `json:"compress-level,omitempty"`
	CompressThreads          *uint8                     `json:"compress-threads,omitempty"`
	CompressWaitThread       *bool                      `json:"compress-wait-thread,omitempty"`
	DecompressThreads        *uint8                     `json:"decompress-threads,omitempty"`
	ThrottleTriggerThreshold *uint8                     `json:"throttle-trigger-threshold,omitempty"`
	CPUThrottleInitial       *uint8                     `json:"cpu-throttle-initial,omitempty"`
	CPUThrottleIncrement     *uint8                     `json:"cpu-throttle-increment,omitempty"`
	CPUThrottleTailslow      *bool                      `json:"cpu-throttle-tailslow,omitempty"`
	TLSCreds                 *StrOrNull                 `json:"tls-creds,omitempty"`
	TLSHostname              *StrOrNull                 `json:"tls-hostname,omitempty"`
	TLSAuthz                 *StrOrNull                 `json:"tls-authz,omitempty"`
	MaxBandwidth             *uint64                    `json:"max-bandwidth,omitempty"`
	DowntimeLimit            *uint64                    `json:"downtime-limit,omitempty"`
	XCheckpointDelay         *uint32                    `json:"x-checkpoint-delay,omitempty"`
	BlockIncremental         *bool                      `json:"block-incremental,omitempty"`
	MultifdChannels          *uint8                     `json:"multifd-channels,omitempty"`
	XbzrleCacheSize          *uint64                    `json:"xbzrle-cache-size,omitempty"`
	MaxPostcopyBandwidth     *uint64                    `json:"max-postcopy-bandwidth,omitempty"`
	MaxCPUThrottle           *uint8                     `json:"max-cpu-throttle,omitempty"`
	MultifdCompression       *MultiFDCompression        `json:"multifd-compression,omitempty"`
	MultifdZlibLevel         *uint8                     `json:"multifd-zlib-level,omitempty"`
	MultifdZstdLevel         *uint8                     `json:"multifd-zstd-level,omitempty"`
	BlockBitmapMapping       []BitmapMigrationNodeAlias `json:"block-bitmap-mapping,omitempty"`
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
	MigrationCapabilityMultifd
	MigrationCapabilityDirtyBitmaps
	MigrationCapabilityPostcopyBlocktime
	MigrationCapabilityLateBlockActivate
	MigrationCapabilityXIgnoreShared
	MigrationCapabilityValidateUUID
	MigrationCapabilityBackgroundSnapshot
	MigrationCapabilityZeroCopySend
	MigrationCapabilityPostcopyPreempt
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
	case MigrationCapabilityMultifd:
		return "multifd"
	case MigrationCapabilityDirtyBitmaps:
		return "dirty-bitmaps"
	case MigrationCapabilityPostcopyBlocktime:
		return "postcopy-blocktime"
	case MigrationCapabilityLateBlockActivate:
		return "late-block-activate"
	case MigrationCapabilityXIgnoreShared:
		return "x-ignore-shared"
	case MigrationCapabilityValidateUUID:
		return "validate-uuid"
	case MigrationCapabilityBackgroundSnapshot:
		return "background-snapshot"
	case MigrationCapabilityZeroCopySend:
		return "zero-copy-send"
	case MigrationCapabilityPostcopyPreempt:
		return "postcopy-preempt"
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
	case MigrationCapabilityMultifd:
		return json.Marshal("multifd")
	case MigrationCapabilityDirtyBitmaps:
		return json.Marshal("dirty-bitmaps")
	case MigrationCapabilityPostcopyBlocktime:
		return json.Marshal("postcopy-blocktime")
	case MigrationCapabilityLateBlockActivate:
		return json.Marshal("late-block-activate")
	case MigrationCapabilityXIgnoreShared:
		return json.Marshal("x-ignore-shared")
	case MigrationCapabilityValidateUUID:
		return json.Marshal("validate-uuid")
	case MigrationCapabilityBackgroundSnapshot:
		return json.Marshal("background-snapshot")
	case MigrationCapabilityZeroCopySend:
		return json.Marshal("zero-copy-send")
	case MigrationCapabilityPostcopyPreempt:
		return json.Marshal("postcopy-preempt")
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
	case "multifd":
		*e = MigrationCapabilityMultifd
	case "dirty-bitmaps":
		*e = MigrationCapabilityDirtyBitmaps
	case "postcopy-blocktime":
		*e = MigrationCapabilityPostcopyBlocktime
	case "late-block-activate":
		*e = MigrationCapabilityLateBlockActivate
	case "x-ignore-shared":
		*e = MigrationCapabilityXIgnoreShared
	case "validate-uuid":
		*e = MigrationCapabilityValidateUUID
	case "background-snapshot":
		*e = MigrationCapabilityBackgroundSnapshot
	case "zero-copy-send":
		*e = MigrationCapabilityZeroCopySend
	case "postcopy-preempt":
		*e = MigrationCapabilityPostcopyPreempt
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
	Vfio                  *VfioStats        `json:"vfio,omitempty"`
	XbzrleCache           *XbzrleCacheStats `json:"xbzrle-cache,omitempty"`
	TotalTime             *int64            `json:"total-time,omitempty"`
	ExpectedDowntime      *int64            `json:"expected-downtime,omitempty"`
	Downtime              *int64            `json:"downtime,omitempty"`
	SetupTime             *int64            `json:"setup-time,omitempty"`
	CPUThrottlePercentage *int64            `json:"cpu-throttle-percentage,omitempty"`
	ErrorDesc             *string           `json:"error-desc,omitempty"`
	BlockedReasons        []string          `json:"blocked-reasons,omitempty"`
	PostcopyBlocktime     *uint32           `json:"postcopy-blocktime,omitempty"`
	PostcopyVcpuBlocktime []uint32          `json:"postcopy-vcpu-blocktime,omitempty"`
	Compression           *CompressionStats `json:"compression,omitempty"`
	SocketAddress         []SocketAddress   `json:"socket-address,omitempty"`
}

// MigrationParameters -> MigrationParameters (struct)

// MigrationParameters implements the "MigrationParameters" QMP API type.
type MigrationParameters struct {
	AnnounceInitial          *uint64                    `json:"announce-initial,omitempty"`
	AnnounceMax              *uint64                    `json:"announce-max,omitempty"`
	AnnounceRounds           *uint64                    `json:"announce-rounds,omitempty"`
	AnnounceStep             *uint64                    `json:"announce-step,omitempty"`
	CompressLevel            *uint8                     `json:"compress-level,omitempty"`
	CompressThreads          *uint8                     `json:"compress-threads,omitempty"`
	CompressWaitThread       *bool                      `json:"compress-wait-thread,omitempty"`
	DecompressThreads        *uint8                     `json:"decompress-threads,omitempty"`
	ThrottleTriggerThreshold *uint8                     `json:"throttle-trigger-threshold,omitempty"`
	CPUThrottleInitial       *uint8                     `json:"cpu-throttle-initial,omitempty"`
	CPUThrottleIncrement     *uint8                     `json:"cpu-throttle-increment,omitempty"`
	CPUThrottleTailslow      *bool                      `json:"cpu-throttle-tailslow,omitempty"`
	TLSCreds                 *string                    `json:"tls-creds,omitempty"`
	TLSHostname              *string                    `json:"tls-hostname,omitempty"`
	TLSAuthz                 *string                    `json:"tls-authz,omitempty"`
	MaxBandwidth             *uint64                    `json:"max-bandwidth,omitempty"`
	DowntimeLimit            *uint64                    `json:"downtime-limit,omitempty"`
	XCheckpointDelay         *uint32                    `json:"x-checkpoint-delay,omitempty"`
	BlockIncremental         *bool                      `json:"block-incremental,omitempty"`
	MultifdChannels          *uint8                     `json:"multifd-channels,omitempty"`
	XbzrleCacheSize          *uint64                    `json:"xbzrle-cache-size,omitempty"`
	MaxPostcopyBandwidth     *uint64                    `json:"max-postcopy-bandwidth,omitempty"`
	MaxCPUThrottle           *uint8                     `json:"max-cpu-throttle,omitempty"`
	MultifdCompression       *MultiFDCompression        `json:"multifd-compression,omitempty"`
	MultifdZlibLevel         *uint8                     `json:"multifd-zlib-level,omitempty"`
	MultifdZstdLevel         *uint8                     `json:"multifd-zstd-level,omitempty"`
	BlockBitmapMapping       []BitmapMigrationNodeAlias `json:"block-bitmap-mapping,omitempty"`
}

// MigrationStats -> MigrationStats (struct)

// MigrationStats implements the "MigrationStats" QMP API type.
type MigrationStats struct {
	Transferred             int64   `json:"transferred"`
	Remaining               int64   `json:"remaining"`
	Total                   int64   `json:"total"`
	Duplicate               int64   `json:"duplicate"`
	Skipped                 int64   `json:"skipped"`
	Normal                  int64   `json:"normal"`
	NormalBytes             int64   `json:"normal-bytes"`
	DirtyPagesRate          int64   `json:"dirty-pages-rate"`
	Mbps                    float64 `json:"mbps"`
	DirtySyncCount          int64   `json:"dirty-sync-count"`
	PostcopyRequests        int64   `json:"postcopy-requests"`
	PageSize                int64   `json:"page-size"`
	MultifdBytes            uint64  `json:"multifd-bytes"`
	PagesPerSecond          uint64  `json:"pages-per-second"`
	PrecopyBytes            uint64  `json:"precopy-bytes"`
	DowntimeBytes           uint64  `json:"downtime-bytes"`
	PostcopyBytes           uint64  `json:"postcopy-bytes"`
	DirtySyncMissedZeroCopy uint64  `json:"dirty-sync-missed-zero-copy"`
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
	MigrationStatusPostcopyPaused
	MigrationStatusPostcopyRecover
	MigrationStatusCompleted
	MigrationStatusFailed
	MigrationStatusColo
	MigrationStatusPreSwitchover
	MigrationStatusDevice
	MigrationStatusWaitUnplug
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
	case MigrationStatusPostcopyPaused:
		return "postcopy-paused"
	case MigrationStatusPostcopyRecover:
		return "postcopy-recover"
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
	case MigrationStatusWaitUnplug:
		return "wait-unplug"
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
	case MigrationStatusPostcopyPaused:
		return json.Marshal("postcopy-paused")
	case MigrationStatusPostcopyRecover:
		return json.Marshal("postcopy-recover")
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
	case MigrationStatusWaitUnplug:
		return json.Marshal("wait-unplug")
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
	case "postcopy-paused":
		*e = MigrationStatusPostcopyPaused
	case "postcopy-recover":
		*e = MigrationStatusPostcopyRecover
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
	case "wait-unplug":
		*e = MigrationStatusWaitUnplug
	default:
		return fmt.Errorf("unknown enum value %q for MigrationStatus", s)
	}
	return nil
}

// MirrorCopyMode -> MirrorCopyMode (enum)

// MirrorCopyMode implements the "MirrorCopyMode" QMP API type.
type MirrorCopyMode int

// Known values of MirrorCopyMode.
const (
	MirrorCopyModeBackground MirrorCopyMode = iota
	MirrorCopyModeWriteBlocking
)

// String implements fmt.Stringer.
func (e MirrorCopyMode) String() string {
	switch e {
	case MirrorCopyModeBackground:
		return "background"
	case MirrorCopyModeWriteBlocking:
		return "write-blocking"
	default:
		return fmt.Sprintf("MirrorCopyMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e MirrorCopyMode) MarshalJSON() ([]byte, error) {
	switch e {
	case MirrorCopyModeBackground:
		return json.Marshal("background")
	case MirrorCopyModeWriteBlocking:
		return json.Marshal("write-blocking")
	default:
		return nil, fmt.Errorf("unknown enum value %q for MirrorCopyMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *MirrorCopyMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "background":
		*e = MirrorCopyModeBackground
	case "write-blocking":
		*e = MirrorCopyModeWriteBlocking
	default:
		return fmt.Errorf("unknown enum value %q for MirrorCopyMode", s)
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
	MirrorSyncModeBitmap
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
	case MirrorSyncModeBitmap:
		return "bitmap"
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
	case MirrorSyncModeBitmap:
		return json.Marshal("bitmap")
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
	case "bitmap":
		*e = MirrorSyncModeBitmap
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

// MultiFDCompression -> MultiFDCompression (enum)

// MultiFDCompression implements the "MultiFDCompression" QMP API type.
type MultiFDCompression int

// Known values of MultiFDCompression.
const (
	MultiFDCompressionNone MultiFDCompression = iota
	MultiFDCompressionZlib
	MultiFDCompressionZstd
)

// String implements fmt.Stringer.
func (e MultiFDCompression) String() string {
	switch e {
	case MultiFDCompressionNone:
		return "none"
	case MultiFDCompressionZlib:
		return "zlib"
	case MultiFDCompressionZstd:
		return "zstd"
	default:
		return fmt.Sprintf("MultiFDCompression(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e MultiFDCompression) MarshalJSON() ([]byte, error) {
	switch e {
	case MultiFDCompressionNone:
		return json.Marshal("none")
	case MultiFDCompressionZlib:
		return json.Marshal("zlib")
	case MultiFDCompressionZstd:
		return json.Marshal("zstd")
	default:
		return nil, fmt.Errorf("unknown enum value %q for MultiFDCompression", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *MultiFDCompression) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = MultiFDCompressionNone
	case "zlib":
		*e = MultiFDCompressionZlib
	case "zstd":
		*e = MultiFDCompressionZstd
	default:
		return fmt.Errorf("unknown enum value %q for MultiFDCompression", s)
	}
	return nil
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

// NbdServerAddOptions -> NBDServerAddOptions (struct)

// NBDServerAddOptions implements the "NbdServerAddOptions" QMP API type.
type NBDServerAddOptions struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Device      string  `json:"device"`
	Writable    *bool   `json:"writable,omitempty"`
	Bitmap      *string `json:"bitmap,omitempty"`
}

// NetClientDriver -> NetClientDriver (enum)

// NetClientDriver implements the "NetClientDriver" QMP API type.
type NetClientDriver int

// Known values of NetClientDriver.
const (
	NetClientDriverNone NetClientDriver = iota
	NetClientDriverNic
	NetClientDriverUser
	NetClientDriverTap
	NetClientDriverL2Tpv3
	NetClientDriverSocket
	NetClientDriverVde
	NetClientDriverBridge
	NetClientDriverHubport
	NetClientDriverNetmap
	NetClientDriverVhostUser
	NetClientDriverVhostVdpa
	NetClientDriverVmnetHost
	NetClientDriverVmnetShared
	NetClientDriverVmnetBridged
)

// String implements fmt.Stringer.
func (e NetClientDriver) String() string {
	switch e {
	case NetClientDriverNone:
		return "none"
	case NetClientDriverNic:
		return "nic"
	case NetClientDriverUser:
		return "user"
	case NetClientDriverTap:
		return "tap"
	case NetClientDriverL2Tpv3:
		return "l2tpv3"
	case NetClientDriverSocket:
		return "socket"
	case NetClientDriverVde:
		return "vde"
	case NetClientDriverBridge:
		return "bridge"
	case NetClientDriverHubport:
		return "hubport"
	case NetClientDriverNetmap:
		return "netmap"
	case NetClientDriverVhostUser:
		return "vhost-user"
	case NetClientDriverVhostVdpa:
		return "vhost-vdpa"
	case NetClientDriverVmnetHost:
		return "vmnet-host"
	case NetClientDriverVmnetShared:
		return "vmnet-shared"
	case NetClientDriverVmnetBridged:
		return "vmnet-bridged"
	default:
		return fmt.Sprintf("NetClientDriver(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e NetClientDriver) MarshalJSON() ([]byte, error) {
	switch e {
	case NetClientDriverNone:
		return json.Marshal("none")
	case NetClientDriverNic:
		return json.Marshal("nic")
	case NetClientDriverUser:
		return json.Marshal("user")
	case NetClientDriverTap:
		return json.Marshal("tap")
	case NetClientDriverL2Tpv3:
		return json.Marshal("l2tpv3")
	case NetClientDriverSocket:
		return json.Marshal("socket")
	case NetClientDriverVde:
		return json.Marshal("vde")
	case NetClientDriverBridge:
		return json.Marshal("bridge")
	case NetClientDriverHubport:
		return json.Marshal("hubport")
	case NetClientDriverNetmap:
		return json.Marshal("netmap")
	case NetClientDriverVhostUser:
		return json.Marshal("vhost-user")
	case NetClientDriverVhostVdpa:
		return json.Marshal("vhost-vdpa")
	case NetClientDriverVmnetHost:
		return json.Marshal("vmnet-host")
	case NetClientDriverVmnetShared:
		return json.Marshal("vmnet-shared")
	case NetClientDriverVmnetBridged:
		return json.Marshal("vmnet-bridged")
	default:
		return nil, fmt.Errorf("unknown enum value %q for NetClientDriver", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NetClientDriver) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = NetClientDriverNone
	case "nic":
		*e = NetClientDriverNic
	case "user":
		*e = NetClientDriverUser
	case "tap":
		*e = NetClientDriverTap
	case "l2tpv3":
		*e = NetClientDriverL2Tpv3
	case "socket":
		*e = NetClientDriverSocket
	case "vde":
		*e = NetClientDriverVde
	case "bridge":
		*e = NetClientDriverBridge
	case "hubport":
		*e = NetClientDriverHubport
	case "netmap":
		*e = NetClientDriverNetmap
	case "vhost-user":
		*e = NetClientDriverVhostUser
	case "vhost-vdpa":
		*e = NetClientDriverVhostVdpa
	case "vmnet-host":
		*e = NetClientDriverVmnetHost
	case "vmnet-shared":
		*e = NetClientDriverVmnetShared
	case "vmnet-bridged":
		*e = NetClientDriverVmnetBridged
	default:
		return fmt.Errorf("unknown enum value %q for NetClientDriver", s)
	}
	return nil
}

// NetFilterDirection -> NetFilterDirection (enum)

// NetFilterDirection implements the "NetFilterDirection" QMP API type.
type NetFilterDirection int

// Known values of NetFilterDirection.
const (
	NetFilterDirectionAll NetFilterDirection = iota
	NetFilterDirectionRx
	NetFilterDirectionTx
)

// String implements fmt.Stringer.
func (e NetFilterDirection) String() string {
	switch e {
	case NetFilterDirectionAll:
		return "all"
	case NetFilterDirectionRx:
		return "rx"
	case NetFilterDirectionTx:
		return "tx"
	default:
		return fmt.Sprintf("NetFilterDirection(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e NetFilterDirection) MarshalJSON() ([]byte, error) {
	switch e {
	case NetFilterDirectionAll:
		return json.Marshal("all")
	case NetFilterDirectionRx:
		return json.Marshal("rx")
	case NetFilterDirectionTx:
		return json.Marshal("tx")
	default:
		return nil, fmt.Errorf("unknown enum value %q for NetFilterDirection", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NetFilterDirection) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "all":
		*e = NetFilterDirectionAll
	case "rx":
		*e = NetFilterDirectionRx
	case "tx":
		*e = NetFilterDirectionTx
	default:
		return fmt.Errorf("unknown enum value %q for NetFilterDirection", s)
	}
	return nil
}

// Netdev -> Netdev (flat union)

// Netdev implements the "Netdev" QMP API type.
//
// Can be one of:
//   - NetdevVariantBridge
//   - NetdevVariantHubport
//   - NetdevVariantL2Tpv3
//   - NetdevVariantNetmap
//   - NetdevVariantNic
//   - NetdevVariantSocket
//   - NetdevVariantTap
//   - NetdevVariantUser
//   - NetdevVariantVde
//   - NetdevVariantVhostUser
//   - NetdevVariantVhostVdpa
//   - NetdevVariantVmnetBridged
//   - NetdevVariantVmnetHost
//   - NetdevVariantVmnetShared
type Netdev interface {
	isNetdev()
}

// NetdevVariantBridge is an implementation of Netdev.
type NetdevVariantBridge struct {
	ID     string  `json:"id"`
	Br     *string `json:"br,omitempty"`
	Helper *string `json:"helper,omitempty"`
}

func (NetdevVariantBridge) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantBridge) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantBridge
	}{
		NetClientDriverBridge,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantHubport is an implementation of Netdev.
type NetdevVariantHubport struct {
	ID     string  `json:"id"`
	Hubid  int32   `json:"hubid"`
	Netdev *string `json:"netdev,omitempty"`
}

func (NetdevVariantHubport) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantHubport) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantHubport
	}{
		NetClientDriverHubport,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantL2Tpv3 is an implementation of Netdev.
type NetdevVariantL2Tpv3 struct {
	ID         string  `json:"id"`
	Src        string  `json:"src"`
	Dst        string  `json:"dst"`
	Srcport    *string `json:"srcport,omitempty"`
	Dstport    *string `json:"dstport,omitempty"`
	Ipv6       *bool   `json:"ipv6,omitempty"`
	UDP        *bool   `json:"udp,omitempty"`
	Cookie64   *bool   `json:"cookie64,omitempty"`
	Counter    *bool   `json:"counter,omitempty"`
	Pincounter *bool   `json:"pincounter,omitempty"`
	Txcookie   *uint64 `json:"txcookie,omitempty"`
	Rxcookie   *uint64 `json:"rxcookie,omitempty"`
	Txsession  uint32  `json:"txsession"`
	Rxsession  *uint32 `json:"rxsession,omitempty"`
	Offset     *uint32 `json:"offset,omitempty"`
}

func (NetdevVariantL2Tpv3) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantL2Tpv3) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantL2Tpv3
	}{
		NetClientDriverL2Tpv3,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantNetmap is an implementation of Netdev.
type NetdevVariantNetmap struct {
	ID      string  `json:"id"`
	Ifname  string  `json:"ifname"`
	Devname *string `json:"devname,omitempty"`
}

func (NetdevVariantNetmap) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantNetmap) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantNetmap
	}{
		NetClientDriverNetmap,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantNic is an implementation of Netdev.
type NetdevVariantNic struct {
	ID      string  `json:"id"`
	Netdev  *string `json:"netdev,omitempty"`
	Macaddr *string `json:"macaddr,omitempty"`
	Model   *string `json:"model,omitempty"`
	Addr    *string `json:"addr,omitempty"`
	Vectors *uint32 `json:"vectors,omitempty"`
}

func (NetdevVariantNic) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantNic) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantNic
	}{
		NetClientDriverNic,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantSocket is an implementation of Netdev.
type NetdevVariantSocket struct {
	ID        string  `json:"id"`
	FD        *string `json:"fd,omitempty"`
	Listen    *string `json:"listen,omitempty"`
	Connect   *string `json:"connect,omitempty"`
	Mcast     *string `json:"mcast,omitempty"`
	Localaddr *string `json:"localaddr,omitempty"`
	UDP       *string `json:"udp,omitempty"`
}

func (NetdevVariantSocket) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantSocket) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantSocket
	}{
		NetClientDriverSocket,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantTap is an implementation of Netdev.
type NetdevVariantTap struct {
	ID         string  `json:"id"`
	Ifname     *string `json:"ifname,omitempty"`
	FD         *string `json:"fd,omitempty"`
	Fds        *string `json:"fds,omitempty"`
	Script     *string `json:"script,omitempty"`
	Downscript *string `json:"downscript,omitempty"`
	Br         *string `json:"br,omitempty"`
	Helper     *string `json:"helper,omitempty"`
	Sndbuf     *uint64 `json:"sndbuf,omitempty"`
	VnetHdr    *bool   `json:"vnet_hdr,omitempty"`
	Vhost      *bool   `json:"vhost,omitempty"`
	Vhostfd    *string `json:"vhostfd,omitempty"`
	Vhostfds   *string `json:"vhostfds,omitempty"`
	Vhostforce *bool   `json:"vhostforce,omitempty"`
	Queues     *uint32 `json:"queues,omitempty"`
	PollUs     *uint32 `json:"poll-us,omitempty"`
}

func (NetdevVariantTap) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantTap) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantTap
	}{
		NetClientDriverTap,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantUser is an implementation of Netdev.
type NetdevVariantUser struct {
	ID             string   `json:"id"`
	Hostname       *string  `json:"hostname,omitempty"`
	Restrict       *bool    `json:"restrict,omitempty"`
	Ipv4           *bool    `json:"ipv4,omitempty"`
	Ipv6           *bool    `json:"ipv6,omitempty"`
	IP             *string  `json:"ip,omitempty"`
	Net            *string  `json:"net,omitempty"`
	Host           *string  `json:"host,omitempty"`
	Tftp           *string  `json:"tftp,omitempty"`
	Bootfile       *string  `json:"bootfile,omitempty"`
	Dhcpstart      *string  `json:"dhcpstart,omitempty"`
	Dns            *string  `json:"dns,omitempty"`
	Dnssearch      []String `json:"dnssearch,omitempty"`
	Domainname     *string  `json:"domainname,omitempty"`
	Ipv6Prefix     *string  `json:"ipv6-prefix,omitempty"`
	Ipv6Prefixlen  *int64   `json:"ipv6-prefixlen,omitempty"`
	Ipv6Host       *string  `json:"ipv6-host,omitempty"`
	Ipv6Dns        *string  `json:"ipv6-dns,omitempty"`
	Smb            *string  `json:"smb,omitempty"`
	Smbserver      *string  `json:"smbserver,omitempty"`
	Hostfwd        []String `json:"hostfwd,omitempty"`
	Guestfwd       []String `json:"guestfwd,omitempty"`
	TftpServerName *string  `json:"tftp-server-name,omitempty"`
}

func (NetdevVariantUser) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantUser) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantUser
	}{
		NetClientDriverUser,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantVde is an implementation of Netdev.
type NetdevVariantVde struct {
	ID    string  `json:"id"`
	Sock  *string `json:"sock,omitempty"`
	Port  *uint16 `json:"port,omitempty"`
	Group *string `json:"group,omitempty"`
	Mode  *uint16 `json:"mode,omitempty"`
}

func (NetdevVariantVde) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantVde) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantVde
	}{
		NetClientDriverVde,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantVhostUser is an implementation of Netdev.
type NetdevVariantVhostUser struct {
	ID         string `json:"id"`
	Chardev    string `json:"chardev"`
	Vhostforce *bool  `json:"vhostforce,omitempty"`
	Queues     *int64 `json:"queues,omitempty"`
}

func (NetdevVariantVhostUser) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantVhostUser) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantVhostUser
	}{
		NetClientDriverVhostUser,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantVhostVdpa is an implementation of Netdev.
type NetdevVariantVhostVdpa struct {
	ID       string  `json:"id"`
	Vhostdev *string `json:"vhostdev,omitempty"`
	Queues   *int64  `json:"queues,omitempty"`
	XSvq     *bool   `json:"x-svq,omitempty"`
}

func (NetdevVariantVhostVdpa) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantVhostVdpa) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantVhostVdpa
	}{
		NetClientDriverVhostVdpa,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantVmnetBridged is an implementation of Netdev.
type NetdevVariantVmnetBridged struct {
	ID       string `json:"id"`
	Ifname   string `json:"ifname"`
	Isolated *bool  `json:"isolated,omitempty"`
}

func (NetdevVariantVmnetBridged) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantVmnetBridged) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantVmnetBridged
	}{
		NetClientDriverVmnetBridged,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantVmnetHost is an implementation of Netdev.
type NetdevVariantVmnetHost struct {
	ID           string  `json:"id"`
	StartAddress *string `json:"start-address,omitempty"`
	EndAddress   *string `json:"end-address,omitempty"`
	SubnetMask   *string `json:"subnet-mask,omitempty"`
	Isolated     *bool   `json:"isolated,omitempty"`
	NetUUID      *string `json:"net-uuid,omitempty"`
}

func (NetdevVariantVmnetHost) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantVmnetHost) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantVmnetHost
	}{
		NetClientDriverVmnetHost,
		s,
	}
	return json.Marshal(v)
}

// NetdevVariantVmnetShared is an implementation of Netdev.
type NetdevVariantVmnetShared struct {
	ID           string  `json:"id"`
	StartAddress *string `json:"start-address,omitempty"`
	EndAddress   *string `json:"end-address,omitempty"`
	SubnetMask   *string `json:"subnet-mask,omitempty"`
	Isolated     *bool   `json:"isolated,omitempty"`
	Nat66Prefix  *string `json:"nat66-prefix,omitempty"`
}

func (NetdevVariantVmnetShared) isNetdev() {}

// MarshalJSON implements json.Marshaler.
func (s NetdevVariantVmnetShared) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
		NetdevVariantVmnetShared
	}{
		NetClientDriverVmnetShared,
		s,
	}
	return json.Marshal(v)
}

func decodeNetdev(bs json.RawMessage) (Netdev, error) {
	v := struct {
		Type NetClientDriver `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case NetClientDriverBridge:
		var ret NetdevVariantBridge
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverHubport:
		var ret NetdevVariantHubport
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverL2Tpv3:
		var ret NetdevVariantL2Tpv3
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverNetmap:
		var ret NetdevVariantNetmap
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverNic:
		var ret NetdevVariantNic
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverSocket:
		var ret NetdevVariantSocket
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverTap:
		var ret NetdevVariantTap
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverUser:
		var ret NetdevVariantUser
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverVde:
		var ret NetdevVariantVde
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverVhostUser:
		var ret NetdevVariantVhostUser
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverVhostVdpa:
		var ret NetdevVariantVhostVdpa
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverVmnetBridged:
		var ret NetdevVariantVmnetBridged
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverVmnetHost:
		var ret NetdevVariantVmnetHost
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NetClientDriverVmnetShared:
		var ret NetdevVariantVmnetShared
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union Netdev", v.Type)
	}
}

// NetfilterInsert -> NetfilterInsert (enum)

// NetfilterInsert implements the "NetfilterInsert" QMP API type.
type NetfilterInsert int

// Known values of NetfilterInsert.
const (
	NetfilterInsertBefore NetfilterInsert = iota
	NetfilterInsertBehind
)

// String implements fmt.Stringer.
func (e NetfilterInsert) String() string {
	switch e {
	case NetfilterInsertBefore:
		return "before"
	case NetfilterInsertBehind:
		return "behind"
	default:
		return fmt.Sprintf("NetfilterInsert(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e NetfilterInsert) MarshalJSON() ([]byte, error) {
	switch e {
	case NetfilterInsertBefore:
		return json.Marshal("before")
	case NetfilterInsertBehind:
		return json.Marshal("behind")
	default:
		return nil, fmt.Errorf("unknown enum value %q for NetfilterInsert", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NetfilterInsert) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "before":
		*e = NetfilterInsertBefore
	case "behind":
		*e = NetfilterInsertBehind
	default:
		return fmt.Errorf("unknown enum value %q for NetfilterInsert", s)
	}
	return nil
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

// NumaOptions -> NumaOptions (flat union)

// NumaOptions implements the "NumaOptions" QMP API type.
//
// Can be one of:
//   - NumaOptionsVariantCPU
//   - NumaOptionsVariantDist
//   - NumaOptionsVariantHmatCache
//   - NumaOptionsVariantHmatLb
//   - NumaOptionsVariantNode
type NumaOptions interface {
	isNumaOptions()
}

// NumaOptionsVariantCPU is an implementation of NumaOptions.
type NumaOptionsVariantCPU struct {
}

func (NumaOptionsVariantCPU) isNumaOptions() {}

// MarshalJSON implements json.Marshaler.
func (s NumaOptionsVariantCPU) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NumaOptionsType `json:"type"`
		NumaOptionsVariantCPU
	}{
		NumaOptionsTypeCPU,
		s,
	}
	return json.Marshal(v)
}

// NumaOptionsVariantDist is an implementation of NumaOptions.
type NumaOptionsVariantDist struct {
	Src uint16 `json:"src"`
	Dst uint16 `json:"dst"`
	Val uint8  `json:"val"`
}

func (NumaOptionsVariantDist) isNumaOptions() {}

// MarshalJSON implements json.Marshaler.
func (s NumaOptionsVariantDist) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NumaOptionsType `json:"type"`
		NumaOptionsVariantDist
	}{
		NumaOptionsTypeDist,
		s,
	}
	return json.Marshal(v)
}

// NumaOptionsVariantHmatCache is an implementation of NumaOptions.
type NumaOptionsVariantHmatCache struct {
	NodeID        uint32                 `json:"node-id"`
	Size          uint64                 `json:"size"`
	Level         uint8                  `json:"level"`
	Associativity HmatCacheAssociativity `json:"associativity"`
	Policy        HmatCacheWritePolicy   `json:"policy"`
	Line          uint16                 `json:"line"`
}

func (NumaOptionsVariantHmatCache) isNumaOptions() {}

// MarshalJSON implements json.Marshaler.
func (s NumaOptionsVariantHmatCache) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NumaOptionsType `json:"type"`
		NumaOptionsVariantHmatCache
	}{
		NumaOptionsTypeHmatCache,
		s,
	}
	return json.Marshal(v)
}

// NumaOptionsVariantHmatLb is an implementation of NumaOptions.
type NumaOptionsVariantHmatLb struct {
	Initiator uint16                `json:"initiator"`
	Target    uint16                `json:"target"`
	Hierarchy HmatLbMemoryHierarchy `json:"hierarchy"`
	DataType  HmatLbDataType        `json:"data-type"`
	Latency   *uint64               `json:"latency,omitempty"`
	Bandwidth *uint64               `json:"bandwidth,omitempty"`
}

func (NumaOptionsVariantHmatLb) isNumaOptions() {}

// MarshalJSON implements json.Marshaler.
func (s NumaOptionsVariantHmatLb) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NumaOptionsType `json:"type"`
		NumaOptionsVariantHmatLb
	}{
		NumaOptionsTypeHmatLb,
		s,
	}
	return json.Marshal(v)
}

// NumaOptionsVariantNode is an implementation of NumaOptions.
type NumaOptionsVariantNode struct {
	Nodeid    *uint16  `json:"nodeid,omitempty"`
	Cpus      []uint16 `json:"cpus,omitempty"`
	Mem       *uint64  `json:"mem,omitempty"`
	Memdev    *string  `json:"memdev,omitempty"`
	Initiator *uint16  `json:"initiator,omitempty"`
}

func (NumaOptionsVariantNode) isNumaOptions() {}

// MarshalJSON implements json.Marshaler.
func (s NumaOptionsVariantNode) MarshalJSON() ([]byte, error) {
	v := struct {
		Type NumaOptionsType `json:"type"`
		NumaOptionsVariantNode
	}{
		NumaOptionsTypeNode,
		s,
	}
	return json.Marshal(v)
}

func decodeNumaOptions(bs json.RawMessage) (NumaOptions, error) {
	v := struct {
		Type NumaOptionsType `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case NumaOptionsTypeCPU:
		var ret NumaOptionsVariantCPU
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NumaOptionsTypeDist:
		var ret NumaOptionsVariantDist
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NumaOptionsTypeHmatCache:
		var ret NumaOptionsVariantHmatCache
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NumaOptionsTypeHmatLb:
		var ret NumaOptionsVariantHmatLb
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case NumaOptionsTypeNode:
		var ret NumaOptionsVariantNode
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union NumaOptions", v.Type)
	}
}

// NumaOptionsType -> NumaOptionsType (enum)

// NumaOptionsType implements the "NumaOptionsType" QMP API type.
type NumaOptionsType int

// Known values of NumaOptionsType.
const (
	NumaOptionsTypeNode NumaOptionsType = iota
	NumaOptionsTypeDist
	NumaOptionsTypeCPU
	NumaOptionsTypeHmatLb
	NumaOptionsTypeHmatCache
)

// String implements fmt.Stringer.
func (e NumaOptionsType) String() string {
	switch e {
	case NumaOptionsTypeNode:
		return "node"
	case NumaOptionsTypeDist:
		return "dist"
	case NumaOptionsTypeCPU:
		return "cpu"
	case NumaOptionsTypeHmatLb:
		return "hmat-lb"
	case NumaOptionsTypeHmatCache:
		return "hmat-cache"
	default:
		return fmt.Sprintf("NumaOptionsType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e NumaOptionsType) MarshalJSON() ([]byte, error) {
	switch e {
	case NumaOptionsTypeNode:
		return json.Marshal("node")
	case NumaOptionsTypeDist:
		return json.Marshal("dist")
	case NumaOptionsTypeCPU:
		return json.Marshal("cpu")
	case NumaOptionsTypeHmatLb:
		return json.Marshal("hmat-lb")
	case NumaOptionsTypeHmatCache:
		return json.Marshal("hmat-cache")
	default:
		return nil, fmt.Errorf("unknown enum value %q for NumaOptionsType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NumaOptionsType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "node":
		*e = NumaOptionsTypeNode
	case "dist":
		*e = NumaOptionsTypeDist
	case "cpu":
		*e = NumaOptionsTypeCPU
	case "hmat-lb":
		*e = NumaOptionsTypeHmatLb
	case "hmat-cache":
		*e = NumaOptionsTypeHmatCache
	default:
		return fmt.Errorf("unknown enum value %q for NumaOptionsType", s)
	}
	return nil
}

// ObjectOptions -> ObjectOptions (flat union)

// ObjectOptions implements the "ObjectOptions" QMP API type.
//
// Can be one of:
//   - ObjectOptionsVariantAuthzList
//   - ObjectOptionsVariantAuthzListfile
//   - ObjectOptionsVariantAuthzPam
//   - ObjectOptionsVariantAuthzSimple
//   - ObjectOptionsVariantCanHostSocketcan
//   - ObjectOptionsVariantColoCompare
//   - ObjectOptionsVariantCryptodevBackend
//   - ObjectOptionsVariantCryptodevBackendBuiltin
//   - ObjectOptionsVariantCryptodevVhostUser
//   - ObjectOptionsVariantDbusVmstate
//   - ObjectOptionsVariantFilterBuffer
//   - ObjectOptionsVariantFilterDump
//   - ObjectOptionsVariantFilterMirror
//   - ObjectOptionsVariantFilterRedirector
//   - ObjectOptionsVariantFilterReplay
//   - ObjectOptionsVariantFilterRewriter
//   - ObjectOptionsVariantInputBarrier
//   - ObjectOptionsVariantInputLinux
//   - ObjectOptionsVariantIothread
//   - ObjectOptionsVariantMainLoop
//   - ObjectOptionsVariantMemoryBackendEpc
//   - ObjectOptionsVariantMemoryBackendFile
//   - ObjectOptionsVariantMemoryBackendMemfd
//   - ObjectOptionsVariantMemoryBackendRAM
//   - ObjectOptionsVariantPrManagerHelper
//   - ObjectOptionsVariantQtest
//   - ObjectOptionsVariantRngBuiltin
//   - ObjectOptionsVariantRngEgd
//   - ObjectOptionsVariantRngRandom
//   - ObjectOptionsVariantSecret
//   - ObjectOptionsVariantSecretKeyring
//   - ObjectOptionsVariantSevGuest
//   - ObjectOptionsVariantThrottleGroup
//   - ObjectOptionsVariantTLSCipherSuites
//   - ObjectOptionsVariantTLSCredsAnon
//   - ObjectOptionsVariantTLSCredsPsk
//   - ObjectOptionsVariantTLSCredsX509
//   - ObjectOptionsVariantXRemoteObject
//   - ObjectOptionsVariantXVfioUserServer
type ObjectOptions interface {
	isObjectOptions()
}

// ObjectOptionsVariantAuthzList is an implementation of ObjectOptions.
type ObjectOptionsVariantAuthzList struct {
	ID     string            `json:"id"`
	Policy *QAuthZListPolicy `json:"policy,omitempty"`
	Rules  []QAuthZListRule  `json:"rules,omitempty"`
}

func (ObjectOptionsVariantAuthzList) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantAuthzList) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantAuthzList
	}{
		ObjectTypeAuthzList,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantAuthzListfile is an implementation of ObjectOptions.
type ObjectOptionsVariantAuthzListfile struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Refresh  *bool  `json:"refresh,omitempty"`
}

func (ObjectOptionsVariantAuthzListfile) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantAuthzListfile) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantAuthzListfile
	}{
		ObjectTypeAuthzListfile,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantAuthzPam is an implementation of ObjectOptions.
type ObjectOptionsVariantAuthzPam struct {
	ID      string `json:"id"`
	Service string `json:"service"`
}

func (ObjectOptionsVariantAuthzPam) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantAuthzPam) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantAuthzPam
	}{
		ObjectTypeAuthzPam,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantAuthzSimple is an implementation of ObjectOptions.
type ObjectOptionsVariantAuthzSimple struct {
	ID       string `json:"id"`
	Identity string `json:"identity"`
}

func (ObjectOptionsVariantAuthzSimple) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantAuthzSimple) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantAuthzSimple
	}{
		ObjectTypeAuthzSimple,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantCanHostSocketcan is an implementation of ObjectOptions.
type ObjectOptionsVariantCanHostSocketcan struct {
	ID     string `json:"id"`
	If     string `json:"if"`
	Canbus string `json:"canbus"`
}

func (ObjectOptionsVariantCanHostSocketcan) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantCanHostSocketcan) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantCanHostSocketcan
	}{
		ObjectTypeCanHostSocketcan,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantColoCompare is an implementation of ObjectOptions.
type ObjectOptionsVariantColoCompare struct {
	ID               string  `json:"id"`
	PrimaryIn        string  `json:"primary_in"`
	SecondaryIn      string  `json:"secondary_in"`
	Outdev           string  `json:"outdev"`
	Iothread         string  `json:"iothread"`
	NotifyDev        *string `json:"notify_dev,omitempty"`
	CompareTimeout   *uint64 `json:"compare_timeout,omitempty"`
	ExpiredScanCycle *uint32 `json:"expired_scan_cycle,omitempty"`
	MaxQueueSize     *uint32 `json:"max_queue_size,omitempty"`
	VnetHdrSupport   *bool   `json:"vnet_hdr_support,omitempty"`
}

func (ObjectOptionsVariantColoCompare) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantColoCompare) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantColoCompare
	}{
		ObjectTypeColoCompare,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantCryptodevBackend is an implementation of ObjectOptions.
type ObjectOptionsVariantCryptodevBackend struct {
	ID     string  `json:"id"`
	Queues *uint32 `json:"queues,omitempty"`
}

func (ObjectOptionsVariantCryptodevBackend) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantCryptodevBackend) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantCryptodevBackend
	}{
		ObjectTypeCryptodevBackend,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantCryptodevBackendBuiltin is an implementation of ObjectOptions.
type ObjectOptionsVariantCryptodevBackendBuiltin struct {
	ID     string  `json:"id"`
	Queues *uint32 `json:"queues,omitempty"`
}

func (ObjectOptionsVariantCryptodevBackendBuiltin) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantCryptodevBackendBuiltin) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantCryptodevBackendBuiltin
	}{
		ObjectTypeCryptodevBackendBuiltin,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantCryptodevVhostUser is an implementation of ObjectOptions.
type ObjectOptionsVariantCryptodevVhostUser struct {
	ID      string `json:"id"`
	Chardev string `json:"chardev"`
}

func (ObjectOptionsVariantCryptodevVhostUser) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantCryptodevVhostUser) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantCryptodevVhostUser
	}{
		ObjectTypeCryptodevVhostUser,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantDbusVmstate is an implementation of ObjectOptions.
type ObjectOptionsVariantDbusVmstate struct {
	ID     string  `json:"id"`
	Addr   string  `json:"addr"`
	IDList *string `json:"id-list,omitempty"`
}

func (ObjectOptionsVariantDbusVmstate) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantDbusVmstate) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantDbusVmstate
	}{
		ObjectTypeDbusVmstate,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantFilterBuffer is an implementation of ObjectOptions.
type ObjectOptionsVariantFilterBuffer struct {
	ID       string `json:"id"`
	Interval uint32 `json:"interval"`
}

func (ObjectOptionsVariantFilterBuffer) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantFilterBuffer) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantFilterBuffer
	}{
		ObjectTypeFilterBuffer,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantFilterDump is an implementation of ObjectOptions.
type ObjectOptionsVariantFilterDump struct {
	ID     string  `json:"id"`
	File   string  `json:"file"`
	Maxlen *uint32 `json:"maxlen,omitempty"`
}

func (ObjectOptionsVariantFilterDump) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantFilterDump) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantFilterDump
	}{
		ObjectTypeFilterDump,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantFilterMirror is an implementation of ObjectOptions.
type ObjectOptionsVariantFilterMirror struct {
	ID             string `json:"id"`
	Outdev         string `json:"outdev"`
	VnetHdrSupport *bool  `json:"vnet_hdr_support,omitempty"`
}

func (ObjectOptionsVariantFilterMirror) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantFilterMirror) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantFilterMirror
	}{
		ObjectTypeFilterMirror,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantFilterRedirector is an implementation of ObjectOptions.
type ObjectOptionsVariantFilterRedirector struct {
	ID             string  `json:"id"`
	Indev          *string `json:"indev,omitempty"`
	Outdev         *string `json:"outdev,omitempty"`
	VnetHdrSupport *bool   `json:"vnet_hdr_support,omitempty"`
}

func (ObjectOptionsVariantFilterRedirector) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantFilterRedirector) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantFilterRedirector
	}{
		ObjectTypeFilterRedirector,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantFilterReplay is an implementation of ObjectOptions.
type ObjectOptionsVariantFilterReplay struct {
	ID       string              `json:"id"`
	Netdev   string              `json:"netdev"`
	Queue    *NetFilterDirection `json:"queue,omitempty"`
	Status   *string             `json:"status,omitempty"`
	Position *string             `json:"position,omitempty"`
	Insert   *NetfilterInsert    `json:"insert,omitempty"`
}

func (ObjectOptionsVariantFilterReplay) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantFilterReplay) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantFilterReplay
	}{
		ObjectTypeFilterReplay,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantFilterRewriter is an implementation of ObjectOptions.
type ObjectOptionsVariantFilterRewriter struct {
	ID             string `json:"id"`
	VnetHdrSupport *bool  `json:"vnet_hdr_support,omitempty"`
}

func (ObjectOptionsVariantFilterRewriter) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantFilterRewriter) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantFilterRewriter
	}{
		ObjectTypeFilterRewriter,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantInputBarrier is an implementation of ObjectOptions.
type ObjectOptionsVariantInputBarrier struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Server  *string `json:"server,omitempty"`
	Port    *string `json:"port,omitempty"`
	XOrigin *string `json:"x-origin,omitempty"`
	YOrigin *string `json:"y-origin,omitempty"`
	Width   *string `json:"width,omitempty"`
	Height  *string `json:"height,omitempty"`
}

func (ObjectOptionsVariantInputBarrier) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantInputBarrier) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantInputBarrier
	}{
		ObjectTypeInputBarrier,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantInputLinux is an implementation of ObjectOptions.
type ObjectOptionsVariantInputLinux struct {
	ID         string          `json:"id"`
	Evdev      string          `json:"evdev"`
	GrabAll    *bool           `json:"grab_all,omitempty"`
	Repeat     *bool           `json:"repeat,omitempty"`
	GrabToggle *GrabToggleKeys `json:"grab-toggle,omitempty"`
}

func (ObjectOptionsVariantInputLinux) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantInputLinux) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantInputLinux
	}{
		ObjectTypeInputLinux,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantIothread is an implementation of ObjectOptions.
type ObjectOptionsVariantIothread struct {
	ID         string `json:"id"`
	PollMaxNs  *int64 `json:"poll-max-ns,omitempty"`
	PollGrow   *int64 `json:"poll-grow,omitempty"`
	PollShrink *int64 `json:"poll-shrink,omitempty"`
}

func (ObjectOptionsVariantIothread) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantIothread) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantIothread
	}{
		ObjectTypeIothread,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantMainLoop is an implementation of ObjectOptions.
type ObjectOptionsVariantMainLoop struct {
	ID string `json:"id"`
}

func (ObjectOptionsVariantMainLoop) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantMainLoop) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantMainLoop
	}{
		ObjectTypeMainLoop,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantMemoryBackendEpc is an implementation of ObjectOptions.
type ObjectOptionsVariantMemoryBackendEpc struct {
	ID string `json:"id"`
}

func (ObjectOptionsVariantMemoryBackendEpc) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantMemoryBackendEpc) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantMemoryBackendEpc
	}{
		ObjectTypeMemoryBackendEpc,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantMemoryBackendFile is an implementation of ObjectOptions.
type ObjectOptionsVariantMemoryBackendFile struct {
	ID          string  `json:"id"`
	Align       *uint64 `json:"align,omitempty"`
	DiscardData *bool   `json:"discard-data,omitempty"`
	MemPath     string  `json:"mem-path"`
	Pmem        *bool   `json:"pmem,omitempty"`
	Readonly    *bool   `json:"readonly,omitempty"`
}

func (ObjectOptionsVariantMemoryBackendFile) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantMemoryBackendFile) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantMemoryBackendFile
	}{
		ObjectTypeMemoryBackendFile,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantMemoryBackendMemfd is an implementation of ObjectOptions.
type ObjectOptionsVariantMemoryBackendMemfd struct {
	ID          string  `json:"id"`
	Hugetlb     *bool   `json:"hugetlb,omitempty"`
	Hugetlbsize *uint64 `json:"hugetlbsize,omitempty"`
	Seal        *bool   `json:"seal,omitempty"`
}

func (ObjectOptionsVariantMemoryBackendMemfd) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantMemoryBackendMemfd) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantMemoryBackendMemfd
	}{
		ObjectTypeMemoryBackendMemfd,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantMemoryBackendRAM is an implementation of ObjectOptions.
type ObjectOptionsVariantMemoryBackendRAM struct {
	ID                             string         `json:"id"`
	Dump                           *bool          `json:"dump,omitempty"`
	HostNodes                      []uint16       `json:"host-nodes,omitempty"`
	Merge                          *bool          `json:"merge,omitempty"`
	Policy                         *HostMemPolicy `json:"policy,omitempty"`
	Prealloc                       *bool          `json:"prealloc,omitempty"`
	PreallocThreads                *uint32        `json:"prealloc-threads,omitempty"`
	Share                          *bool          `json:"share,omitempty"`
	Reserve                        *bool          `json:"reserve,omitempty"`
	Size                           uint64         `json:"size"`
	XUseCanonicalPathForRamblockID *bool          `json:"x-use-canonical-path-for-ramblock-id,omitempty"`
}

func (ObjectOptionsVariantMemoryBackendRAM) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantMemoryBackendRAM) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantMemoryBackendRAM
	}{
		ObjectTypeMemoryBackendRAM,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantPrManagerHelper is an implementation of ObjectOptions.
type ObjectOptionsVariantPrManagerHelper struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

func (ObjectOptionsVariantPrManagerHelper) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantPrManagerHelper) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantPrManagerHelper
	}{
		ObjectTypePrManagerHelper,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantQtest is an implementation of ObjectOptions.
type ObjectOptionsVariantQtest struct {
	ID      string  `json:"id"`
	Chardev string  `json:"chardev"`
	Log     *string `json:"log,omitempty"`
}

func (ObjectOptionsVariantQtest) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantQtest) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantQtest
	}{
		ObjectTypeQtest,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantRngBuiltin is an implementation of ObjectOptions.
type ObjectOptionsVariantRngBuiltin struct {
	ID     string `json:"id"`
	Opened *bool  `json:"opened,omitempty"`
}

func (ObjectOptionsVariantRngBuiltin) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantRngBuiltin) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantRngBuiltin
	}{
		ObjectTypeRngBuiltin,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantRngEgd is an implementation of ObjectOptions.
type ObjectOptionsVariantRngEgd struct {
	ID      string `json:"id"`
	Chardev string `json:"chardev"`
}

func (ObjectOptionsVariantRngEgd) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantRngEgd) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantRngEgd
	}{
		ObjectTypeRngEgd,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantRngRandom is an implementation of ObjectOptions.
type ObjectOptionsVariantRngRandom struct {
	ID       string  `json:"id"`
	Filename *string `json:"filename,omitempty"`
}

func (ObjectOptionsVariantRngRandom) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantRngRandom) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantRngRandom
	}{
		ObjectTypeRngRandom,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantSecret is an implementation of ObjectOptions.
type ObjectOptionsVariantSecret struct {
	ID   string  `json:"id"`
	Data *string `json:"data,omitempty"`
	File *string `json:"file,omitempty"`
}

func (ObjectOptionsVariantSecret) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantSecret) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantSecret
	}{
		ObjectTypeSecret,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantSecretKeyring is an implementation of ObjectOptions.
type ObjectOptionsVariantSecretKeyring struct {
	ID     string `json:"id"`
	Serial int32  `json:"serial"`
}

func (ObjectOptionsVariantSecretKeyring) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantSecretKeyring) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantSecretKeyring
	}{
		ObjectTypeSecretKeyring,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantSevGuest is an implementation of ObjectOptions.
type ObjectOptionsVariantSevGuest struct {
	ID              string  `json:"id"`
	SevDevice       *string `json:"sev-device,omitempty"`
	DhCertFile      *string `json:"dh-cert-file,omitempty"`
	SessionFile     *string `json:"session-file,omitempty"`
	Policy          *uint32 `json:"policy,omitempty"`
	Handle          *uint32 `json:"handle,omitempty"`
	Cbitpos         *uint32 `json:"cbitpos,omitempty"`
	ReducedPhysBits uint32  `json:"reduced-phys-bits"`
	KernelHashes    *bool   `json:"kernel-hashes,omitempty"`
}

func (ObjectOptionsVariantSevGuest) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantSevGuest) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantSevGuest
	}{
		ObjectTypeSevGuest,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantThrottleGroup is an implementation of ObjectOptions.
type ObjectOptionsVariantThrottleGroup struct {
	ID                  string          `json:"id"`
	Limits              *ThrottleLimits `json:"limits,omitempty"`
	XIopsTotal          *int64          `json:"x-iops-total,omitempty"`
	XIopsTotalMax       *int64          `json:"x-iops-total-max,omitempty"`
	XIopsTotalMaxLength *int64          `json:"x-iops-total-max-length,omitempty"`
	XIopsRead           *int64          `json:"x-iops-read,omitempty"`
	XIopsReadMax        *int64          `json:"x-iops-read-max,omitempty"`
	XIopsReadMaxLength  *int64          `json:"x-iops-read-max-length,omitempty"`
	XIopsWrite          *int64          `json:"x-iops-write,omitempty"`
	XIopsWriteMax       *int64          `json:"x-iops-write-max,omitempty"`
	XIopsWriteMaxLength *int64          `json:"x-iops-write-max-length,omitempty"`
	XBpsTotal           *int64          `json:"x-bps-total,omitempty"`
	XBpsTotalMax        *int64          `json:"x-bps-total-max,omitempty"`
	XBpsTotalMaxLength  *int64          `json:"x-bps-total-max-length,omitempty"`
	XBpsRead            *int64          `json:"x-bps-read,omitempty"`
	XBpsReadMax         *int64          `json:"x-bps-read-max,omitempty"`
	XBpsReadMaxLength   *int64          `json:"x-bps-read-max-length,omitempty"`
	XBpsWrite           *int64          `json:"x-bps-write,omitempty"`
	XBpsWriteMax        *int64          `json:"x-bps-write-max,omitempty"`
	XBpsWriteMaxLength  *int64          `json:"x-bps-write-max-length,omitempty"`
	XIopsSize           *int64          `json:"x-iops-size,omitempty"`
}

func (ObjectOptionsVariantThrottleGroup) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantThrottleGroup) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantThrottleGroup
	}{
		ObjectTypeThrottleGroup,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantTLSCipherSuites is an implementation of ObjectOptions.
type ObjectOptionsVariantTLSCipherSuites struct {
	ID         string                   `json:"id"`
	VerifyPeer *bool                    `json:"verify-peer,omitempty"`
	Dir        *string                  `json:"dir,omitempty"`
	Endpoint   *QCryptoTLSCredsEndpoint `json:"endpoint,omitempty"`
	Priority   *string                  `json:"priority,omitempty"`
}

func (ObjectOptionsVariantTLSCipherSuites) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantTLSCipherSuites) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantTLSCipherSuites
	}{
		ObjectTypeTLSCipherSuites,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantTLSCredsAnon is an implementation of ObjectOptions.
type ObjectOptionsVariantTLSCredsAnon struct {
	ID     string `json:"id"`
	Loaded *bool  `json:"loaded,omitempty"`
}

func (ObjectOptionsVariantTLSCredsAnon) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantTLSCredsAnon) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantTLSCredsAnon
	}{
		ObjectTypeTLSCredsAnon,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantTLSCredsPsk is an implementation of ObjectOptions.
type ObjectOptionsVariantTLSCredsPsk struct {
	ID       string  `json:"id"`
	Loaded   *bool   `json:"loaded,omitempty"`
	Username *string `json:"username,omitempty"`
}

func (ObjectOptionsVariantTLSCredsPsk) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantTLSCredsPsk) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantTLSCredsPsk
	}{
		ObjectTypeTLSCredsPsk,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantTLSCredsX509 is an implementation of ObjectOptions.
type ObjectOptionsVariantTLSCredsX509 struct {
	ID          string  `json:"id"`
	Loaded      *bool   `json:"loaded,omitempty"`
	SanityCheck *bool   `json:"sanity-check,omitempty"`
	Passwordid  *string `json:"passwordid,omitempty"`
}

func (ObjectOptionsVariantTLSCredsX509) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantTLSCredsX509) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantTLSCredsX509
	}{
		ObjectTypeTLSCredsX509,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantXRemoteObject is an implementation of ObjectOptions.
type ObjectOptionsVariantXRemoteObject struct {
	ID    string `json:"id"`
	FD    string `json:"fd"`
	Devid string `json:"devid"`
}

func (ObjectOptionsVariantXRemoteObject) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantXRemoteObject) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantXRemoteObject
	}{
		ObjectTypeXRemoteObject,
		s,
	}
	return json.Marshal(v)
}

// ObjectOptionsVariantXVfioUserServer is an implementation of ObjectOptions.
type ObjectOptionsVariantXVfioUserServer struct {
	ID     string        `json:"id"`
	Socket SocketAddress `json:"socket"`
	Device string        `json:"device"`
}

func (ObjectOptionsVariantXVfioUserServer) isObjectOptions() {}

// MarshalJSON implements json.Marshaler.
func (s ObjectOptionsVariantXVfioUserServer) MarshalJSON() ([]byte, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
		ObjectOptionsVariantXVfioUserServer
	}{
		ObjectTypeXVfioUserServer,
		s,
	}
	return json.Marshal(v)
}

func decodeObjectOptions(bs json.RawMessage) (ObjectOptions, error) {
	v := struct {
		QomType ObjectType `json:"qom-type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.QomType {
	case ObjectTypeAuthzList:
		var ret ObjectOptionsVariantAuthzList
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeAuthzListfile:
		var ret ObjectOptionsVariantAuthzListfile
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeAuthzPam:
		var ret ObjectOptionsVariantAuthzPam
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeAuthzSimple:
		var ret ObjectOptionsVariantAuthzSimple
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeCanHostSocketcan:
		var ret ObjectOptionsVariantCanHostSocketcan
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeColoCompare:
		var ret ObjectOptionsVariantColoCompare
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeCryptodevBackend:
		var ret ObjectOptionsVariantCryptodevBackend
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeCryptodevBackendBuiltin:
		var ret ObjectOptionsVariantCryptodevBackendBuiltin
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeCryptodevVhostUser:
		var ret ObjectOptionsVariantCryptodevVhostUser
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeDbusVmstate:
		var ret ObjectOptionsVariantDbusVmstate
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeFilterBuffer:
		var ret ObjectOptionsVariantFilterBuffer
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeFilterDump:
		var ret ObjectOptionsVariantFilterDump
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeFilterMirror:
		var ret ObjectOptionsVariantFilterMirror
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeFilterRedirector:
		var ret ObjectOptionsVariantFilterRedirector
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeFilterReplay:
		var ret ObjectOptionsVariantFilterReplay
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeFilterRewriter:
		var ret ObjectOptionsVariantFilterRewriter
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeInputBarrier:
		var ret ObjectOptionsVariantInputBarrier
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeInputLinux:
		var ret ObjectOptionsVariantInputLinux
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeIothread:
		var ret ObjectOptionsVariantIothread
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeMainLoop:
		var ret ObjectOptionsVariantMainLoop
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeMemoryBackendEpc:
		var ret ObjectOptionsVariantMemoryBackendEpc
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeMemoryBackendFile:
		var ret ObjectOptionsVariantMemoryBackendFile
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeMemoryBackendMemfd:
		var ret ObjectOptionsVariantMemoryBackendMemfd
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeMemoryBackendRAM:
		var ret ObjectOptionsVariantMemoryBackendRAM
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypePrManagerHelper:
		var ret ObjectOptionsVariantPrManagerHelper
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeQtest:
		var ret ObjectOptionsVariantQtest
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeRngBuiltin:
		var ret ObjectOptionsVariantRngBuiltin
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeRngEgd:
		var ret ObjectOptionsVariantRngEgd
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeRngRandom:
		var ret ObjectOptionsVariantRngRandom
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeSecret:
		var ret ObjectOptionsVariantSecret
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeSecretKeyring:
		var ret ObjectOptionsVariantSecretKeyring
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeSevGuest:
		var ret ObjectOptionsVariantSevGuest
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeThrottleGroup:
		var ret ObjectOptionsVariantThrottleGroup
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeTLSCipherSuites:
		var ret ObjectOptionsVariantTLSCipherSuites
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeTLSCredsAnon:
		var ret ObjectOptionsVariantTLSCredsAnon
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeTLSCredsPsk:
		var ret ObjectOptionsVariantTLSCredsPsk
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeTLSCredsX509:
		var ret ObjectOptionsVariantTLSCredsX509
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeXRemoteObject:
		var ret ObjectOptionsVariantXRemoteObject
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case ObjectTypeXVfioUserServer:
		var ret ObjectOptionsVariantXVfioUserServer
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union ObjectOptions", v.QomType)
	}
}

// ObjectPropertyInfo -> ObjectPropertyInfo (struct)

// ObjectPropertyInfo implements the "ObjectPropertyInfo" QMP API type.
type ObjectPropertyInfo struct {
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	Description  *string      `json:"description,omitempty"`
	DefaultValue *interface{} `json:"default-value,omitempty"`
}

// ObjectType -> ObjectType (enum)

// ObjectType implements the "ObjectType" QMP API type.
type ObjectType int

// Known values of ObjectType.
const (
	ObjectTypeAuthzList ObjectType = iota
	ObjectTypeAuthzListfile
	ObjectTypeAuthzPam
	ObjectTypeAuthzSimple
	ObjectTypeCanBus
	ObjectTypeCanHostSocketcan
	ObjectTypeColoCompare
	ObjectTypeCryptodevBackend
	ObjectTypeCryptodevBackendBuiltin
	ObjectTypeCryptodevVhostUser
	ObjectTypeDbusVmstate
	ObjectTypeFilterBuffer
	ObjectTypeFilterDump
	ObjectTypeFilterMirror
	ObjectTypeFilterRedirector
	ObjectTypeFilterReplay
	ObjectTypeFilterRewriter
	ObjectTypeInputBarrier
	ObjectTypeInputLinux
	ObjectTypeIothread
	ObjectTypeMainLoop
	ObjectTypeMemoryBackendEpc
	ObjectTypeMemoryBackendFile
	ObjectTypeMemoryBackendMemfd
	ObjectTypeMemoryBackendRAM
	ObjectTypePefGuest
	ObjectTypePrManagerHelper
	ObjectTypeQtest
	ObjectTypeRngBuiltin
	ObjectTypeRngEgd
	ObjectTypeRngRandom
	ObjectTypeSecret
	ObjectTypeSecretKeyring
	ObjectTypeSevGuest
	ObjectTypeS390PvGuest
	ObjectTypeThrottleGroup
	ObjectTypeTLSCredsAnon
	ObjectTypeTLSCredsPsk
	ObjectTypeTLSCredsX509
	ObjectTypeTLSCipherSuites
	ObjectTypeXRemoteObject
	ObjectTypeXVfioUserServer
)

// String implements fmt.Stringer.
func (e ObjectType) String() string {
	switch e {
	case ObjectTypeAuthzList:
		return "authz-list"
	case ObjectTypeAuthzListfile:
		return "authz-listfile"
	case ObjectTypeAuthzPam:
		return "authz-pam"
	case ObjectTypeAuthzSimple:
		return "authz-simple"
	case ObjectTypeCanBus:
		return "can-bus"
	case ObjectTypeCanHostSocketcan:
		return "can-host-socketcan"
	case ObjectTypeColoCompare:
		return "colo-compare"
	case ObjectTypeCryptodevBackend:
		return "cryptodev-backend"
	case ObjectTypeCryptodevBackendBuiltin:
		return "cryptodev-backend-builtin"
	case ObjectTypeCryptodevVhostUser:
		return "cryptodev-vhost-user"
	case ObjectTypeDbusVmstate:
		return "dbus-vmstate"
	case ObjectTypeFilterBuffer:
		return "filter-buffer"
	case ObjectTypeFilterDump:
		return "filter-dump"
	case ObjectTypeFilterMirror:
		return "filter-mirror"
	case ObjectTypeFilterRedirector:
		return "filter-redirector"
	case ObjectTypeFilterReplay:
		return "filter-replay"
	case ObjectTypeFilterRewriter:
		return "filter-rewriter"
	case ObjectTypeInputBarrier:
		return "input-barrier"
	case ObjectTypeInputLinux:
		return "input-linux"
	case ObjectTypeIothread:
		return "iothread"
	case ObjectTypeMainLoop:
		return "main-loop"
	case ObjectTypeMemoryBackendEpc:
		return "memory-backend-epc"
	case ObjectTypeMemoryBackendFile:
		return "memory-backend-file"
	case ObjectTypeMemoryBackendMemfd:
		return "memory-backend-memfd"
	case ObjectTypeMemoryBackendRAM:
		return "memory-backend-ram"
	case ObjectTypePefGuest:
		return "pef-guest"
	case ObjectTypePrManagerHelper:
		return "pr-manager-helper"
	case ObjectTypeQtest:
		return "qtest"
	case ObjectTypeRngBuiltin:
		return "rng-builtin"
	case ObjectTypeRngEgd:
		return "rng-egd"
	case ObjectTypeRngRandom:
		return "rng-random"
	case ObjectTypeSecret:
		return "secret"
	case ObjectTypeSecretKeyring:
		return "secret_keyring"
	case ObjectTypeSevGuest:
		return "sev-guest"
	case ObjectTypeS390PvGuest:
		return "s390-pv-guest"
	case ObjectTypeThrottleGroup:
		return "throttle-group"
	case ObjectTypeTLSCredsAnon:
		return "tls-creds-anon"
	case ObjectTypeTLSCredsPsk:
		return "tls-creds-psk"
	case ObjectTypeTLSCredsX509:
		return "tls-creds-x509"
	case ObjectTypeTLSCipherSuites:
		return "tls-cipher-suites"
	case ObjectTypeXRemoteObject:
		return "x-remote-object"
	case ObjectTypeXVfioUserServer:
		return "x-vfio-user-server"
	default:
		return fmt.Sprintf("ObjectType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ObjectType) MarshalJSON() ([]byte, error) {
	switch e {
	case ObjectTypeAuthzList:
		return json.Marshal("authz-list")
	case ObjectTypeAuthzListfile:
		return json.Marshal("authz-listfile")
	case ObjectTypeAuthzPam:
		return json.Marshal("authz-pam")
	case ObjectTypeAuthzSimple:
		return json.Marshal("authz-simple")
	case ObjectTypeCanBus:
		return json.Marshal("can-bus")
	case ObjectTypeCanHostSocketcan:
		return json.Marshal("can-host-socketcan")
	case ObjectTypeColoCompare:
		return json.Marshal("colo-compare")
	case ObjectTypeCryptodevBackend:
		return json.Marshal("cryptodev-backend")
	case ObjectTypeCryptodevBackendBuiltin:
		return json.Marshal("cryptodev-backend-builtin")
	case ObjectTypeCryptodevVhostUser:
		return json.Marshal("cryptodev-vhost-user")
	case ObjectTypeDbusVmstate:
		return json.Marshal("dbus-vmstate")
	case ObjectTypeFilterBuffer:
		return json.Marshal("filter-buffer")
	case ObjectTypeFilterDump:
		return json.Marshal("filter-dump")
	case ObjectTypeFilterMirror:
		return json.Marshal("filter-mirror")
	case ObjectTypeFilterRedirector:
		return json.Marshal("filter-redirector")
	case ObjectTypeFilterReplay:
		return json.Marshal("filter-replay")
	case ObjectTypeFilterRewriter:
		return json.Marshal("filter-rewriter")
	case ObjectTypeInputBarrier:
		return json.Marshal("input-barrier")
	case ObjectTypeInputLinux:
		return json.Marshal("input-linux")
	case ObjectTypeIothread:
		return json.Marshal("iothread")
	case ObjectTypeMainLoop:
		return json.Marshal("main-loop")
	case ObjectTypeMemoryBackendEpc:
		return json.Marshal("memory-backend-epc")
	case ObjectTypeMemoryBackendFile:
		return json.Marshal("memory-backend-file")
	case ObjectTypeMemoryBackendMemfd:
		return json.Marshal("memory-backend-memfd")
	case ObjectTypeMemoryBackendRAM:
		return json.Marshal("memory-backend-ram")
	case ObjectTypePefGuest:
		return json.Marshal("pef-guest")
	case ObjectTypePrManagerHelper:
		return json.Marshal("pr-manager-helper")
	case ObjectTypeQtest:
		return json.Marshal("qtest")
	case ObjectTypeRngBuiltin:
		return json.Marshal("rng-builtin")
	case ObjectTypeRngEgd:
		return json.Marshal("rng-egd")
	case ObjectTypeRngRandom:
		return json.Marshal("rng-random")
	case ObjectTypeSecret:
		return json.Marshal("secret")
	case ObjectTypeSecretKeyring:
		return json.Marshal("secret_keyring")
	case ObjectTypeSevGuest:
		return json.Marshal("sev-guest")
	case ObjectTypeS390PvGuest:
		return json.Marshal("s390-pv-guest")
	case ObjectTypeThrottleGroup:
		return json.Marshal("throttle-group")
	case ObjectTypeTLSCredsAnon:
		return json.Marshal("tls-creds-anon")
	case ObjectTypeTLSCredsPsk:
		return json.Marshal("tls-creds-psk")
	case ObjectTypeTLSCredsX509:
		return json.Marshal("tls-creds-x509")
	case ObjectTypeTLSCipherSuites:
		return json.Marshal("tls-cipher-suites")
	case ObjectTypeXRemoteObject:
		return json.Marshal("x-remote-object")
	case ObjectTypeXVfioUserServer:
		return json.Marshal("x-vfio-user-server")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ObjectType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ObjectType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "authz-list":
		*e = ObjectTypeAuthzList
	case "authz-listfile":
		*e = ObjectTypeAuthzListfile
	case "authz-pam":
		*e = ObjectTypeAuthzPam
	case "authz-simple":
		*e = ObjectTypeAuthzSimple
	case "can-bus":
		*e = ObjectTypeCanBus
	case "can-host-socketcan":
		*e = ObjectTypeCanHostSocketcan
	case "colo-compare":
		*e = ObjectTypeColoCompare
	case "cryptodev-backend":
		*e = ObjectTypeCryptodevBackend
	case "cryptodev-backend-builtin":
		*e = ObjectTypeCryptodevBackendBuiltin
	case "cryptodev-vhost-user":
		*e = ObjectTypeCryptodevVhostUser
	case "dbus-vmstate":
		*e = ObjectTypeDbusVmstate
	case "filter-buffer":
		*e = ObjectTypeFilterBuffer
	case "filter-dump":
		*e = ObjectTypeFilterDump
	case "filter-mirror":
		*e = ObjectTypeFilterMirror
	case "filter-redirector":
		*e = ObjectTypeFilterRedirector
	case "filter-replay":
		*e = ObjectTypeFilterReplay
	case "filter-rewriter":
		*e = ObjectTypeFilterRewriter
	case "input-barrier":
		*e = ObjectTypeInputBarrier
	case "input-linux":
		*e = ObjectTypeInputLinux
	case "iothread":
		*e = ObjectTypeIothread
	case "main-loop":
		*e = ObjectTypeMainLoop
	case "memory-backend-epc":
		*e = ObjectTypeMemoryBackendEpc
	case "memory-backend-file":
		*e = ObjectTypeMemoryBackendFile
	case "memory-backend-memfd":
		*e = ObjectTypeMemoryBackendMemfd
	case "memory-backend-ram":
		*e = ObjectTypeMemoryBackendRAM
	case "pef-guest":
		*e = ObjectTypePefGuest
	case "pr-manager-helper":
		*e = ObjectTypePrManagerHelper
	case "qtest":
		*e = ObjectTypeQtest
	case "rng-builtin":
		*e = ObjectTypeRngBuiltin
	case "rng-egd":
		*e = ObjectTypeRngEgd
	case "rng-random":
		*e = ObjectTypeRngRandom
	case "secret":
		*e = ObjectTypeSecret
	case "secret_keyring":
		*e = ObjectTypeSecretKeyring
	case "sev-guest":
		*e = ObjectTypeSevGuest
	case "s390-pv-guest":
		*e = ObjectTypeS390PvGuest
	case "throttle-group":
		*e = ObjectTypeThrottleGroup
	case "tls-creds-anon":
		*e = ObjectTypeTLSCredsAnon
	case "tls-creds-psk":
		*e = ObjectTypeTLSCredsPsk
	case "tls-creds-x509":
		*e = ObjectTypeTLSCredsX509
	case "tls-cipher-suites":
		*e = ObjectTypeTLSCipherSuites
	case "x-remote-object":
		*e = ObjectTypeXRemoteObject
	case "x-vfio-user-server":
		*e = ObjectTypeXVfioUserServer
	default:
		return fmt.Errorf("unknown enum value %q for ObjectType", s)
	}
	return nil
}

// ObjectTypeInfo -> ObjectTypeInfo (struct)

// ObjectTypeInfo implements the "ObjectTypeInfo" QMP API type.
type ObjectTypeInfo struct {
	Name     string  `json:"name"`
	Abstract *bool   `json:"abstract,omitempty"`
	Parent   *string `json:"parent,omitempty"`
}

// OnCbwError -> OnCbwError (enum)

// OnCbwError implements the "OnCbwError" QMP API type.
type OnCbwError int

// Known values of OnCbwError.
const (
	OnCbwErrorBreakGuestWrite OnCbwError = iota
	OnCbwErrorBreakSnapshot
)

// String implements fmt.Stringer.
func (e OnCbwError) String() string {
	switch e {
	case OnCbwErrorBreakGuestWrite:
		return "break-guest-write"
	case OnCbwErrorBreakSnapshot:
		return "break-snapshot"
	default:
		return fmt.Sprintf("OnCbwError(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e OnCbwError) MarshalJSON() ([]byte, error) {
	switch e {
	case OnCbwErrorBreakGuestWrite:
		return json.Marshal("break-guest-write")
	case OnCbwErrorBreakSnapshot:
		return json.Marshal("break-snapshot")
	default:
		return nil, fmt.Errorf("unknown enum value %q for OnCbwError", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *OnCbwError) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "break-guest-write":
		*e = OnCbwErrorBreakGuestWrite
	case "break-snapshot":
		*e = OnCbwErrorBreakSnapshot
	default:
		return fmt.Errorf("unknown enum value %q for OnCbwError", s)
	}
	return nil
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

// PRManagerInfo -> PrManagerInfo (struct)

// PrManagerInfo implements the "PRManagerInfo" QMP API type.
type PrManagerInfo struct {
	ID        string `json:"id"`
	Connected bool   `json:"connected"`
}

// EVENT PR_MANAGER_STATUS_CHANGED

// PanicAction -> PanicAction (enum)

// PanicAction implements the "PanicAction" QMP API type.
type PanicAction int

// Known values of PanicAction.
const (
	PanicActionPause PanicAction = iota
	PanicActionShutdown
	PanicActionExitFailure
	PanicActionNone
)

// String implements fmt.Stringer.
func (e PanicAction) String() string {
	switch e {
	case PanicActionPause:
		return "pause"
	case PanicActionShutdown:
		return "shutdown"
	case PanicActionExitFailure:
		return "exit-failure"
	case PanicActionNone:
		return "none"
	default:
		return fmt.Sprintf("PanicAction(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e PanicAction) MarshalJSON() ([]byte, error) {
	switch e {
	case PanicActionPause:
		return json.Marshal("pause")
	case PanicActionShutdown:
		return json.Marshal("shutdown")
	case PanicActionExitFailure:
		return json.Marshal("exit-failure")
	case PanicActionNone:
		return json.Marshal("none")
	default:
		return nil, fmt.Errorf("unknown enum value %q for PanicAction", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *PanicAction) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "pause":
		*e = PanicActionPause
	case "shutdown":
		*e = PanicActionShutdown
	case "exit-failure":
		*e = PanicActionExitFailure
	case "none":
		*e = PanicActionNone
	default:
		return fmt.Errorf("unknown enum value %q for PanicAction", s)
	}
	return nil
}

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
	Device          int64  `json:"device"`
	Vendor          int64  `json:"vendor"`
	Subsystem       *int64 `json:"subsystem,omitempty"`
	SubsystemVendor *int64 `json:"subsystem-vendor,omitempty"`
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
	IrqPin    int64             `json:"irq_pin"`
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

// PreallocMode -> PreallocMode (enum)

// PreallocMode implements the "PreallocMode" QMP API type.
type PreallocMode int

// Known values of PreallocMode.
const (
	PreallocModeOff PreallocMode = iota
	PreallocModeMetadata
	PreallocModeFalloc
	PreallocModeFull
)

// String implements fmt.Stringer.
func (e PreallocMode) String() string {
	switch e {
	case PreallocModeOff:
		return "off"
	case PreallocModeMetadata:
		return "metadata"
	case PreallocModeFalloc:
		return "falloc"
	case PreallocModeFull:
		return "full"
	default:
		return fmt.Sprintf("PreallocMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e PreallocMode) MarshalJSON() ([]byte, error) {
	switch e {
	case PreallocModeOff:
		return json.Marshal("off")
	case PreallocModeMetadata:
		return json.Marshal("metadata")
	case PreallocModeFalloc:
		return json.Marshal("falloc")
	case PreallocModeFull:
		return json.Marshal("full")
	default:
		return nil, fmt.Errorf("unknown enum value %q for PreallocMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *PreallocMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "off":
		*e = PreallocModeOff
	case "metadata":
		*e = PreallocModeMetadata
	case "falloc":
		*e = PreallocModeFalloc
	case "full":
		*e = PreallocModeFull
	default:
		return fmt.Errorf("unknown enum value %q for PreallocMode", s)
	}
	return nil
}

// QAuthZListFormat -> QAuthZListFormat (enum)

// QAuthZListFormat implements the "QAuthZListFormat" QMP API type.
type QAuthZListFormat int

// Known values of QAuthZListFormat.
const (
	QAuthZListFormatExact QAuthZListFormat = iota
	QAuthZListFormatGlob
)

// String implements fmt.Stringer.
func (e QAuthZListFormat) String() string {
	switch e {
	case QAuthZListFormatExact:
		return "exact"
	case QAuthZListFormatGlob:
		return "glob"
	default:
		return fmt.Sprintf("QAuthZListFormat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QAuthZListFormat) MarshalJSON() ([]byte, error) {
	switch e {
	case QAuthZListFormatExact:
		return json.Marshal("exact")
	case QAuthZListFormatGlob:
		return json.Marshal("glob")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QAuthZListFormat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QAuthZListFormat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "exact":
		*e = QAuthZListFormatExact
	case "glob":
		*e = QAuthZListFormatGlob
	default:
		return fmt.Errorf("unknown enum value %q for QAuthZListFormat", s)
	}
	return nil
}

// QAuthZListPolicy -> QAuthZListPolicy (enum)

// QAuthZListPolicy implements the "QAuthZListPolicy" QMP API type.
type QAuthZListPolicy int

// Known values of QAuthZListPolicy.
const (
	QAuthZListPolicyDeny QAuthZListPolicy = iota
	QAuthZListPolicyAllow
)

// String implements fmt.Stringer.
func (e QAuthZListPolicy) String() string {
	switch e {
	case QAuthZListPolicyDeny:
		return "deny"
	case QAuthZListPolicyAllow:
		return "allow"
	default:
		return fmt.Sprintf("QAuthZListPolicy(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QAuthZListPolicy) MarshalJSON() ([]byte, error) {
	switch e {
	case QAuthZListPolicyDeny:
		return json.Marshal("deny")
	case QAuthZListPolicyAllow:
		return json.Marshal("allow")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QAuthZListPolicy", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QAuthZListPolicy) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "deny":
		*e = QAuthZListPolicyDeny
	case "allow":
		*e = QAuthZListPolicyAllow
	default:
		return fmt.Errorf("unknown enum value %q for QAuthZListPolicy", s)
	}
	return nil
}

// QAuthZListRule -> QAuthZListRule (struct)

// QAuthZListRule implements the "QAuthZListRule" QMP API type.
type QAuthZListRule struct {
	Match  string            `json:"match"`
	Policy QAuthZListPolicy  `json:"policy"`
	Format *QAuthZListFormat `json:"format,omitempty"`
}

// QCryptoBlockAmendOptions -> QCryptoBlockAmendOptions (flat union)

// QCryptoBlockAmendOptions implements the "QCryptoBlockAmendOptions" QMP API type.
//
// Can be one of:
//   - QCryptoBlockAmendOptionsVariantLUKS
type QCryptoBlockAmendOptions interface {
	isQCryptoBlockAmendOptions()
}

// QCryptoBlockAmendOptionsVariantLUKS is an implementation of QCryptoBlockAmendOptions.
type QCryptoBlockAmendOptionsVariantLUKS struct {
	State     QCryptoBlockLUKSKeyslotState `json:"state"`
	NewSecret *string                      `json:"new-secret,omitempty"`
	OldSecret *string                      `json:"old-secret,omitempty"`
	Keyslot   *int64                       `json:"keyslot,omitempty"`
	IterTime  *int64                       `json:"iter-time,omitempty"`
	Secret    *string                      `json:"secret,omitempty"`
}

func (QCryptoBlockAmendOptionsVariantLUKS) isQCryptoBlockAmendOptions() {}

// MarshalJSON implements json.Marshaler.
func (s QCryptoBlockAmendOptionsVariantLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Format QCryptoBlockFormat `json:"format"`
		QCryptoBlockAmendOptionsVariantLUKS
	}{
		QCryptoBlockFormatLUKS,
		s,
	}
	return json.Marshal(v)
}

func decodeQCryptoBlockAmendOptions(bs json.RawMessage) (QCryptoBlockAmendOptions, error) {
	v := struct {
		Format QCryptoBlockFormat `json:"format"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Format {
	case QCryptoBlockFormatLUKS:
		var ret QCryptoBlockAmendOptionsVariantLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union QCryptoBlockAmendOptions", v.Format)
	}
}

// QCryptoBlockCreateOptions -> QCryptoBlockCreateOptions (flat union)

// QCryptoBlockCreateOptions implements the "QCryptoBlockCreateOptions" QMP API type.
//
// Can be one of:
//   - QCryptoBlockCreateOptionsVariantLUKS
//   - QCryptoBlockCreateOptionsVariantQcow
type QCryptoBlockCreateOptions interface {
	isQCryptoBlockCreateOptions()
}

// QCryptoBlockCreateOptionsVariantLUKS is an implementation of QCryptoBlockCreateOptions.
type QCryptoBlockCreateOptionsVariantLUKS struct {
	CipherAlg    *QCryptoCipherAlgorithm `json:"cipher-alg,omitempty"`
	CipherMode   *QCryptoCipherMode      `json:"cipher-mode,omitempty"`
	IvgenAlg     *QCryptoIvGenAlgorithm  `json:"ivgen-alg,omitempty"`
	IvgenHashAlg *QCryptoHashAlgorithm   `json:"ivgen-hash-alg,omitempty"`
	HashAlg      *QCryptoHashAlgorithm   `json:"hash-alg,omitempty"`
	IterTime     *int64                  `json:"iter-time,omitempty"`
}

func (QCryptoBlockCreateOptionsVariantLUKS) isQCryptoBlockCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s QCryptoBlockCreateOptionsVariantLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Format QCryptoBlockFormat `json:"format"`
		QCryptoBlockCreateOptionsVariantLUKS
	}{
		QCryptoBlockFormatLUKS,
		s,
	}
	return json.Marshal(v)
}

// QCryptoBlockCreateOptionsVariantQcow is an implementation of QCryptoBlockCreateOptions.
type QCryptoBlockCreateOptionsVariantQcow struct {
	KeySecret *string `json:"key-secret,omitempty"`
}

func (QCryptoBlockCreateOptionsVariantQcow) isQCryptoBlockCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s QCryptoBlockCreateOptionsVariantQcow) MarshalJSON() ([]byte, error) {
	v := struct {
		Format QCryptoBlockFormat `json:"format"`
		QCryptoBlockCreateOptionsVariantQcow
	}{
		QCryptoBlockFormatQcow,
		s,
	}
	return json.Marshal(v)
}

func decodeQCryptoBlockCreateOptions(bs json.RawMessage) (QCryptoBlockCreateOptions, error) {
	v := struct {
		Format QCryptoBlockFormat `json:"format"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Format {
	case QCryptoBlockFormatLUKS:
		var ret QCryptoBlockCreateOptionsVariantLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case QCryptoBlockFormatQcow:
		var ret QCryptoBlockCreateOptionsVariantQcow
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union QCryptoBlockCreateOptions", v.Format)
	}
}

// QCryptoBlockFormat -> QCryptoBlockFormat (enum)

// QCryptoBlockFormat implements the "QCryptoBlockFormat" QMP API type.
type QCryptoBlockFormat int

// Known values of QCryptoBlockFormat.
const (
	QCryptoBlockFormatQcow QCryptoBlockFormat = iota
	QCryptoBlockFormatLUKS
)

// String implements fmt.Stringer.
func (e QCryptoBlockFormat) String() string {
	switch e {
	case QCryptoBlockFormatQcow:
		return "qcow"
	case QCryptoBlockFormatLUKS:
		return "luks"
	default:
		return fmt.Sprintf("QCryptoBlockFormat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QCryptoBlockFormat) MarshalJSON() ([]byte, error) {
	switch e {
	case QCryptoBlockFormatQcow:
		return json.Marshal("qcow")
	case QCryptoBlockFormatLUKS:
		return json.Marshal("luks")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QCryptoBlockFormat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QCryptoBlockFormat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "qcow":
		*e = QCryptoBlockFormatQcow
	case "luks":
		*e = QCryptoBlockFormatLUKS
	default:
		return fmt.Errorf("unknown enum value %q for QCryptoBlockFormat", s)
	}
	return nil
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

// QCryptoBlockLUKSKeyslotState -> QCryptoBlockLUKSKeyslotState (enum)

// QCryptoBlockLUKSKeyslotState implements the "QCryptoBlockLUKSKeyslotState" QMP API type.
type QCryptoBlockLUKSKeyslotState int

// Known values of QCryptoBlockLUKSKeyslotState.
const (
	QCryptoBlockLUKSKeyslotStateActive QCryptoBlockLUKSKeyslotState = iota
	QCryptoBlockLUKSKeyslotStateInactive
)

// String implements fmt.Stringer.
func (e QCryptoBlockLUKSKeyslotState) String() string {
	switch e {
	case QCryptoBlockLUKSKeyslotStateActive:
		return "active"
	case QCryptoBlockLUKSKeyslotStateInactive:
		return "inactive"
	default:
		return fmt.Sprintf("QCryptoBlockLUKSKeyslotState(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QCryptoBlockLUKSKeyslotState) MarshalJSON() ([]byte, error) {
	switch e {
	case QCryptoBlockLUKSKeyslotStateActive:
		return json.Marshal("active")
	case QCryptoBlockLUKSKeyslotStateInactive:
		return json.Marshal("inactive")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QCryptoBlockLUKSKeyslotState", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QCryptoBlockLUKSKeyslotState) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "active":
		*e = QCryptoBlockLUKSKeyslotStateActive
	case "inactive":
		*e = QCryptoBlockLUKSKeyslotStateInactive
	default:
		return fmt.Errorf("unknown enum value %q for QCryptoBlockLUKSKeyslotState", s)
	}
	return nil
}

// QCryptoCipherAlgorithm -> QCryptoCipherAlgorithm (enum)

// QCryptoCipherAlgorithm implements the "QCryptoCipherAlgorithm" QMP API type.
type QCryptoCipherAlgorithm int

// Known values of QCryptoCipherAlgorithm.
const (
	QCryptoCipherAlgorithmAes128 QCryptoCipherAlgorithm = iota
	QCryptoCipherAlgorithmAes192
	QCryptoCipherAlgorithmAes256
	QCryptoCipherAlgorithmDes
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
	case QCryptoCipherAlgorithmDes:
		return "des"
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
	case QCryptoCipherAlgorithmDes:
		return json.Marshal("des")
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
	case "des":
		*e = QCryptoCipherAlgorithmDes
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

// QCryptoTLSCredsEndpoint -> QCryptoTLSCredsEndpoint (enum)

// QCryptoTLSCredsEndpoint implements the "QCryptoTLSCredsEndpoint" QMP API type.
type QCryptoTLSCredsEndpoint int

// Known values of QCryptoTLSCredsEndpoint.
const (
	QCryptoTLSCredsEndpointClient QCryptoTLSCredsEndpoint = iota
	QCryptoTLSCredsEndpointServer
)

// String implements fmt.Stringer.
func (e QCryptoTLSCredsEndpoint) String() string {
	switch e {
	case QCryptoTLSCredsEndpointClient:
		return "client"
	case QCryptoTLSCredsEndpointServer:
		return "server"
	default:
		return fmt.Sprintf("QCryptoTLSCredsEndpoint(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QCryptoTLSCredsEndpoint) MarshalJSON() ([]byte, error) {
	switch e {
	case QCryptoTLSCredsEndpointClient:
		return json.Marshal("client")
	case QCryptoTLSCredsEndpointServer:
		return json.Marshal("server")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QCryptoTLSCredsEndpoint", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QCryptoTLSCredsEndpoint) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "client":
		*e = QCryptoTLSCredsEndpointClient
	case "server":
		*e = QCryptoTLSCredsEndpointServer
	default:
		return fmt.Errorf("unknown enum value %q for QCryptoTLSCredsEndpoint", s)
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
	QKeyCodeMuhenkan
	QKeyCodeKatakanahiragana
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
	QKeyCodeLang1
	QKeyCodeLang2
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
	case QKeyCodeMuhenkan:
		return "muhenkan"
	case QKeyCodeKatakanahiragana:
		return "katakanahiragana"
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
	case QKeyCodeLang1:
		return "lang1"
	case QKeyCodeLang2:
		return "lang2"
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
	case QKeyCodeMuhenkan:
		return json.Marshal("muhenkan")
	case QKeyCodeKatakanahiragana:
		return json.Marshal("katakanahiragana")
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
	case QKeyCodeLang1:
		return json.Marshal("lang1")
	case QKeyCodeLang2:
		return json.Marshal("lang2")
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
	case "muhenkan":
		*e = QKeyCodeMuhenkan
	case "katakanahiragana":
		*e = QKeyCodeKatakanahiragana
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
	case "lang1":
		*e = QKeyCodeLang1
	case "lang2":
		*e = QKeyCodeLang2
	default:
		return fmt.Errorf("unknown enum value %q for QKeyCode", s)
	}
	return nil
}

// QMPCapability -> QMPCapability (enum)

// QMPCapability implements the "QMPCapability" QMP API type.
type QMPCapability int

// Known values of QMPCapability.
const (
	QMPCapabilityOob QMPCapability = iota
)

// String implements fmt.Stringer.
func (e QMPCapability) String() string {
	switch e {
	case QMPCapabilityOob:
		return "oob"
	default:
		return fmt.Sprintf("QMPCapability(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e QMPCapability) MarshalJSON() ([]byte, error) {
	switch e {
	case QMPCapabilityOob:
		return json.Marshal("oob")
	default:
		return nil, fmt.Errorf("unknown enum value %q for QMPCapability", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *QMPCapability) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "oob":
		*e = QMPCapabilityOob
	default:
		return fmt.Errorf("unknown enum value %q for QMPCapability", s)
	}
	return nil
}

// EVENT QUORUM_FAILURE

// EVENT QUORUM_REPORT_BAD

// Qcow2BitmapInfo -> Qcow2BitmapInfo (struct)

// Qcow2BitmapInfo implements the "Qcow2BitmapInfo" QMP API type.
type Qcow2BitmapInfo struct {
	Name        string                 `json:"name"`
	Granularity uint32                 `json:"granularity"`
	Flags       []Qcow2BitmapInfoFlags `json:"flags"`
}

// Qcow2BitmapInfoFlags -> Qcow2BitmapInfoFlags (enum)

// Qcow2BitmapInfoFlags implements the "Qcow2BitmapInfoFlags" QMP API type.
type Qcow2BitmapInfoFlags int

// Known values of Qcow2BitmapInfoFlags.
const (
	Qcow2BitmapInfoFlagsInUse Qcow2BitmapInfoFlags = iota
	Qcow2BitmapInfoFlagsAuto
)

// String implements fmt.Stringer.
func (e Qcow2BitmapInfoFlags) String() string {
	switch e {
	case Qcow2BitmapInfoFlagsInUse:
		return "in-use"
	case Qcow2BitmapInfoFlagsAuto:
		return "auto"
	default:
		return fmt.Sprintf("Qcow2BitmapInfoFlags(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e Qcow2BitmapInfoFlags) MarshalJSON() ([]byte, error) {
	switch e {
	case Qcow2BitmapInfoFlagsInUse:
		return json.Marshal("in-use")
	case Qcow2BitmapInfoFlagsAuto:
		return json.Marshal("auto")
	default:
		return nil, fmt.Errorf("unknown enum value %q for Qcow2BitmapInfoFlags", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Qcow2BitmapInfoFlags) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "in-use":
		*e = Qcow2BitmapInfoFlagsInUse
	case "auto":
		*e = Qcow2BitmapInfoFlagsAuto
	default:
		return fmt.Errorf("unknown enum value %q for Qcow2BitmapInfoFlags", s)
	}
	return nil
}

// Qcow2CompressionType -> Qcow2CompressionType (enum)

// Qcow2CompressionType implements the "Qcow2CompressionType" QMP API type.
type Qcow2CompressionType int

// Known values of Qcow2CompressionType.
const (
	Qcow2CompressionTypeZlib Qcow2CompressionType = iota
	Qcow2CompressionTypeZstd
)

// String implements fmt.Stringer.
func (e Qcow2CompressionType) String() string {
	switch e {
	case Qcow2CompressionTypeZlib:
		return "zlib"
	case Qcow2CompressionTypeZstd:
		return "zstd"
	default:
		return fmt.Sprintf("Qcow2CompressionType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e Qcow2CompressionType) MarshalJSON() ([]byte, error) {
	switch e {
	case Qcow2CompressionTypeZlib:
		return json.Marshal("zlib")
	case Qcow2CompressionTypeZstd:
		return json.Marshal("zstd")
	default:
		return nil, fmt.Errorf("unknown enum value %q for Qcow2CompressionType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Qcow2CompressionType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "zlib":
		*e = Qcow2CompressionTypeZlib
	case "zstd":
		*e = Qcow2CompressionTypeZstd
	default:
		return fmt.Errorf("unknown enum value %q for Qcow2CompressionType", s)
	}
	return nil
}

// Qcow2OverlapCheckFlags -> Qcow2OverlapCheckFlags (struct)

// Qcow2OverlapCheckFlags implements the "Qcow2OverlapCheckFlags" QMP API type.
type Qcow2OverlapCheckFlags struct {
	Template        *Qcow2OverlapCheckMode `json:"template,omitempty"`
	MainHeader      *bool                  `json:"main-header,omitempty"`
	ActiveL1        *bool                  `json:"active-l1,omitempty"`
	ActiveL2        *bool                  `json:"active-l2,omitempty"`
	RefcountTable   *bool                  `json:"refcount-table,omitempty"`
	RefcountBlock   *bool                  `json:"refcount-block,omitempty"`
	SnapshotTable   *bool                  `json:"snapshot-table,omitempty"`
	InactiveL1      *bool                  `json:"inactive-l1,omitempty"`
	InactiveL2      *bool                  `json:"inactive-l2,omitempty"`
	BitmapDirectory *bool                  `json:"bitmap-directory,omitempty"`
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
//   - Qcow2OverlapChecksVariantFlags
//   - Qcow2OverlapChecksVariantMode
type Qcow2OverlapChecks interface {
	isQcow2OverlapChecks()
}

// Qcow2OverlapChecksVariantFlags is an implementation of Qcow2OverlapChecks
type Qcow2OverlapChecksVariantFlags Qcow2OverlapCheckFlags

func (Qcow2OverlapChecksVariantFlags) isQcow2OverlapChecks() {}

// Qcow2OverlapChecksVariantMode is an implementation of Qcow2OverlapChecks
type Qcow2OverlapChecksVariantMode Qcow2OverlapCheckMode

func (Qcow2OverlapChecksVariantMode) isQcow2OverlapChecks() {}

func decodeQcow2OverlapChecks(bs json.RawMessage) (Qcow2OverlapChecks, error) {

	var flags Qcow2OverlapChecksVariantFlags
	if err := json.Unmarshal([]byte(bs), &flags); err == nil {
		return flags, nil
	}

	var mode Qcow2OverlapChecksVariantMode
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

// EVENT RDMA_GID_STATUS_CHANGED

// EVENT RESET

// EVENT RESUME

// EVENT RTC_CHANGE

// RbdAuthMode -> RbdAuthMode (enum)

// RbdAuthMode implements the "RbdAuthMode" QMP API type.
type RbdAuthMode int

// Known values of RbdAuthMode.
const (
	RbdAuthModeCephx RbdAuthMode = iota
	RbdAuthModeNone
)

// String implements fmt.Stringer.
func (e RbdAuthMode) String() string {
	switch e {
	case RbdAuthModeCephx:
		return "cephx"
	case RbdAuthModeNone:
		return "none"
	default:
		return fmt.Sprintf("RbdAuthMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e RbdAuthMode) MarshalJSON() ([]byte, error) {
	switch e {
	case RbdAuthModeCephx:
		return json.Marshal("cephx")
	case RbdAuthModeNone:
		return json.Marshal("none")
	default:
		return nil, fmt.Errorf("unknown enum value %q for RbdAuthMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *RbdAuthMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "cephx":
		*e = RbdAuthModeCephx
	case "none":
		*e = RbdAuthModeNone
	default:
		return fmt.Errorf("unknown enum value %q for RbdAuthMode", s)
	}
	return nil
}

// RbdEncryptionCreateOptions -> RbdEncryptionCreateOptions (flat union)

// RbdEncryptionCreateOptions implements the "RbdEncryptionCreateOptions" QMP API type.
//
// Can be one of:
//   - RbdEncryptionCreateOptionsVariantLUKS
//   - RbdEncryptionCreateOptionsVariantLUKS2
type RbdEncryptionCreateOptions interface {
	isRbdEncryptionCreateOptions()
}

// RbdEncryptionCreateOptionsVariantLUKS is an implementation of RbdEncryptionCreateOptions.
type RbdEncryptionCreateOptionsVariantLUKS struct {
}

func (RbdEncryptionCreateOptionsVariantLUKS) isRbdEncryptionCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s RbdEncryptionCreateOptionsVariantLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Format RbdImageEncryptionFormat `json:"format"`
		RbdEncryptionCreateOptionsVariantLUKS
	}{
		RbdImageEncryptionFormatLUKS,
		s,
	}
	return json.Marshal(v)
}

// RbdEncryptionCreateOptionsVariantLUKS2 is an implementation of RbdEncryptionCreateOptions.
type RbdEncryptionCreateOptionsVariantLUKS2 struct {
}

func (RbdEncryptionCreateOptionsVariantLUKS2) isRbdEncryptionCreateOptions() {}

// MarshalJSON implements json.Marshaler.
func (s RbdEncryptionCreateOptionsVariantLUKS2) MarshalJSON() ([]byte, error) {
	v := struct {
		Format RbdImageEncryptionFormat `json:"format"`
		RbdEncryptionCreateOptionsVariantLUKS2
	}{
		RbdImageEncryptionFormatLUKS2,
		s,
	}
	return json.Marshal(v)
}

func decodeRbdEncryptionCreateOptions(bs json.RawMessage) (RbdEncryptionCreateOptions, error) {
	v := struct {
		Format RbdImageEncryptionFormat `json:"format"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Format {
	case RbdImageEncryptionFormatLUKS:
		var ret RbdEncryptionCreateOptionsVariantLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case RbdImageEncryptionFormatLUKS2:
		var ret RbdEncryptionCreateOptionsVariantLUKS2
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union RbdEncryptionCreateOptions", v.Format)
	}
}

// RbdEncryptionOptions -> RbdEncryptionOptions (flat union)

// RbdEncryptionOptions implements the "RbdEncryptionOptions" QMP API type.
//
// Can be one of:
//   - RbdEncryptionOptionsVariantLUKS
//   - RbdEncryptionOptionsVariantLUKS2
type RbdEncryptionOptions interface {
	isRbdEncryptionOptions()
}

// RbdEncryptionOptionsVariantLUKS is an implementation of RbdEncryptionOptions.
type RbdEncryptionOptionsVariantLUKS struct {
}

func (RbdEncryptionOptionsVariantLUKS) isRbdEncryptionOptions() {}

// MarshalJSON implements json.Marshaler.
func (s RbdEncryptionOptionsVariantLUKS) MarshalJSON() ([]byte, error) {
	v := struct {
		Format RbdImageEncryptionFormat `json:"format"`
		RbdEncryptionOptionsVariantLUKS
	}{
		RbdImageEncryptionFormatLUKS,
		s,
	}
	return json.Marshal(v)
}

// RbdEncryptionOptionsVariantLUKS2 is an implementation of RbdEncryptionOptions.
type RbdEncryptionOptionsVariantLUKS2 struct {
}

func (RbdEncryptionOptionsVariantLUKS2) isRbdEncryptionOptions() {}

// MarshalJSON implements json.Marshaler.
func (s RbdEncryptionOptionsVariantLUKS2) MarshalJSON() ([]byte, error) {
	v := struct {
		Format RbdImageEncryptionFormat `json:"format"`
		RbdEncryptionOptionsVariantLUKS2
	}{
		RbdImageEncryptionFormatLUKS2,
		s,
	}
	return json.Marshal(v)
}

func decodeRbdEncryptionOptions(bs json.RawMessage) (RbdEncryptionOptions, error) {
	v := struct {
		Format RbdImageEncryptionFormat `json:"format"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Format {
	case RbdImageEncryptionFormatLUKS:
		var ret RbdEncryptionOptionsVariantLUKS
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case RbdImageEncryptionFormatLUKS2:
		var ret RbdEncryptionOptionsVariantLUKS2
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union RbdEncryptionOptions", v.Format)
	}
}

// RbdImageEncryptionFormat -> RbdImageEncryptionFormat (enum)

// RbdImageEncryptionFormat implements the "RbdImageEncryptionFormat" QMP API type.
type RbdImageEncryptionFormat int

// Known values of RbdImageEncryptionFormat.
const (
	RbdImageEncryptionFormatLUKS RbdImageEncryptionFormat = iota
	RbdImageEncryptionFormatLUKS2
)

// String implements fmt.Stringer.
func (e RbdImageEncryptionFormat) String() string {
	switch e {
	case RbdImageEncryptionFormatLUKS:
		return "luks"
	case RbdImageEncryptionFormatLUKS2:
		return "luks2"
	default:
		return fmt.Sprintf("RbdImageEncryptionFormat(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e RbdImageEncryptionFormat) MarshalJSON() ([]byte, error) {
	switch e {
	case RbdImageEncryptionFormatLUKS:
		return json.Marshal("luks")
	case RbdImageEncryptionFormatLUKS2:
		return json.Marshal("luks2")
	default:
		return nil, fmt.Errorf("unknown enum value %q for RbdImageEncryptionFormat", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *RbdImageEncryptionFormat) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "luks":
		*e = RbdImageEncryptionFormatLUKS
	case "luks2":
		*e = RbdImageEncryptionFormatLUKS2
	default:
		return fmt.Errorf("unknown enum value %q for RbdImageEncryptionFormat", s)
	}
	return nil
}

// RebootAction -> RebootAction (enum)

// RebootAction implements the "RebootAction" QMP API type.
type RebootAction int

// Known values of RebootAction.
const (
	RebootActionReset RebootAction = iota
	RebootActionShutdown
)

// String implements fmt.Stringer.
func (e RebootAction) String() string {
	switch e {
	case RebootActionReset:
		return "reset"
	case RebootActionShutdown:
		return "shutdown"
	default:
		return fmt.Sprintf("RebootAction(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e RebootAction) MarshalJSON() ([]byte, error) {
	switch e {
	case RebootActionReset:
		return json.Marshal("reset")
	case RebootActionShutdown:
		return json.Marshal("shutdown")
	default:
		return nil, fmt.Errorf("unknown enum value %q for RebootAction", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *RebootAction) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "reset":
		*e = RebootActionReset
	case "shutdown":
		*e = RebootActionShutdown
	default:
		return fmt.Errorf("unknown enum value %q for RebootAction", s)
	}
	return nil
}

// ReplayInfo -> ReplayInfo (struct)

// ReplayInfo implements the "ReplayInfo" QMP API type.
type ReplayInfo struct {
	Mode     ReplayMode `json:"mode"`
	Filename *string    `json:"filename,omitempty"`
	Icount   int64      `json:"icount"`
}

// ReplayMode -> ReplayMode (enum)

// ReplayMode implements the "ReplayMode" QMP API type.
type ReplayMode int

// Known values of ReplayMode.
const (
	ReplayModeNone ReplayMode = iota
	ReplayModeRecord
	ReplayModePlay
)

// String implements fmt.Stringer.
func (e ReplayMode) String() string {
	switch e {
	case ReplayModeNone:
		return "none"
	case ReplayModeRecord:
		return "record"
	case ReplayModePlay:
		return "play"
	default:
		return fmt.Sprintf("ReplayMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ReplayMode) MarshalJSON() ([]byte, error) {
	switch e {
	case ReplayModeNone:
		return json.Marshal("none")
	case ReplayModeRecord:
		return json.Marshal("record")
	case ReplayModePlay:
		return json.Marshal("play")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ReplayMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ReplayMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = ReplayModeNone
	case "record":
		*e = ReplayModeRecord
	case "play":
		*e = ReplayModePlay
	default:
		return fmt.Errorf("unknown enum value %q for ReplayMode", s)
	}
	return nil
}

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

// S390CrashReason -> S390CrashReason (enum)

// S390CrashReason implements the "S390CrashReason" QMP API type.
type S390CrashReason int

// Known values of S390CrashReason.
const (
	S390CrashReasonUnknown S390CrashReason = iota
	S390CrashReasonDisabledWait
	S390CrashReasonExtintLoop
	S390CrashReasonPgmintLoop
	S390CrashReasonOpintLoop
)

// String implements fmt.Stringer.
func (e S390CrashReason) String() string {
	switch e {
	case S390CrashReasonUnknown:
		return "unknown"
	case S390CrashReasonDisabledWait:
		return "disabled-wait"
	case S390CrashReasonExtintLoop:
		return "extint-loop"
	case S390CrashReasonPgmintLoop:
		return "pgmint-loop"
	case S390CrashReasonOpintLoop:
		return "opint-loop"
	default:
		return fmt.Sprintf("S390CrashReason(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e S390CrashReason) MarshalJSON() ([]byte, error) {
	switch e {
	case S390CrashReasonUnknown:
		return json.Marshal("unknown")
	case S390CrashReasonDisabledWait:
		return json.Marshal("disabled-wait")
	case S390CrashReasonExtintLoop:
		return json.Marshal("extint-loop")
	case S390CrashReasonPgmintLoop:
		return json.Marshal("pgmint-loop")
	case S390CrashReasonOpintLoop:
		return json.Marshal("opint-loop")
	default:
		return nil, fmt.Errorf("unknown enum value %q for S390CrashReason", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *S390CrashReason) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "unknown":
		*e = S390CrashReasonUnknown
	case "disabled-wait":
		*e = S390CrashReasonDisabledWait
	case "extint-loop":
		*e = S390CrashReasonExtintLoop
	case "pgmint-loop":
		*e = S390CrashReasonPgmintLoop
	case "opint-loop":
		*e = S390CrashReasonOpintLoop
	default:
		return fmt.Errorf("unknown enum value %q for S390CrashReason", s)
	}
	return nil
}

// SGXEPCSection -> SgxepcSection (struct)

// SgxepcSection implements the "SGXEPCSection" QMP API type.
type SgxepcSection struct {
	Node int64  `json:"node"`
	Size uint64 `json:"size"`
}

// SGXInfo -> SgxInfo (struct)

// SgxInfo implements the "SGXInfo" QMP API type.
type SgxInfo struct {
	Sgx         bool            `json:"sgx"`
	Sgx1        bool            `json:"sgx1"`
	Sgx2        bool            `json:"sgx2"`
	Flc         bool            `json:"flc"`
	SectionSize uint64          `json:"section-size"`
	Sections    []SgxepcSection `json:"sections"`
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
//   - SchemaInfoVariantAlternate
//   - SchemaInfoVariantArray
//   - SchemaInfoVariantBuiltin
//   - SchemaInfoVariantCommand
//   - SchemaInfoVariantEnum
//   - SchemaInfoVariantEvent
//   - SchemaInfoVariantObject
type SchemaInfo interface {
	isSchemaInfo()
}

// SchemaInfoVariantAlternate is an implementation of SchemaInfo.
type SchemaInfoVariantAlternate struct {
	Name     string                      `json:"name"`
	Features []string                    `json:"features,omitempty"`
	Members  []SchemaInfoAlternateMember `json:"members"`
}

func (SchemaInfoVariantAlternate) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoVariantAlternate) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoVariantAlternate
	}{
		SchemaMetaTypeAlternate,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoVariantArray is an implementation of SchemaInfo.
type SchemaInfoVariantArray struct {
	Name        string   `json:"name"`
	Features    []string `json:"features,omitempty"`
	ElementType string   `json:"element-type"`
}

func (SchemaInfoVariantArray) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoVariantArray) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoVariantArray
	}{
		SchemaMetaTypeArray,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoVariantBuiltin is an implementation of SchemaInfo.
type SchemaInfoVariantBuiltin struct {
	Name     string   `json:"name"`
	Features []string `json:"features,omitempty"`
	JSONType JSONType `json:"json-type"`
}

func (SchemaInfoVariantBuiltin) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoVariantBuiltin) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoVariantBuiltin
	}{
		SchemaMetaTypeBuiltin,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoVariantCommand is an implementation of SchemaInfo.
type SchemaInfoVariantCommand struct {
	Name     string   `json:"name"`
	Features []string `json:"features,omitempty"`
	ArgType  string   `json:"arg-type"`
	RetType  string   `json:"ret-type"`
	AllowOob *bool    `json:"allow-oob,omitempty"`
}

func (SchemaInfoVariantCommand) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoVariantCommand) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoVariantCommand
	}{
		SchemaMetaTypeCommand,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoVariantEnum is an implementation of SchemaInfo.
type SchemaInfoVariantEnum struct {
	Name     string                 `json:"name"`
	Features []string               `json:"features,omitempty"`
	Members  []SchemaInfoEnumMember `json:"members"`
	Values   []string               `json:"values"`
}

func (SchemaInfoVariantEnum) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoVariantEnum) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoVariantEnum
	}{
		SchemaMetaTypeEnum,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoVariantEvent is an implementation of SchemaInfo.
type SchemaInfoVariantEvent struct {
	Name     string   `json:"name"`
	Features []string `json:"features,omitempty"`
	ArgType  string   `json:"arg-type"`
}

func (SchemaInfoVariantEvent) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoVariantEvent) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoVariantEvent
	}{
		SchemaMetaTypeEvent,
		s,
	}
	return json.Marshal(v)
}

// SchemaInfoVariantObject is an implementation of SchemaInfo.
type SchemaInfoVariantObject struct {
	Name     string                    `json:"name"`
	Features []string                  `json:"features,omitempty"`
	Members  []SchemaInfoObjectMember  `json:"members"`
	Tag      *string                   `json:"tag,omitempty"`
	Variants []SchemaInfoObjectVariant `json:"variants,omitempty"`
}

func (SchemaInfoVariantObject) isSchemaInfo() {}

// MarshalJSON implements json.Marshaler.
func (s SchemaInfoVariantObject) MarshalJSON() ([]byte, error) {
	v := struct {
		MetaType SchemaMetaType `json:"meta-type"`
		SchemaInfoVariantObject
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
		var ret SchemaInfoVariantAlternate
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeArray:
		var ret SchemaInfoVariantArray
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeBuiltin:
		var ret SchemaInfoVariantBuiltin
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeCommand:
		var ret SchemaInfoVariantCommand
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeEnum:
		var ret SchemaInfoVariantEnum
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeEvent:
		var ret SchemaInfoVariantEvent
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SchemaMetaTypeObject:
		var ret SchemaInfoVariantObject
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

// SchemaInfoEnumMember -> SchemaInfoEnumMember (struct)

// SchemaInfoEnumMember implements the "SchemaInfoEnumMember" QMP API type.
type SchemaInfoEnumMember struct {
	Name     string   `json:"name"`
	Features []string `json:"features,omitempty"`
}

// SchemaInfoObjectMember -> SchemaInfoObjectMember (struct)

// SchemaInfoObjectMember implements the "SchemaInfoObjectMember" QMP API type.
type SchemaInfoObjectMember struct {
	Name     string       `json:"name"`
	Type     string       `json:"type"`
	Default  *interface{} `json:"default,omitempty"`
	Features []string     `json:"features,omitempty"`
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

// SetPasswordAction -> SetPasswordAction (enum)

// SetPasswordAction implements the "SetPasswordAction" QMP API type.
type SetPasswordAction int

// Known values of SetPasswordAction.
const (
	SetPasswordActionKeep SetPasswordAction = iota
	SetPasswordActionFail
	SetPasswordActionDisconnect
)

// String implements fmt.Stringer.
func (e SetPasswordAction) String() string {
	switch e {
	case SetPasswordActionKeep:
		return "keep"
	case SetPasswordActionFail:
		return "fail"
	case SetPasswordActionDisconnect:
		return "disconnect"
	default:
		return fmt.Sprintf("SetPasswordAction(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e SetPasswordAction) MarshalJSON() ([]byte, error) {
	switch e {
	case SetPasswordActionKeep:
		return json.Marshal("keep")
	case SetPasswordActionFail:
		return json.Marshal("fail")
	case SetPasswordActionDisconnect:
		return json.Marshal("disconnect")
	default:
		return nil, fmt.Errorf("unknown enum value %q for SetPasswordAction", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *SetPasswordAction) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "keep":
		*e = SetPasswordActionKeep
	case "fail":
		*e = SetPasswordActionFail
	case "disconnect":
		*e = SetPasswordActionDisconnect
	default:
		return fmt.Errorf("unknown enum value %q for SetPasswordAction", s)
	}
	return nil
}

// SetPasswordOptions -> SetPasswordOptions (flat union)

// SetPasswordOptions implements the "SetPasswordOptions" QMP API type.
//
// Can be one of:
//   - SetPasswordOptionsVariantVNC
type SetPasswordOptions interface {
	isSetPasswordOptions()
}

// SetPasswordOptionsVariantVNC is an implementation of SetPasswordOptions.
type SetPasswordOptionsVariantVNC struct {
	Password  string             `json:"password"`
	Connected *SetPasswordAction `json:"connected,omitempty"`
	Display   *string            `json:"display,omitempty"`
}

func (SetPasswordOptionsVariantVNC) isSetPasswordOptions() {}

// MarshalJSON implements json.Marshaler.
func (s SetPasswordOptionsVariantVNC) MarshalJSON() ([]byte, error) {
	v := struct {
		Protocol DisplayProtocol `json:"protocol"`
		SetPasswordOptionsVariantVNC
	}{
		DisplayProtocolVNC,
		s,
	}
	return json.Marshal(v)
}

func decodeSetPasswordOptions(bs json.RawMessage) (SetPasswordOptions, error) {
	v := struct {
		Protocol DisplayProtocol `json:"protocol"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Protocol {
	case DisplayProtocolVNC:
		var ret SetPasswordOptionsVariantVNC
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union SetPasswordOptions", v.Protocol)
	}
}

// SevAttestationReport -> SevAttestationReport (struct)

// SevAttestationReport implements the "SevAttestationReport" QMP API type.
type SevAttestationReport struct {
	Data string `json:"data"`
}

// SevCapability -> SevCapability (struct)

// SevCapability implements the "SevCapability" QMP API type.
type SevCapability struct {
	Pdh             string `json:"pdh"`
	CertChain       string `json:"cert-chain"`
	CPU0ID          string `json:"cpu0-id"`
	Cbitpos         int64  `json:"cbitpos"`
	ReducedPhysBits int64  `json:"reduced-phys-bits"`
}

// SevInfo -> SevInfo (struct)

// SevInfo implements the "SevInfo" QMP API type.
type SevInfo struct {
	Enabled  bool     `json:"enabled"`
	ApiMajor uint8    `json:"api-major"`
	ApiMinor uint8    `json:"api-minor"`
	BuildID  uint8    `json:"build-id"`
	Policy   uint32   `json:"policy"`
	State    SevState `json:"state"`
	Handle   uint32   `json:"handle"`
}

// SevLaunchMeasureInfo -> SevLaunchMeasureInfo (struct)

// SevLaunchMeasureInfo implements the "SevLaunchMeasureInfo" QMP API type.
type SevLaunchMeasureInfo struct {
	Data string `json:"data"`
}

// SevState -> SevState (enum)

// SevState implements the "SevState" QMP API type.
type SevState int

// Known values of SevState.
const (
	SevStateUninit SevState = iota
	SevStateLaunchUpdate
	SevStateLaunchSecret
	SevStateRunning
	SevStateSendUpdate
	SevStateReceiveUpdate
)

// String implements fmt.Stringer.
func (e SevState) String() string {
	switch e {
	case SevStateUninit:
		return "uninit"
	case SevStateLaunchUpdate:
		return "launch-update"
	case SevStateLaunchSecret:
		return "launch-secret"
	case SevStateRunning:
		return "running"
	case SevStateSendUpdate:
		return "send-update"
	case SevStateReceiveUpdate:
		return "receive-update"
	default:
		return fmt.Sprintf("SevState(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e SevState) MarshalJSON() ([]byte, error) {
	switch e {
	case SevStateUninit:
		return json.Marshal("uninit")
	case SevStateLaunchUpdate:
		return json.Marshal("launch-update")
	case SevStateLaunchSecret:
		return json.Marshal("launch-secret")
	case SevStateRunning:
		return json.Marshal("running")
	case SevStateSendUpdate:
		return json.Marshal("send-update")
	case SevStateReceiveUpdate:
		return json.Marshal("receive-update")
	default:
		return nil, fmt.Errorf("unknown enum value %q for SevState", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *SevState) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "uninit":
		*e = SevStateUninit
	case "launch-update":
		*e = SevStateLaunchUpdate
	case "launch-secret":
		*e = SevStateLaunchSecret
	case "running":
		*e = SevStateRunning
	case "send-update":
		*e = SevStateSendUpdate
	case "receive-update":
		*e = SevStateReceiveUpdate
	default:
		return fmt.Errorf("unknown enum value %q for SevState", s)
	}
	return nil
}

// SgxEPCDeviceInfo -> SgxEpcDeviceInfo (struct)

// SgxEpcDeviceInfo implements the "SgxEPCDeviceInfo" QMP API type.
type SgxEpcDeviceInfo struct {
	ID      *string `json:"id,omitempty"`
	Memaddr uint64  `json:"memaddr"`
	Size    uint64  `json:"size"`
	Node    int64   `json:"node"`
	Memdev  string  `json:"memdev"`
}

// ShutdownAction -> ShutdownAction (enum)

// ShutdownAction implements the "ShutdownAction" QMP API type.
type ShutdownAction int

// Known values of ShutdownAction.
const (
	ShutdownActionPoweroff ShutdownAction = iota
	ShutdownActionPause
)

// String implements fmt.Stringer.
func (e ShutdownAction) String() string {
	switch e {
	case ShutdownActionPoweroff:
		return "poweroff"
	case ShutdownActionPause:
		return "pause"
	default:
		return fmt.Sprintf("ShutdownAction(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ShutdownAction) MarshalJSON() ([]byte, error) {
	switch e {
	case ShutdownActionPoweroff:
		return json.Marshal("poweroff")
	case ShutdownActionPause:
		return json.Marshal("pause")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ShutdownAction", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ShutdownAction) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "poweroff":
		*e = ShutdownActionPoweroff
	case "pause":
		*e = ShutdownActionPause
	default:
		return fmt.Errorf("unknown enum value %q for ShutdownAction", s)
	}
	return nil
}

// ShutdownCause -> ShutdownCause (enum)

// ShutdownCause implements the "ShutdownCause" QMP API type.
type ShutdownCause int

// Known values of ShutdownCause.
const (
	ShutdownCauseNone ShutdownCause = iota
	ShutdownCauseHostError
	ShutdownCauseHostQMPQuit
	ShutdownCauseHostQMPSystemReset
	ShutdownCauseHostSignal
	ShutdownCauseHostUi
	ShutdownCauseGuestShutdown
	ShutdownCauseGuestReset
	ShutdownCauseGuestPanic
	ShutdownCauseSubsystemReset
)

// String implements fmt.Stringer.
func (e ShutdownCause) String() string {
	switch e {
	case ShutdownCauseNone:
		return "none"
	case ShutdownCauseHostError:
		return "host-error"
	case ShutdownCauseHostQMPQuit:
		return "host-qmp-quit"
	case ShutdownCauseHostQMPSystemReset:
		return "host-qmp-system-reset"
	case ShutdownCauseHostSignal:
		return "host-signal"
	case ShutdownCauseHostUi:
		return "host-ui"
	case ShutdownCauseGuestShutdown:
		return "guest-shutdown"
	case ShutdownCauseGuestReset:
		return "guest-reset"
	case ShutdownCauseGuestPanic:
		return "guest-panic"
	case ShutdownCauseSubsystemReset:
		return "subsystem-reset"
	default:
		return fmt.Sprintf("ShutdownCause(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e ShutdownCause) MarshalJSON() ([]byte, error) {
	switch e {
	case ShutdownCauseNone:
		return json.Marshal("none")
	case ShutdownCauseHostError:
		return json.Marshal("host-error")
	case ShutdownCauseHostQMPQuit:
		return json.Marshal("host-qmp-quit")
	case ShutdownCauseHostQMPSystemReset:
		return json.Marshal("host-qmp-system-reset")
	case ShutdownCauseHostSignal:
		return json.Marshal("host-signal")
	case ShutdownCauseHostUi:
		return json.Marshal("host-ui")
	case ShutdownCauseGuestShutdown:
		return json.Marshal("guest-shutdown")
	case ShutdownCauseGuestReset:
		return json.Marshal("guest-reset")
	case ShutdownCauseGuestPanic:
		return json.Marshal("guest-panic")
	case ShutdownCauseSubsystemReset:
		return json.Marshal("subsystem-reset")
	default:
		return nil, fmt.Errorf("unknown enum value %q for ShutdownCause", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ShutdownCause) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = ShutdownCauseNone
	case "host-error":
		*e = ShutdownCauseHostError
	case "host-qmp-quit":
		*e = ShutdownCauseHostQMPQuit
	case "host-qmp-system-reset":
		*e = ShutdownCauseHostQMPSystemReset
	case "host-signal":
		*e = ShutdownCauseHostSignal
	case "host-ui":
		*e = ShutdownCauseHostUi
	case "guest-shutdown":
		*e = ShutdownCauseGuestShutdown
	case "guest-reset":
		*e = ShutdownCauseGuestReset
	case "guest-panic":
		*e = ShutdownCauseGuestPanic
	case "subsystem-reset":
		*e = ShutdownCauseSubsystemReset
	default:
		return fmt.Errorf("unknown enum value %q for ShutdownCause", s)
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
	Icount      *int64 `json:"icount,omitempty"`
}

// SocketAddress -> SocketAddress (flat union)

// SocketAddress implements the "SocketAddress" QMP API type.
//
// Can be one of:
//   - SocketAddressVariantFD
//   - SocketAddressVariantInet
//   - SocketAddressVariantUnix
//   - SocketAddressVariantVsock
type SocketAddress interface {
	isSocketAddress()
}

// SocketAddressVariantFD is an implementation of SocketAddress.
type SocketAddressVariantFD struct {
	Str string `json:"str"`
}

func (SocketAddressVariantFD) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressVariantFD) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressVariantFD
	}{
		SocketAddressTypeFD,
		s,
	}
	return json.Marshal(v)
}

// SocketAddressVariantInet is an implementation of SocketAddress.
type SocketAddressVariantInet struct {
	Numeric   *bool   `json:"numeric,omitempty"`
	To        *uint16 `json:"to,omitempty"`
	Ipv4      *bool   `json:"ipv4,omitempty"`
	Ipv6      *bool   `json:"ipv6,omitempty"`
	KeepAlive *bool   `json:"keep-alive,omitempty"`
	Mptcp     *bool   `json:"mptcp,omitempty"`
}

func (SocketAddressVariantInet) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressVariantInet) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressVariantInet
	}{
		SocketAddressTypeInet,
		s,
	}
	return json.Marshal(v)
}

// SocketAddressVariantUnix is an implementation of SocketAddress.
type SocketAddressVariantUnix struct {
	Path     string `json:"path"`
	Abstract *bool  `json:"abstract,omitempty"`
	Tight    *bool  `json:"tight,omitempty"`
}

func (SocketAddressVariantUnix) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressVariantUnix) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressVariantUnix
	}{
		SocketAddressTypeUnix,
		s,
	}
	return json.Marshal(v)
}

// SocketAddressVariantVsock is an implementation of SocketAddress.
type SocketAddressVariantVsock struct {
	Cid  string `json:"cid"`
	Port string `json:"port"`
}

func (SocketAddressVariantVsock) isSocketAddress() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressVariantVsock) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressVariantVsock
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
		var ret SocketAddressVariantFD
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SocketAddressTypeInet:
		var ret SocketAddressVariantInet
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SocketAddressTypeUnix:
		var ret SocketAddressVariantUnix
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SocketAddressTypeVsock:
		var ret SocketAddressVariantVsock
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union SocketAddress", v.Type)
	}
}

// SocketAddressLegacy -> SocketAddressLegacy (flat union)

// SocketAddressLegacy implements the "SocketAddressLegacy" QMP API type.
//
// Can be one of:
//   - SocketAddressLegacyVariantFD
//   - SocketAddressLegacyVariantInet
//   - SocketAddressLegacyVariantUnix
//   - SocketAddressLegacyVariantVsock
type SocketAddressLegacy interface {
	isSocketAddressLegacy()
}

// SocketAddressLegacyVariantFD is an implementation of SocketAddressLegacy.
type SocketAddressLegacyVariantFD struct {
	Data String `json:"data"`
}

func (SocketAddressLegacyVariantFD) isSocketAddressLegacy() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressLegacyVariantFD) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressLegacyVariantFD
	}{
		SocketAddressTypeFD,
		s,
	}
	return json.Marshal(v)
}

// SocketAddressLegacyVariantInet is an implementation of SocketAddressLegacy.
type SocketAddressLegacyVariantInet struct {
	Data InetSocketAddress `json:"data"`
}

func (SocketAddressLegacyVariantInet) isSocketAddressLegacy() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressLegacyVariantInet) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressLegacyVariantInet
	}{
		SocketAddressTypeInet,
		s,
	}
	return json.Marshal(v)
}

// SocketAddressLegacyVariantUnix is an implementation of SocketAddressLegacy.
type SocketAddressLegacyVariantUnix struct {
	Data UnixSocketAddress `json:"data"`
}

func (SocketAddressLegacyVariantUnix) isSocketAddressLegacy() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressLegacyVariantUnix) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressLegacyVariantUnix
	}{
		SocketAddressTypeUnix,
		s,
	}
	return json.Marshal(v)
}

// SocketAddressLegacyVariantVsock is an implementation of SocketAddressLegacy.
type SocketAddressLegacyVariantVsock struct {
	Data VsockSocketAddress `json:"data"`
}

func (SocketAddressLegacyVariantVsock) isSocketAddressLegacy() {}

// MarshalJSON implements json.Marshaler.
func (s SocketAddressLegacyVariantVsock) MarshalJSON() ([]byte, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
		SocketAddressLegacyVariantVsock
	}{
		SocketAddressTypeVsock,
		s,
	}
	return json.Marshal(v)
}

func decodeSocketAddressLegacy(bs json.RawMessage) (SocketAddressLegacy, error) {
	v := struct {
		Type SocketAddressType `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case SocketAddressTypeFD:
		var ret SocketAddressLegacyVariantFD
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SocketAddressTypeInet:
		var ret SocketAddressLegacyVariantInet
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SocketAddressTypeUnix:
		var ret SocketAddressLegacyVariantUnix
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case SocketAddressTypeVsock:
		var ret SocketAddressLegacyVariantVsock
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union SocketAddressLegacy", v.Type)
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

// SshHostKeyCheck -> SSHHostKeyCheck (flat union)

// SSHHostKeyCheck implements the "SshHostKeyCheck" QMP API type.
//
// Can be one of:
//   - SSHHostKeyCheckVariantHash
type SSHHostKeyCheck interface {
	isSSHHostKeyCheck()
}

// SSHHostKeyCheckVariantHash is an implementation of SSHHostKeyCheck.
type SSHHostKeyCheckVariantHash struct {
	Type SSHHostKeyCheckHashType `json:"type"`
	Hash string                  `json:"hash"`
}

func (SSHHostKeyCheckVariantHash) isSSHHostKeyCheck() {}

// MarshalJSON implements json.Marshaler.
func (s SSHHostKeyCheckVariantHash) MarshalJSON() ([]byte, error) {
	v := struct {
		Mode SSHHostKeyCheckMode `json:"mode"`
		SSHHostKeyCheckVariantHash
	}{
		SSHHostKeyCheckModeHash,
		s,
	}
	return json.Marshal(v)
}

func decodeSSHHostKeyCheck(bs json.RawMessage) (SSHHostKeyCheck, error) {
	v := struct {
		Mode SSHHostKeyCheckMode `json:"mode"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Mode {
	case SSHHostKeyCheckModeHash:
		var ret SSHHostKeyCheckVariantHash
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union SSHHostKeyCheck", v.Mode)
	}
}

// SshHostKeyCheckHashType -> SSHHostKeyCheckHashType (enum)

// SSHHostKeyCheckHashType implements the "SshHostKeyCheckHashType" QMP API type.
type SSHHostKeyCheckHashType int

// Known values of SSHHostKeyCheckHashType.
const (
	SSHHostKeyCheckHashTypeMd5 SSHHostKeyCheckHashType = iota
	SSHHostKeyCheckHashTypeSha1
	SSHHostKeyCheckHashTypeSha256
)

// String implements fmt.Stringer.
func (e SSHHostKeyCheckHashType) String() string {
	switch e {
	case SSHHostKeyCheckHashTypeMd5:
		return "md5"
	case SSHHostKeyCheckHashTypeSha1:
		return "sha1"
	case SSHHostKeyCheckHashTypeSha256:
		return "sha256"
	default:
		return fmt.Sprintf("SSHHostKeyCheckHashType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e SSHHostKeyCheckHashType) MarshalJSON() ([]byte, error) {
	switch e {
	case SSHHostKeyCheckHashTypeMd5:
		return json.Marshal("md5")
	case SSHHostKeyCheckHashTypeSha1:
		return json.Marshal("sha1")
	case SSHHostKeyCheckHashTypeSha256:
		return json.Marshal("sha256")
	default:
		return nil, fmt.Errorf("unknown enum value %q for SSHHostKeyCheckHashType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *SSHHostKeyCheckHashType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "md5":
		*e = SSHHostKeyCheckHashTypeMd5
	case "sha1":
		*e = SSHHostKeyCheckHashTypeSha1
	case "sha256":
		*e = SSHHostKeyCheckHashTypeSha256
	default:
		return fmt.Errorf("unknown enum value %q for SSHHostKeyCheckHashType", s)
	}
	return nil
}

// SshHostKeyCheckMode -> SSHHostKeyCheckMode (enum)

// SSHHostKeyCheckMode implements the "SshHostKeyCheckMode" QMP API type.
type SSHHostKeyCheckMode int

// Known values of SSHHostKeyCheckMode.
const (
	SSHHostKeyCheckModeNone SSHHostKeyCheckMode = iota
	SSHHostKeyCheckModeHash
	SSHHostKeyCheckModeKnownHosts
)

// String implements fmt.Stringer.
func (e SSHHostKeyCheckMode) String() string {
	switch e {
	case SSHHostKeyCheckModeNone:
		return "none"
	case SSHHostKeyCheckModeHash:
		return "hash"
	case SSHHostKeyCheckModeKnownHosts:
		return "known_hosts"
	default:
		return fmt.Sprintf("SSHHostKeyCheckMode(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e SSHHostKeyCheckMode) MarshalJSON() ([]byte, error) {
	switch e {
	case SSHHostKeyCheckModeNone:
		return json.Marshal("none")
	case SSHHostKeyCheckModeHash:
		return json.Marshal("hash")
	case SSHHostKeyCheckModeKnownHosts:
		return json.Marshal("known_hosts")
	default:
		return nil, fmt.Errorf("unknown enum value %q for SSHHostKeyCheckMode", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *SSHHostKeyCheckMode) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "none":
		*e = SSHHostKeyCheckModeNone
	case "hash":
		*e = SSHHostKeyCheckModeHash
	case "known_hosts":
		*e = SSHHostKeyCheckModeKnownHosts
	default:
		return fmt.Errorf("unknown enum value %q for SSHHostKeyCheckMode", s)
	}
	return nil
}

// Stats -> Stats (struct)

// Stats implements the "Stats" QMP API type.
type Stats struct {
	Name  string     `json:"name"`
	Value StatsValue `json:"value"`
}

// StatsFilter -> StatsFilter (flat union)

// StatsFilter implements the "StatsFilter" QMP API type.
//
// Can be one of:
//   - StatsFilterVariantVcpu
type StatsFilter interface {
	isStatsFilter()
}

// StatsFilterVariantVcpu is an implementation of StatsFilter.
type StatsFilterVariantVcpu struct {
	Providers []StatsRequest `json:"providers,omitempty"`
	Vcpus     []string       `json:"vcpus,omitempty"`
}

func (StatsFilterVariantVcpu) isStatsFilter() {}

// MarshalJSON implements json.Marshaler.
func (s StatsFilterVariantVcpu) MarshalJSON() ([]byte, error) {
	v := struct {
		Target StatsTarget `json:"target"`
		StatsFilterVariantVcpu
	}{
		StatsTargetVcpu,
		s,
	}
	return json.Marshal(v)
}

func decodeStatsFilter(bs json.RawMessage) (StatsFilter, error) {
	v := struct {
		Target StatsTarget `json:"target"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Target {
	case StatsTargetVcpu:
		var ret StatsFilterVariantVcpu
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union StatsFilter", v.Target)
	}
}

// StatsProvider -> StatsProvider (enum)

// StatsProvider implements the "StatsProvider" QMP API type.
type StatsProvider int

// Known values of StatsProvider.
const (
	StatsProviderKVM StatsProvider = iota
)

// String implements fmt.Stringer.
func (e StatsProvider) String() string {
	switch e {
	case StatsProviderKVM:
		return "kvm"
	default:
		return fmt.Sprintf("StatsProvider(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e StatsProvider) MarshalJSON() ([]byte, error) {
	switch e {
	case StatsProviderKVM:
		return json.Marshal("kvm")
	default:
		return nil, fmt.Errorf("unknown enum value %q for StatsProvider", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *StatsProvider) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "kvm":
		*e = StatsProviderKVM
	default:
		return fmt.Errorf("unknown enum value %q for StatsProvider", s)
	}
	return nil
}

// StatsRequest -> StatsRequest (struct)

// StatsRequest implements the "StatsRequest" QMP API type.
type StatsRequest struct {
	Provider StatsProvider `json:"provider"`
	Names    []string      `json:"names,omitempty"`
}

// StatsResult -> StatsResult (struct)

// StatsResult implements the "StatsResult" QMP API type.
type StatsResult struct {
	Provider StatsProvider `json:"provider"`
	QomPath  *string       `json:"qom-path,omitempty"`
	Stats    []Stats       `json:"stats"`
}

// StatsSchema -> StatsSchema (struct)

// StatsSchema implements the "StatsSchema" QMP API type.
type StatsSchema struct {
	Provider StatsProvider      `json:"provider"`
	Target   StatsTarget        `json:"target"`
	Stats    []StatsSchemaValue `json:"stats"`
}

// StatsSchemaValue -> StatsSchemaValue (struct)

// StatsSchemaValue implements the "StatsSchemaValue" QMP API type.
type StatsSchemaValue struct {
	Name       string     `json:"name"`
	Type       StatsType  `json:"type"`
	Unit       *StatsUnit `json:"unit,omitempty"`
	Base       *int8      `json:"base,omitempty"`
	Exponent   int16      `json:"exponent"`
	BucketSize *uint32    `json:"bucket-size,omitempty"`
}

// StatsTarget -> StatsTarget (enum)

// StatsTarget implements the "StatsTarget" QMP API type.
type StatsTarget int

// Known values of StatsTarget.
const (
	StatsTargetVM StatsTarget = iota
	StatsTargetVcpu
)

// String implements fmt.Stringer.
func (e StatsTarget) String() string {
	switch e {
	case StatsTargetVM:
		return "vm"
	case StatsTargetVcpu:
		return "vcpu"
	default:
		return fmt.Sprintf("StatsTarget(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e StatsTarget) MarshalJSON() ([]byte, error) {
	switch e {
	case StatsTargetVM:
		return json.Marshal("vm")
	case StatsTargetVcpu:
		return json.Marshal("vcpu")
	default:
		return nil, fmt.Errorf("unknown enum value %q for StatsTarget", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *StatsTarget) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "vm":
		*e = StatsTargetVM
	case "vcpu":
		*e = StatsTargetVcpu
	default:
		return fmt.Errorf("unknown enum value %q for StatsTarget", s)
	}
	return nil
}

// StatsType -> StatsType (enum)

// StatsType implements the "StatsType" QMP API type.
type StatsType int

// Known values of StatsType.
const (
	StatsTypeCumulative StatsType = iota
	StatsTypeInstant
	StatsTypePeak
	StatsTypeLinearHistogram
	StatsTypeLog2Histogram
)

// String implements fmt.Stringer.
func (e StatsType) String() string {
	switch e {
	case StatsTypeCumulative:
		return "cumulative"
	case StatsTypeInstant:
		return "instant"
	case StatsTypePeak:
		return "peak"
	case StatsTypeLinearHistogram:
		return "linear-histogram"
	case StatsTypeLog2Histogram:
		return "log2-histogram"
	default:
		return fmt.Sprintf("StatsType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e StatsType) MarshalJSON() ([]byte, error) {
	switch e {
	case StatsTypeCumulative:
		return json.Marshal("cumulative")
	case StatsTypeInstant:
		return json.Marshal("instant")
	case StatsTypePeak:
		return json.Marshal("peak")
	case StatsTypeLinearHistogram:
		return json.Marshal("linear-histogram")
	case StatsTypeLog2Histogram:
		return json.Marshal("log2-histogram")
	default:
		return nil, fmt.Errorf("unknown enum value %q for StatsType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *StatsType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "cumulative":
		*e = StatsTypeCumulative
	case "instant":
		*e = StatsTypeInstant
	case "peak":
		*e = StatsTypePeak
	case "linear-histogram":
		*e = StatsTypeLinearHistogram
	case "log2-histogram":
		*e = StatsTypeLog2Histogram
	default:
		return fmt.Errorf("unknown enum value %q for StatsType", s)
	}
	return nil
}

// StatsUnit -> StatsUnit (enum)

// StatsUnit implements the "StatsUnit" QMP API type.
type StatsUnit int

// Known values of StatsUnit.
const (
	StatsUnitBytes StatsUnit = iota
	StatsUnitSeconds
	StatsUnitCycles
	StatsUnitBoolean
)

// String implements fmt.Stringer.
func (e StatsUnit) String() string {
	switch e {
	case StatsUnitBytes:
		return "bytes"
	case StatsUnitSeconds:
		return "seconds"
	case StatsUnitCycles:
		return "cycles"
	case StatsUnitBoolean:
		return "boolean"
	default:
		return fmt.Sprintf("StatsUnit(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e StatsUnit) MarshalJSON() ([]byte, error) {
	switch e {
	case StatsUnitBytes:
		return json.Marshal("bytes")
	case StatsUnitSeconds:
		return json.Marshal("seconds")
	case StatsUnitCycles:
		return json.Marshal("cycles")
	case StatsUnitBoolean:
		return json.Marshal("boolean")
	default:
		return nil, fmt.Errorf("unknown enum value %q for StatsUnit", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *StatsUnit) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "bytes":
		*e = StatsUnitBytes
	case "seconds":
		*e = StatsUnitSeconds
	case "cycles":
		*e = StatsUnitCycles
	case "boolean":
		*e = StatsUnitBoolean
	default:
		return fmt.Errorf("unknown enum value %q for StatsUnit", s)
	}
	return nil
}

// StatsValue -> StatsValue (alternate)

// StatsValue implements the "StatsValue" QMP API type.
//
// Can be one of:
//   - StatsValueVariantBoolean
//   - StatsValueVariantList
//   - StatsValueVariantScalar
type StatsValue interface {
	isStatsValue()
}

// StatsValueVariantBoolean is an implementation of StatsValue
type StatsValueVariantBoolean bool

func (StatsValueVariantBoolean) isStatsValue() {}

// StatsValueVariantList is an implementation of StatsValue
type StatsValueVariantList uint64

func (StatsValueVariantList) isStatsValue() {}

// StatsValueVariantScalar is an implementation of StatsValue
type StatsValueVariantScalar uint64

func (StatsValueVariantScalar) isStatsValue() {}

func decodeStatsValue(bs json.RawMessage) (StatsValue, error) {

	var boolean StatsValueVariantBoolean
	if err := json.Unmarshal([]byte(bs), &boolean); err == nil {
		return boolean, nil
	}
	var list StatsValueVariantList
	if err := json.Unmarshal([]byte(bs), &list); err == nil {
		return list, nil
	}
	var scalar StatsValueVariantScalar
	if err := json.Unmarshal([]byte(bs), &scalar); err == nil {
		return scalar, nil
	}
	return nil, fmt.Errorf("unable to decode %q as a StatsValue", string(bs))
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
//   - StrOrNullVariantN
//   - StrOrNullVariantS
type StrOrNull interface {
	isStrOrNull()
}

// StrOrNullVariantN is a JSON null type, so it must
// also implement the isNullable interface.
type StrOrNullVariantN struct{}

func (StrOrNullVariantN) isNull() bool { return true }
func (StrOrNullVariantN) isStrOrNull() {}

// StrOrNullVariantS is an implementation of StrOrNull
type StrOrNullVariantS string

func (StrOrNullVariantS) isStrOrNull() {}

func decodeStrOrNull(bs json.RawMessage) (StrOrNull, error) {

	// Always try unmarshalling for nil first if it's an option
	// because other types could unmarshal successfully in the case
	// where a Null json type was provided.
	var n *int
	if err := json.Unmarshal([]byte(bs), &n); err == nil {
		if n == nil {
			return StrOrNullVariantN{}, nil
		}
	}
	var s StrOrNullVariantS
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

// SysEmuTarget -> SysEmuTarget (enum)

// SysEmuTarget implements the "SysEmuTarget" QMP API type.
type SysEmuTarget int

// Known values of SysEmuTarget.
const (
	SysEmuTargetAarch64 SysEmuTarget = iota
	SysEmuTargetAlpha
	SysEmuTargetArm
	SysEmuTargetAvr
	SysEmuTargetCris
	SysEmuTargetHppa
	SysEmuTargetI386
	SysEmuTargetLoongarch64
	SysEmuTargetM68K
	SysEmuTargetMicroblaze
	SysEmuTargetMicroblazeel
	SysEmuTargetMIPS
	SysEmuTargetMIPS64
	SysEmuTargetMIPS64El
	SysEmuTargetMipsel
	SysEmuTargetNios2
	SysEmuTargetOr1K
	SysEmuTargetPPC
	SysEmuTargetPPC64
	SysEmuTargetRiscv32
	SysEmuTargetRiscv64
	SysEmuTargetRx
	SysEmuTargetS390X
	SysEmuTargetSh4
	SysEmuTargetSh4Eb
	SysEmuTargetSPARC
	SysEmuTargetSPARC64
	SysEmuTargetTricore
	SysEmuTargetX8664
	SysEmuTargetXtensa
	SysEmuTargetXtensaeb
)

// String implements fmt.Stringer.
func (e SysEmuTarget) String() string {
	switch e {
	case SysEmuTargetAarch64:
		return "aarch64"
	case SysEmuTargetAlpha:
		return "alpha"
	case SysEmuTargetArm:
		return "arm"
	case SysEmuTargetAvr:
		return "avr"
	case SysEmuTargetCris:
		return "cris"
	case SysEmuTargetHppa:
		return "hppa"
	case SysEmuTargetI386:
		return "i386"
	case SysEmuTargetLoongarch64:
		return "loongarch64"
	case SysEmuTargetM68K:
		return "m68k"
	case SysEmuTargetMicroblaze:
		return "microblaze"
	case SysEmuTargetMicroblazeel:
		return "microblazeel"
	case SysEmuTargetMIPS:
		return "mips"
	case SysEmuTargetMIPS64:
		return "mips64"
	case SysEmuTargetMIPS64El:
		return "mips64el"
	case SysEmuTargetMipsel:
		return "mipsel"
	case SysEmuTargetNios2:
		return "nios2"
	case SysEmuTargetOr1K:
		return "or1k"
	case SysEmuTargetPPC:
		return "ppc"
	case SysEmuTargetPPC64:
		return "ppc64"
	case SysEmuTargetRiscv32:
		return "riscv32"
	case SysEmuTargetRiscv64:
		return "riscv64"
	case SysEmuTargetRx:
		return "rx"
	case SysEmuTargetS390X:
		return "s390x"
	case SysEmuTargetSh4:
		return "sh4"
	case SysEmuTargetSh4Eb:
		return "sh4eb"
	case SysEmuTargetSPARC:
		return "sparc"
	case SysEmuTargetSPARC64:
		return "sparc64"
	case SysEmuTargetTricore:
		return "tricore"
	case SysEmuTargetX8664:
		return "x86_64"
	case SysEmuTargetXtensa:
		return "xtensa"
	case SysEmuTargetXtensaeb:
		return "xtensaeb"
	default:
		return fmt.Sprintf("SysEmuTarget(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e SysEmuTarget) MarshalJSON() ([]byte, error) {
	switch e {
	case SysEmuTargetAarch64:
		return json.Marshal("aarch64")
	case SysEmuTargetAlpha:
		return json.Marshal("alpha")
	case SysEmuTargetArm:
		return json.Marshal("arm")
	case SysEmuTargetAvr:
		return json.Marshal("avr")
	case SysEmuTargetCris:
		return json.Marshal("cris")
	case SysEmuTargetHppa:
		return json.Marshal("hppa")
	case SysEmuTargetI386:
		return json.Marshal("i386")
	case SysEmuTargetLoongarch64:
		return json.Marshal("loongarch64")
	case SysEmuTargetM68K:
		return json.Marshal("m68k")
	case SysEmuTargetMicroblaze:
		return json.Marshal("microblaze")
	case SysEmuTargetMicroblazeel:
		return json.Marshal("microblazeel")
	case SysEmuTargetMIPS:
		return json.Marshal("mips")
	case SysEmuTargetMIPS64:
		return json.Marshal("mips64")
	case SysEmuTargetMIPS64El:
		return json.Marshal("mips64el")
	case SysEmuTargetMipsel:
		return json.Marshal("mipsel")
	case SysEmuTargetNios2:
		return json.Marshal("nios2")
	case SysEmuTargetOr1K:
		return json.Marshal("or1k")
	case SysEmuTargetPPC:
		return json.Marshal("ppc")
	case SysEmuTargetPPC64:
		return json.Marshal("ppc64")
	case SysEmuTargetRiscv32:
		return json.Marshal("riscv32")
	case SysEmuTargetRiscv64:
		return json.Marshal("riscv64")
	case SysEmuTargetRx:
		return json.Marshal("rx")
	case SysEmuTargetS390X:
		return json.Marshal("s390x")
	case SysEmuTargetSh4:
		return json.Marshal("sh4")
	case SysEmuTargetSh4Eb:
		return json.Marshal("sh4eb")
	case SysEmuTargetSPARC:
		return json.Marshal("sparc")
	case SysEmuTargetSPARC64:
		return json.Marshal("sparc64")
	case SysEmuTargetTricore:
		return json.Marshal("tricore")
	case SysEmuTargetX8664:
		return json.Marshal("x86_64")
	case SysEmuTargetXtensa:
		return json.Marshal("xtensa")
	case SysEmuTargetXtensaeb:
		return json.Marshal("xtensaeb")
	default:
		return nil, fmt.Errorf("unknown enum value %q for SysEmuTarget", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *SysEmuTarget) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "aarch64":
		*e = SysEmuTargetAarch64
	case "alpha":
		*e = SysEmuTargetAlpha
	case "arm":
		*e = SysEmuTargetArm
	case "avr":
		*e = SysEmuTargetAvr
	case "cris":
		*e = SysEmuTargetCris
	case "hppa":
		*e = SysEmuTargetHppa
	case "i386":
		*e = SysEmuTargetI386
	case "loongarch64":
		*e = SysEmuTargetLoongarch64
	case "m68k":
		*e = SysEmuTargetM68K
	case "microblaze":
		*e = SysEmuTargetMicroblaze
	case "microblazeel":
		*e = SysEmuTargetMicroblazeel
	case "mips":
		*e = SysEmuTargetMIPS
	case "mips64":
		*e = SysEmuTargetMIPS64
	case "mips64el":
		*e = SysEmuTargetMIPS64El
	case "mipsel":
		*e = SysEmuTargetMipsel
	case "nios2":
		*e = SysEmuTargetNios2
	case "or1k":
		*e = SysEmuTargetOr1K
	case "ppc":
		*e = SysEmuTargetPPC
	case "ppc64":
		*e = SysEmuTargetPPC64
	case "riscv32":
		*e = SysEmuTargetRiscv32
	case "riscv64":
		*e = SysEmuTargetRiscv64
	case "rx":
		*e = SysEmuTargetRx
	case "s390x":
		*e = SysEmuTargetS390X
	case "sh4":
		*e = SysEmuTargetSh4
	case "sh4eb":
		*e = SysEmuTargetSh4Eb
	case "sparc":
		*e = SysEmuTargetSPARC
	case "sparc64":
		*e = SysEmuTargetSPARC64
	case "tricore":
		*e = SysEmuTargetTricore
	case "x86_64":
		*e = SysEmuTargetX8664
	case "xtensa":
		*e = SysEmuTargetXtensa
	case "xtensaeb":
		*e = SysEmuTargetXtensaeb
	default:
		return fmt.Errorf("unknown enum value %q for SysEmuTarget", s)
	}
	return nil
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

// TPMPassthroughOptions -> TPMPassthroughOptions (struct)

// TPMPassthroughOptions implements the "TPMPassthroughOptions" QMP API type.
type TPMPassthroughOptions struct {
	Path       *string `json:"path,omitempty"`
	CancelPath *string `json:"cancel-path,omitempty"`
}

// TargetInfo -> TargetInfo (struct)

// TargetInfo implements the "TargetInfo" QMP API type.
type TargetInfo struct {
	Arch SysEmuTarget `json:"arch"`
}

// ThrottleLimits -> ThrottleLimits (struct)

// ThrottleLimits implements the "ThrottleLimits" QMP API type.
type ThrottleLimits struct {
	IopsTotal          *int64 `json:"iops-total,omitempty"`
	IopsTotalMax       *int64 `json:"iops-total-max,omitempty"`
	IopsTotalMaxLength *int64 `json:"iops-total-max-length,omitempty"`
	IopsRead           *int64 `json:"iops-read,omitempty"`
	IopsReadMax        *int64 `json:"iops-read-max,omitempty"`
	IopsReadMaxLength  *int64 `json:"iops-read-max-length,omitempty"`
	IopsWrite          *int64 `json:"iops-write,omitempty"`
	IopsWriteMax       *int64 `json:"iops-write-max,omitempty"`
	IopsWriteMaxLength *int64 `json:"iops-write-max-length,omitempty"`
	BpsTotal           *int64 `json:"bps-total,omitempty"`
	BpsTotalMax        *int64 `json:"bps-total-max,omitempty"`
	BpsTotalMaxLength  *int64 `json:"bps-total-max-length,omitempty"`
	BpsRead            *int64 `json:"bps-read,omitempty"`
	BpsReadMax         *int64 `json:"bps-read-max,omitempty"`
	BpsReadMaxLength   *int64 `json:"bps-read-max-length,omitempty"`
	BpsWrite           *int64 `json:"bps-write,omitempty"`
	BpsWriteMax        *int64 `json:"bps-write-max,omitempty"`
	BpsWriteMaxLength  *int64 `json:"bps-write-max-length,omitempty"`
	IopsSize           *int64 `json:"iops-size,omitempty"`
}

// TpmModel -> TPMModel (enum)

// TPMModel implements the "TpmModel" QMP API type.
type TPMModel int

// Known values of TPMModel.
const (
	TPMModelTPMTis TPMModel = iota
	TPMModelTPMCrb
	TPMModelTPMSpapr
)

// String implements fmt.Stringer.
func (e TPMModel) String() string {
	switch e {
	case TPMModelTPMTis:
		return "tpm-tis"
	case TPMModelTPMCrb:
		return "tpm-crb"
	case TPMModelTPMSpapr:
		return "tpm-spapr"
	default:
		return fmt.Sprintf("TPMModel(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e TPMModel) MarshalJSON() ([]byte, error) {
	switch e {
	case TPMModelTPMTis:
		return json.Marshal("tpm-tis")
	case TPMModelTPMCrb:
		return json.Marshal("tpm-crb")
	case TPMModelTPMSpapr:
		return json.Marshal("tpm-spapr")
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
	case "tpm-crb":
		*e = TPMModelTPMCrb
	case "tpm-spapr":
		*e = TPMModelTPMSpapr
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

// TpmTypeOptions -> TPMTypeOptions (flat union)

// TPMTypeOptions implements the "TpmTypeOptions" QMP API type.
//
// Can be one of:
//   - TPMTypeOptionsVariantEmulator
//   - TPMTypeOptionsVariantPassthrough
type TPMTypeOptions interface {
	isTPMTypeOptions()
}

// TPMTypeOptionsVariantEmulator is an implementation of TPMTypeOptions.
type TPMTypeOptionsVariantEmulator struct {
	Data TPMEmulatorOptions `json:"data"`
}

func (TPMTypeOptionsVariantEmulator) isTPMTypeOptions() {}

// MarshalJSON implements json.Marshaler.
func (s TPMTypeOptionsVariantEmulator) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TPMType `json:"type"`
		TPMTypeOptionsVariantEmulator
	}{
		TPMTypeEmulator,
		s,
	}
	return json.Marshal(v)
}

// TPMTypeOptionsVariantPassthrough is an implementation of TPMTypeOptions.
type TPMTypeOptionsVariantPassthrough struct {
	Data TPMPassthroughOptions `json:"data"`
}

func (TPMTypeOptionsVariantPassthrough) isTPMTypeOptions() {}

// MarshalJSON implements json.Marshaler.
func (s TPMTypeOptionsVariantPassthrough) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TPMType `json:"type"`
		TPMTypeOptionsVariantPassthrough
	}{
		TPMTypePassthrough,
		s,
	}
	return json.Marshal(v)
}

func decodeTPMTypeOptions(bs json.RawMessage) (TPMTypeOptions, error) {
	v := struct {
		Type TPMType `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case TPMTypeEmulator:
		var ret TPMTypeOptionsVariantEmulator
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TPMTypePassthrough:
		var ret TPMTypeOptionsVariantPassthrough
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union TPMTypeOptions", v.Type)
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

// TransactionAction -> TransactionAction (flat union)

// TransactionAction implements the "TransactionAction" QMP API type.
//
// Can be one of:
//   - TransactionActionVariantAbort
//   - TransactionActionVariantBlockDirtyBitmapAdd
//   - TransactionActionVariantBlockDirtyBitmapClear
//   - TransactionActionVariantBlockDirtyBitmapDisable
//   - TransactionActionVariantBlockDirtyBitmapEnable
//   - TransactionActionVariantBlockDirtyBitmapMerge
//   - TransactionActionVariantBlockDirtyBitmapRemove
//   - TransactionActionVariantBlockdevBackup
//   - TransactionActionVariantBlockdevSnapshot
//   - TransactionActionVariantBlockdevSnapshotInternalSync
//   - TransactionActionVariantBlockdevSnapshotSync
//   - TransactionActionVariantDriveBackup
type TransactionAction interface {
	isTransactionAction()
}

// TransactionActionVariantAbort is an implementation of TransactionAction.
type TransactionActionVariantAbort struct {
	Data Abort `json:"data"`
}

func (TransactionActionVariantAbort) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantAbort) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantAbort
	}{
		TransactionActionKindAbort,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantBlockDirtyBitmapAdd is an implementation of TransactionAction.
type TransactionActionVariantBlockDirtyBitmapAdd struct {
	Data BlockDirtyBitmapAdd `json:"data"`
}

func (TransactionActionVariantBlockDirtyBitmapAdd) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantBlockDirtyBitmapAdd) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantBlockDirtyBitmapAdd
	}{
		TransactionActionKindBlockDirtyBitmapAdd,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantBlockDirtyBitmapClear is an implementation of TransactionAction.
type TransactionActionVariantBlockDirtyBitmapClear struct {
	Data BlockDirtyBitmap `json:"data"`
}

func (TransactionActionVariantBlockDirtyBitmapClear) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantBlockDirtyBitmapClear) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantBlockDirtyBitmapClear
	}{
		TransactionActionKindBlockDirtyBitmapClear,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantBlockDirtyBitmapDisable is an implementation of TransactionAction.
type TransactionActionVariantBlockDirtyBitmapDisable struct {
	Data BlockDirtyBitmap `json:"data"`
}

func (TransactionActionVariantBlockDirtyBitmapDisable) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantBlockDirtyBitmapDisable) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantBlockDirtyBitmapDisable
	}{
		TransactionActionKindBlockDirtyBitmapDisable,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantBlockDirtyBitmapEnable is an implementation of TransactionAction.
type TransactionActionVariantBlockDirtyBitmapEnable struct {
	Data BlockDirtyBitmap `json:"data"`
}

func (TransactionActionVariantBlockDirtyBitmapEnable) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantBlockDirtyBitmapEnable) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantBlockDirtyBitmapEnable
	}{
		TransactionActionKindBlockDirtyBitmapEnable,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantBlockDirtyBitmapMerge is an implementation of TransactionAction.
type TransactionActionVariantBlockDirtyBitmapMerge struct {
	Data BlockDirtyBitmapMerge `json:"data"`
}

func (TransactionActionVariantBlockDirtyBitmapMerge) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantBlockDirtyBitmapMerge) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantBlockDirtyBitmapMerge
	}{
		TransactionActionKindBlockDirtyBitmapMerge,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantBlockDirtyBitmapRemove is an implementation of TransactionAction.
type TransactionActionVariantBlockDirtyBitmapRemove struct {
	Data BlockDirtyBitmap `json:"data"`
}

func (TransactionActionVariantBlockDirtyBitmapRemove) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantBlockDirtyBitmapRemove) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantBlockDirtyBitmapRemove
	}{
		TransactionActionKindBlockDirtyBitmapRemove,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantBlockdevBackup is an implementation of TransactionAction.
type TransactionActionVariantBlockdevBackup struct {
	Data BlockdevBackup `json:"data"`
}

func (TransactionActionVariantBlockdevBackup) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantBlockdevBackup) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantBlockdevBackup
	}{
		TransactionActionKindBlockdevBackup,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantBlockdevSnapshot is an implementation of TransactionAction.
type TransactionActionVariantBlockdevSnapshot struct {
	Data BlockdevSnapshot `json:"data"`
}

func (TransactionActionVariantBlockdevSnapshot) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantBlockdevSnapshot) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantBlockdevSnapshot
	}{
		TransactionActionKindBlockdevSnapshot,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantBlockdevSnapshotInternalSync is an implementation of TransactionAction.
type TransactionActionVariantBlockdevSnapshotInternalSync struct {
	Data BlockdevSnapshotInternal `json:"data"`
}

func (TransactionActionVariantBlockdevSnapshotInternalSync) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantBlockdevSnapshotInternalSync) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantBlockdevSnapshotInternalSync
	}{
		TransactionActionKindBlockdevSnapshotInternalSync,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantBlockdevSnapshotSync is an implementation of TransactionAction.
type TransactionActionVariantBlockdevSnapshotSync struct {
	Data BlockdevSnapshotSync `json:"data"`
}

func (TransactionActionVariantBlockdevSnapshotSync) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantBlockdevSnapshotSync) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantBlockdevSnapshotSync
	}{
		TransactionActionKindBlockdevSnapshotSync,
		s,
	}
	return json.Marshal(v)
}

// TransactionActionVariantDriveBackup is an implementation of TransactionAction.
type TransactionActionVariantDriveBackup struct {
	Data DriveBackup `json:"data"`
}

func (TransactionActionVariantDriveBackup) isTransactionAction() {}

// MarshalJSON implements json.Marshaler.
func (s TransactionActionVariantDriveBackup) MarshalJSON() ([]byte, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
		TransactionActionVariantDriveBackup
	}{
		TransactionActionKindDriveBackup,
		s,
	}
	return json.Marshal(v)
}

func decodeTransactionAction(bs json.RawMessage) (TransactionAction, error) {
	v := struct {
		Type TransactionActionKind `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case TransactionActionKindAbort:
		var ret TransactionActionVariantAbort
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindBlockDirtyBitmapAdd:
		var ret TransactionActionVariantBlockDirtyBitmapAdd
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindBlockDirtyBitmapClear:
		var ret TransactionActionVariantBlockDirtyBitmapClear
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindBlockDirtyBitmapDisable:
		var ret TransactionActionVariantBlockDirtyBitmapDisable
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindBlockDirtyBitmapEnable:
		var ret TransactionActionVariantBlockDirtyBitmapEnable
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindBlockDirtyBitmapMerge:
		var ret TransactionActionVariantBlockDirtyBitmapMerge
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindBlockDirtyBitmapRemove:
		var ret TransactionActionVariantBlockDirtyBitmapRemove
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindBlockdevBackup:
		var ret TransactionActionVariantBlockdevBackup
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindBlockdevSnapshot:
		var ret TransactionActionVariantBlockdevSnapshot
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindBlockdevSnapshotInternalSync:
		var ret TransactionActionVariantBlockdevSnapshotInternalSync
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindBlockdevSnapshotSync:
		var ret TransactionActionVariantBlockdevSnapshotSync
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case TransactionActionKindDriveBackup:
		var ret TransactionActionVariantDriveBackup
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union TransactionAction", v.Type)
	}
}

// TransactionActionKind -> TransactionActionKind (enum)

// TransactionActionKind implements the "TransactionActionKind" QMP API type.
type TransactionActionKind int

// Known values of TransactionActionKind.
const (
	TransactionActionKindAbort TransactionActionKind = iota
	TransactionActionKindBlockDirtyBitmapAdd
	TransactionActionKindBlockDirtyBitmapRemove
	TransactionActionKindBlockDirtyBitmapClear
	TransactionActionKindBlockDirtyBitmapEnable
	TransactionActionKindBlockDirtyBitmapDisable
	TransactionActionKindBlockDirtyBitmapMerge
	TransactionActionKindBlockdevBackup
	TransactionActionKindBlockdevSnapshot
	TransactionActionKindBlockdevSnapshotInternalSync
	TransactionActionKindBlockdevSnapshotSync
	TransactionActionKindDriveBackup
)

// String implements fmt.Stringer.
func (e TransactionActionKind) String() string {
	switch e {
	case TransactionActionKindAbort:
		return "abort"
	case TransactionActionKindBlockDirtyBitmapAdd:
		return "block-dirty-bitmap-add"
	case TransactionActionKindBlockDirtyBitmapRemove:
		return "block-dirty-bitmap-remove"
	case TransactionActionKindBlockDirtyBitmapClear:
		return "block-dirty-bitmap-clear"
	case TransactionActionKindBlockDirtyBitmapEnable:
		return "block-dirty-bitmap-enable"
	case TransactionActionKindBlockDirtyBitmapDisable:
		return "block-dirty-bitmap-disable"
	case TransactionActionKindBlockDirtyBitmapMerge:
		return "block-dirty-bitmap-merge"
	case TransactionActionKindBlockdevBackup:
		return "blockdev-backup"
	case TransactionActionKindBlockdevSnapshot:
		return "blockdev-snapshot"
	case TransactionActionKindBlockdevSnapshotInternalSync:
		return "blockdev-snapshot-internal-sync"
	case TransactionActionKindBlockdevSnapshotSync:
		return "blockdev-snapshot-sync"
	case TransactionActionKindDriveBackup:
		return "drive-backup"
	default:
		return fmt.Sprintf("TransactionActionKind(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e TransactionActionKind) MarshalJSON() ([]byte, error) {
	switch e {
	case TransactionActionKindAbort:
		return json.Marshal("abort")
	case TransactionActionKindBlockDirtyBitmapAdd:
		return json.Marshal("block-dirty-bitmap-add")
	case TransactionActionKindBlockDirtyBitmapRemove:
		return json.Marshal("block-dirty-bitmap-remove")
	case TransactionActionKindBlockDirtyBitmapClear:
		return json.Marshal("block-dirty-bitmap-clear")
	case TransactionActionKindBlockDirtyBitmapEnable:
		return json.Marshal("block-dirty-bitmap-enable")
	case TransactionActionKindBlockDirtyBitmapDisable:
		return json.Marshal("block-dirty-bitmap-disable")
	case TransactionActionKindBlockDirtyBitmapMerge:
		return json.Marshal("block-dirty-bitmap-merge")
	case TransactionActionKindBlockdevBackup:
		return json.Marshal("blockdev-backup")
	case TransactionActionKindBlockdevSnapshot:
		return json.Marshal("blockdev-snapshot")
	case TransactionActionKindBlockdevSnapshotInternalSync:
		return json.Marshal("blockdev-snapshot-internal-sync")
	case TransactionActionKindBlockdevSnapshotSync:
		return json.Marshal("blockdev-snapshot-sync")
	case TransactionActionKindDriveBackup:
		return json.Marshal("drive-backup")
	default:
		return nil, fmt.Errorf("unknown enum value %q for TransactionActionKind", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *TransactionActionKind) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "abort":
		*e = TransactionActionKindAbort
	case "block-dirty-bitmap-add":
		*e = TransactionActionKindBlockDirtyBitmapAdd
	case "block-dirty-bitmap-remove":
		*e = TransactionActionKindBlockDirtyBitmapRemove
	case "block-dirty-bitmap-clear":
		*e = TransactionActionKindBlockDirtyBitmapClear
	case "block-dirty-bitmap-enable":
		*e = TransactionActionKindBlockDirtyBitmapEnable
	case "block-dirty-bitmap-disable":
		*e = TransactionActionKindBlockDirtyBitmapDisable
	case "block-dirty-bitmap-merge":
		*e = TransactionActionKindBlockDirtyBitmapMerge
	case "blockdev-backup":
		*e = TransactionActionKindBlockdevBackup
	case "blockdev-snapshot":
		*e = TransactionActionKindBlockdevSnapshot
	case "blockdev-snapshot-internal-sync":
		*e = TransactionActionKindBlockdevSnapshotInternalSync
	case "blockdev-snapshot-sync":
		*e = TransactionActionKindBlockdevSnapshotSync
	case "drive-backup":
		*e = TransactionActionKindDriveBackup
	default:
		return fmt.Errorf("unknown enum value %q for TransactionActionKind", s)
	}
	return nil
}

// TransactionProperties -> TransactionProperties (struct)

// TransactionProperties implements the "TransactionProperties" QMP API type.
type TransactionProperties struct {
	CompletionMode *ActionCompletionMode `json:"completion-mode,omitempty"`
}

// EVENT UNPLUG_PRIMARY

// UnixSocketAddress -> UnixSocketAddress (struct)

// UnixSocketAddress implements the "UnixSocketAddress" QMP API type.
type UnixSocketAddress struct {
	Path     string `json:"path"`
	Abstract *bool  `json:"abstract,omitempty"`
	Tight    *bool  `json:"tight,omitempty"`
}

// UuidInfo -> UUIDInfo (struct)

// UUIDInfo implements the "UuidInfo" QMP API type.
type UUIDInfo struct {
	UUID string `json:"UUID"`
}

// EVENT VFU_CLIENT_HANGUP

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

// VfioStats -> VfioStats (struct)

// VfioStats implements the "VfioStats" QMP API type.
type VfioStats struct {
	Transferred int64 `json:"transferred"`
}

// VirtioMEMDeviceInfo -> VirtioMemDeviceInfo (struct)

// VirtioMemDeviceInfo implements the "VirtioMEMDeviceInfo" QMP API type.
type VirtioMemDeviceInfo struct {
	ID            *string `json:"id,omitempty"`
	Memaddr       uint64  `json:"memaddr"`
	RequestedSize uint64  `json:"requested-size"`
	Size          uint64  `json:"size"`
	MaxSize       uint64  `json:"max-size"`
	BlockSize     uint64  `json:"block-size"`
	Node          int64   `json:"node"`
	Memdev        string  `json:"memdev"`
}

// VirtioPMEMDeviceInfo -> VirtioPmemDeviceInfo (struct)

// VirtioPmemDeviceInfo implements the "VirtioPMEMDeviceInfo" QMP API type.
type VirtioPmemDeviceInfo struct {
	ID      *string `json:"id,omitempty"`
	Memaddr uint64  `json:"memaddr"`
	Size    uint64  `json:"size"`
	Memdev  string  `json:"memdev"`
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
	CacheSize     uint64  `json:"cache-size"`
	Bytes         int64   `json:"bytes"`
	Pages         int64   `json:"pages"`
	CacheMiss     int64   `json:"cache-miss"`
	CacheMissRate float64 `json:"cache-miss-rate"`
	EncodingRate  float64 `json:"encoding-rate"`
	Overflow      int64   `json:"overflow"`
}

// XDbgBlockGraph -> XDbgBlockGraph (struct)

// XDbgBlockGraph implements the "XDbgBlockGraph" QMP API type.
type XDbgBlockGraph struct {
	Nodes []XDbgBlockGraphNode `json:"nodes"`
	Edges []XDbgBlockGraphEdge `json:"edges"`
}

// XDbgBlockGraphEdge -> XDbgBlockGraphEdge (struct)

// XDbgBlockGraphEdge implements the "XDbgBlockGraphEdge" QMP API type.
type XDbgBlockGraphEdge struct {
	Parent     uint64            `json:"parent"`
	Child      uint64            `json:"child"`
	Name       string            `json:"name"`
	Perm       []BlockPermission `json:"perm"`
	SharedPerm []BlockPermission `json:"shared-perm"`
}

// XDbgBlockGraphNode -> XDbgBlockGraphNode (struct)

// XDbgBlockGraphNode implements the "XDbgBlockGraphNode" QMP API type.
type XDbgBlockGraphNode struct {
	ID   uint64                 `json:"id"`
	Type XDbgBlockGraphNodeType `json:"type"`
	Name string                 `json:"name"`
}

// XDbgBlockGraphNodeType -> XDbgBlockGraphNodeType (enum)

// XDbgBlockGraphNodeType implements the "XDbgBlockGraphNodeType" QMP API type.
type XDbgBlockGraphNodeType int

// Known values of XDbgBlockGraphNodeType.
const (
	XDbgBlockGraphNodeTypeBlockBackend XDbgBlockGraphNodeType = iota
	XDbgBlockGraphNodeTypeBlockJob
	XDbgBlockGraphNodeTypeBlockDriver
)

// String implements fmt.Stringer.
func (e XDbgBlockGraphNodeType) String() string {
	switch e {
	case XDbgBlockGraphNodeTypeBlockBackend:
		return "block-backend"
	case XDbgBlockGraphNodeTypeBlockJob:
		return "block-job"
	case XDbgBlockGraphNodeTypeBlockDriver:
		return "block-driver"
	default:
		return fmt.Sprintf("XDbgBlockGraphNodeType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e XDbgBlockGraphNodeType) MarshalJSON() ([]byte, error) {
	switch e {
	case XDbgBlockGraphNodeTypeBlockBackend:
		return json.Marshal("block-backend")
	case XDbgBlockGraphNodeTypeBlockJob:
		return json.Marshal("block-job")
	case XDbgBlockGraphNodeTypeBlockDriver:
		return json.Marshal("block-driver")
	default:
		return nil, fmt.Errorf("unknown enum value %q for XDbgBlockGraphNodeType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *XDbgBlockGraphNodeType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "block-backend":
		*e = XDbgBlockGraphNodeTypeBlockBackend
	case "block-job":
		*e = XDbgBlockGraphNodeTypeBlockJob
	case "block-driver":
		*e = XDbgBlockGraphNodeTypeBlockDriver
	default:
		return fmt.Errorf("unknown enum value %q for XDbgBlockGraphNodeType", s)
	}
	return nil
}

// YankInstance -> YankInstance (flat union)

// YankInstance implements the "YankInstance" QMP API type.
//
// Can be one of:
//   - YankInstanceVariantBlockNode
//   - YankInstanceVariantChardev
type YankInstance interface {
	isYankInstance()
}

// YankInstanceVariantBlockNode is an implementation of YankInstance.
type YankInstanceVariantBlockNode struct {
	NodeName string `json:"node-name"`
}

func (YankInstanceVariantBlockNode) isYankInstance() {}

// MarshalJSON implements json.Marshaler.
func (s YankInstanceVariantBlockNode) MarshalJSON() ([]byte, error) {
	v := struct {
		Type YankInstanceType `json:"type"`
		YankInstanceVariantBlockNode
	}{
		YankInstanceTypeBlockNode,
		s,
	}
	return json.Marshal(v)
}

// YankInstanceVariantChardev is an implementation of YankInstance.
type YankInstanceVariantChardev struct {
	ID string `json:"id"`
}

func (YankInstanceVariantChardev) isYankInstance() {}

// MarshalJSON implements json.Marshaler.
func (s YankInstanceVariantChardev) MarshalJSON() ([]byte, error) {
	v := struct {
		Type YankInstanceType `json:"type"`
		YankInstanceVariantChardev
	}{
		YankInstanceTypeChardev,
		s,
	}
	return json.Marshal(v)
}

func decodeYankInstance(bs json.RawMessage) (YankInstance, error) {
	v := struct {
		Type YankInstanceType `json:"type"`
	}{}
	if err := json.Unmarshal([]byte(bs), &v); err != nil {
		return nil, err
	}
	switch v.Type {
	case YankInstanceTypeBlockNode:
		var ret YankInstanceVariantBlockNode
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	case YankInstanceTypeChardev:
		var ret YankInstanceVariantChardev
		err := json.Unmarshal([]byte(bs), &ret)
		return ret, err
	default:
		return nil, fmt.Errorf("unknown flat union subtype %q for flat union YankInstance", v.Type)
	}
}

// YankInstanceType -> YankInstanceType (enum)

// YankInstanceType implements the "YankInstanceType" QMP API type.
type YankInstanceType int

// Known values of YankInstanceType.
const (
	YankInstanceTypeBlockNode YankInstanceType = iota
	YankInstanceTypeChardev
	YankInstanceTypeMigration
)

// String implements fmt.Stringer.
func (e YankInstanceType) String() string {
	switch e {
	case YankInstanceTypeBlockNode:
		return "block-node"
	case YankInstanceTypeChardev:
		return "chardev"
	case YankInstanceTypeMigration:
		return "migration"
	default:
		return fmt.Sprintf("YankInstanceType(%d)", e)
	}
}

// MarshalJSON implements json.Marshaler.
func (e YankInstanceType) MarshalJSON() ([]byte, error) {
	switch e {
	case YankInstanceTypeBlockNode:
		return json.Marshal("block-node")
	case YankInstanceTypeChardev:
		return json.Marshal("chardev")
	case YankInstanceTypeMigration:
		return json.Marshal("migration")
	default:
		return nil, fmt.Errorf("unknown enum value %q for YankInstanceType", e)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *YankInstanceType) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	switch s {
	case "block-node":
		*e = YankInstanceTypeBlockNode
	case "chardev":
		*e = YankInstanceTypeChardev
	case "migration":
		*e = YankInstanceTypeMigration
	default:
		return fmt.Errorf("unknown enum value %q for YankInstanceType", s)
	}
	return nil
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

// announce-self -> AnnounceSelf (command)

// AnnounceSelf implements the "announce-self" QMP API call.
func (m *Monitor) AnnounceSelf(cmd *AnnounceParameters) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "announce-self",
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
func (m *Monitor) BlockCommit(jobID *string, device string, baseNode *string, base *string, topNode *string, top *string, backingFile *string, speed *int64, onError *BlockdevOnError, filterNodeName *string, autoFinalize *bool, autoDismiss *bool) (err error) {
	cmd := struct {
		JobID          *string          `json:"job-id,omitempty"`
		Device         string           `json:"device"`
		BaseNode       *string          `json:"base-node,omitempty"`
		Base           *string          `json:"base,omitempty"`
		TopNode        *string          `json:"top-node,omitempty"`
		Top            *string          `json:"top,omitempty"`
		BackingFile    *string          `json:"backing-file,omitempty"`
		Speed          *int64           `json:"speed,omitempty"`
		OnError        *BlockdevOnError `json:"on-error,omitempty"`
		FilterNodeName *string          `json:"filter-node-name,omitempty"`
		AutoFinalize   *bool            `json:"auto-finalize,omitempty"`
		AutoDismiss    *bool            `json:"auto-dismiss,omitempty"`
	}{
		jobID,
		device,
		baseNode,
		base,
		topNode,
		top,
		backingFile,
		speed,
		onError,
		filterNodeName,
		autoFinalize,
		autoDismiss,
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
func (m *Monitor) BlockDirtyBitmapAdd(node string, name string, granularity *uint32, persistent *bool, disabled *bool) (err error) {
	cmd := struct {
		Node        string  `json:"node"`
		Name        string  `json:"name"`
		Granularity *uint32 `json:"granularity,omitempty"`
		Persistent  *bool   `json:"persistent,omitempty"`
		Disabled    *bool   `json:"disabled,omitempty"`
	}{
		node,
		name,
		granularity,
		persistent,
		disabled,
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

// block-dirty-bitmap-disable -> BlockDirtyBitmapDisable (command)

// BlockDirtyBitmapDisable implements the "block-dirty-bitmap-disable" QMP API call.
func (m *Monitor) BlockDirtyBitmapDisable(node string, name string) (err error) {
	cmd := struct {
		Node string `json:"node"`
		Name string `json:"name"`
	}{
		node,
		name,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-dirty-bitmap-disable",
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

// block-dirty-bitmap-enable -> BlockDirtyBitmapEnable (command)

// BlockDirtyBitmapEnable implements the "block-dirty-bitmap-enable" QMP API call.
func (m *Monitor) BlockDirtyBitmapEnable(node string, name string) (err error) {
	cmd := struct {
		Node string `json:"node"`
		Name string `json:"name"`
	}{
		node,
		name,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-dirty-bitmap-enable",
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

// block-dirty-bitmap-merge -> BlockDirtyBitmapMerge (command)

// BlockDirtyBitmapMerge implements the "block-dirty-bitmap-merge" QMP API call.
func (m *Monitor) BlockDirtyBitmapMerge(node string, target string, bitmaps []BlockDirtyBitmapOrStr) (err error) {
	cmd := struct {
		Node    string                  `json:"node"`
		Target  string                  `json:"target"`
		Bitmaps []BlockDirtyBitmapOrStr `json:"bitmaps"`
	}{
		node,
		target,
		bitmaps,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-dirty-bitmap-merge",
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

// block-export-add -> BlockExportAdd (command)

// BlockExportAdd implements the "block-export-add" QMP API call.
func (m *Monitor) BlockExportAdd(cmd *BlockExportOptions) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-export-add",
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

// block-export-del -> BlockExportDel (command)

// BlockExportDel implements the "block-export-del" QMP API call.
func (m *Monitor) BlockExportDel(id string, mode *BlockExportRemoveMode) (err error) {
	cmd := struct {
		ID   string                 `json:"id"`
		Mode *BlockExportRemoveMode `json:"mode,omitempty"`
	}{
		id,
		mode,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-export-del",
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

// block-job-dismiss -> BlockJobDismiss (command)

// BlockJobDismiss implements the "block-job-dismiss" QMP API call.
func (m *Monitor) BlockJobDismiss(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-job-dismiss",
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

// block-job-finalize -> BlockJobFinalize (command)

// BlockJobFinalize implements the "block-job-finalize" QMP API call.
func (m *Monitor) BlockJobFinalize(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-job-finalize",
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

// block-latency-histogram-set -> BlockLatencyHistogramSet (command)

// BlockLatencyHistogramSet implements the "block-latency-histogram-set" QMP API call.
func (m *Monitor) BlockLatencyHistogramSet(id string, boundaries []uint64, boundariesRead []uint64, boundariesWrite []uint64, boundariesFlush []uint64) (err error) {
	cmd := struct {
		ID              string   `json:"id"`
		Boundaries      []uint64 `json:"boundaries,omitempty"`
		BoundariesRead  []uint64 `json:"boundaries-read,omitempty"`
		BoundariesWrite []uint64 `json:"boundaries-write,omitempty"`
		BoundariesFlush []uint64 `json:"boundaries-flush,omitempty"`
	}{
		id,
		boundaries,
		boundariesRead,
		boundariesWrite,
		boundariesFlush,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "block-latency-histogram-set",
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
func (m *Monitor) BlockStream(jobID *string, device string, base *string, baseNode *string, backingFile *string, bottom *string, speed *int64, onError *BlockdevOnError, filterNodeName *string, autoFinalize *bool, autoDismiss *bool) (err error) {
	cmd := struct {
		JobID          *string          `json:"job-id,omitempty"`
		Device         string           `json:"device"`
		Base           *string          `json:"base,omitempty"`
		BaseNode       *string          `json:"base-node,omitempty"`
		BackingFile    *string          `json:"backing-file,omitempty"`
		Bottom         *string          `json:"bottom,omitempty"`
		Speed          *int64           `json:"speed,omitempty"`
		OnError        *BlockdevOnError `json:"on-error,omitempty"`
		FilterNodeName *string          `json:"filter-node-name,omitempty"`
		AutoFinalize   *bool            `json:"auto-finalize,omitempty"`
		AutoDismiss    *bool            `json:"auto-dismiss,omitempty"`
	}{
		jobID,
		device,
		base,
		baseNode,
		backingFile,
		bottom,
		speed,
		onError,
		filterNodeName,
		autoFinalize,
		autoDismiss,
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
func (m *Monitor) BlockdevChangeMedium(device *string, id *string, filename string, format *string, force *bool, readOnlyMode *BlockdevChangeReadOnlyMode) (err error) {
	cmd := struct {
		Device       *string                     `json:"device,omitempty"`
		ID           *string                     `json:"id,omitempty"`
		Filename     string                      `json:"filename"`
		Format       *string                     `json:"format,omitempty"`
		Force        *bool                       `json:"force,omitempty"`
		ReadOnlyMode *BlockdevChangeReadOnlyMode `json:"read-only-mode,omitempty"`
	}{
		device,
		id,
		filename,
		format,
		force,
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

// blockdev-create -> BlockdevCreate (command)

// BlockdevCreate implements the "blockdev-create" QMP API call.
func (m *Monitor) BlockdevCreate(jobID string, options BlockdevCreateOptions) (err error) {
	cmd := struct {
		JobID   string                `json:"job-id"`
		Options BlockdevCreateOptions `json:"options"`
	}{
		jobID,
		options,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-create",
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

// blockdev-insert-medium -> BlockdevInsertMedium (command)

// BlockdevInsertMedium implements the "blockdev-insert-medium" QMP API call.
func (m *Monitor) BlockdevInsertMedium(id string, nodeName string) (err error) {
	cmd := struct {
		ID       string `json:"id"`
		NodeName string `json:"node-name"`
	}{
		id,
		nodeName,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-insert-medium",
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
func (m *Monitor) BlockdevMirror(jobID *string, device string, target string, replaces *string, sync MirrorSyncMode, speed *int64, granularity *uint32, bufSize *int64, onSourceError *BlockdevOnError, onTargetError *BlockdevOnError, filterNodeName *string, copyMode *MirrorCopyMode, autoFinalize *bool, autoDismiss *bool) (err error) {
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
		CopyMode       *MirrorCopyMode  `json:"copy-mode,omitempty"`
		AutoFinalize   *bool            `json:"auto-finalize,omitempty"`
		AutoDismiss    *bool            `json:"auto-dismiss,omitempty"`
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
		copyMode,
		autoFinalize,
		autoDismiss,
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

// blockdev-remove-medium -> BlockdevRemoveMedium (command)

// BlockdevRemoveMedium implements the "blockdev-remove-medium" QMP API call.
func (m *Monitor) BlockdevRemoveMedium(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-remove-medium",
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

// blockdev-reopen -> BlockdevReopen (command)

// BlockdevReopen implements the "blockdev-reopen" QMP API call.
func (m *Monitor) BlockdevReopen(options []BlockdevOptions) (err error) {
	cmd := struct {
		Options []BlockdevOptions `json:"options"`
	}{
		options,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "blockdev-reopen",
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

// calc-dirty-rate -> CalcDirtyRate (command)

// CalcDirtyRate implements the "calc-dirty-rate" QMP API call.
func (m *Monitor) CalcDirtyRate(calcTime int64, samplePages *int64, mode *DirtyRateMeasureMode) (err error) {
	cmd := struct {
		CalcTime    int64                 `json:"calc-time"`
		SamplePages *int64                `json:"sample-pages,omitempty"`
		Mode        *DirtyRateMeasureMode `json:"mode,omitempty"`
	}{
		calcTime,
		samplePages,
		mode,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "calc-dirty-rate",
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

// cancel-vcpu-dirty-limit -> CancelVcpuDirtyLimit (command)

// CancelVcpuDirtyLimit implements the "cancel-vcpu-dirty-limit" QMP API call.
func (m *Monitor) CancelVcpuDirtyLimit(cpuIndex *int64) (err error) {
	cmd := struct {
		CPUIndex *int64 `json:"cpu-index,omitempty"`
	}{
		cpuIndex,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "cancel-vcpu-dirty-limit",
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

// device-list-properties -> DeviceListProperties (command)

// DeviceListProperties implements the "device-list-properties" QMP API call.
func (m *Monitor) DeviceListProperties(typename string) (ret []ObjectPropertyInfo, err error) {
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

// display-reload -> DisplayReload (command)

// DisplayReload implements the "display-reload" QMP API call.
func (m *Monitor) DisplayReload(cmd *DisplayReloadOptions) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "display-reload",
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

// display-update -> DisplayUpdate (command)

// DisplayUpdate implements the "display-update" QMP API call.
func (m *Monitor) DisplayUpdate(cmd *DisplayUpdateOptions) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "display-update",
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
func (m *Monitor) ExpirePassword(cmd *ExpirePasswordOptions) (err error) {
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

// job-cancel -> JobCancel (command)

// JobCancel implements the "job-cancel" QMP API call.
func (m *Monitor) JobCancel(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "job-cancel",
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

// job-complete -> JobComplete (command)

// JobComplete implements the "job-complete" QMP API call.
func (m *Monitor) JobComplete(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "job-complete",
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

// job-dismiss -> JobDismiss (command)

// JobDismiss implements the "job-dismiss" QMP API call.
func (m *Monitor) JobDismiss(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "job-dismiss",
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

// job-finalize -> JobFinalize (command)

// JobFinalize implements the "job-finalize" QMP API call.
func (m *Monitor) JobFinalize(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "job-finalize",
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

// job-pause -> JobPause (command)

// JobPause implements the "job-pause" QMP API call.
func (m *Monitor) JobPause(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "job-pause",
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

// job-resume -> JobResume (command)

// JobResume implements the "job-resume" QMP API call.
func (m *Monitor) JobResume(id string) (err error) {
	cmd := struct {
		ID string `json:"id"`
	}{
		id,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "job-resume",
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
func (m *Monitor) Migrate(uri string, blk *bool, inc *bool, detach *bool, resume *bool) (err error) {
	cmd := struct {
		URI    string `json:"uri"`
		Blk    *bool  `json:"blk,omitempty"`
		Inc    *bool  `json:"inc,omitempty"`
		Detach *bool  `json:"detach,omitempty"`
		Resume *bool  `json:"resume,omitempty"`
	}{
		uri,
		blk,
		inc,
		detach,
		resume,
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

// migrate-pause -> MigratePause (command)

// MigratePause implements the "migrate-pause" QMP API call.
func (m *Monitor) MigratePause() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate-pause",
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

// migrate-recover -> MigrateRecover (command)

// MigrateRecover implements the "migrate-recover" QMP API call.
func (m *Monitor) MigrateRecover(uri string) (err error) {
	cmd := struct {
		URI string `json:"uri"`
	}{
		uri,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "migrate-recover",
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

// nbd-server-add -> NBDServerAdd (command)

// NBDServerAdd implements the "nbd-server-add" QMP API call.
func (m *Monitor) NBDServerAdd(cmd *NBDServerAddOptions) (err error) {
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

// nbd-server-remove -> NBDServerRemove (command)

// NBDServerRemove implements the "nbd-server-remove" QMP API call.
func (m *Monitor) NBDServerRemove(name string, mode *BlockExportRemoveMode) (err error) {
	cmd := struct {
		Name string                 `json:"name"`
		Mode *BlockExportRemoveMode `json:"mode,omitempty"`
	}{
		name,
		mode,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "nbd-server-remove",
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
func (m *Monitor) NBDServerStart(addr SocketAddressLegacy, tlsCreds *string, tlsAuthz *string, maxConnections *uint32) (err error) {
	cmd := struct {
		Addr           SocketAddressLegacy `json:"addr"`
		TLSCreds       *string             `json:"tls-creds,omitempty"`
		TLSAuthz       *string             `json:"tls-authz,omitempty"`
		MaxConnections *uint32             `json:"max-connections,omitempty"`
	}{
		addr,
		tlsCreds,
		tlsAuthz,
		maxConnections,
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
func (m *Monitor) NetdevAdd(cmd *Netdev) (err error) {
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
func (m *Monitor) ObjectAdd(cmd *ObjectOptions) (err error) {
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
func (m *Monitor) QMPCapabilities(enable []QMPCapability) (err error) {
	cmd := struct {
		Enable []QMPCapability `json:"enable,omitempty"`
	}{
		enable,
	}
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

// qom-list-properties -> QomListProperties (command)

// QomListProperties implements the "qom-list-properties" QMP API call.
func (m *Monitor) QomListProperties(typename string) (ret []ObjectPropertyInfo, err error) {
	cmd := struct {
		Typename string `json:"typename"`
	}{
		typename,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "qom-list-properties",
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

// query-block-exports -> QueryBlockExports (command)

// QueryBlockExports implements the "query-block-exports" QMP API call.
func (m *Monitor) QueryBlockExports() (ret []BlockExportInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-block-exports",
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

// query-colo-status -> QueryColoStatus (command)

// QueryColoStatus implements the "query-colo-status" QMP API call.
func (m *Monitor) QueryColoStatus() (ret ColoStatus, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-colo-status",
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

// query-cpus-fast -> QueryCpusFast (command)

// QueryCpusFast implements the "query-cpus-fast" QMP API call.
func (m *Monitor) QueryCpusFast() (ret []CPUInfoFast, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-cpus-fast",
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

// query-current-machine -> QueryCurrentMachine (command)

// QueryCurrentMachine implements the "query-current-machine" QMP API call.
func (m *Monitor) QueryCurrentMachine() (ret CurrentMachineParams, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-current-machine",
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

// query-dirty-rate -> QueryDirtyRate (command)

// QueryDirtyRate implements the "query-dirty-rate" QMP API call.
func (m *Monitor) QueryDirtyRate() (ret DirtyRateInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-dirty-rate",
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

// query-display-options -> QueryDisplayOptions (command)

// QueryDisplayOptions implements the "query-display-options" QMP API call.
func (m *Monitor) QueryDisplayOptions() (ret DisplayOptions, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-display-options",
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

// query-jobs -> QueryJobs (command)

// QueryJobs implements the "query-jobs" QMP API call.
func (m *Monitor) QueryJobs() (ret []JobInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-jobs",
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
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
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
func (m *Monitor) QueryNamedBlockNodes(flat *bool) (ret []BlockDeviceInfo, err error) {
	cmd := struct {
		Flat *bool `json:"flat,omitempty"`
	}{
		flat,
	}
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

// query-pr-managers -> QueryPrManagers (command)

// QueryPrManagers implements the "query-pr-managers" QMP API call.
func (m *Monitor) QueryPrManagers() (ret []PrManagerInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-pr-managers",
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
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}

// query-replay -> QueryReplay (command)

// QueryReplay implements the "query-replay" QMP API call.
func (m *Monitor) QueryReplay() (ret ReplayInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-replay",
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

// query-sev -> QuerySev (command)

// QuerySev implements the "query-sev" QMP API call.
func (m *Monitor) QuerySev() (ret SevInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-sev",
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

// query-sev-attestation-report -> QuerySevAttestationReport (command)

// QuerySevAttestationReport implements the "query-sev-attestation-report" QMP API call.
func (m *Monitor) QuerySevAttestationReport(mnonce string) (ret SevAttestationReport, err error) {
	cmd := struct {
		Mnonce string `json:"mnonce"`
	}{
		mnonce,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-sev-attestation-report",
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

// query-sev-capabilities -> QuerySevCapabilities (command)

// QuerySevCapabilities implements the "query-sev-capabilities" QMP API call.
func (m *Monitor) QuerySevCapabilities() (ret SevCapability, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-sev-capabilities",
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

// query-sev-launch-measure -> QuerySevLaunchMeasure (command)

// QuerySevLaunchMeasure implements the "query-sev-launch-measure" QMP API call.
func (m *Monitor) QuerySevLaunchMeasure() (ret SevLaunchMeasureInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-sev-launch-measure",
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

// query-sgx -> QuerySgx (command)

// QuerySgx implements the "query-sgx" QMP API call.
func (m *Monitor) QuerySgx() (ret SgxInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-sgx",
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

// query-sgx-capabilities -> QuerySgxCapabilities (command)

// QuerySgxCapabilities implements the "query-sgx-capabilities" QMP API call.
func (m *Monitor) QuerySgxCapabilities() (ret SgxInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-sgx-capabilities",
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

// query-stats -> QueryStats (command)

// QueryStats implements the "query-stats" QMP API call.
func (m *Monitor) QueryStats(cmd *StatsFilter) (ret []StatsResult, err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-stats",
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

// query-stats-schemas -> QueryStatsSchemas (command)

// QueryStatsSchemas implements the "query-stats-schemas" QMP API call.
func (m *Monitor) QueryStatsSchemas(provider *StatsProvider) (ret []StatsSchema, err error) {
	cmd := struct {
		Provider *StatsProvider `json:"provider,omitempty"`
	}{
		provider,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-stats-schemas",
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

// query-vcpu-dirty-limit -> QueryVcpuDirtyLimit (command)

// QueryVcpuDirtyLimit implements the "query-vcpu-dirty-limit" QMP API call.
func (m *Monitor) QueryVcpuDirtyLimit() (ret []DirtyLimitInfo, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-vcpu-dirty-limit",
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

// query-yank -> QueryYank (command)

// QueryYank implements the "query-yank" QMP API call.
func (m *Monitor) QueryYank() (ret []YankInstance, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "query-yank",
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

// replay-break -> ReplayBreak (command)

// ReplayBreak implements the "replay-break" QMP API call.
func (m *Monitor) ReplayBreak(icount int64) (err error) {
	cmd := struct {
		Icount int64 `json:"icount"`
	}{
		icount,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "replay-break",
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

// replay-delete-break -> ReplayDeleteBreak (command)

// ReplayDeleteBreak implements the "replay-delete-break" QMP API call.
func (m *Monitor) ReplayDeleteBreak() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "replay-delete-break",
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

// replay-seek -> ReplaySeek (command)

// ReplaySeek implements the "replay-seek" QMP API call.
func (m *Monitor) ReplaySeek(icount int64) (err error) {
	cmd := struct {
		Icount int64 `json:"icount"`
	}{
		icount,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "replay-seek",
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
func (m *Monitor) Screendump(filename string, device *string, head *int64, format *ImageFormat) (err error) {
	cmd := struct {
		Filename string       `json:"filename"`
		Device   *string      `json:"device,omitempty"`
		Head     *int64       `json:"head,omitempty"`
		Format   *ImageFormat `json:"format,omitempty"`
	}{
		filename,
		device,
		head,
		format,
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

// set-action -> SetAction (command)

// SetAction implements the "set-action" QMP API call.
func (m *Monitor) SetAction(reboot *RebootAction, shutdown *ShutdownAction, panic *PanicAction, watchdog *WatchdogAction) (err error) {
	cmd := struct {
		Reboot   *RebootAction   `json:"reboot,omitempty"`
		Shutdown *ShutdownAction `json:"shutdown,omitempty"`
		Panic    *PanicAction    `json:"panic,omitempty"`
		Watchdog *WatchdogAction `json:"watchdog,omitempty"`
	}{
		reboot,
		shutdown,
		panic,
		watchdog,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "set-action",
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

// set-numa-node -> SetNumaNode (command)

// SetNumaNode implements the "set-numa-node" QMP API call.
func (m *Monitor) SetNumaNode(cmd *NumaOptions) (err error) {
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "set-numa-node",
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

// set-vcpu-dirty-limit -> SetVcpuDirtyLimit (command)

// SetVcpuDirtyLimit implements the "set-vcpu-dirty-limit" QMP API call.
func (m *Monitor) SetVcpuDirtyLimit(cpuIndex *int64, dirtyRate uint64) (err error) {
	cmd := struct {
		CPUIndex  *int64 `json:"cpu-index,omitempty"`
		DirtyRate uint64 `json:"dirty-rate"`
	}{
		cpuIndex,
		dirtyRate,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "set-vcpu-dirty-limit",
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
func (m *Monitor) SetPassword(cmd *SetPasswordOptions) (err error) {
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

// sev-inject-launch-secret -> SevInjectLaunchSecret (command)

// SevInjectLaunchSecret implements the "sev-inject-launch-secret" QMP API call.
func (m *Monitor) SevInjectLaunchSecret(packetHeader string, secret string, gpa *uint64) (err error) {
	cmd := struct {
		PacketHeader string  `json:"packet-header"`
		Secret       string  `json:"secret"`
		Gpa          *uint64 `json:"gpa,omitempty"`
	}{
		packetHeader,
		secret,
		gpa,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "sev-inject-launch-secret",
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

// snapshot-delete -> SnapshotDelete (command)

// SnapshotDelete implements the "snapshot-delete" QMP API call.
func (m *Monitor) SnapshotDelete(jobID string, tag string, devices []string) (err error) {
	cmd := struct {
		JobID   string   `json:"job-id"`
		Tag     string   `json:"tag"`
		Devices []string `json:"devices"`
	}{
		jobID,
		tag,
		devices,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "snapshot-delete",
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

// snapshot-load -> SnapshotLoad (command)

// SnapshotLoad implements the "snapshot-load" QMP API call.
func (m *Monitor) SnapshotLoad(jobID string, tag string, vmstate string, devices []string) (err error) {
	cmd := struct {
		JobID   string   `json:"job-id"`
		Tag     string   `json:"tag"`
		Vmstate string   `json:"vmstate"`
		Devices []string `json:"devices"`
	}{
		jobID,
		tag,
		vmstate,
		devices,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "snapshot-load",
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

// snapshot-save -> SnapshotSave (command)

// SnapshotSave implements the "snapshot-save" QMP API call.
func (m *Monitor) SnapshotSave(jobID string, tag string, vmstate string, devices []string) (err error) {
	cmd := struct {
		JobID   string   `json:"job-id"`
		Tag     string   `json:"tag"`
		Vmstate string   `json:"vmstate"`
		Devices []string `json:"devices"`
	}{
		jobID,
		tag,
		vmstate,
		devices,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "snapshot-save",
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

// x-blockdev-amend -> XBlockdevAmend (command)

// XBlockdevAmend implements the "x-blockdev-amend" QMP API call.
func (m *Monitor) XBlockdevAmend(jobID string, nodeName string, options BlockdevAmendOptions, force *bool) (err error) {
	cmd := struct {
		JobID    string               `json:"job-id"`
		NodeName string               `json:"node-name"`
		Options  BlockdevAmendOptions `json:"options"`
		Force    *bool                `json:"force,omitempty"`
	}{
		jobID,
		nodeName,
		options,
		force,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-blockdev-amend",
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

// x-blockdev-set-iothread -> XBlockdevSetIothread (command)

// XBlockdevSetIothread implements the "x-blockdev-set-iothread" QMP API call.
func (m *Monitor) XBlockdevSetIothread(nodeName string, iothread StrOrNull, force *bool) (err error) {
	cmd := struct {
		NodeName string    `json:"node-name"`
		Iothread StrOrNull `json:"iothread"`
		Force    *bool     `json:"force,omitempty"`
	}{
		nodeName,
		iothread,
		force,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-blockdev-set-iothread",
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

// x-debug-block-dirty-bitmap-sha256 -> XDebugBlockDirtyBitmapSha256 (command)

// XDebugBlockDirtyBitmapSha256 implements the "x-debug-block-dirty-bitmap-sha256" QMP API call.
func (m *Monitor) XDebugBlockDirtyBitmapSha256(node string, name string) (ret BlockDirtyBitmapSha256, err error) {
	cmd := struct {
		Node string `json:"node"`
		Name string `json:"name"`
	}{
		node,
		name,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-debug-block-dirty-bitmap-sha256",
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

// x-debug-query-block-graph -> XDebugQueryBlockGraph (command)

// XDebugQueryBlockGraph implements the "x-debug-query-block-graph" QMP API call.
func (m *Monitor) XDebugQueryBlockGraph() (ret XDbgBlockGraph, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-debug-query-block-graph",
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

// x-exit-preconfig -> XExitPreconfig (command)

// XExitPreconfig implements the "x-exit-preconfig" QMP API call.
func (m *Monitor) XExitPreconfig() (err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-exit-preconfig",
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

// x-query-irq -> XQueryIrq (command)

// XQueryIrq implements the "x-query-irq" QMP API call.
func (m *Monitor) XQueryIrq() (ret HumanReadableText, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-query-irq",
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

// x-query-jit -> XQueryJit (command)

// XQueryJit implements the "x-query-jit" QMP API call.
func (m *Monitor) XQueryJit() (ret HumanReadableText, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-query-jit",
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

// x-query-numa -> XQueryNuma (command)

// XQueryNuma implements the "x-query-numa" QMP API call.
func (m *Monitor) XQueryNuma() (ret HumanReadableText, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-query-numa",
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

// x-query-opcount -> XQueryOpcount (command)

// XQueryOpcount implements the "x-query-opcount" QMP API call.
func (m *Monitor) XQueryOpcount() (ret HumanReadableText, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-query-opcount",
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

// x-query-profile -> XQueryProfile (command)

// XQueryProfile implements the "x-query-profile" QMP API call.
func (m *Monitor) XQueryProfile() (ret HumanReadableText, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-query-profile",
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

// x-query-ramblock -> XQueryRamblock (command)

// XQueryRamblock implements the "x-query-ramblock" QMP API call.
func (m *Monitor) XQueryRamblock() (ret HumanReadableText, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-query-ramblock",
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

// x-query-rdma -> XQueryRdma (command)

// XQueryRdma implements the "x-query-rdma" QMP API call.
func (m *Monitor) XQueryRdma() (ret HumanReadableText, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-query-rdma",
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

// x-query-roms -> XQueryRoms (command)

// XQueryRoms implements the "x-query-roms" QMP API call.
func (m *Monitor) XQueryRoms() (ret HumanReadableText, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-query-roms",
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

// x-query-usb -> XQueryUsb (command)

// XQueryUsb implements the "x-query-usb" QMP API call.
func (m *Monitor) XQueryUsb() (ret HumanReadableText, err error) {
	cmd := struct {
	}{}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "x-query-usb",
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

// yank -> Yank (command)

// Yank implements the "yank" QMP API call.
func (m *Monitor) Yank(instances []YankInstance) (err error) {
	cmd := struct {
		Instances []YankInstance `json:"instances"`
	}{
		instances,
	}
	bs, err := json.Marshal(map[string]interface{}{
		"execute":   "yank",
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
