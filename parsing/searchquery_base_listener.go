// Code generated from SearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // SearchQuery
import "github.com/antlr4-go/antlr/v4"

// BaseSearchQueryListener is a complete listener for a parse tree produced by SearchQueryParser.
type BaseSearchQueryListener struct{}

var _ SearchQueryListener = &BaseSearchQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSearchQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSearchQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSearchQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSearchQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSearchQueryListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSearchQueryListener) ExitExpr(ctx *ExprContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSearchQueryListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSearchQueryListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseSearchQueryListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseSearchQueryListener) ExitFactor(ctx *FactorContext) {}

// EnterKeywords is called when production keywords is entered.
func (s *BaseSearchQueryListener) EnterKeywords(ctx *KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (s *BaseSearchQueryListener) ExitKeywords(ctx *KeywordsContext) {}

// EnterAlphabets is called when production alphabets is entered.
func (s *BaseSearchQueryListener) EnterAlphabets(ctx *AlphabetsContext) {}

// ExitAlphabets is called when production alphabets is exited.
func (s *BaseSearchQueryListener) ExitAlphabets(ctx *AlphabetsContext) {}
