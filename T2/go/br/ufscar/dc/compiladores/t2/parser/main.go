package main

import (
	"log"
	"os"

	"github.com/DanielKenichi/construcao-compiladores/T2/antlr4/br/ufscar/dc/compiladores/t2/parser/parser"
	"github.com/DanielKenichi/construcao-compiladores/T2/go/br/ufscar/dc/compiladores/t2/parser/errortokens"
	"github.com/DanielKenichi/construcao-compiladores/T2/go/br/ufscar/dc/compiladores/t2/parser/vocabulary"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	inputFile := os.Args[1]
	outputFile := os.Args[2]

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

	lexer := parser.NewT2AlgumaLexerLexer(input)

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

			break
		} else {
			// _, err = output.WriteString(fmt.Sprintf("<'%s',%s>\n", t.GetText(), tokenName))
		}

		if err != nil {
			log.Fatalf("Failed writing to output file: %v", err)
		}
	}

	// -----------------------------------
	// ANALISE SINTÁTICA (T2)
	// -----------------------------------

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewT2AlgumaLexerParser(tokens)

	// Adicionando nosso ErrorListener customizado
	parser.RemoveErrorListeners()
	mcel := NewCustomErrorListener(output)
	parser.AddErrorListener(mcel)

	parser.Programa()

	output.WriteString("Fim da compilacao")
}
