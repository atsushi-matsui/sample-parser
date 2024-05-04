// Code generated from DefaultOrSearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // DefaultOrSearchQuery
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by DefaultOrSearchQueryParser.
type DefaultOrSearchQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DefaultOrSearchQueryParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by DefaultOrSearchQueryParser#orExpression.
	VisitOrExpression(ctx *OrExpressionContext) interface{}

	// Visit a parse tree produced by DefaultOrSearchQueryParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by DefaultOrSearchQueryParser#phrase.
	VisitPhrase(ctx *PhraseContext) interface{}

	// Visit a parse tree produced by DefaultOrSearchQueryParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}
}
