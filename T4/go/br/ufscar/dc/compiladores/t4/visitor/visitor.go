package visitor

import (
	"fmt"

	parser "github.com/DanielKenichi/construcao-compiladores/T4/antlr4/br/ufscar/dc/compiladores/t4/parser/Alguma"
	"github.com/DanielKenichi/construcao-compiladores/T4/go/br/ufscar/dc/compiladores/t4/scope"
	"github.com/antlr4-go/antlr/v4"
)

/*
	Implementacao do nosso visitor para analise semântica.
	O visitor vai acumulando os erros em um slice de strings
*/

type AlgumaVisitor struct {
	parser.BaseT4AlgumaVisitor
	Scopes *scope.Scope
}

func New() *AlgumaVisitor {
	return &AlgumaVisitor{
		Scopes: scope.New(),
	}
}

func semanticError(token antlr.Token, message string) string {
	return fmt.Sprintf("Linha %v: %s\n", token.GetLine(), message)
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
	v.Scopes.NewScope()
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
			result := v.VisitCmdLeia(ctx.CmdLeia())

			cmdResult = append(cmdResult, result...)
			continue
		}

		if ctx.CmdEscreva() != nil {
			result := v.VisitCmdEscreva(ctx.CmdEscreva())

			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdSe() != nil {
			result := v.VisitCmdSe(ctx.CmdSe())

			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdCaso() != nil {
			result := v.VisitCmdCaso(ctx.CmdCaso())

			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdPara() != nil {
			result := v.VisitCmdPara(ctx.CmdPara())

			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdEnquanto() != nil {
			result := v.VisitCmdEnquanto(ctx.CmdEnquanto())

			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdFaca() != nil {
			result := v.VisitCmdFaca(ctx.CmdFaca())

			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdAtribuicao() != nil {
			result := v.VisitCmdAtribuicao(ctx.CmdAtribuicao())

			cmdResult = append(cmdResult, result...)

			continue
		}
	}

	return cmdResult
}

func (v *AlgumaVisitor) VisitCmdLeia(ctx parser.ICmdLeiaContext) []string {
	cmdLeiaResult := make([]string, 0)

	for _, ident := range ctx.AllIdentificador() {
		argToRead := ident.IDENT(0)

		if !v.Scopes.CurrentScope().Exists(argToRead.GetText()) {
			cmdLeiaResult = append(cmdLeiaResult,
				semanticError(argToRead.GetSymbol(), fmt.Sprintf("identificador %v nao declarado", argToRead.GetText())),
			)
		}

	}

	return cmdLeiaResult
}

func (v *AlgumaVisitor) VisitCmdEscreva(ctx parser.ICmdEscrevaContext) []string {
	cmdEscrevaResult := make([]string, 0)

	for _, expressao := range ctx.AllExpressao() {
		expressaoResult := v.VisitExpressao(expressao)

		cmdEscrevaResult = append(cmdEscrevaResult, expressaoResult...)
	}

	return cmdEscrevaResult
}

func (v *AlgumaVisitor) VisitCmdSe(ctx parser.ICmdSeContext) []string {
	cmdSeResult := make([]string, 0)

	result := v.VisitExpressao(ctx.Expressao())

	cmdSeResult = append(cmdSeResult, result...)

	result = v.VisitCmd(ctx.GetSeCmds())

	cmdSeResult = append(cmdSeResult, result...)

	if ctx.SENAO() != nil {
		result = v.VisitCmd(ctx.GetSenaoCmds())

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

func (v *AlgumaVisitor) VisitCmdEnquanto(ctx parser.ICmdEnquantoContext) []string {
	cmdEnquantoResult := make([]string, 0)

	result := v.VisitExpressao(ctx.Expressao())

	cmdEnquantoResult = append(cmdEnquantoResult, result...)

	result = v.VisitCmd(ctx.AllCmd())

	cmdEnquantoResult = append(cmdEnquantoResult, result...)

	return cmdEnquantoResult
}

func (v *AlgumaVisitor) VisitCmdFaca(ctx parser.ICmdFacaContext) []string {
	cmdFacaResult := make([]string, 0)

	result := v.VisitCmd(ctx.AllCmd())

	cmdFacaResult = append(cmdFacaResult, result...)

	result = v.VisitExpressao(ctx.Expressao())

	cmdFacaResult = append(cmdFacaResult, result...)

	return cmdFacaResult
}

func (v *AlgumaVisitor) VisitCmdAtribuicao(ctx parser.ICmdAtribuicaoContext) []string {
	cmdAtibuicaoResult := make([]string, 0)

	result := v.VisitIdentificador(ctx.Identificador())

	cmdAtibuicaoResult = append(cmdAtibuicaoResult, result...)

	result = v.VisitExpressao(ctx.Expressao())

	cmdAtibuicaoResult = append(cmdAtibuicaoResult, result...)

	result = v.VerifyVarAttribution(ctx)

	cmdAtibuicaoResult = append(cmdAtibuicaoResult, result...)

	return cmdAtibuicaoResult
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
			continue
		}

		if ctx.CONSTANTE() != nil {

			result := v.VisitTipo_basico(ctx.Tipo_basico())

			variaveisResult = append(variaveisResult, result...)

			result = v.VisitValor_constante(ctx.Valor_constante())

			variaveisResult = append(variaveisResult, result...)

			result = v.AddConstToSymbolTable(ctx.IDENT(), ctx.Tipo_basico())

			variaveisResult = append(variaveisResult, result...)
			continue
		}

		if ctx.TIPO() != nil {

			result := v.VisitRegistro(ctx.Registro())

			variaveisResult = append(variaveisResult, result...)

			result = v.AddRegToSymbolTable(ctx.IDENT())

			variaveisResult = append(variaveisResult, result...)

			continue
		}
	}

	return variaveisResult
}

func (v *AlgumaVisitor) VisitDeclaracoes_funcoes(ctxs []parser.IDeclaracoes_funcoesContext) []string {
	funcoesResult := make([]string, 0)
	for _, ctx := range ctxs {
		result := v.AddFuncToSymbolTable(ctx.IDENT())
		funcoesResult = append(funcoesResult, result...)
	}

	return funcoesResult
}

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

	_, result := v.VerifyExpression(ctx)

	expressaoResult = append(expressaoResult, result...)

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

		aux++
	}

	return identificadorResult
}

func (v *AlgumaVisitor) VisitTipo(ctx parser.ITipoContext) []string {
	tipoResult := make([]string, 0)

	if ctx.Registro() != nil {
		result := v.VisitRegistro(ctx.Registro())

		tipoResult = append(tipoResult, result...)
	}

	if ctx.Tipo_variavel() != nil {
		result := v.VisitTipo_variavel(ctx.Tipo_variavel())

		tipoResult = append(tipoResult, result...)
	}

	return tipoResult
}

func (v *AlgumaVisitor) VisitTipo_variavel(ctx parser.ITipo_variavelContext) []string {
	tipoVariavelResult := make([]string, 0)

	ident := ctx.IDENT()

	if ident != nil {
		// TODO: Conferindo em todos os escopos, ou só apenas no escopo atual e acima?
		exists := false
		for _, scope := range v.Scopes.Stack {
			if scope.Exists(ident.GetText()) {
				exists = true
			}
		}
		if !exists {
			tipoVariavelResult = append(tipoVariavelResult,
				semanticError(ident.GetSymbol(), fmt.Sprintf("tipo %v nao declarado", ident.GetText())),
			)
		}
	}

	return tipoVariavelResult
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
