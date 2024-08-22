package main

import (
	"log"
	"os"

	parser "github.com/DanielKenichi/construcao-compiladores/T5/antlr4/br/ufscar/dc/compiladores/t5/parser/Alguma"
	"github.com/DanielKenichi/construcao-compiladores/T5/go/br/ufscar/dc/compiladores/t5/errortokens"
	"github.com/DanielKenichi/construcao-compiladores/T5/go/br/ufscar/dc/compiladores/t5/generator"
	"github.com/DanielKenichi/construcao-compiladores/T5/go/br/ufscar/dc/compiladores/t5/visitor"
	"github.com/DanielKenichi/construcao-compiladores/T5/go/br/ufscar/dc/compiladores/t5/vocabulary"
	"github.com/antlr4-go/antlr/v4"
)

//Descomente e arrume o path dessas variaveis para debug

// var inputFile string = "/home/lucky/UFSCAR/Compiladores/construcao-compiladores/corretor/casos-de-teste/2.casos_teste_t2/entrada/1-algoritmo_2-2_apostila_LA_1_erro_linha_3_acusado_linha_10.txt"
// var outputFile string = "./test.out"

var inputFile string
var outputFile string

func main() {

	if inputFile == "" {
		inputFile = os.Args[1]
	}

	if outputFile == "" {
		outputFile = os.Args[2]
	}

	input, err := antlr.NewFileStream(inputFile)

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	output, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("Error opening output file: %v", err)
	}

	// -----------------------------------
	// ANALISE LÉXICA (T1)
	// -----------------------------------

	lexer := parser.NewT5AlgumaLexer(input)

	//Instanciando vocabulario para retornar nome de display dos tokens
	vocabulary := vocabulary.New(lexer.LiteralNames, lexer.SymbolicNames)

	for {
		t := lexer.NextToken()

		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		tokenName := *vocabulary.GetDisplayName(t.GetTokenType())

		if errortokens.IsTokenAError(tokenName) {
			_, err = output.WriteString(errortokens.WriteError(tokenName, t))

			if err != nil {
				log.Fatalf("Failed writing to output file: %v", err)
			}

			output.WriteString("Fim da compilacao\n")
			return
		}

		if err != nil {
			log.Fatalf("Failed writing to output file: %v", err)
		}
	}

	// -----------------------------------
	// ANALISE SINTÁTICA (T2)
	// -----------------------------------
	lexer.Reset()

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewT5AlgumaParser(tokens)

	// Adicionando nosso ErrorListener customizado
	parser.RemoveErrorListeners()
	customErrorListener := NewCustomErrorListener(output)
	parser.AddErrorListener(customErrorListener)

	tree := parser.Programa()

	// -----------------------------------
	// ANALISE SEMÂNTICA (T3) e (T4)
	// -----------------------------------

	algumaVisitor := visitor.New()

	semanticErrors := algumaVisitor.VisitPrograma(tree)

	for _, semanticError := range semanticErrors {
		output.WriteString(semanticError)
	}

	// -----------------------------------
	// GERADOR DE CÓDIGO C
	// -----------------------------------

	algumaGenerator := generator.New()

	programOutput := algumaGenerator.VisitPrograma(tree)

	for _, line := range programOutput {
		output.WriteString(line)
	}

	output.WriteString("// Fim da compilacao\n")
}
