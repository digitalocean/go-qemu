package parse

import (
	"errors"

	"github.com/digitalocean/go-qemu/qapi-schema/internal/token"
)

var (
	// ErrNoMatch indicates parsing failed against the token stream.
	ErrNoMatch = errors.New("no match")

	// ErrNonExhaustiveParse indicates parsing completed but there were
	// still tokens remaining in the input stream.
	ErrNonExhaustiveParse = errors.New("did not consume all tokens")
)

// Parse applies the given parser against the input stream.
func Parse[T any](parser Parser[T], input []token.Token) (T, error) {
	initial := state{input: input}
	val, next, err := parser(initial)
	if err != nil {
		return val, err
	}

	if len(next.remaining()) > 0 {
		return val, ErrNonExhaustiveParse
	}

	return val, nil
}

// Parser is a Parser function that can read an input token stream
// and produce the specified type or an error.
type Parser[T any] func(state) (T, state, error)

// Empty is an empty type to return when a parser successfully parsed
// the input stream and nothing should be returned.
type Empty struct{}

// Exactly consumes the exact token from the input stream or fails.
func Exactly(tok token.Token) Parser[Empty] {
	return func(initial state) (Empty, state, error) {
		if len(initial.remaining()) >= 1 {
			if initial.remaining()[0] == tok {
				return Empty{}, initial.consume(1), nil
			}
		}
		return Empty{}, initial, ErrNoMatch
	}
}

// MatchesToken consumes the token from the input stream if matcher
// returns true.
func MatchesToken(matcher func(token.Token) bool) Parser[token.Token] {
	return func(initial state) (token.Token, state, error) {
		if len(initial.remaining()) > 0 {
			if matcher(initial.remaining()[0]) {
				return initial.remaining()[0], initial.consume(1), nil
			}
		}
		return token.Token{}, initial, ErrNoMatch
	}
}

// Map converts one parser to another parser.
func Map[T any, U any](parser Parser[T], mapper func(T) U) Parser[U] {
	return func(initial state) (U, state, error) {
		val, next, err := parser(initial)
		if err != nil {
			var zero U
			return zero, initial, err
		}
		return mapper(val), next, nil
	}
}

// OneOf tries a list of parsers and if one of them succeeds, returns
// that parser.
func OneOf[T any](parsers ...Parser[T]) Parser[T] {
	return func(initial state) (T, state, error) {
		for _, parser := range parsers {
			val, next, err := parser(initial)
			if err == nil {
				return val, next, nil
			}
		}

		var zero T
		return zero, initial, ErrNoMatch
	}
}

// ManyOf attempts to successfully parse 0 or more times.
func ManyOf[T any](parser Parser[T]) Parser[[]T] {
	return func(initial state) ([]T, state, error) {
		var successes []T
		current := initial
		for {
			val, next, err := parser(current)
			if err != nil {
				break
			}

			successes = append(successes, val)
			current = next
		}
		return successes, current, nil
	}
}

// ConsumeAtLeastOne attempts to parse at least once with the given
// parser and then continues until it fails to parse.
func ConsumeAtLeastOne[T any](parser Parser[T]) Parser[[]T] {
	return func(initial state) ([]T, state, error) {
		val, next, err := parser(initial)
		if err != nil {
			return nil, initial, err
		}

		current := next
		successes := []T{val}
		for {
			val, next, err = parser(current)
			if err != nil {
				break
			}

			current = next
			successes = append(successes, val)
		}

		return successes, current, nil
	}
}

// Seq2 is a tuple of two parsed objects.
type Seq2[A any, B any] struct {
	First  A
	Second B
}

// All2 requires both parsers to succeed and it returns a tuple
// of the parsed objects.
func All2[A any, B any](first Parser[A], second Parser[B]) Parser[Seq2[A, B]] {
	return func(initial state) (Seq2[A, B], state, error) {
		var val Seq2[A, B]

		v1, next, err := first(initial)
		if err != nil {
			return val, initial, err
		}

		v2, next, err := second(next)
		if err != nil {
			return val, initial, err
		}

		val.First = v1
		val.Second = v2

		return val, next, nil
	}
}

// Seq3 is a tuple of three parsed objects.
type Seq3[A any, B any, C any] struct {
	First  A
	Second B
	Third  C
}

// All3 requires all three parsers to succeed and it returns a
// tuple of the parsed objects.
func All3[A any, B any, C any](first Parser[A], second Parser[B], third Parser[C]) Parser[Seq3[A, B, C]] {
	return func(initial state) (Seq3[A, B, C], state, error) {
		var val Seq3[A, B, C]

		v1, next, err := first(initial)
		if err != nil {
			return val, initial, err
		}

		v2, next, err := second(next)
		if err != nil {
			return val, initial, err
		}

		v3, next, err := third(next)
		if err != nil {
			return val, initial, err
		}

		val.First = v1
		val.Second = v2
		val.Third = v3

		return val, next, nil
	}
}

// Nothing is a parser that always succeeds and doesn't consume input.
func Nothing() Parser[Empty] {
	return func(initial state) (Empty, state, error) {
		return Empty{}, initial, nil
	}
}

type state struct {
	input  []token.Token
	offset int
}

func (s state) remaining() []token.Token {
	return s.input[s.offset:]
}

func (s state) consume(n int) state {
	s.offset += n
	return s
}
