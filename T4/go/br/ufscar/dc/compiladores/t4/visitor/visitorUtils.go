package visitor

import (
	"fmt"

	parser "github.com/DanielKenichi/construcao-compiladores/T4/antlr4/br/ufscar/dc/compiladores/t4/parser/Alguma"
	"github.com/DanielKenichi/construcao-compiladores/T4/go/br/ufscar/dc/compiladores/t4/symboltable"
	"github.com/antlr4-go/antlr/v4"
)

func (v *AlgumaVisitor) AddVarToSymbolTable(variavel parser.IVariavelContext, table *symboltable.SymbolTable) []string {

	addSymbolTableResult := make([]string, 0)

	var varType symboltable.Type = symboltable.INVALIDO

	if variavel.Tipo().Registro() != nil {
		varType = symboltable.REGISTRO_VAR
	} else if variavel.Tipo().Tipo_variavel().PONTEIRO() != nil {
		varType = symboltable.PONTEIRO
	} else if variavel.Tipo().Tipo_variavel().Tipo_basico() != nil {
		basicType := variavel.Tipo().Tipo_variavel().Tipo_basico()

		varType = MapBasicTypeToSymbolType(basicType)
	} else { //Tipo variavel eh um ident de um registro
		varType = symboltable.REGISTRO_VAR
	}

	for _, idents := range variavel.AllIdentificador() {
		for _, identName := range idents.AllIDENT() {
			result := v.AddIdentifierToSymbolTable(identName, varType, table)

			addSymbolTableResult = append(addSymbolTableResult, result...)

			if varType == symboltable.REGISTRO_VAR {
				if variavel.Tipo().Registro() != nil {
					regTable := table.GetSymbol(identName.GetText()).InnerTable

					result = v.AddRegInnerVarsToSymbolTable(variavel.Tipo().Registro(), regTable)

					addSymbolTableResult = append(addSymbolTableResult, result...)
				} else {
					regSymbol := v.GetRegTypeSymbol(variavel.Tipo().Tipo_variavel().IDENT().GetText())

					table.AddInnerTableToRegVar(identName.GetText(), *regSymbol.InnerTable)
				}
			}
		}
	}

	return addSymbolTableResult
}

func (v *AlgumaVisitor) GetRegTypeSymbol(name string) *symboltable.Symbol {

	for _, scope := range v.Scopes.Stack {
		if scope.Exists(name) {
			symbol := scope.GetSymbol(name)
			if symbol.SymbolType == symboltable.REGISTRO {
				return symbol
			}
		}
	}

	return nil
}

func (v *AlgumaVisitor) GetRegVarSymbol(name string) *symboltable.Symbol {

	for _, scope := range v.Scopes.Stack {
		if scope.Exists(name) {
			symbol := scope.GetSymbol(name)
			if symbol.SymbolType == symboltable.REGISTRO_VAR {
				return symbol
			}
		}
	}

	return nil
}

func (v *AlgumaVisitor) GetFuncVarSymbol(name string) *symboltable.Symbol {

	for _, scope := range v.Scopes.Stack {
		if scope.Exists(name) {
			symbol := scope.GetSymbol(name)
			if symbol.SymbolType == symboltable.FUNCAO || symbol.SymbolType == symboltable.PROCEDIMENTO {
				return symbol
			}
		}
	}

	return nil
}

func (v *AlgumaVisitor) AddConstToSymbolTable(identifier antlr.TerminalNode, basicType parser.ITipo_basicoContext) []string {

	varType := MapBasicTypeToSymbolType(basicType)

	result := v.AddIdentifierToSymbolTable(identifier, varType, v.Scopes.CurrentScope())

	return result
}

func (v *AlgumaVisitor) AddFuncToSymbolTable(function parser.IDeclaracoes_funcoesContext) []string {

	addFuncResult := make([]string, 0)

	var varType symboltable.Type

	if function.Tipo_variavel() != nil {
		varType = symboltable.FUNCAO
	} else {
		varType = symboltable.PROCEDIMENTO
	}

	result := v.AddIdentifierToSymbolTable(function.IDENT(), varType, v.Scopes.CurrentScope())

	addFuncResult = append(addFuncResult, result...)

	funcTable := v.Scopes.CurrentScope().GetSymbol(function.IDENT().GetText()).InnerTable

	result = v.AddParamsToSymbolTable(function, funcTable)

	addFuncResult = append(addFuncResult, result...)

	if function.Tipo_variavel() != nil {
		result = v.AddReturnTypeToSymbolTable(function.IDENT(), function.Tipo_variavel(), v.Scopes.CurrentScope())

		addFuncResult = append(addFuncResult, result...)
	}

	return addFuncResult
}

func (v *AlgumaVisitor) AddParamsToSymbolTable(function parser.IDeclaracoes_funcoesContext, table *symboltable.SymbolTable) []string {
	addParamsResult := make([]string, 0)

	for _, parametro := range function.Parametros().AllParametro() {
		var paramType symboltable.Type

		if parametro.Tipo_variavel().PONTEIRO() != nil {
			paramType = symboltable.PONTEIRO
		} else if parametro.Tipo_variavel().Tipo_basico() != nil {
			basicType := parametro.Tipo_variavel().Tipo_basico()

			paramType = MapBasicTypeToSymbolType(basicType)
		} else { //Tipo variavel eh um ident de um registro
			paramType = symboltable.REGISTRO_VAR
		}

		for _, idents := range parametro.AllIdentificador() {
			for _, identName := range idents.AllIDENT() {
				result := v.AddIdentifierToSymbolTable(identName, paramType, table)

				addParamsResult = append(addParamsResult, result...)
			}
		}
	}

	return addParamsResult
}

func (v *AlgumaVisitor) AddReturnTypeToSymbolTable(ident antlr.TerminalNode, varType parser.ITipo_variavelContext, table *symboltable.SymbolTable) []string {
	returnTypeResult := make([]string, 0)

	var paramType string

	if varType.Tipo_basico() != nil {
		paramType = varType.Tipo_basico().GetText()
	} else {
		ident := varType.IDENT()
		exists := v.VerifyIdentOnValidScopes(ident)

		if !exists {
			returnTypeResult = append(returnTypeResult,
				semanticError(ident.GetSymbol(), fmt.Sprintf("tipo %v nao declarado", ident.GetText())),
			)
			return returnTypeResult
		}

		paramType = ident.GetText()
	}

	table.AddReturnTypeToSymbol(ident.GetText(), paramType)

	return returnTypeResult
}

func (v *AlgumaVisitor) AddRegTypeToSymbolTable(identifier antlr.TerminalNode, reg parser.IRegistroContext) []string {
	addRegResult := make([]string, 0)

	varType := symboltable.REGISTRO

	result := v.AddIdentifierToSymbolTable(identifier, varType, v.Scopes.CurrentScope())

	addRegResult = append(addRegResult, result...)

	regTable := v.Scopes.CurrentScope().GetSymbol(identifier.GetText()).InnerTable

	result = v.AddRegInnerVarsToSymbolTable(reg, regTable)

	addRegResult = append(addRegResult, result...)

	return addRegResult

}

func (v *AlgumaVisitor) AddRegInnerVarsToSymbolTable(reg parser.IRegistroContext, table *symboltable.SymbolTable) []string {

	result := make([]string, 0)

	for _, variavel := range reg.AllVariavel() {
		result = v.AddVarToSymbolTable(variavel, table)

		if len(result) > 0 {
			return result
		}
	}

	return result

}

func MapBasicTypeToSymbolType(basicType parser.ITipo_basicoContext) symboltable.Type {

	if basicType == nil {
		return symboltable.INVALIDO
	}

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

func MapStringToType(t string) symboltable.Type {

	if t == "literal" {
		return symboltable.LITERAL
	} else if t == "inteiro" {
		return symboltable.INTEIRO
	} else if t == "real" {
		return symboltable.REAL
	} else if t == "logico" {
		return symboltable.LOGICO
	}

	return symboltable.REGISTRO_VAR
}

func (v *AlgumaVisitor) AddIdentifierToSymbolTable(identifier antlr.TerminalNode, varType symboltable.Type, table *symboltable.SymbolTable) []string {
	result := make([]string, 0)

	for _, scope := range v.Scopes.Stack {
		if scope.Exists(identifier.GetText()) {
			result = append(result,
				semanticError(identifier.GetSymbol(), fmt.Sprintf("identificador %v ja declarado anteriormente", identifier.GetText())),
			)
			return result
		}
	}

	table.AddSymbol(identifier.GetText(), varType)

	return result
}

func IsTypesCompatible(t1, t2 symboltable.Type) bool {

	if t1 == symboltable.INVALIDO || t2 == symboltable.INVALIDO {
		return false
	}

	return t1 == t2 ||
		(t1 == symboltable.INTEIRO && t2 == symboltable.REAL) ||
		(t1 == symboltable.REAL && t2 == symboltable.INTEIRO) ||
		(t1 == symboltable.PONTEIRO && t2 == symboltable.ENDERECO) ||
		(t1 == symboltable.ENDERECO && t2 == symboltable.PONTEIRO)

}

func (v *AlgumaVisitor) VerifyVarAttribution(atribuicao parser.ICmdAtribuicaoContext) []string {

	varAttrResult := v.VerifyIdentifier(atribuicao.Identificador())

	if len(varAttrResult) > 0 {
		return varAttrResult
	}

	var attrType symboltable.Type

	// TODO: definir o tipo do ponteiro, para permimitir mudar o seu valor
	/*
		Exemplo:
		declare valor: inteiro { int normal }
		declare ponteiro: ^inteiro { ponteiro para int }
		valor <- 100 { normal }
		ponteiro <- &valor { normal }
		^ponteiro <- 100 { esperado: normal }
		^ponteiro <- "100" { esperado: erro }
	*/

	attrType, varAttrResult = v.VerifyExpression(atribuicao.Expressao())

	/*Se tiver um identificador nao declarado na expressao
	o resultado vai ter pelo menos um erro semantico.
	Dessa forma, retornar ele antes para nao ter tambem erro de atribuicao
	incompativel, uma vez que um identificador nao declarado retorna
	tipo invalido para a expressao */

	if len(varAttrResult) > 0 {
		return varAttrResult
	}

	identSufix := ""
	if atribuicao.Identificador().ABRECHAVE(0) != nil {
		identSufix = atribuicao.Identificador().ABRECHAVE(0).GetText()
		identSufix += atribuicao.Identificador().Exp_aritmetica(0).GetText()
		identSufix += atribuicao.Identificador().FECHACHAVE(0).GetText()
	}

	var identifier antlr.TerminalNode
	scopeToVerify := v.Scopes.CurrentScope()
	identPrefix := ""
	//Se tiver mais de um ident, estamos lidando com um registro
	for i, ident := range atribuicao.Identificador().AllIDENT() {
		identifier = ident
		regVarSymbol := v.GetRegVarSymbol(ident.GetText())

		if regVarSymbol != nil && atribuicao.Identificador().IDENT(i+1) != nil {
			identPrefix += ident.GetText() + "."
			scopeToVerify = regVarSymbol.InnerTable
		}
	}

	identType := scopeToVerify.GetType(identifier.GetText())

	if !IsTypesCompatible(identType, attrType) {

		if identType == symboltable.PONTEIRO {
			identPrefix = atribuicao.PONTEIRO().GetText()
		}

		varAttrResult = append(varAttrResult,
			semanticError(identifier.GetSymbol(), fmt.Sprintf("atribuicao nao compativel para %v%v%v", identPrefix, identifier.GetText(), identSufix)),
		)
	}

	return varAttrResult
}

func (v *AlgumaVisitor) VerifyIdentOnValidScopes(ident antlr.TerminalNode) bool {
	exists := false

	var idx int

	for i, scope := range v.Scopes.Stack {
		if scope == v.Scopes.CurrentScope() {
			idx = i
			break
		}
	}

	for _, scope := range v.Scopes.Stack[:idx+1] {
		if scope.Exists(ident.GetText()) {
			exists = true
		}
	}

	return exists
}

func (v *AlgumaVisitor) VerifyIdentifier(ident parser.IIdentificadorContext) []string {
	identResult := make([]string, 0)

	scopeToVerify := v.Scopes.CurrentScope()
	prefix := ""

	//Se tiver mais de um ident dentro de um identificador, estamos lidando com um registro
	for _, argToVerify := range ident.AllIDENT() {

		if !scopeToVerify.Exists(argToVerify.GetText()) {
			identResult = append(identResult,
				semanticError(argToVerify.GetSymbol(), fmt.Sprintf("identificador %v nao declarado", prefix+argToVerify.GetText())),
			)
		}

		regSymbol := v.GetRegVarSymbol(argToVerify.GetText())

		if regSymbol != nil {
			scopeToVerify = regSymbol.InnerTable
		}

		prefix += argToVerify.GetText() + "."
	}

	if len(identResult) > 1 {
		identResult = identResult[len(identResult)-1:]
	}

	return identResult
}

func (v *AlgumaVisitor) VerifyExpression(expression parser.IExpressaoContext) (symboltable.Type, []string) {
	var expType symboltable.Type
	result := make([]string, 0)
	aux := 0

	for _, termoLogico := range expression.AllTermo_logico() {

		currentType, termoLogicoResult := v.VerifyTermoLogico(termoLogico)

		result = append(result, termoLogicoResult...)

		if aux != 0 && !IsTypesCompatible(currentType, expType) {
			return symboltable.INVALIDO, result
		}

		expType = currentType
		aux++
	}

	if len(expression.AllOp_logico_1()) > 0 {
		expType = symboltable.LOGICO
	}

	return expType, result
}

func (v *AlgumaVisitor) VerifyTermoLogico(termoLogico parser.ITermo_logicoContext) (symboltable.Type, []string) {

	var lastType symboltable.Type
	result := make([]string, 0)
	aux := 0

	for _, fatorLogico := range termoLogico.AllFator_logico() {
		currentType, parcelaLogicaResult := v.VerifyParcelaLogica(fatorLogico.Parcela_logica())

		result = append(result, parcelaLogicaResult...)

		if aux != 0 && !IsTypesCompatible(currentType, lastType) {
			return symboltable.INVALIDO, result
		}

		lastType = currentType
		aux++
	}

	if len(termoLogico.AllOp_logico_2()) > 0 {
		lastType = symboltable.LOGICO
	}

	return lastType, result
}

func (v *AlgumaVisitor) VerifyParcelaLogica(parcelaLogica parser.IParcela_logicaContext) (symboltable.Type, []string) {
	result := make([]string, 0)
	var parcelaType symboltable.Type = symboltable.INVALIDO

	if parcelaLogica.VERDADEIRO() != nil || parcelaLogica.FALSO() != nil {
		return symboltable.LOGICO, result
	}

	aux := 0

	for _, expArit := range parcelaLogica.Exp_relacional().AllExp_aritmetica() {

		currentType, expAritResult := v.VerifyExpressaoAritmetica(expArit)

		result = append(result, expAritResult...)

		if aux != 0 && !IsTypesCompatible(parcelaType, currentType) {
			return symboltable.INVALIDO, result
		}

		parcelaType = currentType
		aux++
	}

	if parcelaLogica.Exp_relacional().Op_relacional() != nil {
		parcelaType = symboltable.LOGICO
	}

	return parcelaType, result
}

func (v *AlgumaVisitor) VerifyExpressaoAritmetica(expAritmetica parser.IExp_aritmeticaContext) (symboltable.Type, []string) {
	result := make([]string, 0)
	var lastType symboltable.Type = symboltable.INVALIDO
	aux := 0

	for _, termo := range expAritmetica.AllTermo() {

		currentType, termoResult := v.VerifyTermo(termo)

		result = append(result, termoResult...)

		if aux != 0 && !IsTypesCompatible(currentType, lastType) {
			return symboltable.INVALIDO, result
		}

		lastType = currentType

		aux++
	}

	return lastType, result
}

func (v *AlgumaVisitor) VerifyTermo(termo parser.ITermoContext) (symboltable.Type, []string) {
	result := make([]string, 0)
	var lastType symboltable.Type = symboltable.INVALIDO
	aux := 0

	for _, fator := range termo.AllFator() {
		currentType, fatorResult := v.VerifyFator(fator)

		result = append(result, fatorResult...)

		if aux != 0 && !IsTypesCompatible(currentType, lastType) {
			return symboltable.INVALIDO, result
		}

		lastType = currentType

		aux++
	}

	return lastType, result
}

func (v *AlgumaVisitor) VerifyFator(fator parser.IFatorContext) (symboltable.Type, []string) {

	result := make([]string, 0)
	var lastType symboltable.Type = symboltable.INVALIDO
	aux := 0

	for _, parcela := range fator.AllParcela() {

		currentType, parcelaResult := v.VerifyParcela(parcela)

		result = append(result, parcelaResult...)

		if aux != 0 && !IsTypesCompatible(currentType, lastType) {
			return symboltable.INVALIDO, result
		}

		lastType = currentType

		aux++
	}

	return lastType, result
}

func (v *AlgumaVisitor) VerifyParcela(parcela parser.IParcelaContext) (symboltable.Type, []string) {
	parcelaResult := make([]string, 0)

	if parcela.Parcela_unario() != nil {

		if parcela.Parcela_unario().Identificador() != nil {
			identifier := parcela.Parcela_unario().Identificador()
			identResult := v.VerifyIdentifier(identifier)

			if len(identResult) > 0 {
				parcelaResult = append(parcelaResult, identResult...)

				return symboltable.INVALIDO, parcelaResult
			}

			return v.GetIdentifierType(identifier), parcelaResult

		} else if parcela.Parcela_unario().NUM_INT() != nil {
			return symboltable.INTEIRO, parcelaResult
		} else if parcela.Parcela_unario().NUM_REAL() != nil {
			return symboltable.REAL, parcelaResult
		} else if parcela.Parcela_unario().IDENT() != nil {
			ident := parcela.Parcela_unario().IDENT()

			result := v.VerifyFuncCall(ident, parcela.Parcela_unario().AllExpressao())

			parcelaResult = append(parcelaResult, result...)

			funcSymbol := v.GetFuncVarSymbol(ident.GetText())

			return MapStringToType(funcSymbol.ReturnType), parcelaResult
		} else if parcela.Parcela_unario().IDENT() == nil && parcela.Parcela_unario().ABREPAR() != nil {
			return v.VerifyExpression(parcela.Parcela_unario().Expressao(0))
		}
	}

	if parcela.Parcela_nao_unario() != nil {
		if parcela.Parcela_nao_unario().ENDERECO() != nil {
			return symboltable.ENDERECO, parcelaResult
		}

		if parcela.Parcela_nao_unario().CADEIA() != nil {
			return symboltable.LITERAL, parcelaResult
		}
	}

	return symboltable.INVALIDO, parcelaResult
}

func (v *AlgumaVisitor) GetIdentifierType(identifier parser.IIdentificadorContext) symboltable.Type {

	scopeToVerify := v.Scopes.CurrentScope()
	var identName string

	for i, ident := range identifier.AllIDENT() {
		identName = ident.GetText()
		regVarSymbol := v.GetRegVarSymbol(identName)

		if regVarSymbol != nil && identifier.IDENT(i+1) != nil {
			scopeToVerify = regVarSymbol.InnerTable
		}
	}

	return scopeToVerify.GetType(identName)
}

func (v *AlgumaVisitor) GetIdentType(ident antlr.TerminalNode) symboltable.Type {

	var idx int

	for i, scope := range v.Scopes.Stack {
		if scope == v.Scopes.CurrentScope() {
			idx = i
			break
		}
	}

	for _, scope := range v.Scopes.Stack[:idx+1] {
		if scope.Exists(ident.GetText()) {
			return scope.GetType(ident.GetText())
		}
	}

	return symboltable.INVALIDO
}

func (v *AlgumaVisitor) VerifyFuncCall(ident antlr.TerminalNode, expressoes []parser.IExpressaoContext) []string {

	funcCallResult := make([]string, 0)

	exists := v.VerifyIdentOnValidScopes(ident)

	if !exists {
		funcCallResult = append(funcCallResult,
			semanticError(ident.GetSymbol(), fmt.Sprintf("identificador %v nao declarado", ident.GetText())),
		)
	}

	funcSymbol := v.GetFuncVarSymbol(ident.GetText())
	paramTypes := make([]symboltable.Type, 0)

	for _, paramType := range funcSymbol.InnerTable.Table {
		paramTypes = append(paramTypes, paramType.SymbolType)
	}

	if len(expressoes) != len(paramTypes) {
		funcCallResult = append(funcCallResult,
			semanticError(ident.GetSymbol(), fmt.Sprintf("incompatibilidade de parametros na chamada de %v", ident.GetText())),
		)

		return funcCallResult
	}

	for i, expressao := range expressoes {
		exprType, result := v.VerifyExpression(expressao)

		funcCallResult = append(funcCallResult, result...)

		if len(result) > 0 {
			return funcCallResult
		}

		if exprType != paramTypes[i] {
			funcCallResult = append(funcCallResult,
				semanticError(ident.GetSymbol(), fmt.Sprintf("incompatibilidade de parametros na chamada de %v", ident.GetText())),
			)

			return funcCallResult
		}
	}

	return funcCallResult
}
