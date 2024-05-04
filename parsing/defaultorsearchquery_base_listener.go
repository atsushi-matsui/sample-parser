// Code generated from DefaultOrSearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // DefaultOrSearchQuery
import "github.com/antlr4-go/antlr/v4"

// BaseDefaultOrSearchQueryListener is a complete listener for a parse tree produced by DefaultOrSearchQueryParser.
type BaseDefaultOrSearchQueryListener struct{}

var _ DefaultOrSearchQueryListener = &BaseDefaultOrSearchQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDefaultOrSearchQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDefaultOrSearchQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDefaultOrSearchQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDefaultOrSearchQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseDefaultOrSearchQueryListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseDefaultOrSearchQueryListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseDefaultOrSearchQueryListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseDefaultOrSearchQueryListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseDefaultOrSearchQueryListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseDefaultOrSearchQueryListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterPhrase is called when production phrase is entered.
func (s *BaseDefaultOrSearchQueryListener) EnterPhrase(ctx *PhraseContext) {}

// ExitPhrase is called when production phrase is exited.
func (s *BaseDefaultOrSearchQueryListener) ExitPhrase(ctx *PhraseContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseDefaultOrSearchQueryListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseDefaultOrSearchQueryListener) ExitKeyword(ctx *KeywordContext) {}
