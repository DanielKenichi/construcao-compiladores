// Code generated from T5Alguma.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // T5Alguma

import "github.com/antlr4-go/antlr/v4"

type BaseT5AlgumaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseT5AlgumaVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitDeclaracoes(ctx *DeclaracoesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitDeclaracoes_variaveis(ctx *Declaracoes_variaveisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitVariavel(ctx *VariavelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitIdentificador(ctx *IdentificadorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitExp_aritmetica(ctx *Exp_aritmeticaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitTermo(ctx *TermoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitFator(ctx *FatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitOp1(ctx *Op1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitOp2(ctx *Op2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitOp3(ctx *Op3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitParcela(ctx *ParcelaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitParcela_unario(ctx *Parcela_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitOp_unario(ctx *Op_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitParcela_nao_unario(ctx *Parcela_nao_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitTipo(ctx *TipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitTipo_basico(ctx *Tipo_basicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitTipo_variavel(ctx *Tipo_variavelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitValor_constante(ctx *Valor_constanteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitRegistro(ctx *RegistroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitParametro(ctx *ParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitParametros(ctx *ParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitDeclaracoes_funcoes(ctx *Declaracoes_funcoesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCorpo(ctx *CorpoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmd(ctx *CmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmdLeia(ctx *CmdLeiaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmdEscreva(ctx *CmdEscrevaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmdSe(ctx *CmdSeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmdCaso(ctx *CmdCasoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmdPara(ctx *CmdParaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmdEnquanto(ctx *CmdEnquantoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmdFaca(ctx *CmdFacaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmdAtribuicao(ctx *CmdAtribuicaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmdChamada(ctx *CmdChamadaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitCmdRetorne(ctx *CmdRetorneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitSelecao(ctx *SelecaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitItem_selecao(ctx *Item_selecaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitConstantes(ctx *ConstantesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitNumero_intervalo(ctx *Numero_intervaloContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitExp_relacional(ctx *Exp_relacionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitOp_relacional(ctx *Op_relacionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitExpressao(ctx *ExpressaoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitTermo_logico(ctx *Termo_logicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitFator_logico(ctx *Fator_logicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitParcela_logica(ctx *Parcela_logicaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitOp_logico_1(ctx *Op_logico_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT5AlgumaVisitor) VisitOp_logico_2(ctx *Op_logico_2Context) interface{} {
	return v.VisitChildren(ctx)
}
