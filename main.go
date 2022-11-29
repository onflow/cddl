package main

//go:generate antlr -Dlanguage=Go -o parser -package parser -listener -no-visitor CDDL.g4

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.org/onflow/cddl_gen/parser"
)

type TreeShapeListener struct {
	*parser.BaseCDDLListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (*TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	println(ctx.GetText())
}

func main() {
	filename := os.Args[1]
	contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	is := antlr.NewInputStream(string(contents))
	lexer := parser.NewCDDLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCDDLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Cddl()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
