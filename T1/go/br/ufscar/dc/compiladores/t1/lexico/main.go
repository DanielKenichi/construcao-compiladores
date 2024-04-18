package main

import (
	"log"
	"os"

	"github.com/DanielKenichi/construcao-compiladores/T1/antlr4/br/ufscar/dc/compiladores/t1/lexico/parser"
	"github.com/DanielKenichi/construcao-compiladores/T1/go/br/ufscar/dc/compiladores/t1/lexico/vocabulary"
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

	lexer := parser.NewT1AlgumaLexer(input)

	vocabulary := vocabulary.New(lexer.LiteralNames, lexer.SymbolicNames)

	for {
		t := lexer.NextToken()

		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		_, err = output.WriteString("<'" + t.GetText() + "'," + *vocabulary.GetDisplayName(t.GetTokenType()) + ">\n")

		if err != nil {
			log.Fatalf("Failed writing to output file: %v", err)
		}
	}

}
