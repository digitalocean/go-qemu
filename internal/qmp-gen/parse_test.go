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
		in, out string
	}{
		{
			// Basic translation
			in:  `{'foo': 42, "bar": 'baz'}`,
			out: `{"foo": 42, "bar": "baz"}`,
		},
		{
			// Idempotency
			in:  `{"foo": 42, "bar": "baz"}`,
			out: `{"foo": 42, "bar": "baz"}`,
		},
		{
			// Lone comment
			in:  `# foo`,
			out: ``,
		},
		{
			// Whole line comment
			in: `# This is a test
{'foo': 42, 'bar': 'baz'}
# This is another`,
			out: `{"foo": 42, "bar": "baz"}`,
		},
		{
			// Inline comment
			in: `{'foo': 42, # This is a test
'bar': 'baz'} # This is another`,
			// there's a trailing space on the next line.
			out: `{"foo": 42, 
"bar": "baz"}`,
		},
		{
			// Comment in a string
			in:  `{'foo': 'look, # a comment!'}`,
			out: `{"foo": "look, # a comment!"}`,
		},
	}

	for i, test := range tests {
		actual := strings.TrimSpace(string(pyToJSON([]byte(test.in))))
		if actual != test.out {
			t.Errorf(`Case %d has wrong output
Input:
  %s
Want:
  %s
Got:
  %s`, i+1, indent(test.in), indent(test.out), indent(actual))
		}
	}
}
