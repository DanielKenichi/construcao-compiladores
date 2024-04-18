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
	staticData.LiteralNames = []string{
		"", "", "'algoritmo'", "'declare'", "'literal'", "'inteiro'", "'leia'",
		"'escreva'", "'fim_algoritmo'", "':'", "'('", "')'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "COMENTARIO", "ALGORITMO", "DECLARE", "LITERAL", "INTEIRO", "LEIA",
		"ESCREVA", "FIM_ALGORITMO", "DOISPONTOS", "ABREPAR", "FECHAPAR", "VIRGULA",
		"IDENT", "CADEIA", "WS",
	}
	staticData.RuleNames = []string{
		"COMENTARIO", "ALGORITMO", "DECLARE", "LITERAL", "INTEIRO", "LEIA",
		"ESCREVA", "FIM_ALGORITMO", "DOISPONTOS", "ABREPAR", "FECHAPAR", "VIRGULA",
		"IDENT", "CADEIA", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 131, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 5, 0, 34, 8, 0, 10, 0, 12, 0, 37, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 5, 12, 115, 8, 12, 10, 12, 12, 12, 118, 9, 12, 1, 13, 1, 13, 5, 13,
		122, 8, 13, 10, 13, 12, 13, 125, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 2, 35, 123, 0, 15, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 1, 0, 4, 1, 0,
		10, 10, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 3, 0, 9,
		10, 13, 13, 32, 32, 133, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 53, 1, 0, 0, 0,
		7, 61, 1, 0, 0, 0, 9, 69, 1, 0, 0, 0, 11, 77, 1, 0, 0, 0, 13, 82, 1, 0,
		0, 0, 15, 90, 1, 0, 0, 0, 17, 104, 1, 0, 0, 0, 19, 106, 1, 0, 0, 0, 21,
		108, 1, 0, 0, 0, 23, 110, 1, 0, 0, 0, 25, 112, 1, 0, 0, 0, 27, 119, 1,
		0, 0, 0, 29, 128, 1, 0, 0, 0, 31, 35, 5, 123, 0, 0, 32, 34, 8, 0, 0, 0,
		33, 32, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 35, 33, 1,
		0, 0, 0, 36, 38, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 39, 5, 32, 0, 0, 39,
		40, 5, 125, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 6, 0, 0, 0, 42, 2, 1, 0,
		0, 0, 43, 44, 5, 97, 0, 0, 44, 45, 5, 108, 0, 0, 45, 46, 5, 103, 0, 0,
		46, 47, 5, 111, 0, 0, 47, 48, 5, 114, 0, 0, 48, 49, 5, 105, 0, 0, 49, 50,
		5, 116, 0, 0, 50, 51, 5, 109, 0, 0, 51, 52, 5, 111, 0, 0, 52, 4, 1, 0,
		0, 0, 53, 54, 5, 100, 0, 0, 54, 55, 5, 101, 0, 0, 55, 56, 5, 99, 0, 0,
		56, 57, 5, 108, 0, 0, 57, 58, 5, 97, 0, 0, 58, 59, 5, 114, 0, 0, 59, 60,
		5, 101, 0, 0, 60, 6, 1, 0, 0, 0, 61, 62, 5, 108, 0, 0, 62, 63, 5, 105,
		0, 0, 63, 64, 5, 116, 0, 0, 64, 65, 5, 101, 0, 0, 65, 66, 5, 114, 0, 0,
		66, 67, 5, 97, 0, 0, 67, 68, 5, 108, 0, 0, 68, 8, 1, 0, 0, 0, 69, 70, 5,
		105, 0, 0, 70, 71, 5, 110, 0, 0, 71, 72, 5, 116, 0, 0, 72, 73, 5, 101,
		0, 0, 73, 74, 5, 105, 0, 0, 74, 75, 5, 114, 0, 0, 75, 76, 5, 111, 0, 0,
		76, 10, 1, 0, 0, 0, 77, 78, 5, 108, 0, 0, 78, 79, 5, 101, 0, 0, 79, 80,
		5, 105, 0, 0, 80, 81, 5, 97, 0, 0, 81, 12, 1, 0, 0, 0, 82, 83, 5, 101,
		0, 0, 83, 84, 5, 115, 0, 0, 84, 85, 5, 99, 0, 0, 85, 86, 5, 114, 0, 0,
		86, 87, 5, 101, 0, 0, 87, 88, 5, 118, 0, 0, 88, 89, 5, 97, 0, 0, 89, 14,
		1, 0, 0, 0, 90, 91, 5, 102, 0, 0, 91, 92, 5, 105, 0, 0, 92, 93, 5, 109,
		0, 0, 93, 94, 5, 95, 0, 0, 94, 95, 5, 97, 0, 0, 95, 96, 5, 108, 0, 0, 96,
		97, 5, 103, 0, 0, 97, 98, 5, 111, 0, 0, 98, 99, 5, 114, 0, 0, 99, 100,
		5, 105, 0, 0, 100, 101, 5, 116, 0, 0, 101, 102, 5, 109, 0, 0, 102, 103,
		5, 111, 0, 0, 103, 16, 1, 0, 0, 0, 104, 105, 5, 58, 0, 0, 105, 18, 1, 0,
		0, 0, 106, 107, 5, 40, 0, 0, 107, 20, 1, 0, 0, 0, 108, 109, 5, 41, 0, 0,
		109, 22, 1, 0, 0, 0, 110, 111, 5, 44, 0, 0, 111, 24, 1, 0, 0, 0, 112, 116,
		7, 1, 0, 0, 113, 115, 7, 2, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0,
		0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 26, 1, 0, 0, 0,
		118, 116, 1, 0, 0, 0, 119, 123, 5, 34, 0, 0, 120, 122, 8, 0, 0, 0, 121,
		120, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 123, 121,
		1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 127, 5, 34,
		0, 0, 127, 28, 1, 0, 0, 0, 128, 129, 7, 3, 0, 0, 129, 130, 6, 14, 1, 0,
		130, 30, 1, 0, 0, 0, 4, 0, 35, 116, 123, 2, 1, 0, 0, 1, 14, 1,
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
	T1AlgumaLexerCOMENTARIO    = 1
	T1AlgumaLexerALGORITMO     = 2
	T1AlgumaLexerDECLARE       = 3
	T1AlgumaLexerLITERAL       = 4
	T1AlgumaLexerINTEIRO       = 5
	T1AlgumaLexerLEIA          = 6
	T1AlgumaLexerESCREVA       = 7
	T1AlgumaLexerFIM_ALGORITMO = 8
	T1AlgumaLexerDOISPONTOS    = 9
	T1AlgumaLexerABREPAR       = 10
	T1AlgumaLexerFECHAPAR      = 11
	T1AlgumaLexerVIRGULA       = 12
	T1AlgumaLexerIDENT         = 13
	T1AlgumaLexerCADEIA        = 14
	T1AlgumaLexerWS            = 15
)

func (l *T1AlgumaLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 0:
		l.COMENTARIO_Action(localctx, actionIndex)

	case 14:
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
