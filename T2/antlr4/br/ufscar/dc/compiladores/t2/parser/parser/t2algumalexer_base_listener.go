// Code generated from T2AlgumaLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // T2AlgumaLexer

import "github.com/antlr4-go/antlr/v4"

// BaseT2AlgumaLexerListener is a complete listener for a parse tree produced by T2AlgumaLexerParser.
type BaseT2AlgumaLexerListener struct{}

var _ T2AlgumaLexerListener = &BaseT2AlgumaLexerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseT2AlgumaLexerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseT2AlgumaLexerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseT2AlgumaLexerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseT2AlgumaLexerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseT2AlgumaLexerListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseT2AlgumaLexerListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterDeclaracoes is called when production declaracoes is entered.
func (s *BaseT2AlgumaLexerListener) EnterDeclaracoes(ctx *DeclaracoesContext) {}

// ExitDeclaracoes is called when production declaracoes is exited.
func (s *BaseT2AlgumaLexerListener) ExitDeclaracoes(ctx *DeclaracoesContext) {}

// EnterDeclaracoes_variaveis is called when production declaracoes_variaveis is entered.
func (s *BaseT2AlgumaLexerListener) EnterDeclaracoes_variaveis(ctx *Declaracoes_variaveisContext) {}

// ExitDeclaracoes_variaveis is called when production declaracoes_variaveis is exited.
func (s *BaseT2AlgumaLexerListener) ExitDeclaracoes_variaveis(ctx *Declaracoes_variaveisContext) {}

// EnterVariavel is called when production variavel is entered.
func (s *BaseT2AlgumaLexerListener) EnterVariavel(ctx *VariavelContext) {}

// ExitVariavel is called when production variavel is exited.
func (s *BaseT2AlgumaLexerListener) ExitVariavel(ctx *VariavelContext) {}

// EnterIdentificador is called when production identificador is entered.
func (s *BaseT2AlgumaLexerListener) EnterIdentificador(ctx *IdentificadorContext) {}

// ExitIdentificador is called when production identificador is exited.
func (s *BaseT2AlgumaLexerListener) ExitIdentificador(ctx *IdentificadorContext) {}

// EnterExp_aritmetica is called when production exp_aritmetica is entered.
func (s *BaseT2AlgumaLexerListener) EnterExp_aritmetica(ctx *Exp_aritmeticaContext) {}

// ExitExp_aritmetica is called when production exp_aritmetica is exited.
func (s *BaseT2AlgumaLexerListener) ExitExp_aritmetica(ctx *Exp_aritmeticaContext) {}

// EnterTermo is called when production termo is entered.
func (s *BaseT2AlgumaLexerListener) EnterTermo(ctx *TermoContext) {}

// ExitTermo is called when production termo is exited.
func (s *BaseT2AlgumaLexerListener) ExitTermo(ctx *TermoContext) {}

// EnterFator is called when production fator is entered.
func (s *BaseT2AlgumaLexerListener) EnterFator(ctx *FatorContext) {}

// ExitFator is called when production fator is exited.
func (s *BaseT2AlgumaLexerListener) ExitFator(ctx *FatorContext) {}

// EnterOp1 is called when production op1 is entered.
func (s *BaseT2AlgumaLexerListener) EnterOp1(ctx *Op1Context) {}

// ExitOp1 is called when production op1 is exited.
func (s *BaseT2AlgumaLexerListener) ExitOp1(ctx *Op1Context) {}

// EnterOp2 is called when production op2 is entered.
func (s *BaseT2AlgumaLexerListener) EnterOp2(ctx *Op2Context) {}

// ExitOp2 is called when production op2 is exited.
func (s *BaseT2AlgumaLexerListener) ExitOp2(ctx *Op2Context) {}

// EnterOp3 is called when production op3 is entered.
func (s *BaseT2AlgumaLexerListener) EnterOp3(ctx *Op3Context) {}

// ExitOp3 is called when production op3 is exited.
func (s *BaseT2AlgumaLexerListener) ExitOp3(ctx *Op3Context) {}

// EnterParcela is called when production parcela is entered.
func (s *BaseT2AlgumaLexerListener) EnterParcela(ctx *ParcelaContext) {}

// ExitParcela is called when production parcela is exited.
func (s *BaseT2AlgumaLexerListener) ExitParcela(ctx *ParcelaContext) {}

// EnterParcela_unario is called when production parcela_unario is entered.
func (s *BaseT2AlgumaLexerListener) EnterParcela_unario(ctx *Parcela_unarioContext) {}

// ExitParcela_unario is called when production parcela_unario is exited.
func (s *BaseT2AlgumaLexerListener) ExitParcela_unario(ctx *Parcela_unarioContext) {}

// EnterOp_unario is called when production op_unario is entered.
func (s *BaseT2AlgumaLexerListener) EnterOp_unario(ctx *Op_unarioContext) {}

// ExitOp_unario is called when production op_unario is exited.
func (s *BaseT2AlgumaLexerListener) ExitOp_unario(ctx *Op_unarioContext) {}

// EnterParcela_nao_unario is called when production parcela_nao_unario is entered.
func (s *BaseT2AlgumaLexerListener) EnterParcela_nao_unario(ctx *Parcela_nao_unarioContext) {}

// ExitParcela_nao_unario is called when production parcela_nao_unario is exited.
func (s *BaseT2AlgumaLexerListener) ExitParcela_nao_unario(ctx *Parcela_nao_unarioContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseT2AlgumaLexerListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseT2AlgumaLexerListener) ExitTipo(ctx *TipoContext) {}

// EnterTipo_basico is called when production tipo_basico is entered.
func (s *BaseT2AlgumaLexerListener) EnterTipo_basico(ctx *Tipo_basicoContext) {}

// ExitTipo_basico is called when production tipo_basico is exited.
func (s *BaseT2AlgumaLexerListener) ExitTipo_basico(ctx *Tipo_basicoContext) {}

// EnterTipo_variavel is called when production tipo_variavel is entered.
func (s *BaseT2AlgumaLexerListener) EnterTipo_variavel(ctx *Tipo_variavelContext) {}

// ExitTipo_variavel is called when production tipo_variavel is exited.
func (s *BaseT2AlgumaLexerListener) ExitTipo_variavel(ctx *Tipo_variavelContext) {}

// EnterValor_constante is called when production valor_constante is entered.
func (s *BaseT2AlgumaLexerListener) EnterValor_constante(ctx *Valor_constanteContext) {}

// ExitValor_constante is called when production valor_constante is exited.
func (s *BaseT2AlgumaLexerListener) ExitValor_constante(ctx *Valor_constanteContext) {}

// EnterRegistro is called when production registro is entered.
func (s *BaseT2AlgumaLexerListener) EnterRegistro(ctx *RegistroContext) {}

// ExitRegistro is called when production registro is exited.
func (s *BaseT2AlgumaLexerListener) ExitRegistro(ctx *RegistroContext) {}

// EnterParametro is called when production parametro is entered.
func (s *BaseT2AlgumaLexerListener) EnterParametro(ctx *ParametroContext) {}

// ExitParametro is called when production parametro is exited.
func (s *BaseT2AlgumaLexerListener) ExitParametro(ctx *ParametroContext) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseT2AlgumaLexerListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseT2AlgumaLexerListener) ExitParametros(ctx *ParametrosContext) {}

// EnterDeclaracoes_funcoes is called when production declaracoes_funcoes is entered.
func (s *BaseT2AlgumaLexerListener) EnterDeclaracoes_funcoes(ctx *Declaracoes_funcoesContext) {}

// ExitDeclaracoes_funcoes is called when production declaracoes_funcoes is exited.
func (s *BaseT2AlgumaLexerListener) ExitDeclaracoes_funcoes(ctx *Declaracoes_funcoesContext) {}

// EnterCorpo is called when production corpo is entered.
func (s *BaseT2AlgumaLexerListener) EnterCorpo(ctx *CorpoContext) {}

// ExitCorpo is called when production corpo is exited.
func (s *BaseT2AlgumaLexerListener) ExitCorpo(ctx *CorpoContext) {}

// EnterCmd is called when production cmd is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmd(ctx *CmdContext) {}

// ExitCmd is called when production cmd is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmd(ctx *CmdContext) {}

// EnterCmdLeia is called when production cmdLeia is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmdLeia(ctx *CmdLeiaContext) {}

// ExitCmdLeia is called when production cmdLeia is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmdLeia(ctx *CmdLeiaContext) {}

// EnterCmdEscreva is called when production cmdEscreva is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmdEscreva(ctx *CmdEscrevaContext) {}

// ExitCmdEscreva is called when production cmdEscreva is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmdEscreva(ctx *CmdEscrevaContext) {}

// EnterCmdSe is called when production cmdSe is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmdSe(ctx *CmdSeContext) {}

// ExitCmdSe is called when production cmdSe is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmdSe(ctx *CmdSeContext) {}

// EnterCmdCaso is called when production cmdCaso is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmdCaso(ctx *CmdCasoContext) {}

// ExitCmdCaso is called when production cmdCaso is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmdCaso(ctx *CmdCasoContext) {}

// EnterCmdPara is called when production cmdPara is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmdPara(ctx *CmdParaContext) {}

// ExitCmdPara is called when production cmdPara is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmdPara(ctx *CmdParaContext) {}

// EnterCmdEnquanto is called when production cmdEnquanto is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmdEnquanto(ctx *CmdEnquantoContext) {}

// ExitCmdEnquanto is called when production cmdEnquanto is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmdEnquanto(ctx *CmdEnquantoContext) {}

// EnterCmdFaca is called when production cmdFaca is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmdFaca(ctx *CmdFacaContext) {}

// ExitCmdFaca is called when production cmdFaca is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmdFaca(ctx *CmdFacaContext) {}

// EnterCmdAtribuicao is called when production cmdAtribuicao is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmdAtribuicao(ctx *CmdAtribuicaoContext) {}

// ExitCmdAtribuicao is called when production cmdAtribuicao is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmdAtribuicao(ctx *CmdAtribuicaoContext) {}

// EnterCmdChamada is called when production cmdChamada is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmdChamada(ctx *CmdChamadaContext) {}

// ExitCmdChamada is called when production cmdChamada is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmdChamada(ctx *CmdChamadaContext) {}

// EnterCmdRetorne is called when production cmdRetorne is entered.
func (s *BaseT2AlgumaLexerListener) EnterCmdRetorne(ctx *CmdRetorneContext) {}

// ExitCmdRetorne is called when production cmdRetorne is exited.
func (s *BaseT2AlgumaLexerListener) ExitCmdRetorne(ctx *CmdRetorneContext) {}

// EnterSelecao is called when production selecao is entered.
func (s *BaseT2AlgumaLexerListener) EnterSelecao(ctx *SelecaoContext) {}

// ExitSelecao is called when production selecao is exited.
func (s *BaseT2AlgumaLexerListener) ExitSelecao(ctx *SelecaoContext) {}

// EnterItem_selecao is called when production item_selecao is entered.
func (s *BaseT2AlgumaLexerListener) EnterItem_selecao(ctx *Item_selecaoContext) {}

// ExitItem_selecao is called when production item_selecao is exited.
func (s *BaseT2AlgumaLexerListener) ExitItem_selecao(ctx *Item_selecaoContext) {}

// EnterConstantes is called when production constantes is entered.
func (s *BaseT2AlgumaLexerListener) EnterConstantes(ctx *ConstantesContext) {}

// ExitConstantes is called when production constantes is exited.
func (s *BaseT2AlgumaLexerListener) ExitConstantes(ctx *ConstantesContext) {}

// EnterNumero_intervalo is called when production numero_intervalo is entered.
func (s *BaseT2AlgumaLexerListener) EnterNumero_intervalo(ctx *Numero_intervaloContext) {}

// ExitNumero_intervalo is called when production numero_intervalo is exited.
func (s *BaseT2AlgumaLexerListener) ExitNumero_intervalo(ctx *Numero_intervaloContext) {}

// EnterExp_relacional is called when production exp_relacional is entered.
func (s *BaseT2AlgumaLexerListener) EnterExp_relacional(ctx *Exp_relacionalContext) {}

// ExitExp_relacional is called when production exp_relacional is exited.
func (s *BaseT2AlgumaLexerListener) ExitExp_relacional(ctx *Exp_relacionalContext) {}

// EnterOp_relacional is called when production op_relacional is entered.
func (s *BaseT2AlgumaLexerListener) EnterOp_relacional(ctx *Op_relacionalContext) {}

// ExitOp_relacional is called when production op_relacional is exited.
func (s *BaseT2AlgumaLexerListener) ExitOp_relacional(ctx *Op_relacionalContext) {}

// EnterExpressao is called when production expressao is entered.
func (s *BaseT2AlgumaLexerListener) EnterExpressao(ctx *ExpressaoContext) {}

// ExitExpressao is called when production expressao is exited.
func (s *BaseT2AlgumaLexerListener) ExitExpressao(ctx *ExpressaoContext) {}

// EnterTermo_logico is called when production termo_logico is entered.
func (s *BaseT2AlgumaLexerListener) EnterTermo_logico(ctx *Termo_logicoContext) {}

// ExitTermo_logico is called when production termo_logico is exited.
func (s *BaseT2AlgumaLexerListener) ExitTermo_logico(ctx *Termo_logicoContext) {}

// EnterFator_logico is called when production fator_logico is entered.
func (s *BaseT2AlgumaLexerListener) EnterFator_logico(ctx *Fator_logicoContext) {}

// ExitFator_logico is called when production fator_logico is exited.
func (s *BaseT2AlgumaLexerListener) ExitFator_logico(ctx *Fator_logicoContext) {}

// EnterParcela_logica is called when production parcela_logica is entered.
func (s *BaseT2AlgumaLexerListener) EnterParcela_logica(ctx *Parcela_logicaContext) {}

// ExitParcela_logica is called when production parcela_logica is exited.
func (s *BaseT2AlgumaLexerListener) ExitParcela_logica(ctx *Parcela_logicaContext) {}

// EnterOp_logico_1 is called when production op_logico_1 is entered.
func (s *BaseT2AlgumaLexerListener) EnterOp_logico_1(ctx *Op_logico_1Context) {}

// ExitOp_logico_1 is called when production op_logico_1 is exited.
func (s *BaseT2AlgumaLexerListener) ExitOp_logico_1(ctx *Op_logico_1Context) {}

// EnterOp_logico_2 is called when production op_logico_2 is entered.
func (s *BaseT2AlgumaLexerListener) EnterOp_logico_2(ctx *Op_logico_2Context) {}

// ExitOp_logico_2 is called when production op_logico_2 is exited.
func (s *BaseT2AlgumaLexerListener) ExitOp_logico_2(ctx *Op_logico_2Context) {}
