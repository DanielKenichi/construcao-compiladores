// Code generated from T3Alguma.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // T3Alguma

import "github.com/antlr4-go/antlr/v4"

type BaseT3AlgumaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseT3AlgumaVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitDeclaracoes(ctx *DeclaracoesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitDeclaracoes_variaveis(ctx *Declaracoes_variaveisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitVariavel(ctx *VariavelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitIdentificador(ctx *IdentificadorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitExp_aritmetica(ctx *Exp_aritmeticaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitTermo(ctx *TermoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitFator(ctx *FatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitOp1(ctx *Op1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitOp2(ctx *Op2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitOp3(ctx *Op3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitParcela(ctx *ParcelaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitParcela_unario(ctx *Parcela_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitOp_unario(ctx *Op_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitParcela_nao_unario(ctx *Parcela_nao_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitTipo(ctx *TipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitTipo_basico(ctx *Tipo_basicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitTipo_variavel(ctx *Tipo_variavelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitValor_constante(ctx *Valor_constanteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitRegistro(ctx *RegistroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitParametro(ctx *ParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitParametros(ctx *ParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitDeclaracoes_funcoes(ctx *Declaracoes_funcoesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCorpo(ctx *CorpoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmd(ctx *CmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmdLeia(ctx *CmdLeiaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmdEscreva(ctx *CmdEscrevaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmdSe(ctx *CmdSeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmdCaso(ctx *CmdCasoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmdPara(ctx *CmdParaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmdEnquanto(ctx *CmdEnquantoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmdFaca(ctx *CmdFacaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmdAtribuicao(ctx *CmdAtribuicaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmdChamada(ctx *CmdChamadaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitCmdRetorne(ctx *CmdRetorneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitSelecao(ctx *SelecaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitItem_selecao(ctx *Item_selecaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitConstantes(ctx *ConstantesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitNumero_intervalo(ctx *Numero_intervaloContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitExp_relacional(ctx *Exp_relacionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitOp_relacional(ctx *Op_relacionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitExpressao(ctx *ExpressaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitTermo_logico(ctx *Termo_logicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitFator_logico(ctx *Fator_logicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitParcela_logica(ctx *Parcela_logicaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitOp_logico_1(ctx *Op_logico_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT3AlgumaVisitor) VisitOp_logico_2(ctx *Op_logico_2Context) interface{} {
	return v.VisitChildren(ctx)
}
