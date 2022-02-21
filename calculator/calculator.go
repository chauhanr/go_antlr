package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chauhanr/go_antlr/calculator/parser"
)

func main() {
	is := antlr.NewInputStream("1 + 2 * 3")
	lexer := parser.NewCalcLexer(is)

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q) \n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}

}
