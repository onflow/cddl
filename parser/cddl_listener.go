// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // CDDL
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// CDDLListener is a complete listener for a parse tree produced by CDDLParser.
type CDDLListener interface {
	antlr.ParseTreeListener

	// EnterCddl is called when entering the cddl production.
	EnterCddl(c *CddlContext)

	// EnterRule is called when entering the rule production.
	EnterRule(c *RuleContext)

	// EnterTypeRule is called when entering the typeRule production.
	EnterTypeRule(c *TypeRuleContext)

	// EnterGroupRule is called when entering the groupRule production.
	EnterGroupRule(c *GroupRuleContext)

	// EnterAssignType is called when entering the assignType production.
	EnterAssignType(c *AssignTypeContext)

	// EnterAssignGroup is called when entering the assignGroup production.
	EnterAssignGroup(c *AssignGroupContext)

	// EnterGenericParam is called when entering the genericParam production.
	EnterGenericParam(c *GenericParamContext)

	// EnterGenericArg is called when entering the genericArg production.
	EnterGenericArg(c *GenericArgContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterType1 is called when entering the type1 production.
	EnterType1(c *Type1Context)

	// EnterType2 is called when entering the type2 production.
	EnterType2(c *Type2Context)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterGroupChoice is called when entering the groupChoice production.
	EnterGroupChoice(c *GroupChoiceContext)

	// EnterGroupEntry is called when entering the groupEntry production.
	EnterGroupEntry(c *GroupEntryContext)

	// EnterMemberKey is called when entering the memberKey production.
	EnterMemberKey(c *MemberKeyContext)

	// EnterOptComma is called when entering the optComma production.
	EnterOptComma(c *OptCommaContext)

	// ExitCddl is called when exiting the cddl production.
	ExitCddl(c *CddlContext)

	// ExitRule is called when exiting the rule production.
	ExitRule(c *RuleContext)

	// ExitTypeRule is called when exiting the typeRule production.
	ExitTypeRule(c *TypeRuleContext)

	// ExitGroupRule is called when exiting the groupRule production.
	ExitGroupRule(c *GroupRuleContext)

	// ExitAssignType is called when exiting the assignType production.
	ExitAssignType(c *AssignTypeContext)

	// ExitAssignGroup is called when exiting the assignGroup production.
	ExitAssignGroup(c *AssignGroupContext)

	// ExitGenericParam is called when exiting the genericParam production.
	ExitGenericParam(c *GenericParamContext)

	// ExitGenericArg is called when exiting the genericArg production.
	ExitGenericArg(c *GenericArgContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitType1 is called when exiting the type1 production.
	ExitType1(c *Type1Context)

	// ExitType2 is called when exiting the type2 production.
	ExitType2(c *Type2Context)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitGroupChoice is called when exiting the groupChoice production.
	ExitGroupChoice(c *GroupChoiceContext)

	// ExitGroupEntry is called when exiting the groupEntry production.
	ExitGroupEntry(c *GroupEntryContext)

	// ExitMemberKey is called when exiting the memberKey production.
	ExitMemberKey(c *MemberKeyContext)

	// ExitOptComma is called when exiting the optComma production.
	ExitOptComma(c *OptCommaContext)
}
