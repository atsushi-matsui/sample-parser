// Code generated from SearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // SearchQuery
import "github.com/antlr4-go/antlr/v4"

// SearchQueryListener is a complete listener for a parse tree produced by SearchQueryParser.
type SearchQueryListener interface {
	antlr.ParseTreeListener

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterKeywords is called when entering the keywords production.
	EnterKeywords(c *KeywordsContext)

	// EnterAlphabets is called when entering the alphabets production.
	EnterAlphabets(c *AlphabetsContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitKeywords is called when exiting the keywords production.
	ExitKeywords(c *KeywordsContext)

	// ExitAlphabets is called when exiting the alphabets production.
	ExitAlphabets(c *AlphabetsContext)
}
