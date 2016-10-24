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

package gen

import (
	"bytes"
	"go/format"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		out  []byte
	}{
		{
			name: "StructInfo",
			in: []byte(`
##
# @StatusInfo:
#
# Information about VCPU run state
#
# @running: true if all VCPUs are runnable, false if not runnable
#
# @singlestep: true if VCPUs are in single-step mode
#
# @status: the virtual machine @RunState
#
# Since:  0.14.0
#
# Notes: @singlestep is enabled through the GDB stub
##
{ 'struct': 'StatusInfo',
  'data': {'running': 'bool', 'singlestep': 'bool', 'status': 'RunState'} }
			`),
			out: []byte(`
// StatusInfo -> StatusInfo (struct)

// StatusInfo implements the "StatusInfo" QMP API type.
type StatusInfo struct {
	Running    bool     'json:"running"'
	Singlestep bool     'json:"singlestep"'
	Status     RunState 'json:"status"'
}
			`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace single quote with backtick to match generated Go
			// source code
			tt.out = bytes.TrimSpace(bytes.Replace(tt.out, []byte("'"), []byte("`"), -1))

			// Do not check needed types, third parameter is false
			source := testGenerate(t, tt.in, false)

			if want, got := tt.out, source; !bytes.Contains(got, want) {
				t.Fatalf("generated code does not match actual code:\n- want: %s\n-  got: %s",
					indent(string(want)), indent(string(got)))
			}
		})
	}
}

func TestGenerateNeededTypes(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		out  []byte
	}{
		{
			name: "no output: no command that uses StructInfo",
			in: []byte(`
##
# @StatusInfo:
#
# Information about VCPU run state
#
# @running: true if all VCPUs are runnable, false if not runnable
#
# @singlestep: true if VCPUs are in single-step mode
#
# @status: the virtual machine @RunState
#
# Since:  0.14.0
#
# Notes: @singlestep is enabled through the GDB stub
##
{ 'struct': 'StatusInfo',
  'data': {'running': 'bool', 'singlestep': 'bool', 'status': 'RunState'} }
			`),
		},
		{
			name: "StatusInfo with query-status command",
			in: []byte(`
##
# @StatusInfo:
#
# Information about VCPU run state
#
# @running: true if all VCPUs are runnable, false if not runnable
#
# @singlestep: true if VCPUs are in single-step mode
#
# @status: the virtual machine @RunState
#
# Since:  0.14.0
#
# Notes: @singlestep is enabled through the GDB stub
##
{ 'struct': 'StatusInfo',
  'data': {'running': 'bool', 'singlestep': 'bool', 'status': 'RunState'} }

##
# @query-status:
#
# Query the run status of all VCPUs
#
# Returns: @StatusInfo reflecting all VCPUs
#
# Since:  0.14.0
##
{ 'command': 'query-status', 'returns': 'StatusInfo' }
			`),
			out: []byte(`
// StatusInfo -> StatusInfo (struct)

// StatusInfo implements the "StatusInfo" QMP API type.
type StatusInfo struct {
	Running    bool     'json:"running"'
	Singlestep bool     'json:"singlestep"'
	Status     RunState 'json:"status"'
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
		Res json.RawMessage 'json:"return"'
	}{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(res.Res), &ret); err != nil {
		return
	}
	return
}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace single quote with backtick to match generated Go
			// source code
			tt.out = bytes.TrimSpace(bytes.Replace(tt.out, []byte("'"), []byte("`"), -1))

			// Check needed types, third parameter is true
			source := testGenerate(t, tt.in, true)

			if want, got := tt.out, source; !bytes.Contains(got, want) {
				t.Fatalf("generated code does not match actual code:\n- want: %s\n-  got: %s",
					indent(string(want)), indent(string(got)))
			}
		})
	}
}

func indent(s string) string {
	return strings.Replace(s, "\n", "\n  ", -1)
}

func TestPyToJSON(t *testing.T) {
	tests := []struct {
		name, in, out string
	}{
		{
			name: "Basic translation",
			in:   `{'foo': 42, "bar": 'baz'}`,
			out:  `{"foo": 42, "bar": "baz"}`,
		},
		{
			name: "Idempotency",
			in:   `{"foo": 42, "bar": "baz"}`,
			out:  `{"foo": 42, "bar": "baz"}`,
		},
		{
			name: "Lone comment",
			in:   `# foo`,
			out:  ``,
		},
		{
			name: "Whole line comment",
			in: `# This is a test
{'foo': 42, 'bar': 'baz'}
# This is another`,
			out: `{"foo": 42, "bar": "baz"}`,
		},
		{
			name: "Inline comment",
			in: `{'foo': 42, # This is a test
'bar': 'baz'} # This is another`,
			// there's a trailing space on the next line.
			out: `{"foo": 42, 
"bar": "baz"}`,
		},
		{
			name: "Comment in a string",
			in:   `{'foo': 'look, # a comment!'}`,
			out:  `{"foo": "look, # a comment!"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := strings.TrimSpace(string(pyToJSON([]byte(test.in))))
			if actual != test.out {
				t.Errorf(`Wrong output
Input:
  %s
Want:
  %s
Got:
  %s`, indent(test.in), indent(test.out), indent(actual))
			}
		})
	}
}

func testGenerate(t *testing.T, in []byte, checkNeededTypes bool) []byte {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(in)
	}))
	defer s.Close()

	defs, err := readDefinitions(s.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	symbols, err := parse(defs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	need := symbols
	if checkNeededTypes {
		need = neededTypes(symbols)
	}

	bs, err := renderAPI("templates/", symbols, need)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	formatted, err := format.Source(bs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	return formatted
}
