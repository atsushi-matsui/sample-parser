package parser

import (
	"local/sample-parser/parsing"

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
	phrase := ctx.PHRASE()
	keyword := ctx.AllKEYWORD_CHARACTER()

	if orExp != nil {
		lq := l.PopQuery()
		rq := l.PopQuery()
		l.PushQuery(buildOrQuery(lq, rq))
	}
	if phrase != nil && len(keyword) > 0 {
		phraseStr := phrase.GetText()
		phraseStr = phraseStr[1 : len(phraseStr)-1]
		keyWordStr := ""
		for _, k := range keyword {
			keyWordStr += k.GetText()
		}
		l.PushQuery(buildOrQuery(builderMultiMatchQuery(phraseStr), builderMultiMatchQuery(keyWordStr)))
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
	phrase := ctx.PHRASE()

	if phrase != nil {
		phrase := phrase.GetText()
		phrase = phrase[1 : len(phrase)-1]
		l.PushQuery(builderMultiMatchQuery(phrase))
	}

}

func (l *DefaultOrListener) ExitKeyword(ctx *parsing.KeywordContext) {
	l.PushQuery(builderMultiMatchQuery(ctx.GetText()))
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
