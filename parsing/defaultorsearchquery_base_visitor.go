// Code generated from DefaultOrSearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // DefaultOrSearchQuery
import "github.com/antlr4-go/antlr/v4"

type BaseDefaultOrSearchQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDefaultOrSearchQueryVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefaultOrSearchQueryVisitor) VisitOrExpression(ctx *OrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefaultOrSearchQueryVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefaultOrSearchQueryVisitor) VisitPhrase(ctx *PhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefaultOrSearchQueryVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}
