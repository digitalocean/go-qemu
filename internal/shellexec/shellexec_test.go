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
	"reflect"
	"testing"
)

func TestSystemCommand(t *testing.T) {
	sys := &SystemCommand{}

	cmd := sys.Prepare("echo", "hello", "world")

	if reflect.DeepEqual(sys, cmd) {
		t.Fatalf("Prepare should create a copy; SystemCommands should not be equal:\n- %#v\n- %#v",
			sys, cmd)
	}

	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if want, got := "hello world\n", string(out); want != got {
		t.Fatalf("unexpected output:\n- want: %q\n-  got: %q",
			want, got)
	}
}
