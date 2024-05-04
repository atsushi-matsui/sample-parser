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
	lexer := parsing.NewDefaultOrSearchQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parsing.NewDefaultOrSearchQueryParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	listener := parser.NewDefaultOrListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.AndExpression())

	res := listener.PopQuery()
	resJSON, _ := json.Marshal(res)
	fmt.Printf("res: %s\n", resJSON)
}
