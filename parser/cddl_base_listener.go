// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // CDDL
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseCDDLListener is a complete listener for a parse tree produced by CDDLParser.
type BaseCDDLListener struct{}

var _ CDDLListener = &BaseCDDLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCDDLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCDDLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCDDLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCDDLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCddl is called when production cddl is entered.
func (s *BaseCDDLListener) EnterCddl(ctx *CddlContext) {}

// ExitCddl is called when production cddl is exited.
func (s *BaseCDDLListener) ExitCddl(ctx *CddlContext) {}

// EnterRule is called when production rule is entered.
func (s *BaseCDDLListener) EnterRule(ctx *RuleContext) {}

// ExitRule is called when production rule is exited.
func (s *BaseCDDLListener) ExitRule(ctx *RuleContext) {}

// EnterTypeRule is called when production typeRule is entered.
func (s *BaseCDDLListener) EnterTypeRule(ctx *TypeRuleContext) {}

// ExitTypeRule is called when production typeRule is exited.
func (s *BaseCDDLListener) ExitTypeRule(ctx *TypeRuleContext) {}

// EnterGroupRule is called when production groupRule is entered.
func (s *BaseCDDLListener) EnterGroupRule(ctx *GroupRuleContext) {}

// ExitGroupRule is called when production groupRule is exited.
func (s *BaseCDDLListener) ExitGroupRule(ctx *GroupRuleContext) {}

// EnterAssignType is called when production assignType is entered.
func (s *BaseCDDLListener) EnterAssignType(ctx *AssignTypeContext) {}

// ExitAssignType is called when production assignType is exited.
func (s *BaseCDDLListener) ExitAssignType(ctx *AssignTypeContext) {}

// EnterAssignGroup is called when production assignGroup is entered.
func (s *BaseCDDLListener) EnterAssignGroup(ctx *AssignGroupContext) {}

// ExitAssignGroup is called when production assignGroup is exited.
func (s *BaseCDDLListener) ExitAssignGroup(ctx *AssignGroupContext) {}

// EnterGenericParam is called when production genericParam is entered.
func (s *BaseCDDLListener) EnterGenericParam(ctx *GenericParamContext) {}

// ExitGenericParam is called when production genericParam is exited.
func (s *BaseCDDLListener) ExitGenericParam(ctx *GenericParamContext) {}

// EnterGenericArg is called when production genericArg is entered.
func (s *BaseCDDLListener) EnterGenericArg(ctx *GenericArgContext) {}

// ExitGenericArg is called when production genericArg is exited.
func (s *BaseCDDLListener) ExitGenericArg(ctx *GenericArgContext) {}

// EnterType is called when production type is entered.
func (s *BaseCDDLListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseCDDLListener) ExitType(ctx *TypeContext) {}

// EnterType1 is called when production type1 is entered.
func (s *BaseCDDLListener) EnterType1(ctx *Type1Context) {}

// ExitType1 is called when production type1 is exited.
func (s *BaseCDDLListener) ExitType1(ctx *Type1Context) {}

// EnterType2 is called when production type2 is entered.
func (s *BaseCDDLListener) EnterType2(ctx *Type2Context) {}

// ExitType2 is called when production type2 is exited.
func (s *BaseCDDLListener) ExitType2(ctx *Type2Context) {}

// EnterGroup is called when production group is entered.
func (s *BaseCDDLListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseCDDLListener) ExitGroup(ctx *GroupContext) {}

// EnterGroupChoice is called when production groupChoice is entered.
func (s *BaseCDDLListener) EnterGroupChoice(ctx *GroupChoiceContext) {}

// ExitGroupChoice is called when production groupChoice is exited.
func (s *BaseCDDLListener) ExitGroupChoice(ctx *GroupChoiceContext) {}

// EnterGroupEntry is called when production groupEntry is entered.
func (s *BaseCDDLListener) EnterGroupEntry(ctx *GroupEntryContext) {}

// ExitGroupEntry is called when production groupEntry is exited.
func (s *BaseCDDLListener) ExitGroupEntry(ctx *GroupEntryContext) {}

// EnterMemberKey is called when production memberKey is entered.
func (s *BaseCDDLListener) EnterMemberKey(ctx *MemberKeyContext) {}

// ExitMemberKey is called when production memberKey is exited.
func (s *BaseCDDLListener) ExitMemberKey(ctx *MemberKeyContext) {}

// EnterOptComma is called when production optComma is entered.
func (s *BaseCDDLListener) EnterOptComma(ctx *OptCommaContext) {}

// ExitOptComma is called when production optComma is exited.
func (s *BaseCDDLListener) ExitOptComma(ctx *OptCommaContext) {}
