package token

// Type describes the type of token.
type Type string

// Types of tokens.
const (
	Illegal = "illegal"
	EOF     = "EOF"

	LeftCurly    = "{"
	RightCurly   = "}"
	LeftSquare   = "["
	RightSquare  = "]"
	Comma        = ","
	SingleQuote  = "'"
	Colon        = ":"
	AlphaNumeric = "alphanum"
)

// Token is a lexed token.
type Token struct {
	Type    Type
	Literal string
}
