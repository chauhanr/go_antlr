package main

import (
	"fmt"
	"io/ioutil"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chauhanr/go_antlr/translator/parser"
)

type translatorListener struct {
	*parser.BaseJavaParserListener
}

func (l *translatorListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	fmt.Println("interface I" + ctx.Identifier().GetText() + "{")
}

func (l *translatorListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	fmt.Println("}")
}

func (l *translatorListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	ts := ctx.GetParser().GetTokenStream()
	r_type := ts.GetTextFromInterval(ctx.TypeTypeOrVoid().GetSourceInterval())
	args := ts.GetTextFromInterval(ctx.FormalParameters().GetSourceInterval())
	fmt.Println("\t" + r_type + " " + ctx.Identifier().GetText() + args + ";")
}

func main() {
	b, err := ioutil.ReadFile("Demo.java")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	is := antlr.NewInputStream(string(b))
	lexer := parser.NewJavaLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewJavaParser(stream)
	var l translatorListener
	antlr.ParseTreeWalkerDefault.Walk(&l, p.CompilationUnit())

}
