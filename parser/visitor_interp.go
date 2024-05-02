package parser

import (
	"fmt"
	"local/sample-parser/parsing"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeVisitor struct {
	*parsing.BaseSearchQueryVisitor
}

func NewTreeShapeVisitor() *TreeShapeVisitor {
	return new(TreeShapeVisitor)
}

type Children []interface{}

type Operator string

type Expr struct {
	Operator Operator
	Children Children
}

type Term struct {
	Operator Operator
	Children Children
}

type Factor struct {
	Operator Operator
	Children Children
}

type Keywords struct {
	Children Children
}

type Alphabets string

func (v *TreeShapeVisitor) VisitExpr(ctx *parsing.ExprContext) interface{} {
	if ctx.GetChildCount() == 3 {
		opc := ctx.GetChild(1).(antlr.ParseTree).GetText()
		v1 := ctx.GetChild(0) // Cast ctx.GetChild(0) to antlr.ParseTree
		v2 := ctx.GetChild(2)
		return &Expr{
			Operator: Operator(opc),
			Children: Children{v1, v2},
		}
	} else {
		fmt.Printf("ctx.GetChild(0): %s\n", ctx.GetChild(0).(antlr.ParseTree).GetText())
		return ctx.GetChild(0)
	}
}

func (v *TreeShapeVisitor) VisitTerm(ctx *parsing.TermContext) interface{} {
	if ctx.GetChildCount() == 3 {
		opc := ctx.GetChild(1).(antlr.ParseTree).GetText()
		v1 := ctx.GetChild(0)
		v2 := ctx.GetChild(2)
		return &Term{
			Operator: Operator(opc),
			Children: Children{v1, v2},
		}
	} else {
		return ctx.GetChild(0)
	}
}

func (v *TreeShapeVisitor) VisitFactor(ctx *parsing.FactorContext) interface{} {
	if ctx.GetChildCount() == 2 {
		opc := ctx.GetChild(0).(antlr.ParseTree).GetText()
		v1 := ctx.GetChild(1)
		return &Factor{
			Operator: Operator(opc),
			Children: Children{v1},
		}
	} else {
		return ctx.GetChild(0)
	}
}

func (v *TreeShapeVisitor) VisitKeywords(ctx *parsing.KeywordsContext) interface{} {
	if ctx.GetChildCount() == 3 {
		return ctx.GetChild(1)
	} else {
		return ctx.GetChild(0)
	}
}

func (v *TreeShapeVisitor) VisitAlphabets(ctx *parsing.AlphabetsContext) interface{} {
	return Alphabets(ctx.GetText())
}
