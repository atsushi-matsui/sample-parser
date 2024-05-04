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
		"NL",
	}
	staticData.RuleNames = []string{
		"AND", "OR", "NOT", "WHITE_SPACE", "KEYWORD_CHARACTER", "DOUBLE_QUOTE",
		"NL",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 42, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 28, 8, 3, 11, 3, 12, 3, 29, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 4, 6, 37, 8, 6, 11, 6, 12, 6, 38, 1, 6, 1, 6, 0,
		0, 7, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 1, 0, 3, 3, 0, 9, 9,
		32, 32, 12288, 12288, 5, 0, 9, 10, 13, 13, 32, 32, 34, 34, 12288, 12288,
		2, 0, 10, 10, 13, 13, 43, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 1, 15, 1, 0, 0, 0, 3, 19, 1, 0, 0, 0, 5, 22, 1, 0, 0, 0, 7,
		27, 1, 0, 0, 0, 9, 31, 1, 0, 0, 0, 11, 33, 1, 0, 0, 0, 13, 36, 1, 0, 0,
		0, 15, 16, 5, 65, 0, 0, 16, 17, 5, 78, 0, 0, 17, 18, 5, 68, 0, 0, 18, 2,
		1, 0, 0, 0, 19, 20, 5, 79, 0, 0, 20, 21, 5, 82, 0, 0, 21, 4, 1, 0, 0, 0,
		22, 23, 5, 78, 0, 0, 23, 24, 5, 79, 0, 0, 24, 25, 5, 84, 0, 0, 25, 6, 1,
		0, 0, 0, 26, 28, 7, 0, 0, 0, 27, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29,
		27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 8, 1, 0, 0, 0, 31, 32, 8, 1, 0,
		0, 32, 10, 1, 0, 0, 0, 33, 34, 5, 34, 0, 0, 34, 12, 1, 0, 0, 0, 35, 37,
		7, 2, 0, 0, 36, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0,
		38, 39, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 41, 6, 6, 0, 0, 41, 14, 1,
		0, 0, 0, 3, 0, 29, 38, 1, 6, 0, 0,
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
	DefaultOrSearchQueryLexerNL                = 7
)
