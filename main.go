package main

import (
	"fmt"
	"os"

	"antlr4test/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseCustomJsonListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewCustomJsonLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCustomJsonParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Json()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
