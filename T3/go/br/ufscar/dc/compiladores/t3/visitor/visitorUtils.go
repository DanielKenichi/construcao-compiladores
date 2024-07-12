package visitor

import (
	"fmt"

	parser "github.com/DanielKenichi/construcao-compiladores/T3/antlr4/br/ufscar/dc/compiladores/t3/parser/Alguma"
	"github.com/DanielKenichi/construcao-compiladores/T3/go/br/ufscar/dc/compiladores/t3/symboltable"
	"github.com/antlr4-go/antlr/v4"
)

func (v *AlgumaVisitor) AddVarToSymbolTable(variavel parser.IVariavelContext) []string {

	addSymbolTableResult := make([]string, 0)

	var varType symboltable.Type = symboltable.INVALIDO

	if variavel.Tipo().Registro() != nil {
		varType = symboltable.REGISTRO
	} else if variavel.Tipo().Tipo_variavel().PONTEIRO() != nil {
		varType = symboltable.PONTEIRO
	} else {
		basicType := variavel.Tipo().Tipo_variavel().Tipo_basico()

		varType = MapBasicTypeToSymbolType(basicType)
	}

	for _, idents := range variavel.AllIdentificador() {
		for _, identName := range idents.AllIDENT() {
			result := v.AddIdentifierToSymbolTable(identName, varType)

			addSymbolTableResult = append(addSymbolTableResult, result...)
		}
	}

	return addSymbolTableResult
}

func (v *AlgumaVisitor) AddConstToSymbolTable(identifier antlr.TerminalNode, basicType parser.ITipo_basicoContext) []string {

	varType := MapBasicTypeToSymbolType(basicType)

	result := v.AddIdentifierToSymbolTable(identifier, varType)

	return result
}

func MapBasicTypeToSymbolType(basicType parser.ITipo_basicoContext) symboltable.Type {

	if basicType.LITERAL() != nil {
		return symboltable.LITERAL
	} else if basicType.INTEIRO() != nil {
		return symboltable.INTEIRO
	} else if basicType.REAL() != nil {
		return symboltable.REAL
	} else if basicType.LOGICO() != nil {
		return symboltable.LOGICO
	}

	return symboltable.INVALIDO
}

func (v *AlgumaVisitor) AddIdentifierToSymbolTable(identifier antlr.TerminalNode, varType symboltable.Type) []string {
	result := make([]string, 0)

	if v.Scopes.CurrentScope().Exists(identifier.GetText()) {
		result = append(result,
			semanticError(identifier.GetSymbol(), fmt.Sprintf("identificador %v ja declarado anteriormente", identifier.GetText())),
		)
	}

	v.Scopes.CurrentScope().AddSymbol(identifier.GetText(), varType)

	return result
}
