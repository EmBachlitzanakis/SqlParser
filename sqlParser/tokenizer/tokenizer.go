// tokenizer/tokenizer.go
package tokenizer

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenSelect
	TokenFrom
	TokenIdentifier
	TokenComma
	TokenWhitespace
	TokenUnknown
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	input   string
	pos     int
	length  int
	current rune
}

func NewTokenizer(input string) *Tokenizer {
	t := &Tokenizer{input: input, length: len(input)}
	t.read()
	return t
}

func (t *Tokenizer) read() {
	if t.pos >= t.length {
		t.current = 0 // EOF
	} else {
		t.current = rune(t.input[t.pos])
	}
	t.pos++
}

// Pos returns the current position of the tokenizer in the input string.
func (t *Tokenizer) Pos() int {
	return t.pos
}

func (t *Tokenizer) NextToken() Token {
	for unicode.IsSpace(t.current) {
		t.read()
	}

	if t.current == 0 {
		return Token{Type: TokenEOF}
	}

	switch t.current {
	case ',':
		t.read()
		return Token{Type: TokenComma, Value: ","}
	}

	start := t.pos - 1
	if unicode.IsLetter(t.current) {
		for unicode.IsLetter(t.current) || unicode.IsDigit(t.current) || t.current == '_' {
			t.read()
		}
		value := t.input[start : t.pos-1]
		switch strings.ToUpper(value) {
		case "SELECT":
			return Token{Type: TokenSelect, Value: value}
		case "FROM":
			return Token{Type: TokenFrom, Value: value}
		default:
			return Token{Type: TokenIdentifier, Value: value}
		}
	}

	return Token{Type: TokenUnknown, Value: string(t.current)}
}
