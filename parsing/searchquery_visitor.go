// Code generated from SearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // SearchQuery
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SearchQueryParser.
type SearchQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SearchQueryParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#keywords.
	VisitKeywords(ctx *KeywordsContext) interface{}

	// Visit a parse tree produced by SearchQueryParser#alphabets.
	VisitAlphabets(ctx *AlphabetsContext) interface{}
}
