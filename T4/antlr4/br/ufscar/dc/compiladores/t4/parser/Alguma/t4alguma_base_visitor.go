// Code generated from T4Alguma.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // T4Alguma

import "github.com/antlr4-go/antlr/v4"

type BaseT4AlgumaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseT4AlgumaVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitDeclaracoes(ctx *DeclaracoesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitDeclaracoes_variaveis(ctx *Declaracoes_variaveisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitVariavel(ctx *VariavelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitIdentificador(ctx *IdentificadorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitExp_aritmetica(ctx *Exp_aritmeticaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitTermo(ctx *TermoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitFator(ctx *FatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitOp1(ctx *Op1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitOp2(ctx *Op2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitOp3(ctx *Op3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitParcela(ctx *ParcelaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitParcela_unario(ctx *Parcela_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitOp_unario(ctx *Op_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitParcela_nao_unario(ctx *Parcela_nao_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitTipo(ctx *TipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitTipo_basico(ctx *Tipo_basicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitTipo_variavel(ctx *Tipo_variavelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitValor_constante(ctx *Valor_constanteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitRegistro(ctx *RegistroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitParametro(ctx *ParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitParametros(ctx *ParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitDeclaracoes_funcoes(ctx *Declaracoes_funcoesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCorpo(ctx *CorpoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmd(ctx *CmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmdLeia(ctx *CmdLeiaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmdEscreva(ctx *CmdEscrevaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmdSe(ctx *CmdSeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmdCaso(ctx *CmdCasoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmdPara(ctx *CmdParaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmdEnquanto(ctx *CmdEnquantoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmdFaca(ctx *CmdFacaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmdAtribuicao(ctx *CmdAtribuicaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmdChamada(ctx *CmdChamadaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitCmdRetorne(ctx *CmdRetorneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitSelecao(ctx *SelecaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitItem_selecao(ctx *Item_selecaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitConstantes(ctx *ConstantesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitNumero_intervalo(ctx *Numero_intervaloContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitExp_relacional(ctx *Exp_relacionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitOp_relacional(ctx *Op_relacionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitExpressao(ctx *ExpressaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitTermo_logico(ctx *Termo_logicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitFator_logico(ctx *Fator_logicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitParcela_logica(ctx *Parcela_logicaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitOp_logico_1(ctx *Op_logico_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT4AlgumaVisitor) VisitOp_logico_2(ctx *Op_logico_2Context) interface{} {
	return v.VisitChildren(ctx)
}
