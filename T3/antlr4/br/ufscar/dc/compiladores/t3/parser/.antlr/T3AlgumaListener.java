// Generated from /home/rodrigo/UFSCAR/Compiladores/construcao-compiladores/T3/antlr4/br/ufscar/dc/compiladores/t3/parser/T3Alguma.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link T3AlgumaParser}.
 */
public interface T3AlgumaListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#programa}.
	 * @param ctx the parse tree
	 */
	void enterPrograma(T3AlgumaParser.ProgramaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#programa}.
	 * @param ctx the parse tree
	 */
	void exitPrograma(T3AlgumaParser.ProgramaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#declaracoes}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracoes(T3AlgumaParser.DeclaracoesContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#declaracoes}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracoes(T3AlgumaParser.DeclaracoesContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#declaracoes_variaveis}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracoes_variaveis(T3AlgumaParser.Declaracoes_variaveisContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#declaracoes_variaveis}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracoes_variaveis(T3AlgumaParser.Declaracoes_variaveisContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#variavel}.
	 * @param ctx the parse tree
	 */
	void enterVariavel(T3AlgumaParser.VariavelContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#variavel}.
	 * @param ctx the parse tree
	 */
	void exitVariavel(T3AlgumaParser.VariavelContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#identificador}.
	 * @param ctx the parse tree
	 */
	void enterIdentificador(T3AlgumaParser.IdentificadorContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#identificador}.
	 * @param ctx the parse tree
	 */
	void exitIdentificador(T3AlgumaParser.IdentificadorContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#exp_aritmetica}.
	 * @param ctx the parse tree
	 */
	void enterExp_aritmetica(T3AlgumaParser.Exp_aritmeticaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#exp_aritmetica}.
	 * @param ctx the parse tree
	 */
	void exitExp_aritmetica(T3AlgumaParser.Exp_aritmeticaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#termo}.
	 * @param ctx the parse tree
	 */
	void enterTermo(T3AlgumaParser.TermoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#termo}.
	 * @param ctx the parse tree
	 */
	void exitTermo(T3AlgumaParser.TermoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#fator}.
	 * @param ctx the parse tree
	 */
	void enterFator(T3AlgumaParser.FatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#fator}.
	 * @param ctx the parse tree
	 */
	void exitFator(T3AlgumaParser.FatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#op1}.
	 * @param ctx the parse tree
	 */
	void enterOp1(T3AlgumaParser.Op1Context ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#op1}.
	 * @param ctx the parse tree
	 */
	void exitOp1(T3AlgumaParser.Op1Context ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#op2}.
	 * @param ctx the parse tree
	 */
	void enterOp2(T3AlgumaParser.Op2Context ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#op2}.
	 * @param ctx the parse tree
	 */
	void exitOp2(T3AlgumaParser.Op2Context ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#op3}.
	 * @param ctx the parse tree
	 */
	void enterOp3(T3AlgumaParser.Op3Context ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#op3}.
	 * @param ctx the parse tree
	 */
	void exitOp3(T3AlgumaParser.Op3Context ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#parcela}.
	 * @param ctx the parse tree
	 */
	void enterParcela(T3AlgumaParser.ParcelaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#parcela}.
	 * @param ctx the parse tree
	 */
	void exitParcela(T3AlgumaParser.ParcelaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#parcela_unario}.
	 * @param ctx the parse tree
	 */
	void enterParcela_unario(T3AlgumaParser.Parcela_unarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#parcela_unario}.
	 * @param ctx the parse tree
	 */
	void exitParcela_unario(T3AlgumaParser.Parcela_unarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#op_unario}.
	 * @param ctx the parse tree
	 */
	void enterOp_unario(T3AlgumaParser.Op_unarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#op_unario}.
	 * @param ctx the parse tree
	 */
	void exitOp_unario(T3AlgumaParser.Op_unarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#parcela_nao_unario}.
	 * @param ctx the parse tree
	 */
	void enterParcela_nao_unario(T3AlgumaParser.Parcela_nao_unarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#parcela_nao_unario}.
	 * @param ctx the parse tree
	 */
	void exitParcela_nao_unario(T3AlgumaParser.Parcela_nao_unarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo(T3AlgumaParser.TipoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo(T3AlgumaParser.TipoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#tipo_basico}.
	 * @param ctx the parse tree
	 */
	void enterTipo_basico(T3AlgumaParser.Tipo_basicoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#tipo_basico}.
	 * @param ctx the parse tree
	 */
	void exitTipo_basico(T3AlgumaParser.Tipo_basicoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#tipo_variavel}.
	 * @param ctx the parse tree
	 */
	void enterTipo_variavel(T3AlgumaParser.Tipo_variavelContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#tipo_variavel}.
	 * @param ctx the parse tree
	 */
	void exitTipo_variavel(T3AlgumaParser.Tipo_variavelContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#valor_constante}.
	 * @param ctx the parse tree
	 */
	void enterValor_constante(T3AlgumaParser.Valor_constanteContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#valor_constante}.
	 * @param ctx the parse tree
	 */
	void exitValor_constante(T3AlgumaParser.Valor_constanteContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#registro}.
	 * @param ctx the parse tree
	 */
	void enterRegistro(T3AlgumaParser.RegistroContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#registro}.
	 * @param ctx the parse tree
	 */
	void exitRegistro(T3AlgumaParser.RegistroContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#parametro}.
	 * @param ctx the parse tree
	 */
	void enterParametro(T3AlgumaParser.ParametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#parametro}.
	 * @param ctx the parse tree
	 */
	void exitParametro(T3AlgumaParser.ParametroContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#parametros}.
	 * @param ctx the parse tree
	 */
	void enterParametros(T3AlgumaParser.ParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#parametros}.
	 * @param ctx the parse tree
	 */
	void exitParametros(T3AlgumaParser.ParametrosContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#declaracoes_funcoes}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracoes_funcoes(T3AlgumaParser.Declaracoes_funcoesContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#declaracoes_funcoes}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracoes_funcoes(T3AlgumaParser.Declaracoes_funcoesContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#corpo}.
	 * @param ctx the parse tree
	 */
	void enterCorpo(T3AlgumaParser.CorpoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#corpo}.
	 * @param ctx the parse tree
	 */
	void exitCorpo(T3AlgumaParser.CorpoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmd}.
	 * @param ctx the parse tree
	 */
	void enterCmd(T3AlgumaParser.CmdContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmd}.
	 * @param ctx the parse tree
	 */
	void exitCmd(T3AlgumaParser.CmdContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmdLeia}.
	 * @param ctx the parse tree
	 */
	void enterCmdLeia(T3AlgumaParser.CmdLeiaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmdLeia}.
	 * @param ctx the parse tree
	 */
	void exitCmdLeia(T3AlgumaParser.CmdLeiaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmdEscreva}.
	 * @param ctx the parse tree
	 */
	void enterCmdEscreva(T3AlgumaParser.CmdEscrevaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmdEscreva}.
	 * @param ctx the parse tree
	 */
	void exitCmdEscreva(T3AlgumaParser.CmdEscrevaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmdSe}.
	 * @param ctx the parse tree
	 */
	void enterCmdSe(T3AlgumaParser.CmdSeContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmdSe}.
	 * @param ctx the parse tree
	 */
	void exitCmdSe(T3AlgumaParser.CmdSeContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmdCaso}.
	 * @param ctx the parse tree
	 */
	void enterCmdCaso(T3AlgumaParser.CmdCasoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmdCaso}.
	 * @param ctx the parse tree
	 */
	void exitCmdCaso(T3AlgumaParser.CmdCasoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmdPara}.
	 * @param ctx the parse tree
	 */
	void enterCmdPara(T3AlgumaParser.CmdParaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmdPara}.
	 * @param ctx the parse tree
	 */
	void exitCmdPara(T3AlgumaParser.CmdParaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmdEnquanto}.
	 * @param ctx the parse tree
	 */
	void enterCmdEnquanto(T3AlgumaParser.CmdEnquantoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmdEnquanto}.
	 * @param ctx the parse tree
	 */
	void exitCmdEnquanto(T3AlgumaParser.CmdEnquantoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmdFaca}.
	 * @param ctx the parse tree
	 */
	void enterCmdFaca(T3AlgumaParser.CmdFacaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmdFaca}.
	 * @param ctx the parse tree
	 */
	void exitCmdFaca(T3AlgumaParser.CmdFacaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmdAtribuicao}.
	 * @param ctx the parse tree
	 */
	void enterCmdAtribuicao(T3AlgumaParser.CmdAtribuicaoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmdAtribuicao}.
	 * @param ctx the parse tree
	 */
	void exitCmdAtribuicao(T3AlgumaParser.CmdAtribuicaoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmdChamada}.
	 * @param ctx the parse tree
	 */
	void enterCmdChamada(T3AlgumaParser.CmdChamadaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmdChamada}.
	 * @param ctx the parse tree
	 */
	void exitCmdChamada(T3AlgumaParser.CmdChamadaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#cmdRetorne}.
	 * @param ctx the parse tree
	 */
	void enterCmdRetorne(T3AlgumaParser.CmdRetorneContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#cmdRetorne}.
	 * @param ctx the parse tree
	 */
	void exitCmdRetorne(T3AlgumaParser.CmdRetorneContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#selecao}.
	 * @param ctx the parse tree
	 */
	void enterSelecao(T3AlgumaParser.SelecaoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#selecao}.
	 * @param ctx the parse tree
	 */
	void exitSelecao(T3AlgumaParser.SelecaoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#item_selecao}.
	 * @param ctx the parse tree
	 */
	void enterItem_selecao(T3AlgumaParser.Item_selecaoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#item_selecao}.
	 * @param ctx the parse tree
	 */
	void exitItem_selecao(T3AlgumaParser.Item_selecaoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#constantes}.
	 * @param ctx the parse tree
	 */
	void enterConstantes(T3AlgumaParser.ConstantesContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#constantes}.
	 * @param ctx the parse tree
	 */
	void exitConstantes(T3AlgumaParser.ConstantesContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#numero_intervalo}.
	 * @param ctx the parse tree
	 */
	void enterNumero_intervalo(T3AlgumaParser.Numero_intervaloContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#numero_intervalo}.
	 * @param ctx the parse tree
	 */
	void exitNumero_intervalo(T3AlgumaParser.Numero_intervaloContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#exp_relacional}.
	 * @param ctx the parse tree
	 */
	void enterExp_relacional(T3AlgumaParser.Exp_relacionalContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#exp_relacional}.
	 * @param ctx the parse tree
	 */
	void exitExp_relacional(T3AlgumaParser.Exp_relacionalContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#op_relacional}.
	 * @param ctx the parse tree
	 */
	void enterOp_relacional(T3AlgumaParser.Op_relacionalContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#op_relacional}.
	 * @param ctx the parse tree
	 */
	void exitOp_relacional(T3AlgumaParser.Op_relacionalContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#expressao}.
	 * @param ctx the parse tree
	 */
	void enterExpressao(T3AlgumaParser.ExpressaoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#expressao}.
	 * @param ctx the parse tree
	 */
	void exitExpressao(T3AlgumaParser.ExpressaoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#termo_logico}.
	 * @param ctx the parse tree
	 */
	void enterTermo_logico(T3AlgumaParser.Termo_logicoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#termo_logico}.
	 * @param ctx the parse tree
	 */
	void exitTermo_logico(T3AlgumaParser.Termo_logicoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#fator_logico}.
	 * @param ctx the parse tree
	 */
	void enterFator_logico(T3AlgumaParser.Fator_logicoContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#fator_logico}.
	 * @param ctx the parse tree
	 */
	void exitFator_logico(T3AlgumaParser.Fator_logicoContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#parcela_logica}.
	 * @param ctx the parse tree
	 */
	void enterParcela_logica(T3AlgumaParser.Parcela_logicaContext ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#parcela_logica}.
	 * @param ctx the parse tree
	 */
	void exitParcela_logica(T3AlgumaParser.Parcela_logicaContext ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#op_logico_1}.
	 * @param ctx the parse tree
	 */
	void enterOp_logico_1(T3AlgumaParser.Op_logico_1Context ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#op_logico_1}.
	 * @param ctx the parse tree
	 */
	void exitOp_logico_1(T3AlgumaParser.Op_logico_1Context ctx);
	/**
	 * Enter a parse tree produced by {@link T3AlgumaParser#op_logico_2}.
	 * @param ctx the parse tree
	 */
	void enterOp_logico_2(T3AlgumaParser.Op_logico_2Context ctx);
	/**
	 * Exit a parse tree produced by {@link T3AlgumaParser#op_logico_2}.
	 * @param ctx the parse tree
	 */
	void exitOp_logico_2(T3AlgumaParser.Op_logico_2Context ctx);
}