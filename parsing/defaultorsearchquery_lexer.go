// Code generated from DefaultOrSearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type DefaultOrSearchQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var DefaultOrSearchQueryLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func defaultorsearchquerylexerLexerInit() {
	staticData := &DefaultOrSearchQueryLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'AND'", "'OR'", "'NOT'", "", "", "'\"'",
	}
	staticData.SymbolicNames = []string{
		"", "AND", "OR", "NOT", "WHITE_SPACE", "KEYWORD_CHARACTER", "DOUBLE_QUOTE",
		"PHRASE", "NL",
	}
	staticData.RuleNames = []string{
		"AND", "OR", "NOT", "WHITE_SPACE", "KEYWORD_CHARACTER", "DOUBLE_QUOTE",
		"PHRASE", "NL",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 53, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 30, 8, 3, 11, 3, 12,
		3, 31, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 40, 8, 6, 10, 6, 12, 6,
		43, 9, 6, 1, 6, 1, 6, 1, 7, 4, 7, 48, 8, 7, 11, 7, 12, 7, 49, 1, 7, 1,
		7, 1, 41, 0, 8, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 1, 0,
		3, 3, 0, 9, 9, 32, 32, 12288, 12288, 5, 0, 9, 10, 13, 13, 32, 32, 34, 34,
		12288, 12288, 2, 0, 10, 10, 13, 13, 55, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 21, 1, 0,
		0, 0, 5, 24, 1, 0, 0, 0, 7, 29, 1, 0, 0, 0, 9, 33, 1, 0, 0, 0, 11, 35,
		1, 0, 0, 0, 13, 37, 1, 0, 0, 0, 15, 47, 1, 0, 0, 0, 17, 18, 5, 65, 0, 0,
		18, 19, 5, 78, 0, 0, 19, 20, 5, 68, 0, 0, 20, 2, 1, 0, 0, 0, 21, 22, 5,
		79, 0, 0, 22, 23, 5, 82, 0, 0, 23, 4, 1, 0, 0, 0, 24, 25, 5, 78, 0, 0,
		25, 26, 5, 79, 0, 0, 26, 27, 5, 84, 0, 0, 27, 6, 1, 0, 0, 0, 28, 30, 7,
		0, 0, 0, 29, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31,
		32, 1, 0, 0, 0, 32, 8, 1, 0, 0, 0, 33, 34, 8, 1, 0, 0, 34, 10, 1, 0, 0,
		0, 35, 36, 5, 34, 0, 0, 36, 12, 1, 0, 0, 0, 37, 41, 3, 11, 5, 0, 38, 40,
		9, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0,
		41, 39, 1, 0, 0, 0, 42, 44, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 3,
		11, 5, 0, 45, 14, 1, 0, 0, 0, 46, 48, 7, 2, 0, 0, 47, 46, 1, 0, 0, 0, 48,
		49, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0,
		0, 51, 52, 6, 7, 0, 0, 52, 16, 1, 0, 0, 0, 4, 0, 31, 41, 49, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// DefaultOrSearchQueryLexerInit initializes any static state used to implement DefaultOrSearchQueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewDefaultOrSearchQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func DefaultOrSearchQueryLexerInit() {
	staticData := &DefaultOrSearchQueryLexerLexerStaticData
	staticData.once.Do(defaultorsearchquerylexerLexerInit)
}

// NewDefaultOrSearchQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewDefaultOrSearchQueryLexer(input antlr.CharStream) *DefaultOrSearchQueryLexer {
	DefaultOrSearchQueryLexerInit()
	l := new(DefaultOrSearchQueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &DefaultOrSearchQueryLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "DefaultOrSearchQuery.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DefaultOrSearchQueryLexer tokens.
const (
	DefaultOrSearchQueryLexerAND               = 1
	DefaultOrSearchQueryLexerOR                = 2
	DefaultOrSearchQueryLexerNOT               = 3
	DefaultOrSearchQueryLexerWHITE_SPACE       = 4
	DefaultOrSearchQueryLexerKEYWORD_CHARACTER = 5
	DefaultOrSearchQueryLexerDOUBLE_QUOTE      = 6
	DefaultOrSearchQueryLexerPHRASE            = 7
	DefaultOrSearchQueryLexerNL                = 8
)
