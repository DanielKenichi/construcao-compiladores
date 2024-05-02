package main

import (
	"os"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type customErrorListener struct {
	output     *os.File
	syntar_err bool
}

func NewCustomErrorListener(output *os.File) *customErrorListener {
	return &customErrorListener{
		output: output,
	}
}

func (l *customErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// Não será necessário para o T2, pode deixar vazio
}

func (l *customErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// Não será necessário para o T2, pode deixar vazio
}

func (l *customErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	// Não será necessário para o T2, pode deixar vazio
}

func (l *customErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, charPositionInLine int, msg string, e antlr.RecognitionException) {
	// Reportar apenas o primeiro erro, ignorar os outros
	if !l.syntar_err {
		t := offendingSymbol.(antlr.Token)

		// Modificação da string "EOF" para conformidade com os casos de testes
		tokenText := t.GetText()
		if tokenText == "<EOF>" {
			tokenText = "EOF"
		}

		// Impressão do erro
		l.output.WriteString("Linha " + strconv.Itoa(line) + ": erro sintatico proximo a " + tokenText + "\n")
		l.syntar_err = true
	}
}
