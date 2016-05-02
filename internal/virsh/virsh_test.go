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

package virsh

import (
	"os/exec"
	"testing"

	"github.com/digitalocean/go-qemu/internal/shellexec"
)

func TestVirshEchoShell(t *testing.T) {
	if !virshInstalled(t) {
		t.Skip("virsh not found in path, skipping test")
	}

	testVirshOutput(
		t,
		&shellexec.SystemPreparer{},
	)
}

func TestVirshEchoMock(t *testing.T) {
	testVirshOutput(
		t,
		&outputPreparer{
			out: []byte("hello world\n"),
		},
	)
}

func TestVirshErrorMock(t *testing.T) {
	const want = "some error"

	prep := &outputPreparer{
		err: &exec.ExitError{
			Stderr: []byte("error: SOME ERROR"),
		},
	}

	_, gotErr := Virsh(prep, "")
	if want, got := want, gotErr.Error(); want != got {
		t.Fatalf("unexpected error output:\n- want: %q\n-  got: %q",
			want, got)
	}
}

func testVirshOutput(t *testing.T, prep shellexec.Preparer) {
	out, err := Virsh(prep, "qemu:///system", "echo", "hello", "world")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if want, got := "hello world\n", string(out); want != got {
		t.Fatalf("unexpected command output:\n- want: %q\n-  got: %q",
			want, got)
	}
}

func virshInstalled(t *testing.T) bool {
	_, err := exec.LookPath("virsh")
	if err == nil {
		return true
	}

	eerr, ok := err.(*exec.Error)
	if !ok || eerr.Err != exec.ErrNotFound {
		t.Fatalf("unexpected error: %v", err)
	}

	return false
}

var _ shellexec.Preparer = &outputPreparer{}

type outputPreparer struct {
	out []byte
	err error
}

func (p *outputPreparer) Prepare(_ string, _ ...string) shellexec.Command {
	return &outputCommand{
		out: p.out,
		err: p.err,
	}
}

var _ shellexec.Command = &outputCommand{}

type outputCommand struct {
	out []byte
	err error
	shellexec.Command
}

func (cmd *outputCommand) Output() ([]byte, error) { return cmd.out, cmd.err }
