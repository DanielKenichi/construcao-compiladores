// Code generated from T1AlgumaLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"", "", "'algoritmo'", "'declare'", "'var'", "'constante'", "'literal'",
		"'inteiro'", "'real'", "'logico'", "'verdadeiro'", "'falso'", "'e'",
		"'ou'", "'nao'", "'se'", "'fim_se'", "'entao'", "'senao'", "'caso'",
		"'seja'", "'fim_caso'", "'para'", "'fim_para'", "'ate'", "'faca'", "'enquanto'",
		"'fim_enquanto'", "'tipo'", "'registro'", "'fim_registro'", "'procedimento'",
		"'fim_procedimento'", "'funcao'", "'fim_funcao'", "'retorne'", "'leia'",
		"'escreva'", "'fim_algoritmo'", "'..'", "'<'", "'<='", "'>'", "'>='",
		"'='", "'<>'", "':'", "'('", "')'", "'['", "']'", "','", "'\"'", "'/'",
		"'%'", "'+'", "'-'", "'*'", "'<-'", "'^'", "'&'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "COMENTARIO", "ALGORITMO", "DECLARE", "VAR", "CONSTANTE", "LITERAL",
		"INTEIRO", "REAL", "LOGICO", "VERDADEIRO", "FALSO", "E", "OR", "NAO",
		"SE", "FIM_SE", "ENTAO", "SENAO", "CASO", "SEJA", "FIM_CASO", "PARA",
		"FIM_PARA", "ATE", "FACA", "ENQUANTO", "FIM_ENQUANTO", "TIPO", "REGISTRO",
		"FIM_REGISTRO", "PROCEDIMENTO", "FIM_PROCEDIMENTO", "FUNCAO", "FIM_FUNCAO",
		"RETORNE", "LEIA", "ESCREVA", "FIM_ALGORITMO", "INTERVALO", "MENOR",
		"MENORIGUAL", "MAIOR", "MAIORIGUAL", "IGUAL", "DIFERENTE", "DOISPONTOS",
		"ABREPAR", "FECHAPAR", "ABRECHAVE", "FECHACHAVE", "VIRGULA", "ASPAS",
		"DIVISAO", "RESTO", "SOMA", "SUBTRACAO", "MULTIPLICACAO", "ATRIBUICAO",
		"PONTEIRO", "ENDERECO", "PONTO", "NUM_INT", "NUM_REAL", "IDENT", "CADEIA",
		"ERRO_COMENTARIO_ABERTO", "ERRO_CADEIA_ABERTA", "WS", "ERRO",
	}
	staticData.RuleNames = []string{
		"COMENTARIO", "ALGORITMO", "DECLARE", "VAR", "CONSTANTE", "LITERAL",
		"INTEIRO", "REAL", "LOGICO", "VERDADEIRO", "FALSO", "E", "OR", "NAO",
		"SE", "FIM_SE", "ENTAO", "SENAO", "CASO", "SEJA", "FIM_CASO", "PARA",
		"FIM_PARA", "ATE", "FACA", "ENQUANTO", "FIM_ENQUANTO", "TIPO", "REGISTRO",
		"FIM_REGISTRO", "PROCEDIMENTO", "FIM_PROCEDIMENTO", "FUNCAO", "FIM_FUNCAO",
		"RETORNE", "LEIA", "ESCREVA", "FIM_ALGORITMO", "INTERVALO", "MENOR",
		"MENORIGUAL", "MAIOR", "MAIORIGUAL", "IGUAL", "DIFERENTE", "DOISPONTOS",
		"ABREPAR", "FECHAPAR", "ABRECHAVE", "FECHACHAVE", "VIRGULA", "ASPAS",
		"DIVISAO", "RESTO", "SOMA", "SUBTRACAO", "MULTIPLICACAO", "ATRIBUICAO",
		"PONTEIRO", "ENDERECO", "PONTO", "NUM_INT", "NUM_REAL", "IDENT", "CADEIA",
		"ERRO_COMENTARIO_ABERTO", "ERRO_CADEIA_ABERTA", "WS", "ERRO",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 69, 539, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 1, 0, 1, 0, 5, 0, 142, 8, 0, 10, 0, 12, 0, 145, 9, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40,
		1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1,
		45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50,
		1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1,
		55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60,
		1, 60, 1, 61, 4, 61, 484, 8, 61, 11, 61, 12, 61, 485, 1, 62, 4, 62, 489,
		8, 62, 11, 62, 12, 62, 490, 1, 62, 1, 62, 4, 62, 495, 8, 62, 11, 62, 12,
		62, 496, 3, 62, 499, 8, 62, 1, 63, 1, 63, 5, 63, 503, 8, 63, 10, 63, 12,
		63, 506, 9, 63, 1, 64, 1, 64, 5, 64, 510, 8, 64, 10, 64, 12, 64, 513, 9,
		64, 1, 64, 1, 64, 1, 65, 1, 65, 5, 65, 519, 8, 65, 10, 65, 12, 65, 522,
		9, 65, 1, 65, 1, 65, 1, 66, 1, 66, 5, 66, 528, 8, 66, 10, 66, 12, 66, 531,
		9, 66, 1, 66, 1, 66, 1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 4, 143, 511, 520,
		529, 0, 69, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91,
		46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54,
		109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62,
		125, 63, 127, 64, 129, 65, 131, 66, 133, 67, 135, 68, 137, 69, 1, 0, 6,
		1, 0, 10, 10, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 2, 0, 10, 10, 125, 125, 2, 0, 10, 10, 34, 34, 3, 0, 9, 10, 13, 13,
		32, 32, 547, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0,
		0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0,
		0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0,
		0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1,
		0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99,
		1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0,
		0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1,
		0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0,
		121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0,
		0, 0, 0, 129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135,
		1, 0, 0, 0, 0, 137, 1, 0, 0, 0, 1, 139, 1, 0, 0, 0, 3, 149, 1, 0, 0, 0,
		5, 159, 1, 0, 0, 0, 7, 167, 1, 0, 0, 0, 9, 171, 1, 0, 0, 0, 11, 181, 1,
		0, 0, 0, 13, 189, 1, 0, 0, 0, 15, 197, 1, 0, 0, 0, 17, 202, 1, 0, 0, 0,
		19, 209, 1, 0, 0, 0, 21, 220, 1, 0, 0, 0, 23, 226, 1, 0, 0, 0, 25, 228,
		1, 0, 0, 0, 27, 231, 1, 0, 0, 0, 29, 235, 1, 0, 0, 0, 31, 238, 1, 0, 0,
		0, 33, 245, 1, 0, 0, 0, 35, 251, 1, 0, 0, 0, 37, 257, 1, 0, 0, 0, 39, 262,
		1, 0, 0, 0, 41, 267, 1, 0, 0, 0, 43, 276, 1, 0, 0, 0, 45, 281, 1, 0, 0,
		0, 47, 290, 1, 0, 0, 0, 49, 294, 1, 0, 0, 0, 51, 299, 1, 0, 0, 0, 53, 308,
		1, 0, 0, 0, 55, 321, 1, 0, 0, 0, 57, 326, 1, 0, 0, 0, 59, 335, 1, 0, 0,
		0, 61, 348, 1, 0, 0, 0, 63, 361, 1, 0, 0, 0, 65, 378, 1, 0, 0, 0, 67, 385,
		1, 0, 0, 0, 69, 396, 1, 0, 0, 0, 71, 404, 1, 0, 0, 0, 73, 409, 1, 0, 0,
		0, 75, 417, 1, 0, 0, 0, 77, 431, 1, 0, 0, 0, 79, 434, 1, 0, 0, 0, 81, 436,
		1, 0, 0, 0, 83, 439, 1, 0, 0, 0, 85, 441, 1, 0, 0, 0, 87, 444, 1, 0, 0,
		0, 89, 446, 1, 0, 0, 0, 91, 449, 1, 0, 0, 0, 93, 451, 1, 0, 0, 0, 95, 453,
		1, 0, 0, 0, 97, 455, 1, 0, 0, 0, 99, 457, 1, 0, 0, 0, 101, 459, 1, 0, 0,
		0, 103, 461, 1, 0, 0, 0, 105, 463, 1, 0, 0, 0, 107, 465, 1, 0, 0, 0, 109,
		467, 1, 0, 0, 0, 111, 469, 1, 0, 0, 0, 113, 471, 1, 0, 0, 0, 115, 473,
		1, 0, 0, 0, 117, 476, 1, 0, 0, 0, 119, 478, 1, 0, 0, 0, 121, 480, 1, 0,
		0, 0, 123, 483, 1, 0, 0, 0, 125, 488, 1, 0, 0, 0, 127, 500, 1, 0, 0, 0,
		129, 507, 1, 0, 0, 0, 131, 516, 1, 0, 0, 0, 133, 525, 1, 0, 0, 0, 135,
		534, 1, 0, 0, 0, 137, 537, 1, 0, 0, 0, 139, 143, 5, 123, 0, 0, 140, 142,
		8, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 144, 1, 0,
		0, 0, 143, 141, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0,
		146, 147, 5, 125, 0, 0, 147, 148, 6, 0, 0, 0, 148, 2, 1, 0, 0, 0, 149,
		150, 5, 97, 0, 0, 150, 151, 5, 108, 0, 0, 151, 152, 5, 103, 0, 0, 152,
		153, 5, 111, 0, 0, 153, 154, 5, 114, 0, 0, 154, 155, 5, 105, 0, 0, 155,
		156, 5, 116, 0, 0, 156, 157, 5, 109, 0, 0, 157, 158, 5, 111, 0, 0, 158,
		4, 1, 0, 0, 0, 159, 160, 5, 100, 0, 0, 160, 161, 5, 101, 0, 0, 161, 162,
		5, 99, 0, 0, 162, 163, 5, 108, 0, 0, 163, 164, 5, 97, 0, 0, 164, 165, 5,
		114, 0, 0, 165, 166, 5, 101, 0, 0, 166, 6, 1, 0, 0, 0, 167, 168, 5, 118,
		0, 0, 168, 169, 5, 97, 0, 0, 169, 170, 5, 114, 0, 0, 170, 8, 1, 0, 0, 0,
		171, 172, 5, 99, 0, 0, 172, 173, 5, 111, 0, 0, 173, 174, 5, 110, 0, 0,
		174, 175, 5, 115, 0, 0, 175, 176, 5, 116, 0, 0, 176, 177, 5, 97, 0, 0,
		177, 178, 5, 110, 0, 0, 178, 179, 5, 116, 0, 0, 179, 180, 5, 101, 0, 0,
		180, 10, 1, 0, 0, 0, 181, 182, 5, 108, 0, 0, 182, 183, 5, 105, 0, 0, 183,
		184, 5, 116, 0, 0, 184, 185, 5, 101, 0, 0, 185, 186, 5, 114, 0, 0, 186,
		187, 5, 97, 0, 0, 187, 188, 5, 108, 0, 0, 188, 12, 1, 0, 0, 0, 189, 190,
		5, 105, 0, 0, 190, 191, 5, 110, 0, 0, 191, 192, 5, 116, 0, 0, 192, 193,
		5, 101, 0, 0, 193, 194, 5, 105, 0, 0, 194, 195, 5, 114, 0, 0, 195, 196,
		5, 111, 0, 0, 196, 14, 1, 0, 0, 0, 197, 198, 5, 114, 0, 0, 198, 199, 5,
		101, 0, 0, 199, 200, 5, 97, 0, 0, 200, 201, 5, 108, 0, 0, 201, 16, 1, 0,
		0, 0, 202, 203, 5, 108, 0, 0, 203, 204, 5, 111, 0, 0, 204, 205, 5, 103,
		0, 0, 205, 206, 5, 105, 0, 0, 206, 207, 5, 99, 0, 0, 207, 208, 5, 111,
		0, 0, 208, 18, 1, 0, 0, 0, 209, 210, 5, 118, 0, 0, 210, 211, 5, 101, 0,
		0, 211, 212, 5, 114, 0, 0, 212, 213, 5, 100, 0, 0, 213, 214, 5, 97, 0,
		0, 214, 215, 5, 100, 0, 0, 215, 216, 5, 101, 0, 0, 216, 217, 5, 105, 0,
		0, 217, 218, 5, 114, 0, 0, 218, 219, 5, 111, 0, 0, 219, 20, 1, 0, 0, 0,
		220, 221, 5, 102, 0, 0, 221, 222, 5, 97, 0, 0, 222, 223, 5, 108, 0, 0,
		223, 224, 5, 115, 0, 0, 224, 225, 5, 111, 0, 0, 225, 22, 1, 0, 0, 0, 226,
		227, 5, 101, 0, 0, 227, 24, 1, 0, 0, 0, 228, 229, 5, 111, 0, 0, 229, 230,
		5, 117, 0, 0, 230, 26, 1, 0, 0, 0, 231, 232, 5, 110, 0, 0, 232, 233, 5,
		97, 0, 0, 233, 234, 5, 111, 0, 0, 234, 28, 1, 0, 0, 0, 235, 236, 5, 115,
		0, 0, 236, 237, 5, 101, 0, 0, 237, 30, 1, 0, 0, 0, 238, 239, 5, 102, 0,
		0, 239, 240, 5, 105, 0, 0, 240, 241, 5, 109, 0, 0, 241, 242, 5, 95, 0,
		0, 242, 243, 5, 115, 0, 0, 243, 244, 5, 101, 0, 0, 244, 32, 1, 0, 0, 0,
		245, 246, 5, 101, 0, 0, 246, 247, 5, 110, 0, 0, 247, 248, 5, 116, 0, 0,
		248, 249, 5, 97, 0, 0, 249, 250, 5, 111, 0, 0, 250, 34, 1, 0, 0, 0, 251,
		252, 5, 115, 0, 0, 252, 253, 5, 101, 0, 0, 253, 254, 5, 110, 0, 0, 254,
		255, 5, 97, 0, 0, 255, 256, 5, 111, 0, 0, 256, 36, 1, 0, 0, 0, 257, 258,
		5, 99, 0, 0, 258, 259, 5, 97, 0, 0, 259, 260, 5, 115, 0, 0, 260, 261, 5,
		111, 0, 0, 261, 38, 1, 0, 0, 0, 262, 263, 5, 115, 0, 0, 263, 264, 5, 101,
		0, 0, 264, 265, 5, 106, 0, 0, 265, 266, 5, 97, 0, 0, 266, 40, 1, 0, 0,
		0, 267, 268, 5, 102, 0, 0, 268, 269, 5, 105, 0, 0, 269, 270, 5, 109, 0,
		0, 270, 271, 5, 95, 0, 0, 271, 272, 5, 99, 0, 0, 272, 273, 5, 97, 0, 0,
		273, 274, 5, 115, 0, 0, 274, 275, 5, 111, 0, 0, 275, 42, 1, 0, 0, 0, 276,
		277, 5, 112, 0, 0, 277, 278, 5, 97, 0, 0, 278, 279, 5, 114, 0, 0, 279,
		280, 5, 97, 0, 0, 280, 44, 1, 0, 0, 0, 281, 282, 5, 102, 0, 0, 282, 283,
		5, 105, 0, 0, 283, 284, 5, 109, 0, 0, 284, 285, 5, 95, 0, 0, 285, 286,
		5, 112, 0, 0, 286, 287, 5, 97, 0, 0, 287, 288, 5, 114, 0, 0, 288, 289,
		5, 97, 0, 0, 289, 46, 1, 0, 0, 0, 290, 291, 5, 97, 0, 0, 291, 292, 5, 116,
		0, 0, 292, 293, 5, 101, 0, 0, 293, 48, 1, 0, 0, 0, 294, 295, 5, 102, 0,
		0, 295, 296, 5, 97, 0, 0, 296, 297, 5, 99, 0, 0, 297, 298, 5, 97, 0, 0,
		298, 50, 1, 0, 0, 0, 299, 300, 5, 101, 0, 0, 300, 301, 5, 110, 0, 0, 301,
		302, 5, 113, 0, 0, 302, 303, 5, 117, 0, 0, 303, 304, 5, 97, 0, 0, 304,
		305, 5, 110, 0, 0, 305, 306, 5, 116, 0, 0, 306, 307, 5, 111, 0, 0, 307,
		52, 1, 0, 0, 0, 308, 309, 5, 102, 0, 0, 309, 310, 5, 105, 0, 0, 310, 311,
		5, 109, 0, 0, 311, 312, 5, 95, 0, 0, 312, 313, 5, 101, 0, 0, 313, 314,
		5, 110, 0, 0, 314, 315, 5, 113, 0, 0, 315, 316, 5, 117, 0, 0, 316, 317,
		5, 97, 0, 0, 317, 318, 5, 110, 0, 0, 318, 319, 5, 116, 0, 0, 319, 320,
		5, 111, 0, 0, 320, 54, 1, 0, 0, 0, 321, 322, 5, 116, 0, 0, 322, 323, 5,
		105, 0, 0, 323, 324, 5, 112, 0, 0, 324, 325, 5, 111, 0, 0, 325, 56, 1,
		0, 0, 0, 326, 327, 5, 114, 0, 0, 327, 328, 5, 101, 0, 0, 328, 329, 5, 103,
		0, 0, 329, 330, 5, 105, 0, 0, 330, 331, 5, 115, 0, 0, 331, 332, 5, 116,
		0, 0, 332, 333, 5, 114, 0, 0, 333, 334, 5, 111, 0, 0, 334, 58, 1, 0, 0,
		0, 335, 336, 5, 102, 0, 0, 336, 337, 5, 105, 0, 0, 337, 338, 5, 109, 0,
		0, 338, 339, 5, 95, 0, 0, 339, 340, 5, 114, 0, 0, 340, 341, 5, 101, 0,
		0, 341, 342, 5, 103, 0, 0, 342, 343, 5, 105, 0, 0, 343, 344, 5, 115, 0,
		0, 344, 345, 5, 116, 0, 0, 345, 346, 5, 114, 0, 0, 346, 347, 5, 111, 0,
		0, 347, 60, 1, 0, 0, 0, 348, 349, 5, 112, 0, 0, 349, 350, 5, 114, 0, 0,
		350, 351, 5, 111, 0, 0, 351, 352, 5, 99, 0, 0, 352, 353, 5, 101, 0, 0,
		353, 354, 5, 100, 0, 0, 354, 355, 5, 105, 0, 0, 355, 356, 5, 109, 0, 0,
		356, 357, 5, 101, 0, 0, 357, 358, 5, 110, 0, 0, 358, 359, 5, 116, 0, 0,
		359, 360, 5, 111, 0, 0, 360, 62, 1, 0, 0, 0, 361, 362, 5, 102, 0, 0, 362,
		363, 5, 105, 0, 0, 363, 364, 5, 109, 0, 0, 364, 365, 5, 95, 0, 0, 365,
		366, 5, 112, 0, 0, 366, 367, 5, 114, 0, 0, 367, 368, 5, 111, 0, 0, 368,
		369, 5, 99, 0, 0, 369, 370, 5, 101, 0, 0, 370, 371, 5, 100, 0, 0, 371,
		372, 5, 105, 0, 0, 372, 373, 5, 109, 0, 0, 373, 374, 5, 101, 0, 0, 374,
		375, 5, 110, 0, 0, 375, 376, 5, 116, 0, 0, 376, 377, 5, 111, 0, 0, 377,
		64, 1, 0, 0, 0, 378, 379, 5, 102, 0, 0, 379, 380, 5, 117, 0, 0, 380, 381,
		5, 110, 0, 0, 381, 382, 5, 99, 0, 0, 382, 383, 5, 97, 0, 0, 383, 384, 5,
		111, 0, 0, 384, 66, 1, 0, 0, 0, 385, 386, 5, 102, 0, 0, 386, 387, 5, 105,
		0, 0, 387, 388, 5, 109, 0, 0, 388, 389, 5, 95, 0, 0, 389, 390, 5, 102,
		0, 0, 390, 391, 5, 117, 0, 0, 391, 392, 5, 110, 0, 0, 392, 393, 5, 99,
		0, 0, 393, 394, 5, 97, 0, 0, 394, 395, 5, 111, 0, 0, 395, 68, 1, 0, 0,
		0, 396, 397, 5, 114, 0, 0, 397, 398, 5, 101, 0, 0, 398, 399, 5, 116, 0,
		0, 399, 400, 5, 111, 0, 0, 400, 401, 5, 114, 0, 0, 401, 402, 5, 110, 0,
		0, 402, 403, 5, 101, 0, 0, 403, 70, 1, 0, 0, 0, 404, 405, 5, 108, 0, 0,
		405, 406, 5, 101, 0, 0, 406, 407, 5, 105, 0, 0, 407, 408, 5, 97, 0, 0,
		408, 72, 1, 0, 0, 0, 409, 410, 5, 101, 0, 0, 410, 411, 5, 115, 0, 0, 411,
		412, 5, 99, 0, 0, 412, 413, 5, 114, 0, 0, 413, 414, 5, 101, 0, 0, 414,
		415, 5, 118, 0, 0, 415, 416, 5, 97, 0, 0, 416, 74, 1, 0, 0, 0, 417, 418,
		5, 102, 0, 0, 418, 419, 5, 105, 0, 0, 419, 420, 5, 109, 0, 0, 420, 421,
		5, 95, 0, 0, 421, 422, 5, 97, 0, 0, 422, 423, 5, 108, 0, 0, 423, 424, 5,
		103, 0, 0, 424, 425, 5, 111, 0, 0, 425, 426, 5, 114, 0, 0, 426, 427, 5,
		105, 0, 0, 427, 428, 5, 116, 0, 0, 428, 429, 5, 109, 0, 0, 429, 430, 5,
		111, 0, 0, 430, 76, 1, 0, 0, 0, 431, 432, 5, 46, 0, 0, 432, 433, 5, 46,
		0, 0, 433, 78, 1, 0, 0, 0, 434, 435, 5, 60, 0, 0, 435, 80, 1, 0, 0, 0,
		436, 437, 5, 60, 0, 0, 437, 438, 5, 61, 0, 0, 438, 82, 1, 0, 0, 0, 439,
		440, 5, 62, 0, 0, 440, 84, 1, 0, 0, 0, 441, 442, 5, 62, 0, 0, 442, 443,
		5, 61, 0, 0, 443, 86, 1, 0, 0, 0, 444, 445, 5, 61, 0, 0, 445, 88, 1, 0,
		0, 0, 446, 447, 5, 60, 0, 0, 447, 448, 5, 62, 0, 0, 448, 90, 1, 0, 0, 0,
		449, 450, 5, 58, 0, 0, 450, 92, 1, 0, 0, 0, 451, 452, 5, 40, 0, 0, 452,
		94, 1, 0, 0, 0, 453, 454, 5, 41, 0, 0, 454, 96, 1, 0, 0, 0, 455, 456, 5,
		91, 0, 0, 456, 98, 1, 0, 0, 0, 457, 458, 5, 93, 0, 0, 458, 100, 1, 0, 0,
		0, 459, 460, 5, 44, 0, 0, 460, 102, 1, 0, 0, 0, 461, 462, 5, 34, 0, 0,
		462, 104, 1, 0, 0, 0, 463, 464, 5, 47, 0, 0, 464, 106, 1, 0, 0, 0, 465,
		466, 5, 37, 0, 0, 466, 108, 1, 0, 0, 0, 467, 468, 5, 43, 0, 0, 468, 110,
		1, 0, 0, 0, 469, 470, 5, 45, 0, 0, 470, 112, 1, 0, 0, 0, 471, 472, 5, 42,
		0, 0, 472, 114, 1, 0, 0, 0, 473, 474, 5, 60, 0, 0, 474, 475, 5, 45, 0,
		0, 475, 116, 1, 0, 0, 0, 476, 477, 5, 94, 0, 0, 477, 118, 1, 0, 0, 0, 478,
		479, 5, 38, 0, 0, 479, 120, 1, 0, 0, 0, 480, 481, 5, 46, 0, 0, 481, 122,
		1, 0, 0, 0, 482, 484, 2, 48, 57, 0, 483, 482, 1, 0, 0, 0, 484, 485, 1,
		0, 0, 0, 485, 483, 1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486, 124, 1, 0, 0,
		0, 487, 489, 2, 48, 57, 0, 488, 487, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0,
		490, 488, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 498, 1, 0, 0, 0, 492,
		494, 5, 46, 0, 0, 493, 495, 2, 48, 57, 0, 494, 493, 1, 0, 0, 0, 495, 496,
		1, 0, 0, 0, 496, 494, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0, 497, 499, 1, 0,
		0, 0, 498, 492, 1, 0, 0, 0, 498, 499, 1, 0, 0, 0, 499, 126, 1, 0, 0, 0,
		500, 504, 7, 1, 0, 0, 501, 503, 7, 2, 0, 0, 502, 501, 1, 0, 0, 0, 503,
		506, 1, 0, 0, 0, 504, 502, 1, 0, 0, 0, 504, 505, 1, 0, 0, 0, 505, 128,
		1, 0, 0, 0, 506, 504, 1, 0, 0, 0, 507, 511, 5, 34, 0, 0, 508, 510, 8, 0,
		0, 0, 509, 508, 1, 0, 0, 0, 510, 513, 1, 0, 0, 0, 511, 512, 1, 0, 0, 0,
		511, 509, 1, 0, 0, 0, 512, 514, 1, 0, 0, 0, 513, 511, 1, 0, 0, 0, 514,
		515, 5, 34, 0, 0, 515, 130, 1, 0, 0, 0, 516, 520, 5, 123, 0, 0, 517, 519,
		8, 3, 0, 0, 518, 517, 1, 0, 0, 0, 519, 522, 1, 0, 0, 0, 520, 521, 1, 0,
		0, 0, 520, 518, 1, 0, 0, 0, 521, 523, 1, 0, 0, 0, 522, 520, 1, 0, 0, 0,
		523, 524, 5, 10, 0, 0, 524, 132, 1, 0, 0, 0, 525, 529, 5, 34, 0, 0, 526,
		528, 8, 4, 0, 0, 527, 526, 1, 0, 0, 0, 528, 531, 1, 0, 0, 0, 529, 530,
		1, 0, 0, 0, 529, 527, 1, 0, 0, 0, 530, 532, 1, 0, 0, 0, 531, 529, 1, 0,
		0, 0, 532, 533, 5, 10, 0, 0, 533, 134, 1, 0, 0, 0, 534, 535, 7, 5, 0, 0,
		535, 536, 6, 67, 1, 0, 536, 136, 1, 0, 0, 0, 537, 538, 9, 0, 0, 0, 538,
		138, 1, 0, 0, 0, 10, 0, 143, 485, 490, 496, 498, 504, 511, 520, 529, 2,
		1, 0, 0, 1, 67, 1,
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
	T1AlgumaLexerCOMENTARIO             = 1
	T1AlgumaLexerALGORITMO              = 2
	T1AlgumaLexerDECLARE                = 3
	T1AlgumaLexerVAR                    = 4
	T1AlgumaLexerCONSTANTE              = 5
	T1AlgumaLexerLITERAL                = 6
	T1AlgumaLexerINTEIRO                = 7
	T1AlgumaLexerREAL                   = 8
	T1AlgumaLexerLOGICO                 = 9
	T1AlgumaLexerVERDADEIRO             = 10
	T1AlgumaLexerFALSO                  = 11
	T1AlgumaLexerE                      = 12
	T1AlgumaLexerOR                     = 13
	T1AlgumaLexerNAO                    = 14
	T1AlgumaLexerSE                     = 15
	T1AlgumaLexerFIM_SE                 = 16
	T1AlgumaLexerENTAO                  = 17
	T1AlgumaLexerSENAO                  = 18
	T1AlgumaLexerCASO                   = 19
	T1AlgumaLexerSEJA                   = 20
	T1AlgumaLexerFIM_CASO               = 21
	T1AlgumaLexerPARA                   = 22
	T1AlgumaLexerFIM_PARA               = 23
	T1AlgumaLexerATE                    = 24
	T1AlgumaLexerFACA                   = 25
	T1AlgumaLexerENQUANTO               = 26
	T1AlgumaLexerFIM_ENQUANTO           = 27
	T1AlgumaLexerTIPO                   = 28
	T1AlgumaLexerREGISTRO               = 29
	T1AlgumaLexerFIM_REGISTRO           = 30
	T1AlgumaLexerPROCEDIMENTO           = 31
	T1AlgumaLexerFIM_PROCEDIMENTO       = 32
	T1AlgumaLexerFUNCAO                 = 33
	T1AlgumaLexerFIM_FUNCAO             = 34
	T1AlgumaLexerRETORNE                = 35
	T1AlgumaLexerLEIA                   = 36
	T1AlgumaLexerESCREVA                = 37
	T1AlgumaLexerFIM_ALGORITMO          = 38
	T1AlgumaLexerINTERVALO              = 39
	T1AlgumaLexerMENOR                  = 40
	T1AlgumaLexerMENORIGUAL             = 41
	T1AlgumaLexerMAIOR                  = 42
	T1AlgumaLexerMAIORIGUAL             = 43
	T1AlgumaLexerIGUAL                  = 44
	T1AlgumaLexerDIFERENTE              = 45
	T1AlgumaLexerDOISPONTOS             = 46
	T1AlgumaLexerABREPAR                = 47
	T1AlgumaLexerFECHAPAR               = 48
	T1AlgumaLexerABRECHAVE              = 49
	T1AlgumaLexerFECHACHAVE             = 50
	T1AlgumaLexerVIRGULA                = 51
	T1AlgumaLexerASPAS                  = 52
	T1AlgumaLexerDIVISAO                = 53
	T1AlgumaLexerRESTO                  = 54
	T1AlgumaLexerSOMA                   = 55
	T1AlgumaLexerSUBTRACAO              = 56
	T1AlgumaLexerMULTIPLICACAO          = 57
	T1AlgumaLexerATRIBUICAO             = 58
	T1AlgumaLexerPONTEIRO               = 59
	T1AlgumaLexerENDERECO               = 60
	T1AlgumaLexerPONTO                  = 61
	T1AlgumaLexerNUM_INT                = 62
	T1AlgumaLexerNUM_REAL               = 63
	T1AlgumaLexerIDENT                  = 64
	T1AlgumaLexerCADEIA                 = 65
	T1AlgumaLexerERRO_COMENTARIO_ABERTO = 66
	T1AlgumaLexerERRO_CADEIA_ABERTA     = 67
	T1AlgumaLexerWS                     = 68
	T1AlgumaLexerERRO                   = 69
)

func (l *T1AlgumaLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 0:
		l.COMENTARIO_Action(localctx, actionIndex)

	case 67:
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