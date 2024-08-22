package generator

import (
	parser "github.com/DanielKenichi/construcao-compiladores/T5/antlr4/br/ufscar/dc/compiladores/t5/parser/Alguma"
	"github.com/DanielKenichi/construcao-compiladores/T5/go/br/ufscar/dc/compiladores/t5/symboltable"
	"github.com/antlr4-go/antlr/v4"
)

func (g *AlgumaGenerator) AddVarToSymbolTable(variavel parser.IVariavelContext, table *symboltable.SymbolTable) []string {

	addSymbolTableResult := make([]string, 0)

	var varType symboltable.Type = symboltable.INVALIDO

	// if variavel.Tipo().Registro() != nil {
	// 	varType = symboltable.REGISTRO_VAR
	// } else if variavel.Tipo().Tipo_variavel().PONTEIRO() != nil {
	// 	varType = symboltable.PONTEIRO
	// } else if variavel.Tipo().Tipo_variavel().Tipo_basico() != nil {
	// 	basicType := variavel.Tipo().Tipo_variavel().Tipo_basico()

	// 	varType = MapBasicTypeToSymbolType(basicType)
	// } else { //Tipo variavel eh um ident de um registro
	// 	varType = symboltable.REGISTRO_VAR
	// }

	if variavel.Tipo().Tipo_variavel().Tipo_basico() != nil {
		basicType := variavel.Tipo().Tipo_variavel().Tipo_basico()
		varType = MapBasicTypeToSymbolType(basicType)
	}

	for _, idents := range variavel.AllIdentificador() {
		for _, identName := range idents.AllIDENT() {
			result := g.AddIdentifierToSymbolTable(identName, varType, table)

			addSymbolTableResult = append(addSymbolTableResult, result...)

			// if varType == symboltable.REGISTRO_VAR {
			// 	if variavel.Tipo().Registro() != nil {
			// 		regTable := table.GetSymbol(identName.GetText()).InnerTable

			// 		result = g.AddRegInnerVarsToSymbolTable(variavel.Tipo().Registro(), regTable)

			// 		addSymbolTableResult = append(addSymbolTableResult, result...)
			// 	} else {
			// 		regSymbol := g.GetRegTypeSymbol(variavel.Tipo().Tipo_variavel().IDENT().GetText())

			// 		table.AddInnerTableToRegVar(identName.GetText(), *regSymbol.InnerTable)
			// 	}
			// }
		}
	}

	return addSymbolTableResult
}

func (g *AlgumaGenerator) AddConstToSymbolTable(identifier antlr.TerminalNode, basicType parser.ITipo_basicoContext) []string {
	varType := MapBasicTypeToSymbolType(basicType)

	result := g.AddIdentifierToSymbolTable(identifier, varType, g.Scopes.CurrentScope())

	return result
}

func (g *AlgumaGenerator) VerifyExpression(expression parser.IExpressaoContext) []string {
	result := make([]string, 0)
	for index, termoLogico := range expression.AllTermo_logico() {
		termoLogicoResult := g.VerifyTermoLogico(termoLogico)

		if index > 0 {
			if expression.Op_logico_1(index-1) != nil {
				result = append(result, (" || ")) // OU
			}
		}

		result = append(result, termoLogicoResult...)
	}

	return result
}

func (g *AlgumaGenerator) VerifyTermoLogico(termoLogico parser.ITermo_logicoContext) []string {
	parcelaLogicaResult := make([]string, 0)
	result := make([]string, 0)

	if termoLogico.Fator_logico(0).NAO() != nil {
		result = append(result, ("!("))
	}

	result = append(result, g.VerifyParcelaLogica(termoLogico.Fator_logico(0).Parcela_logica())...)

	if termoLogico.Fator_logico(0).NAO() != nil {
		result = append(result, ")")
	}
	parcelaLogicaResult = append(parcelaLogicaResult, result...)

	for index, fatorLogico := range termoLogico.AllFator_logico()[1:] {
		if termoLogico.Op_logico_2(index) != nil {
			result = []string{(" && ")} // E
		}
		if termoLogico.Fator_logico(index).NAO() != nil {
			result = append(result, ("!("))
		}

		result = append(result, g.VerifyParcelaLogica(fatorLogico.Parcela_logica())...)

		if termoLogico.Fator_logico(index).NAO() != nil {
			result = append(result, (")"))
		}

		parcelaLogicaResult = append(parcelaLogicaResult, result...)
	}

	return parcelaLogicaResult
}

func (g *AlgumaGenerator) VerifyParcelaLogica(parcelaLogica parser.IParcela_logicaContext) []string {
	result := make([]string, 0)

	if parcelaLogica.VERDADEIRO() != nil {
		result = []string{"1"}
		return result
	}
	if parcelaLogica.FALSO() != nil {
		result = []string{"0"}
		return result
	}

	expArit := parcelaLogica.Exp_relacional().Exp_aritmetica(0)
	expAritResult := g.VerifyExpressaoAritmetica(expArit)
	result = append(result, expAritResult...)

	if parcelaLogica.Exp_relacional().Op_relacional() != nil {
		opRelResult := g.VerifyOpRelacional(parcelaLogica.Exp_relacional().Op_relacional())
		result = append(result, opRelResult...)

		expAritResult = g.VerifyExpressaoAritmetica(parcelaLogica.Exp_relacional().Exp_aritmetica(1))
		result = append(result, expAritResult...)
	}

	return result
}

func (g *AlgumaGenerator) VerifyExpressaoAritmetica(expAritmetica parser.IExp_aritmeticaContext) []string {
	result := make([]string, 0)

	for i, termo := range expAritmetica.AllTermo() {
		termoResult := g.VerifyTermo(termo)

		if i > 0 {
			if expAritmetica.Op1(i-1) != nil {
				result = append(result, " ", expAritmetica.Op1(i-1).GetText(), " ")
			}
		}

		result = append(result, termoResult...)
	}

	return result
}

func (g *AlgumaGenerator) VerifyTermo(termo parser.ITermoContext) []string {
	result := make([]string, 0)

	for i, fator := range termo.AllFator() {
		fatorResult := g.VerifyFator(fator)

		if i > 0 {
			if termo.Op2(i) != nil {
				result = append(result, " ", termo.Op2(i-1).GetText(), " ")
			}
		}

		result = append(result, fatorResult...)
	}

	return result
}

func (g *AlgumaGenerator) VerifyFator(fator parser.IFatorContext) []string {
	result := make([]string, 0)

	for i, parcela := range fator.AllParcela() {
		parcelaResult := g.VerifyParcela(parcela)

		if i > 0 {
			if fator.Op3(i-1) != nil {
				result = append(result, " ", fator.Op3(i).GetText(), " ")
			}
		}

		result = append(result, parcelaResult...)
	}

	return result
}

func (g *AlgumaGenerator) VerifyParcela(parcela parser.IParcelaContext) []string {
	parcelaResult := make([]string, 0)

	if parcela.Op_unario() != nil {
		parcelaResult = append(parcelaResult, parcela.Op_unario().GetText())
	}

	if parcela.Parcela_unario() != nil {
		if parcela.Parcela_unario().Identificador() != nil {
			// if parcela.Parcela_unario().PONTEIRO() != nil {
			// 	parcelaResult = append(parcelaResult, "*")
			// }
			identifier := parcela.Parcela_unario().Identificador().GetText()
			parcelaResult = append(parcelaResult, identifier)

			// if parcela.Parcela_unario().Identificador() != nil {
			// 	identifier := parcela.Parcela_unario().Identificador()

			// 	result := []string{identifier.GetText()}
			// 	parcelaResult = append(parcelaResult, result...)

			// 	return g.GetIdentifierType(identifier), parcelaResult
		} else if parcela.Parcela_unario().NUM_INT() != nil {
			result := parcela.Parcela_unario().NUM_INT().GetText()
			parcelaResult = append(parcelaResult, result)
		} else if parcela.Parcela_unario().NUM_REAL() != nil {
			result := parcela.Parcela_unario().NUM_REAL().GetText()
			parcelaResult = append(parcelaResult, result)
		}

		// } else if parcela.Parcela_unario().NUM_INT() != nil {
		// 	return symboltable.INTEIRO, parcelaResult
		// } else if parcela.Parcela_unario().NUM_REAL() != nil {
		// 	return symboltable.REAL, parcelaResult
		// } else if parcela.Parcela_unario().IDENT() != nil {
		// 	ident := parcela.Parcela_unario().IDENT()

		// 	result := g.VerifyFuncCall(ident, parcela.Parcela_unario().AllExpressao())

		// 	parcelaResult = append(parcelaResult, result...)

		// 	funcSymbol := g.GetFuncVarSymbol(ident.GetText())

		// 	return MapStringToType(funcSymbol.ReturnType), parcelaResult
		// } else if parcela.Parcela_unario().IDENT() == nil && parcela.Parcela_unario().ABREPAR() != nil {
		// 	return g.VerifyExpression(parcela.Parcela_unario().Expressao(0))
		// }
	}

	if parcela.Parcela_nao_unario() != nil {
		parcelaResult = append(parcelaResult, parcela.Parcela_nao_unario().GetText())
	}

	return parcelaResult
}

func (g *AlgumaGenerator) VerifyOpRelacional(opRel parser.IOp_relacionalContext) []string {
	result := make([]string, 0)
	opC := g.MapOpRelToOpRelC(opRel)

	if opRel != nil {
		result = append(result, opC)
	}

	return result
}

func (g *AlgumaGenerator) GetIdentifierType(identifier parser.IIdentificadorContext) symboltable.Type {
	scopeToVerify := g.Scopes.CurrentScope()
	var identName string

	for _, ident := range identifier.AllIDENT() {
		identName = ident.GetText()
		// regVarSymbol := g.GetRegVarSymbol(identName)

		// if regVarSymbol != nil && identifier.IDENT(i+1) != nil {
		// 	scopeToVerify = regVarSymbol.InnerTable
		// }
	}

	return scopeToVerify.GetType(identName)
}

func (g *AlgumaGenerator) GetAllScopesType(ident string) symboltable.Type {
	var idx int

	for i, scope := range g.Scopes.Stack {
		if scope == g.Scopes.CurrentScope() {
			idx = i
			break
		}
	}

	for _, scope := range g.Scopes.Stack[:idx+1] {
		if scope.Exists(ident) {
			return scope.GetType(ident)
		}
	}

	return symboltable.INVALIDO
}

func (g *AlgumaGenerator) GetParcelaUnarioType(ctx parser.IParcela_unarioContext) symboltable.Type {
	expectedType := symboltable.INVALIDO

	if ctx.Identificador() != nil {
		nome := ctx.Identificador().IDENT(0).GetText()

		if len(ctx.Identificador().AllPONTO()) > 0 {
			// TODO: registros
		}

		return g.GetAllScopesType(nome)
	}
	if ctx.PONTEIRO() != nil {
		expectedType = symboltable.PONTEIRO
	}
	if ctx.NUM_INT() != nil {
		expectedType = symboltable.INTEIRO
	}
	if ctx.NUM_REAL() != nil {
		expectedType = symboltable.REAL
	}

	return expectedType
}

func (g *AlgumaGenerator) GetParcelaNaoUnarioType(ctx parser.IParcela_nao_unarioContext) symboltable.Type {
	expectedType := symboltable.INVALIDO

	if ctx.Identificador() != nil {
		expectedType = g.GetIdentifierType(ctx.Identificador())
	}
	if ctx.ENDERECO() != nil {
		expectedType = symboltable.ENDERECO
	}
	if ctx.CADEIA() != nil {
		expectedType = symboltable.LITERAL
	}

	return expectedType
}

func (g *AlgumaGenerator) GetParcelaType(ctx parser.IParcelaContext) symboltable.Type {
	if ctx.Parcela_unario() != nil {
		return g.GetParcelaUnarioType(ctx.Parcela_unario())
	} else if (ctx.Parcela_nao_unario()) != nil {
		return g.GetParcelaNaoUnarioType(ctx.Parcela_nao_unario())
	}

	return symboltable.INVALIDO
}

func (g *AlgumaGenerator) GetFatorType(ctx parser.IFatorContext) symboltable.Type {
	expectedType := g.GetParcelaType(ctx.Parcela(0))

	if expectedType != symboltable.INVALIDO {
		for _, parcela := range ctx.AllParcela() {
			actualType := g.GetParcelaType(parcela)

			if !g.TypesCompatible(expectedType, actualType) {
				return symboltable.INVALIDO
			}
		}
	}

	return expectedType
}

func (g *AlgumaGenerator) GetTermoType(ctx parser.ITermoContext) symboltable.Type {
	expectedType := g.GetFatorType(ctx.Fator(0))

	if expectedType != symboltable.INVALIDO {
		for _, fator := range ctx.AllFator() {
			actualType := g.GetFatorType(fator)

			if !g.TypesCompatible(expectedType, actualType) {
				return symboltable.INVALIDO
			}
		}
	}

	return expectedType
}

func (g *AlgumaGenerator) GetExpAritmeticaType(ctx parser.IExp_aritmeticaContext) symboltable.Type {
	expectedType := g.GetTermoType(ctx.Termo(0))

	if expectedType != symboltable.INVALIDO {
		for _, termo := range ctx.AllTermo() {
			actualType := g.GetTermoType(termo)

			if !g.TypesCompatible(expectedType, actualType) {
				return symboltable.INVALIDO
			}
		}
	}

	return expectedType
}

func (g *AlgumaGenerator) GetExpRelacionalType(ctx parser.IExp_relacionalContext) symboltable.Type {
	expectedType := g.GetExpAritmeticaType(ctx.Exp_aritmetica(0))

	if expectedType != symboltable.INVALIDO {
		for _, expArit := range ctx.AllExp_aritmetica() {
			actualType := g.GetExpAritmeticaType(expArit)

			if !g.TypesCompatible(expectedType, actualType) {
				return symboltable.INVALIDO
			}
		}

		if ctx.Op_relacional() != nil {
			return symboltable.LOGICO
		}
	}

	return expectedType
}

func (g *AlgumaGenerator) GetParcelaLogicaType(ctx parser.IParcela_logicaContext) symboltable.Type {
	if ctx.Exp_relacional() != nil {
		return g.GetExpRelacionalType(ctx.Exp_relacional())
	}

	return symboltable.INVALIDO
}

func (g *AlgumaGenerator) GetFatorLogicoType(ctx parser.IFator_logicoContext) symboltable.Type {
	if ctx.Parcela_logica() != nil {
		return g.GetParcelaLogicaType(ctx.Parcela_logica())
	}

	return symboltable.INVALIDO
}

func (g *AlgumaGenerator) GetTermoLogicoType(ctx parser.ITermo_logicoContext) symboltable.Type {
	expectedType := g.GetFatorLogicoType(ctx.Fator_logico(0))

	if expectedType != symboltable.INVALIDO {
		for _, fatorLog := range ctx.AllFator_logico() {
			actualType := g.GetFatorLogicoType(fatorLog)

			if !g.TypesCompatible(expectedType, actualType) {
				return symboltable.INVALIDO
			}
		}

		if len(ctx.AllOp_logico_2()) > 0 {
			return symboltable.LOGICO
		}
	}

	return expectedType
}

func (g *AlgumaGenerator) GetExpressaoType(ctx parser.IExpressaoContext) symboltable.Type {
	expectedType := g.GetTermoLogicoType(ctx.Termo_logico(0))

	if expectedType != symboltable.INVALIDO {
		for _, termoLog := range ctx.AllTermo_logico() {
			actualType := g.GetTermoLogicoType(termoLog)

			if !g.TypesCompatible(expectedType, actualType) {
				return symboltable.INVALIDO
			}
		}

		if len(ctx.AllOp_logico_1()) > 0 {
			return symboltable.LOGICO
		}
	}

	return expectedType
}

func (g *AlgumaGenerator) TypesCompatible(t1, t2 symboltable.Type) bool {
	if t1 == symboltable.INVALIDO || t2 == symboltable.INVALIDO {
		return false
	}

	return t1 == t2 ||
		(t1 == symboltable.INTEIRO && t2 == symboltable.REAL) ||
		(t1 == symboltable.REAL && t2 == symboltable.INTEIRO) ||
		(t1 == symboltable.PONTEIRO && t2 == symboltable.ENDERECO) ||
		(t1 == symboltable.ENDERECO && t2 == symboltable.PONTEIRO)

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

func (g *AlgumaGenerator) AddIdentifierToSymbolTable(identifier antlr.TerminalNode, varType symboltable.Type, table *symboltable.SymbolTable) []string {
	result := make([]string, 0)

	table.AddSymbol(identifier.GetText(), varType)

	return result
}

func (g *AlgumaGenerator) MapTypeToTypeC(tipo parser.ITipo_basicoContext) string {
	originalType := tipo
	typeC := ""

	if originalType.INTEIRO() != nil {
		typeC = "int"
	} else if originalType.REAL() != nil {
		typeC = "float"
	} else if originalType.LITERAL() != nil {
		typeC = "char"
	} else if originalType.LOGICO() != nil {
		typeC = "int"
	}

	return typeC
}

func (g *AlgumaGenerator) MapTypeToOperatorC(tipo symboltable.Type) string {
	originalType := tipo
	operatorC := ""

	if originalType == symboltable.INTEIRO {
		operatorC = "%d"
	} else if originalType == symboltable.REAL {
		operatorC = "%f"
	} else if originalType == symboltable.LITERAL {
		operatorC = "%s"
	} else if originalType == symboltable.LOGICO {
		operatorC = "%d"
	}

	return operatorC
}

func (g *AlgumaGenerator) MapOpRelToOpRelC(parser parser.IOp_relacionalContext) string {
	originalOp := parser
	opC := ""

	if originalOp.IGUAL() != nil {
		opC = " == "
	} else if originalOp.DIFERENTE() != nil {
		opC = " != "
	} else if originalOp.MAIOR() != nil {
		opC = " > "
	} else if originalOp.MAIORIGUAL() != nil {
		opC = " >= "
	} else if originalOp.MENOR() != nil {
		opC = " < "
	} else if originalOp.MENORIGUAL() != nil {
		opC = " <= "
	}

	return opC
}