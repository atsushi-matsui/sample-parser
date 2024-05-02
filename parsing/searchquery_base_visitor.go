// Code generated from SearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // SearchQuery
import "github.com/antlr4-go/antlr/v4"

type BaseSearchQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSearchQueryVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitKeywords(ctx *KeywordsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSearchQueryVisitor) VisitAlphabets(ctx *AlphabetsContext) interface{} {
	return v.VisitChildren(ctx)
}
