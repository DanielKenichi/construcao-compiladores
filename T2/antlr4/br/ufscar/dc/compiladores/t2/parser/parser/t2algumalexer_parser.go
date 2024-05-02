// Code generated from T2AlgumaLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // T2AlgumaLexer

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

type T2AlgumaLexerParser struct {
	*antlr.BaseParser
}

var T2AlgumaLexerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func t2algumalexerParserInit() {
	staticData := &T2AlgumaLexerParserStaticData
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
		4, 1, 69, 528, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 102, 8, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 117,
		8, 2, 1, 3, 1, 3, 1, 3, 5, 3, 122, 8, 3, 10, 3, 12, 3, 125, 9, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 133, 8, 4, 10, 4, 12, 4, 136, 9, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 142, 8, 4, 10, 4, 12, 4, 145, 9, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 5, 5, 151, 8, 5, 10, 5, 12, 5, 154, 9, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 5, 6, 160, 8, 6, 10, 6, 12, 6, 163, 9, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 5, 7, 169, 8, 7, 10, 7, 12, 7, 172, 9, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 3, 11, 181, 8, 11, 1, 11, 1, 11, 3, 11, 185, 8, 11,
		1, 12, 3, 12, 188, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5,
		12, 196, 8, 12, 10, 12, 12, 12, 199, 9, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 209, 8, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 3, 14, 216, 8, 14, 1, 15, 1, 15, 3, 15, 220, 8, 15, 1, 16, 1,
		16, 1, 17, 3, 17, 225, 8, 17, 1, 17, 1, 17, 3, 17, 229, 8, 17, 1, 18, 1,
		18, 1, 19, 1, 19, 5, 19, 235, 8, 19, 10, 19, 12, 19, 238, 9, 19, 1, 19,
		1, 19, 1, 20, 3, 20, 243, 8, 20, 1, 20, 1, 20, 1, 20, 5, 20, 248, 8, 20,
		10, 20, 12, 20, 251, 9, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5,
		21, 259, 8, 21, 10, 21, 12, 21, 262, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 268, 8, 22, 1, 22, 1, 22, 5, 22, 272, 8, 22, 10, 22, 12, 22, 275,
		9, 22, 1, 22, 5, 22, 278, 8, 22, 10, 22, 12, 22, 281, 9, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 3, 22, 288, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		5, 22, 294, 8, 22, 10, 22, 12, 22, 297, 9, 22, 1, 22, 5, 22, 300, 8, 22,
		10, 22, 12, 22, 303, 9, 22, 1, 22, 1, 22, 3, 22, 307, 8, 22, 1, 23, 5,
		23, 310, 8, 23, 10, 23, 12, 23, 313, 9, 23, 1, 23, 5, 23, 316, 8, 23, 10,
		23, 12, 23, 319, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 3, 24, 331, 8, 24, 1, 25, 1, 25, 1, 25, 3, 25, 336,
		8, 25, 1, 25, 1, 25, 1, 25, 3, 25, 341, 8, 25, 1, 25, 5, 25, 344, 8, 25,
		10, 25, 12, 25, 347, 9, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 5, 26, 356, 8, 26, 10, 26, 12, 26, 359, 9, 26, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 27, 1, 27, 5, 27, 367, 8, 27, 10, 27, 12, 27, 370, 9, 27, 1,
		27, 1, 27, 5, 27, 374, 8, 27, 10, 27, 12, 27, 377, 9, 27, 3, 27, 379, 8,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 389,
		8, 28, 10, 28, 12, 28, 392, 9, 28, 3, 28, 394, 8, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 406, 8, 29,
		10, 29, 12, 29, 409, 9, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 5,
		30, 417, 8, 30, 10, 30, 12, 30, 420, 9, 30, 1, 30, 1, 30, 1, 31, 1, 31,
		5, 31, 426, 8, 31, 10, 31, 12, 31, 429, 9, 31, 1, 31, 1, 31, 1, 31, 1,
		32, 3, 32, 435, 8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 5, 33, 446, 8, 33, 10, 33, 12, 33, 449, 9, 33, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 34, 1, 35, 5, 35, 457, 8, 35, 10, 35, 12, 35, 460,
		9, 35, 1, 36, 1, 36, 1, 36, 5, 36, 465, 8, 36, 10, 36, 12, 36, 468, 9,
		36, 1, 37, 1, 37, 1, 37, 5, 37, 473, 8, 37, 10, 37, 12, 37, 476, 9, 37,
		1, 38, 3, 38, 479, 8, 38, 1, 38, 1, 38, 1, 38, 3, 38, 484, 8, 38, 1, 38,
		3, 38, 487, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 493, 8, 39, 1, 40,
		1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 501, 8, 41, 10, 41, 12, 41, 504,
		9, 41, 1, 42, 1, 42, 1, 42, 1, 42, 5, 42, 510, 8, 42, 10, 42, 12, 42, 513,
		9, 42, 1, 43, 3, 43, 516, 8, 43, 1, 43, 1, 43, 1, 44, 1, 44, 3, 44, 522,
		8, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 0, 0, 47, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
		84, 86, 88, 90, 92, 0, 6, 1, 0, 55, 56, 2, 0, 53, 53, 57, 57, 1, 0, 6,
		9, 3, 0, 10, 11, 62, 63, 65, 65, 1, 0, 40, 45, 1, 0, 10, 11, 548, 0, 94,
		1, 0, 0, 0, 2, 101, 1, 0, 0, 0, 4, 116, 1, 0, 0, 0, 6, 118, 1, 0, 0, 0,
		8, 129, 1, 0, 0, 0, 10, 146, 1, 0, 0, 0, 12, 155, 1, 0, 0, 0, 14, 164,
		1, 0, 0, 0, 16, 173, 1, 0, 0, 0, 18, 175, 1, 0, 0, 0, 20, 177, 1, 0, 0,
		0, 22, 184, 1, 0, 0, 0, 24, 208, 1, 0, 0, 0, 26, 210, 1, 0, 0, 0, 28, 215,
		1, 0, 0, 0, 30, 219, 1, 0, 0, 0, 32, 221, 1, 0, 0, 0, 34, 224, 1, 0, 0,
		0, 36, 230, 1, 0, 0, 0, 38, 232, 1, 0, 0, 0, 40, 242, 1, 0, 0, 0, 42, 255,
		1, 0, 0, 0, 44, 306, 1, 0, 0, 0, 46, 311, 1, 0, 0, 0, 48, 330, 1, 0, 0,
		0, 50, 332, 1, 0, 0, 0, 52, 350, 1, 0, 0, 0, 54, 362, 1, 0, 0, 0, 56, 382,
		1, 0, 0, 0, 58, 397, 1, 0, 0, 0, 60, 412, 1, 0, 0, 0, 62, 423, 1, 0, 0,
		0, 64, 434, 1, 0, 0, 0, 66, 440, 1, 0, 0, 0, 68, 452, 1, 0, 0, 0, 70, 458,
		1, 0, 0, 0, 72, 461, 1, 0, 0, 0, 74, 469, 1, 0, 0, 0, 76, 478, 1, 0, 0,
		0, 78, 488, 1, 0, 0, 0, 80, 494, 1, 0, 0, 0, 82, 496, 1, 0, 0, 0, 84, 505,
		1, 0, 0, 0, 86, 515, 1, 0, 0, 0, 88, 521, 1, 0, 0, 0, 90, 523, 1, 0, 0,
		0, 92, 525, 1, 0, 0, 0, 94, 95, 3, 2, 1, 0, 95, 96, 5, 2, 0, 0, 96, 97,
		3, 46, 23, 0, 97, 98, 5, 38, 0, 0, 98, 1, 1, 0, 0, 0, 99, 102, 3, 4, 2,
		0, 100, 102, 3, 44, 22, 0, 101, 99, 1, 0, 0, 0, 101, 100, 1, 0, 0, 0, 102,
		3, 1, 0, 0, 0, 103, 104, 5, 3, 0, 0, 104, 117, 3, 6, 3, 0, 105, 106, 5,
		5, 0, 0, 106, 107, 5, 64, 0, 0, 107, 108, 5, 46, 0, 0, 108, 109, 3, 32,
		16, 0, 109, 110, 5, 44, 0, 0, 110, 111, 3, 36, 18, 0, 111, 117, 1, 0, 0,
		0, 112, 113, 5, 28, 0, 0, 113, 114, 5, 64, 0, 0, 114, 115, 5, 46, 0, 0,
		115, 117, 3, 38, 19, 0, 116, 103, 1, 0, 0, 0, 116, 105, 1, 0, 0, 0, 116,
		112, 1, 0, 0, 0, 117, 5, 1, 0, 0, 0, 118, 123, 3, 8, 4, 0, 119, 120, 5,
		51, 0, 0, 120, 122, 3, 8, 4, 0, 121, 119, 1, 0, 0, 0, 122, 125, 1, 0, 0,
		0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125,
		123, 1, 0, 0, 0, 126, 127, 5, 46, 0, 0, 127, 128, 3, 30, 15, 0, 128, 7,
		1, 0, 0, 0, 129, 134, 5, 64, 0, 0, 130, 131, 5, 61, 0, 0, 131, 133, 5,
		64, 0, 0, 132, 130, 1, 0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132, 1, 0, 0,
		0, 134, 135, 1, 0, 0, 0, 135, 143, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 137,
		138, 5, 49, 0, 0, 138, 139, 3, 10, 5, 0, 139, 140, 5, 50, 0, 0, 140, 142,
		1, 0, 0, 0, 141, 137, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0,
		0, 0, 143, 144, 1, 0, 0, 0, 144, 9, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146,
		152, 3, 12, 6, 0, 147, 148, 3, 16, 8, 0, 148, 149, 3, 12, 6, 0, 149, 151,
		1, 0, 0, 0, 150, 147, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0,
		0, 0, 152, 153, 1, 0, 0, 0, 153, 11, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0,
		155, 161, 3, 14, 7, 0, 156, 157, 3, 18, 9, 0, 157, 158, 3, 14, 7, 0, 158,
		160, 1, 0, 0, 0, 159, 156, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159,
		1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 13, 1, 0, 0, 0, 163, 161, 1, 0,
		0, 0, 164, 170, 3, 22, 11, 0, 165, 166, 3, 20, 10, 0, 166, 167, 3, 22,
		11, 0, 167, 169, 1, 0, 0, 0, 168, 165, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0,
		170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 15, 1, 0, 0, 0, 172, 170,
		1, 0, 0, 0, 173, 174, 7, 0, 0, 0, 174, 17, 1, 0, 0, 0, 175, 176, 7, 1,
		0, 0, 176, 19, 1, 0, 0, 0, 177, 178, 5, 54, 0, 0, 178, 21, 1, 0, 0, 0,
		179, 181, 3, 26, 13, 0, 180, 179, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		182, 1, 0, 0, 0, 182, 185, 3, 24, 12, 0, 183, 185, 3, 28, 14, 0, 184, 180,
		1, 0, 0, 0, 184, 183, 1, 0, 0, 0, 185, 23, 1, 0, 0, 0, 186, 188, 5, 59,
		0, 0, 187, 186, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0,
		189, 209, 3, 8, 4, 0, 190, 191, 5, 64, 0, 0, 191, 192, 5, 47, 0, 0, 192,
		197, 3, 82, 41, 0, 193, 194, 5, 51, 0, 0, 194, 196, 3, 82, 41, 0, 195,
		193, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198,
		1, 0, 0, 0, 198, 200, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 201, 5, 48,
		0, 0, 201, 209, 1, 0, 0, 0, 202, 209, 5, 62, 0, 0, 203, 209, 5, 63, 0,
		0, 204, 205, 5, 47, 0, 0, 205, 206, 3, 82, 41, 0, 206, 207, 5, 48, 0, 0,
		207, 209, 1, 0, 0, 0, 208, 187, 1, 0, 0, 0, 208, 190, 1, 0, 0, 0, 208,
		202, 1, 0, 0, 0, 208, 203, 1, 0, 0, 0, 208, 204, 1, 0, 0, 0, 209, 25, 1,
		0, 0, 0, 210, 211, 5, 56, 0, 0, 211, 27, 1, 0, 0, 0, 212, 213, 5, 60, 0,
		0, 213, 216, 3, 8, 4, 0, 214, 216, 5, 65, 0, 0, 215, 212, 1, 0, 0, 0, 215,
		214, 1, 0, 0, 0, 216, 29, 1, 0, 0, 0, 217, 220, 3, 38, 19, 0, 218, 220,
		3, 34, 17, 0, 219, 217, 1, 0, 0, 0, 219, 218, 1, 0, 0, 0, 220, 31, 1, 0,
		0, 0, 221, 222, 7, 2, 0, 0, 222, 33, 1, 0, 0, 0, 223, 225, 5, 59, 0, 0,
		224, 223, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226,
		229, 3, 32, 16, 0, 227, 229, 5, 64, 0, 0, 228, 226, 1, 0, 0, 0, 228, 227,
		1, 0, 0, 0, 229, 35, 1, 0, 0, 0, 230, 231, 7, 3, 0, 0, 231, 37, 1, 0, 0,
		0, 232, 236, 5, 29, 0, 0, 233, 235, 3, 6, 3, 0, 234, 233, 1, 0, 0, 0, 235,
		238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 239,
		1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 240, 5, 30, 0, 0, 240, 39, 1, 0,
		0, 0, 241, 243, 5, 4, 0, 0, 242, 241, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0,
		243, 244, 1, 0, 0, 0, 244, 249, 3, 8, 4, 0, 245, 246, 5, 51, 0, 0, 246,
		248, 3, 8, 4, 0, 247, 245, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249, 247,
		1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 252, 1, 0, 0, 0, 251, 249, 1, 0,
		0, 0, 252, 253, 5, 46, 0, 0, 253, 254, 3, 34, 17, 0, 254, 41, 1, 0, 0,
		0, 255, 260, 3, 40, 20, 0, 256, 257, 5, 51, 0, 0, 257, 259, 3, 40, 20,
		0, 258, 256, 1, 0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260,
		261, 1, 0, 0, 0, 261, 43, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263, 264, 5,
		31, 0, 0, 264, 265, 5, 64, 0, 0, 265, 267, 5, 47, 0, 0, 266, 268, 3, 42,
		21, 0, 267, 266, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0,
		269, 273, 5, 48, 0, 0, 270, 272, 3, 4, 2, 0, 271, 270, 1, 0, 0, 0, 272,
		275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 279,
		1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 278, 3, 48, 24, 0, 277, 276, 1,
		0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0,
		0, 280, 282, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 307, 5, 32, 0, 0, 283,
		284, 5, 33, 0, 0, 284, 285, 5, 64, 0, 0, 285, 287, 5, 47, 0, 0, 286, 288,
		3, 42, 21, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 1,
		0, 0, 0, 289, 290, 5, 48, 0, 0, 290, 291, 5, 46, 0, 0, 291, 295, 3, 34,
		17, 0, 292, 294, 3, 4, 2, 0, 293, 292, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0,
		295, 293, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 301, 1, 0, 0, 0, 297,
		295, 1, 0, 0, 0, 298, 300, 3, 48, 24, 0, 299, 298, 1, 0, 0, 0, 300, 303,
		1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 1, 0,
		0, 0, 303, 301, 1, 0, 0, 0, 304, 305, 5, 34, 0, 0, 305, 307, 1, 0, 0, 0,
		306, 263, 1, 0, 0, 0, 306, 283, 1, 0, 0, 0, 307, 45, 1, 0, 0, 0, 308, 310,
		3, 4, 2, 0, 309, 308, 1, 0, 0, 0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0,
		0, 0, 311, 312, 1, 0, 0, 0, 312, 317, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0,
		314, 316, 3, 48, 24, 0, 315, 314, 1, 0, 0, 0, 316, 319, 1, 0, 0, 0, 317,
		315, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 47, 1, 0, 0, 0, 319, 317, 1,
		0, 0, 0, 320, 331, 3, 50, 25, 0, 321, 331, 3, 52, 26, 0, 322, 331, 3, 54,
		27, 0, 323, 331, 3, 56, 28, 0, 324, 331, 3, 58, 29, 0, 325, 331, 3, 60,
		30, 0, 326, 331, 3, 62, 31, 0, 327, 331, 3, 64, 32, 0, 328, 331, 3, 66,
		33, 0, 329, 331, 3, 68, 34, 0, 330, 320, 1, 0, 0, 0, 330, 321, 1, 0, 0,
		0, 330, 322, 1, 0, 0, 0, 330, 323, 1, 0, 0, 0, 330, 324, 1, 0, 0, 0, 330,
		325, 1, 0, 0, 0, 330, 326, 1, 0, 0, 0, 330, 327, 1, 0, 0, 0, 330, 328,
		1, 0, 0, 0, 330, 329, 1, 0, 0, 0, 331, 49, 1, 0, 0, 0, 332, 333, 5, 36,
		0, 0, 333, 335, 5, 47, 0, 0, 334, 336, 5, 59, 0, 0, 335, 334, 1, 0, 0,
		0, 335, 336, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 345, 3, 8, 4, 0, 338,
		340, 5, 51, 0, 0, 339, 341, 5, 59, 0, 0, 340, 339, 1, 0, 0, 0, 340, 341,
		1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 344, 3, 8, 4, 0, 343, 338, 1, 0,
		0, 0, 344, 347, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0,
		346, 348, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 348, 349, 5, 48, 0, 0, 349,
		51, 1, 0, 0, 0, 350, 351, 5, 37, 0, 0, 351, 352, 5, 47, 0, 0, 352, 357,
		3, 82, 41, 0, 353, 354, 5, 51, 0, 0, 354, 356, 3, 82, 41, 0, 355, 353,
		1, 0, 0, 0, 356, 359, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 357, 358, 1, 0,
		0, 0, 358, 360, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 360, 361, 5, 48, 0, 0,
		361, 53, 1, 0, 0, 0, 362, 363, 5, 15, 0, 0, 363, 364, 3, 82, 41, 0, 364,
		368, 5, 17, 0, 0, 365, 367, 3, 48, 24, 0, 366, 365, 1, 0, 0, 0, 367, 370,
		1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 378, 1, 0,
		0, 0, 370, 368, 1, 0, 0, 0, 371, 375, 5, 18, 0, 0, 372, 374, 3, 48, 24,
		0, 373, 372, 1, 0, 0, 0, 374, 377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375,
		376, 1, 0, 0, 0, 376, 379, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 371,
		1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 381, 5, 16,
		0, 0, 381, 55, 1, 0, 0, 0, 382, 383, 5, 19, 0, 0, 383, 384, 3, 10, 5, 0,
		384, 385, 5, 20, 0, 0, 385, 393, 3, 70, 35, 0, 386, 390, 5, 18, 0, 0, 387,
		389, 3, 48, 24, 0, 388, 387, 1, 0, 0, 0, 389, 392, 1, 0, 0, 0, 390, 388,
		1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392, 390, 1, 0,
		0, 0, 393, 386, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0,
		395, 396, 5, 21, 0, 0, 396, 57, 1, 0, 0, 0, 397, 398, 5, 22, 0, 0, 398,
		399, 5, 64, 0, 0, 399, 400, 5, 58, 0, 0, 400, 401, 3, 10, 5, 0, 401, 402,
		5, 24, 0, 0, 402, 403, 3, 10, 5, 0, 403, 407, 5, 25, 0, 0, 404, 406, 3,
		48, 24, 0, 405, 404, 1, 0, 0, 0, 406, 409, 1, 0, 0, 0, 407, 405, 1, 0,
		0, 0, 407, 408, 1, 0, 0, 0, 408, 410, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0,
		410, 411, 5, 23, 0, 0, 411, 59, 1, 0, 0, 0, 412, 413, 5, 26, 0, 0, 413,
		414, 3, 82, 41, 0, 414, 418, 5, 25, 0, 0, 415, 417, 3, 48, 24, 0, 416,
		415, 1, 0, 0, 0, 417, 420, 1, 0, 0, 0, 418, 416, 1, 0, 0, 0, 418, 419,
		1, 0, 0, 0, 419, 421, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 421, 422, 5, 27,
		0, 0, 422, 61, 1, 0, 0, 0, 423, 427, 5, 25, 0, 0, 424, 426, 3, 48, 24,
		0, 425, 424, 1, 0, 0, 0, 426, 429, 1, 0, 0, 0, 427, 425, 1, 0, 0, 0, 427,
		428, 1, 0, 0, 0, 428, 430, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 430, 431,
		5, 24, 0, 0, 431, 432, 3, 82, 41, 0, 432, 63, 1, 0, 0, 0, 433, 435, 5,
		59, 0, 0, 434, 433, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 436, 1, 0, 0,
		0, 436, 437, 3, 8, 4, 0, 437, 438, 5, 58, 0, 0, 438, 439, 3, 82, 41, 0,
		439, 65, 1, 0, 0, 0, 440, 441, 5, 64, 0, 0, 441, 442, 5, 47, 0, 0, 442,
		447, 3, 82, 41, 0, 443, 444, 5, 51, 0, 0, 444, 446, 3, 82, 41, 0, 445,
		443, 1, 0, 0, 0, 446, 449, 1, 0, 0, 0, 447, 445, 1, 0, 0, 0, 447, 448,
		1, 0, 0, 0, 448, 450, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 450, 451, 5, 48,
		0, 0, 451, 67, 1, 0, 0, 0, 452, 453, 5, 35, 0, 0, 453, 454, 3, 82, 41,
		0, 454, 69, 1, 0, 0, 0, 455, 457, 3, 72, 36, 0, 456, 455, 1, 0, 0, 0, 457,
		460, 1, 0, 0, 0, 458, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 71, 1,
		0, 0, 0, 460, 458, 1, 0, 0, 0, 461, 462, 3, 74, 37, 0, 462, 466, 5, 46,
		0, 0, 463, 465, 3, 48, 24, 0, 464, 463, 1, 0, 0, 0, 465, 468, 1, 0, 0,
		0, 466, 464, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 73, 1, 0, 0, 0, 468,
		466, 1, 0, 0, 0, 469, 474, 3, 76, 38, 0, 470, 471, 5, 51, 0, 0, 471, 473,
		3, 76, 38, 0, 472, 470, 1, 0, 0, 0, 473, 476, 1, 0, 0, 0, 474, 472, 1,
		0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 75, 1, 0, 0, 0, 476, 474, 1, 0, 0,
		0, 477, 479, 3, 26, 13, 0, 478, 477, 1, 0, 0, 0, 478, 479, 1, 0, 0, 0,
		479, 480, 1, 0, 0, 0, 480, 486, 5, 62, 0, 0, 481, 483, 5, 39, 0, 0, 482,
		484, 3, 26, 13, 0, 483, 482, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484, 485,
		1, 0, 0, 0, 485, 487, 5, 62, 0, 0, 486, 481, 1, 0, 0, 0, 486, 487, 1, 0,
		0, 0, 487, 77, 1, 0, 0, 0, 488, 492, 3, 10, 5, 0, 489, 490, 3, 80, 40,
		0, 490, 491, 3, 10, 5, 0, 491, 493, 1, 0, 0, 0, 492, 489, 1, 0, 0, 0, 492,
		493, 1, 0, 0, 0, 493, 79, 1, 0, 0, 0, 494, 495, 7, 4, 0, 0, 495, 81, 1,
		0, 0, 0, 496, 502, 3, 84, 42, 0, 497, 498, 3, 90, 45, 0, 498, 499, 3, 84,
		42, 0, 499, 501, 1, 0, 0, 0, 500, 497, 1, 0, 0, 0, 501, 504, 1, 0, 0, 0,
		502, 500, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503, 83, 1, 0, 0, 0, 504, 502,
		1, 0, 0, 0, 505, 511, 3, 86, 43, 0, 506, 507, 3, 92, 46, 0, 507, 508, 3,
		86, 43, 0, 508, 510, 1, 0, 0, 0, 509, 506, 1, 0, 0, 0, 510, 513, 1, 0,
		0, 0, 511, 509, 1, 0, 0, 0, 511, 512, 1, 0, 0, 0, 512, 85, 1, 0, 0, 0,
		513, 511, 1, 0, 0, 0, 514, 516, 5, 14, 0, 0, 515, 514, 1, 0, 0, 0, 515,
		516, 1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517, 518, 3, 88, 44, 0, 518, 87,
		1, 0, 0, 0, 519, 522, 7, 5, 0, 0, 520, 522, 3, 78, 39, 0, 521, 519, 1,
		0, 0, 0, 521, 520, 1, 0, 0, 0, 522, 89, 1, 0, 0, 0, 523, 524, 5, 13, 0,
		0, 524, 91, 1, 0, 0, 0, 525, 526, 5, 12, 0, 0, 526, 93, 1, 0, 0, 0, 56,
		101, 116, 123, 134, 143, 152, 161, 170, 180, 184, 187, 197, 208, 215, 219,
		224, 228, 236, 242, 249, 260, 267, 273, 279, 287, 295, 301, 306, 311, 317,
		330, 335, 340, 345, 357, 368, 375, 378, 390, 393, 407, 418, 427, 434, 447,
		458, 466, 474, 478, 483, 486, 492, 502, 511, 515, 521,
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

// T2AlgumaLexerParserInit initializes any static state used to implement T2AlgumaLexerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewT2AlgumaLexerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func T2AlgumaLexerParserInit() {
	staticData := &T2AlgumaLexerParserStaticData
	staticData.once.Do(t2algumalexerParserInit)
}

// NewT2AlgumaLexerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewT2AlgumaLexerParser(input antlr.TokenStream) *T2AlgumaLexerParser {
	T2AlgumaLexerParserInit()
	this := new(T2AlgumaLexerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &T2AlgumaLexerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "T2AlgumaLexer.g4"

	return this
}

// T2AlgumaLexerParser tokens.
const (
	T2AlgumaLexerParserEOF                    = antlr.TokenEOF
	T2AlgumaLexerParserCOMENTARIO             = 1
	T2AlgumaLexerParserALGORITMO              = 2
	T2AlgumaLexerParserDECLARE                = 3
	T2AlgumaLexerParserVAR                    = 4
	T2AlgumaLexerParserCONSTANTE              = 5
	T2AlgumaLexerParserLITERAL                = 6
	T2AlgumaLexerParserINTEIRO                = 7
	T2AlgumaLexerParserREAL                   = 8
	T2AlgumaLexerParserLOGICO                 = 9
	T2AlgumaLexerParserVERDADEIRO             = 10
	T2AlgumaLexerParserFALSO                  = 11
	T2AlgumaLexerParserE                      = 12
	T2AlgumaLexerParserOU                     = 13
	T2AlgumaLexerParserNAO                    = 14
	T2AlgumaLexerParserSE                     = 15
	T2AlgumaLexerParserFIM_SE                 = 16
	T2AlgumaLexerParserENTAO                  = 17
	T2AlgumaLexerParserSENAO                  = 18
	T2AlgumaLexerParserCASO                   = 19
	T2AlgumaLexerParserSEJA                   = 20
	T2AlgumaLexerParserFIM_CASO               = 21
	T2AlgumaLexerParserPARA                   = 22
	T2AlgumaLexerParserFIM_PARA               = 23
	T2AlgumaLexerParserATE                    = 24
	T2AlgumaLexerParserFACA                   = 25
	T2AlgumaLexerParserENQUANTO               = 26
	T2AlgumaLexerParserFIM_ENQUANTO           = 27
	T2AlgumaLexerParserTIPO                   = 28
	T2AlgumaLexerParserREGISTRO               = 29
	T2AlgumaLexerParserFIM_REGISTRO           = 30
	T2AlgumaLexerParserPROCEDIMENTO           = 31
	T2AlgumaLexerParserFIM_PROCEDIMENTO       = 32
	T2AlgumaLexerParserFUNCAO                 = 33
	T2AlgumaLexerParserFIM_FUNCAO             = 34
	T2AlgumaLexerParserRETORNE                = 35
	T2AlgumaLexerParserLEIA                   = 36
	T2AlgumaLexerParserESCREVA                = 37
	T2AlgumaLexerParserFIM_ALGORITMO          = 38
	T2AlgumaLexerParserINTERVALO              = 39
	T2AlgumaLexerParserMENOR                  = 40
	T2AlgumaLexerParserMENORIGUAL             = 41
	T2AlgumaLexerParserMAIOR                  = 42
	T2AlgumaLexerParserMAIORIGUAL             = 43
	T2AlgumaLexerParserIGUAL                  = 44
	T2AlgumaLexerParserDIFERENTE              = 45
	T2AlgumaLexerParserDOISPONTOS             = 46
	T2AlgumaLexerParserABREPAR                = 47
	T2AlgumaLexerParserFECHAPAR               = 48
	T2AlgumaLexerParserABRECHAVE              = 49
	T2AlgumaLexerParserFECHACHAVE             = 50
	T2AlgumaLexerParserVIRGULA                = 51
	T2AlgumaLexerParserASPAS                  = 52
	T2AlgumaLexerParserDIVISAO                = 53
	T2AlgumaLexerParserRESTO                  = 54
	T2AlgumaLexerParserSOMA                   = 55
	T2AlgumaLexerParserSUBTRACAO              = 56
	T2AlgumaLexerParserMULTIPLICACAO          = 57
	T2AlgumaLexerParserATRIBUICAO             = 58
	T2AlgumaLexerParserPONTEIRO               = 59
	T2AlgumaLexerParserENDERECO               = 60
	T2AlgumaLexerParserPONTO                  = 61
	T2AlgumaLexerParserNUM_INT                = 62
	T2AlgumaLexerParserNUM_REAL               = 63
	T2AlgumaLexerParserIDENT                  = 64
	T2AlgumaLexerParserCADEIA                 = 65
	T2AlgumaLexerParserERRO_COMENTARIO_ABERTO = 66
	T2AlgumaLexerParserERRO_CADEIA_ABERTA     = 67
	T2AlgumaLexerParserWS                     = 68
	T2AlgumaLexerParserERRO                   = 69
)

// T2AlgumaLexerParser rules.
const (
	T2AlgumaLexerParserRULE_programa              = 0
	T2AlgumaLexerParserRULE_declaracoes           = 1
	T2AlgumaLexerParserRULE_declaracoes_variaveis = 2
	T2AlgumaLexerParserRULE_variavel              = 3
	T2AlgumaLexerParserRULE_identificador         = 4
	T2AlgumaLexerParserRULE_exp_aritmetica        = 5
	T2AlgumaLexerParserRULE_termo                 = 6
	T2AlgumaLexerParserRULE_fator                 = 7
	T2AlgumaLexerParserRULE_op1                   = 8
	T2AlgumaLexerParserRULE_op2                   = 9
	T2AlgumaLexerParserRULE_op3                   = 10
	T2AlgumaLexerParserRULE_parcela               = 11
	T2AlgumaLexerParserRULE_parcela_unario        = 12
	T2AlgumaLexerParserRULE_op_unario             = 13
	T2AlgumaLexerParserRULE_parcela_nao_unario    = 14
	T2AlgumaLexerParserRULE_tipo                  = 15
	T2AlgumaLexerParserRULE_tipo_basico           = 16
	T2AlgumaLexerParserRULE_tipo_variavel         = 17
	T2AlgumaLexerParserRULE_valor_constante       = 18
	T2AlgumaLexerParserRULE_registro              = 19
	T2AlgumaLexerParserRULE_parametro             = 20
	T2AlgumaLexerParserRULE_parametros            = 21
	T2AlgumaLexerParserRULE_declaracoes_funcoes   = 22
	T2AlgumaLexerParserRULE_corpo                 = 23
	T2AlgumaLexerParserRULE_cmd                   = 24
	T2AlgumaLexerParserRULE_cmdLeia               = 25
	T2AlgumaLexerParserRULE_cmdEscreva            = 26
	T2AlgumaLexerParserRULE_cmdSe                 = 27
	T2AlgumaLexerParserRULE_cmdCaso               = 28
	T2AlgumaLexerParserRULE_cmdPara               = 29
	T2AlgumaLexerParserRULE_cmdEnquanto           = 30
	T2AlgumaLexerParserRULE_cmdFaca               = 31
	T2AlgumaLexerParserRULE_cmdAtribuicao         = 32
	T2AlgumaLexerParserRULE_cmdChamada            = 33
	T2AlgumaLexerParserRULE_cmdRetorne            = 34
	T2AlgumaLexerParserRULE_selecao               = 35
	T2AlgumaLexerParserRULE_item_selecao          = 36
	T2AlgumaLexerParserRULE_constantes            = 37
	T2AlgumaLexerParserRULE_numero_intervalo      = 38
	T2AlgumaLexerParserRULE_exp_relacional        = 39
	T2AlgumaLexerParserRULE_op_relacional         = 40
	T2AlgumaLexerParserRULE_expressao             = 41
	T2AlgumaLexerParserRULE_termo_logico          = 42
	T2AlgumaLexerParserRULE_fator_logico          = 43
	T2AlgumaLexerParserRULE_parcela_logica        = 44
	T2AlgumaLexerParserRULE_op_logico_1           = 45
	T2AlgumaLexerParserRULE_op_logico_2           = 46
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
	p.RuleIndex = T2AlgumaLexerParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_programa

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
	return s.GetToken(T2AlgumaLexerParserALGORITMO, 0)
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
	return s.GetToken(T2AlgumaLexerParserFIM_ALGORITMO, 0)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (p *T2AlgumaLexerParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, T2AlgumaLexerParserRULE_programa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Declaracoes()
	}
	{
		p.SetState(95)
		p.Match(T2AlgumaLexerParserALGORITMO)
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
		p.Match(T2AlgumaLexerParserFIM_ALGORITMO)
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
	Declaracoes_variaveis() IDeclaracoes_variaveisContext
	Declaracoes_funcoes() IDeclaracoes_funcoesContext

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
	p.RuleIndex = T2AlgumaLexerParserRULE_declaracoes
	return p
}

func InitEmptyDeclaracoesContext(p *DeclaracoesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_declaracoes
}

func (*DeclaracoesContext) IsDeclaracoesContext() {}

func NewDeclaracoesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracoesContext {
	var p = new(DeclaracoesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_declaracoes

	return p
}

func (s *DeclaracoesContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracoesContext) Declaracoes_variaveis() IDeclaracoes_variaveisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracoes_variaveisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracoes_variaveisContext)
}

func (s *DeclaracoesContext) Declaracoes_funcoes() IDeclaracoes_funcoesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracoes_funcoesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterDeclaracoes(s)
	}
}

func (s *DeclaracoesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitDeclaracoes(s)
	}
}

func (p *T2AlgumaLexerParser) Declaracoes() (localctx IDeclaracoesContext) {
	localctx = NewDeclaracoesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, T2AlgumaLexerParserRULE_declaracoes)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T2AlgumaLexerParserDECLARE, T2AlgumaLexerParserCONSTANTE, T2AlgumaLexerParserTIPO:
		{
			p.SetState(99)
			p.Declaracoes_variaveis()
		}

	case T2AlgumaLexerParserPROCEDIMENTO, T2AlgumaLexerParserFUNCAO:
		{
			p.SetState(100)
			p.Declaracoes_funcoes()
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
	p.RuleIndex = T2AlgumaLexerParserRULE_declaracoes_variaveis
	return p
}

func InitEmptyDeclaracoes_variaveisContext(p *Declaracoes_variaveisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_declaracoes_variaveis
}

func (*Declaracoes_variaveisContext) IsDeclaracoes_variaveisContext() {}

func NewDeclaracoes_variaveisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracoes_variaveisContext {
	var p = new(Declaracoes_variaveisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_declaracoes_variaveis

	return p
}

func (s *Declaracoes_variaveisContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracoes_variaveisContext) DECLARE() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserDECLARE, 0)
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
	return s.GetToken(T2AlgumaLexerParserCONSTANTE, 0)
}

func (s *Declaracoes_variaveisContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserIDENT, 0)
}

func (s *Declaracoes_variaveisContext) DOISPONTOS() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserDOISPONTOS, 0)
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
	return s.GetToken(T2AlgumaLexerParserIGUAL, 0)
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
	return s.GetToken(T2AlgumaLexerParserTIPO, 0)
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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterDeclaracoes_variaveis(s)
	}
}

func (s *Declaracoes_variaveisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitDeclaracoes_variaveis(s)
	}
}

func (p *T2AlgumaLexerParser) Declaracoes_variaveis() (localctx IDeclaracoes_variaveisContext) {
	localctx = NewDeclaracoes_variaveisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, T2AlgumaLexerParserRULE_declaracoes_variaveis)
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T2AlgumaLexerParserDECLARE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Match(T2AlgumaLexerParserDECLARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Variavel()
		}

	case T2AlgumaLexerParserCONSTANTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Match(T2AlgumaLexerParserCONSTANTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(T2AlgumaLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(T2AlgumaLexerParserDOISPONTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Tipo_basico()
		}
		{
			p.SetState(109)
			p.Match(T2AlgumaLexerParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Valor_constante()
		}

	case T2AlgumaLexerParserTIPO:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.Match(T2AlgumaLexerParserTIPO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(T2AlgumaLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(T2AlgumaLexerParserDOISPONTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_variavel
	return p
}

func InitEmptyVariavelContext(p *VariavelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_variavel
}

func (*VariavelContext) IsVariavelContext() {}

func NewVariavelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariavelContext {
	var p = new(VariavelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_variavel

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
	return s.GetToken(T2AlgumaLexerParserDOISPONTOS, 0)
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
	return s.GetTokens(T2AlgumaLexerParserVIRGULA)
}

func (s *VariavelContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserVIRGULA, i)
}

func (s *VariavelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariavelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariavelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterVariavel(s)
	}
}

func (s *VariavelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitVariavel(s)
	}
}

func (p *T2AlgumaLexerParser) Variavel() (localctx IVariavelContext) {
	localctx = NewVariavelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, T2AlgumaLexerParserRULE_variavel)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Identificador()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserVIRGULA {
		{
			p.SetState(119)
			p.Match(T2AlgumaLexerParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Identificador()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
		p.Match(T2AlgumaLexerParserDOISPONTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_identificador
	return p
}

func InitEmptyIdentificadorContext(p *IdentificadorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_identificador
}

func (*IdentificadorContext) IsIdentificadorContext() {}

func NewIdentificadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentificadorContext {
	var p = new(IdentificadorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_identificador

	return p
}

func (s *IdentificadorContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentificadorContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(T2AlgumaLexerParserIDENT)
}

func (s *IdentificadorContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserIDENT, i)
}

func (s *IdentificadorContext) AllPONTO() []antlr.TerminalNode {
	return s.GetTokens(T2AlgumaLexerParserPONTO)
}

func (s *IdentificadorContext) PONTO(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserPONTO, i)
}

func (s *IdentificadorContext) AllABRECHAVE() []antlr.TerminalNode {
	return s.GetTokens(T2AlgumaLexerParserABRECHAVE)
}

func (s *IdentificadorContext) ABRECHAVE(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserABRECHAVE, i)
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
	return s.GetTokens(T2AlgumaLexerParserFECHACHAVE)
}

func (s *IdentificadorContext) FECHACHAVE(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFECHACHAVE, i)
}

func (s *IdentificadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentificadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentificadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterIdentificador(s)
	}
}

func (s *IdentificadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitIdentificador(s)
	}
}

func (p *T2AlgumaLexerParser) Identificador() (localctx IIdentificadorContext) {
	localctx = NewIdentificadorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, T2AlgumaLexerParserRULE_identificador)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(T2AlgumaLexerParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserPONTO {
		{
			p.SetState(130)
			p.Match(T2AlgumaLexerParserPONTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Match(T2AlgumaLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserABRECHAVE {
		{
			p.SetState(137)
			p.Match(T2AlgumaLexerParserABRECHAVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Exp_aritmetica()
		}
		{
			p.SetState(139)
			p.Match(T2AlgumaLexerParserFECHACHAVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(145)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_exp_aritmetica
	return p
}

func InitEmptyExp_aritmeticaContext(p *Exp_aritmeticaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_exp_aritmetica
}

func (*Exp_aritmeticaContext) IsExp_aritmeticaContext() {}

func NewExp_aritmeticaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_aritmeticaContext {
	var p = new(Exp_aritmeticaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_exp_aritmetica

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterExp_aritmetica(s)
	}
}

func (s *Exp_aritmeticaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitExp_aritmetica(s)
	}
}

func (p *T2AlgumaLexerParser) Exp_aritmetica() (localctx IExp_aritmeticaContext) {
	localctx = NewExp_aritmeticaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, T2AlgumaLexerParserRULE_exp_aritmetica)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Termo()
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(147)
				p.Op1()
			}
			{
				p.SetState(148)
				p.Termo()
			}

		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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
	p.RuleIndex = T2AlgumaLexerParserRULE_termo
	return p
}

func InitEmptyTermoContext(p *TermoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_termo
}

func (*TermoContext) IsTermoContext() {}

func NewTermoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermoContext {
	var p = new(TermoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_termo

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterTermo(s)
	}
}

func (s *TermoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitTermo(s)
	}
}

func (p *T2AlgumaLexerParser) Termo() (localctx ITermoContext) {
	localctx = NewTermoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, T2AlgumaLexerParserRULE_termo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Fator()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserDIVISAO || _la == T2AlgumaLexerParserMULTIPLICACAO {
		{
			p.SetState(156)
			p.Op2()
		}
		{
			p.SetState(157)
			p.Fator()
		}

		p.SetState(163)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_fator
	return p
}

func InitEmptyFatorContext(p *FatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_fator
}

func (*FatorContext) IsFatorContext() {}

func NewFatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FatorContext {
	var p = new(FatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_fator

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterFator(s)
	}
}

func (s *FatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitFator(s)
	}
}

func (p *T2AlgumaLexerParser) Fator() (localctx IFatorContext) {
	localctx = NewFatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, T2AlgumaLexerParserRULE_fator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Parcela()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserRESTO {
		{
			p.SetState(165)
			p.Op3()
		}
		{
			p.SetState(166)
			p.Parcela()
		}

		p.SetState(172)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_op1
	return p
}

func InitEmptyOp1Context(p *Op1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_op1
}

func (*Op1Context) IsOp1Context() {}

func NewOp1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op1Context {
	var p = new(Op1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_op1

	return p
}

func (s *Op1Context) GetParser() antlr.Parser { return s.parser }

func (s *Op1Context) SOMA() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserSOMA, 0)
}

func (s *Op1Context) SUBTRACAO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserSUBTRACAO, 0)
}

func (s *Op1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterOp1(s)
	}
}

func (s *Op1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitOp1(s)
	}
}

func (p *T2AlgumaLexerParser) Op1() (localctx IOp1Context) {
	localctx = NewOp1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, T2AlgumaLexerParserRULE_op1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		_la = p.GetTokenStream().LA(1)

		if !(_la == T2AlgumaLexerParserSOMA || _la == T2AlgumaLexerParserSUBTRACAO) {
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
	p.RuleIndex = T2AlgumaLexerParserRULE_op2
	return p
}

func InitEmptyOp2Context(p *Op2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_op2
}

func (*Op2Context) IsOp2Context() {}

func NewOp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op2Context {
	var p = new(Op2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_op2

	return p
}

func (s *Op2Context) GetParser() antlr.Parser { return s.parser }

func (s *Op2Context) MULTIPLICACAO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserMULTIPLICACAO, 0)
}

func (s *Op2Context) DIVISAO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserDIVISAO, 0)
}

func (s *Op2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterOp2(s)
	}
}

func (s *Op2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitOp2(s)
	}
}

func (p *T2AlgumaLexerParser) Op2() (localctx IOp2Context) {
	localctx = NewOp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, T2AlgumaLexerParserRULE_op2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		_la = p.GetTokenStream().LA(1)

		if !(_la == T2AlgumaLexerParserDIVISAO || _la == T2AlgumaLexerParserMULTIPLICACAO) {
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
	p.RuleIndex = T2AlgumaLexerParserRULE_op3
	return p
}

func InitEmptyOp3Context(p *Op3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_op3
}

func (*Op3Context) IsOp3Context() {}

func NewOp3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op3Context {
	var p = new(Op3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_op3

	return p
}

func (s *Op3Context) GetParser() antlr.Parser { return s.parser }

func (s *Op3Context) RESTO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserRESTO, 0)
}

func (s *Op3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterOp3(s)
	}
}

func (s *Op3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitOp3(s)
	}
}

func (p *T2AlgumaLexerParser) Op3() (localctx IOp3Context) {
	localctx = NewOp3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, T2AlgumaLexerParserRULE_op3)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(T2AlgumaLexerParserRESTO)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela
	return p
}

func InitEmptyParcelaContext(p *ParcelaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela
}

func (*ParcelaContext) IsParcelaContext() {}

func NewParcelaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParcelaContext {
	var p = new(ParcelaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterParcela(s)
	}
}

func (s *ParcelaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitParcela(s)
	}
}

func (p *T2AlgumaLexerParser) Parcela() (localctx IParcelaContext) {
	localctx = NewParcelaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, T2AlgumaLexerParserRULE_parcela)
	var _la int

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T2AlgumaLexerParserABREPAR, T2AlgumaLexerParserSUBTRACAO, T2AlgumaLexerParserPONTEIRO, T2AlgumaLexerParserNUM_INT, T2AlgumaLexerParserNUM_REAL, T2AlgumaLexerParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T2AlgumaLexerParserSUBTRACAO {
			{
				p.SetState(179)
				p.Op_unario()
			}

		}
		{
			p.SetState(182)
			p.Parcela_unario()
		}

	case T2AlgumaLexerParserENDERECO, T2AlgumaLexerParserCADEIA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela_unario
	return p
}

func InitEmptyParcela_unarioContext(p *Parcela_unarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela_unario
}

func (*Parcela_unarioContext) IsParcela_unarioContext() {}

func NewParcela_unarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parcela_unarioContext {
	var p = new(Parcela_unarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela_unario

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
	return s.GetToken(T2AlgumaLexerParserPONTEIRO, 0)
}

func (s *Parcela_unarioContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserIDENT, 0)
}

func (s *Parcela_unarioContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserABREPAR, 0)
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
	return s.GetToken(T2AlgumaLexerParserFECHAPAR, 0)
}

func (s *Parcela_unarioContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T2AlgumaLexerParserVIRGULA)
}

func (s *Parcela_unarioContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserVIRGULA, i)
}

func (s *Parcela_unarioContext) NUM_INT() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserNUM_INT, 0)
}

func (s *Parcela_unarioContext) NUM_REAL() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserNUM_REAL, 0)
}

func (s *Parcela_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parcela_unarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parcela_unarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterParcela_unario(s)
	}
}

func (s *Parcela_unarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitParcela_unario(s)
	}
}

func (p *T2AlgumaLexerParser) Parcela_unario() (localctx IParcela_unarioContext) {
	localctx = NewParcela_unarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, T2AlgumaLexerParserRULE_parcela_unario)
	var _la int

	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T2AlgumaLexerParserPONTEIRO {
			{
				p.SetState(186)
				p.Match(T2AlgumaLexerParserPONTEIRO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(189)
			p.Identificador()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(T2AlgumaLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Match(T2AlgumaLexerParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Expressao()
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == T2AlgumaLexerParserVIRGULA {
			{
				p.SetState(193)
				p.Match(T2AlgumaLexerParserVIRGULA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(194)
				p.Expressao()
			}

			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(200)
			p.Match(T2AlgumaLexerParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(202)
			p.Match(T2AlgumaLexerParserNUM_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(203)
			p.Match(T2AlgumaLexerParserNUM_REAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(204)
			p.Match(T2AlgumaLexerParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.Expressao()
		}
		{
			p.SetState(206)
			p.Match(T2AlgumaLexerParserFECHAPAR)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_op_unario
	return p
}

func InitEmptyOp_unarioContext(p *Op_unarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_op_unario
}

func (*Op_unarioContext) IsOp_unarioContext() {}

func NewOp_unarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_unarioContext {
	var p = new(Op_unarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_op_unario

	return p
}

func (s *Op_unarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_unarioContext) SUBTRACAO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserSUBTRACAO, 0)
}

func (s *Op_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_unarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_unarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterOp_unario(s)
	}
}

func (s *Op_unarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitOp_unario(s)
	}
}

func (p *T2AlgumaLexerParser) Op_unario() (localctx IOp_unarioContext) {
	localctx = NewOp_unarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, T2AlgumaLexerParserRULE_op_unario)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(T2AlgumaLexerParserSUBTRACAO)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela_nao_unario
	return p
}

func InitEmptyParcela_nao_unarioContext(p *Parcela_nao_unarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela_nao_unario
}

func (*Parcela_nao_unarioContext) IsParcela_nao_unarioContext() {}

func NewParcela_nao_unarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parcela_nao_unarioContext {
	var p = new(Parcela_nao_unarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela_nao_unario

	return p
}

func (s *Parcela_nao_unarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Parcela_nao_unarioContext) ENDERECO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserENDERECO, 0)
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
	return s.GetToken(T2AlgumaLexerParserCADEIA, 0)
}

func (s *Parcela_nao_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parcela_nao_unarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parcela_nao_unarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterParcela_nao_unario(s)
	}
}

func (s *Parcela_nao_unarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitParcela_nao_unario(s)
	}
}

func (p *T2AlgumaLexerParser) Parcela_nao_unario() (localctx IParcela_nao_unarioContext) {
	localctx = NewParcela_nao_unarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, T2AlgumaLexerParserRULE_parcela_nao_unario)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T2AlgumaLexerParserENDERECO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Match(T2AlgumaLexerParserENDERECO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Identificador()
		}

	case T2AlgumaLexerParserCADEIA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Match(T2AlgumaLexerParserCADEIA)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_tipo

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *T2AlgumaLexerParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, T2AlgumaLexerParserRULE_tipo)
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T2AlgumaLexerParserREGISTRO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Registro()
		}

	case T2AlgumaLexerParserLITERAL, T2AlgumaLexerParserINTEIRO, T2AlgumaLexerParserREAL, T2AlgumaLexerParserLOGICO, T2AlgumaLexerParserPONTEIRO, T2AlgumaLexerParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_tipo_basico
	return p
}

func InitEmptyTipo_basicoContext(p *Tipo_basicoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_tipo_basico
}

func (*Tipo_basicoContext) IsTipo_basicoContext() {}

func NewTipo_basicoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_basicoContext {
	var p = new(Tipo_basicoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_tipo_basico

	return p
}

func (s *Tipo_basicoContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_basicoContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserLITERAL, 0)
}

func (s *Tipo_basicoContext) INTEIRO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserINTEIRO, 0)
}

func (s *Tipo_basicoContext) REAL() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserREAL, 0)
}

func (s *Tipo_basicoContext) LOGICO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserLOGICO, 0)
}

func (s *Tipo_basicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_basicoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_basicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterTipo_basico(s)
	}
}

func (s *Tipo_basicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitTipo_basico(s)
	}
}

func (p *T2AlgumaLexerParser) Tipo_basico() (localctx ITipo_basicoContext) {
	localctx = NewTipo_basicoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, T2AlgumaLexerParserRULE_tipo_basico)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_tipo_variavel
	return p
}

func InitEmptyTipo_variavelContext(p *Tipo_variavelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_tipo_variavel
}

func (*Tipo_variavelContext) IsTipo_variavelContext() {}

func NewTipo_variavelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_variavelContext {
	var p = new(Tipo_variavelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_tipo_variavel

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
	return s.GetToken(T2AlgumaLexerParserIDENT, 0)
}

func (s *Tipo_variavelContext) PONTEIRO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserPONTEIRO, 0)
}

func (s *Tipo_variavelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_variavelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_variavelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterTipo_variavel(s)
	}
}

func (s *Tipo_variavelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitTipo_variavel(s)
	}
}

func (p *T2AlgumaLexerParser) Tipo_variavel() (localctx ITipo_variavelContext) {
	localctx = NewTipo_variavelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, T2AlgumaLexerParserRULE_tipo_variavel)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T2AlgumaLexerParserPONTEIRO {
		{
			p.SetState(223)
			p.Match(T2AlgumaLexerParserPONTEIRO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T2AlgumaLexerParserLITERAL, T2AlgumaLexerParserINTEIRO, T2AlgumaLexerParserREAL, T2AlgumaLexerParserLOGICO:
		{
			p.SetState(226)
			p.Tipo_basico()
		}

	case T2AlgumaLexerParserIDENT:
		{
			p.SetState(227)
			p.Match(T2AlgumaLexerParserIDENT)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_valor_constante
	return p
}

func InitEmptyValor_constanteContext(p *Valor_constanteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_valor_constante
}

func (*Valor_constanteContext) IsValor_constanteContext() {}

func NewValor_constanteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Valor_constanteContext {
	var p = new(Valor_constanteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_valor_constante

	return p
}

func (s *Valor_constanteContext) GetParser() antlr.Parser { return s.parser }

func (s *Valor_constanteContext) CADEIA() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserCADEIA, 0)
}

func (s *Valor_constanteContext) NUM_INT() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserNUM_INT, 0)
}

func (s *Valor_constanteContext) NUM_REAL() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserNUM_REAL, 0)
}

func (s *Valor_constanteContext) VERDADEIRO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserVERDADEIRO, 0)
}

func (s *Valor_constanteContext) FALSO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFALSO, 0)
}

func (s *Valor_constanteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Valor_constanteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Valor_constanteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterValor_constante(s)
	}
}

func (s *Valor_constanteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitValor_constante(s)
	}
}

func (p *T2AlgumaLexerParser) Valor_constante() (localctx IValor_constanteContext) {
	localctx = NewValor_constanteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, T2AlgumaLexerParserRULE_valor_constante)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_registro
	return p
}

func InitEmptyRegistroContext(p *RegistroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_registro
}

func (*RegistroContext) IsRegistroContext() {}

func NewRegistroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegistroContext {
	var p = new(RegistroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_registro

	return p
}

func (s *RegistroContext) GetParser() antlr.Parser { return s.parser }

func (s *RegistroContext) REGISTRO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserREGISTRO, 0)
}

func (s *RegistroContext) FIM_REGISTRO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFIM_REGISTRO, 0)
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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterRegistro(s)
	}
}

func (s *RegistroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitRegistro(s)
	}
}

func (p *T2AlgumaLexerParser) Registro() (localctx IRegistroContext) {
	localctx = NewRegistroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, T2AlgumaLexerParserRULE_registro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(T2AlgumaLexerParserREGISTRO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserIDENT {
		{
			p.SetState(233)
			p.Variavel()
		}

		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(239)
		p.Match(T2AlgumaLexerParserFIM_REGISTRO)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_parametro
	return p
}

func InitEmptyParametroContext(p *ParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_parametro
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_parametro

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
	return s.GetToken(T2AlgumaLexerParserDOISPONTOS, 0)
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
	return s.GetToken(T2AlgumaLexerParserVAR, 0)
}

func (s *ParametroContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T2AlgumaLexerParserVIRGULA)
}

func (s *ParametroContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserVIRGULA, i)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterParametro(s)
	}
}

func (s *ParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitParametro(s)
	}
}

func (p *T2AlgumaLexerParser) Parametro() (localctx IParametroContext) {
	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, T2AlgumaLexerParserRULE_parametro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T2AlgumaLexerParserVAR {
		{
			p.SetState(241)
			p.Match(T2AlgumaLexerParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(244)
		p.Identificador()
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserVIRGULA {
		{
			p.SetState(245)
			p.Match(T2AlgumaLexerParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Identificador()
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(252)
		p.Match(T2AlgumaLexerParserDOISPONTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_parametros

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
	return s.GetTokens(T2AlgumaLexerParserVIRGULA)
}

func (s *ParametrosContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserVIRGULA, i)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterParametros(s)
	}
}

func (s *ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitParametros(s)
	}
}

func (p *T2AlgumaLexerParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, T2AlgumaLexerParserRULE_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Parametro()
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserVIRGULA {
		{
			p.SetState(256)
			p.Match(T2AlgumaLexerParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.Parametro()
		}

		p.SetState(262)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_declaracoes_funcoes
	return p
}

func InitEmptyDeclaracoes_funcoesContext(p *Declaracoes_funcoesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_declaracoes_funcoes
}

func (*Declaracoes_funcoesContext) IsDeclaracoes_funcoesContext() {}

func NewDeclaracoes_funcoesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracoes_funcoesContext {
	var p = new(Declaracoes_funcoesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_declaracoes_funcoes

	return p
}

func (s *Declaracoes_funcoesContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracoes_funcoesContext) PROCEDIMENTO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserPROCEDIMENTO, 0)
}

func (s *Declaracoes_funcoesContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserIDENT, 0)
}

func (s *Declaracoes_funcoesContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserABREPAR, 0)
}

func (s *Declaracoes_funcoesContext) FECHAPAR() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFECHAPAR, 0)
}

func (s *Declaracoes_funcoesContext) FIM_PROCEDIMENTO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFIM_PROCEDIMENTO, 0)
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
	return s.GetToken(T2AlgumaLexerParserFUNCAO, 0)
}

func (s *Declaracoes_funcoesContext) DOISPONTOS() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserDOISPONTOS, 0)
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
	return s.GetToken(T2AlgumaLexerParserFIM_FUNCAO, 0)
}

func (s *Declaracoes_funcoesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracoes_funcoesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracoes_funcoesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterDeclaracoes_funcoes(s)
	}
}

func (s *Declaracoes_funcoesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitDeclaracoes_funcoes(s)
	}
}

func (p *T2AlgumaLexerParser) Declaracoes_funcoes() (localctx IDeclaracoes_funcoesContext) {
	localctx = NewDeclaracoes_funcoesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, T2AlgumaLexerParserRULE_declaracoes_funcoes)
	var _la int

	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T2AlgumaLexerParserPROCEDIMENTO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Match(T2AlgumaLexerParserPROCEDIMENTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Match(T2AlgumaLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(T2AlgumaLexerParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T2AlgumaLexerParserVAR || _la == T2AlgumaLexerParserIDENT {
			{
				p.SetState(266)
				p.Parametros()
			}

		}
		{
			p.SetState(269)
			p.Match(T2AlgumaLexerParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268435496) != 0 {
			{
				p.SetState(270)
				p.Declaracoes_variaveis()
			}

			p.SetState(275)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
			{
				p.SetState(276)
				p.Cmd()
			}

			p.SetState(281)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(282)
			p.Match(T2AlgumaLexerParserFIM_PROCEDIMENTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T2AlgumaLexerParserFUNCAO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.Match(T2AlgumaLexerParserFUNCAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(T2AlgumaLexerParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Match(T2AlgumaLexerParserABREPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T2AlgumaLexerParserVAR || _la == T2AlgumaLexerParserIDENT {
			{
				p.SetState(286)
				p.Parametros()
			}

		}
		{
			p.SetState(289)
			p.Match(T2AlgumaLexerParserFECHAPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(T2AlgumaLexerParserDOISPONTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.Tipo_variavel()
		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268435496) != 0 {
			{
				p.SetState(292)
				p.Declaracoes_variaveis()
			}

			p.SetState(297)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
			{
				p.SetState(298)
				p.Cmd()
			}

			p.SetState(303)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(304)
			p.Match(T2AlgumaLexerParserFIM_FUNCAO)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_corpo
	return p
}

func InitEmptyCorpoContext(p *CorpoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_corpo
}

func (*CorpoContext) IsCorpoContext() {}

func NewCorpoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CorpoContext {
	var p = new(CorpoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_corpo

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCorpo(s)
	}
}

func (s *CorpoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCorpo(s)
	}
}

func (p *T2AlgumaLexerParser) Corpo() (localctx ICorpoContext) {
	localctx = NewCorpoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, T2AlgumaLexerParserRULE_corpo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268435496) != 0 {
		{
			p.SetState(308)
			p.Declaracoes_variaveis()
		}

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(314)
			p.Cmd()
		}

		p.SetState(319)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_cmd
	return p
}

func InitEmptyCmdContext(p *CmdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmd
}

func (*CmdContext) IsCmdContext() {}

func NewCmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdContext {
	var p = new(CmdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmd

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmd(s)
	}
}

func (s *CmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmd(s)
	}
}

func (p *T2AlgumaLexerParser) Cmd() (localctx ICmdContext) {
	localctx = NewCmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, T2AlgumaLexerParserRULE_cmd)
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)
			p.CmdLeia()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(321)
			p.CmdEscreva()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(322)
			p.CmdSe()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(323)
			p.CmdCaso()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(324)
			p.CmdPara()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(325)
			p.CmdEnquanto()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(326)
			p.CmdFaca()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(327)
			p.CmdAtribuicao()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(328)
			p.CmdChamada()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(329)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdLeia
	return p
}

func InitEmptyCmdLeiaContext(p *CmdLeiaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdLeia
}

func (*CmdLeiaContext) IsCmdLeiaContext() {}

func NewCmdLeiaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdLeiaContext {
	var p = new(CmdLeiaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdLeia

	return p
}

func (s *CmdLeiaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdLeiaContext) LEIA() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserLEIA, 0)
}

func (s *CmdLeiaContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserABREPAR, 0)
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
	return s.GetToken(T2AlgumaLexerParserFECHAPAR, 0)
}

func (s *CmdLeiaContext) AllPONTEIRO() []antlr.TerminalNode {
	return s.GetTokens(T2AlgumaLexerParserPONTEIRO)
}

func (s *CmdLeiaContext) PONTEIRO(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserPONTEIRO, i)
}

func (s *CmdLeiaContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T2AlgumaLexerParserVIRGULA)
}

func (s *CmdLeiaContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserVIRGULA, i)
}

func (s *CmdLeiaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdLeiaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdLeiaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmdLeia(s)
	}
}

func (s *CmdLeiaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmdLeia(s)
	}
}

func (p *T2AlgumaLexerParser) CmdLeia() (localctx ICmdLeiaContext) {
	localctx = NewCmdLeiaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, T2AlgumaLexerParserRULE_cmdLeia)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(T2AlgumaLexerParserLEIA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(333)
		p.Match(T2AlgumaLexerParserABREPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T2AlgumaLexerParserPONTEIRO {
		{
			p.SetState(334)
			p.Match(T2AlgumaLexerParserPONTEIRO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(337)
		p.Identificador()
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserVIRGULA {
		{
			p.SetState(338)
			p.Match(T2AlgumaLexerParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T2AlgumaLexerParserPONTEIRO {
			{
				p.SetState(339)
				p.Match(T2AlgumaLexerParserPONTEIRO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(342)
			p.Identificador()
		}

		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(348)
		p.Match(T2AlgumaLexerParserFECHAPAR)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdEscreva
	return p
}

func InitEmptyCmdEscrevaContext(p *CmdEscrevaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdEscreva
}

func (*CmdEscrevaContext) IsCmdEscrevaContext() {}

func NewCmdEscrevaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdEscrevaContext {
	var p = new(CmdEscrevaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdEscreva

	return p
}

func (s *CmdEscrevaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdEscrevaContext) ESCREVA() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserESCREVA, 0)
}

func (s *CmdEscrevaContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserABREPAR, 0)
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
	return s.GetToken(T2AlgumaLexerParserFECHAPAR, 0)
}

func (s *CmdEscrevaContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T2AlgumaLexerParserVIRGULA)
}

func (s *CmdEscrevaContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserVIRGULA, i)
}

func (s *CmdEscrevaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdEscrevaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdEscrevaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmdEscreva(s)
	}
}

func (s *CmdEscrevaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmdEscreva(s)
	}
}

func (p *T2AlgumaLexerParser) CmdEscreva() (localctx ICmdEscrevaContext) {
	localctx = NewCmdEscrevaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, T2AlgumaLexerParserRULE_cmdEscreva)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(T2AlgumaLexerParserESCREVA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.Match(T2AlgumaLexerParserABREPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(352)
		p.Expressao()
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserVIRGULA {
		{
			p.SetState(353)
			p.Match(T2AlgumaLexerParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.Expressao()
		}

		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(360)
		p.Match(T2AlgumaLexerParserFECHAPAR)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdSe
	return p
}

func InitEmptyCmdSeContext(p *CmdSeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdSe
}

func (*CmdSeContext) IsCmdSeContext() {}

func NewCmdSeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdSeContext {
	var p = new(CmdSeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdSe

	return p
}

func (s *CmdSeContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdSeContext) SE() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserSE, 0)
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
	return s.GetToken(T2AlgumaLexerParserENTAO, 0)
}

func (s *CmdSeContext) FIM_SE() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFIM_SE, 0)
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
	return s.GetToken(T2AlgumaLexerParserSENAO, 0)
}

func (s *CmdSeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdSeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdSeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmdSe(s)
	}
}

func (s *CmdSeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmdSe(s)
	}
}

func (p *T2AlgumaLexerParser) CmdSe() (localctx ICmdSeContext) {
	localctx = NewCmdSeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, T2AlgumaLexerParserRULE_cmdSe)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(T2AlgumaLexerParserSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.Expressao()
	}
	{
		p.SetState(364)
		p.Match(T2AlgumaLexerParserENTAO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(365)
			p.Cmd()
		}

		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T2AlgumaLexerParserSENAO {
		{
			p.SetState(371)
			p.Match(T2AlgumaLexerParserSENAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
			{
				p.SetState(372)
				p.Cmd()
			}

			p.SetState(377)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(380)
		p.Match(T2AlgumaLexerParserFIM_SE)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdCaso
	return p
}

func InitEmptyCmdCasoContext(p *CmdCasoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdCaso
}

func (*CmdCasoContext) IsCmdCasoContext() {}

func NewCmdCasoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdCasoContext {
	var p = new(CmdCasoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdCaso

	return p
}

func (s *CmdCasoContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdCasoContext) CASO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserCASO, 0)
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
	return s.GetToken(T2AlgumaLexerParserSEJA, 0)
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
	return s.GetToken(T2AlgumaLexerParserFIM_CASO, 0)
}

func (s *CmdCasoContext) SENAO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserSENAO, 0)
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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmdCaso(s)
	}
}

func (s *CmdCasoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmdCaso(s)
	}
}

func (p *T2AlgumaLexerParser) CmdCaso() (localctx ICmdCasoContext) {
	localctx = NewCmdCasoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, T2AlgumaLexerParserRULE_cmdCaso)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.Match(T2AlgumaLexerParserCASO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.Exp_aritmetica()
	}
	{
		p.SetState(384)
		p.Match(T2AlgumaLexerParserSEJA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)
		p.Selecao()
	}
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T2AlgumaLexerParserSENAO {
		{
			p.SetState(386)
			p.Match(T2AlgumaLexerParserSENAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
			{
				p.SetState(387)
				p.Cmd()
			}

			p.SetState(392)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(395)
		p.Match(T2AlgumaLexerParserFIM_CASO)
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

	// Getter signatures
	PARA() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ATRIBUICAO() antlr.TerminalNode
	AllExp_aritmetica() []IExp_aritmeticaContext
	Exp_aritmetica(i int) IExp_aritmeticaContext
	ATE() antlr.TerminalNode
	FACA() antlr.TerminalNode
	FIM_PARA() antlr.TerminalNode
	AllCmd() []ICmdContext
	Cmd(i int) ICmdContext

	// IsCmdParaContext differentiates from other interfaces.
	IsCmdParaContext()
}

type CmdParaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdParaContext() *CmdParaContext {
	var p = new(CmdParaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdPara
	return p
}

func InitEmptyCmdParaContext(p *CmdParaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdPara
}

func (*CmdParaContext) IsCmdParaContext() {}

func NewCmdParaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdParaContext {
	var p = new(CmdParaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdPara

	return p
}

func (s *CmdParaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdParaContext) PARA() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserPARA, 0)
}

func (s *CmdParaContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserIDENT, 0)
}

func (s *CmdParaContext) ATRIBUICAO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserATRIBUICAO, 0)
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

func (s *CmdParaContext) ATE() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserATE, 0)
}

func (s *CmdParaContext) FACA() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFACA, 0)
}

func (s *CmdParaContext) FIM_PARA() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFIM_PARA, 0)
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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmdPara(s)
	}
}

func (s *CmdParaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmdPara(s)
	}
}

func (p *T2AlgumaLexerParser) CmdPara() (localctx ICmdParaContext) {
	localctx = NewCmdParaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, T2AlgumaLexerParserRULE_cmdPara)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(T2AlgumaLexerParserPARA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Match(T2AlgumaLexerParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(399)
		p.Match(T2AlgumaLexerParserATRIBUICAO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
		p.Exp_aritmetica()
	}
	{
		p.SetState(401)
		p.Match(T2AlgumaLexerParserATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.Exp_aritmetica()
	}
	{
		p.SetState(403)
		p.Match(T2AlgumaLexerParserFACA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(404)
			p.Cmd()
		}

		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(410)
		p.Match(T2AlgumaLexerParserFIM_PARA)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdEnquanto
	return p
}

func InitEmptyCmdEnquantoContext(p *CmdEnquantoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdEnquanto
}

func (*CmdEnquantoContext) IsCmdEnquantoContext() {}

func NewCmdEnquantoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdEnquantoContext {
	var p = new(CmdEnquantoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdEnquanto

	return p
}

func (s *CmdEnquantoContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdEnquantoContext) ENQUANTO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserENQUANTO, 0)
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
	return s.GetToken(T2AlgumaLexerParserFACA, 0)
}

func (s *CmdEnquantoContext) FIM_ENQUANTO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFIM_ENQUANTO, 0)
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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmdEnquanto(s)
	}
}

func (s *CmdEnquantoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmdEnquanto(s)
	}
}

func (p *T2AlgumaLexerParser) CmdEnquanto() (localctx ICmdEnquantoContext) {
	localctx = NewCmdEnquantoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, T2AlgumaLexerParserRULE_cmdEnquanto)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)
		p.Match(T2AlgumaLexerParserENQUANTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.Expressao()
	}
	{
		p.SetState(414)
		p.Match(T2AlgumaLexerParserFACA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(415)
			p.Cmd()
		}

		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(421)
		p.Match(T2AlgumaLexerParserFIM_ENQUANTO)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdFaca
	return p
}

func InitEmptyCmdFacaContext(p *CmdFacaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdFaca
}

func (*CmdFacaContext) IsCmdFacaContext() {}

func NewCmdFacaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdFacaContext {
	var p = new(CmdFacaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdFaca

	return p
}

func (s *CmdFacaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdFacaContext) FACA() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFACA, 0)
}

func (s *CmdFacaContext) ATE() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserATE, 0)
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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmdFaca(s)
	}
}

func (s *CmdFacaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmdFaca(s)
	}
}

func (p *T2AlgumaLexerParser) CmdFaca() (localctx ICmdFacaContext) {
	localctx = NewCmdFacaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, T2AlgumaLexerParserRULE_cmdFaca)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(423)
		p.Match(T2AlgumaLexerParserFACA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(424)
			p.Cmd()
		}

		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(430)
		p.Match(T2AlgumaLexerParserATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdAtribuicao
	return p
}

func InitEmptyCmdAtribuicaoContext(p *CmdAtribuicaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdAtribuicao
}

func (*CmdAtribuicaoContext) IsCmdAtribuicaoContext() {}

func NewCmdAtribuicaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdAtribuicaoContext {
	var p = new(CmdAtribuicaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdAtribuicao

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
	return s.GetToken(T2AlgumaLexerParserATRIBUICAO, 0)
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
	return s.GetToken(T2AlgumaLexerParserPONTEIRO, 0)
}

func (s *CmdAtribuicaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdAtribuicaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdAtribuicaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmdAtribuicao(s)
	}
}

func (s *CmdAtribuicaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmdAtribuicao(s)
	}
}

func (p *T2AlgumaLexerParser) CmdAtribuicao() (localctx ICmdAtribuicaoContext) {
	localctx = NewCmdAtribuicaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, T2AlgumaLexerParserRULE_cmdAtribuicao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T2AlgumaLexerParserPONTEIRO {
		{
			p.SetState(433)
			p.Match(T2AlgumaLexerParserPONTEIRO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(436)
		p.Identificador()
	}
	{
		p.SetState(437)
		p.Match(T2AlgumaLexerParserATRIBUICAO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(438)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdChamada
	return p
}

func InitEmptyCmdChamadaContext(p *CmdChamadaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdChamada
}

func (*CmdChamadaContext) IsCmdChamadaContext() {}

func NewCmdChamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdChamadaContext {
	var p = new(CmdChamadaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdChamada

	return p
}

func (s *CmdChamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdChamadaContext) IDENT() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserIDENT, 0)
}

func (s *CmdChamadaContext) ABREPAR() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserABREPAR, 0)
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
	return s.GetToken(T2AlgumaLexerParserFECHAPAR, 0)
}

func (s *CmdChamadaContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(T2AlgumaLexerParserVIRGULA)
}

func (s *CmdChamadaContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserVIRGULA, i)
}

func (s *CmdChamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdChamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdChamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmdChamada(s)
	}
}

func (s *CmdChamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmdChamada(s)
	}
}

func (p *T2AlgumaLexerParser) CmdChamada() (localctx ICmdChamadaContext) {
	localctx = NewCmdChamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, T2AlgumaLexerParserRULE_cmdChamada)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.Match(T2AlgumaLexerParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.Match(T2AlgumaLexerParserABREPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(442)
		p.Expressao()
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserVIRGULA {
		{
			p.SetState(443)
			p.Match(T2AlgumaLexerParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.Expressao()
		}

		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(450)
		p.Match(T2AlgumaLexerParserFECHAPAR)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdRetorne
	return p
}

func InitEmptyCmdRetorneContext(p *CmdRetorneContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdRetorne
}

func (*CmdRetorneContext) IsCmdRetorneContext() {}

func NewCmdRetorneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdRetorneContext {
	var p = new(CmdRetorneContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_cmdRetorne

	return p
}

func (s *CmdRetorneContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdRetorneContext) RETORNE() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserRETORNE, 0)
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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterCmdRetorne(s)
	}
}

func (s *CmdRetorneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitCmdRetorne(s)
	}
}

func (p *T2AlgumaLexerParser) CmdRetorne() (localctx ICmdRetorneContext) {
	localctx = NewCmdRetorneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, T2AlgumaLexerParserRULE_cmdRetorne)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.Match(T2AlgumaLexerParserRETORNE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(453)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_selecao
	return p
}

func InitEmptySelecaoContext(p *SelecaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_selecao
}

func (*SelecaoContext) IsSelecaoContext() {}

func NewSelecaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelecaoContext {
	var p = new(SelecaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_selecao

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterSelecao(s)
	}
}

func (s *SelecaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitSelecao(s)
	}
}

func (p *T2AlgumaLexerParser) Selecao() (localctx ISelecaoContext) {
	localctx = NewSelecaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, T2AlgumaLexerParserRULE_selecao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserSUBTRACAO || _la == T2AlgumaLexerParserNUM_INT {
		{
			p.SetState(455)
			p.Item_selecao()
		}

		p.SetState(460)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_item_selecao
	return p
}

func InitEmptyItem_selecaoContext(p *Item_selecaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_item_selecao
}

func (*Item_selecaoContext) IsItem_selecaoContext() {}

func NewItem_selecaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Item_selecaoContext {
	var p = new(Item_selecaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_item_selecao

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
	return s.GetToken(T2AlgumaLexerParserDOISPONTOS, 0)
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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterItem_selecao(s)
	}
}

func (s *Item_selecaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitItem_selecao(s)
	}
}

func (p *T2AlgumaLexerParser) Item_selecao() (localctx IItem_selecaoContext) {
	localctx = NewItem_selecaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, T2AlgumaLexerParserRULE_item_selecao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(461)
		p.Constantes()
	}
	{
		p.SetState(462)
		p.Match(T2AlgumaLexerParserDOISPONTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&580542146808977) != 0 {
		{
			p.SetState(463)
			p.Cmd()
		}

		p.SetState(468)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_constantes
	return p
}

func InitEmptyConstantesContext(p *ConstantesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_constantes
}

func (*ConstantesContext) IsConstantesContext() {}

func NewConstantesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantesContext {
	var p = new(ConstantesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_constantes

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
	return s.GetTokens(T2AlgumaLexerParserVIRGULA)
}

func (s *ConstantesContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserVIRGULA, i)
}

func (s *ConstantesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterConstantes(s)
	}
}

func (s *ConstantesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitConstantes(s)
	}
}

func (p *T2AlgumaLexerParser) Constantes() (localctx IConstantesContext) {
	localctx = NewConstantesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, T2AlgumaLexerParserRULE_constantes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)
		p.Numero_intervalo()
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserVIRGULA {
		{
			p.SetState(470)
			p.Match(T2AlgumaLexerParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.Numero_intervalo()
		}

		p.SetState(476)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_numero_intervalo
	return p
}

func InitEmptyNumero_intervaloContext(p *Numero_intervaloContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_numero_intervalo
}

func (*Numero_intervaloContext) IsNumero_intervaloContext() {}

func NewNumero_intervaloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numero_intervaloContext {
	var p = new(Numero_intervaloContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_numero_intervalo

	return p
}

func (s *Numero_intervaloContext) GetParser() antlr.Parser { return s.parser }

func (s *Numero_intervaloContext) AllNUM_INT() []antlr.TerminalNode {
	return s.GetTokens(T2AlgumaLexerParserNUM_INT)
}

func (s *Numero_intervaloContext) NUM_INT(i int) antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserNUM_INT, i)
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
	return s.GetToken(T2AlgumaLexerParserINTERVALO, 0)
}

func (s *Numero_intervaloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numero_intervaloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numero_intervaloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterNumero_intervalo(s)
	}
}

func (s *Numero_intervaloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitNumero_intervalo(s)
	}
}

func (p *T2AlgumaLexerParser) Numero_intervalo() (localctx INumero_intervaloContext) {
	localctx = NewNumero_intervaloContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, T2AlgumaLexerParserRULE_numero_intervalo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T2AlgumaLexerParserSUBTRACAO {
		{
			p.SetState(477)
			p.Op_unario()
		}

	}
	{
		p.SetState(480)
		p.Match(T2AlgumaLexerParserNUM_INT)
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

	if _la == T2AlgumaLexerParserINTERVALO {
		{
			p.SetState(481)
			p.Match(T2AlgumaLexerParserINTERVALO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T2AlgumaLexerParserSUBTRACAO {
			{
				p.SetState(482)
				p.Op_unario()
			}

		}
		{
			p.SetState(485)
			p.Match(T2AlgumaLexerParserNUM_INT)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_exp_relacional
	return p
}

func InitEmptyExp_relacionalContext(p *Exp_relacionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_exp_relacional
}

func (*Exp_relacionalContext) IsExp_relacionalContext() {}

func NewExp_relacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_relacionalContext {
	var p = new(Exp_relacionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_exp_relacional

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterExp_relacional(s)
	}
}

func (s *Exp_relacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitExp_relacional(s)
	}
}

func (p *T2AlgumaLexerParser) Exp_relacional() (localctx IExp_relacionalContext) {
	localctx = NewExp_relacionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, T2AlgumaLexerParserRULE_exp_relacional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(488)
		p.Exp_aritmetica()
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69269232549888) != 0 {
		{
			p.SetState(489)
			p.Op_relacional()
		}
		{
			p.SetState(490)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_op_relacional
	return p
}

func InitEmptyOp_relacionalContext(p *Op_relacionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_op_relacional
}

func (*Op_relacionalContext) IsOp_relacionalContext() {}

func NewOp_relacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_relacionalContext {
	var p = new(Op_relacionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_op_relacional

	return p
}

func (s *Op_relacionalContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_relacionalContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserIGUAL, 0)
}

func (s *Op_relacionalContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserDIFERENTE, 0)
}

func (s *Op_relacionalContext) MAIORIGUAL() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserMAIORIGUAL, 0)
}

func (s *Op_relacionalContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserMENORIGUAL, 0)
}

func (s *Op_relacionalContext) MAIOR() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserMAIOR, 0)
}

func (s *Op_relacionalContext) MENOR() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserMENOR, 0)
}

func (s *Op_relacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_relacionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_relacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterOp_relacional(s)
	}
}

func (s *Op_relacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitOp_relacional(s)
	}
}

func (p *T2AlgumaLexerParser) Op_relacional() (localctx IOp_relacionalContext) {
	localctx = NewOp_relacionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, T2AlgumaLexerParserRULE_op_relacional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(494)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_expressao
	return p
}

func InitEmptyExpressaoContext(p *ExpressaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_expressao
}

func (*ExpressaoContext) IsExpressaoContext() {}

func NewExpressaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressaoContext {
	var p = new(ExpressaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_expressao

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterExpressao(s)
	}
}

func (s *ExpressaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitExpressao(s)
	}
}

func (p *T2AlgumaLexerParser) Expressao() (localctx IExpressaoContext) {
	localctx = NewExpressaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, T2AlgumaLexerParserRULE_expressao)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.Termo_logico()
	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserOU {
		{
			p.SetState(497)
			p.Op_logico_1()
		}
		{
			p.SetState(498)
			p.Termo_logico()
		}

		p.SetState(504)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_termo_logico
	return p
}

func InitEmptyTermo_logicoContext(p *Termo_logicoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_termo_logico
}

func (*Termo_logicoContext) IsTermo_logicoContext() {}

func NewTermo_logicoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Termo_logicoContext {
	var p = new(Termo_logicoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_termo_logico

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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterTermo_logico(s)
	}
}

func (s *Termo_logicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitTermo_logico(s)
	}
}

func (p *T2AlgumaLexerParser) Termo_logico() (localctx ITermo_logicoContext) {
	localctx = NewTermo_logicoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, T2AlgumaLexerParserRULE_termo_logico)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)
		p.Fator_logico()
	}
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T2AlgumaLexerParserE {
		{
			p.SetState(506)
			p.Op_logico_2()
		}
		{
			p.SetState(507)
			p.Fator_logico()
		}

		p.SetState(513)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_fator_logico
	return p
}

func InitEmptyFator_logicoContext(p *Fator_logicoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_fator_logico
}

func (*Fator_logicoContext) IsFator_logicoContext() {}

func NewFator_logicoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fator_logicoContext {
	var p = new(Fator_logicoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_fator_logico

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
	return s.GetToken(T2AlgumaLexerParserNAO, 0)
}

func (s *Fator_logicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fator_logicoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fator_logicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterFator_logico(s)
	}
}

func (s *Fator_logicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitFator_logico(s)
	}
}

func (p *T2AlgumaLexerParser) Fator_logico() (localctx IFator_logicoContext) {
	localctx = NewFator_logicoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, T2AlgumaLexerParserRULE_fator_logico)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T2AlgumaLexerParserNAO {
		{
			p.SetState(514)
			p.Match(T2AlgumaLexerParserNAO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(517)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela_logica
	return p
}

func InitEmptyParcela_logicaContext(p *Parcela_logicaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela_logica
}

func (*Parcela_logicaContext) IsParcela_logicaContext() {}

func NewParcela_logicaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parcela_logicaContext {
	var p = new(Parcela_logicaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_parcela_logica

	return p
}

func (s *Parcela_logicaContext) GetParser() antlr.Parser { return s.parser }

func (s *Parcela_logicaContext) VERDADEIRO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserVERDADEIRO, 0)
}

func (s *Parcela_logicaContext) FALSO() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserFALSO, 0)
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
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterParcela_logica(s)
	}
}

func (s *Parcela_logicaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitParcela_logica(s)
	}
}

func (p *T2AlgumaLexerParser) Parcela_logica() (localctx IParcela_logicaContext) {
	localctx = NewParcela_logicaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, T2AlgumaLexerParserRULE_parcela_logica)
	var _la int

	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T2AlgumaLexerParserVERDADEIRO, T2AlgumaLexerParserFALSO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(519)
			_la = p.GetTokenStream().LA(1)

			if !(_la == T2AlgumaLexerParserVERDADEIRO || _la == T2AlgumaLexerParserFALSO) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case T2AlgumaLexerParserABREPAR, T2AlgumaLexerParserSUBTRACAO, T2AlgumaLexerParserPONTEIRO, T2AlgumaLexerParserENDERECO, T2AlgumaLexerParserNUM_INT, T2AlgumaLexerParserNUM_REAL, T2AlgumaLexerParserIDENT, T2AlgumaLexerParserCADEIA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(520)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_op_logico_1
	return p
}

func InitEmptyOp_logico_1Context(p *Op_logico_1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_op_logico_1
}

func (*Op_logico_1Context) IsOp_logico_1Context() {}

func NewOp_logico_1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_logico_1Context {
	var p = new(Op_logico_1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_op_logico_1

	return p
}

func (s *Op_logico_1Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_logico_1Context) OU() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserOU, 0)
}

func (s *Op_logico_1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_logico_1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_logico_1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterOp_logico_1(s)
	}
}

func (s *Op_logico_1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitOp_logico_1(s)
	}
}

func (p *T2AlgumaLexerParser) Op_logico_1() (localctx IOp_logico_1Context) {
	localctx = NewOp_logico_1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, T2AlgumaLexerParserRULE_op_logico_1)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(523)
		p.Match(T2AlgumaLexerParserOU)
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
	p.RuleIndex = T2AlgumaLexerParserRULE_op_logico_2
	return p
}

func InitEmptyOp_logico_2Context(p *Op_logico_2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T2AlgumaLexerParserRULE_op_logico_2
}

func (*Op_logico_2Context) IsOp_logico_2Context() {}

func NewOp_logico_2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_logico_2Context {
	var p = new(Op_logico_2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T2AlgumaLexerParserRULE_op_logico_2

	return p
}

func (s *Op_logico_2Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_logico_2Context) E() antlr.TerminalNode {
	return s.GetToken(T2AlgumaLexerParserE, 0)
}

func (s *Op_logico_2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_logico_2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_logico_2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.EnterOp_logico_2(s)
	}
}

func (s *Op_logico_2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(T2AlgumaLexerListener); ok {
		listenerT.ExitOp_logico_2(s)
	}
}

func (p *T2AlgumaLexerParser) Op_logico_2() (localctx IOp_logico_2Context) {
	localctx = NewOp_logico_2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, T2AlgumaLexerParserRULE_op_logico_2)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(525)
		p.Match(T2AlgumaLexerParserE)
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
