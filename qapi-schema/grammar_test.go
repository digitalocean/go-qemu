package qapischema

import (
	"reflect"
	"testing"

	"github.com/alecthomas/repr"

	"github.com/digitalocean/go-qemu/qapi-schema/internal/parse"
	"github.com/digitalocean/go-qemu/qapi-schema/internal/token"
)

func TestTrueParser(t *testing.T) {
	tests := []struct {
		name    string
		input   []token.Token
		want    bool
		wantErr error
	}{
		{
			name: "true",
			input: []token.Token{
				{Type: token.AlphaNumeric, Literal: "true"},
			},
			want: true,
		},
		{
			name: "no match",
			input: []token.Token{
				{Type: token.AlphaNumeric, Literal: "troo"},
			},
			want:    false,
			wantErr: parse.ErrNoMatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse.Parse(trueParser, tt.input)
			if tt.wantErr != err {
				t.Errorf("want err %v, got err %v", tt.wantErr, err)
			}

			if tt.want != got {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestFalseParser(t *testing.T) {
	tests := []struct {
		name    string
		input   []token.Token
		want    bool
		wantErr error
	}{
		{
			name: "false",
			input: []token.Token{
				{Type: token.AlphaNumeric, Literal: "false"},
			},
			want: false,
		},
		{
			name: "no match",
			input: []token.Token{
				{Type: token.AlphaNumeric, Literal: "true"},
			},
			want:    false,
			wantErr: parse.ErrNoMatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse.Parse(falseParser, tt.input)
			if tt.wantErr != err {
				t.Errorf("want err %v, got err %v", tt.wantErr, err)
			}

			if tt.want != got {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestBoolParser(t *testing.T) {
	tests := []struct {
		name    string
		input   []token.Token
		want    bool
		wantErr error
	}{
		{
			name:  "true",
			input: []token.Token{{Type: token.AlphaNumeric, Literal: "true"}},
			want:  true,
		},
		{
			name:  "false",
			input: []token.Token{{Type: token.AlphaNumeric, Literal: "false"}},
			want:  false,
		},
		{
			name:    "no match",
			input:   []token.Token{{Type: token.AlphaNumeric, Literal: "tralse"}},
			want:    false,
			wantErr: parse.ErrNoMatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse.Parse(boolParser, tt.input)

			if tt.wantErr != err {
				t.Errorf("want err %v, got err %v", tt.wantErr, err)
			}

			if tt.want != got {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestStrParser(t *testing.T) {
	tests := []struct {
		name    string
		input   []token.Token
		want    string
		wantErr error
	}{
		{
			name: "basic string",
			input: []token.Token{
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "hello"},
				{Type: token.SingleQuote, Literal: "'"},
			},
			want: "hello",
		},
		{
			name: "incomplete",
			input: []token.Token{
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "hello"},
			},
			wantErr: parse.ErrNoMatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse.Parse(strParser, tt.input)

			if tt.wantErr != err {
				t.Errorf("want err %v, got err %v", tt.wantErr, err)
			}

			if tt.want != got {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}

func TestAtLeastOneStrParser(t *testing.T) {
	tests := []struct {
		name    string
		input   []token.Token
		want    []string
		wantErr error
	}{
		{
			name: "one",
			input: []token.Token{
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "hello"},
				{Type: token.SingleQuote, Literal: "'"},
			},
			want: []string{"hello"},
		},
		{
			name: "two",
			input: []token.Token{
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "hello"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.Comma, Literal: ","},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "world"},
				{Type: token.SingleQuote, Literal: "'"},
			},
			want: []string{"hello", "world"},
		},
		{
			name: "many",
			input: []token.Token{
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "hello"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.Comma, Literal: ","},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "world"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.Comma, Literal: ","},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "hi"},
				{Type: token.SingleQuote, Literal: "'"},
			},
			want: []string{"hello", "world", "hi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse.Parse(atLeastOneStrParser, tt.input)
			if tt.wantErr != err {
				t.Errorf("want err %v, got err %v", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestStrArrayParser(t *testing.T) {
	tests := []struct {
		name    string
		input   []token.Token
		want    []string
		wantErr error
	}{
		{
			name: "one",
			input: []token.Token{
				{Type: token.LeftSquare, Literal: "["},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "hello"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.RightSquare, Literal: "]"},
			},
			want: []string{"hello"},
		},
		{
			name: "two",
			input: []token.Token{
				{Type: token.LeftSquare, Literal: "["},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "hello"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.Comma, Literal: ","},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "goodbye"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.RightSquare, Literal: "]"},
			},
			want: []string{"hello", "goodbye"},
		},
		{
			name: "three",
			input: []token.Token{
				{Type: token.LeftSquare, Literal: "["},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "hello"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.Comma, Literal: ","},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "goodbye"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.Comma, Literal: ","},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "ok"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.RightSquare, Literal: "]"},
			},
			want: []string{"hello", "goodbye", "ok"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse.Parse(strArrayParser, tt.input)
			if tt.wantErr != err {
				t.Errorf("want err %v, got err %v", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestIncludeParser(t *testing.T) {
	tests := []struct {
		name    string
		input   []token.Token
		want    string
		wantErr error
	}{
		{
			name: "include foo-bar.json",
			input: []token.Token{
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "include"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.Colon, Literal: ":"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "foo-bar.json"},
				{Type: token.SingleQuote, Literal: "'"},
			},
			want: "foo-bar.json",
		},
		{
			name: "no match",
			input: []token.Token{
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "import"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.Colon, Literal: ":"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "foo-bar.json"},
				{Type: token.SingleQuote, Literal: "'"},
			},
			wantErr: parse.ErrNoMatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse.Parse(includeParser, tt.input)
			if tt.wantErr != err {
				t.Errorf("want err %v, got err %v", tt.wantErr, err)
			}

			if tt.want != got {
				t.Errorf("want %q, got %v", tt.want, got)
			}
		})
	}
}

func TestIncludeObjParser(t *testing.T) {
	tests := []struct {
		name    string
		input   []token.Token
		want    *Tree
		wantErr error
	}{
		{
			name: "include foo-bar.json",
			input: []token.Token{
				{Type: token.LeftCurly, Literal: "{"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "include"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.Colon, Literal: ":"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.AlphaNumeric, Literal: "foo-bar.json"},
				{Type: token.SingleQuote, Literal: "'"},
				{Type: token.RightCurly, Literal: "}"},
			},
			want: &Tree{Node: Include("foo-bar.json")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse.Parse(includeObjParser, tt.input)
			if tt.wantErr != err {
				t.Errorf("want err %v, got err %v", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %s, got %s", repr.String(tt.want), repr.String(got))
			}
		})
	}
}
