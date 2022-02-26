package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chauhanr/go_antlr/calculator/parser"
)

type calcListener struct {
	*parser.BaseCalcListener
	stack []int
}

// push the i into the stack for calculation
func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

/*
   1. if we pop an empty stack return panic
   2. get the last value on the stack and return
   3. return the last element before that shorten the stack.
*/
func (l *calcListener) pop() int {
	if len(l.stack) == 0 {
		panic("stack is empty unable to pop")
	}
	result := l.stack[len(l.stack)-1]
	// remove the last element just popped
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()
	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexepected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()
	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexepected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}
	l.push(i)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("-> ")
		expr, _ := reader.ReadString('\n')
		expr = strings.Replace(expr, "\n", "", -1)
		readExprCall(expr)
	}
}

func readExprCall(expr string) {
	is := antlr.NewInputStream(expr)
	//listTokens(is)
	v := CalculateExpression(is)
	fmt.Printf("-> %d\n", v)
}

func CalculateExpression(is *antlr.InputStream) int {
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewCalcParser(stream)
	// finally parse the expression
	var l calcListener
	antlr.ParseTreeWalkerDefault.Walk(&l, p.Start())
	return l.pop()
}

// THis method will just determine the tokens that are parser/ lexer generates
// and returns name and value
func listTokens(is *antlr.InputStream) {
	lexer := parser.NewCalcLexer(is)

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q) \n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
