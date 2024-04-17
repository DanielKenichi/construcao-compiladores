package main

import (
	"fmt"
	"os"

	"github.com/DanielKenichi/construcao-compiladores/T1/antlr4/br/ufscar/dc/compiladores/t1/lexico/parser"
)

func main() {
	aa := parser.T1AlgumaLexer{}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	fmt.Println(inputFile)
	fmt.Println(outputFile)
	fmt.Println(aa)

}
