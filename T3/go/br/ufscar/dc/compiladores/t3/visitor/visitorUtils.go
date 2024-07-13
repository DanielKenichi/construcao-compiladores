package visitor

import (
	"fmt"
	"log"

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

	if v.Scopes.CurrentScope().Exists(identifier.GetText()) {
		result = append(result,
			semanticError(identifier.GetSymbol(), fmt.Sprintf("identificador %v ja declarado anteriormente", identifier.GetText())),
		)
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
		(t1 == symboltable.REAL && t2 == symboltable.INTEIRO)
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

	if atribuicao.PONTEIRO() != nil {
		attrType = symboltable.PONTEIRO
	} else {
		attrType, result = v.VerifyExpression(atribuicao.Expressao())
	}

	if len(result) > 0 {
		return result
	}

	identType := v.Scopes.CurrentScope().GetType(identifier.GetText())

	if !IsTypesCompatible(identType, attrType) {
		result = append(result,
			semanticError(identifier.GetSymbol(), fmt.Sprintf("atribuicao nao compativel para %v", identifier.GetText())),
		)
	}

	return result
}

func (v *AlgumaVisitor) VerifyExpression(expression parser.IExpressaoContext) (symboltable.Type, []string) {
	var expType symboltable.Type
	result := make([]string, 0)
	for _, termoLogico := range expression.AllTermo_logico() {
		for _, fatorLogico := range termoLogico.AllFator_logico() {
			expType, result = v.VerifyParcelaLogica(fatorLogico.Parcela_logica())

			if expType == symboltable.INVALIDO {
				log.Printf("invalid parcel: %v", result)

				return symboltable.INVALIDO, result
			}
		}
	}

	return expType, result
}

func (v *AlgumaVisitor) VerifyParcelaLogica(parcelaLogica parser.IParcela_logicaContext) (symboltable.Type, []string) {

	result := make([]string, 0)
	var parcelaType symboltable.Type

	if parcelaLogica.VERDADEIRO() != nil || parcelaLogica.FALSO() != nil {
		return symboltable.LOGICO, result
	}

	for _, expArit := range parcelaLogica.Exp_relacional().AllExp_aritmetica() {
		currentType, expAritResult := v.VerifyExpressaoAritmetica(expArit)

		result = append(result, expAritResult...)

		if currentType == symboltable.INVALIDO {
			log.Printf("invalid exparit: %v", result)

			return symboltable.INVALIDO, result
		}

		parcelaType = currentType
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
			log.Printf("invalid termo: %v", result)
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
			log.Printf("invalid fator: %v", result)

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
			log.Print("Verify parcela ident")

			identifier := parcela.Parcela_unario().Identificador().IDENT(0)

			if !v.Scopes.CurrentScope().Exists(identifier.GetText()) {

				result = append(result,
					semanticError(identifier.GetSymbol(), fmt.Sprintf("identificador %v nao declarado", identifier.GetText())),
				)

				log.Printf("invalid ident %v %v", identifier.GetText(), result)

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
		return symboltable.PONTEIRO, result
	}

	return symboltable.INVALIDO, result
}
