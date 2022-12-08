/*
 * CDDL - A CDDL implementation for Go
 *
 * Copyright 2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package parser

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

//go:generate antlr -Dlanguage=Go -o . -package parser -no-listener -visitor CDDL.g4

type SyntaxError struct {
	Line    int
	Column  int
	Message string
}

type ParseError struct {
	SyntaxErrors []SyntaxError
}

func (p ParseError) Error() string {
	var b strings.Builder
	b.WriteString("parsing failed:\n")
	for _, syntaxError := range p.SyntaxErrors {
		_, _ = fmt.Fprintf(
			&b,
			"%d:%d: %s\n",
			syntaxError.Line,
			syntaxError.Column,
			syntaxError.Message,
		)
	}
	return b.String()
}

type errorListener struct {
	*antlr.DefaultErrorListener
	syntaxErrors []SyntaxError
}

func (l *errorListener) SyntaxError(
	_ antlr.Recognizer,
	_ interface{},
	line, column int,
	msg string,
	_ antlr.RecognitionException,
) {

	l.syntaxErrors = append(
		l.syntaxErrors,
		SyntaxError{
			Line:    line,
			Column:  column,
			Message: msg,
		},
	)
}

func (l *errorListener) ParseError() error {
	if len(l.syntaxErrors) > 0 {
		return ParseError{
			SyntaxErrors: l.syntaxErrors,
		}
	}

	return nil
}

var _ antlr.ErrorListener = &errorListener{}

func ParseCDDL(cddl string) (ICddlContext, error) {
	errorListener := &errorListener{}

	is := antlr.NewInputStream(cddl)
	lexer := NewCDDLLexer(is)
	// remove the lexer's default console error listener
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewCDDLParser(stream)
	// remove the default console error listener
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	return p.Cddl(), errorListener.ParseError()
}

type Position struct {
	Offset int
	Line   int
	Column int
}

func PositionFromToken(token antlr.Token) Position {
	return Position{
		Offset: token.GetStart(),
		Line:   token.GetLine(),
		Column: token.GetColumn(),
	}
}
