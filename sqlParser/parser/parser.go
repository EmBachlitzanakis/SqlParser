// parser/parser.go
package parser

import (
	"SqlParser/tokenizer"
	"fmt"
)

type Parser struct {
	tokenizer *tokenizer.Tokenizer
	current   tokenizer.Token
	err       error
}

func NewParser(input string) *Parser {
	tokenizer := tokenizer.NewTokenizer(input)
	return &Parser{tokenizer: tokenizer}
}

func (p *Parser) read() {
	p.current = p.tokenizer.NextToken()
}

func (p *Parser) reportError(message string) {
	p.err = fmt.Errorf("%s at position %d", message, p.tokenizer.Pos())
}

func (p *Parser) Parse() error {
	p.read()
	if p.current.Type != tokenizer.TokenSelect {
		p.reportError("expected SELECT keyword")
		return p.err
	}
	fmt.Println("Parsed: SELECT")

	p.read()
	for p.current.Type == tokenizer.TokenIdentifier {
		fmt.Printf("Parsed identifier: %s\n", p.current.Value)
		p.read()
		if p.current.Type == tokenizer.TokenComma {
			p.read()
		} else {
			break
		}
	}

	if p.current.Type != tokenizer.TokenFrom {
		p.reportError("expected FROM keyword")
		return p.err
	}
	fmt.Println("Parsed: FROM")

	p.read()
	if p.current.Type != tokenizer.TokenIdentifier {
		p.reportError("expected table name")
		return p.err
	}
	fmt.Printf("Parsed table name: %s\n", p.current.Value)

	return nil
}
