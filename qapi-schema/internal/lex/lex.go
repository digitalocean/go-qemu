package lex

import (
	"unicode/utf8"

	"github.com/digitalocean/go-qemu/qapi-schema/internal/token"
)

// Lex lexes the input document into tokens.
func Lex(input string) []token.Token {
	lexer := newLexer(input)

	var tokens []token.Token
	for {
		t := lexer.NextToken()
		tokens = append(tokens, t)
		if t.Type == token.EOF {
			break
		}
	}
	return tokens
}

type lexer struct {
	input        string
	position     int
	readPosition int
	ch           rune
}

func newLexer(input string) *lexer {
	l := lexer{
		input: input,
	}

	l.readRune()
	return &l
}

func (l *lexer) NextToken() token.Token {
	var tok token.Token

	l.ignoreCommentsAndWhitespace()

	switch l.ch {
	case '{':
		tok = token.Token{Type: token.LeftCurly, Literal: "{"}
	case '}':
		tok = token.Token{Type: token.RightCurly, Literal: "}"}
	case '[':
		tok = token.Token{Type: token.LeftSquare, Literal: "["}
	case ']':
		tok = token.Token{Type: token.RightSquare, Literal: "]"}
	case ',':
		tok = token.Token{Type: token.Comma, Literal: ","}
	case '\'':
		tok = token.Token{Type: token.SingleQuote, Literal: "'"}
	case ':':
		tok = token.Token{Type: token.Colon, Literal: ":"}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isAlphaNumeric(l.ch) {
			tok.Literal = l.readAlphaNumeric()
			tok.Type = token.AlphaNumeric
			return tok
		}
	}

	l.readRune()
	return tok
}

func (l *lexer) readRune() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
		return
	}

	ch, size := utf8.DecodeRuneInString(l.input[l.readPosition:])
	l.ch = ch
	l.position = l.readPosition
	l.readPosition += size

}

func (l *lexer) readAlphaNumeric() string {
	position := l.position
	for isAlphaNumeric(l.ch) {
		l.readRune()
	}
	return l.input[position:l.position]
}

func (l *lexer) ignoreCommentsAndWhitespace() {
	for {
		switch {
		case l.ch == '#':
			l.ignoreComment()
		case isWhitespace(l.ch):
			l.ignoreWhitespace()
		default:
			return
		}
	}
}

var endOfLine = []rune{'\n', '\r'}

func (l *lexer) ignoreComment() {
	for !containsRune(endOfLine, l.ch) {
		l.readRune()
	}
}

func (l *lexer) ignoreWhitespace() {
	for isWhitespace(l.ch) {
		l.readRune()
	}
}

var whitespace = []rune{' ', '\n', '\t', '\r'}

func isWhitespace(r rune) bool {
	return containsRune(whitespace, r)
}

var additionAlphaNumeric = []rune{'_', '-', '.', '/', '*'}

func isAlphaNumeric(r rune) bool {
	return r >= '0' && r <= '9' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || containsRune(additionAlphaNumeric, r)
}

func containsRune(haystack []rune, needle rune) bool {
	for _, r := range haystack {
		if r == needle {
			return true
		}
	}
	return false
}
