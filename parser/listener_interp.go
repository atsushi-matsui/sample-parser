package parser

import (
	"fmt"
	"local/sample-parser/parsing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type TreeShapeListener struct {
	*parsing.BaseSearchQueryListener

	stackQuery []*types.Query
}

func (t *TreeShapeListener) PushQuery(q *types.Query) {
	t.stackQuery = append(t.stackQuery, q)
}

func (t *TreeShapeListener) PopQuery() *types.Query {
	if len(t.stackQuery) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := t.stackQuery[len(t.stackQuery)-1]

	// Remove the last element from the stack.
	t.stackQuery = t.stackQuery[:len(t.stackQuery)-1]

	return result
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) ExitExpr(ctx *parsing.ExprContext) {

	expr := ctx.Expr()
	ope := ctx.OR_OPERATOR()
	term := ctx.Term()
	if expr != nil && ope != nil && term != nil {
		fmt.Printf("Expr: %s\n", expr.GetText())
		fmt.Printf("OR_OPERATOR: %s\n", ope.GetText())
		fmt.Printf("Term: %s\n", term.GetText())

		tq := t.PopQuery()
		eq := t.PopQuery()
		t.PushQuery(&types.Query{
			Bool: &types.BoolQuery{
				Should:             []types.Query{*eq, *tq},
				MinimumShouldMatch: 1,
			},
		})
	} else {
		fmt.Printf("Term: %s\n", term.GetText())
	}
}

func (t *TreeShapeListener) ExitTerm(ctx *parsing.TermContext) {
	term := ctx.Term()
	ope := ctx.AND_OPERATOR()
	factor := ctx.Factor()

	if ope != nil && term != nil && factor != nil {
		fmt.Printf("Term: %s\n", term.GetText())
		fmt.Printf("AND_OPERATOR: %s\n", ope.GetText())
		fmt.Printf("Factor: %s\n", factor.GetText())

		fq := t.PopQuery()
		tq := t.PopQuery()
		t.PushQuery(&types.Query{
			Bool: &types.BoolQuery{
				Must: []types.Query{*tq, *fq},
			},
		})

	} else {
		fmt.Printf("Factor: %s\n", factor.GetText())
	}
}

func (t *TreeShapeListener) ExitFactor(ctx *parsing.FactorContext) {
	key := ctx.Keywords()
	not := ctx.NOT_OPERATOR()

	if not != nil && key != nil {
		fmt.Printf("Keywords: %s\n", key.GetText())
		fmt.Printf("NOT_OPERATOR: %s\n", not.GetText())

		fk := t.PopQuery()
		t.PushQuery(&types.Query{
			Bool: &types.BoolQuery{
				MustNot: []types.Query{*fk},
			},
		})
	} else {
		fmt.Printf("Keywords: %s\n", key.GetText())
	}
}

func (t *TreeShapeListener) ExitKeywords(ctx *parsing.KeywordsContext) {
	//fmt.Printf("Keywords: %s\n", ctx.GetText())
	expr := ctx.Expr()
	alp := ctx.Alphabets()
	if expr != nil {
		fmt.Printf("Expr: %s\n", expr.GetText())
	}
	if alp != nil {
		fmt.Printf("Alphabets: %s\n", alp.GetText())
	}
}

func (t *TreeShapeListener) ExitAlphabets(ctx *parsing.AlphabetsContext) {
	alp := ctx.ALPHABETS()
	if alp != nil {
		fmt.Printf("ALPHABETS: %s\n", alp.GetText())
		t.PushQuery(&types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Query: alp.GetText(),
			},
		})
	}
}
