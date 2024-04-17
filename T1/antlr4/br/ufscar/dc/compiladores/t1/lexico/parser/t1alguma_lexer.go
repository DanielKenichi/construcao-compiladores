// Code generated from T1AlgumaLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

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

type T1AlgumaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var T1AlgumaLexerLexerStaticData struct {
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

func t1algumalexerLexerInit() {
	staticData := &T1AlgumaLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.SymbolicNames = []string{
		"", "COMENTARIO", "WS",
	}
	staticData.RuleNames = []string{
		"COMENTARIO", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 2, 20, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 5, 0, 8, 8, 0,
		10, 0, 12, 0, 11, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 9, 0, 2, 1, 1, 3, 2, 1, 0, 2, 1, 0, 10, 10, 3, 0, 9, 10, 13, 13, 32,
		32, 20, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 1, 5, 1, 0, 0, 0, 3, 17, 1,
		0, 0, 0, 5, 9, 5, 123, 0, 0, 6, 8, 8, 0, 0, 0, 7, 6, 1, 0, 0, 0, 8, 11,
		1, 0, 0, 0, 9, 10, 1, 0, 0, 0, 9, 7, 1, 0, 0, 0, 10, 12, 1, 0, 0, 0, 11,
		9, 1, 0, 0, 0, 12, 13, 5, 32, 0, 0, 13, 14, 5, 125, 0, 0, 14, 15, 1, 0,
		0, 0, 15, 16, 6, 0, 0, 0, 16, 2, 1, 0, 0, 0, 17, 18, 7, 1, 0, 0, 18, 19,
		6, 1, 1, 0, 19, 4, 1, 0, 0, 0, 2, 0, 9, 2, 1, 0, 0, 1, 1, 1,
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

// T1AlgumaLexerInit initializes any static state used to implement T1AlgumaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewT1AlgumaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func T1AlgumaLexerInit() {
	staticData := &T1AlgumaLexerLexerStaticData
	staticData.once.Do(t1algumalexerLexerInit)
}

// NewT1AlgumaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewT1AlgumaLexer(input antlr.CharStream) *T1AlgumaLexer {
	T1AlgumaLexerInit()
	l := new(T1AlgumaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &T1AlgumaLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "T1AlgumaLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// T1AlgumaLexer tokens.
const (
	T1AlgumaLexerCOMENTARIO = 1
	T1AlgumaLexerWS         = 2
)

func (l *T1AlgumaLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 0:
		l.COMENTARIO_Action(localctx, actionIndex)

	case 1:
		l.WS_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *T1AlgumaLexer) COMENTARIO_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		l.Skip()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *T1AlgumaLexer) WS_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 1:
		l.Skip()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
