package visitor

import (
	"fmt"

	parser "github.com/DanielKenichi/construcao-compiladores/T3/antlr4/br/ufscar/dc/compiladores/t3/parser/Alguma"
	"github.com/DanielKenichi/construcao-compiladores/T3/go/br/ufscar/dc/compiladores/t3/scope"
	"github.com/DanielKenichi/construcao-compiladores/T3/go/br/ufscar/dc/compiladores/t3/symboltable"
	"github.com/antlr4-go/antlr/v4"
)

type AlgumaVisitor struct {
	parser.BaseT3AlgumaVisitor
	Scopes *scope.Scope
}

func New() *AlgumaVisitor {
	return &AlgumaVisitor{
		Scopes: scope.New(),
	}
}

func semanticError(token antlr.Token, message string) string {
	return fmt.Sprintf("Linha %v: %s", token.GetLine(), message)
}

func (v *AlgumaVisitor) VisitPrograma(ctx parser.IProgramaContext) []string {

	var programaResult = make([]string, 0)

	result := v.VisitDeclaracoes(ctx.Declaracoes())

	programaResult = append(programaResult, result...)

	result = v.VisitCorpo(ctx.Corpo())

	programaResult = append(programaResult, result...)

	return programaResult
}

func (v *AlgumaVisitor) VisitDeclaracoes(ctx parser.IDeclaracoesContext) []string {

	v.Scopes.NewScope()

	declaracoesResult := make([]string, 0)

	result := v.VisitDeclaracoes_variaveis(ctx.AllDeclaracoes_variaveis())

	declaracoesResult = append(declaracoesResult, result...)

	result = v.VisitDeclaracoes_funcoes(ctx.AllDeclaracoes_funcoes())

	declaracoesResult = append(declaracoesResult, result...)

	return declaracoesResult
}

func (v *AlgumaVisitor) VisitCorpo(ctx parser.ICorpoContext) []string {
	corpoResult := make([]string, 0)

	result := v.VisitDeclaracoes_variaveis(ctx.AllDeclaracoes_variaveis())

	corpoResult = append(corpoResult, result...)

	result = v.VisitCmd(ctx.AllCmd())

	corpoResult = append(corpoResult, result...)

	return corpoResult
}

func (v *AlgumaVisitor) VisitCmd(ctxs []parser.ICmdContext) []string {
	cmdResult := make([]string, 0)

	for _, ctx := range ctxs {

		if ctx.CmdLeia() != nil {
			return cmdResult
		}

		if ctx.CmdEscreva() != nil {
			return cmdResult
		}

		if ctx.CmdSe() != nil {
			result := v.VisitCmdSe(ctx.CmdSe())

			cmdResult = append(cmdResult, result...)

			return cmdResult
		}

		if ctx.CmdCaso() != nil {
			result := v.VisitCmdCaso(ctx.CmdCaso())

			cmdResult = append(cmdResult, result...)

			return cmdResult
		}

		if ctx.CmdPara() != nil {
			result := v.VisitCmdPara(ctx.CmdPara())

			cmdResult = append(cmdResult, result...)

			return cmdResult
		}

		if ctx.CmdEnquanto() != nil {

		}
	}

	return cmdResult
}

func (v *AlgumaVisitor) VisitCmdSe(ctx parser.ICmdSeContext) []string {
	cmdSeResult := make([]string, 0)

	result := v.VisitExpressao(ctx.Expressao())

	cmdSeResult = append(cmdSeResult, result...)

	result = v.VisitCmd(ctx.AllCmd())

	cmdSeResult = append(cmdSeResult, result...)

	if ctx.SENAO() != nil {
		result = v.VisitCmd(ctx.AllCmd())

		cmdSeResult = append(cmdSeResult, result...)
	}

	return cmdSeResult
}

func (v *AlgumaVisitor) VisitCmdCaso(ctx parser.ICmdCasoContext) []string {
	cmdCasoResult := make([]string, 0)

	result := v.VisitExp_aritmetica(ctx.Exp_aritmetica())

	cmdCasoResult = append(cmdCasoResult, result...)

	result = v.VisitSelecao(ctx.Selecao())

	cmdCasoResult = append(cmdCasoResult, result...)

	return cmdCasoResult
}

func (v *AlgumaVisitor) VisitCmdPara(ctx parser.ICmdParaContext) []string {
	cmdParaResult := make([]string, 0)

	result := v.VisitExp_aritmetica(ctx.GetExp1())

	cmdParaResult = append(cmdParaResult, result...)

	result = v.VisitExp_aritmetica(ctx.GetExp2())

	cmdParaResult = append(cmdParaResult, result...)

	result = v.VisitCmd(ctx.AllCmd())

	cmdParaResult = append(cmdParaResult, result...)

	return cmdParaResult
}

func (v *AlgumaVisitor) VisitSelecao(ctx parser.ISelecaoContext) []string {
	selecaoResult := make([]string, 0)

	for _, itemSelecao := range ctx.AllItem_selecao() {
		result := v.VisitItem_selecao(itemSelecao)

		selecaoResult = append(selecaoResult, result...)
	}

	return selecaoResult
}

func (v *AlgumaVisitor) VisitItem_selecao(ctx parser.IItem_selecaoContext) []string {
	itemSelecaoResult := make([]string, 0)

	result := v.VisitConstantes(ctx.Constantes())

	itemSelecaoResult = append(itemSelecaoResult, result...)

	result = v.VisitCmd(ctx.AllCmd())

	itemSelecaoResult = append(itemSelecaoResult, result...)

	return itemSelecaoResult
}

func (v *AlgumaVisitor) VisitConstantes(ctx parser.IConstantesContext) []string {
	return make([]string, 0)
}

func (v *AlgumaVisitor) VisitDeclaracoes_variaveis(ctxs []parser.IDeclaracoes_variaveisContext) []string {

	variaveisResult := make([]string, 0)

	for _, ctx := range ctxs {

		if ctx.DECLARE() != nil {
			result := v.VisitVariavel(ctx.Variavel())

			variaveisResult = append(variaveisResult, result...)

			result = v.AddVarToSymbolTable(ctx.Variavel())

			variaveisResult = append(variaveisResult, result...)
			break
		}

		if ctx.CONSTANTE() != nil {

			result := v.VisitTipo_basico(ctx.Tipo_basico())

			variaveisResult = append(variaveisResult, result...)

			result = v.VisitValor_constante(ctx.Valor_constante())

			variaveisResult = append(variaveisResult, result...)

			result = v.AddConstToSymbolTable(ctx.IDENT(), ctx.Tipo_basico())

			variaveisResult = append(variaveisResult, result...)
			break
		}

		if ctx.TIPO() != nil {

			result := v.VisitRegistro(ctx.Registro())

			variaveisResult = append(variaveisResult, result...)

			result = v.AddIdentifierToSymbolTable(ctx.IDENT(), symboltable.REGISTRO)

			variaveisResult = append(variaveisResult, result...)

			break
		}
	}

	return variaveisResult
}

func (v *AlgumaVisitor) VisitDeclaracoes_funcoes(ctx []parser.IDeclaracoes_funcoesContext) []string {
	funcoesResult := make([]string, 0)

	return funcoesResult
}

/*
A partir desse no nessa branch da arvore nao seria necessario para os casos de testes do T3 aprofundar mais
daria apenas para retornar um slice de string vazio
mas fica a implementação apenas para exemplo caso seja necessario futuramente
*/

func (v *AlgumaVisitor) VisitVariavel(ctx parser.IVariavelContext) []string {
	variavelResult := make([]string, 0)

	result := v.VisitIdentificador(ctx.Identificador(0))

	variavelResult = append(variavelResult, result...)

	aux := 0
	for ctx.VIRGULA(aux) != nil {
		aux++

		result = v.VisitIdentificador(ctx.Identificador(aux))

		variavelResult = append(variavelResult, result...)
	}

	result = v.VisitTipo(ctx.Tipo())

	variavelResult = append(variavelResult, result...)

	return variavelResult
}

func (v *AlgumaVisitor) VisitExpressao(ctx parser.IExpressaoContext) []string {
	expressaoResult := make([]string, 0)

	return expressaoResult
}

func (v *AlgumaVisitor) VisitTipo_basico(ctx parser.ITipo_basicoContext) []string {
	tipoBasicoResult := make([]string, 0)

	return tipoBasicoResult
}

func (v *AlgumaVisitor) VisitValor_constante(ctx parser.IValor_constanteContext) []string {
	valorConstanteResult := make([]string, 0)

	return valorConstanteResult
}

func (v *AlgumaVisitor) VisitRegistro(ctx parser.IRegistroContext) []string {
	registroResult := make([]string, 0)

	return registroResult
}

func (v *AlgumaVisitor) VisitIdentificador(ctx parser.IIdentificadorContext) []string {
	identificadorResult := make([]string, 0)

	aux := 0

	for ctx.ABRECHAVE(aux) != nil {
		result := v.VisitExp_aritmetica(ctx.Exp_aritmetica(aux))

		identificadorResult = append(identificadorResult, result...)
	}

	return identificadorResult
}

func (v *AlgumaVisitor) VisitTipo(ctx parser.ITipoContext) []string {
	tipoResult := make([]string, 0)

	return tipoResult
}

func (v *AlgumaVisitor) VisitExp_aritmetica(ctx parser.IExp_aritmeticaContext) []string {
	expResult := make([]string, 0)

	result := v.VisitTermo(ctx.Termo(0))

	expResult = append(expResult, result...)

	aux := 0

	for ctx.Op1(aux) != nil {
		result = v.VisitOp1(ctx.Op1(aux))

		expResult = append(expResult, result...)

		aux++

		result = v.VisitTermo(ctx.Termo(aux))

		expResult = append(expResult, result...)
	}

	return expResult
}

func (v *AlgumaVisitor) VisitTermo(ctx parser.ITermoContext) []string {
	termoResult := make([]string, 0)

	result := v.VisitFator(ctx.Fator(0))

	termoResult = append(termoResult, result...)

	aux := 0

	for ctx.Op2(aux) != nil {
		result = v.VisitOp2(ctx.Op2(aux))

		termoResult = append(termoResult, result...)

		aux++

		result = v.VisitFator(ctx.Fator(aux))

		termoResult = append(termoResult, result...)
	}

	return termoResult
}

func (v *AlgumaVisitor) VisitOp1(ctx parser.IOp1Context) []string {
	op1Result := make([]string, 0)

	return op1Result
}

func (v *AlgumaVisitor) VisitFator(ctx parser.IFatorContext) []string {
	fatorResult := make([]string, 0)

	return fatorResult
}

func (v *AlgumaVisitor) VisitOp2(ctx parser.IOp2Context) []string {
	op2Result := make([]string, 0)

	return op2Result
}
