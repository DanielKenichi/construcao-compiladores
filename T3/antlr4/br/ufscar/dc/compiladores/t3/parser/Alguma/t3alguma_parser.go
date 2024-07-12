// Code generated from T3Alguma.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // T3Alguma

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type T3AlgumaParser struct {
	*antlr.BaseParser
}

var T3AlgumaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func t3algumaParserInit() {
	staticData := &T3AlgumaParserStaticData
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
		"INTEIRO", "REAL", "LOGICO", "VERDADEIRO", "FALSO", "E", "OU", "NAO",
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
		"programa", "declaracoes", "declaracoes_variaveis", "variavel", "identificador",
		"exp_aritmetica", "termo", "fator", "op1", "op2", "op3", "parcela",
		"parcela_unario", "op_unario", "parcela_nao_unario", "tipo", "tipo_basico",
		"tipo_variavel", "valor_constante", "registro", "parametro", "parametros",
		"declaracoes_funcoes", "corpo", "cmd", "cmdLeia", "cmdEscreva", "cmdSe",
		"cmdCaso", "cmdPara", "cmdEnquanto", "cmdFaca", "cmdAtribuicao", "cmdChamada",
		"cmdRetorne", "selecao", "item_selecao", "constantes", "numero_intervalo",
		"exp_relacional", "op_relacional", "expressao", "termo_logico", "fator_logico",
		"parcela_logica", "op_logico_1", "op_logico_2",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 69, 531, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 102, 8, 1, 10, 1, 12, 1, 105,
		9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 120, 8, 2, 1, 3, 1, 3, 1, 3, 5, 3, 125, 8, 3, 10, 3,
		12, 3, 128, 9, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 136, 8, 4,
		10, 4, 12, 4, 139, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 145, 8, 4, 10, 4,
		12, 4, 148, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 154, 8, 5, 10, 5, 12, 5,
		157, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 163, 8, 6, 10, 6, 12, 6, 166,
		9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 172, 8, 7, 10, 7, 12, 7, 175, 9, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 3, 11, 184, 8, 11, 1, 11,
		1, 11, 3, 11, 188, 8, 11, 1, 12, 3, 12, 191, 8, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 5, 12, 199, 8, 12, 10, 12, 12, 12, 202, 9, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 212, 8, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14, 219, 8, 14, 1, 15, 1, 15, 3,
		15, 223, 8, 15, 1, 16, 1, 16, 1, 17, 3, 17, 228, 8, 17, 1, 17, 1, 17, 3,
		17, 232, 8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 238, 8, 19, 10, 19,
		12, 19, 241, 9, 19, 1, 19, 1, 19, 1, 20, 3, 20, 246, 8, 20, 1, 20, 1, 20,
		1, 20, 5, 20, 251, 8, 20, 10, 20, 12, 20, 254, 9, 20, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 5, 21, 262, 8, 21, 10, 21, 12, 21, 265, 9, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 271, 8, 22, 1, 22, 1, 22, 5, 22, 275,
		8, 22, 10, 22, 12, 22, 278, 9, 22, 1, 22, 5, 22, 281, 8, 22, 10, 22, 12,
		22, 284, 9, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 291, 8, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 5, 22, 297, 8, 22, 10, 22, 12, 22, 300, 9, 22,
		1, 22, 5, 22, 303, 8, 22, 10, 22, 12, 22, 306, 9, 22, 1, 22, 1, 22, 3,
		22, 310, 8, 22, 1, 23, 5, 23, 313, 8, 23, 10, 23, 12, 23, 316, 9, 23, 1,
		23, 5, 23, 319, 8, 23, 10, 23, 12, 23, 322, 9, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 334, 8, 24, 1,
		25, 1, 25, 1, 25, 3, 25, 339, 8, 25, 1, 25, 1, 25, 1, 25, 3, 25, 344, 8,
		25, 1, 25, 5, 25, 347, 8, 25, 10, 25, 12, 25, 350, 9, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 359, 8, 26, 10, 26, 12, 26, 362,
		9, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 370, 8, 27, 10,
		27, 12, 27, 373, 9, 27, 1, 27, 1, 27, 5, 27, 377, 8, 27, 10, 27, 12, 27,
		380, 9, 27, 3, 27, 382, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 5, 28, 392, 8, 28, 10, 28, 12, 28, 395, 9, 28, 3, 28, 397,
		8, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 5, 29, 409, 8, 29, 10, 29, 12, 29, 412, 9, 29, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 5, 30, 420, 8, 30, 10, 30, 12, 30, 423, 9, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 5, 31, 429, 8, 31, 10, 31, 12, 31, 432, 9, 31,
		1, 31, 1, 31, 1, 31, 1, 32, 3, 32, 438, 8, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 449, 8, 33, 10, 33, 12, 33,
		452, 9, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 5, 35, 460, 8, 35,
		10, 35, 12, 35, 463, 9, 35, 1, 36, 1, 36, 1, 36, 5, 36, 468, 8, 36, 10,
		36, 12, 36, 471, 9, 36, 1, 37, 1, 37, 1, 37, 5, 37, 476, 8, 37, 10, 37,
		12, 37, 479, 9, 37, 1, 38, 3, 38, 482, 8, 38, 1, 38, 1, 38, 1, 38, 3, 38,
		487, 8, 38, 1, 38, 3, 38, 490, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39,
		496, 8, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 504, 8, 41,
		10, 41, 12, 41, 507, 9, 41, 1, 42, 1, 42, 1, 42, 1, 42, 5, 42, 513, 8,
		42, 10, 42, 12, 42, 516, 9, 42, 1, 43, 3, 43, 519, 8, 43, 1, 43, 1, 43,
		1, 44, 1, 44, 3, 44, 525, 8, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 0,
		0, 47, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 0, 6, 1, 0, 55, 56, 2, 0, 53,
		53, 57, 57, 1, 0, 6, 9, 3, 0, 10, 11, 62, 63, 65, 65, 1, 0, 40, 45, 1,
		0, 10, 11, 552, 0, 94, 1, 0, 0, 0, 2, 103, 1, 0, 0, 0, 4, 119, 1, 0, 0,
		0, 6, 121, 1, 0, 0, 0, 8, 132, 1, 0, 0, 0, 10, 149, 1, 0, 0, 0, 12, 158,
		1, 0, 0, 0, 14, 167, 1, 0, 0, 0, 16, 176, 1, 0, 0, 0, 18, 178, 1, 0, 0,
		0, 20, 180, 1, 0, 0, 0, 22, 187, 1, 0, 0, 0, 24, 211, 1, 0, 0, 0, 26, 213,
		1, 0, 0, 0, 28, 218, 1, 0, 0, 0, 30, 222, 1, 0, 0, 0, 32, 224, 1, 0, 0,
		0, 34, 227, 1, 0, 0, 0, 36, 233, 1, 0, 0, 0, 38, 235, 1, 0, 0, 0, 40, 245,
		1, 0, 0, 0, 42, 258, 1, 0, 0, 0, 44, 309, 1, 0, 0, 0, 46, 314, 1, 0, 0,
		0, 48, 333, 1, 0, 0, 0, 50, 335, 1, 0, 0, 0, 52, 353, 1, 0, 0, 0, 54, 365,
		1, 0, 0, 0, 56, 385, 1, 0, 0, 0, 58, 400, 1, 0, 0, 0, 60, 415, 1, 0, 0,
		0, 62, 426, 1, 0, 0, 0, 64, 437, 1, 0, 0, 0, 66, 443, 1, 0, 0, 0, 68, 455,
		1, 0, 0, 0, 70, 461, 1, 0, 0, 0, 72, 464, 1, 0, 0, 0, 74, 472, 1, 0, 0,
		0, 76, 481, 1, 0, 0, 0, 78, 491, 1, 0, 0, 0, 80, 497, 1, 0, 0, 0, 82, 499,
		1, 0, 0, 0, 84, 508, 1, 0, 0, 0, 86, 518, 1, 0, 0, 0, 88, 524, 1, 0, 0,
		0, 90, 526, 1, 0, 0, 0, 92, 528, 1, 0, 0, 0, 94, 95, 3, 2, 1, 0, 95, 96,
		5, 2, 0, 0, 96, 97, 3, 46, 23, 0, 97, 98, 5, 38, 0, 0, 98, 1, 1, 0, 0,
		0, 99, 102, 3, 4, 2, 0, 100, 102, 3, 44, 22, 0, 101, 99, 1, 0, 0, 0, 101,
		100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104,
		1, 0, 0, 0, 104, 3, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 107, 5, 3, 0,
		0, 107, 120, 3, 6, 3, 0, 108, 109, 5, 5, 0, 0, 109, 110, 5, 64, 0, 0, 110,
		111, 5, 46, 0, 0, 111, 112, 3, 32, 16, 0, 112, 113, 5, 44, 0, 0, 113, 114,
		3, 36, 18, 0, 114, 120, 1, 0, 0, 0, 115, 116, 5, 28, 0, 0, 116, 117, 5,
		64, 0, 0, 117, 118, 5, 46, 0, 0, 118, 120, 3, 38, 19, 0, 119, 106, 1, 0,
		0, 0, 119, 108, 1, 0, 0, 0, 119, 115, 1, 0, 0, 0, 120, 5, 1, 0, 0, 0, 121,
		126, 3, 8, 4, 0, 122, 123, 5, 51, 0, 0, 123, 125, 3, 8, 4, 0, 124, 122,
		1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0,
		0, 0, 127, 129, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 130, 5, 46, 0, 0,
		130, 131, 3, 30, 15, 0, 131, 7, 1, 0, 0, 0, 132, 137, 5, 64, 0, 0, 133,
		134, 5, 61, 0, 0, 134, 136, 5, 64, 0, 0, 135, 133, 1, 0, 0, 0, 136, 139,
		1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 146, 1, 0,
		0, 0, 139, 137, 1, 0, 0, 0, 140, 141, 5, 49, 0, 0, 141, 142, 3, 10, 5,
		0, 142, 143, 5, 50, 0, 0, 143, 145, 1, 0, 0, 0, 144, 140, 1, 0, 0, 0, 145,
		148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 9, 1,
		0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 155, 3, 12, 6, 0, 150, 151, 3, 16,
		8, 0, 151, 152, 3, 12, 6, 0, 152, 154, 1, 0, 0, 0, 153, 150, 1, 0, 0, 0,
		154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156,
		11, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 164, 3, 14, 7, 0, 159, 160,
		3, 18, 9, 0, 160, 161, 3, 14, 7, 0, 161, 163, 1, 0, 0, 0, 162, 159, 1,
		0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0,
		0, 165, 13, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 173, 3, 22, 11, 0, 168,
		169, 3, 20, 10, 0, 169, 170, 3, 22, 11, 0, 170, 172, 1, 0, 0, 0, 171, 168,
		1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0,
		0, 0, 174, 15, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 177, 7, 0, 0, 0,
		177, 17, 1, 0, 0, 0, 178, 179, 7, 1, 0, 0, 179, 19, 1, 0, 0, 0, 180, 181,
		5, 54, 0, 0, 181, 21, 1, 0, 0, 0, 182, 184, 3, 26, 13, 0, 183, 182, 1,
		0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 188, 3, 24, 12,
		0, 186, 188, 3, 28, 14, 0, 187, 183, 1, 0, 0, 0, 187, 186, 1, 0, 0, 0,
		188, 23, 1, 0, 0, 0, 189, 191, 5, 59, 0, 0, 190, 189, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 212, 3, 8, 4, 0, 193, 194,
		5, 64, 0, 0, 194, 195, 5, 47, 0, 0, 195, 200, 3, 82, 41, 0, 196, 197, 5,
		51, 0, 0, 197, 199, 3, 82, 41, 0, 198, 196, 1, 0, 0, 0, 199, 202, 1, 0,
		0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 203, 1, 0, 0, 0,
		202, 200, 1, 0, 0, 0, 203, 204, 5, 48, 0, 0, 204, 212, 1, 0, 0, 0, 205,
		212, 5, 62, 0, 0, 206, 212, 5, 63, 0, 0, 207, 208, 5, 47, 0, 0, 208, 209,
		3, 82, 41, 0, 209, 210, 5, 48, 0, 0, 210, 212, 1, 0, 0, 0, 211, 190, 1,
		0, 0, 0, 211, 193, 1, 0, 0, 0, 211, 205, 1, 0, 0, 0, 211, 206, 1, 0, 0,
		0, 211, 207, 1, 0, 0, 0, 212, 25, 1, 0, 0, 0, 213, 214, 5, 56, 0, 0, 214,
		27, 1, 0, 0, 0, 215, 216, 5, 60, 0, 0, 216, 219, 3, 8, 4, 0, 217, 219,
		5, 65, 0, 0, 218, 215, 1, 0, 0, 0, 218, 217, 1, 0, 0, 0, 219, 29, 1, 0,
		0, 0, 220, 223, 3, 38, 19, 0, 221, 223, 3, 34, 17, 0, 222, 220, 1, 0, 0,
		0, 222, 221, 1, 0, 0, 0, 223, 31, 1, 0, 0, 0, 224, 225, 7, 2, 0, 0, 225,
		33, 1, 0, 0, 0, 226, 228, 5, 59, 0, 0, 227, 226, 1, 0, 0, 0, 227, 228,
		1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 232, 3, 32, 16, 0, 230, 232, 5,
		64, 0, 0, 231, 229, 1, 0, 0, 0, 231, 230, 1, 0, 0, 0, 232, 35, 1, 0, 0,
		0, 233, 234, 7, 3, 0, 0, 234, 37, 1, 0, 0, 0, 235, 239, 5, 29, 0, 0, 236,
		238, 3, 6, 3, 0, 237, 236, 1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 242, 1, 0, 0, 0, 241, 239, 1, 0,
		0, 0, 242, 243, 5, 30, 0, 0, 243, 39, 1, 0, 0, 0, 244, 246, 5, 4, 0, 0,
		245, 244, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247,
		252, 3, 8, 4, 0, 248, 249, 5, 51, 0, 0, 249, 251, 3, 8, 4, 0, 250, 248,
		1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0,
		0, 0, 253, 255, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 256, 5, 46, 0, 0,
		256, 257, 3, 34, 17, 0, 257, 41, 1, 0, 0, 0, 258, 263, 3, 40, 20, 0, 259,
		260, 5, 51, 0, 0, 260, 262, 3, 40, 20, 0, 261, 259, 1, 0, 0, 0, 262, 265,
		1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 43, 1, 0,
		0, 0, 265, 263, 1, 0, 0, 0, 266, 267, 5, 31, 0, 0, 267, 268, 5, 64, 0,
		0, 268, 270, 5, 47, 0, 0, 269, 271, 3, 42, 21, 0, 270, 269, 1, 0, 0, 0,
		270, 271, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 276, 5, 48, 0, 0, 273,
		275, 3, 4, 2, 0, 274, 273, 1, 0, 0, 0, 275, 278, 1, 0, 0, 0, 276, 274,
		1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 282, 1, 0, 0, 0, 278, 276, 1, 0,
		0, 0, 279, 281, 3, 48, 24, 0, 280, 279, 1, 0, 0, 0, 281, 284, 1, 0, 0,
		0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 285, 1, 0, 0, 0, 284,
		282, 1, 0, 0, 0, 285, 310, 5, 32, 0, 0, 286, 287, 5, 33, 0, 0, 287, 288,
		5, 64, 0, 0, 288, 290, 5, 47, 0, 0, 289, 291, 3, 42, 21, 0, 290, 289, 1,
		0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293, 5, 48, 0,
		0, 293, 294, 5, 46, 0, 0, 294, 298, 3, 34, 17, 0, 295, 297, 3, 4, 2, 0,
		296, 295, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298,
		299, 1, 0, 0, 0, 299, 304, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301, 303,
		3, 48, 24, 0, 302, 301, 1, 0, 0, 0, 303, 306, 1, 0, 0, 0, 304, 302, 1,
		0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 307, 1, 0, 0, 0, 306, 304, 1, 0, 0,
		0, 307, 308, 5, 34, 0, 0, 308, 310, 1, 0, 0, 0, 309, 266, 1, 0, 0, 0, 309,
		286, 1, 0, 0, 0, 310, 45, 1, 0, 0, 0, 311, 313, 3, 4, 2, 0, 312, 311, 1,
		0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0,
		0, 315, 320, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317, 319, 3, 48, 24, 0,
		318, 317, 1, 0, 0, 0, 319, 322, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320,
		321, 1, 0, 0, 0, 321, 47, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 323, 334, 3,
		50, 25, 0, 324, 334, 3, 52, 26, 0, 325, 334, 3, 54, 27, 0, 326, 334, 3,
		56, 28, 0, 327, 334, 3, 58, 29, 0, 328, 334, 3, 60, 30, 0, 329, 334, 3,
		62, 31, 0, 330, 334, 3, 64, 32, 0, 331, 334, 3, 66, 33, 0, 332, 334, 3,
		68, 34, 0, 333, 323, 1, 0, 0, 0, 333, 324, 1, 0, 0, 0, 333, 325, 1, 0,
		0, 0, 333, 326, 1, 0, 0, 0, 333, 327, 1, 0, 0, 0, 333, 328, 1, 0, 0, 0,
		333, 329, 1, 0, 0, 0, 333, 330, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333,
		332, 1, 0, 0, 0, 334, 49, 1, 0, 0, 0, 335, 336, 5, 36, 0, 0, 336, 338,
		5, 47, 0, 0, 337, 339, 5, 59, 0, 0, 338, 337, 1, 0, 0, 0, 338, 339, 1,
		0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 348, 3, 8, 4, 0, 341, 343, 5, 51, 0,
		0, 342, 344, 5, 59, 0, 0, 343, 342, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344,
		345, 1, 0, 0, 0, 345, 347, 3, 8, 4, 0, 346, 341, 1, 0, 0, 0, 347, 350,
		1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 351, 1, 0,
		0, 0, 350, 348, 1, 0, 0, 0, 351, 352, 5, 48, 0, 0, 352, 51, 1, 0, 0, 0,
		353, 354, 5, 37, 0, 0, 354, 355, 5, 47, 0, 0, 355, 360, 3, 82, 41, 0, 356,
		357, 5, 51, 0, 0, 357, 359, 3, 82, 41, 0, 358, 356, 1, 0, 0, 0, 359, 362,
		1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 363, 1, 0,
		0, 0, 362, 360, 1, 0, 0, 0, 363, 364, 5, 48, 0, 0, 364, 53, 1, 0, 0, 0,
		365, 366, 5, 15, 0, 0, 366, 367, 3, 82, 41, 0, 367, 371, 5, 17, 0, 0, 368,
		370, 3, 48, 24, 0, 369, 368, 1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369,
		1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 381, 1, 0, 0, 0, 373, 371, 1, 0,
		0, 0, 374, 378, 5, 18, 0, 0, 375, 377, 3, 48, 24, 0, 376, 375, 1, 0, 0,
		0, 377, 380, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379,
		382, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 381, 374, 1, 0, 0, 0, 381, 382,
		1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 384, 5, 16, 0, 0, 384, 55, 1, 0,
		0, 0, 385, 386, 5, 19, 0, 0, 386, 387, 3, 10, 5, 0, 387, 388, 5, 20, 0,
		0, 388, 396, 3, 70, 35, 0, 389, 393, 5, 18, 0, 0, 390, 392, 3, 48, 24,
		0, 391, 390, 1, 0, 0, 0, 392, 395, 1, 0, 0, 0, 393, 391, 1, 0, 0, 0, 393,
		394, 1, 0, 0, 0, 394, 397, 1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 396, 389,
		1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 399, 5, 21,
		0, 0, 399, 57, 1, 0, 0, 0, 400, 401, 5, 22, 0, 0, 401, 402, 5, 64, 0, 0,
		402, 403, 5, 58, 0, 0, 403, 404, 3, 10, 5, 0, 404, 405, 5, 24, 0, 0, 405,
		406, 3, 10, 5, 0, 406, 410, 5, 25, 0, 0, 407, 409, 3, 48, 24, 0, 408, 407,
		1, 0, 0, 0, 409, 412, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 410, 411, 1, 0,
		0, 0, 411, 413, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 413, 414, 5, 23, 0, 0,
		414, 59, 1, 0, 0, 0, 415, 416, 5, 26, 0, 0, 416, 417, 3, 82, 41, 0, 417,
		421, 5, 25, 0, 0, 418, 420, 3, 48, 24, 0, 419, 418, 1, 0, 0, 0, 420, 423,
		1, 0, 0, 0, 421, 419, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 422, 424, 1, 0,
		0, 0, 423, 421, 1, 0, 0, 0, 424, 425, 5, 27, 0, 0, 425, 61, 1, 0, 0, 0,
		426, 430, 5, 25, 0, 0, 427, 429, 3, 48, 24, 0, 428, 427, 1, 0, 0, 0, 429,
		432, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431, 433,
		1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 433, 434, 5, 24, 0, 0, 434, 435, 3, 82,
		41, 0, 435, 63, 1, 0, 0, 0, 436, 438, 5, 59, 0, 0, 437, 436, 1, 0, 0, 0,
		437, 438, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 440, 3, 8, 4, 0, 440,
		441, 5, 58, 0, 0, 441, 442, 3, 82, 41, 0, 442, 65, 1, 0, 0, 0, 443, 444,
		5, 64, 0, 0, 444, 445, 5, 47, 0, 0, 445, 450, 3, 82, 41, 0, 446, 447, 5,
		51, 0, 0, 447, 449, 3, 82, 41, 0, 448, 446, 1, 0, 0, 0, 449, 452, 1, 0,
		0, 0, 450, 448, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 453, 1, 0, 0, 0,
		452, 450, 1, 0, 0, 0, 453, 454, 5, 48, 0, 0, 454, 67, 1, 0, 0, 0, 455,
		456, 5, 35, 0, 0, 456, 457, 3, 82, 41, 0, 457, 69, 1, 0, 0, 0, 458, 460,
		3, 72, 36, 0, 459, 458, 1, 0, 0, 0, 460, 463, 1, 0, 0, 0, 461, 459, 1,
		0, 0, 0, 461, 462, 1, 0, 0, 0, 462, 71, 1, 0, 0, 0, 463, 461, 1, 0, 0,
		0, 464, 465, 3, 74, 37, 0, 465, 469, 5, 46, 0, 0, 466, 468, 3, 48, 24,
		0, 467, 466, 1, 0, 0, 0, 468, 471, 1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 469,
		470, 1, 0, 0, 0, 470, 73, 1, 0, 0, 0, 471, 469, 1, 0, 0, 0, 472, 477, 3,
		76, 38, 0, 473, 474, 5, 51, 0, 0, 474, 476, 3, 76, 38, 0, 475, 473, 1,
		0, 0, 0, 476, 479, 1, 0, 0, 0, 477, 475, 1, 0, 0, 0, 477, 478, 1, 0, 0,
		0, 478, 75, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 480, 482, 3, 26, 13, 0, 481,
		480, 1, 0, 0, 0, 481, 482, 1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 489,
		5, 62, 0, 0, 484, 486, 5, 39, 0, 0, 485, 487, 3, 26, 13, 0, 486, 485, 1,
		0, 0, 0, 486, 487, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 490, 5, 62, 0,
		0, 489, 484, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490, 77, 1, 0, 0, 0, 491,
		495, 3, 10, 5, 0, 492, 493, 3, 80, 40, 0, 493, 494, 3, 10, 5, 0, 494, 496,
		1, 0, 0, 0, 495, 492, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 79, 1, 0,
		0, 0, 497, 498, 7, 4, 0, 0, 498, 81, 1, 0, 0, 0, 499, 505, 3, 84, 42, 0,
		500, 501, 3, 90, 45, 0, 501, 502, 3, 84, 42, 0, 502, 504, 1, 0, 0, 0, 503,
		500, 1, 0, 0, 0, 504, 507, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 505, 506,
		1, 0, 0, 0, 506, 83, 1, 0, 0, 0, 507, 505, 1, 0, 0, 0, 508, 514, 3, 86,
		43, 0, 509, 510, 3, 92, 46, 0, 510, 511, 3, 86, 43, 0, 511, 513, 1, 0,
		0, 0, 512, 509, 1, 0, 0, 0, 513, 516, 1, 0, 0, 0, 514, 512, 1, 0, 0, 0,
		514, 515, 1, 0, 0, 0, 515, 85, 1, 0, 0, 0, 516, 514, 1, 0, 0, 0, 517, 519,
		5, 14, 0, 0, 518, 517, 1, 0, 0, 0, 518, 519, 1, 0, 0, 0, 519, 520, 1, 0,
		0, 0, 520, 521, 3, 88, 44, 0, 521, 87, 1, 0, 0, 0, 522, 525, 7, 5, 0, 0,
		523, 525, 3, 78, 39, 0, 524, 522, 1, 0, 0, 0, 524, 523, 1, 0, 0, 0, 525,
		89, 1, 0, 0, 0, 526, 527, 5, 13, 0, 0, 527, 91, 1, 0, 0, 0, 528, 529, 5,
		12, 0, 0, 529, 93, 1, 0, 0, 0, 57, 101, 103, 119, 126, 137, 146, 155, 164,
		173, 183, 187, 190, 200, 211, 218, 222, 227, 231, 239, 245, 252, 263, 270,
		276, 282, 290, 298, 304, 309, 314, 320, 333, 338, 343, 348, 360, 371, 378,
		381, 393, 396, 410, 421, 430, 437, 450, 461, 469, 477, 481, 486, 489, 495,
		505, 514, 518, 524,
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

// T3AlgumaParserInit initializes any static state used to implement T3AlgumaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewT3AlgumaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func T3AlgumaParserInit() {
	staticData := &T3AlgumaParserStaticData
	staticData.once.Do(t3algumaParserInit)
}

// NewT3AlgumaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewT3AlgumaParser(input antlr.TokenStream) *T3AlgumaParser {
	T3AlgumaParserInit()
	this := new(T3AlgumaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &T3AlgumaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "T3Alguma.g4"

	return this
}

// T3AlgumaParser tokens.
const (
	T3AlgumaParserEOF                    = antlr.TokenEOF
	T3AlgumaParserCOMENTARIO             = 1
	T3AlgumaParserALGORITMO              = 2
	T3AlgumaParserDECLARE                = 3
	T3AlgumaParserVAR                    = 4
	T3AlgumaParserCONSTANTE              = 5
	T3AlgumaParserLITERAL                = 6
	T3AlgumaParserINTEIRO                = 7
	T3AlgumaParserREAL                   = 8
	T3AlgumaParserLOGICO                 = 9
	T3AlgumaParserVERDADEIRO             = 10
	T3AlgumaParserFALSO                  = 11
	T3AlgumaParserE                      = 12
	T3AlgumaParserOU                     = 13
	T3AlgumaParserNAO                    = 14
	T3AlgumaParserSE                     = 15
	T3AlgumaParserFIM_SE                 = 16
	T3AlgumaParserENTAO                  = 17
	T3AlgumaParserSENAO                  = 18
	T3AlgumaParserCASO                   = 19
	T3AlgumaParserSEJA                   = 20
	T3AlgumaParserFIM_CASO               = 21
	T3AlgumaParserPARA                   = 22
	T3AlgumaParserFIM_PARA               = 23
	T3AlgumaParserATE                    = 24
	T3AlgumaParserFACA                   = 25
	T3AlgumaParserENQUANTO               = 26
	T3AlgumaParserFIM_ENQUANTO           = 27
	T3AlgumaParserTIPO                   = 28
	T3AlgumaParserREGISTRO               = 29
	T3AlgumaParserFIM_REGISTRO           = 30
	T3AlgumaParserPROCEDIMENTO           = 31
	T3AlgumaParserFIM_PROCEDIMENTO       = 32
	T3AlgumaParserFUNCAO                 = 33
	T3AlgumaParserFIM_FUNCAO             = 34
	T3AlgumaParserRETORNE                = 35
	T3AlgumaParserLEIA                   = 36
	T3AlgumaParserESCREVA                = 37
	T3AlgumaParserFIM_ALGORITMO          = 38
	T3AlgumaParserINTERVALO              = 39
	T3AlgumaParserMENOR                  = 40
	T3AlgumaParserMENORIGUAL             = 41
	T3AlgumaParserMAIOR                  = 42
	T3AlgumaParserMAIORIGUAL             = 43
	T3AlgumaParserIGUAL                  = 44
	T3AlgumaParserDIFERENTE              = 45
	T3AlgumaParserDOISPONTOS             = 46
	T3AlgumaParserABREPAR                = 47
	T3AlgumaParserFECHAPAR               = 48
	T3AlgumaParserABRECHAVE              = 49
	T3AlgumaParserFECHACHAVE             = 50
	T3AlgumaParserVIRGULA                = 51
	T3AlgumaParserASPAS                  = 52
	T3AlgumaParserDIVISAO                = 53
	T3AlgumaParserRESTO                  = 54
	T3AlgumaParserSOMA                   = 55
	T3AlgumaParserSUBTRACAO              = 56
	T3AlgumaParserMULTIPLICACAO          = 57
	T3AlgumaParserATRIBUICAO             = 58
	T3AlgumaParserPONTEIRO               = 59
	T3AlgumaParserENDERECO               = 60
	T3AlgumaParserPONTO                  = 61
	T3AlgumaParserNUM_INT                = 62
	T3AlgumaParserNUM_REAL               = 63
	T3AlgumaParserIDENT                  = 64
	T3AlgumaParserCADEIA                 = 65
	T3AlgumaParserERRO_COMENTARIO_ABERTO = 66
	T3AlgumaParserERRO_CADEIA_ABERTA     = 67
	T3AlgumaParserWS                     = 68
	T3AlgumaParserERRO                   = 69
)

// T3AlgumaParser rules.
const (
	T3AlgumaParserRULE_programa              = 0
	T3AlgumaParserRULE_declaracoes           = 1
	T3AlgumaParserRULE_declaracoes_variaveis = 2
	T3AlgumaParserRULE_variavel              = 3
	T3AlgumaParserRULE_identificador         = 4
	T3AlgumaParserRULE_exp_aritmetica        = 5
	T3AlgumaParserRULE_termo                 = 6
	T3AlgumaParserRULE_fator                 = 7
	T3AlgumaParserRULE_op1                   = 8
	T3AlgumaParserRULE_op2                   = 9
	T3AlgumaParserRULE_op3                   = 10
	T3AlgumaParserRULE_parcela               = 11
	T3AlgumaParserRULE_parcela_unario        = 12
	T3AlgumaParserRULE_op_unario             = 13
	T3AlgumaParserRULE_parcela_nao_unario    = 14
	T3AlgumaParserRULE_tipo                  = 15
	T3AlgumaParserRULE_tipo_basico           = 16
	T3AlgumaParserRULE_tipo_variavel         = 17
	T3AlgumaParserRULE_valor_constante       = 18
	T3AlgumaParserRULE_registro              = 19
	T3AlgumaParserRULE_parametro             = 20
	T3AlgumaParserRULE_parametros            = 21
	T3AlgumaParserRULE_declaracoes_funcoes   = 22
	T3AlgumaParserRULE_corpo                 = 23
	T3AlgumaParserRULE_cmd                   = 24
	T3AlgumaParserRULE_cmdLeia               = 25
	T3AlgumaParserRULE_cmdEscreva            = 26
	T3AlgumaParserRULE_cmdSe                 = 27
	T3AlgumaParserRULE_cmdCaso               = 28
	T3AlgumaParserRULE_cmdPara               = 29
	T3AlgumaParserRULE_cmdEnquanto           = 30
	T3AlgumaParserRULE_cmdFaca               = 31
	T3AlgumaParserRULE_cmdAtribuicao         = 32
	T3AlgumaParserRULE_cmdChamada            = 33
	T3AlgumaParserRULE_cmdRetorne            = 34
	T3AlgumaParserRULE_selecao               = 35
	T3AlgumaParserRULE_item_selecao          = 36
	T3AlgumaParserRULE_constantes            = 37
	T3AlgumaParserRULE_numero_intervalo      = 38
	T3AlgumaParserRULE_exp_relacional        = 39
	T3AlgumaParserRULE_op_relacional         = 40
	T3AlgumaParserRULE_expressao             = 41
	T3AlgumaParserRULE_termo_logico          = 42
	T3AlgumaParserRULE_fator_logico          = 43
	T3AlgumaParserRULE_parcela_logica        = 44
	T3AlgumaParserRULE_op_logico_1           = 45
	T3AlgumaParserRULE_op_logico_2           = 46
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaracoes() IDeclaracoesContext
	ALGORITMO() antlr.TerminalNode
	Corpo() ICorpoContext
	FIM_ALGORITMO() antlr.TerminalNode

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) Declaracoes() IDeclaracoesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracoesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracoesContext)
}

func (s *ProgramaContext) ALGORITMO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserALGORITMO, 0)
}

func (s *ProgramaContext) Corpo() ICorpoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICorpoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICorpoContext)
}

func (s *ProgramaContext) FIM_ALGORITMO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFIM_ALGORITMO, 0)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, T3AlgumaParserRULE_programa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Declaracoes()
	}
	{
		p.SetState(95)
		p.Match(T3AlgumaParserALGORITMO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Corpo()
	}
	{
		p.SetState(97)
		p.Match(T3AlgumaParserFIM_ALGORITMO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracoesContext is an interface to support dynamic dispatch.
type IDeclaracoesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaracoes_variaveis() []IDeclaracoes_variaveisContext
	Declaracoes_variaveis(i int) IDeclaracoes_variaveisContext
	AllDeclaracoes_funcoes() []IDeclaracoes_funcoesContext
	Declaracoes_funcoes(i int) IDeclaracoes_funcoesContext

	// IsDeclaracoesContext differentiates from other interfaces.
	IsDeclaracoesContext()
}

type DeclaracoesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracoesContext() *DeclaracoesContext {
	var p = new(DeclaracoesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_declaracoes
	return p
}

func InitEmptyDeclaracoesContext(p *DeclaracoesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_declaracoes
}

func (*DeclaracoesContext) IsDeclaracoesContext() {}

func NewDeclaracoesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracoesContext {
	var p = new(DeclaracoesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_declaracoes

	return p
}

func (s *DeclaracoesContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracoesContext) AllDeclaracoes_variaveis() []IDeclaracoes_variaveisContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracoes_variaveisContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracoes_variaveisContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracoes_variaveisContext); ok {
			tst[i] = t.(IDeclaracoes_variaveisContext)
			i++
		}
	}

	return tst
}

func (s *DeclaracoesContext) Declaracoes_variaveis(i int) IDeclaracoes_variaveisContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracoes_variaveisContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracoes_variaveisContext)
}

func (s *DeclaracoesContext) AllDeclaracoes_funcoes() []IDeclaracoes_funcoesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracoes_funcoesContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracoes_funcoesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracoes_funcoesContext); ok {
			tst[i] = t.(IDeclaracoes_funcoesContext)
			i++
		}
	}

	return tst
}

func (s *DeclaracoesContext) Declaracoes_funcoes(i int) IDeclaracoes_funcoesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracoes_funcoesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracoes_funcoesContext)
}

func (s *DeclaracoesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracoesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracoesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterDeclaracoes(s)
	}
}

func (s *DeclaracoesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitDeclaracoes(s)
	}
}

func (s *DeclaracoesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitDeclaracoes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Declaracoes() (localctx IDeclaracoesContext) {
	localctx = NewDeclaracoesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, T3AlgumaParserRULE_declaracoes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&11005853736) != 0 {
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case T3AlgumaParserDECLARE, T3AlgumaParserCONSTANTE, T3AlgumaParserTIPO:
			{
				p.SetState(99)
				p.Declaracoes_variaveis()
			}

		case T3AlgumaParserPROCEDIMENTO, T3AlgumaParserFUNCAO:
			{
				p.SetState(100)
				p.Declaracoes_funcoes()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracoes_variaveisContext is an interface to support dynamic dispatch.
type IDeclaracoes_variaveisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECLARE() antlr.TerminalNode
	Variavel() IVariavelContext
	CONSTANTE() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	DOISPONTOS() antlr.TerminalNode
	Tipo_basico() ITipo_basicoContext
	IGUAL() antlr.TerminalNode
	Valor_constante() IValor_constanteContext
	TIPO() antlr.TerminalNode
	Registro() IRegistroContext

	// IsDeclaracoes_variaveisContext differentiates from other interfaces.
	IsDeclaracoes_variaveisContext()
}

type Declaracoes_variaveisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracoes_variaveisContext() *Declaracoes_variaveisContext {
	var p = new(Declaracoes_variaveisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_declaracoes_variaveis
	return p
}

func InitEmptyDeclaracoes_variaveisContext(p *Declaracoes_variaveisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_declaracoes_variaveis
}

func (*Declaracoes_variaveisContext) IsDeclaracoes_variaveisContext() {}

func NewDeclaracoes_variaveisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracoes_variaveisContext {
	var p = new(Declaracoes_variaveisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_declaracoes_variaveis

	return p
}

func (s *Declaracoes_variaveisContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracoes_variaveisContext) DECLARE() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserDECLARE, 0)
}

func (s *Declaracoes_variaveisContext) Variavel() IVariavelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariavelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariavelContext)
}

func (s *Declaracoes_variaveisContext) CONSTANTE() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserCONSTANTE, 0)
}

func (s *Declaracoes_variaveisContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserIDENT, 0)
}

func (s *Declaracoes_variaveisContext) DOISPONTOS() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserDOISPONTOS, 0)
}

func (s *Declaracoes_variaveisContext) Tipo_basico() ITipo_basicoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_basicoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_basicoContext)
}

func (s *Declaracoes_variaveisContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserIGUAL, 0)
}

func (s *Declaracoes_variaveisContext) Valor_constante() IValor_constanteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValor_constanteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValor_constanteContext)
}

func (s *Declaracoes_variaveisContext) TIPO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserTIPO, 0)
}

func (s *Declaracoes_variaveisContext) Registro() IRegistroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegistroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegistroContext)
}

func (s *Declaracoes_variaveisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracoes_variaveisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracoes_variaveisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterDeclaracoes_variaveis(s)
	}
}

func (s *Declaracoes_variaveisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitDeclaracoes_variaveis(s)
	}
}

func (s *Declaracoes_variaveisContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitDeclaracoes_variaveis(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Declaracoes_variaveis() (localctx IDeclaracoes_variaveisContext) {
	localctx = NewDeclaracoes_variaveisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, T3AlgumaParserRULE_declaracoes_variaveis)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T3AlgumaParserDECLARE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(T3AlgumaParserDECLARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Variavel()
		}

	case T3AlgumaParserCONSTANTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Match(T3AlgumaParserCONSTANTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(T3AlgumaParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Match(T3AlgumaParserDOISPONTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Tipo_basico()
		}
		{
			p.SetState(112)
			p.Match(T3AlgumaParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Valor_constante()
		}

	case T3AlgumaParserTIPO:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.Match(T3AlgumaParserTIPO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(T3AlgumaParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(T3AlgumaParserDOISPONTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Registro()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariavelContext is an interface to support dynamic dispatch.
type IVariavelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentificador() []IIdentificadorContext
	Identificador(i int) IIdentificadorContext
	DOISPONTOS() antlr.TerminalNode
	Tipo() ITipoContext
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsVariavelContext differentiates from other interfaces.
	IsVariavelContext()
}

type VariavelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariavelContext() *VariavelContext {
	var p = new(VariavelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_variavel
	return p
}

func InitEmptyVariavelContext(p *VariavelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_variavel
}

func (*VariavelContext) IsVariavelContext() {}

func NewVariavelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariavelContext {
	var p = new(VariavelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_variavel

	return p
}

func (s *VariavelContext) GetParser() antlr.Parser { return s.parser }

func (s *VariavelContext) AllIdentificador() []IIdentificadorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentificadorContext); ok {
			len++
		}
	}

	tst := make([]IIdentificadorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentificadorContext); ok {
			tst[i] = t.(IIdentificadorContext)
			i++
		}
	}

	return tst
}

func (s *VariavelContext) Identificador(i int) IIdentificadorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *VariavelContext) DOISPONTOS() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserDOISPONTOS, 0)
}

func (s *VariavelContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *VariavelContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserVIRGULA)
}

func (s *VariavelContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVIRGULA, i)
}

func (s *VariavelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariavelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariavelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterVariavel(s)
	}
}

func (s *VariavelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitVariavel(s)
	}
}

func (s *VariavelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitVariavel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Variavel() (localctx IVariavelContext) {
	localctx = NewVariavelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, T3AlgumaParserRULE_variavel)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Identificador()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserVIRGULA {
		{
			p.SetState(122)
			p.Match(T3AlgumaParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Identificador()
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(129)
		p.Match(T3AlgumaParserDOISPONTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Tipo()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentificadorContext is an interface to support dynamic dispatch.
type IIdentificadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllPONTO() []antlr.TerminalNode
	PONTO(i int) antlr.TerminalNode
	AllABRECHAVE() []antlr.TerminalNode
	ABRECHAVE(i int) antlr.TerminalNode
	AllExp_aritmetica() []IExp_aritmeticaContext
	Exp_aritmetica(i int) IExp_aritmeticaContext
	AllFECHACHAVE() []antlr.TerminalNode
	FECHACHAVE(i int) antlr.TerminalNode

	// IsIdentificadorContext differentiates from other interfaces.
	IsIdentificadorContext()
}

type IdentificadorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentificadorContext() *IdentificadorContext {
	var p = new(IdentificadorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_identificador
	return p
}

func InitEmptyIdentificadorContext(p *IdentificadorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_identificador
}

func (*IdentificadorContext) IsIdentificadorContext() {}

func NewIdentificadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentificadorContext {
	var p = new(IdentificadorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_identificador

	return p
}

func (s *IdentificadorContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentificadorContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserIDENT)
}

func (s *IdentificadorContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserIDENT, i)
}

func (s *IdentificadorContext) AllPONTO() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserPONTO)
}

func (s *IdentificadorContext) PONTO(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserPONTO, i)
}

func (s *IdentificadorContext) AllABRECHAVE() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserABRECHAVE)
}

func (s *IdentificadorContext) ABRECHAVE(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserABRECHAVE, i)
}

func (s *IdentificadorContext) AllExp_aritmetica() []IExp_aritmeticaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			len++
		}
	}

	tst := make([]IExp_aritmeticaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExp_aritmeticaContext); ok {
			tst[i] = t.(IExp_aritmeticaContext)
			i++
		}
	}

	return tst
}

func (s *IdentificadorContext) Exp_aritmetica(i int) IExp_aritmeticaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_aritmeticaContext)
}

func (s *IdentificadorContext) AllFECHACHAVE() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserFECHACHAVE)
}

func (s *IdentificadorContext) FECHACHAVE(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFECHACHAVE, i)
}

func (s *IdentificadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentificadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentificadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterIdentificador(s)
	}
}

func (s *IdentificadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitIdentificador(s)
	}
}

func (s *IdentificadorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitIdentificador(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Identificador() (localctx IIdentificadorContext) {
	localctx = NewIdentificadorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, T3AlgumaParserRULE_identificador)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(T3AlgumaParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserPONTO {
		{
			p.SetState(133)
			p.Match(T3AlgumaParserPONTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Match(T3AlgumaParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserABRECHAVE {
		{
			p.SetState(140)
			p.Match(T3AlgumaParserABRECHAVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Exp_aritmetica()
		}
		{
			p.SetState(142)
			p.Match(T3AlgumaParserFECHACHAVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExp_aritmeticaContext is an interface to support dynamic dispatch.
type IExp_aritmeticaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTermo() []ITermoContext
	Termo(i int) ITermoContext
	AllOp1() []IOp1Context
	Op1(i int) IOp1Context

	// IsExp_aritmeticaContext differentiates from other interfaces.
	IsExp_aritmeticaContext()
}

type Exp_aritmeticaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp_aritmeticaContext() *Exp_aritmeticaContext {
	var p = new(Exp_aritmeticaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_exp_aritmetica
	return p
}

func InitEmptyExp_aritmeticaContext(p *Exp_aritmeticaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_exp_aritmetica
}

func (*Exp_aritmeticaContext) IsExp_aritmeticaContext() {}

func NewExp_aritmeticaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_aritmeticaContext {
	var p = new(Exp_aritmeticaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_exp_aritmetica

	return p
}

func (s *Exp_aritmeticaContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp_aritmeticaContext) AllTermo() []ITermoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermoContext); ok {
			len++
		}
	}

	tst := make([]ITermoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermoContext); ok {
			tst[i] = t.(ITermoContext)
			i++
		}
	}

	return tst
}

func (s *Exp_aritmeticaContext) Termo(i int) ITermoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermoContext)
}

func (s *Exp_aritmeticaContext) AllOp1() []IOp1Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp1Context); ok {
			len++
		}
	}

	tst := make([]IOp1Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp1Context); ok {
			tst[i] = t.(IOp1Context)
			i++
		}
	}

	return tst
}

func (s *Exp_aritmeticaContext) Op1(i int) IOp1Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp1Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp1Context)
}

func (s *Exp_aritmeticaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp_aritmeticaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp_aritmeticaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterExp_aritmetica(s)
	}
}

func (s *Exp_aritmeticaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitExp_aritmetica(s)
	}
}

func (s *Exp_aritmeticaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitExp_aritmetica(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Exp_aritmetica() (localctx IExp_aritmeticaContext) {
	localctx = NewExp_aritmeticaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, T3AlgumaParserRULE_exp_aritmetica)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Termo()
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(150)
				p.Op1()
			}
			{
				p.SetState(151)
				p.Termo()
			}

		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermoContext is an interface to support dynamic dispatch.
type ITermoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFator() []IFatorContext
	Fator(i int) IFatorContext
	AllOp2() []IOp2Context
	Op2(i int) IOp2Context

	// IsTermoContext differentiates from other interfaces.
	IsTermoContext()
}

type TermoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermoContext() *TermoContext {
	var p = new(TermoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_termo
	return p
}

func InitEmptyTermoContext(p *TermoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_termo
}

func (*TermoContext) IsTermoContext() {}

func NewTermoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermoContext {
	var p = new(TermoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_termo

	return p
}

func (s *TermoContext) GetParser() antlr.Parser { return s.parser }

func (s *TermoContext) AllFator() []IFatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFatorContext); ok {
			len++
		}
	}

	tst := make([]IFatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFatorContext); ok {
			tst[i] = t.(IFatorContext)
			i++
		}
	}

	return tst
}

func (s *TermoContext) Fator(i int) IFatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFatorContext)
}

func (s *TermoContext) AllOp2() []IOp2Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp2Context); ok {
			len++
		}
	}

	tst := make([]IOp2Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp2Context); ok {
			tst[i] = t.(IOp2Context)
			i++
		}
	}

	return tst
}

func (s *TermoContext) Op2(i int) IOp2Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp2Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp2Context)
}

func (s *TermoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterTermo(s)
	}
}

func (s *TermoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitTermo(s)
	}
}

func (s *TermoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitTermo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Termo() (localctx ITermoContext) {
	localctx = NewTermoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, T3AlgumaParserRULE_termo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Fator()
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserDIVISAO || _la == T3AlgumaParserMULTIPLICACAO {
		{
			p.SetState(159)
			p.Op2()
		}
		{
			p.SetState(160)
			p.Fator()
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFatorContext is an interface to support dynamic dispatch.
type IFatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParcela() []IParcelaContext
	Parcela(i int) IParcelaContext
	AllOp3() []IOp3Context
	Op3(i int) IOp3Context

	// IsFatorContext differentiates from other interfaces.
	IsFatorContext()
}

type FatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFatorContext() *FatorContext {
	var p = new(FatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_fator
	return p
}

func InitEmptyFatorContext(p *FatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_fator
}

func (*FatorContext) IsFatorContext() {}

func NewFatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FatorContext {
	var p = new(FatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_fator

	return p
}

func (s *FatorContext) GetParser() antlr.Parser { return s.parser }

func (s *FatorContext) AllParcela() []IParcelaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParcelaContext); ok {
			len++
		}
	}

	tst := make([]IParcelaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParcelaContext); ok {
			tst[i] = t.(IParcelaContext)
			i++
		}
	}

	return tst
}

func (s *FatorContext) Parcela(i int) IParcelaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParcelaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParcelaContext)
}

func (s *FatorContext) AllOp3() []IOp3Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp3Context); ok {
			len++
		}
	}

	tst := make([]IOp3Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp3Context); ok {
			tst[i] = t.(IOp3Context)
			i++
		}
	}

	return tst
}

func (s *FatorContext) Op3(i int) IOp3Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp3Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp3Context)
}

func (s *FatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterFator(s)
	}
}

func (s *FatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitFator(s)
	}
}

func (s *FatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitFator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Fator() (localctx IFatorContext) {
	localctx = NewFatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, T3AlgumaParserRULE_fator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Parcela()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserRESTO {
		{
			p.SetState(168)
			p.Op3()
		}
		{
			p.SetState(169)
			p.Parcela()
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp1Context is an interface to support dynamic dispatch.
type IOp1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SOMA() antlr.TerminalNode
	SUBTRACAO() antlr.TerminalNode

	// IsOp1Context differentiates from other interfaces.
	IsOp1Context()
}

type Op1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp1Context() *Op1Context {
	var p = new(Op1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op1
	return p
}

func InitEmptyOp1Context(p *Op1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op1
}

func (*Op1Context) IsOp1Context() {}

func NewOp1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op1Context {
	var p = new(Op1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_op1

	return p
}

func (s *Op1Context) GetParser() antlr.Parser { return s.parser }

func (s *Op1Context) SOMA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserSOMA, 0)
}

func (s *Op1Context) SUBTRACAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserSUBTRACAO, 0)
}

func (s *Op1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterOp1(s)
	}
}

func (s *Op1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitOp1(s)
	}
}

func (s *Op1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitOp1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Op1() (localctx IOp1Context) {
	localctx = NewOp1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, T3AlgumaParserRULE_op1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		_la = p.GetTokenStream().LA(1)

		if !(_la == T3AlgumaParserSOMA || _la == T3AlgumaParserSUBTRACAO) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp2Context is an interface to support dynamic dispatch.
type IOp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULTIPLICACAO() antlr.TerminalNode
	DIVISAO() antlr.TerminalNode

	// IsOp2Context differentiates from other interfaces.
	IsOp2Context()
}

type Op2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp2Context() *Op2Context {
	var p = new(Op2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op2
	return p
}

func InitEmptyOp2Context(p *Op2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op2
}

func (*Op2Context) IsOp2Context() {}

func NewOp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op2Context {
	var p = new(Op2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_op2

	return p
}

func (s *Op2Context) GetParser() antlr.Parser { return s.parser }

func (s *Op2Context) MULTIPLICACAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserMULTIPLICACAO, 0)
}

func (s *Op2Context) DIVISAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserDIVISAO, 0)
}

func (s *Op2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterOp2(s)
	}
}

func (s *Op2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitOp2(s)
	}
}

func (s *Op2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitOp2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Op2() (localctx IOp2Context) {
	localctx = NewOp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, T3AlgumaParserRULE_op2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		_la = p.GetTokenStream().LA(1)

		if !(_la == T3AlgumaParserDIVISAO || _la == T3AlgumaParserMULTIPLICACAO) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp3Context is an interface to support dynamic dispatch.
type IOp3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESTO() antlr.TerminalNode

	// IsOp3Context differentiates from other interfaces.
	IsOp3Context()
}

type Op3Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp3Context() *Op3Context {
	var p = new(Op3Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op3
	return p
}

func InitEmptyOp3Context(p *Op3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op3
}

func (*Op3Context) IsOp3Context() {}

func NewOp3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op3Context {
	var p = new(Op3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_op3

	return p
}

func (s *Op3Context) GetParser() antlr.Parser { return s.parser }

func (s *Op3Context) RESTO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserRESTO, 0)
}

func (s *Op3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterOp3(s)
	}
}

func (s *Op3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitOp3(s)
	}
}

func (s *Op3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitOp3(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Op3() (localctx IOp3Context) {
	localctx = NewOp3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, T3AlgumaParserRULE_op3)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(T3AlgumaParserRESTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParcelaContext is an interface to support dynamic dispatch.
type IParcelaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Parcela_unario() IParcela_unarioContext
	Op_unario() IOp_unarioContext
	Parcela_nao_unario() IParcela_nao_unarioContext

	// IsParcelaContext differentiates from other interfaces.
	IsParcelaContext()
}

type ParcelaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParcelaContext() *ParcelaContext {
	var p = new(ParcelaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parcela
	return p
}

func InitEmptyParcelaContext(p *ParcelaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parcela
}

func (*ParcelaContext) IsParcelaContext() {}

func NewParcelaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParcelaContext {
	var p = new(ParcelaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_parcela

	return p
}

func (s *ParcelaContext) GetParser() antlr.Parser { return s.parser }

func (s *ParcelaContext) Parcela_unario() IParcela_unarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParcela_unarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParcela_unarioContext)
}

func (s *ParcelaContext) Op_unario() IOp_unarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_unarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_unarioContext)
}

func (s *ParcelaContext) Parcela_nao_unario() IParcela_nao_unarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParcela_nao_unarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParcela_nao_unarioContext)
}

func (s *ParcelaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParcelaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParcelaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterParcela(s)
	}
}

func (s *ParcelaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitParcela(s)
	}
}

func (s *ParcelaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitParcela(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Parcela() (localctx IParcelaContext) {
	localctx = NewParcelaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, T3AlgumaParserRULE_parcela)
	var _la int

	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T3AlgumaParserABREPAR, T3AlgumaParserSUBTRACAO, T3AlgumaParserPONTEIRO, T3AlgumaParserNUM_INT, T3AlgumaParserNUM_REAL, T3AlgumaParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T3AlgumaParserSUBTRACAO {
			{
				p.SetState(182)
				p.Op_unario()
			}

		}
		{
			p.SetState(185)
			p.Parcela_unario()
		}

	case T3AlgumaParserENDERECO, T3AlgumaParserCADEIA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Parcela_nao_unario()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParcela_unarioContext is an interface to support dynamic dispatch.
type IParcela_unarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() IIdentificadorContext
	PONTEIRO() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ABREPAR() antlr.TerminalNode
	AllExpressao() []IExpressaoContext
	Expressao(i int) IExpressaoContext
	FECHAPAR() antlr.TerminalNode
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode
	NUM_INT() antlr.TerminalNode
	NUM_REAL() antlr.TerminalNode

	// IsParcela_unarioContext differentiates from other interfaces.
	IsParcela_unarioContext()
}

type Parcela_unarioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParcela_unarioContext() *Parcela_unarioContext {
	var p = new(Parcela_unarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parcela_unario
	return p
}

func InitEmptyParcela_unarioContext(p *Parcela_unarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parcela_unario
}

func (*Parcela_unarioContext) IsParcela_unarioContext() {}

func NewParcela_unarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parcela_unarioContext {
	var p = new(Parcela_unarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_parcela_unario

	return p
}

func (s *Parcela_unarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Parcela_unarioContext) Identificador() IIdentificadorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *Parcela_unarioContext) PONTEIRO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserPONTEIRO, 0)
}

func (s *Parcela_unarioContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserIDENT, 0)
}

func (s *Parcela_unarioContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserABREPAR, 0)
}

func (s *Parcela_unarioContext) AllExpressao() []IExpressaoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressaoContext); ok {
			len++
		}
	}

	tst := make([]IExpressaoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressaoContext); ok {
			tst[i] = t.(IExpressaoContext)
			i++
		}
	}

	return tst
}

func (s *Parcela_unarioContext) Expressao(i int) IExpressaoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *Parcela_unarioContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFECHAPAR, 0)
}

func (s *Parcela_unarioContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserVIRGULA)
}

func (s *Parcela_unarioContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVIRGULA, i)
}

func (s *Parcela_unarioContext) NUM_INT() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserNUM_INT, 0)
}

func (s *Parcela_unarioContext) NUM_REAL() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserNUM_REAL, 0)
}

func (s *Parcela_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parcela_unarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parcela_unarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterParcela_unario(s)
	}
}

func (s *Parcela_unarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitParcela_unario(s)
	}
}

func (s *Parcela_unarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitParcela_unario(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Parcela_unario() (localctx IParcela_unarioContext) {
	localctx = NewParcela_unarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, T3AlgumaParserRULE_parcela_unario)
	var _la int

	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T3AlgumaParserPONTEIRO {
			{
				p.SetState(189)
				p.Match(T3AlgumaParserPONTEIRO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(192)
			p.Identificador()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Match(T3AlgumaParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Match(T3AlgumaParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Expressao()
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == T3AlgumaParserVIRGULA {
			{
				p.SetState(196)
				p.Match(T3AlgumaParserVIRGULA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(197)
				p.Expressao()
			}

			p.SetState(202)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(203)
			p.Match(T3AlgumaParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(205)
			p.Match(T3AlgumaParserNUM_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(206)
			p.Match(T3AlgumaParserNUM_REAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(207)
			p.Match(T3AlgumaParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Expressao()
		}
		{
			p.SetState(209)
			p.Match(T3AlgumaParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_unarioContext is an interface to support dynamic dispatch.
type IOp_unarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUBTRACAO() antlr.TerminalNode

	// IsOp_unarioContext differentiates from other interfaces.
	IsOp_unarioContext()
}

type Op_unarioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_unarioContext() *Op_unarioContext {
	var p = new(Op_unarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op_unario
	return p
}

func InitEmptyOp_unarioContext(p *Op_unarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op_unario
}

func (*Op_unarioContext) IsOp_unarioContext() {}

func NewOp_unarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_unarioContext {
	var p = new(Op_unarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_op_unario

	return p
}

func (s *Op_unarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_unarioContext) SUBTRACAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserSUBTRACAO, 0)
}

func (s *Op_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_unarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_unarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterOp_unario(s)
	}
}

func (s *Op_unarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitOp_unario(s)
	}
}

func (s *Op_unarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitOp_unario(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Op_unario() (localctx IOp_unarioContext) {
	localctx = NewOp_unarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, T3AlgumaParserRULE_op_unario)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(T3AlgumaParserSUBTRACAO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParcela_nao_unarioContext is an interface to support dynamic dispatch.
type IParcela_nao_unarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENDERECO() antlr.TerminalNode
	Identificador() IIdentificadorContext
	CADEIA() antlr.TerminalNode

	// IsParcela_nao_unarioContext differentiates from other interfaces.
	IsParcela_nao_unarioContext()
}

type Parcela_nao_unarioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParcela_nao_unarioContext() *Parcela_nao_unarioContext {
	var p = new(Parcela_nao_unarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parcela_nao_unario
	return p
}

func InitEmptyParcela_nao_unarioContext(p *Parcela_nao_unarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parcela_nao_unario
}

func (*Parcela_nao_unarioContext) IsParcela_nao_unarioContext() {}

func NewParcela_nao_unarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parcela_nao_unarioContext {
	var p = new(Parcela_nao_unarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_parcela_nao_unario

	return p
}

func (s *Parcela_nao_unarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Parcela_nao_unarioContext) ENDERECO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserENDERECO, 0)
}

func (s *Parcela_nao_unarioContext) Identificador() IIdentificadorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *Parcela_nao_unarioContext) CADEIA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserCADEIA, 0)
}

func (s *Parcela_nao_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parcela_nao_unarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parcela_nao_unarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterParcela_nao_unario(s)
	}
}

func (s *Parcela_nao_unarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitParcela_nao_unario(s)
	}
}

func (s *Parcela_nao_unarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitParcela_nao_unario(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Parcela_nao_unario() (localctx IParcela_nao_unarioContext) {
	localctx = NewParcela_nao_unarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, T3AlgumaParserRULE_parcela_nao_unario)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T3AlgumaParserENDERECO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.Match(T3AlgumaParserENDERECO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Identificador()
		}

	case T3AlgumaParserCADEIA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(T3AlgumaParserCADEIA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Registro() IRegistroContext
	Tipo_variavel() ITipo_variavelContext

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) Registro() IRegistroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegistroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegistroContext)
}

func (s *TipoContext) Tipo_variavel() ITipo_variavelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_variavelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_variavelContext)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (s *TipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, T3AlgumaParserRULE_tipo)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T3AlgumaParserREGISTRO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Registro()
		}

	case T3AlgumaParserLITERAL, T3AlgumaParserINTEIRO, T3AlgumaParserREAL, T3AlgumaParserLOGICO, T3AlgumaParserPONTEIRO, T3AlgumaParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.Tipo_variavel()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipo_basicoContext is an interface to support dynamic dispatch.
type ITipo_basicoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LITERAL() antlr.TerminalNode
	INTEIRO() antlr.TerminalNode
	REAL() antlr.TerminalNode
	LOGICO() antlr.TerminalNode

	// IsTipo_basicoContext differentiates from other interfaces.
	IsTipo_basicoContext()
}

type Tipo_basicoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipo_basicoContext() *Tipo_basicoContext {
	var p = new(Tipo_basicoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_tipo_basico
	return p
}

func InitEmptyTipo_basicoContext(p *Tipo_basicoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_tipo_basico
}

func (*Tipo_basicoContext) IsTipo_basicoContext() {}

func NewTipo_basicoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_basicoContext {
	var p = new(Tipo_basicoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_tipo_basico

	return p
}

func (s *Tipo_basicoContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_basicoContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserLITERAL, 0)
}

func (s *Tipo_basicoContext) INTEIRO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserINTEIRO, 0)
}

func (s *Tipo_basicoContext) REAL() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserREAL, 0)
}

func (s *Tipo_basicoContext) LOGICO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserLOGICO, 0)
}

func (s *Tipo_basicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_basicoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_basicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterTipo_basico(s)
	}
}

func (s *Tipo_basicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitTipo_basico(s)
	}
}

func (s *Tipo_basicoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitTipo_basico(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Tipo_basico() (localctx ITipo_basicoContext) {
	localctx = NewTipo_basicoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, T3AlgumaParserRULE_tipo_basico)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&960) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipo_variavelContext is an interface to support dynamic dispatch.
type ITipo_variavelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo_basico() ITipo_basicoContext
	IDENT() antlr.TerminalNode
	PONTEIRO() antlr.TerminalNode

	// IsTipo_variavelContext differentiates from other interfaces.
	IsTipo_variavelContext()
}

type Tipo_variavelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipo_variavelContext() *Tipo_variavelContext {
	var p = new(Tipo_variavelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_tipo_variavel
	return p
}

func InitEmptyTipo_variavelContext(p *Tipo_variavelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_tipo_variavel
}

func (*Tipo_variavelContext) IsTipo_variavelContext() {}

func NewTipo_variavelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_variavelContext {
	var p = new(Tipo_variavelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_tipo_variavel

	return p
}

func (s *Tipo_variavelContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_variavelContext) Tipo_basico() ITipo_basicoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_basicoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_basicoContext)
}

func (s *Tipo_variavelContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserIDENT, 0)
}

func (s *Tipo_variavelContext) PONTEIRO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserPONTEIRO, 0)
}

func (s *Tipo_variavelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_variavelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_variavelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterTipo_variavel(s)
	}
}

func (s *Tipo_variavelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitTipo_variavel(s)
	}
}

func (s *Tipo_variavelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitTipo_variavel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Tipo_variavel() (localctx ITipo_variavelContext) {
	localctx = NewTipo_variavelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, T3AlgumaParserRULE_tipo_variavel)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T3AlgumaParserPONTEIRO {
		{
			p.SetState(226)
			p.Match(T3AlgumaParserPONTEIRO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T3AlgumaParserLITERAL, T3AlgumaParserINTEIRO, T3AlgumaParserREAL, T3AlgumaParserLOGICO:
		{
			p.SetState(229)
			p.Tipo_basico()
		}

	case T3AlgumaParserIDENT:
		{
			p.SetState(230)
			p.Match(T3AlgumaParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValor_constanteContext is an interface to support dynamic dispatch.
type IValor_constanteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CADEIA() antlr.TerminalNode
	NUM_INT() antlr.TerminalNode
	NUM_REAL() antlr.TerminalNode
	VERDADEIRO() antlr.TerminalNode
	FALSO() antlr.TerminalNode

	// IsValor_constanteContext differentiates from other interfaces.
	IsValor_constanteContext()
}

type Valor_constanteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValor_constanteContext() *Valor_constanteContext {
	var p = new(Valor_constanteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_valor_constante
	return p
}

func InitEmptyValor_constanteContext(p *Valor_constanteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_valor_constante
}

func (*Valor_constanteContext) IsValor_constanteContext() {}

func NewValor_constanteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Valor_constanteContext {
	var p = new(Valor_constanteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_valor_constante

	return p
}

func (s *Valor_constanteContext) GetParser() antlr.Parser { return s.parser }

func (s *Valor_constanteContext) CADEIA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserCADEIA, 0)
}

func (s *Valor_constanteContext) NUM_INT() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserNUM_INT, 0)
}

func (s *Valor_constanteContext) NUM_REAL() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserNUM_REAL, 0)
}

func (s *Valor_constanteContext) VERDADEIRO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVERDADEIRO, 0)
}

func (s *Valor_constanteContext) FALSO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFALSO, 0)
}

func (s *Valor_constanteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Valor_constanteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Valor_constanteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterValor_constante(s)
	}
}

func (s *Valor_constanteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitValor_constante(s)
	}
}

func (s *Valor_constanteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitValor_constante(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Valor_constante() (localctx IValor_constanteContext) {
	localctx = NewValor_constanteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, T3AlgumaParserRULE_valor_constante)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-10)) & ^0x3f) == 0 && ((int64(1)<<(_la-10))&49539595901075459) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRegistroContext is an interface to support dynamic dispatch.
type IRegistroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGISTRO() antlr.TerminalNode
	FIM_REGISTRO() antlr.TerminalNode
	AllVariavel() []IVariavelContext
	Variavel(i int) IVariavelContext

	// IsRegistroContext differentiates from other interfaces.
	IsRegistroContext()
}

type RegistroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegistroContext() *RegistroContext {
	var p = new(RegistroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_registro
	return p
}

func InitEmptyRegistroContext(p *RegistroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_registro
}

func (*RegistroContext) IsRegistroContext() {}

func NewRegistroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegistroContext {
	var p = new(RegistroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_registro

	return p
}

func (s *RegistroContext) GetParser() antlr.Parser { return s.parser }

func (s *RegistroContext) REGISTRO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserREGISTRO, 0)
}

func (s *RegistroContext) FIM_REGISTRO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFIM_REGISTRO, 0)
}

func (s *RegistroContext) AllVariavel() []IVariavelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariavelContext); ok {
			len++
		}
	}

	tst := make([]IVariavelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariavelContext); ok {
			tst[i] = t.(IVariavelContext)
			i++
		}
	}

	return tst
}

func (s *RegistroContext) Variavel(i int) IVariavelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariavelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariavelContext)
}

func (s *RegistroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegistroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegistroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterRegistro(s)
	}
}

func (s *RegistroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitRegistro(s)
	}
}

func (s *RegistroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitRegistro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Registro() (localctx IRegistroContext) {
	localctx = NewRegistroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, T3AlgumaParserRULE_registro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(T3AlgumaParserREGISTRO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserIDENT {
		{
			p.SetState(236)
			p.Variavel()
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(242)
		p.Match(T3AlgumaParserFIM_REGISTRO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametroContext is an interface to support dynamic dispatch.
type IParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentificador() []IIdentificadorContext
	Identificador(i int) IIdentificadorContext
	DOISPONTOS() antlr.TerminalNode
	Tipo_variavel() ITipo_variavelContext
	VAR() antlr.TerminalNode
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsParametroContext differentiates from other interfaces.
	IsParametroContext()
}

type ParametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametroContext() *ParametroContext {
	var p = new(ParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parametro
	return p
}

func InitEmptyParametroContext(p *ParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parametro
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_parametro

	return p
}

func (s *ParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroContext) AllIdentificador() []IIdentificadorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentificadorContext); ok {
			len++
		}
	}

	tst := make([]IIdentificadorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentificadorContext); ok {
			tst[i] = t.(IIdentificadorContext)
			i++
		}
	}

	return tst
}

func (s *ParametroContext) Identificador(i int) IIdentificadorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *ParametroContext) DOISPONTOS() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserDOISPONTOS, 0)
}

func (s *ParametroContext) Tipo_variavel() ITipo_variavelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_variavelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_variavelContext)
}

func (s *ParametroContext) VAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVAR, 0)
}

func (s *ParametroContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserVIRGULA)
}

func (s *ParametroContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVIRGULA, i)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterParametro(s)
	}
}

func (s *ParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitParametro(s)
	}
}

func (s *ParametroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitParametro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Parametro() (localctx IParametroContext) {
	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, T3AlgumaParserRULE_parametro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T3AlgumaParserVAR {
		{
			p.SetState(244)
			p.Match(T3AlgumaParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(247)
		p.Identificador()
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserVIRGULA {
		{
			p.SetState(248)
			p.Match(T3AlgumaParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Identificador()
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(255)
		p.Match(T3AlgumaParserDOISPONTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.Tipo_variavel()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParametro() []IParametroContext
	Parametro(i int) IParametroContext
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *ParametrosContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserVIRGULA)
}

func (s *ParametrosContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVIRGULA, i)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterParametros(s)
	}
}

func (s *ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitParametros(s)
	}
}

func (s *ParametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitParametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, T3AlgumaParserRULE_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Parametro()
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserVIRGULA {
		{
			p.SetState(259)
			p.Match(T3AlgumaParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Parametro()
		}

		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracoes_funcoesContext is an interface to support dynamic dispatch.
type IDeclaracoes_funcoesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROCEDIMENTO() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ABREPAR() antlr.TerminalNode
	FECHAPAR() antlr.TerminalNode
	FIM_PROCEDIMENTO() antlr.TerminalNode
	Parametros() IParametrosContext
	AllDeclaracoes_variaveis() []IDeclaracoes_variaveisContext
	Declaracoes_variaveis(i int) IDeclaracoes_variaveisContext
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext
	FUNCAO() antlr.TerminalNode
	DOISPONTOS() antlr.TerminalNode
	Tipo_variavel() ITipo_variavelContext
	FIM_FUNCAO() antlr.TerminalNode

	// IsDeclaracoes_funcoesContext differentiates from other interfaces.
	IsDeclaracoes_funcoesContext()
}

type Declaracoes_funcoesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracoes_funcoesContext() *Declaracoes_funcoesContext {
	var p = new(Declaracoes_funcoesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_declaracoes_funcoes
	return p
}

func InitEmptyDeclaracoes_funcoesContext(p *Declaracoes_funcoesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_declaracoes_funcoes
}

func (*Declaracoes_funcoesContext) IsDeclaracoes_funcoesContext() {}

func NewDeclaracoes_funcoesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracoes_funcoesContext {
	var p = new(Declaracoes_funcoesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_declaracoes_funcoes

	return p
}

func (s *Declaracoes_funcoesContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracoes_funcoesContext) PROCEDIMENTO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserPROCEDIMENTO, 0)
}

func (s *Declaracoes_funcoesContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserIDENT, 0)
}

func (s *Declaracoes_funcoesContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserABREPAR, 0)
}

func (s *Declaracoes_funcoesContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFECHAPAR, 0)
}

func (s *Declaracoes_funcoesContext) FIM_PROCEDIMENTO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFIM_PROCEDIMENTO, 0)
}

func (s *Declaracoes_funcoesContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *Declaracoes_funcoesContext) AllDeclaracoes_variaveis() []IDeclaracoes_variaveisContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracoes_variaveisContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracoes_variaveisContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracoes_variaveisContext); ok {
			tst[i] = t.(IDeclaracoes_variaveisContext)
			i++
		}
	}

	return tst
}

func (s *Declaracoes_funcoesContext) Declaracoes_variaveis(i int) IDeclaracoes_variaveisContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracoes_variaveisContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracoes_variaveisContext)
}

func (s *Declaracoes_funcoesContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *Declaracoes_funcoesContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *Declaracoes_funcoesContext) FUNCAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFUNCAO, 0)
}

func (s *Declaracoes_funcoesContext) DOISPONTOS() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserDOISPONTOS, 0)
}

func (s *Declaracoes_funcoesContext) Tipo_variavel() ITipo_variavelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_variavelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_variavelContext)
}

func (s *Declaracoes_funcoesContext) FIM_FUNCAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFIM_FUNCAO, 0)
}

func (s *Declaracoes_funcoesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracoes_funcoesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracoes_funcoesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterDeclaracoes_funcoes(s)
	}
}

func (s *Declaracoes_funcoesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitDeclaracoes_funcoes(s)
	}
}

func (s *Declaracoes_funcoesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitDeclaracoes_funcoes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Declaracoes_funcoes() (localctx IDeclaracoes_funcoesContext) {
	localctx = NewDeclaracoes_funcoesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, T3AlgumaParserRULE_declaracoes_funcoes)
	var _la int

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T3AlgumaParserPROCEDIMENTO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.Match(T3AlgumaParserPROCEDIMENTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(T3AlgumaParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Match(T3AlgumaParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T3AlgumaParserVAR || _la == T3AlgumaParserIDENT {
			{
				p.SetState(269)
				p.Parametros()
			}

		}
		{
			p.SetState(272)
			p.Match(T3AlgumaParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268435496) != 0 {
			{
				p.SetState(273)
				p.Declaracoes_variaveis()
			}

			p.SetState(278)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
			{
				p.SetState(279)
				p.Cmd()
			}

			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(285)
			p.Match(T3AlgumaParserFIM_PROCEDIMENTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T3AlgumaParserFUNCAO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Match(T3AlgumaParserFUNCAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.Match(T3AlgumaParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
			p.Match(T3AlgumaParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T3AlgumaParserVAR || _la == T3AlgumaParserIDENT {
			{
				p.SetState(289)
				p.Parametros()
			}

		}
		{
			p.SetState(292)
			p.Match(T3AlgumaParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.Match(T3AlgumaParserDOISPONTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Tipo_variavel()
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268435496) != 0 {
			{
				p.SetState(295)
				p.Declaracoes_variaveis()
			}

			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
			{
				p.SetState(301)
				p.Cmd()
			}

			p.SetState(306)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(307)
			p.Match(T3AlgumaParserFIM_FUNCAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICorpoContext is an interface to support dynamic dispatch.
type ICorpoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaracoes_variaveis() []IDeclaracoes_variaveisContext
	Declaracoes_variaveis(i int) IDeclaracoes_variaveisContext
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCorpoContext differentiates from other interfaces.
	IsCorpoContext()
}

type CorpoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCorpoContext() *CorpoContext {
	var p = new(CorpoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_corpo
	return p
}

func InitEmptyCorpoContext(p *CorpoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_corpo
}

func (*CorpoContext) IsCorpoContext() {}

func NewCorpoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CorpoContext {
	var p = new(CorpoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_corpo

	return p
}

func (s *CorpoContext) GetParser() antlr.Parser { return s.parser }

func (s *CorpoContext) AllDeclaracoes_variaveis() []IDeclaracoes_variaveisContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracoes_variaveisContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracoes_variaveisContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracoes_variaveisContext); ok {
			tst[i] = t.(IDeclaracoes_variaveisContext)
			i++
		}
	}

	return tst
}

func (s *CorpoContext) Declaracoes_variaveis(i int) IDeclaracoes_variaveisContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracoes_variaveisContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracoes_variaveisContext)
}

func (s *CorpoContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CorpoContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CorpoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CorpoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CorpoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCorpo(s)
	}
}

func (s *CorpoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCorpo(s)
	}
}

func (s *CorpoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCorpo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Corpo() (localctx ICorpoContext) {
	localctx = NewCorpoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, T3AlgumaParserRULE_corpo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268435496) != 0 {
		{
			p.SetState(311)
			p.Declaracoes_variaveis()
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(317)
			p.Cmd()
		}

		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdContext is an interface to support dynamic dispatch.
type ICmdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CmdLeia() ICmdLeiaContext
	CmdEscreva() ICmdEscrevaContext
	CmdSe() ICmdSeContext
	CmdCaso() ICmdCasoContext
	CmdPara() ICmdParaContext
	CmdEnquanto() ICmdEnquantoContext
	CmdFaca() ICmdFacaContext
	CmdAtribuicao() ICmdAtribuicaoContext
	CmdChamada() ICmdChamadaContext
	CmdRetorne() ICmdRetorneContext

	// IsCmdContext differentiates from other interfaces.
	IsCmdContext()
}

type CmdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdContext() *CmdContext {
	var p = new(CmdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmd
	return p
}

func InitEmptyCmdContext(p *CmdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmd
}

func (*CmdContext) IsCmdContext() {}

func NewCmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdContext {
	var p = new(CmdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmd

	return p
}

func (s *CmdContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdContext) CmdLeia() ICmdLeiaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdLeiaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdLeiaContext)
}

func (s *CmdContext) CmdEscreva() ICmdEscrevaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdEscrevaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdEscrevaContext)
}

func (s *CmdContext) CmdSe() ICmdSeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdSeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdSeContext)
}

func (s *CmdContext) CmdCaso() ICmdCasoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdCasoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdCasoContext)
}

func (s *CmdContext) CmdPara() ICmdParaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdParaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdParaContext)
}

func (s *CmdContext) CmdEnquanto() ICmdEnquantoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdEnquantoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdEnquantoContext)
}

func (s *CmdContext) CmdFaca() ICmdFacaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdFacaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdFacaContext)
}

func (s *CmdContext) CmdAtribuicao() ICmdAtribuicaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdAtribuicaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdAtribuicaoContext)
}

func (s *CmdContext) CmdChamada() ICmdChamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdChamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdChamadaContext)
}

func (s *CmdContext) CmdRetorne() ICmdRetorneContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdRetorneContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdRetorneContext)
}

func (s *CmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmd(s)
	}
}

func (s *CmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmd(s)
	}
}

func (s *CmdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Cmd() (localctx ICmdContext) {
	localctx = NewCmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, T3AlgumaParserRULE_cmd)
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)
			p.CmdLeia()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.CmdEscreva()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(325)
			p.CmdSe()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(326)
			p.CmdCaso()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(327)
			p.CmdPara()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(328)
			p.CmdEnquanto()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(329)
			p.CmdFaca()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(330)
			p.CmdAtribuicao()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(331)
			p.CmdChamada()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(332)
			p.CmdRetorne()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdLeiaContext is an interface to support dynamic dispatch.
type ICmdLeiaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEIA() antlr.TerminalNode
	ABREPAR() antlr.TerminalNode
	AllIdentificador() []IIdentificadorContext
	Identificador(i int) IIdentificadorContext
	FECHAPAR() antlr.TerminalNode
	AllPONTEIRO() []antlr.TerminalNode
	PONTEIRO(i int) antlr.TerminalNode
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsCmdLeiaContext differentiates from other interfaces.
	IsCmdLeiaContext()
}

type CmdLeiaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdLeiaContext() *CmdLeiaContext {
	var p = new(CmdLeiaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdLeia
	return p
}

func InitEmptyCmdLeiaContext(p *CmdLeiaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdLeia
}

func (*CmdLeiaContext) IsCmdLeiaContext() {}

func NewCmdLeiaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdLeiaContext {
	var p = new(CmdLeiaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmdLeia

	return p
}

func (s *CmdLeiaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdLeiaContext) LEIA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserLEIA, 0)
}

func (s *CmdLeiaContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserABREPAR, 0)
}

func (s *CmdLeiaContext) AllIdentificador() []IIdentificadorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentificadorContext); ok {
			len++
		}
	}

	tst := make([]IIdentificadorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentificadorContext); ok {
			tst[i] = t.(IIdentificadorContext)
			i++
		}
	}

	return tst
}

func (s *CmdLeiaContext) Identificador(i int) IIdentificadorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *CmdLeiaContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFECHAPAR, 0)
}

func (s *CmdLeiaContext) AllPONTEIRO() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserPONTEIRO)
}

func (s *CmdLeiaContext) PONTEIRO(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserPONTEIRO, i)
}

func (s *CmdLeiaContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserVIRGULA)
}

func (s *CmdLeiaContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVIRGULA, i)
}

func (s *CmdLeiaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdLeiaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdLeiaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmdLeia(s)
	}
}

func (s *CmdLeiaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmdLeia(s)
	}
}

func (s *CmdLeiaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmdLeia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) CmdLeia() (localctx ICmdLeiaContext) {
	localctx = NewCmdLeiaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, T3AlgumaParserRULE_cmdLeia)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(T3AlgumaParserLEIA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.Match(T3AlgumaParserABREPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T3AlgumaParserPONTEIRO {
		{
			p.SetState(337)
			p.Match(T3AlgumaParserPONTEIRO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(340)
		p.Identificador()
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserVIRGULA {
		{
			p.SetState(341)
			p.Match(T3AlgumaParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T3AlgumaParserPONTEIRO {
			{
				p.SetState(342)
				p.Match(T3AlgumaParserPONTEIRO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(345)
			p.Identificador()
		}

		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(351)
		p.Match(T3AlgumaParserFECHAPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdEscrevaContext is an interface to support dynamic dispatch.
type ICmdEscrevaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ESCREVA() antlr.TerminalNode
	ABREPAR() antlr.TerminalNode
	AllExpressao() []IExpressaoContext
	Expressao(i int) IExpressaoContext
	FECHAPAR() antlr.TerminalNode
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsCmdEscrevaContext differentiates from other interfaces.
	IsCmdEscrevaContext()
}

type CmdEscrevaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdEscrevaContext() *CmdEscrevaContext {
	var p = new(CmdEscrevaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdEscreva
	return p
}

func InitEmptyCmdEscrevaContext(p *CmdEscrevaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdEscreva
}

func (*CmdEscrevaContext) IsCmdEscrevaContext() {}

func NewCmdEscrevaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdEscrevaContext {
	var p = new(CmdEscrevaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmdEscreva

	return p
}

func (s *CmdEscrevaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdEscrevaContext) ESCREVA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserESCREVA, 0)
}

func (s *CmdEscrevaContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserABREPAR, 0)
}

func (s *CmdEscrevaContext) AllExpressao() []IExpressaoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressaoContext); ok {
			len++
		}
	}

	tst := make([]IExpressaoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressaoContext); ok {
			tst[i] = t.(IExpressaoContext)
			i++
		}
	}

	return tst
}

func (s *CmdEscrevaContext) Expressao(i int) IExpressaoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdEscrevaContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFECHAPAR, 0)
}

func (s *CmdEscrevaContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserVIRGULA)
}

func (s *CmdEscrevaContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVIRGULA, i)
}

func (s *CmdEscrevaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdEscrevaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdEscrevaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmdEscreva(s)
	}
}

func (s *CmdEscrevaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmdEscreva(s)
	}
}

func (s *CmdEscrevaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmdEscreva(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) CmdEscreva() (localctx ICmdEscrevaContext) {
	localctx = NewCmdEscrevaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, T3AlgumaParserRULE_cmdEscreva)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(T3AlgumaParserESCREVA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(354)
		p.Match(T3AlgumaParserABREPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(355)
		p.Expressao()
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserVIRGULA {
		{
			p.SetState(356)
			p.Match(T3AlgumaParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Expressao()
		}

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(363)
		p.Match(T3AlgumaParserFECHAPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdSeContext is an interface to support dynamic dispatch.
type ICmdSeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SE() antlr.TerminalNode
	Expressao() IExpressaoContext
	ENTAO() antlr.TerminalNode
	FIM_SE() antlr.TerminalNode
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext
	SENAO() antlr.TerminalNode

	// IsCmdSeContext differentiates from other interfaces.
	IsCmdSeContext()
}

type CmdSeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdSeContext() *CmdSeContext {
	var p = new(CmdSeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdSe
	return p
}

func InitEmptyCmdSeContext(p *CmdSeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdSe
}

func (*CmdSeContext) IsCmdSeContext() {}

func NewCmdSeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdSeContext {
	var p = new(CmdSeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmdSe

	return p
}

func (s *CmdSeContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdSeContext) SE() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserSE, 0)
}

func (s *CmdSeContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdSeContext) ENTAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserENTAO, 0)
}

func (s *CmdSeContext) FIM_SE() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFIM_SE, 0)
}

func (s *CmdSeContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CmdSeContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdSeContext) SENAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserSENAO, 0)
}

func (s *CmdSeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdSeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdSeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmdSe(s)
	}
}

func (s *CmdSeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmdSe(s)
	}
}

func (s *CmdSeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmdSe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) CmdSe() (localctx ICmdSeContext) {
	localctx = NewCmdSeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, T3AlgumaParserRULE_cmdSe)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(T3AlgumaParserSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Expressao()
	}
	{
		p.SetState(367)
		p.Match(T3AlgumaParserENTAO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(368)
			p.Cmd()
		}

		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T3AlgumaParserSENAO {
		{
			p.SetState(374)
			p.Match(T3AlgumaParserSENAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
			{
				p.SetState(375)
				p.Cmd()
			}

			p.SetState(380)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(383)
		p.Match(T3AlgumaParserFIM_SE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdCasoContext is an interface to support dynamic dispatch.
type ICmdCasoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASO() antlr.TerminalNode
	Exp_aritmetica() IExp_aritmeticaContext
	SEJA() antlr.TerminalNode
	Selecao() ISelecaoContext
	FIM_CASO() antlr.TerminalNode
	SENAO() antlr.TerminalNode
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCmdCasoContext differentiates from other interfaces.
	IsCmdCasoContext()
}

type CmdCasoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdCasoContext() *CmdCasoContext {
	var p = new(CmdCasoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdCaso
	return p
}

func InitEmptyCmdCasoContext(p *CmdCasoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdCaso
}

func (*CmdCasoContext) IsCmdCasoContext() {}

func NewCmdCasoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdCasoContext {
	var p = new(CmdCasoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmdCaso

	return p
}

func (s *CmdCasoContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdCasoContext) CASO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserCASO, 0)
}

func (s *CmdCasoContext) Exp_aritmetica() IExp_aritmeticaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_aritmeticaContext)
}

func (s *CmdCasoContext) SEJA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserSEJA, 0)
}

func (s *CmdCasoContext) Selecao() ISelecaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelecaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelecaoContext)
}

func (s *CmdCasoContext) FIM_CASO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFIM_CASO, 0)
}

func (s *CmdCasoContext) SENAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserSENAO, 0)
}

func (s *CmdCasoContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CmdCasoContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdCasoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdCasoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdCasoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmdCaso(s)
	}
}

func (s *CmdCasoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmdCaso(s)
	}
}

func (s *CmdCasoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmdCaso(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) CmdCaso() (localctx ICmdCasoContext) {
	localctx = NewCmdCasoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, T3AlgumaParserRULE_cmdCaso)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(T3AlgumaParserCASO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Exp_aritmetica()
	}
	{
		p.SetState(387)
		p.Match(T3AlgumaParserSEJA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
		p.Selecao()
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T3AlgumaParserSENAO {
		{
			p.SetState(389)
			p.Match(T3AlgumaParserSENAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
			{
				p.SetState(390)
				p.Cmd()
			}

			p.SetState(395)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(398)
		p.Match(T3AlgumaParserFIM_CASO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdParaContext is an interface to support dynamic dispatch.
type ICmdParaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp1 returns the exp1 rule contexts.
	GetExp1() IExp_aritmeticaContext

	// GetExp2 returns the exp2 rule contexts.
	GetExp2() IExp_aritmeticaContext

	// SetExp1 sets the exp1 rule contexts.
	SetExp1(IExp_aritmeticaContext)

	// SetExp2 sets the exp2 rule contexts.
	SetExp2(IExp_aritmeticaContext)

	// Getter signatures
	PARA() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ATRIBUICAO() antlr.TerminalNode
	ATE() antlr.TerminalNode
	FACA() antlr.TerminalNode
	FIM_PARA() antlr.TerminalNode
	AllExp_aritmetica() []IExp_aritmeticaContext
	Exp_aritmetica(i int) IExp_aritmeticaContext
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCmdParaContext differentiates from other interfaces.
	IsCmdParaContext()
}

type CmdParaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	exp1   IExp_aritmeticaContext
	exp2   IExp_aritmeticaContext
}

func NewEmptyCmdParaContext() *CmdParaContext {
	var p = new(CmdParaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdPara
	return p
}

func InitEmptyCmdParaContext(p *CmdParaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdPara
}

func (*CmdParaContext) IsCmdParaContext() {}

func NewCmdParaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdParaContext {
	var p = new(CmdParaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmdPara

	return p
}

func (s *CmdParaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdParaContext) GetExp1() IExp_aritmeticaContext { return s.exp1 }

func (s *CmdParaContext) GetExp2() IExp_aritmeticaContext { return s.exp2 }

func (s *CmdParaContext) SetExp1(v IExp_aritmeticaContext) { s.exp1 = v }

func (s *CmdParaContext) SetExp2(v IExp_aritmeticaContext) { s.exp2 = v }

func (s *CmdParaContext) PARA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserPARA, 0)
}

func (s *CmdParaContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserIDENT, 0)
}

func (s *CmdParaContext) ATRIBUICAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserATRIBUICAO, 0)
}

func (s *CmdParaContext) ATE() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserATE, 0)
}

func (s *CmdParaContext) FACA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFACA, 0)
}

func (s *CmdParaContext) FIM_PARA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFIM_PARA, 0)
}

func (s *CmdParaContext) AllExp_aritmetica() []IExp_aritmeticaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			len++
		}
	}

	tst := make([]IExp_aritmeticaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExp_aritmeticaContext); ok {
			tst[i] = t.(IExp_aritmeticaContext)
			i++
		}
	}

	return tst
}

func (s *CmdParaContext) Exp_aritmetica(i int) IExp_aritmeticaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_aritmeticaContext)
}

func (s *CmdParaContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CmdParaContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdParaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdParaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdParaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmdPara(s)
	}
}

func (s *CmdParaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmdPara(s)
	}
}

func (s *CmdParaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmdPara(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) CmdPara() (localctx ICmdParaContext) {
	localctx = NewCmdParaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, T3AlgumaParserRULE_cmdPara)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.Match(T3AlgumaParserPARA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(401)
		p.Match(T3AlgumaParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.Match(T3AlgumaParserATRIBUICAO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)

		var _x = p.Exp_aritmetica()

		localctx.(*CmdParaContext).exp1 = _x
	}
	{
		p.SetState(404)
		p.Match(T3AlgumaParserATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)

		var _x = p.Exp_aritmetica()

		localctx.(*CmdParaContext).exp2 = _x
	}
	{
		p.SetState(406)
		p.Match(T3AlgumaParserFACA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(407)
			p.Cmd()
		}

		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(413)
		p.Match(T3AlgumaParserFIM_PARA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdEnquantoContext is an interface to support dynamic dispatch.
type ICmdEnquantoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENQUANTO() antlr.TerminalNode
	Expressao() IExpressaoContext
	FACA() antlr.TerminalNode
	FIM_ENQUANTO() antlr.TerminalNode
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCmdEnquantoContext differentiates from other interfaces.
	IsCmdEnquantoContext()
}

type CmdEnquantoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdEnquantoContext() *CmdEnquantoContext {
	var p = new(CmdEnquantoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdEnquanto
	return p
}

func InitEmptyCmdEnquantoContext(p *CmdEnquantoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdEnquanto
}

func (*CmdEnquantoContext) IsCmdEnquantoContext() {}

func NewCmdEnquantoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdEnquantoContext {
	var p = new(CmdEnquantoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmdEnquanto

	return p
}

func (s *CmdEnquantoContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdEnquantoContext) ENQUANTO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserENQUANTO, 0)
}

func (s *CmdEnquantoContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdEnquantoContext) FACA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFACA, 0)
}

func (s *CmdEnquantoContext) FIM_ENQUANTO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFIM_ENQUANTO, 0)
}

func (s *CmdEnquantoContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CmdEnquantoContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdEnquantoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdEnquantoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdEnquantoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmdEnquanto(s)
	}
}

func (s *CmdEnquantoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmdEnquanto(s)
	}
}

func (s *CmdEnquantoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmdEnquanto(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) CmdEnquanto() (localctx ICmdEnquantoContext) {
	localctx = NewCmdEnquantoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, T3AlgumaParserRULE_cmdEnquanto)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(415)
		p.Match(T3AlgumaParserENQUANTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(416)
		p.Expressao()
	}
	{
		p.SetState(417)
		p.Match(T3AlgumaParserFACA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(418)
			p.Cmd()
		}

		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(424)
		p.Match(T3AlgumaParserFIM_ENQUANTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdFacaContext is an interface to support dynamic dispatch.
type ICmdFacaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FACA() antlr.TerminalNode
	ATE() antlr.TerminalNode
	Expressao() IExpressaoContext
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCmdFacaContext differentiates from other interfaces.
	IsCmdFacaContext()
}

type CmdFacaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdFacaContext() *CmdFacaContext {
	var p = new(CmdFacaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdFaca
	return p
}

func InitEmptyCmdFacaContext(p *CmdFacaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdFaca
}

func (*CmdFacaContext) IsCmdFacaContext() {}

func NewCmdFacaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdFacaContext {
	var p = new(CmdFacaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmdFaca

	return p
}

func (s *CmdFacaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdFacaContext) FACA() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFACA, 0)
}

func (s *CmdFacaContext) ATE() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserATE, 0)
}

func (s *CmdFacaContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdFacaContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *CmdFacaContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *CmdFacaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdFacaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdFacaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmdFaca(s)
	}
}

func (s *CmdFacaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmdFaca(s)
	}
}

func (s *CmdFacaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmdFaca(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) CmdFaca() (localctx ICmdFacaContext) {
	localctx = NewCmdFacaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, T3AlgumaParserRULE_cmdFaca)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(T3AlgumaParserFACA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(427)
			p.Cmd()
		}

		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(433)
		p.Match(T3AlgumaParserATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(434)
		p.Expressao()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdAtribuicaoContext is an interface to support dynamic dispatch.
type ICmdAtribuicaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() IIdentificadorContext
	ATRIBUICAO() antlr.TerminalNode
	Expressao() IExpressaoContext
	PONTEIRO() antlr.TerminalNode

	// IsCmdAtribuicaoContext differentiates from other interfaces.
	IsCmdAtribuicaoContext()
}

type CmdAtribuicaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdAtribuicaoContext() *CmdAtribuicaoContext {
	var p = new(CmdAtribuicaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdAtribuicao
	return p
}

func InitEmptyCmdAtribuicaoContext(p *CmdAtribuicaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdAtribuicao
}

func (*CmdAtribuicaoContext) IsCmdAtribuicaoContext() {}

func NewCmdAtribuicaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdAtribuicaoContext {
	var p = new(CmdAtribuicaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmdAtribuicao

	return p
}

func (s *CmdAtribuicaoContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdAtribuicaoContext) Identificador() IIdentificadorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificadorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificadorContext)
}

func (s *CmdAtribuicaoContext) ATRIBUICAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserATRIBUICAO, 0)
}

func (s *CmdAtribuicaoContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdAtribuicaoContext) PONTEIRO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserPONTEIRO, 0)
}

func (s *CmdAtribuicaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdAtribuicaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdAtribuicaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmdAtribuicao(s)
	}
}

func (s *CmdAtribuicaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmdAtribuicao(s)
	}
}

func (s *CmdAtribuicaoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmdAtribuicao(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) CmdAtribuicao() (localctx ICmdAtribuicaoContext) {
	localctx = NewCmdAtribuicaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, T3AlgumaParserRULE_cmdAtribuicao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T3AlgumaParserPONTEIRO {
		{
			p.SetState(436)
			p.Match(T3AlgumaParserPONTEIRO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(439)
		p.Identificador()
	}
	{
		p.SetState(440)
		p.Match(T3AlgumaParserATRIBUICAO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.Expressao()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdChamadaContext is an interface to support dynamic dispatch.
type ICmdChamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	ABREPAR() antlr.TerminalNode
	AllExpressao() []IExpressaoContext
	Expressao(i int) IExpressaoContext
	FECHAPAR() antlr.TerminalNode
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsCmdChamadaContext differentiates from other interfaces.
	IsCmdChamadaContext()
}

type CmdChamadaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdChamadaContext() *CmdChamadaContext {
	var p = new(CmdChamadaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdChamada
	return p
}

func InitEmptyCmdChamadaContext(p *CmdChamadaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdChamada
}

func (*CmdChamadaContext) IsCmdChamadaContext() {}

func NewCmdChamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdChamadaContext {
	var p = new(CmdChamadaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmdChamada

	return p
}

func (s *CmdChamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdChamadaContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserIDENT, 0)
}

func (s *CmdChamadaContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserABREPAR, 0)
}

func (s *CmdChamadaContext) AllExpressao() []IExpressaoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressaoContext); ok {
			len++
		}
	}

	tst := make([]IExpressaoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressaoContext); ok {
			tst[i] = t.(IExpressaoContext)
			i++
		}
	}

	return tst
}

func (s *CmdChamadaContext) Expressao(i int) IExpressaoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdChamadaContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFECHAPAR, 0)
}

func (s *CmdChamadaContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserVIRGULA)
}

func (s *CmdChamadaContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVIRGULA, i)
}

func (s *CmdChamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdChamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdChamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmdChamada(s)
	}
}

func (s *CmdChamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmdChamada(s)
	}
}

func (s *CmdChamadaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmdChamada(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) CmdChamada() (localctx ICmdChamadaContext) {
	localctx = NewCmdChamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, T3AlgumaParserRULE_cmdChamada)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)
		p.Match(T3AlgumaParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(444)
		p.Match(T3AlgumaParserABREPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(445)
		p.Expressao()
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserVIRGULA {
		{
			p.SetState(446)
			p.Match(T3AlgumaParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
			p.Expressao()
		}

		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(453)
		p.Match(T3AlgumaParserFECHAPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmdRetorneContext is an interface to support dynamic dispatch.
type ICmdRetorneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETORNE() antlr.TerminalNode
	Expressao() IExpressaoContext

	// IsCmdRetorneContext differentiates from other interfaces.
	IsCmdRetorneContext()
}

type CmdRetorneContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdRetorneContext() *CmdRetorneContext {
	var p = new(CmdRetorneContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdRetorne
	return p
}

func InitEmptyCmdRetorneContext(p *CmdRetorneContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_cmdRetorne
}

func (*CmdRetorneContext) IsCmdRetorneContext() {}

func NewCmdRetorneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdRetorneContext {
	var p = new(CmdRetorneContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_cmdRetorne

	return p
}

func (s *CmdRetorneContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdRetorneContext) RETORNE() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserRETORNE, 0)
}

func (s *CmdRetorneContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *CmdRetorneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdRetorneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdRetorneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterCmdRetorne(s)
	}
}

func (s *CmdRetorneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitCmdRetorne(s)
	}
}

func (s *CmdRetorneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitCmdRetorne(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) CmdRetorne() (localctx ICmdRetorneContext) {
	localctx = NewCmdRetorneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, T3AlgumaParserRULE_cmdRetorne)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Match(T3AlgumaParserRETORNE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.Expressao()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelecaoContext is an interface to support dynamic dispatch.
type ISelecaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllItem_selecao() []IItem_selecaoContext
	Item_selecao(i int) IItem_selecaoContext

	// IsSelecaoContext differentiates from other interfaces.
	IsSelecaoContext()
}

type SelecaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelecaoContext() *SelecaoContext {
	var p = new(SelecaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_selecao
	return p
}

func InitEmptySelecaoContext(p *SelecaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_selecao
}

func (*SelecaoContext) IsSelecaoContext() {}

func NewSelecaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelecaoContext {
	var p = new(SelecaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_selecao

	return p
}

func (s *SelecaoContext) GetParser() antlr.Parser { return s.parser }

func (s *SelecaoContext) AllItem_selecao() []IItem_selecaoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IItem_selecaoContext); ok {
			len++
		}
	}

	tst := make([]IItem_selecaoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IItem_selecaoContext); ok {
			tst[i] = t.(IItem_selecaoContext)
			i++
		}
	}

	return tst
}

func (s *SelecaoContext) Item_selecao(i int) IItem_selecaoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItem_selecaoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItem_selecaoContext)
}

func (s *SelecaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelecaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelecaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterSelecao(s)
	}
}

func (s *SelecaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitSelecao(s)
	}
}

func (s *SelecaoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitSelecao(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Selecao() (localctx ISelecaoContext) {
	localctx = NewSelecaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, T3AlgumaParserRULE_selecao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserSUBTRACAO || _la == T3AlgumaParserNUM_INT {
		{
			p.SetState(458)
			p.Item_selecao()
		}

		p.SetState(463)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IItem_selecaoContext is an interface to support dynamic dispatch.
type IItem_selecaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constantes() IConstantesContext
	DOISPONTOS() antlr.TerminalNode
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsItem_selecaoContext differentiates from other interfaces.
	IsItem_selecaoContext()
}

type Item_selecaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItem_selecaoContext() *Item_selecaoContext {
	var p = new(Item_selecaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_item_selecao
	return p
}

func InitEmptyItem_selecaoContext(p *Item_selecaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_item_selecao
}

func (*Item_selecaoContext) IsItem_selecaoContext() {}

func NewItem_selecaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Item_selecaoContext {
	var p = new(Item_selecaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_item_selecao

	return p
}

func (s *Item_selecaoContext) GetParser() antlr.Parser { return s.parser }

func (s *Item_selecaoContext) Constantes() IConstantesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantesContext)
}

func (s *Item_selecaoContext) DOISPONTOS() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserDOISPONTOS, 0)
}

func (s *Item_selecaoContext) AllCmd() []ICmdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICmdContext); ok {
			len++
		}
	}

	tst := make([]ICmdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICmdContext); ok {
			tst[i] = t.(ICmdContext)
			i++
		}
	}

	return tst
}

func (s *Item_selecaoContext) Cmd(i int) ICmdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *Item_selecaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Item_selecaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Item_selecaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterItem_selecao(s)
	}
}

func (s *Item_selecaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitItem_selecao(s)
	}
}

func (s *Item_selecaoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitItem_selecao(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Item_selecao() (localctx IItem_selecaoContext) {
	localctx = NewItem_selecaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, T3AlgumaParserRULE_item_selecao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		p.Constantes()
	}
	{
		p.SetState(465)
		p.Match(T3AlgumaParserDOISPONTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(466)
			p.Cmd()
		}

		p.SetState(471)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantesContext is an interface to support dynamic dispatch.
type IConstantesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumero_intervalo() []INumero_intervaloContext
	Numero_intervalo(i int) INumero_intervaloContext
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsConstantesContext differentiates from other interfaces.
	IsConstantesContext()
}

type ConstantesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantesContext() *ConstantesContext {
	var p = new(ConstantesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_constantes
	return p
}

func InitEmptyConstantesContext(p *ConstantesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_constantes
}

func (*ConstantesContext) IsConstantesContext() {}

func NewConstantesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantesContext {
	var p = new(ConstantesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_constantes

	return p
}

func (s *ConstantesContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantesContext) AllNumero_intervalo() []INumero_intervaloContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumero_intervaloContext); ok {
			len++
		}
	}

	tst := make([]INumero_intervaloContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumero_intervaloContext); ok {
			tst[i] = t.(INumero_intervaloContext)
			i++
		}
	}

	return tst
}

func (s *ConstantesContext) Numero_intervalo(i int) INumero_intervaloContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumero_intervaloContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumero_intervaloContext)
}

func (s *ConstantesContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserVIRGULA)
}

func (s *ConstantesContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVIRGULA, i)
}

func (s *ConstantesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterConstantes(s)
	}
}

func (s *ConstantesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitConstantes(s)
	}
}

func (s *ConstantesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitConstantes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Constantes() (localctx IConstantesContext) {
	localctx = NewConstantesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, T3AlgumaParserRULE_constantes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		p.Numero_intervalo()
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserVIRGULA {
		{
			p.SetState(473)
			p.Match(T3AlgumaParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(474)
			p.Numero_intervalo()
		}

		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumero_intervaloContext is an interface to support dynamic dispatch.
type INumero_intervaloContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNUM_INT() []antlr.TerminalNode
	NUM_INT(i int) antlr.TerminalNode
	AllOp_unario() []IOp_unarioContext
	Op_unario(i int) IOp_unarioContext
	INTERVALO() antlr.TerminalNode

	// IsNumero_intervaloContext differentiates from other interfaces.
	IsNumero_intervaloContext()
}

type Numero_intervaloContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumero_intervaloContext() *Numero_intervaloContext {
	var p = new(Numero_intervaloContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_numero_intervalo
	return p
}

func InitEmptyNumero_intervaloContext(p *Numero_intervaloContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_numero_intervalo
}

func (*Numero_intervaloContext) IsNumero_intervaloContext() {}

func NewNumero_intervaloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numero_intervaloContext {
	var p = new(Numero_intervaloContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_numero_intervalo

	return p
}

func (s *Numero_intervaloContext) GetParser() antlr.Parser { return s.parser }

func (s *Numero_intervaloContext) AllNUM_INT() []antlr.TerminalNode {
	return s.GetTokens(T3AlgumaParserNUM_INT)
}

func (s *Numero_intervaloContext) NUM_INT(i int) antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserNUM_INT, i)
}

func (s *Numero_intervaloContext) AllOp_unario() []IOp_unarioContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp_unarioContext); ok {
			len++
		}
	}

	tst := make([]IOp_unarioContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp_unarioContext); ok {
			tst[i] = t.(IOp_unarioContext)
			i++
		}
	}

	return tst
}

func (s *Numero_intervaloContext) Op_unario(i int) IOp_unarioContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_unarioContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_unarioContext)
}

func (s *Numero_intervaloContext) INTERVALO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserINTERVALO, 0)
}

func (s *Numero_intervaloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numero_intervaloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numero_intervaloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterNumero_intervalo(s)
	}
}

func (s *Numero_intervaloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitNumero_intervalo(s)
	}
}

func (s *Numero_intervaloContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitNumero_intervalo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Numero_intervalo() (localctx INumero_intervaloContext) {
	localctx = NewNumero_intervaloContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, T3AlgumaParserRULE_numero_intervalo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T3AlgumaParserSUBTRACAO {
		{
			p.SetState(480)
			p.Op_unario()
		}

	}
	{
		p.SetState(483)
		p.Match(T3AlgumaParserNUM_INT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T3AlgumaParserINTERVALO {
		{
			p.SetState(484)
			p.Match(T3AlgumaParserINTERVALO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(486)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T3AlgumaParserSUBTRACAO {
			{
				p.SetState(485)
				p.Op_unario()
			}

		}
		{
			p.SetState(488)
			p.Match(T3AlgumaParserNUM_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExp_relacionalContext is an interface to support dynamic dispatch.
type IExp_relacionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp_aritmetica() []IExp_aritmeticaContext
	Exp_aritmetica(i int) IExp_aritmeticaContext
	Op_relacional() IOp_relacionalContext

	// IsExp_relacionalContext differentiates from other interfaces.
	IsExp_relacionalContext()
}

type Exp_relacionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp_relacionalContext() *Exp_relacionalContext {
	var p = new(Exp_relacionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_exp_relacional
	return p
}

func InitEmptyExp_relacionalContext(p *Exp_relacionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_exp_relacional
}

func (*Exp_relacionalContext) IsExp_relacionalContext() {}

func NewExp_relacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_relacionalContext {
	var p = new(Exp_relacionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_exp_relacional

	return p
}

func (s *Exp_relacionalContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp_relacionalContext) AllExp_aritmetica() []IExp_aritmeticaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			len++
		}
	}

	tst := make([]IExp_aritmeticaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExp_aritmeticaContext); ok {
			tst[i] = t.(IExp_aritmeticaContext)
			i++
		}
	}

	return tst
}

func (s *Exp_relacionalContext) Exp_aritmetica(i int) IExp_aritmeticaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_aritmeticaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_aritmeticaContext)
}

func (s *Exp_relacionalContext) Op_relacional() IOp_relacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_relacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_relacionalContext)
}

func (s *Exp_relacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp_relacionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp_relacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterExp_relacional(s)
	}
}

func (s *Exp_relacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitExp_relacional(s)
	}
}

func (s *Exp_relacionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitExp_relacional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Exp_relacional() (localctx IExp_relacionalContext) {
	localctx = NewExp_relacionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, T3AlgumaParserRULE_exp_relacional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		p.Exp_aritmetica()
	}
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69269232549888) != 0 {
		{
			p.SetState(492)
			p.Op_relacional()
		}
		{
			p.SetState(493)
			p.Exp_aritmetica()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_relacionalContext is an interface to support dynamic dispatch.
type IOp_relacionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IGUAL() antlr.TerminalNode
	DIFERENTE() antlr.TerminalNode
	MAIORIGUAL() antlr.TerminalNode
	MENORIGUAL() antlr.TerminalNode
	MAIOR() antlr.TerminalNode
	MENOR() antlr.TerminalNode

	// IsOp_relacionalContext differentiates from other interfaces.
	IsOp_relacionalContext()
}

type Op_relacionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_relacionalContext() *Op_relacionalContext {
	var p = new(Op_relacionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op_relacional
	return p
}

func InitEmptyOp_relacionalContext(p *Op_relacionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op_relacional
}

func (*Op_relacionalContext) IsOp_relacionalContext() {}

func NewOp_relacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_relacionalContext {
	var p = new(Op_relacionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_op_relacional

	return p
}

func (s *Op_relacionalContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_relacionalContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserIGUAL, 0)
}

func (s *Op_relacionalContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserDIFERENTE, 0)
}

func (s *Op_relacionalContext) MAIORIGUAL() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserMAIORIGUAL, 0)
}

func (s *Op_relacionalContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserMENORIGUAL, 0)
}

func (s *Op_relacionalContext) MAIOR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserMAIOR, 0)
}

func (s *Op_relacionalContext) MENOR() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserMENOR, 0)
}

func (s *Op_relacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_relacionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_relacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterOp_relacional(s)
	}
}

func (s *Op_relacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitOp_relacional(s)
	}
}

func (s *Op_relacionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitOp_relacional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Op_relacional() (localctx IOp_relacionalContext) {
	localctx = NewOp_relacionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, T3AlgumaParserRULE_op_relacional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(497)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69269232549888) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressaoContext is an interface to support dynamic dispatch.
type IExpressaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTermo_logico() []ITermo_logicoContext
	Termo_logico(i int) ITermo_logicoContext
	AllOp_logico_1() []IOp_logico_1Context
	Op_logico_1(i int) IOp_logico_1Context

	// IsExpressaoContext differentiates from other interfaces.
	IsExpressaoContext()
}

type ExpressaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressaoContext() *ExpressaoContext {
	var p = new(ExpressaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_expressao
	return p
}

func InitEmptyExpressaoContext(p *ExpressaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_expressao
}

func (*ExpressaoContext) IsExpressaoContext() {}

func NewExpressaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressaoContext {
	var p = new(ExpressaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_expressao

	return p
}

func (s *ExpressaoContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressaoContext) AllTermo_logico() []ITermo_logicoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermo_logicoContext); ok {
			len++
		}
	}

	tst := make([]ITermo_logicoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermo_logicoContext); ok {
			tst[i] = t.(ITermo_logicoContext)
			i++
		}
	}

	return tst
}

func (s *ExpressaoContext) Termo_logico(i int) ITermo_logicoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermo_logicoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermo_logicoContext)
}

func (s *ExpressaoContext) AllOp_logico_1() []IOp_logico_1Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp_logico_1Context); ok {
			len++
		}
	}

	tst := make([]IOp_logico_1Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp_logico_1Context); ok {
			tst[i] = t.(IOp_logico_1Context)
			i++
		}
	}

	return tst
}

func (s *ExpressaoContext) Op_logico_1(i int) IOp_logico_1Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_logico_1Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_logico_1Context)
}

func (s *ExpressaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterExpressao(s)
	}
}

func (s *ExpressaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitExpressao(s)
	}
}

func (s *ExpressaoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitExpressao(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Expressao() (localctx IExpressaoContext) {
	localctx = NewExpressaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, T3AlgumaParserRULE_expressao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)
		p.Termo_logico()
	}
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserOU {
		{
			p.SetState(500)
			p.Op_logico_1()
		}
		{
			p.SetState(501)
			p.Termo_logico()
		}

		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermo_logicoContext is an interface to support dynamic dispatch.
type ITermo_logicoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFator_logico() []IFator_logicoContext
	Fator_logico(i int) IFator_logicoContext
	AllOp_logico_2() []IOp_logico_2Context
	Op_logico_2(i int) IOp_logico_2Context

	// IsTermo_logicoContext differentiates from other interfaces.
	IsTermo_logicoContext()
}

type Termo_logicoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermo_logicoContext() *Termo_logicoContext {
	var p = new(Termo_logicoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_termo_logico
	return p
}

func InitEmptyTermo_logicoContext(p *Termo_logicoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_termo_logico
}

func (*Termo_logicoContext) IsTermo_logicoContext() {}

func NewTermo_logicoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Termo_logicoContext {
	var p = new(Termo_logicoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_termo_logico

	return p
}

func (s *Termo_logicoContext) GetParser() antlr.Parser { return s.parser }

func (s *Termo_logicoContext) AllFator_logico() []IFator_logicoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFator_logicoContext); ok {
			len++
		}
	}

	tst := make([]IFator_logicoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFator_logicoContext); ok {
			tst[i] = t.(IFator_logicoContext)
			i++
		}
	}

	return tst
}

func (s *Termo_logicoContext) Fator_logico(i int) IFator_logicoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFator_logicoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFator_logicoContext)
}

func (s *Termo_logicoContext) AllOp_logico_2() []IOp_logico_2Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOp_logico_2Context); ok {
			len++
		}
	}

	tst := make([]IOp_logico_2Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOp_logico_2Context); ok {
			tst[i] = t.(IOp_logico_2Context)
			i++
		}
	}

	return tst
}

func (s *Termo_logicoContext) Op_logico_2(i int) IOp_logico_2Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_logico_2Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_logico_2Context)
}

func (s *Termo_logicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Termo_logicoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Termo_logicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterTermo_logico(s)
	}
}

func (s *Termo_logicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitTermo_logico(s)
	}
}

func (s *Termo_logicoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitTermo_logico(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Termo_logico() (localctx ITermo_logicoContext) {
	localctx = NewTermo_logicoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, T3AlgumaParserRULE_termo_logico)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(508)
		p.Fator_logico()
	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T3AlgumaParserE {
		{
			p.SetState(509)
			p.Op_logico_2()
		}
		{
			p.SetState(510)
			p.Fator_logico()
		}

		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFator_logicoContext is an interface to support dynamic dispatch.
type IFator_logicoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Parcela_logica() IParcela_logicaContext
	NAO() antlr.TerminalNode

	// IsFator_logicoContext differentiates from other interfaces.
	IsFator_logicoContext()
}

type Fator_logicoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFator_logicoContext() *Fator_logicoContext {
	var p = new(Fator_logicoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_fator_logico
	return p
}

func InitEmptyFator_logicoContext(p *Fator_logicoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_fator_logico
}

func (*Fator_logicoContext) IsFator_logicoContext() {}

func NewFator_logicoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fator_logicoContext {
	var p = new(Fator_logicoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_fator_logico

	return p
}

func (s *Fator_logicoContext) GetParser() antlr.Parser { return s.parser }

func (s *Fator_logicoContext) Parcela_logica() IParcela_logicaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParcela_logicaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParcela_logicaContext)
}

func (s *Fator_logicoContext) NAO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserNAO, 0)
}

func (s *Fator_logicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fator_logicoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fator_logicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterFator_logico(s)
	}
}

func (s *Fator_logicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitFator_logico(s)
	}
}

func (s *Fator_logicoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitFator_logico(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Fator_logico() (localctx IFator_logicoContext) {
	localctx = NewFator_logicoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, T3AlgumaParserRULE_fator_logico)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T3AlgumaParserNAO {
		{
			p.SetState(517)
			p.Match(T3AlgumaParserNAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(520)
		p.Parcela_logica()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParcela_logicaContext is an interface to support dynamic dispatch.
type IParcela_logicaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VERDADEIRO() antlr.TerminalNode
	FALSO() antlr.TerminalNode
	Exp_relacional() IExp_relacionalContext

	// IsParcela_logicaContext differentiates from other interfaces.
	IsParcela_logicaContext()
}

type Parcela_logicaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParcela_logicaContext() *Parcela_logicaContext {
	var p = new(Parcela_logicaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parcela_logica
	return p
}

func InitEmptyParcela_logicaContext(p *Parcela_logicaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_parcela_logica
}

func (*Parcela_logicaContext) IsParcela_logicaContext() {}

func NewParcela_logicaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parcela_logicaContext {
	var p = new(Parcela_logicaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_parcela_logica

	return p
}

func (s *Parcela_logicaContext) GetParser() antlr.Parser { return s.parser }

func (s *Parcela_logicaContext) VERDADEIRO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserVERDADEIRO, 0)
}

func (s *Parcela_logicaContext) FALSO() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserFALSO, 0)
}

func (s *Parcela_logicaContext) Exp_relacional() IExp_relacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_relacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_relacionalContext)
}

func (s *Parcela_logicaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parcela_logicaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parcela_logicaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterParcela_logica(s)
	}
}

func (s *Parcela_logicaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitParcela_logica(s)
	}
}

func (s *Parcela_logicaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitParcela_logica(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Parcela_logica() (localctx IParcela_logicaContext) {
	localctx = NewParcela_logicaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, T3AlgumaParserRULE_parcela_logica)
	var _la int

	p.SetState(524)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T3AlgumaParserVERDADEIRO, T3AlgumaParserFALSO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(522)
			_la = p.GetTokenStream().LA(1)

			if !(_la == T3AlgumaParserVERDADEIRO || _la == T3AlgumaParserFALSO) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case T3AlgumaParserABREPAR, T3AlgumaParserSUBTRACAO, T3AlgumaParserPONTEIRO, T3AlgumaParserENDERECO, T3AlgumaParserNUM_INT, T3AlgumaParserNUM_REAL, T3AlgumaParserIDENT, T3AlgumaParserCADEIA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(523)
			p.Exp_relacional()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_logico_1Context is an interface to support dynamic dispatch.
type IOp_logico_1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OU() antlr.TerminalNode

	// IsOp_logico_1Context differentiates from other interfaces.
	IsOp_logico_1Context()
}

type Op_logico_1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_logico_1Context() *Op_logico_1Context {
	var p = new(Op_logico_1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op_logico_1
	return p
}

func InitEmptyOp_logico_1Context(p *Op_logico_1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op_logico_1
}

func (*Op_logico_1Context) IsOp_logico_1Context() {}

func NewOp_logico_1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_logico_1Context {
	var p = new(Op_logico_1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_op_logico_1

	return p
}

func (s *Op_logico_1Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_logico_1Context) OU() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserOU, 0)
}

func (s *Op_logico_1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_logico_1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_logico_1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterOp_logico_1(s)
	}
}

func (s *Op_logico_1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitOp_logico_1(s)
	}
}

func (s *Op_logico_1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitOp_logico_1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Op_logico_1() (localctx IOp_logico_1Context) {
	localctx = NewOp_logico_1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, T3AlgumaParserRULE_op_logico_1)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(526)
		p.Match(T3AlgumaParserOU)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_logico_2Context is an interface to support dynamic dispatch.
type IOp_logico_2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	E() antlr.TerminalNode

	// IsOp_logico_2Context differentiates from other interfaces.
	IsOp_logico_2Context()
}

type Op_logico_2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_logico_2Context() *Op_logico_2Context {
	var p = new(Op_logico_2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op_logico_2
	return p
}

func InitEmptyOp_logico_2Context(p *Op_logico_2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T3AlgumaParserRULE_op_logico_2
}

func (*Op_logico_2Context) IsOp_logico_2Context() {}

func NewOp_logico_2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_logico_2Context {
	var p = new(Op_logico_2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T3AlgumaParserRULE_op_logico_2

	return p
}

func (s *Op_logico_2Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_logico_2Context) E() antlr.TerminalNode {
	return s.GetToken(T3AlgumaParserE, 0)
}

func (s *Op_logico_2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_logico_2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_logico_2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.EnterOp_logico_2(s)
	}
}

func (s *Op_logico_2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T3AlgumaListener); ok {
		listenerT.ExitOp_logico_2(s)
	}
}

func (s *Op_logico_2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T3AlgumaVisitor:
		return t.VisitOp_logico_2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T3AlgumaParser) Op_logico_2() (localctx IOp_logico_2Context) {
	localctx = NewOp_logico_2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, T3AlgumaParserRULE_op_logico_2)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
		p.Match(T3AlgumaParserE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
