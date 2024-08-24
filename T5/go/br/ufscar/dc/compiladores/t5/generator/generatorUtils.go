package generator

import (
	"log"

	parser "github.com/DanielKenichi/construcao-compiladores/T5/antlr4/br/ufscar/dc/compiladores/t5/parser/Alguma"
	"github.com/DanielKenichi/construcao-compiladores/T5/go/br/ufscar/dc/compiladores/t5/symboltable"
	"github.com/DanielKenichi/construcao-compiladores/T5/go/br/ufscar/dc/compiladores/t5/visitor"
)

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
			if termo.Op2(i-1) != nil {
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
			identifier := parcela.Parcela_unario().Identificador().GetText()
			parcelaResult = append(parcelaResult, identifier)
		} else if parcela.Parcela_unario().NUM_INT() != nil {
			result := parcela.Parcela_unario().NUM_INT().GetText()
			parcelaResult = append(parcelaResult, result)
		} else if parcela.Parcela_unario().NUM_REAL() != nil {
			result := parcela.Parcela_unario().NUM_REAL().GetText()
			parcelaResult = append(parcelaResult, result)
		} else if parcela.Parcela_unario().IDENT() == nil && parcela.Parcela_unario().ABREPAR() != nil {
			return g.VerifyExpression(parcela.Parcela_unario().Expressao(0))
		} else if parcela.Parcela_unario().IDENT() != nil {
			ident := parcela.Parcela_unario().IDENT()

			parcelaResult = append(parcelaResult, ident.GetText())
			parcelaResult = append(parcelaResult, "(")

			for i, expressao := range parcela.Parcela_unario().AllExpressao() {
				parcelaResult = append(parcelaResult, g.VerifyExpression(expressao)...)

				if parcela.Parcela_unario().Expressao(i+1) != nil {
					parcelaResult = append(parcelaResult, ", ")
				}

			}

			parcelaResult = append(parcelaResult, ")")
		}
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
	scopeToVerify := g.Visitor.Scopes.CurrentScope()
	var identName string

	for i, ident := range identifier.AllIDENT() {
		identName = ident.GetText()
		regVarSymbol := g.Visitor.GetRegVarSymbol(identName)

		if regVarSymbol != nil && identifier.IDENT(i+1) != nil {
			scopeToVerify = regVarSymbol.InnerTable
		}
	}

	return scopeToVerify.GetType(identName)
}

func (g *AlgumaGenerator) GetAllScopesType(ident string) symboltable.Type {
	var idx int

	for i, scope := range g.Visitor.Scopes.Stack {
		if scope == g.Visitor.Scopes.CurrentScope() {
			idx = i
			break
		}
	}

	for _, scope := range g.Visitor.Scopes.Stack[:idx+1] {
		if g.checkFunction != "" {
			if scope.Exists(g.checkFunction) {
				functionSymbol := scope.GetSymbol(g.checkFunction)

				if functionSymbol.InnerTable.Exists(ident) {
					return functionSymbol.InnerTable.GetType(ident)
				}
			} else {
				continue
			}
		}

		if scope.Exists(ident) {
			return scope.GetType(ident)
		}
	}

	return symboltable.INVALIDO
}

func (g *AlgumaGenerator) GetFunctionReturnType(function string) symboltable.Type {
	functionSymbol := g.Visitor.GetFuncVarSymbol(function)

	return visitor.MapStringToType(functionSymbol.ReturnType)
}

func (g *AlgumaGenerator) GetParcelaUnarioType(ctx parser.IParcela_unarioContext) symboltable.Type {
	expectedType := symboltable.INVALIDO

	log.Printf("ParcelaUnario %v", ctx.GetText())

	if ctx.Identificador() != nil {
		nome := ctx.Identificador().IDENT(0).GetText()
		log.Printf("Identifier %v", nome)

		expectedType = g.GetAllScopesType(nome)

		if len(ctx.Identificador().AllPONTO()) > 0 {
			expectedType = g.GetIdentifierType(ctx.Identificador())
		}
	}

	if ctx.IDENT() != nil {
		return g.GetFunctionReturnType(ctx.IDENT().GetText())
	}

	if ctx.PONTEIRO() != nil {
		return symboltable.PONTEIRO
	}
	if ctx.NUM_INT() != nil {
		return symboltable.INTEIRO
	}
	if ctx.NUM_REAL() != nil {
		return symboltable.REAL
	}

	if ctx.ABREPAR() != nil {
		return g.GetExpressaoType(ctx.Expressao(0))
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
		parcelaUnariotype := g.GetParcelaUnarioType(ctx.Parcela_unario())
		log.Printf("ParcelaUnarioType %v", parcelaUnariotype)
		return parcelaUnariotype
	} else if (ctx.Parcela_nao_unario()) != nil {
		parcelaNUnarioType := g.GetParcelaNaoUnarioType(ctx.Parcela_nao_unario())
		log.Printf("ParcelaNaoUnarioType %v", parcelaNUnarioType)
		return parcelaNUnarioType
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

	log.Printf("ParcelaType %v", expectedType)

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
		expType := g.GetExpRelacionalType(ctx.Exp_relacional())

		return expType
	}

	return symboltable.INVALIDO
}

func (g *AlgumaGenerator) GetFatorLogicoType(ctx parser.IFator_logicoContext) symboltable.Type {
	if ctx.Parcela_logica() != nil {
		parcelaType := g.GetParcelaLogicaType(ctx.Parcela_logica())

		return parcelaType
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

func (g *AlgumaGenerator) MapParamToTypeC(tipo parser.ITipo_basicoContext) string {
	originalType := tipo
	typeC := ""

	if originalType.INTEIRO() != nil {
		typeC = "int"
	} else if originalType.REAL() != nil {
		typeC = "float"
	} else if originalType.LITERAL() != nil {
		typeC = "char*"
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
