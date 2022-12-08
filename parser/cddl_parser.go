// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // CDDL
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CDDLParser struct {
	*antlr.BaseParser
}

var cddlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cddlParserInit() {
	staticData := &cddlParserStaticData
	staticData.literalNames = []string{
		"", "'='", "'/='", "'//='", "'<'", "','", "'>'", "'/'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "'~'", "'&'", "'#'", "'//'", "'^'", "'=>'",
		"':'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "TAG", "MAJOR", "RANGEOP", "CTLOP", "OCCUR", "VALUE",
		"ID", "S",
	}
	staticData.ruleNames = []string{
		"cddl", "rule", "assignType", "assignGroup", "genericParam", "genericArg",
		"type", "type1", "type2", "group", "groupChoice", "groupEntry", "memberKey",
		"optComma",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 186, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 4, 0, 30, 8, 0, 11,
		0, 12, 0, 31, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 38, 8, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 45, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 50, 8, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 60, 8, 4, 10, 4, 12, 4, 63,
		9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 71, 8, 5, 10, 5, 12, 5,
		74, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 81, 8, 6, 10, 6, 12, 6, 84,
		9, 6, 1, 7, 1, 7, 1, 7, 3, 7, 89, 8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 94, 8,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 3, 8, 111, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 3, 8, 121, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 3, 8, 130, 8, 8, 1, 9, 1, 9, 1, 9, 5, 9, 135, 8, 9, 10, 9, 12, 9, 138,
		9, 9, 1, 10, 1, 10, 1, 10, 5, 10, 143, 8, 10, 10, 10, 12, 10, 146, 9, 10,
		1, 11, 3, 11, 149, 8, 11, 1, 11, 3, 11, 152, 8, 11, 1, 11, 1, 11, 3, 11,
		156, 8, 11, 1, 11, 1, 11, 3, 11, 160, 8, 11, 1, 11, 3, 11, 163, 8, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 169, 8, 11, 1, 12, 1, 12, 3, 12, 173,
		8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 181, 8, 12, 1,
		13, 3, 13, 184, 8, 13, 1, 13, 0, 0, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 0, 3, 1, 0, 1, 2, 2, 0, 1, 1, 3, 3, 1, 0, 23, 24, 205,
		0, 29, 1, 0, 0, 0, 2, 49, 1, 0, 0, 0, 4, 51, 1, 0, 0, 0, 6, 53, 1, 0, 0,
		0, 8, 55, 1, 0, 0, 0, 10, 66, 1, 0, 0, 0, 12, 77, 1, 0, 0, 0, 14, 85, 1,
		0, 0, 0, 16, 129, 1, 0, 0, 0, 18, 131, 1, 0, 0, 0, 20, 144, 1, 0, 0, 0,
		22, 168, 1, 0, 0, 0, 24, 180, 1, 0, 0, 0, 26, 183, 1, 0, 0, 0, 28, 30,
		3, 2, 1, 0, 29, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0,
		31, 32, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 34, 5, 0, 0, 1, 34, 1, 1, 0,
		0, 0, 35, 37, 5, 27, 0, 0, 36, 38, 3, 8, 4, 0, 37, 36, 1, 0, 0, 0, 37,
		38, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 3, 4, 2, 0, 40, 41, 3, 12,
		6, 0, 41, 50, 1, 0, 0, 0, 42, 44, 5, 27, 0, 0, 43, 45, 3, 8, 4, 0, 44,
		43, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 3, 6, 3,
		0, 47, 48, 3, 22, 11, 0, 48, 50, 1, 0, 0, 0, 49, 35, 1, 0, 0, 0, 49, 42,
		1, 0, 0, 0, 50, 3, 1, 0, 0, 0, 51, 52, 7, 0, 0, 0, 52, 5, 1, 0, 0, 0, 53,
		54, 7, 1, 0, 0, 54, 7, 1, 0, 0, 0, 55, 56, 5, 4, 0, 0, 56, 61, 5, 27, 0,
		0, 57, 58, 5, 5, 0, 0, 58, 60, 5, 27, 0, 0, 59, 57, 1, 0, 0, 0, 60, 63,
		1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 64, 1, 0, 0, 0,
		63, 61, 1, 0, 0, 0, 64, 65, 5, 6, 0, 0, 65, 9, 1, 0, 0, 0, 66, 67, 5, 4,
		0, 0, 67, 72, 3, 14, 7, 0, 68, 69, 5, 5, 0, 0, 69, 71, 3, 14, 7, 0, 70,
		68, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0,
		0, 73, 75, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 76, 5, 6, 0, 0, 76, 11,
		1, 0, 0, 0, 77, 82, 3, 14, 7, 0, 78, 79, 5, 7, 0, 0, 79, 81, 3, 14, 7,
		0, 80, 78, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83,
		1, 0, 0, 0, 83, 13, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 88, 3, 16, 8, 0,
		86, 87, 7, 2, 0, 0, 87, 89, 3, 16, 8, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1,
		0, 0, 0, 89, 15, 1, 0, 0, 0, 90, 130, 5, 26, 0, 0, 91, 93, 5, 27, 0, 0,
		92, 94, 3, 10, 5, 0, 93, 92, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 130, 1,
		0, 0, 0, 95, 96, 5, 8, 0, 0, 96, 97, 3, 12, 6, 0, 97, 98, 5, 9, 0, 0, 98,
		130, 1, 0, 0, 0, 99, 100, 5, 10, 0, 0, 100, 101, 3, 18, 9, 0, 101, 102,
		5, 11, 0, 0, 102, 130, 1, 0, 0, 0, 103, 104, 5, 12, 0, 0, 104, 105, 3,
		18, 9, 0, 105, 106, 5, 13, 0, 0, 106, 130, 1, 0, 0, 0, 107, 108, 5, 14,
		0, 0, 108, 110, 5, 27, 0, 0, 109, 111, 3, 10, 5, 0, 110, 109, 1, 0, 0,
		0, 110, 111, 1, 0, 0, 0, 111, 130, 1, 0, 0, 0, 112, 113, 5, 15, 0, 0, 113,
		114, 5, 8, 0, 0, 114, 115, 3, 18, 9, 0, 115, 116, 5, 9, 0, 0, 116, 130,
		1, 0, 0, 0, 117, 118, 5, 15, 0, 0, 118, 120, 5, 27, 0, 0, 119, 121, 3,
		10, 5, 0, 120, 119, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 130, 1, 0, 0,
		0, 122, 123, 5, 21, 0, 0, 123, 124, 5, 8, 0, 0, 124, 125, 3, 12, 6, 0,
		125, 126, 5, 9, 0, 0, 126, 130, 1, 0, 0, 0, 127, 130, 5, 22, 0, 0, 128,
		130, 5, 16, 0, 0, 129, 90, 1, 0, 0, 0, 129, 91, 1, 0, 0, 0, 129, 95, 1,
		0, 0, 0, 129, 99, 1, 0, 0, 0, 129, 103, 1, 0, 0, 0, 129, 107, 1, 0, 0,
		0, 129, 112, 1, 0, 0, 0, 129, 117, 1, 0, 0, 0, 129, 122, 1, 0, 0, 0, 129,
		127, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 17, 1, 0, 0, 0, 131, 136, 3,
		20, 10, 0, 132, 133, 5, 17, 0, 0, 133, 135, 3, 20, 10, 0, 134, 132, 1,
		0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0,
		0, 137, 19, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 140, 3, 22, 11, 0, 140,
		141, 3, 26, 13, 0, 141, 143, 1, 0, 0, 0, 142, 139, 1, 0, 0, 0, 143, 146,
		1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 21, 1, 0,
		0, 0, 146, 144, 1, 0, 0, 0, 147, 149, 5, 25, 0, 0, 148, 147, 1, 0, 0, 0,
		148, 149, 1, 0, 0, 0, 149, 151, 1, 0, 0, 0, 150, 152, 3, 24, 12, 0, 151,
		150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 169,
		3, 12, 6, 0, 154, 156, 5, 25, 0, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1,
		0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 159, 5, 27, 0, 0, 158, 160, 3, 10,
		5, 0, 159, 158, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 169, 1, 0, 0, 0,
		161, 163, 5, 25, 0, 0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163,
		164, 1, 0, 0, 0, 164, 165, 5, 8, 0, 0, 165, 166, 3, 18, 9, 0, 166, 167,
		5, 9, 0, 0, 167, 169, 1, 0, 0, 0, 168, 148, 1, 0, 0, 0, 168, 155, 1, 0,
		0, 0, 168, 162, 1, 0, 0, 0, 169, 23, 1, 0, 0, 0, 170, 172, 3, 14, 7, 0,
		171, 173, 5, 18, 0, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173,
		174, 1, 0, 0, 0, 174, 175, 5, 19, 0, 0, 175, 181, 1, 0, 0, 0, 176, 177,
		5, 27, 0, 0, 177, 181, 5, 20, 0, 0, 178, 179, 5, 26, 0, 0, 179, 181, 5,
		20, 0, 0, 180, 170, 1, 0, 0, 0, 180, 176, 1, 0, 0, 0, 180, 178, 1, 0, 0,
		0, 181, 25, 1, 0, 0, 0, 182, 184, 5, 5, 0, 0, 183, 182, 1, 0, 0, 0, 183,
		184, 1, 0, 0, 0, 184, 27, 1, 0, 0, 0, 23, 31, 37, 44, 49, 61, 72, 82, 88,
		93, 110, 120, 129, 136, 144, 148, 151, 155, 159, 162, 168, 172, 180, 183,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CDDLParserInit initializes any static state used to implement CDDLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCDDLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CDDLParserInit() {
	staticData := &cddlParserStaticData
	staticData.once.Do(cddlParserInit)
}

// NewCDDLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCDDLParser(input antlr.TokenStream) *CDDLParser {
	CDDLParserInit()
	this := new(CDDLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cddlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// CDDLParser tokens.
const (
	CDDLParserEOF     = antlr.TokenEOF
	CDDLParserT__0    = 1
	CDDLParserT__1    = 2
	CDDLParserT__2    = 3
	CDDLParserT__3    = 4
	CDDLParserT__4    = 5
	CDDLParserT__5    = 6
	CDDLParserT__6    = 7
	CDDLParserT__7    = 8
	CDDLParserT__8    = 9
	CDDLParserT__9    = 10
	CDDLParserT__10   = 11
	CDDLParserT__11   = 12
	CDDLParserT__12   = 13
	CDDLParserT__13   = 14
	CDDLParserT__14   = 15
	CDDLParserT__15   = 16
	CDDLParserT__16   = 17
	CDDLParserT__17   = 18
	CDDLParserT__18   = 19
	CDDLParserT__19   = 20
	CDDLParserTAG     = 21
	CDDLParserMAJOR   = 22
	CDDLParserRANGEOP = 23
	CDDLParserCTLOP   = 24
	CDDLParserOCCUR   = 25
	CDDLParserVALUE   = 26
	CDDLParserID      = 27
	CDDLParserS       = 28
)

// CDDLParser rules.
const (
	CDDLParserRULE_cddl         = 0
	CDDLParserRULE_rule         = 1
	CDDLParserRULE_assignType   = 2
	CDDLParserRULE_assignGroup  = 3
	CDDLParserRULE_genericParam = 4
	CDDLParserRULE_genericArg   = 5
	CDDLParserRULE_type         = 6
	CDDLParserRULE_type1        = 7
	CDDLParserRULE_type2        = 8
	CDDLParserRULE_group        = 9
	CDDLParserRULE_groupChoice  = 10
	CDDLParserRULE_groupEntry   = 11
	CDDLParserRULE_memberKey    = 12
	CDDLParserRULE_optComma     = 13
)

// ICddlContext is an interface to support dynamic dispatch.
type ICddlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_rule returns the _rule rule contexts.
	Get_rule() IRuleContext

	// Set_rule sets the _rule rule contexts.
	Set_rule(IRuleContext)

	// GetRules returns the rules rule context list.
	GetRules() []IRuleContext

	// SetRules sets the rules rule context list.
	SetRules([]IRuleContext)

	// IsCddlContext differentiates from other interfaces.
	IsCddlContext()
}

type CddlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_rule  IRuleContext
	rules  []IRuleContext
}

func NewEmptyCddlContext() *CddlContext {
	var p = new(CddlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_cddl
	return p
}

func (*CddlContext) IsCddlContext() {}

func NewCddlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CddlContext {
	var p = new(CddlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_cddl

	return p
}

func (s *CddlContext) GetParser() antlr.Parser { return s.parser }

func (s *CddlContext) Get_rule() IRuleContext { return s._rule }

func (s *CddlContext) Set_rule(v IRuleContext) { s._rule = v }

func (s *CddlContext) GetRules() []IRuleContext { return s.rules }

func (s *CddlContext) SetRules(v []IRuleContext) { s.rules = v }

func (s *CddlContext) EOF() antlr.TerminalNode {
	return s.GetToken(CDDLParserEOF, 0)
}

func (s *CddlContext) AllRule_() []IRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleContext); ok {
			len++
		}
	}

	tst := make([]IRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleContext); ok {
			tst[i] = t.(IRuleContext)
			i++
		}
	}

	return tst
}

func (s *CddlContext) Rule_(i int) IRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleContext)
}

func (s *CddlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CddlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CddlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitCddl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Cddl() (localctx ICddlContext) {
	this := p
	_ = this

	localctx = NewCddlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CDDLParserRULE_cddl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CDDLParserID {
		{
			p.SetState(28)

			var _x = p.Rule_()

			localctx.(*CddlContext)._rule = _x
		}
		localctx.(*CddlContext).rules = append(localctx.(*CddlContext).rules, localctx.(*CddlContext)._rule)

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Match(CDDLParserEOF)
	}

	return localctx
}

// IRuleContext is an interface to support dynamic dispatch.
type IRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleContext differentiates from other interfaces.
	IsRuleContext()
}

type RuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleContext() *RuleContext {
	var p = new(RuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_rule
	return p
}

func (*RuleContext) IsRuleContext() {}

func NewRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContext {
	var p = new(RuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_rule

	return p
}

func (s *RuleContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContext) CopyFrom(ctx *RuleContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GroupRuleContext struct {
	*RuleContext
}

func NewGroupRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupRuleContext {
	var p = new(GroupRuleContext)

	p.RuleContext = NewEmptyRuleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RuleContext))

	return p
}

func (s *GroupRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupRuleContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *GroupRuleContext) AssignGroup() IAssignGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignGroupContext)
}

func (s *GroupRuleContext) GroupEntry() IGroupEntryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupEntryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupEntryContext)
}

func (s *GroupRuleContext) GenericParam() IGenericParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericParamContext)
}

func (s *GroupRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGroupRule(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeRuleContext struct {
	*RuleContext
}

func NewTypeRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeRuleContext {
	var p = new(TypeRuleContext)

	p.RuleContext = NewEmptyRuleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RuleContext))

	return p
}

func (s *TypeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRuleContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *TypeRuleContext) AssignType() IAssignTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignTypeContext)
}

func (s *TypeRuleContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeRuleContext) GenericParam() IGenericParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericParamContext)
}

func (s *TypeRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitTypeRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Rule_() (localctx IRuleContext) {
	this := p
	_ = this

	localctx = NewRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CDDLParserRULE_rule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.Match(CDDLParserID)
		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(36)
				p.GenericParam()
			}

		}
		{
			p.SetState(39)
			p.AssignType()
		}
		{
			p.SetState(40)
			p.Type_()
		}

	case 2:
		localctx = NewGroupRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(CDDLParserID)
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(43)
				p.GenericParam()
			}

		}
		{
			p.SetState(46)
			p.AssignGroup()
		}
		{
			p.SetState(47)
			p.GroupEntry()
		}

	}

	return localctx
}

// IAssignTypeContext is an interface to support dynamic dispatch.
type IAssignTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignTypeContext differentiates from other interfaces.
	IsAssignTypeContext()
}

type AssignTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignTypeContext() *AssignTypeContext {
	var p = new(AssignTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_assignType
	return p
}

func (*AssignTypeContext) IsAssignTypeContext() {}

func NewAssignTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignTypeContext {
	var p = new(AssignTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_assignType

	return p
}

func (s *AssignTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *AssignTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitAssignType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) AssignType() (localctx IAssignTypeContext) {
	this := p
	_ = this

	localctx = NewAssignTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CDDLParserRULE_assignType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CDDLParserT__0 || _la == CDDLParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssignGroupContext is an interface to support dynamic dispatch.
type IAssignGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignGroupContext differentiates from other interfaces.
	IsAssignGroupContext()
}

type AssignGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignGroupContext() *AssignGroupContext {
	var p = new(AssignGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_assignGroup
	return p
}

func (*AssignGroupContext) IsAssignGroupContext() {}

func NewAssignGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignGroupContext {
	var p = new(AssignGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_assignGroup

	return p
}

func (s *AssignGroupContext) GetParser() antlr.Parser { return s.parser }
func (s *AssignGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitAssignGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) AssignGroup() (localctx IAssignGroupContext) {
	this := p
	_ = this

	localctx = NewAssignGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CDDLParserRULE_assignGroup)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CDDLParserT__0 || _la == CDDLParserT__2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGenericParamContext is an interface to support dynamic dispatch.
type IGenericParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetIds returns the ids token list.
	GetIds() []antlr.Token

	// SetIds sets the ids token list.
	SetIds([]antlr.Token)

	// IsGenericParamContext differentiates from other interfaces.
	IsGenericParamContext()
}

type GenericParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_ID    antlr.Token
	ids    []antlr.Token
}

func NewEmptyGenericParamContext() *GenericParamContext {
	var p = new(GenericParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_genericParam
	return p
}

func (*GenericParamContext) IsGenericParamContext() {}

func NewGenericParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericParamContext {
	var p = new(GenericParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_genericParam

	return p
}

func (s *GenericParamContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericParamContext) Get_ID() antlr.Token { return s._ID }

func (s *GenericParamContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *GenericParamContext) GetIds() []antlr.Token { return s.ids }

func (s *GenericParamContext) SetIds(v []antlr.Token) { s.ids = v }

func (s *GenericParamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CDDLParserID)
}

func (s *GenericParamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CDDLParserID, i)
}

func (s *GenericParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGenericParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) GenericParam() (localctx IGenericParamContext) {
	this := p
	_ = this

	localctx = NewGenericParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CDDLParserRULE_genericParam)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(CDDLParserT__3)
	}
	{
		p.SetState(56)

		var _m = p.Match(CDDLParserID)

		localctx.(*GenericParamContext)._ID = _m
	}
	localctx.(*GenericParamContext).ids = append(localctx.(*GenericParamContext).ids, localctx.(*GenericParamContext)._ID)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDDLParserT__4 {
		{
			p.SetState(57)
			p.Match(CDDLParserT__4)
		}
		{
			p.SetState(58)

			var _m = p.Match(CDDLParserID)

			localctx.(*GenericParamContext)._ID = _m
		}
		localctx.(*GenericParamContext).ids = append(localctx.(*GenericParamContext).ids, localctx.(*GenericParamContext)._ID)

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
		p.Match(CDDLParserT__5)
	}

	return localctx
}

// IGenericArgContext is an interface to support dynamic dispatch.
type IGenericArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_type1 returns the _type1 rule contexts.
	Get_type1() IType1Context

	// Set_type1 sets the _type1 rule contexts.
	Set_type1(IType1Context)

	// GetTypes returns the types rule context list.
	GetTypes() []IType1Context

	// SetTypes sets the types rule context list.
	SetTypes([]IType1Context)

	// IsGenericArgContext differentiates from other interfaces.
	IsGenericArgContext()
}

type GenericArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_type1 IType1Context
	types  []IType1Context
}

func NewEmptyGenericArgContext() *GenericArgContext {
	var p = new(GenericArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_genericArg
	return p
}

func (*GenericArgContext) IsGenericArgContext() {}

func NewGenericArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericArgContext {
	var p = new(GenericArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_genericArg

	return p
}

func (s *GenericArgContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericArgContext) Get_type1() IType1Context { return s._type1 }

func (s *GenericArgContext) Set_type1(v IType1Context) { s._type1 = v }

func (s *GenericArgContext) GetTypes() []IType1Context { return s.types }

func (s *GenericArgContext) SetTypes(v []IType1Context) { s.types = v }

func (s *GenericArgContext) AllType1() []IType1Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType1Context); ok {
			len++
		}
	}

	tst := make([]IType1Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType1Context); ok {
			tst[i] = t.(IType1Context)
			i++
		}
	}

	return tst
}

func (s *GenericArgContext) Type1(i int) IType1Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType1Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType1Context)
}

func (s *GenericArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGenericArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) GenericArg() (localctx IGenericArgContext) {
	this := p
	_ = this

	localctx = NewGenericArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CDDLParserRULE_genericArg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(CDDLParserT__3)
	}
	{
		p.SetState(67)

		var _x = p.Type1()

		localctx.(*GenericArgContext)._type1 = _x
	}
	localctx.(*GenericArgContext).types = append(localctx.(*GenericArgContext).types, localctx.(*GenericArgContext)._type1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDDLParserT__4 {
		{
			p.SetState(68)
			p.Match(CDDLParserT__4)
		}
		{
			p.SetState(69)

			var _x = p.Type1()

			localctx.(*GenericArgContext)._type1 = _x
		}
		localctx.(*GenericArgContext).types = append(localctx.(*GenericArgContext).types, localctx.(*GenericArgContext)._type1)

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Match(CDDLParserT__5)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_type1 returns the _type1 rule contexts.
	Get_type1() IType1Context

	// Set_type1 sets the _type1 rule contexts.
	Set_type1(IType1Context)

	// GetTypes returns the types rule context list.
	GetTypes() []IType1Context

	// SetTypes sets the types rule context list.
	SetTypes([]IType1Context)

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_type1 IType1Context
	types  []IType1Context
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Get_type1() IType1Context { return s._type1 }

func (s *TypeContext) Set_type1(v IType1Context) { s._type1 = v }

func (s *TypeContext) GetTypes() []IType1Context { return s.types }

func (s *TypeContext) SetTypes(v []IType1Context) { s.types = v }

func (s *TypeContext) AllType1() []IType1Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType1Context); ok {
			len++
		}
	}

	tst := make([]IType1Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType1Context); ok {
			tst[i] = t.(IType1Context)
			i++
		}
	}

	return tst
}

func (s *TypeContext) Type1(i int) IType1Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType1Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType1Context)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CDDLParserRULE_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)

		var _x = p.Type1()

		localctx.(*TypeContext)._type1 = _x
	}
	localctx.(*TypeContext).types = append(localctx.(*TypeContext).types, localctx.(*TypeContext)._type1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDDLParserT__6 {
		{
			p.SetState(78)
			p.Match(CDDLParserT__6)
		}
		{
			p.SetState(79)

			var _x = p.Type1()

			localctx.(*TypeContext)._type1 = _x
		}
		localctx.(*TypeContext).types = append(localctx.(*TypeContext).types, localctx.(*TypeContext)._type1)

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IType1Context is an interface to support dynamic dispatch.
type IType1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType1Context differentiates from other interfaces.
	IsType1Context()
}

type Type1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType1Context() *Type1Context {
	var p = new(Type1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_type1
	return p
}

func (*Type1Context) IsType1Context() {}

func NewType1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type1Context {
	var p = new(Type1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_type1

	return p
}

func (s *Type1Context) GetParser() antlr.Parser { return s.parser }

func (s *Type1Context) AllType2() []IType2Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType2Context); ok {
			len++
		}
	}

	tst := make([]IType2Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType2Context); ok {
			tst[i] = t.(IType2Context)
			i++
		}
	}

	return tst
}

func (s *Type1Context) Type2(i int) IType2Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType2Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType2Context)
}

func (s *Type1Context) RANGEOP() antlr.TerminalNode {
	return s.GetToken(CDDLParserRANGEOP, 0)
}

func (s *Type1Context) CTLOP() antlr.TerminalNode {
	return s.GetToken(CDDLParserCTLOP, 0)
}

func (s *Type1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitType1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Type1() (localctx IType1Context) {
	this := p
	_ = this

	localctx = NewType1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CDDLParserRULE_type1)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Type2()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserRANGEOP || _la == CDDLParserCTLOP {
		{
			p.SetState(86)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CDDLParserRANGEOP || _la == CDDLParserCTLOP) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(87)
			p.Type2()
		}

	}

	return localctx
}

// IType2Context is an interface to support dynamic dispatch.
type IType2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType2Context differentiates from other interfaces.
	IsType2Context()
}

type Type2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType2Context() *Type2Context {
	var p = new(Type2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_type2
	return p
}

func (*Type2Context) IsType2Context() {}

func NewType2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type2Context {
	var p = new(Type2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_type2

	return p
}

func (s *Type2Context) GetParser() antlr.Parser { return s.parser }

func (s *Type2Context) CopyFrom(ctx *Type2Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Type2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MapExprContext struct {
	*Type2Context
}

func NewMapExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapExprContext {
	var p = new(MapExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *MapExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapExprContext) Group() IGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *MapExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitMapExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ChoiceExprContext struct {
	*Type2Context
}

func NewChoiceExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChoiceExprContext {
	var p = new(ChoiceExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *ChoiceExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceExprContext) Group() IGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *ChoiceExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitChoiceExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueExprContext struct {
	*Type2Context
	value antlr.Token
}

func NewValueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueExprContext {
	var p = new(ValueExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *ValueExprContext) GetValue() antlr.Token { return s.value }

func (s *ValueExprContext) SetValue(v antlr.Token) { s.value = v }

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) VALUE() antlr.TerminalNode {
	return s.GetToken(CDDLParserVALUE, 0)
}

func (s *ValueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitValueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupExprContext struct {
	*Type2Context
}

func NewGroupExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupExprContext {
	var p = new(GroupExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *GroupExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupExprContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *GroupExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGroupExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayExprContext struct {
	*Type2Context
}

func NewArrayExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExprContext {
	var p = new(ArrayExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *ArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExprContext) Group() IGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *ArrayExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitArrayExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MajorExprContext struct {
	*Type2Context
	major antlr.Token
}

func NewMajorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MajorExprContext {
	var p = new(MajorExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *MajorExprContext) GetMajor() antlr.Token { return s.major }

func (s *MajorExprContext) SetMajor(v antlr.Token) { s.major = v }

func (s *MajorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MajorExprContext) MAJOR() antlr.TerminalNode {
	return s.GetToken(CDDLParserMAJOR, 0)
}

func (s *MajorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitMajorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	*Type2Context
	id  antlr.Token
	arg IGenericArgContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *IdExprContext) GetId() antlr.Token { return s.id }

func (s *IdExprContext) SetId(v antlr.Token) { s.id = v }

func (s *IdExprContext) GetArg() IGenericArgContext { return s.arg }

func (s *IdExprContext) SetArg(v IGenericArgContext) { s.arg = v }

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *IdExprContext) GenericArg() IGenericArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericArgContext)
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AnyExprContext struct {
	*Type2Context
}

func NewAnyExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnyExprContext {
	var p = new(AnyExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *AnyExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitAnyExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnwrapExprContext struct {
	*Type2Context
	id  antlr.Token
	arg IGenericArgContext
}

func NewUnwrapExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnwrapExprContext {
	var p = new(UnwrapExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *UnwrapExprContext) GetId() antlr.Token { return s.id }

func (s *UnwrapExprContext) SetId(v antlr.Token) { s.id = v }

func (s *UnwrapExprContext) GetArg() IGenericArgContext { return s.arg }

func (s *UnwrapExprContext) SetArg(v IGenericArgContext) { s.arg = v }

func (s *UnwrapExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnwrapExprContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *UnwrapExprContext) GenericArg() IGenericArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericArgContext)
}

func (s *UnwrapExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitUnwrapExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ChoiceIDExprContext struct {
	*Type2Context
	id  antlr.Token
	arg IGenericArgContext
}

func NewChoiceIDExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChoiceIDExprContext {
	var p = new(ChoiceIDExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *ChoiceIDExprContext) GetId() antlr.Token { return s.id }

func (s *ChoiceIDExprContext) SetId(v antlr.Token) { s.id = v }

func (s *ChoiceIDExprContext) GetArg() IGenericArgContext { return s.arg }

func (s *ChoiceIDExprContext) SetArg(v IGenericArgContext) { s.arg = v }

func (s *ChoiceIDExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceIDExprContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *ChoiceIDExprContext) GenericArg() IGenericArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericArgContext)
}

func (s *ChoiceIDExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitChoiceIDExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TaggedExprContext struct {
	*Type2Context
	tag antlr.Token
}

func NewTaggedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TaggedExprContext {
	var p = new(TaggedExprContext)

	p.Type2Context = NewEmptyType2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type2Context))

	return p
}

func (s *TaggedExprContext) GetTag() antlr.Token { return s.tag }

func (s *TaggedExprContext) SetTag(v antlr.Token) { s.tag = v }

func (s *TaggedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TaggedExprContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TaggedExprContext) TAG() antlr.TerminalNode {
	return s.GetToken(CDDLParserTAG, 0)
}

func (s *TaggedExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitTaggedExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Type2() (localctx IType2Context) {
	this := p
	_ = this

	localctx = NewType2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CDDLParserRULE_type2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)

			var _m = p.Match(CDDLParserVALUE)

			localctx.(*ValueExprContext).value = _m
		}

	case 2:
		localctx = NewIdExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)

			var _m = p.Match(CDDLParserID)

			localctx.(*IdExprContext).id = _m
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(92)

				var _x = p.GenericArg()

				localctx.(*IdExprContext).arg = _x
			}

		}

	case 3:
		localctx = NewGroupExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.Match(CDDLParserT__7)
		}
		{
			p.SetState(96)
			p.Type_()
		}
		{
			p.SetState(97)
			p.Match(CDDLParserT__8)
		}

	case 4:
		localctx = NewMapExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(99)
			p.Match(CDDLParserT__9)
		}
		{
			p.SetState(100)
			p.Group()
		}
		{
			p.SetState(101)
			p.Match(CDDLParserT__10)
		}

	case 5:
		localctx = NewArrayExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(103)
			p.Match(CDDLParserT__11)
		}
		{
			p.SetState(104)
			p.Group()
		}
		{
			p.SetState(105)
			p.Match(CDDLParserT__12)
		}

	case 6:
		localctx = NewUnwrapExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(107)
			p.Match(CDDLParserT__13)
		}
		{
			p.SetState(108)

			var _m = p.Match(CDDLParserID)

			localctx.(*UnwrapExprContext).id = _m
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(109)

				var _x = p.GenericArg()

				localctx.(*UnwrapExprContext).arg = _x
			}

		}

	case 7:
		localctx = NewChoiceExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(112)
			p.Match(CDDLParserT__14)
		}
		{
			p.SetState(113)
			p.Match(CDDLParserT__7)
		}
		{
			p.SetState(114)
			p.Group()
		}
		{
			p.SetState(115)
			p.Match(CDDLParserT__8)
		}

	case 8:
		localctx = NewChoiceIDExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(117)
			p.Match(CDDLParserT__14)
		}
		{
			p.SetState(118)

			var _m = p.Match(CDDLParserID)

			localctx.(*ChoiceIDExprContext).id = _m
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(119)

				var _x = p.GenericArg()

				localctx.(*ChoiceIDExprContext).arg = _x
			}

		}

	case 9:
		localctx = NewTaggedExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(122)

			var _m = p.Match(CDDLParserTAG)

			localctx.(*TaggedExprContext).tag = _m
		}
		{
			p.SetState(123)
			p.Match(CDDLParserT__7)
		}
		{
			p.SetState(124)
			p.Type_()
		}
		{
			p.SetState(125)
			p.Match(CDDLParserT__8)
		}

	case 10:
		localctx = NewMajorExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(127)

			var _m = p.Match(CDDLParserMAJOR)

			localctx.(*MajorExprContext).major = _m
		}

	case 11:
		localctx = NewAnyExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(128)
			p.Match(CDDLParserT__15)
		}

	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_groupChoice returns the _groupChoice rule contexts.
	Get_groupChoice() IGroupChoiceContext

	// Set_groupChoice sets the _groupChoice rule contexts.
	Set_groupChoice(IGroupChoiceContext)

	// GetGroups returns the groups rule context list.
	GetGroups() []IGroupChoiceContext

	// SetGroups sets the groups rule context list.
	SetGroups([]IGroupChoiceContext)

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_groupChoice IGroupChoiceContext
	groups       []IGroupChoiceContext
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) Get_groupChoice() IGroupChoiceContext { return s._groupChoice }

func (s *GroupContext) Set_groupChoice(v IGroupChoiceContext) { s._groupChoice = v }

func (s *GroupContext) GetGroups() []IGroupChoiceContext { return s.groups }

func (s *GroupContext) SetGroups(v []IGroupChoiceContext) { s.groups = v }

func (s *GroupContext) AllGroupChoice() []IGroupChoiceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupChoiceContext); ok {
			len++
		}
	}

	tst := make([]IGroupChoiceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupChoiceContext); ok {
			tst[i] = t.(IGroupChoiceContext)
			i++
		}
	}

	return tst
}

func (s *GroupContext) GroupChoice(i int) IGroupChoiceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupChoiceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupChoiceContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) Group() (localctx IGroupContext) {
	this := p
	_ = this

	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CDDLParserRULE_group)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)

		var _x = p.GroupChoice()

		localctx.(*GroupContext)._groupChoice = _x
	}
	localctx.(*GroupContext).groups = append(localctx.(*GroupContext).groups, localctx.(*GroupContext)._groupChoice)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDDLParserT__16 {
		{
			p.SetState(132)
			p.Match(CDDLParserT__16)
		}
		{
			p.SetState(133)

			var _x = p.GroupChoice()

			localctx.(*GroupContext)._groupChoice = _x
		}
		localctx.(*GroupContext).groups = append(localctx.(*GroupContext).groups, localctx.(*GroupContext)._groupChoice)

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGroupChoiceContext is an interface to support dynamic dispatch.
type IGroupChoiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEntry returns the entry rule contexts.
	GetEntry() IGroupEntryContext

	// SetEntry sets the entry rule contexts.
	SetEntry(IGroupEntryContext)

	// IsGroupChoiceContext differentiates from other interfaces.
	IsGroupChoiceContext()
}

type GroupChoiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	entry  IGroupEntryContext
}

func NewEmptyGroupChoiceContext() *GroupChoiceContext {
	var p = new(GroupChoiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_groupChoice
	return p
}

func (*GroupChoiceContext) IsGroupChoiceContext() {}

func NewGroupChoiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupChoiceContext {
	var p = new(GroupChoiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_groupChoice

	return p
}

func (s *GroupChoiceContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupChoiceContext) GetEntry() IGroupEntryContext { return s.entry }

func (s *GroupChoiceContext) SetEntry(v IGroupEntryContext) { s.entry = v }

func (s *GroupChoiceContext) AllOptComma() []IOptCommaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptCommaContext); ok {
			len++
		}
	}

	tst := make([]IOptCommaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptCommaContext); ok {
			tst[i] = t.(IOptCommaContext)
			i++
		}
	}

	return tst
}

func (s *GroupChoiceContext) OptComma(i int) IOptCommaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptCommaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptCommaContext)
}

func (s *GroupChoiceContext) AllGroupEntry() []IGroupEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupEntryContext); ok {
			len++
		}
	}

	tst := make([]IGroupEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupEntryContext); ok {
			tst[i] = t.(IGroupEntryContext)
			i++
		}
	}

	return tst
}

func (s *GroupChoiceContext) GroupEntry(i int) IGroupEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupEntryContext)
}

func (s *GroupChoiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupChoiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupChoiceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGroupChoice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) GroupChoice() (localctx IGroupChoiceContext) {
	this := p
	_ = this

	localctx = NewGroupChoiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CDDLParserRULE_groupChoice)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&241292544) != 0 {
		{
			p.SetState(139)

			var _x = p.GroupEntry()

			localctx.(*GroupChoiceContext).entry = _x
		}
		{
			p.SetState(140)
			p.OptComma()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGroupEntryContext is an interface to support dynamic dispatch.
type IGroupEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupEntryContext differentiates from other interfaces.
	IsGroupEntryContext()
}

type GroupEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupEntryContext() *GroupEntryContext {
	var p = new(GroupEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_groupEntry
	return p
}

func (*GroupEntryContext) IsGroupEntryContext() {}

func NewGroupEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupEntryContext {
	var p = new(GroupEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_groupEntry

	return p
}

func (s *GroupEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupEntryContext) CopyFrom(ctx *GroupEntryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *GroupEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GroupsEntryContext struct {
	*GroupEntryContext
}

func NewGroupsEntryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupsEntryContext {
	var p = new(GroupsEntryContext)

	p.GroupEntryContext = NewEmptyGroupEntryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GroupEntryContext))

	return p
}

func (s *GroupsEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupsEntryContext) Group() IGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *GroupsEntryContext) OCCUR() antlr.TerminalNode {
	return s.GetToken(CDDLParserOCCUR, 0)
}

func (s *GroupsEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitGroupsEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

type NameEntryContext struct {
	*GroupEntryContext
	id  antlr.Token
	arg IGenericArgContext
}

func NewNameEntryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameEntryContext {
	var p = new(NameEntryContext)

	p.GroupEntryContext = NewEmptyGroupEntryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GroupEntryContext))

	return p
}

func (s *NameEntryContext) GetId() antlr.Token { return s.id }

func (s *NameEntryContext) SetId(v antlr.Token) { s.id = v }

func (s *NameEntryContext) GetArg() IGenericArgContext { return s.arg }

func (s *NameEntryContext) SetArg(v IGenericArgContext) { s.arg = v }

func (s *NameEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameEntryContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *NameEntryContext) OCCUR() antlr.TerminalNode {
	return s.GetToken(CDDLParserOCCUR, 0)
}

func (s *NameEntryContext) GenericArg() IGenericArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericArgContext)
}

func (s *NameEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitNameEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeEntryContext struct {
	*GroupEntryContext
}

func NewTypeEntryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeEntryContext {
	var p = new(TypeEntryContext)

	p.GroupEntryContext = NewEmptyGroupEntryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GroupEntryContext))

	return p
}

func (s *TypeEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeEntryContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeEntryContext) OCCUR() antlr.TerminalNode {
	return s.GetToken(CDDLParserOCCUR, 0)
}

func (s *TypeEntryContext) MemberKey() IMemberKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberKeyContext)
}

func (s *TypeEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitTypeEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) GroupEntry() (localctx IGroupEntryContext) {
	this := p
	_ = this

	localctx = NewGroupEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CDDLParserRULE_groupEntry)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserOCCUR {
			{
				p.SetState(147)
				p.Match(CDDLParserOCCUR)
			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(150)
				p.MemberKey()
			}

		}
		{
			p.SetState(153)
			p.Type_()
		}

	case 2:
		localctx = NewNameEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserOCCUR {
			{
				p.SetState(154)
				p.Match(CDDLParserOCCUR)
			}

		}
		{
			p.SetState(157)

			var _m = p.Match(CDDLParserID)

			localctx.(*NameEntryContext).id = _m
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__3 {
			{
				p.SetState(158)

				var _x = p.GenericArg()

				localctx.(*NameEntryContext).arg = _x
			}

		}

	case 3:
		localctx = NewGroupsEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserOCCUR {
			{
				p.SetState(161)
				p.Match(CDDLParserOCCUR)
			}

		}
		{
			p.SetState(164)
			p.Match(CDDLParserT__7)
		}
		{
			p.SetState(165)
			p.Group()
		}
		{
			p.SetState(166)
			p.Match(CDDLParserT__8)
		}

	}

	return localctx
}

// IMemberKeyContext is an interface to support dynamic dispatch.
type IMemberKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberKeyContext differentiates from other interfaces.
	IsMemberKeyContext()
}

type MemberKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberKeyContext() *MemberKeyContext {
	var p = new(MemberKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_memberKey
	return p
}

func (*MemberKeyContext) IsMemberKeyContext() {}

func NewMemberKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberKeyContext {
	var p = new(MemberKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_memberKey

	return p
}

func (s *MemberKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberKeyContext) CopyFrom(ctx *MemberKeyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MemberKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeMemberContext struct {
	*MemberKeyContext
	cut antlr.Token
}

func NewTypeMemberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeMemberContext {
	var p = new(TypeMemberContext)

	p.MemberKeyContext = NewEmptyMemberKeyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MemberKeyContext))

	return p
}

func (s *TypeMemberContext) GetCut() antlr.Token { return s.cut }

func (s *TypeMemberContext) SetCut(v antlr.Token) { s.cut = v }

func (s *TypeMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeMemberContext) Type1() IType1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType1Context)
}

func (s *TypeMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitTypeMember(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueMemberContext struct {
	*MemberKeyContext
	value antlr.Token
}

func NewValueMemberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueMemberContext {
	var p = new(ValueMemberContext)

	p.MemberKeyContext = NewEmptyMemberKeyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MemberKeyContext))

	return p
}

func (s *ValueMemberContext) GetValue() antlr.Token { return s.value }

func (s *ValueMemberContext) SetValue(v antlr.Token) { s.value = v }

func (s *ValueMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueMemberContext) VALUE() antlr.TerminalNode {
	return s.GetToken(CDDLParserVALUE, 0)
}

func (s *ValueMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitValueMember(s)

	default:
		return t.VisitChildren(s)
	}
}

type NameMemberContext struct {
	*MemberKeyContext
	id antlr.Token
}

func NewNameMemberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameMemberContext {
	var p = new(NameMemberContext)

	p.MemberKeyContext = NewEmptyMemberKeyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MemberKeyContext))

	return p
}

func (s *NameMemberContext) GetId() antlr.Token { return s.id }

func (s *NameMemberContext) SetId(v antlr.Token) { s.id = v }

func (s *NameMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameMemberContext) ID() antlr.TerminalNode {
	return s.GetToken(CDDLParserID, 0)
}

func (s *NameMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitNameMember(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) MemberKey() (localctx IMemberKeyContext) {
	this := p
	_ = this

	localctx = NewMemberKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CDDLParserRULE_memberKey)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeMemberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Type1()
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CDDLParserT__17 {
			{
				p.SetState(171)

				var _m = p.Match(CDDLParserT__17)

				localctx.(*TypeMemberContext).cut = _m
			}

		}
		{
			p.SetState(174)
			p.Match(CDDLParserT__18)
		}

	case 2:
		localctx = NewNameMemberContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)

			var _m = p.Match(CDDLParserID)

			localctx.(*NameMemberContext).id = _m
		}
		{
			p.SetState(177)
			p.Match(CDDLParserT__19)
		}

	case 3:
		localctx = NewValueMemberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(178)

			var _m = p.Match(CDDLParserVALUE)

			localctx.(*ValueMemberContext).value = _m
		}
		{
			p.SetState(179)
			p.Match(CDDLParserT__19)
		}

	}

	return localctx
}

// IOptCommaContext is an interface to support dynamic dispatch.
type IOptCommaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptCommaContext differentiates from other interfaces.
	IsOptCommaContext()
}

type OptCommaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptCommaContext() *OptCommaContext {
	var p = new(OptCommaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDDLParserRULE_optComma
	return p
}

func (*OptCommaContext) IsOptCommaContext() {}

func NewOptCommaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptCommaContext {
	var p = new(OptCommaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDDLParserRULE_optComma

	return p
}

func (s *OptCommaContext) GetParser() antlr.Parser { return s.parser }
func (s *OptCommaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptCommaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptCommaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDDLVisitor:
		return t.VisitOptComma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDDLParser) OptComma() (localctx IOptCommaContext) {
	this := p
	_ = this

	localctx = NewOptCommaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CDDLParserRULE_optComma)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDDLParserT__4 {
		{
			p.SetState(182)
			p.Match(CDDLParserT__4)
		}

	}

	return localctx
}
