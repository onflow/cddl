// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // CDDL
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by CDDLParser.
type CDDLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CDDLParser#cddl.
	VisitCddl(ctx *CddlContext) interface{}

	// Visit a parse tree produced by CDDLParser#TypeRule.
	VisitTypeRule(ctx *TypeRuleContext) interface{}

	// Visit a parse tree produced by CDDLParser#GroupRule.
	VisitGroupRule(ctx *GroupRuleContext) interface{}

	// Visit a parse tree produced by CDDLParser#assignType.
	VisitAssignType(ctx *AssignTypeContext) interface{}

	// Visit a parse tree produced by CDDLParser#assignGroup.
	VisitAssignGroup(ctx *AssignGroupContext) interface{}

	// Visit a parse tree produced by CDDLParser#genericParam.
	VisitGenericParam(ctx *GenericParamContext) interface{}

	// Visit a parse tree produced by CDDLParser#genericArg.
	VisitGenericArg(ctx *GenericArgContext) interface{}

	// Visit a parse tree produced by CDDLParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by CDDLParser#type1.
	VisitType1(ctx *Type1Context) interface{}

	// Visit a parse tree produced by CDDLParser#ValueExpr.
	VisitValueExpr(ctx *ValueExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#GroupExpr.
	VisitGroupExpr(ctx *GroupExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#MapExpr.
	VisitMapExpr(ctx *MapExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#ArrayExpr.
	VisitArrayExpr(ctx *ArrayExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#UnwrapExpr.
	VisitUnwrapExpr(ctx *UnwrapExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#ChoiceExpr.
	VisitChoiceExpr(ctx *ChoiceExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#ChoiceIDExpr.
	VisitChoiceIDExpr(ctx *ChoiceIDExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#TaggedExpr.
	VisitTaggedExpr(ctx *TaggedExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#MajorExpr.
	VisitMajorExpr(ctx *MajorExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#AnyExpr.
	VisitAnyExpr(ctx *AnyExprContext) interface{}

	// Visit a parse tree produced by CDDLParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by CDDLParser#groupChoice.
	VisitGroupChoice(ctx *GroupChoiceContext) interface{}

	// Visit a parse tree produced by CDDLParser#TypeEntry.
	VisitTypeEntry(ctx *TypeEntryContext) interface{}

	// Visit a parse tree produced by CDDLParser#NameEntry.
	VisitNameEntry(ctx *NameEntryContext) interface{}

	// Visit a parse tree produced by CDDLParser#GroupsEntry.
	VisitGroupsEntry(ctx *GroupsEntryContext) interface{}

	// Visit a parse tree produced by CDDLParser#TypeMember.
	VisitTypeMember(ctx *TypeMemberContext) interface{}

	// Visit a parse tree produced by CDDLParser#NameMember.
	VisitNameMember(ctx *NameMemberContext) interface{}

	// Visit a parse tree produced by CDDLParser#ValueMember.
	VisitValueMember(ctx *ValueMemberContext) interface{}

	// Visit a parse tree produced by CDDLParser#optComma.
	VisitOptComma(ctx *OptCommaContext) interface{}
}
