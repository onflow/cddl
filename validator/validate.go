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

package validator

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.org/onflow/cddl/parser"
)

var predeclaredIDs = map[string]struct{}{
	"any": {},

	"uint": {},
	"nint": {},
	"int":  {},

	"bstr":  {},
	"bytes": {},
	"tstr":  {},
	"text":  {},

	"tdate":      {},
	"time":       {},
	"number":     {},
	"biguint":    {},
	"bignint":    {},
	"bigint":     {},
	"integer":    {},
	"unsigned":   {},
	"decfrac":    {},
	"bigfloat":   {},
	"eb64url":    {},
	"eb64legacy": {},
	"eb16":       {},
	"encoded":    {},
	"uri":        {},
	"b64url":     {},
	"b64legacy":  {},
	"regexp":     {},
	"mime":       {},
	"cbor":       {},

	"float16":    {},
	"float32":    {},
	"float64":    {},
	"float16-32": {},
	"float32-64": {},
	"float":      {},

	"false":     {},
	"true":      {},
	"bool":      {},
	"nil":       {},
	"null":      {},
	"undefined": {},
}

type RuleUse struct {
	ID string
	parser.Position
}

type Validator struct {
	definedRules map[string]struct{}
	usedRules    []RuleUse
}

var _ parser.CDDLVisitor = &Validator{}

func (v *Validator) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Validator) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		ruleChild, ok := child.(antlr.RuleNode)
		if !ok {
			continue
		}

		result := ruleChild.Accept(v)
		if result != nil {
			return result
		}
	}

	return nil
}

func (v *Validator) VisitTerminal(_ antlr.TerminalNode) interface{} {
	return nil
}

func (v *Validator) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	return nil
}

func (v *Validator) VisitCddl(ctx *parser.CddlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitTypeRule(ctx *parser.TypeRuleContext) interface{} {
	id := ctx.ID().GetText()
	v.defineRule(id)

	return v.VisitChildren(ctx)
}

func (v *Validator) VisitGroupRule(ctx *parser.GroupRuleContext) interface{} {
	id := ctx.ID().GetText()
	v.defineRule(id)

	return v.VisitChildren(ctx)
}

func (v *Validator) VisitAssignType(ctx *parser.AssignTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitAssignGroup(ctx *parser.AssignGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitGenericParam(ctx *parser.GenericParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitGenericArg(ctx *parser.GenericArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitType(ctx *parser.TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitType1(ctx *parser.Type1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitValueExpr(ctx *parser.ValueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	v.addRuleUse(ctx.ID())
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitGroupExpr(ctx *parser.GroupExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitMapExpr(ctx *parser.MapExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitArrayExpr(ctx *parser.ArrayExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitUnwrapExpr(ctx *parser.UnwrapExprContext) interface{} {
	v.addRuleUse(ctx.ID())
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitChoiceExpr(ctx *parser.ChoiceExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitChoiceIDExpr(ctx *parser.ChoiceIDExprContext) interface{} {
	v.addRuleUse(ctx.ID())
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitTaggedExpr(ctx *parser.TaggedExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitMajorExpr(ctx *parser.MajorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitAnyExpr(ctx *parser.AnyExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitGroup(ctx *parser.GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitGroupChoice(ctx *parser.GroupChoiceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitTypeEntry(ctx *parser.TypeEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitNameEntry(ctx *parser.NameEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitGroupsEntry(ctx *parser.GroupsEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitTypeMember(ctx *parser.TypeMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitNameMember(ctx *parser.NameMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitValueMember(ctx *parser.ValueMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitOptComma(ctx *parser.OptCommaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) defineRule(id string) {
	// TODO: report redeclaration
	v.definedRules[id] = struct{}{}
}

func (v *Validator) addRuleUse(id antlr.TerminalNode) {
	v.usedRules = append(v.usedRules,
		RuleUse{
			ID:       id.GetText(),
			Position: parser.PositionFromToken(id.GetSymbol()),
		},
	)
}

type UndefinedRuleUsageError struct {
	ID string
	parser.Position
}

var _ error = UndefinedRuleUsageError{}

func (e UndefinedRuleUsageError) Error() string {
	return fmt.Sprintf(
		"%d:%d: use of undefined rule: %s",
		e.Line,
		e.Column,
		e.ID,
	)
}

func Validate(cddl parser.ICddlContext) error {
	validator := &Validator{
		definedRules: map[string]struct{}{},
	}

	cddl.Accept(validator)

	for _, ruleUse := range validator.usedRules {
		id := ruleUse.ID

		_, ok := validator.definedRules[id]
		if ok {
			continue
		}

		_, ok = predeclaredIDs[id]
		if ok {
			continue
		}

		// TODO: report all
		return UndefinedRuleUsageError{
			ID:       id,
			Position: ruleUse.Position,
		}
	}

	return nil
}
