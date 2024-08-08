// Generated from /home/lucky/UFSCAR/Compiladores/construcao-compiladores/T4/antlr4/br/ufscar/dc/compiladores/t4/parser/T4Alguma.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link T4AlgumaParser}.
 */
public interface T4AlgumaListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#programa}.
	 * @param ctx the parse tree
	 */
	void enterPrograma(T4AlgumaParser.ProgramaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#programa}.
	 * @param ctx the parse tree
	 */
	void exitPrograma(T4AlgumaParser.ProgramaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#declaracoes}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracoes(T4AlgumaParser.DeclaracoesContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#declaracoes}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracoes(T4AlgumaParser.DeclaracoesContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#declaracoes_variaveis}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracoes_variaveis(T4AlgumaParser.Declaracoes_variaveisContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#declaracoes_variaveis}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracoes_variaveis(T4AlgumaParser.Declaracoes_variaveisContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#variavel}.
	 * @param ctx the parse tree
	 */
	void enterVariavel(T4AlgumaParser.VariavelContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#variavel}.
	 * @param ctx the parse tree
	 */
	void exitVariavel(T4AlgumaParser.VariavelContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#identificador}.
	 * @param ctx the parse tree
	 */
	void enterIdentificador(T4AlgumaParser.IdentificadorContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#identificador}.
	 * @param ctx the parse tree
	 */
	void exitIdentificador(T4AlgumaParser.IdentificadorContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#exp_aritmetica}.
	 * @param ctx the parse tree
	 */
	void enterExp_aritmetica(T4AlgumaParser.Exp_aritmeticaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#exp_aritmetica}.
	 * @param ctx the parse tree
	 */
	void exitExp_aritmetica(T4AlgumaParser.Exp_aritmeticaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#termo}.
	 * @param ctx the parse tree
	 */
	void enterTermo(T4AlgumaParser.TermoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#termo}.
	 * @param ctx the parse tree
	 */
	void exitTermo(T4AlgumaParser.TermoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#fator}.
	 * @param ctx the parse tree
	 */
	void enterFator(T4AlgumaParser.FatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#fator}.
	 * @param ctx the parse tree
	 */
	void exitFator(T4AlgumaParser.FatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#op1}.
	 * @param ctx the parse tree
	 */
	void enterOp1(T4AlgumaParser.Op1Context ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#op1}.
	 * @param ctx the parse tree
	 */
	void exitOp1(T4AlgumaParser.Op1Context ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#op2}.
	 * @param ctx the parse tree
	 */
	void enterOp2(T4AlgumaParser.Op2Context ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#op2}.
	 * @param ctx the parse tree
	 */
	void exitOp2(T4AlgumaParser.Op2Context ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#op3}.
	 * @param ctx the parse tree
	 */
	void enterOp3(T4AlgumaParser.Op3Context ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#op3}.
	 * @param ctx the parse tree
	 */
	void exitOp3(T4AlgumaParser.Op3Context ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#parcela}.
	 * @param ctx the parse tree
	 */
	void enterParcela(T4AlgumaParser.ParcelaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#parcela}.
	 * @param ctx the parse tree
	 */
	void exitParcela(T4AlgumaParser.ParcelaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#parcela_unario}.
	 * @param ctx the parse tree
	 */
	void enterParcela_unario(T4AlgumaParser.Parcela_unarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#parcela_unario}.
	 * @param ctx the parse tree
	 */
	void exitParcela_unario(T4AlgumaParser.Parcela_unarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#op_unario}.
	 * @param ctx the parse tree
	 */
	void enterOp_unario(T4AlgumaParser.Op_unarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#op_unario}.
	 * @param ctx the parse tree
	 */
	void exitOp_unario(T4AlgumaParser.Op_unarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#parcela_nao_unario}.
	 * @param ctx the parse tree
	 */
	void enterParcela_nao_unario(T4AlgumaParser.Parcela_nao_unarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#parcela_nao_unario}.
	 * @param ctx the parse tree
	 */
	void exitParcela_nao_unario(T4AlgumaParser.Parcela_nao_unarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo(T4AlgumaParser.TipoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo(T4AlgumaParser.TipoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#tipo_basico}.
	 * @param ctx the parse tree
	 */
	void enterTipo_basico(T4AlgumaParser.Tipo_basicoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#tipo_basico}.
	 * @param ctx the parse tree
	 */
	void exitTipo_basico(T4AlgumaParser.Tipo_basicoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#tipo_variavel}.
	 * @param ctx the parse tree
	 */
	void enterTipo_variavel(T4AlgumaParser.Tipo_variavelContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#tipo_variavel}.
	 * @param ctx the parse tree
	 */
	void exitTipo_variavel(T4AlgumaParser.Tipo_variavelContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#valor_constante}.
	 * @param ctx the parse tree
	 */
	void enterValor_constante(T4AlgumaParser.Valor_constanteContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#valor_constante}.
	 * @param ctx the parse tree
	 */
	void exitValor_constante(T4AlgumaParser.Valor_constanteContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#registro}.
	 * @param ctx the parse tree
	 */
	void enterRegistro(T4AlgumaParser.RegistroContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#registro}.
	 * @param ctx the parse tree
	 */
	void exitRegistro(T4AlgumaParser.RegistroContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#parametro}.
	 * @param ctx the parse tree
	 */
	void enterParametro(T4AlgumaParser.ParametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#parametro}.
	 * @param ctx the parse tree
	 */
	void exitParametro(T4AlgumaParser.ParametroContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#parametros}.
	 * @param ctx the parse tree
	 */
	void enterParametros(T4AlgumaParser.ParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#parametros}.
	 * @param ctx the parse tree
	 */
	void exitParametros(T4AlgumaParser.ParametrosContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#declaracoes_funcoes}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracoes_funcoes(T4AlgumaParser.Declaracoes_funcoesContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#declaracoes_funcoes}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracoes_funcoes(T4AlgumaParser.Declaracoes_funcoesContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#corpo}.
	 * @param ctx the parse tree
	 */
	void enterCorpo(T4AlgumaParser.CorpoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#corpo}.
	 * @param ctx the parse tree
	 */
	void exitCorpo(T4AlgumaParser.CorpoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmd}.
	 * @param ctx the parse tree
	 */
	void enterCmd(T4AlgumaParser.CmdContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmd}.
	 * @param ctx the parse tree
	 */
	void exitCmd(T4AlgumaParser.CmdContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmdLeia}.
	 * @param ctx the parse tree
	 */
	void enterCmdLeia(T4AlgumaParser.CmdLeiaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmdLeia}.
	 * @param ctx the parse tree
	 */
	void exitCmdLeia(T4AlgumaParser.CmdLeiaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmdEscreva}.
	 * @param ctx the parse tree
	 */
	void enterCmdEscreva(T4AlgumaParser.CmdEscrevaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmdEscreva}.
	 * @param ctx the parse tree
	 */
	void exitCmdEscreva(T4AlgumaParser.CmdEscrevaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmdSe}.
	 * @param ctx the parse tree
	 */
	void enterCmdSe(T4AlgumaParser.CmdSeContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmdSe}.
	 * @param ctx the parse tree
	 */
	void exitCmdSe(T4AlgumaParser.CmdSeContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmdCaso}.
	 * @param ctx the parse tree
	 */
	void enterCmdCaso(T4AlgumaParser.CmdCasoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmdCaso}.
	 * @param ctx the parse tree
	 */
	void exitCmdCaso(T4AlgumaParser.CmdCasoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmdPara}.
	 * @param ctx the parse tree
	 */
	void enterCmdPara(T4AlgumaParser.CmdParaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmdPara}.
	 * @param ctx the parse tree
	 */
	void exitCmdPara(T4AlgumaParser.CmdParaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmdEnquanto}.
	 * @param ctx the parse tree
	 */
	void enterCmdEnquanto(T4AlgumaParser.CmdEnquantoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmdEnquanto}.
	 * @param ctx the parse tree
	 */
	void exitCmdEnquanto(T4AlgumaParser.CmdEnquantoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmdFaca}.
	 * @param ctx the parse tree
	 */
	void enterCmdFaca(T4AlgumaParser.CmdFacaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmdFaca}.
	 * @param ctx the parse tree
	 */
	void exitCmdFaca(T4AlgumaParser.CmdFacaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmdAtribuicao}.
	 * @param ctx the parse tree
	 */
	void enterCmdAtribuicao(T4AlgumaParser.CmdAtribuicaoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmdAtribuicao}.
	 * @param ctx the parse tree
	 */
	void exitCmdAtribuicao(T4AlgumaParser.CmdAtribuicaoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmdChamada}.
	 * @param ctx the parse tree
	 */
	void enterCmdChamada(T4AlgumaParser.CmdChamadaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmdChamada}.
	 * @param ctx the parse tree
	 */
	void exitCmdChamada(T4AlgumaParser.CmdChamadaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#cmdRetorne}.
	 * @param ctx the parse tree
	 */
	void enterCmdRetorne(T4AlgumaParser.CmdRetorneContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#cmdRetorne}.
	 * @param ctx the parse tree
	 */
	void exitCmdRetorne(T4AlgumaParser.CmdRetorneContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#selecao}.
	 * @param ctx the parse tree
	 */
	void enterSelecao(T4AlgumaParser.SelecaoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#selecao}.
	 * @param ctx the parse tree
	 */
	void exitSelecao(T4AlgumaParser.SelecaoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#item_selecao}.
	 * @param ctx the parse tree
	 */
	void enterItem_selecao(T4AlgumaParser.Item_selecaoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#item_selecao}.
	 * @param ctx the parse tree
	 */
	void exitItem_selecao(T4AlgumaParser.Item_selecaoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#constantes}.
	 * @param ctx the parse tree
	 */
	void enterConstantes(T4AlgumaParser.ConstantesContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#constantes}.
	 * @param ctx the parse tree
	 */
	void exitConstantes(T4AlgumaParser.ConstantesContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#numero_intervalo}.
	 * @param ctx the parse tree
	 */
	void enterNumero_intervalo(T4AlgumaParser.Numero_intervaloContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#numero_intervalo}.
	 * @param ctx the parse tree
	 */
	void exitNumero_intervalo(T4AlgumaParser.Numero_intervaloContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#exp_relacional}.
	 * @param ctx the parse tree
	 */
	void enterExp_relacional(T4AlgumaParser.Exp_relacionalContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#exp_relacional}.
	 * @param ctx the parse tree
	 */
	void exitExp_relacional(T4AlgumaParser.Exp_relacionalContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#op_relacional}.
	 * @param ctx the parse tree
	 */
	void enterOp_relacional(T4AlgumaParser.Op_relacionalContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#op_relacional}.
	 * @param ctx the parse tree
	 */
	void exitOp_relacional(T4AlgumaParser.Op_relacionalContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#expressao}.
	 * @param ctx the parse tree
	 */
	void enterExpressao(T4AlgumaParser.ExpressaoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#expressao}.
	 * @param ctx the parse tree
	 */
	void exitExpressao(T4AlgumaParser.ExpressaoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#termo_logico}.
	 * @param ctx the parse tree
	 */
	void enterTermo_logico(T4AlgumaParser.Termo_logicoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#termo_logico}.
	 * @param ctx the parse tree
	 */
	void exitTermo_logico(T4AlgumaParser.Termo_logicoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#fator_logico}.
	 * @param ctx the parse tree
	 */
	void enterFator_logico(T4AlgumaParser.Fator_logicoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#fator_logico}.
	 * @param ctx the parse tree
	 */
	void exitFator_logico(T4AlgumaParser.Fator_logicoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#parcela_logica}.
	 * @param ctx the parse tree
	 */
	void enterParcela_logica(T4AlgumaParser.Parcela_logicaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#parcela_logica}.
	 * @param ctx the parse tree
	 */
	void exitParcela_logica(T4AlgumaParser.Parcela_logicaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#op_logico_1}.
	 * @param ctx the parse tree
	 */
	void enterOp_logico_1(T4AlgumaParser.Op_logico_1Context ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#op_logico_1}.
	 * @param ctx the parse tree
	 */
	void exitOp_logico_1(T4AlgumaParser.Op_logico_1Context ctx);
	/**
	 * Enter a parse tree produced by {@link T4AlgumaParser#op_logico_2}.
	 * @param ctx the parse tree
	 */
	void enterOp_logico_2(T4AlgumaParser.Op_logico_2Context ctx);
	/**
	 * Exit a parse tree produced by {@link T4AlgumaParser#op_logico_2}.
	 * @param ctx the parse tree
	 */
	void exitOp_logico_2(T4AlgumaParser.Op_logico_2Context ctx);
}