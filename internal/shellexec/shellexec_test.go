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

package shellexec

import (
	"testing"
)

func TestSystemPreparer(t *testing.T) {
	testHelloWorldOutput(
		t,
		&SystemPreparer{},
	)
}

func TestCustomPreparer(t *testing.T) {
	testHelloWorldOutput(
		t,
		&customPreparer{},
	)
}

func testHelloWorldOutput(t *testing.T, prep Preparer) {
	cmd := prep.Prepare("echo", "hello", "world")

	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if want, got := "hello world\n", string(out); want != got {
		t.Fatalf("unexpected output:\n- want: %q\n-  got: %q",
			want, got)
	}
}

var _ Preparer = &customPreparer{}

type customPreparer struct{}

func (p *customPreparer) Prepare(_ string, _ ...string) Command {
	return &customCommand{
		out:     []byte("hello world\n"),
		Command: NoopCommand(),
	}
}

var _ Command = &customCommand{}

type customCommand struct {
	out []byte
	Command
}

func (cmd *customCommand) Output() ([]byte, error) {
	return cmd.out, nil
}
