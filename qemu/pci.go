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

// PCIDevice represents a QEMU PCI device.
type PCIDevice struct {
	Bus       int    `json:"bus"`
	QdevID    string `json:"qdev_id"`
	Slot      int    `json:"slot"`
	ClassInfo struct {
		Class int    `json:"class"`
		Desc  string `json:"desc"`
	} `json:"class_info"`
	ID struct {
		Device int `json:"device"`
		Vendor int `json:"vendor"`
	} `json:"id"`
	Function int `json:"function"`
	Regions  []struct {
		Prefetch  bool   `json:"prefetch"`
		MemType64 bool   `json:"mem_type_64"`
		Bar       int    `json:"bar"`
		Size      int    `json:"size"`
		Address   int64  `json:"address"`
		Type      string `json:"type"`
	} `json:"regions"`
}
