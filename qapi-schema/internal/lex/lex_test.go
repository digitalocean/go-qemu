package lex

import (
	"testing"

	"github.com/digitalocean/go-qemu/qapi-schema/internal/token"
)

func TestLex(t *testing.T) {
	input := `##
# this is a comment
## 1. hello, world

{ 'include': 'block-job.json' }

{ 'enum': 'arch',
  'data': ['x86', 'x86_64', 'arm64'] }`

	wantTokens := []token.Token{
		{Type: token.LeftCurly, Literal: "{"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.AlphaNumeric, Literal: "include"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.Colon, Literal: ":"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.AlphaNumeric, Literal: "block-job.json"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.RightCurly, Literal: "}"},
		{Type: token.LeftCurly, Literal: "{"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.AlphaNumeric, Literal: "enum"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.Colon, Literal: ":"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.AlphaNumeric, Literal: "arch"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.Comma, Literal: ","},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.AlphaNumeric, Literal: "data"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.Colon, Literal: ":"},
		{Type: token.LeftSquare, Literal: "["},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.AlphaNumeric, Literal: "x86"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.Comma, Literal: ","},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.AlphaNumeric, Literal: "x86_64"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.Comma, Literal: ","},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.AlphaNumeric, Literal: "arm64"},
		{Type: token.SingleQuote, Literal: "'"},
		{Type: token.RightSquare, Literal: "]"},
		{Type: token.RightCurly, Literal: "}"},
		{Type: token.EOF, Literal: ""},
	}

	lex := newLexer(input)

	for num, tt := range wantTokens {
		got := lex.NextToken()

		if tt.Type != got.Type {
			t.Errorf("test [%d] wrong token type - want %q, got %q", num, tt.Type, got.Type)
		}

		if tt.Literal != got.Literal {
			t.Errorf("test [%d] wrong token literal - want %q, got %q", num, tt.Literal, got.Literal)
		}
	}
}
