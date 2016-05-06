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

// A Preparer can prepare a Command for execution using the input binary
// name and arguments.
type Preparer interface {
	Prepare(name string, args ...string) Command
}

// A Command is a command which is typically executed by shelling out.
type Command interface {
	Exited() bool
	Kill() error
	Output() ([]byte, error)
	Start() error
	StdoutPipe() (io.ReadCloser, error)
	Wait() error
}

var _ Preparer = &SystemPreparer{}

// A SystemPreparer is a Preparer which prepares Commands that shells out
// using package os/exec.
type SystemPreparer struct{}

// Prepare prepares a Command which shells out to a binary with arguments,
// using package os/exec.
func (SystemPreparer) Prepare(name string, args ...string) Command {
	return &systemCommand{
		Cmd: exec.Command(name, args...),
	}
}

var _ Command = &systemCommand{}

// A systemCommand is a Command which shells out to perform actions.
type systemCommand struct {
	*exec.Cmd
}

// Exited reports whether the command has exited.
func (cmd *systemCommand) Exited() bool {
	if cmd.Cmd.ProcessState == nil {
		return true
	}

	return cmd.Cmd.ProcessState.Exited()
}

// Kill kills the process wrapped by the systemCommand.
func (cmd *systemCommand) Kill() error {
	return cmd.Cmd.Process.Kill()
}

// NoopCommand returns a Command which does nothing and returns no error
// for all Command methods.
//
// The returned Command should only be used in tests.
func NoopCommand() Command {
	return &noopCommand{}
}

var _ Command = &noopCommand{}

type noopCommand struct{}

func (noopCommand) Exited() bool                       { return false }
func (noopCommand) Kill() error                        { return nil }
func (noopCommand) Output() ([]byte, error)            { return nil, nil }
func (noopCommand) Start() error                       { return nil }
func (noopCommand) StdoutPipe() (io.ReadCloser, error) { return nil, nil }
func (noopCommand) Wait() error                        { return nil }
