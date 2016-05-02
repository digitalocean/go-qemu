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

// Package shellexec provides the Command interface to allowing testing of
// actions which shell out.
package shellexec

import (
	"io"
	"os/exec"
)

// A Command is a command which is typically executed by shelling out.
type Command interface {
	// Prepare copies a Command and prepares it for use with the input binary
	// name and arguments.
	Prepare(name string, args ...string) Command

	// Standard methods used by os/exec.Command.
	Kill() error
	Output() ([]byte, error)
	Start() error
	StdoutPipe() (io.ReadCloser, error)
	Wait() error
}

var _ Command = &SystemCommand{}

// A SystemCommand is a wrapper type for os/exec.Command.
type SystemCommand struct {
	*exec.Cmd
}

// Prepare prepares a copy of a SystemCommand for shelling out to
// a binary with arguments.
func (cmd SystemCommand) Prepare(name string, args ...string) Command {
	cmd.Cmd = exec.Command(name, args...)
	return &cmd
}

// Kill kills the process wrapped by an os/exec.Command.
func (cmd *SystemCommand) Kill() error {
	return cmd.Cmd.Process.Kill()
}
