// main.go
package main

import (
	"SqlParser/parser"
	"fmt"
)

func main() {
	input := "SELECT name, age FROM users"
	p := parser.NewParser(input)
	p.Parse()

	fmt.Printf("\n")
	input = "SELECT name, age users" // Missing FROM keyword
	p = parser.NewParser(input)
	err := p.Parse()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
