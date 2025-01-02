package main

import (
	"fmt"
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

type Parser struct {
	tokenizer *Tokenizer
	current   Token
	err       error
}

func NewParser(input string) *Parser {
	tokenizer := NewTokenizer(input)
	return &Parser{tokenizer: tokenizer}
}

func (p *Parser) read() {
	p.current = p.tokenizer.NextToken()
}

func (p *Parser) reportError(message string) {
	p.err = fmt.Errorf("%s at position %d", message, p.tokenizer.pos)
}

func (p *Parser) Parse() error {
	p.read()
	if p.current.Type != TokenSelect {
		p.reportError("expected SELECT keyword")
		return p.err
	}
	fmt.Println("Parsed: SELECT")

	p.read()
	for p.current.Type == TokenIdentifier {
		fmt.Printf("Parsed identifier: %s\n", p.current.Value)
		p.read()
		if p.current.Type == TokenComma {
			p.read()
		} else {
			break
		}
	}

	if p.current.Type != TokenFrom {
		p.reportError("expected FROM keyword")
		return p.err
	}
	fmt.Println("Parsed: FROM")

	p.read()
	if p.current.Type != TokenIdentifier {
		p.reportError("expected table name")
		return p.err
	}
	fmt.Printf("Parsed table name: %s\n", p.current.Value)

	return nil
}

func main() {
	input := "SELECT name, age FROM users"
	parser := NewParser(input)
	parser.Parse()

	fmt.Printf("\n")
	input = "SELECT name, age users" // Missing FROM keyword
	parser = NewParser(input)
	err := parser.Parse()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
