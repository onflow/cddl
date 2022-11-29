package main

//go:generate antlr -Dlanguage=Go -o parser -package parser -no-listener -no-visitor CDDL.g4

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.org/onflow/cddl_gen/parser"
)

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

	fmt.Printf("#%+v\n", tree)
}
