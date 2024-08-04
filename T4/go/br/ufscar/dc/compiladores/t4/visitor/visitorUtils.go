package visitor

import (
	"fmt"

	parser "github.com/DanielKenichi/construcao-compiladores/T4/antlr4/br/ufscar/dc/compiladores/t4/parser/Alguma"
	"github.com/DanielKenichi/construcao-compiladores/T4/go/br/ufscar/dc/compiladores/t4/symboltable"
	"github.com/antlr4-go/antlr/v4"
)

func (v *AlgumaVisitor) AddVarToSymbolTable(variavel parser.IVariavelContext) []string {

	addSymbolTableResult := make([]string, 0)

	var varType symboltable.Type = symboltable.INVALIDO
	var tipoType symboltable.Type = symboltable.INVALIDO

	// Se o tipo estiver no ctx como Registro, marca a variável como registro_var
	for _, scope := range v.Scopes.Stack {
		if scope.Exists(variavel.Tipo().GetText()) {
			tipoType = scope.GetType(variavel.Tipo().GetText())
		}
	}
	if tipoType == symboltable.REGISTRO {
		varType = symboltable.REGISTRO_VAR
		// v.AddRegVarToSymbolTable(variavel.Tipo().Registro())
	} else if variavel.Tipo().Registro() != nil {
		varType = symboltable.REGISTRO
	} else if variavel.Tipo().Tipo_variavel().PONTEIRO() != nil {
		varType = symboltable.PONTEIRO
	} else {
		basicType := variavel.Tipo().Tipo_variavel().Tipo_basico()

		varType = MapBasicTypeToSymbolType(basicType)
	}

	for _, idents := range variavel.AllIdentificador() {
		for _, identName := range idents.AllIDENT() {
			// log.Printf("Adicionando variavel %v do tipo %v", identName, varType)
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

func (v *AlgumaVisitor) AddFuncToSymbolTable(identifier antlr.TerminalNode) []string {

	varType := symboltable.FUNCAO

	result := v.AddIdentifierToSymbolTable(identifier, varType)

	return result

}

func (v *AlgumaVisitor) AddRegToSymbolTable(identifier antlr.TerminalNode) []string {

	varType := symboltable.REGISTRO

	result := v.AddIdentifierToSymbolTable(identifier, varType)

	return result

}

func (v *AlgumaVisitor) AddRegVarToSymbolTable(ctx parser.IRegistroContext) []string {

	result := make([]string, 0)
	// log.Printf("Adicionando registro %v", ctx.GetText())
	// TODO: Declarar cada `struct.var: tipo` na tabela de símbolos com o seu tipo dentro da struct

	// for _, variavel := range ctx.AllVariavel() {
	// 	log.Printf("Adicionando variavel do registro: %v", variavel.GetText())
	// 	// result = v.AddVarToSymbolTable(variavel)

	// 	if len(result) > 0 {
	// 		return result
	// 	}
	// }

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

func (v *AlgumaVisitor) AddIdentifierToSymbolTable(identifier antlr.TerminalNode, varType symboltable.Type) []string {
	result := make([]string, 0)

	for _, scope := range v.Scopes.Stack {
		if scope.Exists(identifier.GetText()) {
			result = append(result,
				semanticError(identifier.GetSymbol(), fmt.Sprintf("identificador %v ja declarado anteriormente", identifier.GetText())),
			)
		}
	}

	v.Scopes.CurrentScope().AddSymbol(identifier.GetText(), varType)

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
	result := make([]string, 0)

	identifier := atribuicao.Identificador().IDENT(0)

	if !v.Scopes.CurrentScope().Exists(identifier.GetText()) {

		result = append(result,
			semanticError(identifier.GetSymbol(), fmt.Sprintf("identificador %v nao declarado", identifier.GetText())),
		)

		return result
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

	attrType, result = v.VerifyExpression(atribuicao.Expressao())

	/*Se tiver um identificador nao declarado na expressao
	o resultado vai ter pelo menos um erro semantico.
	Dessa forma, retornar ele antes para nao ter tambem erro de atribuicao
	incompativel, uma vez que um identificador nao declarado retorna
	tipo invalido para a expressao */

	if len(result) > 0 {
		return result
	}

	identSufix := ""
	if atribuicao.Identificador().ABRECHAVE(0) != nil {
		identSufix = atribuicao.Identificador().ABRECHAVE(0).GetText()
		identSufix += atribuicao.Identificador().Exp_aritmetica(0).GetText()
		identSufix += atribuicao.Identificador().FECHACHAVE(0).GetText()
	}

	if atribuicao.Identificador().PONTO(0) != nil {
		identSufix = atribuicao.Identificador().PONTO(0).GetText()
		identSufix += atribuicao.Identificador().IDENT(1).GetText()

	}

	identType := v.Scopes.CurrentScope().GetType(identifier.GetText())

	// log.Printf("%v (%v) = %v (%v)", identifier, identType, atribuicao.Expressao().GetText(), attrType)
	// log.Printf("%v = %v", identifier.GetText(), atribuicao.Expressao().GetText())
	// log.Printf("%v = %v", identType, attrType)

	// if atribuicao.Identificador().PONTO(0) != nil {
	// 	log.Printf("Atribuicao para registro - %v%v%v",
	// 		atribuicao.Identificador().IDENT(0).GetText(),
	// 		atribuicao.Identificador().PONTO(0).GetText(),
	// 		atribuicao.Identificador().IDENT(1).GetText())
	// }

	if !IsTypesCompatible(identType, attrType) {
		identPrefix := ""

		if identType == symboltable.PONTEIRO {
			identPrefix = atribuicao.PONTEIRO().GetText()
		}

		result = append(result,
			semanticError(identifier.GetSymbol(), fmt.Sprintf("atribuicao nao compativel para %v%v%v", identPrefix, identifier.GetText(), identSufix)),
		)
	}

	return result
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
	result := make([]string, 0)

	if parcela.Parcela_unario() != nil {

		if parcela.Parcela_unario().Identificador() != nil {

			identifier := parcela.Parcela_unario().Identificador().IDENT(0)

			if !v.Scopes.CurrentScope().Exists(identifier.GetText()) {

				result = append(result,
					// TODO: Identificador não declarado, mas agora envolvendo também ponteiros, registros, funções
					semanticError(identifier.GetSymbol(), fmt.Sprintf("identificador %v nao declarado", identifier.GetText())),
				)

				return symboltable.INVALIDO, result
			}

			return v.Scopes.CurrentScope().GetType(identifier.GetText()), result

		} else if parcela.Parcela_unario().NUM_INT() != nil {
			return symboltable.INTEIRO, result
		} else if parcela.Parcela_unario().NUM_REAL() != nil {
			return symboltable.REAL, result
		} else if parcela.Parcela_unario().IDENT() == nil && parcela.Parcela_unario().ABREPAR() != nil {
			return v.VerifyExpression(parcela.Parcela_unario().Expressao(0))
		}
	}

	if parcela.Parcela_nao_unario() != nil {
		if parcela.Parcela_nao_unario().ENDERECO() != nil {
			return symboltable.ENDERECO, result
		}

		if parcela.Parcela_nao_unario().CADEIA() != nil {
			return symboltable.LITERAL, result
		}
	}

	return symboltable.INVALIDO, result
}
