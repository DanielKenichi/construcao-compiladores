// Code generated from T3Alguma.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // T3Alguma

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by T3AlgumaParser.
type T3AlgumaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by T3AlgumaParser#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#declaracoes.
	VisitDeclaracoes(ctx *DeclaracoesContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#declaracoes_variaveis.
	VisitDeclaracoes_variaveis(ctx *Declaracoes_variaveisContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#variavel.
	VisitVariavel(ctx *VariavelContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#identificador.
	VisitIdentificador(ctx *IdentificadorContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#exp_aritmetica.
	VisitExp_aritmetica(ctx *Exp_aritmeticaContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#termo.
	VisitTermo(ctx *TermoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#fator.
	VisitFator(ctx *FatorContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#op1.
	VisitOp1(ctx *Op1Context) interface{}

	// Visit a parse tree produced by T3AlgumaParser#op2.
	VisitOp2(ctx *Op2Context) interface{}

	// Visit a parse tree produced by T3AlgumaParser#op3.
	VisitOp3(ctx *Op3Context) interface{}

	// Visit a parse tree produced by T3AlgumaParser#parcela.
	VisitParcela(ctx *ParcelaContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#parcela_unario.
	VisitParcela_unario(ctx *Parcela_unarioContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#op_unario.
	VisitOp_unario(ctx *Op_unarioContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#parcela_nao_unario.
	VisitParcela_nao_unario(ctx *Parcela_nao_unarioContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#tipo.
	VisitTipo(ctx *TipoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#tipo_basico.
	VisitTipo_basico(ctx *Tipo_basicoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#tipo_variavel.
	VisitTipo_variavel(ctx *Tipo_variavelContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#valor_constante.
	VisitValor_constante(ctx *Valor_constanteContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#registro.
	VisitRegistro(ctx *RegistroContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#parametro.
	VisitParametro(ctx *ParametroContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#parametros.
	VisitParametros(ctx *ParametrosContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#declaracoes_funcoes.
	VisitDeclaracoes_funcoes(ctx *Declaracoes_funcoesContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#corpo.
	VisitCorpo(ctx *CorpoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmd.
	VisitCmd(ctx *CmdContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmdLeia.
	VisitCmdLeia(ctx *CmdLeiaContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmdEscreva.
	VisitCmdEscreva(ctx *CmdEscrevaContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmdSe.
	VisitCmdSe(ctx *CmdSeContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmdCaso.
	VisitCmdCaso(ctx *CmdCasoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmdPara.
	VisitCmdPara(ctx *CmdParaContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmdEnquanto.
	VisitCmdEnquanto(ctx *CmdEnquantoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmdFaca.
	VisitCmdFaca(ctx *CmdFacaContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmdAtribuicao.
	VisitCmdAtribuicao(ctx *CmdAtribuicaoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmdChamada.
	VisitCmdChamada(ctx *CmdChamadaContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#cmdRetorne.
	VisitCmdRetorne(ctx *CmdRetorneContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#selecao.
	VisitSelecao(ctx *SelecaoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#item_selecao.
	VisitItem_selecao(ctx *Item_selecaoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#constantes.
	VisitConstantes(ctx *ConstantesContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#numero_intervalo.
	VisitNumero_intervalo(ctx *Numero_intervaloContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#exp_relacional.
	VisitExp_relacional(ctx *Exp_relacionalContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#op_relacional.
	VisitOp_relacional(ctx *Op_relacionalContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#expressao.
	VisitExpressao(ctx *ExpressaoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#termo_logico.
	VisitTermo_logico(ctx *Termo_logicoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#fator_logico.
	VisitFator_logico(ctx *Fator_logicoContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#parcela_logica.
	VisitParcela_logica(ctx *Parcela_logicaContext) interface{}

	// Visit a parse tree produced by T3AlgumaParser#op_logico_1.
	VisitOp_logico_1(ctx *Op_logico_1Context) interface{}

	// Visit a parse tree produced by T3AlgumaParser#op_logico_2.
	VisitOp_logico_2(ctx *Op_logico_2Context) interface{}
}
