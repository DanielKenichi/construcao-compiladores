package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DanielKenichi/construcao-compiladores/T1/antlr4/br/ufscar/dc/compiladores/t1/lexico/parser"
	"github.com/DanielKenichi/construcao-compiladores/T1/go/br/ufscar/dc/compiladores/t1/lexico/errortokens"
	"github.com/DanielKenichi/construcao-compiladores/T1/go/br/ufscar/dc/compiladores/t1/lexico/vocabulary"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	//Leitura dos argumentos
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	//Nova stream de input
	input, err := antlr.NewFileStream(inputFile)

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	//Abertura do arquivo de output
	output, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("Error opening output file: %v", err)
	}

	lexer := parser.NewT1AlgumaLexer(input)

	//Instanciando vocabulario para retornar nome de display dos tokens
	vocabulary := vocabulary.New(lexer.LiteralNames, lexer.SymbolicNames)

	for {
		//leitura do proximo token
		t := lexer.NextToken()

		//Para a leitura caso tenha chego no fim do arquivo
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		//Recuperando no vocabulario o nome do token
		tokenName := *vocabulary.GetDisplayName(t.GetTokenType())

		//Verificando se o Token eh um erro lexico
		if errortokens.IsTokenAError(tokenName) {
			//escrevendo erro no arquivo de output e parando leitura
			_, err = output.WriteString(errortokens.WriteError(tokenName, t))

			if err != nil {
				log.Fatalf("Failed writing to output file: %v", err)
			}

			break
		} else {
			//escrevendo token no arquivo de output
			_, err = output.WriteString(fmt.Sprintf("<'%s',%s>\n", t.GetText(), tokenName))
		}

		if err != nil {
			log.Fatalf("Failed writing to output file: %v", err)
		}
	}

}
