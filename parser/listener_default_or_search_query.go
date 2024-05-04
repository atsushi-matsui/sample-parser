package parser

import (
	"local/sample-parser/parsing"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
)

type DefaultOrListener struct {
	*parsing.BaseDefaultOrSearchQueryListener

	stackQuery []*types.Query
}

func (t *DefaultOrListener) PushQuery(q *types.Query) {
	t.stackQuery = append(t.stackQuery, q)
}

func (t *DefaultOrListener) PopQuery() *types.Query {
	if len(t.stackQuery) < 1 {
		panic("stack is empty unable to pop")
	}

	result := t.stackQuery[len(t.stackQuery)-1]

	t.stackQuery = t.stackQuery[:len(t.stackQuery)-1]

	return result
}

func NewDefaultOrListener() *DefaultOrListener {
	return new(DefaultOrListener)
}


func (l *DefaultOrListener) ExitAndExpression(ctx *parsing.AndExpressionContext) {
	andExp := ctx.AndExpression()
	if andExp != nil {
		lq := l.PopQuery()
		rq := l.PopQuery()
		l.PushQuery(buildAndQuery(lq, rq))
	}
}

func (l *DefaultOrListener) ExitOrExpression(ctx *parsing.OrExpressionContext) {
	orExp := ctx.OrExpression()
	if orExp != nil{
		lq := l.PopQuery()
		rq := l.PopQuery()
		l.PushQuery(buildOrQuery(lq, rq))
	}
}

func (l *DefaultOrListener) ExitNotExpression(ctx *parsing.NotExpressionContext) {
	notOpe := ctx.NOT()
	if notOpe != nil {
		q := l.PopQuery()
		l.PushQuery(buildNotQuery(q))
	}
}

func (l *DefaultOrListener) ExitPhrase(ctx *parsing.PhraseContext) {
	quotes := ctx.AllDOUBLE_QUOTE()
	if len(quotes) == 2 {
		phrase := strings.TrimLeft(ctx.GetText(), "\"")
		phrase = strings.TrimRight(phrase, "\"")
		l.PushQuery(builderMultiMatchQuery(phrase))
	
	}
}

func (l *DefaultOrListener) ExitKeyword(ctx *parsing.KeywordContext) {
	kws := ctx.AllKEYWORD_CHARACTER()
	if kws != nil {
		str := ""
		for _, kw := range kws {
			str += kw.GetText()
		}
		l.PushQuery(builderMultiMatchQuery(str))
	}
}


func builderMultiMatchQuery(str string) *types.Query {
	return &types.Query{
		MultiMatch: &types.MultiMatchQuery{
			Query: str,
			Type:  &textquerytype.Phrase,
		},
	}
}

func buildAndQuery(lq, rq *types.Query) *types.Query {
	return &types.Query{
		Bool: &types.BoolQuery{
			Must: []types.Query{*lq, *rq},
		},
	}
}

func buildOrQuery(lq, rq *types.Query) *types.Query {
	return &types.Query{
		Bool: &types.BoolQuery{
			Should:             []types.Query{*lq, *rq},
			MinimumShouldMatch: 1,
		},
	}
}

func buildNotQuery(q *types.Query) *types.Query {
	return &types.Query{
		Bool: &types.BoolQuery{
			MustNot: []types.Query{*q},
		},
	}
}
