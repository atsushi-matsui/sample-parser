// Code generated from SearchQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

type SearchQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SearchQueryLexerLexerStaticData struct {
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

func searchquerylexerLexerInit() {
	staticData := &SearchQueryLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "", "'OR'", "'AND'", "'NOT'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "ALPHABETS", "OR_OPERATOR", "AND_OPERATOR", "NOT_OPERATOR",
		"WHITE_SPACES",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "ALPHABETS", "OR_OPERATOR", "AND_OPERATOR", "NOT_OPERATOR",
		"WHITE_SPACES", "DIGITS", "HEX", "STRING", "AND_OR", "AND_OR_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 87, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 4, 6, 44, 8,
		6, 11, 6, 12, 6, 45, 1, 6, 1, 6, 1, 7, 4, 7, 51, 8, 7, 11, 7, 12, 7, 52,
		1, 8, 1, 8, 1, 8, 4, 8, 58, 8, 8, 11, 8, 12, 8, 59, 1, 9, 1, 9, 3, 9, 64,
		8, 9, 1, 9, 1, 9, 5, 9, 68, 8, 9, 10, 9, 12, 9, 71, 9, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 3, 10, 78, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 4, 11,
		84, 8, 11, 11, 11, 12, 11, 85, 0, 0, 12, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 0, 17, 0, 19, 0, 21, 0, 23, 0, 1, 0, 5, 3, 0, 9, 10,
		13, 13, 32, 32, 1, 0, 48, 57, 3, 0, 48, 57, 65, 70, 97, 102, 4, 0, 48,
		57, 65, 90, 97, 122, 126, 126, 5, 0, 43, 43, 45, 46, 48, 57, 65, 90, 97,
		122, 90, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 1, 25,
		1, 0, 0, 0, 3, 27, 1, 0, 0, 0, 5, 29, 1, 0, 0, 0, 7, 31, 1, 0, 0, 0, 9,
		34, 1, 0, 0, 0, 11, 38, 1, 0, 0, 0, 13, 43, 1, 0, 0, 0, 15, 50, 1, 0, 0,
		0, 17, 57, 1, 0, 0, 0, 19, 63, 1, 0, 0, 0, 21, 77, 1, 0, 0, 0, 23, 83,
		1, 0, 0, 0, 25, 26, 5, 40, 0, 0, 26, 2, 1, 0, 0, 0, 27, 28, 5, 41, 0, 0,
		28, 4, 1, 0, 0, 0, 29, 30, 3, 23, 11, 0, 30, 6, 1, 0, 0, 0, 31, 32, 5,
		79, 0, 0, 32, 33, 5, 82, 0, 0, 33, 8, 1, 0, 0, 0, 34, 35, 5, 65, 0, 0,
		35, 36, 5, 78, 0, 0, 36, 37, 5, 68, 0, 0, 37, 10, 1, 0, 0, 0, 38, 39, 5,
		78, 0, 0, 39, 40, 5, 79, 0, 0, 40, 41, 5, 84, 0, 0, 41, 12, 1, 0, 0, 0,
		42, 44, 7, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 43, 1,
		0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 6, 6, 0, 0, 48,
		14, 1, 0, 0, 0, 49, 51, 7, 1, 0, 0, 50, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0,
		0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 16, 1, 0, 0, 0, 54, 55,
		5, 37, 0, 0, 55, 56, 7, 2, 0, 0, 56, 58, 7, 2, 0, 0, 57, 54, 1, 0, 0, 0,
		58, 59, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 18, 1,
		0, 0, 0, 61, 64, 7, 3, 0, 0, 62, 64, 3, 17, 8, 0, 63, 61, 1, 0, 0, 0, 63,
		62, 1, 0, 0, 0, 64, 69, 1, 0, 0, 0, 65, 68, 7, 4, 0, 0, 66, 68, 3, 17,
		8, 0, 67, 65, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 20, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0,
		72, 73, 5, 65, 0, 0, 73, 74, 5, 78, 0, 0, 74, 78, 5, 68, 0, 0, 75, 76,
		5, 79, 0, 0, 76, 78, 5, 82, 0, 0, 77, 72, 1, 0, 0, 0, 77, 75, 1, 0, 0,
		0, 78, 22, 1, 0, 0, 0, 79, 80, 3, 21, 10, 0, 80, 81, 7, 4, 0, 0, 81, 84,
		1, 0, 0, 0, 82, 84, 3, 17, 8, 0, 83, 79, 1, 0, 0, 0, 83, 82, 1, 0, 0, 0,
		84, 85, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 24, 1,
		0, 0, 0, 10, 0, 45, 52, 59, 63, 67, 69, 77, 83, 85, 1, 6, 0, 0,
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

// SearchQueryLexerInit initializes any static state used to implement SearchQueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSearchQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SearchQueryLexerInit() {
	staticData := &SearchQueryLexerLexerStaticData
	staticData.once.Do(searchquerylexerLexerInit)
}

// NewSearchQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSearchQueryLexer(input antlr.CharStream) *SearchQueryLexer {
	SearchQueryLexerInit()
	l := new(SearchQueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SearchQueryLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SearchQuery.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SearchQueryLexer tokens.
const (
	SearchQueryLexerT__0         = 1
	SearchQueryLexerT__1         = 2
	SearchQueryLexerALPHABETS    = 3
	SearchQueryLexerOR_OPERATOR  = 4
	SearchQueryLexerAND_OPERATOR = 5
	SearchQueryLexerNOT_OPERATOR = 6
	SearchQueryLexerWHITE_SPACES = 7
)
