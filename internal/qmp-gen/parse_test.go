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
	"strings"
	"testing"
)

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
