// +build !linux

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

import "fmt"

func (mon LibvirtGoMonitor) Connect() error {
	return notSupportedError()
}

func (mon *LibvirtGoMonitor) Disconnect() error {
	return notSupportedError()
}

func (mon LibvirtGoMonitor) Run(cmd []byte) ([]byte, error) {
	return nil, notSupportedError()
}

func (mon *LibvirtGoMonitor) Events() (<-chan Event, error) {
	return nil, notSupportedError()
}

func NewLibvirtGoMonitor(uri, domain string) Monitor {
	return &LibvirtGoMonitor{}
}

func notSupportedError() error {
	return fmt.Errorf("libvirt-go is only supported on Linux")
}
