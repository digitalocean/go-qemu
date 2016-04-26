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

// Package qmp enables interaction with QEMU instances
// via the QEMU Machine Protocol (QMP).
package qmp

// Monitor represents a QEMU Machine Protocol socket.
// See: http://wiki.qemu.org/QMP
type Monitor interface {
	Connect() error
	Disconnect() error
	Run(command []byte) (out []byte, err error)
	Events() (events <-chan Event, err error)
}

// Cmd represents a QMP command.
type Cmd struct {
	// Name of the command to run
	Execute string `json:"execute"`

	// Optional arguments for the above command.
	Args interface{} `json:"arguments,omitempty"`
}

// Event represents a QEMU QMP event.
// See http://wiki.qemu.org/QMP
type Event struct {
	// Event name, e.g., BLOCK_JOB_COMPLETE
	Event string

	// Arbitrary event data
	Data map[string]interface{}

	// Event timestamp, provided by QEMU.
	Timestamp struct {
		Seconds      int64
		Microseconds int64
	}
}
