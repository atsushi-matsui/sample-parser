package main

import (
	"encoding/json"
	"fmt"
	"os"

	"local/sample-parser/parser"
	"local/sample-parser/parsing" // Note that with modules you may not be able to use a relative immport path

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	//lexer := parsing.NewSearchQueryLexer(input)
	lexer := parsing.NewDefaultOrSearchQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	//p := parsing.NewSearchQueryParser(stream)
	p := parsing.NewDefaultOrSearchQueryParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//listener := parser.NewTreeShapeListener()
	listener := parser.NewDefaultOrListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.AndExpression())

	res := listener.PopQuery()
	resJSON, _ := json.Marshal(res)
	fmt.Printf("res: %s\n", resJSON)

	//if p.GetError() != nil {
	//	fmt.Println(p.GetError())
	//} else {
	//	// p := &parsing.BaseSearchQueryVisitor{}
	//	tree := p.Expr()
	//
	//	visitor := parser.NewTreeShapeVisitor()
	//
	//	//visitor.Visit(tree)
	//	tree.Accept(visitor)
	//	fmt.Println(tree.GetChild(0).(antlr.ParseTree).GetText()) // Convert the result to string using GetText()
	//	fmt.Println(tree.GetChild(0).GetChild(0).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(1))
	//	//fmt.Println(tree.GetChild(2).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(0).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(1).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(2).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(0).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(1).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(2).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(2).GetChild(0).GetChild(0).GetChild(0).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(2).GetChild(0).GetChild(0).GetChild(1).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(2).GetChild(0).GetChild(0).GetChild(2).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(0).GetChild(0).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(0).GetChild(1).(antlr.ParseTree).GetText())
	//	//fmt.Println(tree.GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(2).GetChild(0).GetChild(0).GetChild(1).GetChild(0).GetChild(2).(antlr.ParseTree).GetText())
	//}
}
