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

// Package qmptest provides types which assist in testing interactions with
// package qmp.
package qmptest

import (
	"encoding/json"

	"github.com/digitalocean/go-qemu/qmp"
)

// A RunFunc is a function invoked to process a qmp.Command.  Any data
// returned by results is marshaled into JSON and returned to the caller.
type RunFunc func(cmd qmp.Command) (results interface{}, err error)

// NewMonitor creates a qmp.Monitor that invokes runFunc whenever its Run method
// is called.  All other methods are a no-op, and return nil.
func NewMonitor(runFunc RunFunc) qmp.Monitor {
	return &monitor{
		fn: runFunc,
	}
}

var _ qmp.Monitor = &monitor{}

type monitor struct {
	fn RunFunc
	noopMonitor
}

func (t *monitor) Run(raw []byte) ([]byte, error) {
	var cmd qmp.Command
	if err := json.Unmarshal(raw, &cmd); err != nil {
		return nil, err
	}

	result, err := t.fn(cmd)
	if err != nil {
		return nil, err
	}

	return json.Marshal(result)
}

var _ qmp.Monitor = &noopMonitor{}

type noopMonitor struct{}

func (noopMonitor) Connect() error                    { return nil }
func (noopMonitor) Disconnect() error                 { return nil }
func (noopMonitor) Run(_ []byte) ([]byte, error)      { return nil, nil }
func (noopMonitor) Events() (<-chan qmp.Event, error) { return nil, nil }
