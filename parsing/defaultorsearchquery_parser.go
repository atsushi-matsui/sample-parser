// Code generated from DefaultOrSearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // DefaultOrSearchQuery
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DefaultOrSearchQueryParser struct {
	*antlr.BaseParser
}

var DefaultOrSearchQueryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func defaultorsearchqueryParserInit() {
	staticData := &DefaultOrSearchQueryParserStaticData
	staticData.LiteralNames = []string{
		"", "'AND'", "'OR'", "'NOT'", "", "", "'\"'",
	}
	staticData.SymbolicNames = []string{
		"", "AND", "OR", "NOT", "WHITE_SPACE", "KEYWORD_CHARACTER", "DOUBLE_QUOTE",
		"PHRASE", "NL",
	}
	staticData.RuleNames = []string{
		"andExpression", "orExpression", "notExpression", "phrase", "keyword",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 86, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 4, 0, 16, 8, 0, 11, 0, 12, 0, 17, 1, 0, 1,
		0, 4, 0, 22, 8, 0, 11, 0, 12, 0, 23, 1, 0, 5, 0, 27, 8, 0, 10, 0, 12, 0,
		30, 9, 0, 1, 1, 1, 1, 1, 1, 4, 1, 35, 8, 1, 11, 1, 12, 1, 36, 1, 1, 1,
		1, 1, 1, 4, 1, 42, 8, 1, 11, 1, 12, 1, 43, 3, 1, 46, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 52, 8, 1, 1, 1, 5, 1, 55, 8, 1, 10, 1, 12, 1, 58, 9,
		1, 1, 2, 1, 2, 1, 2, 4, 2, 63, 8, 2, 11, 2, 12, 2, 64, 1, 2, 3, 2, 68,
		8, 2, 1, 3, 1, 3, 3, 3, 72, 8, 3, 1, 4, 4, 4, 75, 8, 4, 11, 4, 12, 4, 76,
		1, 4, 4, 4, 80, 8, 4, 11, 4, 12, 4, 81, 3, 4, 84, 8, 4, 1, 4, 0, 2, 0,
		2, 5, 0, 2, 4, 6, 8, 0, 0, 95, 0, 10, 1, 0, 0, 0, 2, 45, 1, 0, 0, 0, 4,
		67, 1, 0, 0, 0, 6, 71, 1, 0, 0, 0, 8, 83, 1, 0, 0, 0, 10, 11, 6, 0, -1,
		0, 11, 12, 3, 2, 1, 0, 12, 28, 1, 0, 0, 0, 13, 15, 10, 1, 0, 0, 14, 16,
		5, 4, 0, 0, 15, 14, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0,
		17, 18, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 21, 5, 1, 0, 0, 20, 22, 5,
		4, 0, 0, 21, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23,
		24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 27, 3, 2, 1, 0, 26, 13, 1, 0, 0,
		0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 1, 1,
		0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 32, 6, 1, -1, 0, 32, 46, 3, 4, 2, 0, 33,
		35, 5, 5, 0, 0, 34, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0,
		0, 36, 37, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 46, 5, 7, 0, 0, 39, 41,
		5, 7, 0, 0, 40, 42, 5, 5, 0, 0, 41, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0,
		43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 46, 1, 0, 0, 0, 45, 31, 1,
		0, 0, 0, 45, 34, 1, 0, 0, 0, 45, 39, 1, 0, 0, 0, 46, 56, 1, 0, 0, 0, 47,
		48, 10, 3, 0, 0, 48, 51, 5, 4, 0, 0, 49, 50, 5, 2, 0, 0, 50, 52, 5, 4,
		0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 55,
		3, 4, 2, 0, 54, 47, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0,
		56, 57, 1, 0, 0, 0, 57, 3, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 68, 3, 6,
		3, 0, 60, 62, 5, 3, 0, 0, 61, 63, 5, 4, 0, 0, 62, 61, 1, 0, 0, 0, 63, 64,
		1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0,
		66, 68, 3, 6, 3, 0, 67, 59, 1, 0, 0, 0, 67, 60, 1, 0, 0, 0, 68, 5, 1, 0,
		0, 0, 69, 72, 3, 8, 4, 0, 70, 72, 5, 7, 0, 0, 71, 69, 1, 0, 0, 0, 71, 70,
		1, 0, 0, 0, 72, 7, 1, 0, 0, 0, 73, 75, 5, 5, 0, 0, 74, 73, 1, 0, 0, 0,
		75, 76, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 84, 1,
		0, 0, 0, 78, 80, 5, 6, 0, 0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81,
		79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 84, 1, 0, 0, 0, 83, 74, 1, 0, 0,
		0, 83, 79, 1, 0, 0, 0, 84, 9, 1, 0, 0, 0, 14, 17, 23, 28, 36, 43, 45, 51,
		56, 64, 67, 71, 76, 81, 83,
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

// DefaultOrSearchQueryParserInit initializes any static state used to implement DefaultOrSearchQueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDefaultOrSearchQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DefaultOrSearchQueryParserInit() {
	staticData := &DefaultOrSearchQueryParserStaticData
	staticData.once.Do(defaultorsearchqueryParserInit)
}

// NewDefaultOrSearchQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDefaultOrSearchQueryParser(input antlr.TokenStream) *DefaultOrSearchQueryParser {
	DefaultOrSearchQueryParserInit()
	this := new(DefaultOrSearchQueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DefaultOrSearchQueryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "DefaultOrSearchQuery.g4"

	return this
}

// DefaultOrSearchQueryParser tokens.
const (
	DefaultOrSearchQueryParserEOF               = antlr.TokenEOF
	DefaultOrSearchQueryParserAND               = 1
	DefaultOrSearchQueryParserOR                = 2
	DefaultOrSearchQueryParserNOT               = 3
	DefaultOrSearchQueryParserWHITE_SPACE       = 4
	DefaultOrSearchQueryParserKEYWORD_CHARACTER = 5
	DefaultOrSearchQueryParserDOUBLE_QUOTE      = 6
	DefaultOrSearchQueryParserPHRASE            = 7
	DefaultOrSearchQueryParserNL                = 8
)

// DefaultOrSearchQueryParser rules.
const (
	DefaultOrSearchQueryParserRULE_andExpression = 0
	DefaultOrSearchQueryParserRULE_orExpression  = 1
	DefaultOrSearchQueryParserRULE_notExpression = 2
	DefaultOrSearchQueryParserRULE_phrase        = 3
	DefaultOrSearchQueryParserRULE_keyword       = 4
)

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OrExpression() IOrExpressionContext
	AndExpression() IAndExpressionContext
	AND() antlr.TerminalNode
	AllWHITE_SPACE() []antlr.TerminalNode
	WHITE_SPACE(i int) antlr.TerminalNode

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DefaultOrSearchQueryParserRULE_andExpression
	return p
}

func InitEmptyAndExpressionContext(p *AndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DefaultOrSearchQueryParserRULE_andExpression
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefaultOrSearchQueryParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) OrExpression() IOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *AndExpressionContext) AndExpression() IAndExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *AndExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserAND, 0)
}

func (s *AndExpressionContext) AllWHITE_SPACE() []antlr.TerminalNode {
	return s.GetTokens(DefaultOrSearchQueryParserWHITE_SPACE)
}

func (s *AndExpressionContext) WHITE_SPACE(i int) antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserWHITE_SPACE, i)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultOrSearchQueryListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultOrSearchQueryListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefaultOrSearchQueryVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefaultOrSearchQueryParser) AndExpression() (localctx IAndExpressionContext) {
	return p.andExpression(0)
}

func (p *DefaultOrSearchQueryParser) andExpression(_p int) (localctx IAndExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, DefaultOrSearchQueryParserRULE_andExpression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(11)
		p.orExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, DefaultOrSearchQueryParserRULE_andExpression)
			p.SetState(13)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(15)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == DefaultOrSearchQueryParserWHITE_SPACE {
				{
					p.SetState(14)
					p.Match(DefaultOrSearchQueryParserWHITE_SPACE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(17)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(19)
				p.Match(DefaultOrSearchQueryParserAND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(21)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == DefaultOrSearchQueryParserWHITE_SPACE {
				{
					p.SetState(20)
					p.Match(DefaultOrSearchQueryParserWHITE_SPACE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(23)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(25)
				p.orExpression(0)
			}

		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NotExpression() INotExpressionContext
	PHRASE() antlr.TerminalNode
	AllKEYWORD_CHARACTER() []antlr.TerminalNode
	KEYWORD_CHARACTER(i int) antlr.TerminalNode
	OrExpression() IOrExpressionContext
	AllWHITE_SPACE() []antlr.TerminalNode
	WHITE_SPACE(i int) antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DefaultOrSearchQueryParserRULE_orExpression
	return p
}

func InitEmptyOrExpressionContext(p *OrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DefaultOrSearchQueryParserRULE_orExpression
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefaultOrSearchQueryParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) NotExpression() INotExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotExpressionContext)
}

func (s *OrExpressionContext) PHRASE() antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserPHRASE, 0)
}

func (s *OrExpressionContext) AllKEYWORD_CHARACTER() []antlr.TerminalNode {
	return s.GetTokens(DefaultOrSearchQueryParserKEYWORD_CHARACTER)
}

func (s *OrExpressionContext) KEYWORD_CHARACTER(i int) antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserKEYWORD_CHARACTER, i)
}

func (s *OrExpressionContext) OrExpression() IOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *OrExpressionContext) AllWHITE_SPACE() []antlr.TerminalNode {
	return s.GetTokens(DefaultOrSearchQueryParserWHITE_SPACE)
}

func (s *OrExpressionContext) WHITE_SPACE(i int) antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserWHITE_SPACE, i)
}

func (s *OrExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserOR, 0)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultOrSearchQueryListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultOrSearchQueryListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefaultOrSearchQueryVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefaultOrSearchQueryParser) OrExpression() (localctx IOrExpressionContext) {
	return p.orExpression(0)
}

func (p *DefaultOrSearchQueryParser) orExpression(_p int) (localctx IOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, DefaultOrSearchQueryParserRULE_orExpression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(32)
			p.NotExpression()
		}

	case 2:
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DefaultOrSearchQueryParserKEYWORD_CHARACTER {
			{
				p.SetState(33)
				p.Match(DefaultOrSearchQueryParserKEYWORD_CHARACTER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(36)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(38)
			p.Match(DefaultOrSearchQueryParserPHRASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(39)
			p.Match(DefaultOrSearchQueryParserPHRASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(40)
					p.Match(DefaultOrSearchQueryParserKEYWORD_CHARACTER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, DefaultOrSearchQueryParserRULE_orExpression)
			p.SetState(47)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(48)
				p.Match(DefaultOrSearchQueryParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == DefaultOrSearchQueryParserOR {
				{
					p.SetState(49)
					p.Match(DefaultOrSearchQueryParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(50)
					p.Match(DefaultOrSearchQueryParserWHITE_SPACE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(53)
				p.NotExpression()
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INotExpressionContext is an interface to support dynamic dispatch.
type INotExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Phrase() IPhraseContext
	NOT() antlr.TerminalNode
	AllWHITE_SPACE() []antlr.TerminalNode
	WHITE_SPACE(i int) antlr.TerminalNode

	// IsNotExpressionContext differentiates from other interfaces.
	IsNotExpressionContext()
}

type NotExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotExpressionContext() *NotExpressionContext {
	var p = new(NotExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DefaultOrSearchQueryParserRULE_notExpression
	return p
}

func InitEmptyNotExpressionContext(p *NotExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DefaultOrSearchQueryParserRULE_notExpression
}

func (*NotExpressionContext) IsNotExpressionContext() {}

func NewNotExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotExpressionContext {
	var p = new(NotExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefaultOrSearchQueryParserRULE_notExpression

	return p
}

func (s *NotExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NotExpressionContext) Phrase() IPhraseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPhraseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPhraseContext)
}

func (s *NotExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserNOT, 0)
}

func (s *NotExpressionContext) AllWHITE_SPACE() []antlr.TerminalNode {
	return s.GetTokens(DefaultOrSearchQueryParserWHITE_SPACE)
}

func (s *NotExpressionContext) WHITE_SPACE(i int) antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserWHITE_SPACE, i)
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultOrSearchQueryListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultOrSearchQueryListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

func (s *NotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefaultOrSearchQueryVisitor:
		return t.VisitNotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefaultOrSearchQueryParser) NotExpression() (localctx INotExpressionContext) {
	localctx = NewNotExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DefaultOrSearchQueryParserRULE_notExpression)
	var _la int

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DefaultOrSearchQueryParserKEYWORD_CHARACTER, DefaultOrSearchQueryParserDOUBLE_QUOTE, DefaultOrSearchQueryParserPHRASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Phrase()
		}

	case DefaultOrSearchQueryParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Match(DefaultOrSearchQueryParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DefaultOrSearchQueryParserWHITE_SPACE {
			{
				p.SetState(61)
				p.Match(DefaultOrSearchQueryParserWHITE_SPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(66)
			p.Phrase()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPhraseContext is an interface to support dynamic dispatch.
type IPhraseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword() IKeywordContext
	PHRASE() antlr.TerminalNode

	// IsPhraseContext differentiates from other interfaces.
	IsPhraseContext()
}

type PhraseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhraseContext() *PhraseContext {
	var p = new(PhraseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DefaultOrSearchQueryParserRULE_phrase
	return p
}

func InitEmptyPhraseContext(p *PhraseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DefaultOrSearchQueryParserRULE_phrase
}

func (*PhraseContext) IsPhraseContext() {}

func NewPhraseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhraseContext {
	var p = new(PhraseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefaultOrSearchQueryParserRULE_phrase

	return p
}

func (s *PhraseContext) GetParser() antlr.Parser { return s.parser }

func (s *PhraseContext) Keyword() IKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *PhraseContext) PHRASE() antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserPHRASE, 0)
}

func (s *PhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhraseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhraseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultOrSearchQueryListener); ok {
		listenerT.EnterPhrase(s)
	}
}

func (s *PhraseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultOrSearchQueryListener); ok {
		listenerT.ExitPhrase(s)
	}
}

func (s *PhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefaultOrSearchQueryVisitor:
		return t.VisitPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefaultOrSearchQueryParser) Phrase() (localctx IPhraseContext) {
	localctx = NewPhraseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DefaultOrSearchQueryParserRULE_phrase)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DefaultOrSearchQueryParserKEYWORD_CHARACTER, DefaultOrSearchQueryParserDOUBLE_QUOTE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Keyword()
		}

	case DefaultOrSearchQueryParserPHRASE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Match(DefaultOrSearchQueryParserPHRASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllKEYWORD_CHARACTER() []antlr.TerminalNode
	KEYWORD_CHARACTER(i int) antlr.TerminalNode
	AllDOUBLE_QUOTE() []antlr.TerminalNode
	DOUBLE_QUOTE(i int) antlr.TerminalNode

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DefaultOrSearchQueryParserRULE_keyword
	return p
}

func InitEmptyKeywordContext(p *KeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DefaultOrSearchQueryParserRULE_keyword
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefaultOrSearchQueryParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) AllKEYWORD_CHARACTER() []antlr.TerminalNode {
	return s.GetTokens(DefaultOrSearchQueryParserKEYWORD_CHARACTER)
}

func (s *KeywordContext) KEYWORD_CHARACTER(i int) antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserKEYWORD_CHARACTER, i)
}

func (s *KeywordContext) AllDOUBLE_QUOTE() []antlr.TerminalNode {
	return s.GetTokens(DefaultOrSearchQueryParserDOUBLE_QUOTE)
}

func (s *KeywordContext) DOUBLE_QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(DefaultOrSearchQueryParserDOUBLE_QUOTE, i)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultOrSearchQueryListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefaultOrSearchQueryListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (s *KeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefaultOrSearchQueryVisitor:
		return t.VisitKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefaultOrSearchQueryParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DefaultOrSearchQueryParserRULE_keyword)
	var _alt int

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DefaultOrSearchQueryParserKEYWORD_CHARACTER:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(73)
					p.Match(DefaultOrSearchQueryParserKEYWORD_CHARACTER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case DefaultOrSearchQueryParserDOUBLE_QUOTE:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(78)
					p.Match(DefaultOrSearchQueryParserDOUBLE_QUOTE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *DefaultOrSearchQueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *AndExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AndExpressionContext)
		}
		return p.AndExpression_Sempred(t, predIndex)

	case 1:
		var t *OrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*OrExpressionContext)
		}
		return p.OrExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DefaultOrSearchQueryParser) AndExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *DefaultOrSearchQueryParser) OrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
