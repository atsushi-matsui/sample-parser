// Code generated from DefaultOrSearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // DefaultOrSearchQuery
import "github.com/antlr4-go/antlr/v4"

// DefaultOrSearchQueryListener is a complete listener for a parse tree produced by DefaultOrSearchQueryParser.
type DefaultOrSearchQueryListener interface {
	antlr.ParseTreeListener

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterOrExpression is called when entering the orExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterPhrase is called when entering the phrase production.
	EnterPhrase(c *PhraseContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitOrExpression is called when exiting the orExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitPhrase is called when exiting the phrase production.
	ExitPhrase(c *PhraseContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)
}
