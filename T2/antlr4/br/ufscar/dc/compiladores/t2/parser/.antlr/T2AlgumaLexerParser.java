// Generated from /home/rodrigo/UFSCAR/Compiladores/construcao-compiladores/T2/antlr4/br/ufscar/dc/compiladores/t2/lexico/T2AlgumaLexer.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class T2AlgumaLexerParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		COMENTARIO=1, ALGORITMO=2, DECLARE=3, VAR=4, CONSTANTE=5, LITERAL=6, INTEIRO=7, 
		REAL=8, LOGICO=9, VERDADEIRO=10, FALSO=11, E=12, OU=13, NAO=14, SE=15, 
		FIM_SE=16, ENTAO=17, SENAO=18, CASO=19, SEJA=20, FIM_CASO=21, PARA=22, 
		FIM_PARA=23, ATE=24, FACA=25, ENQUANTO=26, FIM_ENQUANTO=27, TIPO=28, REGISTRO=29, 
		FIM_REGISTRO=30, PROCEDIMENTO=31, FIM_PROCEDIMENTO=32, FUNCAO=33, FIM_FUNCAO=34, 
		RETORNE=35, LEIA=36, ESCREVA=37, FIM_ALGORITMO=38, INTERVALO=39, MENOR=40, 
		MENORIGUAL=41, MAIOR=42, MAIORIGUAL=43, IGUAL=44, DIFERENTE=45, DOISPONTOS=46, 
		ABREPAR=47, FECHAPAR=48, ABRECHAVE=49, FECHACHAVE=50, VIRGULA=51, ASPAS=52, 
		DIVISAO=53, RESTO=54, SOMA=55, SUBTRACAO=56, MULTIPLICACAO=57, ATRIBUICAO=58, 
		PONTEIRO=59, ENDERECO=60, PONTO=61, NUM_INT=62, NUM_REAL=63, IDENT=64, 
		CADEIA=65, ERRO_COMENTARIO_ABERTO=66, ERRO_CADEIA_ABERTA=67, WS=68, ERRO=69;
	public static final int
		RULE_programa = 0, RULE_declaracoes = 1, RULE_declaracoes_variaveis = 2, 
		RULE_variavel = 3, RULE_identificador = 4, RULE_exp_aritmetica = 5, RULE_termo = 6, 
		RULE_fator = 7, RULE_op1 = 8, RULE_op2 = 9, RULE_op3 = 10, RULE_parcela = 11, 
		RULE_parcela_unario = 12, RULE_op_unario = 13, RULE_parcela_nao_unario = 14, 
		RULE_tipo = 15, RULE_tipo_basico = 16, RULE_tipo_variavel = 17, RULE_valor_constante = 18, 
		RULE_registro = 19, RULE_parametro = 20, RULE_parametros = 21, RULE_declaracoes_funcoes = 22, 
		RULE_corpo = 23, RULE_cmd = 24, RULE_cmdLeia = 25, RULE_cmdEscreva = 26, 
		RULE_cmdSe = 27, RULE_cmdCaso = 28, RULE_cmdPara = 29, RULE_cmdEnquanto = 30, 
		RULE_cmdFaca = 31, RULE_cmdAtribuicao = 32, RULE_cmdChamada = 33, RULE_cmdRetorne = 34, 
		RULE_selecao = 35, RULE_item_selecao = 36, RULE_constantes = 37, RULE_numero_intervalo = 38, 
		RULE_exp_relacional = 39, RULE_op_relacional = 40, RULE_expressao = 41, 
		RULE_termo_logico = 42, RULE_fator_logico = 43, RULE_parcela_logica = 44, 
		RULE_op_logico_1 = 45, RULE_op_logico_2 = 46;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "declaracoes", "declaracoes_variaveis", "variavel", "identificador", 
			"exp_aritmetica", "termo", "fator", "op1", "op2", "op3", "parcela", "parcela_unario", 
			"op_unario", "parcela_nao_unario", "tipo", "tipo_basico", "tipo_variavel", 
			"valor_constante", "registro", "parametro", "parametros", "declaracoes_funcoes", 
			"corpo", "cmd", "cmdLeia", "cmdEscreva", "cmdSe", "cmdCaso", "cmdPara", 
			"cmdEnquanto", "cmdFaca", "cmdAtribuicao", "cmdChamada", "cmdRetorne", 
			"selecao", "item_selecao", "constantes", "numero_intervalo", "exp_relacional", 
			"op_relacional", "expressao", "termo_logico", "fator_logico", "parcela_logica", 
			"op_logico_1", "op_logico_2"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'algoritmo'", "'declare'", "'var'", "'constante'", "'literal'", 
			"'inteiro'", "'real'", "'logico'", "'verdadeiro'", "'falso'", "'e'", 
			"'ou'", "'nao'", "'se'", "'fim_se'", "'entao'", "'senao'", "'caso'", 
			"'seja'", "'fim_caso'", "'para'", "'fim_para'", "'ate'", "'faca'", "'enquanto'", 
			"'fim_enquanto'", "'tipo'", "'registro'", "'fim_registro'", "'procedimento'", 
			"'fim_procedimento'", "'funcao'", "'fim_funcao'", "'retorne'", "'leia'", 
			"'escreva'", "'fim_algoritmo'", "'..'", "'<'", "'<='", "'>'", "'>='", 
			"'='", "'<>'", "':'", "'('", "')'", "'['", "']'", "','", "'\"'", "'/'", 
			"'%'", "'+'", "'-'", "'*'", "'<-'", "'^'", "'&'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "COMENTARIO", "ALGORITMO", "DECLARE", "VAR", "CONSTANTE", "LITERAL", 
			"INTEIRO", "REAL", "LOGICO", "VERDADEIRO", "FALSO", "E", "OU", "NAO", 
			"SE", "FIM_SE", "ENTAO", "SENAO", "CASO", "SEJA", "FIM_CASO", "PARA", 
			"FIM_PARA", "ATE", "FACA", "ENQUANTO", "FIM_ENQUANTO", "TIPO", "REGISTRO", 
			"FIM_REGISTRO", "PROCEDIMENTO", "FIM_PROCEDIMENTO", "FUNCAO", "FIM_FUNCAO", 
			"RETORNE", "LEIA", "ESCREVA", "FIM_ALGORITMO", "INTERVALO", "MENOR", 
			"MENORIGUAL", "MAIOR", "MAIORIGUAL", "IGUAL", "DIFERENTE", "DOISPONTOS", 
			"ABREPAR", "FECHAPAR", "ABRECHAVE", "FECHACHAVE", "VIRGULA", "ASPAS", 
			"DIVISAO", "RESTO", "SOMA", "SUBTRACAO", "MULTIPLICACAO", "ATRIBUICAO", 
			"PONTEIRO", "ENDERECO", "PONTO", "NUM_INT", "NUM_REAL", "IDENT", "CADEIA", 
			"ERRO_COMENTARIO_ABERTO", "ERRO_CADEIA_ABERTA", "WS", "ERRO"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "T2AlgumaLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public T2AlgumaLexerParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramaContext extends ParserRuleContext {
		public DeclaracoesContext declaracoes() {
			return getRuleContext(DeclaracoesContext.class,0);
		}
		public TerminalNode ALGORITMO() { return getToken(T2AlgumaLexerParser.ALGORITMO, 0); }
		public CorpoContext corpo() {
			return getRuleContext(CorpoContext.class,0);
		}
		public TerminalNode FIM_ALGORITMO() { return getToken(T2AlgumaLexerParser.FIM_ALGORITMO, 0); }
		public ProgramaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_programa; }
	}

	public final ProgramaContext programa() throws RecognitionException {
		ProgramaContext _localctx = new ProgramaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_programa);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			declaracoes();
			setState(95);
			match(ALGORITMO);
			setState(96);
			corpo();
			setState(97);
			match(FIM_ALGORITMO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclaracoesContext extends ParserRuleContext {
		public Declaracoes_variaveisContext declaracoes_variaveis() {
			return getRuleContext(Declaracoes_variaveisContext.class,0);
		}
		public Declaracoes_funcoesContext declaracoes_funcoes() {
			return getRuleContext(Declaracoes_funcoesContext.class,0);
		}
		public DeclaracoesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracoes; }
	}

	public final DeclaracoesContext declaracoes() throws RecognitionException {
		DeclaracoesContext _localctx = new DeclaracoesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_declaracoes);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(101);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DECLARE:
			case CONSTANTE:
			case TIPO:
				{
				setState(99);
				declaracoes_variaveis();
				}
				break;
			case PROCEDIMENTO:
			case FUNCAO:
				{
				setState(100);
				declaracoes_funcoes();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Declaracoes_variaveisContext extends ParserRuleContext {
		public TerminalNode DECLARE() { return getToken(T2AlgumaLexerParser.DECLARE, 0); }
		public VariavelContext variavel() {
			return getRuleContext(VariavelContext.class,0);
		}
		public TerminalNode CONSTANTE() { return getToken(T2AlgumaLexerParser.CONSTANTE, 0); }
		public TerminalNode IDENT() { return getToken(T2AlgumaLexerParser.IDENT, 0); }
		public TerminalNode DOISPONTOS() { return getToken(T2AlgumaLexerParser.DOISPONTOS, 0); }
		public Tipo_basicoContext tipo_basico() {
			return getRuleContext(Tipo_basicoContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(T2AlgumaLexerParser.IGUAL, 0); }
		public Valor_constanteContext valor_constante() {
			return getRuleContext(Valor_constanteContext.class,0);
		}
		public TerminalNode TIPO() { return getToken(T2AlgumaLexerParser.TIPO, 0); }
		public RegistroContext registro() {
			return getRuleContext(RegistroContext.class,0);
		}
		public Declaracoes_variaveisContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracoes_variaveis; }
	}

	public final Declaracoes_variaveisContext declaracoes_variaveis() throws RecognitionException {
		Declaracoes_variaveisContext _localctx = new Declaracoes_variaveisContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_declaracoes_variaveis);
		try {
			setState(116);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DECLARE:
				enterOuterAlt(_localctx, 1);
				{
				setState(103);
				match(DECLARE);
				setState(104);
				variavel();
				}
				break;
			case CONSTANTE:
				enterOuterAlt(_localctx, 2);
				{
				setState(105);
				match(CONSTANTE);
				setState(106);
				match(IDENT);
				setState(107);
				match(DOISPONTOS);
				setState(108);
				tipo_basico();
				setState(109);
				match(IGUAL);
				setState(110);
				valor_constante();
				}
				break;
			case TIPO:
				enterOuterAlt(_localctx, 3);
				{
				setState(112);
				match(TIPO);
				setState(113);
				match(IDENT);
				setState(114);
				match(DOISPONTOS);
				setState(115);
				registro();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VariavelContext extends ParserRuleContext {
		public List<IdentificadorContext> identificador() {
			return getRuleContexts(IdentificadorContext.class);
		}
		public IdentificadorContext identificador(int i) {
			return getRuleContext(IdentificadorContext.class,i);
		}
		public TerminalNode DOISPONTOS() { return getToken(T2AlgumaLexerParser.DOISPONTOS, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public List<TerminalNode> VIRGULA() { return getTokens(T2AlgumaLexerParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T2AlgumaLexerParser.VIRGULA, i);
		}
		public VariavelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variavel; }
	}

	public final VariavelContext variavel() throws RecognitionException {
		VariavelContext _localctx = new VariavelContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_variavel);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(118);
			identificador();
			setState(123);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(119);
				match(VIRGULA);
				setState(120);
				identificador();
				}
				}
				setState(125);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(126);
			match(DOISPONTOS);
			setState(127);
			tipo();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentificadorContext extends ParserRuleContext {
		public List<TerminalNode> IDENT() { return getTokens(T2AlgumaLexerParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(T2AlgumaLexerParser.IDENT, i);
		}
		public List<TerminalNode> PONTO() { return getTokens(T2AlgumaLexerParser.PONTO); }
		public TerminalNode PONTO(int i) {
			return getToken(T2AlgumaLexerParser.PONTO, i);
		}
		public List<TerminalNode> ABRECHAVE() { return getTokens(T2AlgumaLexerParser.ABRECHAVE); }
		public TerminalNode ABRECHAVE(int i) {
			return getToken(T2AlgumaLexerParser.ABRECHAVE, i);
		}
		public List<Exp_aritmeticaContext> exp_aritmetica() {
			return getRuleContexts(Exp_aritmeticaContext.class);
		}
		public Exp_aritmeticaContext exp_aritmetica(int i) {
			return getRuleContext(Exp_aritmeticaContext.class,i);
		}
		public List<TerminalNode> FECHACHAVE() { return getTokens(T2AlgumaLexerParser.FECHACHAVE); }
		public TerminalNode FECHACHAVE(int i) {
			return getToken(T2AlgumaLexerParser.FECHACHAVE, i);
		}
		public IdentificadorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identificador; }
	}

	public final IdentificadorContext identificador() throws RecognitionException {
		IdentificadorContext _localctx = new IdentificadorContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_identificador);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
			match(IDENT);
			setState(134);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PONTO) {
				{
				{
				setState(130);
				match(PONTO);
				setState(131);
				match(IDENT);
				}
				}
				setState(136);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(143);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ABRECHAVE) {
				{
				{
				setState(137);
				match(ABRECHAVE);
				setState(138);
				exp_aritmetica();
				setState(139);
				match(FECHACHAVE);
				}
				}
				setState(145);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Exp_aritmeticaContext extends ParserRuleContext {
		public List<TermoContext> termo() {
			return getRuleContexts(TermoContext.class);
		}
		public TermoContext termo(int i) {
			return getRuleContext(TermoContext.class,i);
		}
		public List<Op1Context> op1() {
			return getRuleContexts(Op1Context.class);
		}
		public Op1Context op1(int i) {
			return getRuleContext(Op1Context.class,i);
		}
		public Exp_aritmeticaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exp_aritmetica; }
	}

	public final Exp_aritmeticaContext exp_aritmetica() throws RecognitionException {
		Exp_aritmeticaContext _localctx = new Exp_aritmeticaContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_exp_aritmetica);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(146);
			termo();
			setState(152);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(147);
					op1();
					setState(148);
					termo();
					}
					} 
				}
				setState(154);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TermoContext extends ParserRuleContext {
		public List<FatorContext> fator() {
			return getRuleContexts(FatorContext.class);
		}
		public FatorContext fator(int i) {
			return getRuleContext(FatorContext.class,i);
		}
		public List<Op2Context> op2() {
			return getRuleContexts(Op2Context.class);
		}
		public Op2Context op2(int i) {
			return getRuleContext(Op2Context.class,i);
		}
		public TermoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_termo; }
	}

	public final TermoContext termo() throws RecognitionException {
		TermoContext _localctx = new TermoContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_termo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(155);
			fator();
			setState(161);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DIVISAO || _la==MULTIPLICACAO) {
				{
				{
				setState(156);
				op2();
				setState(157);
				fator();
				}
				}
				setState(163);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FatorContext extends ParserRuleContext {
		public List<ParcelaContext> parcela() {
			return getRuleContexts(ParcelaContext.class);
		}
		public ParcelaContext parcela(int i) {
			return getRuleContext(ParcelaContext.class,i);
		}
		public List<Op3Context> op3() {
			return getRuleContexts(Op3Context.class);
		}
		public Op3Context op3(int i) {
			return getRuleContext(Op3Context.class,i);
		}
		public FatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fator; }
	}

	public final FatorContext fator() throws RecognitionException {
		FatorContext _localctx = new FatorContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_fator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(164);
			parcela();
			setState(170);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==RESTO) {
				{
				{
				setState(165);
				op3();
				setState(166);
				parcela();
				}
				}
				setState(172);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Op1Context extends ParserRuleContext {
		public TerminalNode SOMA() { return getToken(T2AlgumaLexerParser.SOMA, 0); }
		public TerminalNode SUBTRACAO() { return getToken(T2AlgumaLexerParser.SUBTRACAO, 0); }
		public Op1Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op1; }
	}

	public final Op1Context op1() throws RecognitionException {
		Op1Context _localctx = new Op1Context(_ctx, getState());
		enterRule(_localctx, 16, RULE_op1);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(173);
			_la = _input.LA(1);
			if ( !(_la==SOMA || _la==SUBTRACAO) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Op2Context extends ParserRuleContext {
		public TerminalNode MULTIPLICACAO() { return getToken(T2AlgumaLexerParser.MULTIPLICACAO, 0); }
		public TerminalNode DIVISAO() { return getToken(T2AlgumaLexerParser.DIVISAO, 0); }
		public Op2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op2; }
	}

	public final Op2Context op2() throws RecognitionException {
		Op2Context _localctx = new Op2Context(_ctx, getState());
		enterRule(_localctx, 18, RULE_op2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			_la = _input.LA(1);
			if ( !(_la==DIVISAO || _la==MULTIPLICACAO) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Op3Context extends ParserRuleContext {
		public TerminalNode RESTO() { return getToken(T2AlgumaLexerParser.RESTO, 0); }
		public Op3Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op3; }
	}

	public final Op3Context op3() throws RecognitionException {
		Op3Context _localctx = new Op3Context(_ctx, getState());
		enterRule(_localctx, 20, RULE_op3);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			match(RESTO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParcelaContext extends ParserRuleContext {
		public Parcela_unarioContext parcela_unario() {
			return getRuleContext(Parcela_unarioContext.class,0);
		}
		public Op_unarioContext op_unario() {
			return getRuleContext(Op_unarioContext.class,0);
		}
		public Parcela_nao_unarioContext parcela_nao_unario() {
			return getRuleContext(Parcela_nao_unarioContext.class,0);
		}
		public ParcelaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parcela; }
	}

	public final ParcelaContext parcela() throws RecognitionException {
		ParcelaContext _localctx = new ParcelaContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_parcela);
		int _la;
		try {
			setState(184);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ABREPAR:
			case SUBTRACAO:
			case PONTEIRO:
			case NUM_INT:
			case NUM_REAL:
			case IDENT:
				enterOuterAlt(_localctx, 1);
				{
				setState(180);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SUBTRACAO) {
					{
					setState(179);
					op_unario();
					}
				}

				setState(182);
				parcela_unario();
				}
				break;
			case ENDERECO:
			case CADEIA:
				enterOuterAlt(_localctx, 2);
				{
				setState(183);
				parcela_nao_unario();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Parcela_unarioContext extends ParserRuleContext {
		public IdentificadorContext identificador() {
			return getRuleContext(IdentificadorContext.class,0);
		}
		public TerminalNode PONTEIRO() { return getToken(T2AlgumaLexerParser.PONTEIRO, 0); }
		public TerminalNode IDENT() { return getToken(T2AlgumaLexerParser.IDENT, 0); }
		public TerminalNode ABREPAR() { return getToken(T2AlgumaLexerParser.ABREPAR, 0); }
		public List<ExpressaoContext> expressao() {
			return getRuleContexts(ExpressaoContext.class);
		}
		public ExpressaoContext expressao(int i) {
			return getRuleContext(ExpressaoContext.class,i);
		}
		public TerminalNode FECHAPAR() { return getToken(T2AlgumaLexerParser.FECHAPAR, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(T2AlgumaLexerParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T2AlgumaLexerParser.VIRGULA, i);
		}
		public TerminalNode NUM_INT() { return getToken(T2AlgumaLexerParser.NUM_INT, 0); }
		public TerminalNode NUM_REAL() { return getToken(T2AlgumaLexerParser.NUM_REAL, 0); }
		public Parcela_unarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parcela_unario; }
	}

	public final Parcela_unarioContext parcela_unario() throws RecognitionException {
		Parcela_unarioContext _localctx = new Parcela_unarioContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_parcela_unario);
		int _la;
		try {
			setState(208);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(187);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PONTEIRO) {
					{
					setState(186);
					match(PONTEIRO);
					}
				}

				setState(189);
				identificador();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(190);
				match(IDENT);
				setState(191);
				match(ABREPAR);
				setState(192);
				expressao();
				setState(197);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==VIRGULA) {
					{
					{
					setState(193);
					match(VIRGULA);
					setState(194);
					expressao();
					}
					}
					setState(199);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(200);
				match(FECHAPAR);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(202);
				match(NUM_INT);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(203);
				match(NUM_REAL);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(204);
				match(ABREPAR);
				setState(205);
				expressao();
				setState(206);
				match(FECHAPAR);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Op_unarioContext extends ParserRuleContext {
		public TerminalNode SUBTRACAO() { return getToken(T2AlgumaLexerParser.SUBTRACAO, 0); }
		public Op_unarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_unario; }
	}

	public final Op_unarioContext op_unario() throws RecognitionException {
		Op_unarioContext _localctx = new Op_unarioContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_op_unario);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(210);
			match(SUBTRACAO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Parcela_nao_unarioContext extends ParserRuleContext {
		public TerminalNode ENDERECO() { return getToken(T2AlgumaLexerParser.ENDERECO, 0); }
		public IdentificadorContext identificador() {
			return getRuleContext(IdentificadorContext.class,0);
		}
		public TerminalNode CADEIA() { return getToken(T2AlgumaLexerParser.CADEIA, 0); }
		public Parcela_nao_unarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parcela_nao_unario; }
	}

	public final Parcela_nao_unarioContext parcela_nao_unario() throws RecognitionException {
		Parcela_nao_unarioContext _localctx = new Parcela_nao_unarioContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_parcela_nao_unario);
		try {
			setState(215);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENDERECO:
				enterOuterAlt(_localctx, 1);
				{
				setState(212);
				match(ENDERECO);
				setState(213);
				identificador();
				}
				break;
			case CADEIA:
				enterOuterAlt(_localctx, 2);
				{
				setState(214);
				match(CADEIA);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TipoContext extends ParserRuleContext {
		public RegistroContext registro() {
			return getRuleContext(RegistroContext.class,0);
		}
		public Tipo_variavelContext tipo_variavel() {
			return getRuleContext(Tipo_variavelContext.class,0);
		}
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_tipo);
		try {
			setState(219);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case REGISTRO:
				enterOuterAlt(_localctx, 1);
				{
				setState(217);
				registro();
				}
				break;
			case LITERAL:
			case INTEIRO:
			case REAL:
			case LOGICO:
			case PONTEIRO:
			case IDENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(218);
				tipo_variavel();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_basicoContext extends ParserRuleContext {
		public TerminalNode LITERAL() { return getToken(T2AlgumaLexerParser.LITERAL, 0); }
		public TerminalNode INTEIRO() { return getToken(T2AlgumaLexerParser.INTEIRO, 0); }
		public TerminalNode REAL() { return getToken(T2AlgumaLexerParser.REAL, 0); }
		public TerminalNode LOGICO() { return getToken(T2AlgumaLexerParser.LOGICO, 0); }
		public Tipo_basicoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_basico; }
	}

	public final Tipo_basicoContext tipo_basico() throws RecognitionException {
		Tipo_basicoContext _localctx = new Tipo_basicoContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_tipo_basico);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(221);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 960L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_variavelContext extends ParserRuleContext {
		public Tipo_basicoContext tipo_basico() {
			return getRuleContext(Tipo_basicoContext.class,0);
		}
		public TerminalNode IDENT() { return getToken(T2AlgumaLexerParser.IDENT, 0); }
		public TerminalNode PONTEIRO() { return getToken(T2AlgumaLexerParser.PONTEIRO, 0); }
		public Tipo_variavelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_variavel; }
	}

	public final Tipo_variavelContext tipo_variavel() throws RecognitionException {
		Tipo_variavelContext _localctx = new Tipo_variavelContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_tipo_variavel);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(224);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PONTEIRO) {
				{
				setState(223);
				match(PONTEIRO);
				}
			}

			setState(228);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LITERAL:
			case INTEIRO:
			case REAL:
			case LOGICO:
				{
				setState(226);
				tipo_basico();
				}
				break;
			case IDENT:
				{
				setState(227);
				match(IDENT);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Valor_constanteContext extends ParserRuleContext {
		public TerminalNode CADEIA() { return getToken(T2AlgumaLexerParser.CADEIA, 0); }
		public TerminalNode NUM_INT() { return getToken(T2AlgumaLexerParser.NUM_INT, 0); }
		public TerminalNode NUM_REAL() { return getToken(T2AlgumaLexerParser.NUM_REAL, 0); }
		public TerminalNode VERDADEIRO() { return getToken(T2AlgumaLexerParser.VERDADEIRO, 0); }
		public TerminalNode FALSO() { return getToken(T2AlgumaLexerParser.FALSO, 0); }
		public Valor_constanteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valor_constante; }
	}

	public final Valor_constanteContext valor_constante() throws RecognitionException {
		Valor_constanteContext _localctx = new Valor_constanteContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_valor_constante);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(230);
			_la = _input.LA(1);
			if ( !(((((_la - 10)) & ~0x3f) == 0 && ((1L << (_la - 10)) & 49539595901075459L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RegistroContext extends ParserRuleContext {
		public TerminalNode REGISTRO() { return getToken(T2AlgumaLexerParser.REGISTRO, 0); }
		public TerminalNode FIM_REGISTRO() { return getToken(T2AlgumaLexerParser.FIM_REGISTRO, 0); }
		public List<VariavelContext> variavel() {
			return getRuleContexts(VariavelContext.class);
		}
		public VariavelContext variavel(int i) {
			return getRuleContext(VariavelContext.class,i);
		}
		public RegistroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_registro; }
	}

	public final RegistroContext registro() throws RecognitionException {
		RegistroContext _localctx = new RegistroContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_registro);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(232);
			match(REGISTRO);
			setState(236);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENT) {
				{
				{
				setState(233);
				variavel();
				}
				}
				setState(238);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(239);
			match(FIM_REGISTRO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametroContext extends ParserRuleContext {
		public List<IdentificadorContext> identificador() {
			return getRuleContexts(IdentificadorContext.class);
		}
		public IdentificadorContext identificador(int i) {
			return getRuleContext(IdentificadorContext.class,i);
		}
		public TerminalNode DOISPONTOS() { return getToken(T2AlgumaLexerParser.DOISPONTOS, 0); }
		public Tipo_variavelContext tipo_variavel() {
			return getRuleContext(Tipo_variavelContext.class,0);
		}
		public TerminalNode VAR() { return getToken(T2AlgumaLexerParser.VAR, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(T2AlgumaLexerParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T2AlgumaLexerParser.VIRGULA, i);
		}
		public ParametroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametro; }
	}

	public final ParametroContext parametro() throws RecognitionException {
		ParametroContext _localctx = new ParametroContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_parametro);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(242);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==VAR) {
				{
				setState(241);
				match(VAR);
				}
			}

			setState(244);
			identificador();
			setState(249);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(245);
				match(VIRGULA);
				setState(246);
				identificador();
				}
				}
				setState(251);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(252);
			match(DOISPONTOS);
			setState(253);
			tipo_variavel();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametrosContext extends ParserRuleContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public List<TerminalNode> VIRGULA() { return getTokens(T2AlgumaLexerParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T2AlgumaLexerParser.VIRGULA, i);
		}
		public ParametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametros; }
	}

	public final ParametrosContext parametros() throws RecognitionException {
		ParametrosContext _localctx = new ParametrosContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(255);
			parametro();
			setState(260);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(256);
				match(VIRGULA);
				setState(257);
				parametro();
				}
				}
				setState(262);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Declaracoes_funcoesContext extends ParserRuleContext {
		public TerminalNode PROCEDIMENTO() { return getToken(T2AlgumaLexerParser.PROCEDIMENTO, 0); }
		public TerminalNode IDENT() { return getToken(T2AlgumaLexerParser.IDENT, 0); }
		public TerminalNode ABREPAR() { return getToken(T2AlgumaLexerParser.ABREPAR, 0); }
		public TerminalNode FECHAPAR() { return getToken(T2AlgumaLexerParser.FECHAPAR, 0); }
		public TerminalNode FIM_PROCEDIMENTO() { return getToken(T2AlgumaLexerParser.FIM_PROCEDIMENTO, 0); }
		public ParametrosContext parametros() {
			return getRuleContext(ParametrosContext.class,0);
		}
		public List<Declaracoes_variaveisContext> declaracoes_variaveis() {
			return getRuleContexts(Declaracoes_variaveisContext.class);
		}
		public Declaracoes_variaveisContext declaracoes_variaveis(int i) {
			return getRuleContext(Declaracoes_variaveisContext.class,i);
		}
		public List<CmdContext> cmd() {
			return getRuleContexts(CmdContext.class);
		}
		public CmdContext cmd(int i) {
			return getRuleContext(CmdContext.class,i);
		}
		public TerminalNode FUNCAO() { return getToken(T2AlgumaLexerParser.FUNCAO, 0); }
		public TerminalNode DOISPONTOS() { return getToken(T2AlgumaLexerParser.DOISPONTOS, 0); }
		public Tipo_variavelContext tipo_variavel() {
			return getRuleContext(Tipo_variavelContext.class,0);
		}
		public TerminalNode FIM_FUNCAO() { return getToken(T2AlgumaLexerParser.FIM_FUNCAO, 0); }
		public Declaracoes_funcoesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracoes_funcoes; }
	}

	public final Declaracoes_funcoesContext declaracoes_funcoes() throws RecognitionException {
		Declaracoes_funcoesContext _localctx = new Declaracoes_funcoesContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_declaracoes_funcoes);
		int _la;
		try {
			setState(306);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PROCEDIMENTO:
				enterOuterAlt(_localctx, 1);
				{
				setState(263);
				match(PROCEDIMENTO);
				setState(264);
				match(IDENT);
				setState(265);
				match(ABREPAR);
				setState(267);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==VAR || _la==IDENT) {
					{
					setState(266);
					parametros();
					}
				}

				setState(269);
				match(FECHAPAR);
				setState(273);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 268435496L) != 0)) {
					{
					{
					setState(270);
					declaracoes_variaveis();
					}
					}
					setState(275);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(279);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
					{
					{
					setState(276);
					cmd();
					}
					}
					setState(281);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(282);
				match(FIM_PROCEDIMENTO);
				}
				break;
			case FUNCAO:
				enterOuterAlt(_localctx, 2);
				{
				setState(283);
				match(FUNCAO);
				setState(284);
				match(IDENT);
				setState(285);
				match(ABREPAR);
				setState(287);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==VAR || _la==IDENT) {
					{
					setState(286);
					parametros();
					}
				}

				setState(289);
				match(FECHAPAR);
				setState(290);
				match(DOISPONTOS);
				setState(291);
				tipo_variavel();
				setState(295);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 268435496L) != 0)) {
					{
					{
					setState(292);
					declaracoes_variaveis();
					}
					}
					setState(297);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(301);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
					{
					{
					setState(298);
					cmd();
					}
					}
					setState(303);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(304);
				match(FIM_FUNCAO);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CorpoContext extends ParserRuleContext {
		public List<Declaracoes_variaveisContext> declaracoes_variaveis() {
			return getRuleContexts(Declaracoes_variaveisContext.class);
		}
		public Declaracoes_variaveisContext declaracoes_variaveis(int i) {
			return getRuleContext(Declaracoes_variaveisContext.class,i);
		}
		public List<CmdContext> cmd() {
			return getRuleContexts(CmdContext.class);
		}
		public CmdContext cmd(int i) {
			return getRuleContext(CmdContext.class,i);
		}
		public CorpoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_corpo; }
	}

	public final CorpoContext corpo() throws RecognitionException {
		CorpoContext _localctx = new CorpoContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_corpo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(311);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 268435496L) != 0)) {
				{
				{
				setState(308);
				declaracoes_variaveis();
				}
				}
				setState(313);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(317);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(314);
				cmd();
				}
				}
				setState(319);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdContext extends ParserRuleContext {
		public CmdLeiaContext cmdLeia() {
			return getRuleContext(CmdLeiaContext.class,0);
		}
		public CmdEscrevaContext cmdEscreva() {
			return getRuleContext(CmdEscrevaContext.class,0);
		}
		public CmdSeContext cmdSe() {
			return getRuleContext(CmdSeContext.class,0);
		}
		public CmdCasoContext cmdCaso() {
			return getRuleContext(CmdCasoContext.class,0);
		}
		public CmdParaContext cmdPara() {
			return getRuleContext(CmdParaContext.class,0);
		}
		public CmdEnquantoContext cmdEnquanto() {
			return getRuleContext(CmdEnquantoContext.class,0);
		}
		public CmdFacaContext cmdFaca() {
			return getRuleContext(CmdFacaContext.class,0);
		}
		public CmdAtribuicaoContext cmdAtribuicao() {
			return getRuleContext(CmdAtribuicaoContext.class,0);
		}
		public CmdChamadaContext cmdChamada() {
			return getRuleContext(CmdChamadaContext.class,0);
		}
		public CmdRetorneContext cmdRetorne() {
			return getRuleContext(CmdRetorneContext.class,0);
		}
		public CmdContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmd; }
	}

	public final CmdContext cmd() throws RecognitionException {
		CmdContext _localctx = new CmdContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_cmd);
		try {
			setState(330);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(320);
				cmdLeia();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(321);
				cmdEscreva();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(322);
				cmdSe();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(323);
				cmdCaso();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(324);
				cmdPara();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(325);
				cmdEnquanto();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(326);
				cmdFaca();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(327);
				cmdAtribuicao();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(328);
				cmdChamada();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(329);
				cmdRetorne();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdLeiaContext extends ParserRuleContext {
		public TerminalNode LEIA() { return getToken(T2AlgumaLexerParser.LEIA, 0); }
		public TerminalNode ABREPAR() { return getToken(T2AlgumaLexerParser.ABREPAR, 0); }
		public List<IdentificadorContext> identificador() {
			return getRuleContexts(IdentificadorContext.class);
		}
		public IdentificadorContext identificador(int i) {
			return getRuleContext(IdentificadorContext.class,i);
		}
		public TerminalNode FECHAPAR() { return getToken(T2AlgumaLexerParser.FECHAPAR, 0); }
		public List<TerminalNode> PONTEIRO() { return getTokens(T2AlgumaLexerParser.PONTEIRO); }
		public TerminalNode PONTEIRO(int i) {
			return getToken(T2AlgumaLexerParser.PONTEIRO, i);
		}
		public List<TerminalNode> VIRGULA() { return getTokens(T2AlgumaLexerParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T2AlgumaLexerParser.VIRGULA, i);
		}
		public CmdLeiaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdLeia; }
	}

	public final CmdLeiaContext cmdLeia() throws RecognitionException {
		CmdLeiaContext _localctx = new CmdLeiaContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_cmdLeia);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(332);
			match(LEIA);
			setState(333);
			match(ABREPAR);
			setState(335);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PONTEIRO) {
				{
				setState(334);
				match(PONTEIRO);
				}
			}

			setState(337);
			identificador();
			setState(345);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(338);
				match(VIRGULA);
				setState(340);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PONTEIRO) {
					{
					setState(339);
					match(PONTEIRO);
					}
				}

				setState(342);
				identificador();
				}
				}
				setState(347);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(348);
			match(FECHAPAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdEscrevaContext extends ParserRuleContext {
		public TerminalNode ESCREVA() { return getToken(T2AlgumaLexerParser.ESCREVA, 0); }
		public TerminalNode ABREPAR() { return getToken(T2AlgumaLexerParser.ABREPAR, 0); }
		public List<ExpressaoContext> expressao() {
			return getRuleContexts(ExpressaoContext.class);
		}
		public ExpressaoContext expressao(int i) {
			return getRuleContext(ExpressaoContext.class,i);
		}
		public TerminalNode FECHAPAR() { return getToken(T2AlgumaLexerParser.FECHAPAR, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(T2AlgumaLexerParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T2AlgumaLexerParser.VIRGULA, i);
		}
		public CmdEscrevaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdEscreva; }
	}

	public final CmdEscrevaContext cmdEscreva() throws RecognitionException {
		CmdEscrevaContext _localctx = new CmdEscrevaContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_cmdEscreva);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(350);
			match(ESCREVA);
			setState(351);
			match(ABREPAR);
			setState(352);
			expressao();
			setState(357);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(353);
				match(VIRGULA);
				setState(354);
				expressao();
				}
				}
				setState(359);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(360);
			match(FECHAPAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdSeContext extends ParserRuleContext {
		public TerminalNode SE() { return getToken(T2AlgumaLexerParser.SE, 0); }
		public ExpressaoContext expressao() {
			return getRuleContext(ExpressaoContext.class,0);
		}
		public TerminalNode ENTAO() { return getToken(T2AlgumaLexerParser.ENTAO, 0); }
		public TerminalNode FIM_SE() { return getToken(T2AlgumaLexerParser.FIM_SE, 0); }
		public List<CmdContext> cmd() {
			return getRuleContexts(CmdContext.class);
		}
		public CmdContext cmd(int i) {
			return getRuleContext(CmdContext.class,i);
		}
		public TerminalNode SENAO() { return getToken(T2AlgumaLexerParser.SENAO, 0); }
		public CmdSeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdSe; }
	}

	public final CmdSeContext cmdSe() throws RecognitionException {
		CmdSeContext _localctx = new CmdSeContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_cmdSe);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(362);
			match(SE);
			setState(363);
			expressao();
			setState(364);
			match(ENTAO);
			setState(368);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(365);
				cmd();
				}
				}
				setState(370);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(378);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SENAO) {
				{
				setState(371);
				match(SENAO);
				setState(375);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
					{
					{
					setState(372);
					cmd();
					}
					}
					setState(377);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(380);
			match(FIM_SE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdCasoContext extends ParserRuleContext {
		public TerminalNode CASO() { return getToken(T2AlgumaLexerParser.CASO, 0); }
		public Exp_aritmeticaContext exp_aritmetica() {
			return getRuleContext(Exp_aritmeticaContext.class,0);
		}
		public TerminalNode SEJA() { return getToken(T2AlgumaLexerParser.SEJA, 0); }
		public SelecaoContext selecao() {
			return getRuleContext(SelecaoContext.class,0);
		}
		public TerminalNode FIM_CASO() { return getToken(T2AlgumaLexerParser.FIM_CASO, 0); }
		public TerminalNode SENAO() { return getToken(T2AlgumaLexerParser.SENAO, 0); }
		public List<CmdContext> cmd() {
			return getRuleContexts(CmdContext.class);
		}
		public CmdContext cmd(int i) {
			return getRuleContext(CmdContext.class,i);
		}
		public CmdCasoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdCaso; }
	}

	public final CmdCasoContext cmdCaso() throws RecognitionException {
		CmdCasoContext _localctx = new CmdCasoContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_cmdCaso);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(382);
			match(CASO);
			setState(383);
			exp_aritmetica();
			setState(384);
			match(SEJA);
			setState(385);
			selecao();
			setState(393);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SENAO) {
				{
				setState(386);
				match(SENAO);
				setState(390);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
					{
					{
					setState(387);
					cmd();
					}
					}
					setState(392);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(395);
			match(FIM_CASO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdParaContext extends ParserRuleContext {
		public TerminalNode PARA() { return getToken(T2AlgumaLexerParser.PARA, 0); }
		public TerminalNode IDENT() { return getToken(T2AlgumaLexerParser.IDENT, 0); }
		public TerminalNode ATRIBUICAO() { return getToken(T2AlgumaLexerParser.ATRIBUICAO, 0); }
		public List<Exp_aritmeticaContext> exp_aritmetica() {
			return getRuleContexts(Exp_aritmeticaContext.class);
		}
		public Exp_aritmeticaContext exp_aritmetica(int i) {
			return getRuleContext(Exp_aritmeticaContext.class,i);
		}
		public TerminalNode ATE() { return getToken(T2AlgumaLexerParser.ATE, 0); }
		public TerminalNode FACA() { return getToken(T2AlgumaLexerParser.FACA, 0); }
		public TerminalNode FIM_PARA() { return getToken(T2AlgumaLexerParser.FIM_PARA, 0); }
		public List<CmdContext> cmd() {
			return getRuleContexts(CmdContext.class);
		}
		public CmdContext cmd(int i) {
			return getRuleContext(CmdContext.class,i);
		}
		public CmdParaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdPara; }
	}

	public final CmdParaContext cmdPara() throws RecognitionException {
		CmdParaContext _localctx = new CmdParaContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_cmdPara);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(397);
			match(PARA);
			setState(398);
			match(IDENT);
			setState(399);
			match(ATRIBUICAO);
			setState(400);
			exp_aritmetica();
			setState(401);
			match(ATE);
			setState(402);
			exp_aritmetica();
			setState(403);
			match(FACA);
			setState(407);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(404);
				cmd();
				}
				}
				setState(409);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(410);
			match(FIM_PARA);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdEnquantoContext extends ParserRuleContext {
		public TerminalNode ENQUANTO() { return getToken(T2AlgumaLexerParser.ENQUANTO, 0); }
		public ExpressaoContext expressao() {
			return getRuleContext(ExpressaoContext.class,0);
		}
		public TerminalNode FACA() { return getToken(T2AlgumaLexerParser.FACA, 0); }
		public TerminalNode FIM_ENQUANTO() { return getToken(T2AlgumaLexerParser.FIM_ENQUANTO, 0); }
		public List<CmdContext> cmd() {
			return getRuleContexts(CmdContext.class);
		}
		public CmdContext cmd(int i) {
			return getRuleContext(CmdContext.class,i);
		}
		public CmdEnquantoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdEnquanto; }
	}

	public final CmdEnquantoContext cmdEnquanto() throws RecognitionException {
		CmdEnquantoContext _localctx = new CmdEnquantoContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_cmdEnquanto);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(412);
			match(ENQUANTO);
			setState(413);
			expressao();
			setState(414);
			match(FACA);
			setState(418);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(415);
				cmd();
				}
				}
				setState(420);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(421);
			match(FIM_ENQUANTO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdFacaContext extends ParserRuleContext {
		public TerminalNode FACA() { return getToken(T2AlgumaLexerParser.FACA, 0); }
		public TerminalNode ATE() { return getToken(T2AlgumaLexerParser.ATE, 0); }
		public ExpressaoContext expressao() {
			return getRuleContext(ExpressaoContext.class,0);
		}
		public List<CmdContext> cmd() {
			return getRuleContexts(CmdContext.class);
		}
		public CmdContext cmd(int i) {
			return getRuleContext(CmdContext.class,i);
		}
		public CmdFacaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdFaca; }
	}

	public final CmdFacaContext cmdFaca() throws RecognitionException {
		CmdFacaContext _localctx = new CmdFacaContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_cmdFaca);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(423);
			match(FACA);
			setState(427);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(424);
				cmd();
				}
				}
				setState(429);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(430);
			match(ATE);
			setState(431);
			expressao();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdAtribuicaoContext extends ParserRuleContext {
		public IdentificadorContext identificador() {
			return getRuleContext(IdentificadorContext.class,0);
		}
		public TerminalNode ATRIBUICAO() { return getToken(T2AlgumaLexerParser.ATRIBUICAO, 0); }
		public ExpressaoContext expressao() {
			return getRuleContext(ExpressaoContext.class,0);
		}
		public TerminalNode PONTEIRO() { return getToken(T2AlgumaLexerParser.PONTEIRO, 0); }
		public CmdAtribuicaoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdAtribuicao; }
	}

	public final CmdAtribuicaoContext cmdAtribuicao() throws RecognitionException {
		CmdAtribuicaoContext _localctx = new CmdAtribuicaoContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_cmdAtribuicao);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(434);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PONTEIRO) {
				{
				setState(433);
				match(PONTEIRO);
				}
			}

			setState(436);
			identificador();
			setState(437);
			match(ATRIBUICAO);
			setState(438);
			expressao();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdChamadaContext extends ParserRuleContext {
		public TerminalNode IDENT() { return getToken(T2AlgumaLexerParser.IDENT, 0); }
		public TerminalNode ABREPAR() { return getToken(T2AlgumaLexerParser.ABREPAR, 0); }
		public List<ExpressaoContext> expressao() {
			return getRuleContexts(ExpressaoContext.class);
		}
		public ExpressaoContext expressao(int i) {
			return getRuleContext(ExpressaoContext.class,i);
		}
		public TerminalNode FECHAPAR() { return getToken(T2AlgumaLexerParser.FECHAPAR, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(T2AlgumaLexerParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T2AlgumaLexerParser.VIRGULA, i);
		}
		public CmdChamadaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdChamada; }
	}

	public final CmdChamadaContext cmdChamada() throws RecognitionException {
		CmdChamadaContext _localctx = new CmdChamadaContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_cmdChamada);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(440);
			match(IDENT);
			setState(441);
			match(ABREPAR);
			setState(442);
			expressao();
			setState(447);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(443);
				match(VIRGULA);
				setState(444);
				expressao();
				}
				}
				setState(449);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(450);
			match(FECHAPAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CmdRetorneContext extends ParserRuleContext {
		public TerminalNode RETORNE() { return getToken(T2AlgumaLexerParser.RETORNE, 0); }
		public ExpressaoContext expressao() {
			return getRuleContext(ExpressaoContext.class,0);
		}
		public CmdRetorneContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdRetorne; }
	}

	public final CmdRetorneContext cmdRetorne() throws RecognitionException {
		CmdRetorneContext _localctx = new CmdRetorneContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_cmdRetorne);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(452);
			match(RETORNE);
			setState(453);
			expressao();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SelecaoContext extends ParserRuleContext {
		public List<Item_selecaoContext> item_selecao() {
			return getRuleContexts(Item_selecaoContext.class);
		}
		public Item_selecaoContext item_selecao(int i) {
			return getRuleContext(Item_selecaoContext.class,i);
		}
		public SelecaoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selecao; }
	}

	public final SelecaoContext selecao() throws RecognitionException {
		SelecaoContext _localctx = new SelecaoContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_selecao);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(458);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SUBTRACAO || _la==NUM_INT) {
				{
				{
				setState(455);
				item_selecao();
				}
				}
				setState(460);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Item_selecaoContext extends ParserRuleContext {
		public ConstantesContext constantes() {
			return getRuleContext(ConstantesContext.class,0);
		}
		public TerminalNode DOISPONTOS() { return getToken(T2AlgumaLexerParser.DOISPONTOS, 0); }
		public List<CmdContext> cmd() {
			return getRuleContexts(CmdContext.class);
		}
		public CmdContext cmd(int i) {
			return getRuleContext(CmdContext.class,i);
		}
		public Item_selecaoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_item_selecao; }
	}

	public final Item_selecaoContext item_selecao() throws RecognitionException {
		Item_selecaoContext _localctx = new Item_selecaoContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_item_selecao);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(461);
			constantes();
			setState(462);
			match(DOISPONTOS);
			setState(466);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(463);
				cmd();
				}
				}
				setState(468);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConstantesContext extends ParserRuleContext {
		public List<Numero_intervaloContext> numero_intervalo() {
			return getRuleContexts(Numero_intervaloContext.class);
		}
		public Numero_intervaloContext numero_intervalo(int i) {
			return getRuleContext(Numero_intervaloContext.class,i);
		}
		public List<TerminalNode> VIRGULA() { return getTokens(T2AlgumaLexerParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T2AlgumaLexerParser.VIRGULA, i);
		}
		public ConstantesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constantes; }
	}

	public final ConstantesContext constantes() throws RecognitionException {
		ConstantesContext _localctx = new ConstantesContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_constantes);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(469);
			numero_intervalo();
			setState(474);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(470);
				match(VIRGULA);
				setState(471);
				numero_intervalo();
				}
				}
				setState(476);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Numero_intervaloContext extends ParserRuleContext {
		public List<TerminalNode> NUM_INT() { return getTokens(T2AlgumaLexerParser.NUM_INT); }
		public TerminalNode NUM_INT(int i) {
			return getToken(T2AlgumaLexerParser.NUM_INT, i);
		}
		public List<Op_unarioContext> op_unario() {
			return getRuleContexts(Op_unarioContext.class);
		}
		public Op_unarioContext op_unario(int i) {
			return getRuleContext(Op_unarioContext.class,i);
		}
		public TerminalNode INTERVALO() { return getToken(T2AlgumaLexerParser.INTERVALO, 0); }
		public Numero_intervaloContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numero_intervalo; }
	}

	public final Numero_intervaloContext numero_intervalo() throws RecognitionException {
		Numero_intervaloContext _localctx = new Numero_intervaloContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_numero_intervalo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(478);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SUBTRACAO) {
				{
				setState(477);
				op_unario();
				}
			}

			setState(480);
			match(NUM_INT);
			setState(486);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERVALO) {
				{
				setState(481);
				match(INTERVALO);
				setState(483);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SUBTRACAO) {
					{
					setState(482);
					op_unario();
					}
				}

				setState(485);
				match(NUM_INT);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Exp_relacionalContext extends ParserRuleContext {
		public List<Exp_aritmeticaContext> exp_aritmetica() {
			return getRuleContexts(Exp_aritmeticaContext.class);
		}
		public Exp_aritmeticaContext exp_aritmetica(int i) {
			return getRuleContext(Exp_aritmeticaContext.class,i);
		}
		public Op_relacionalContext op_relacional() {
			return getRuleContext(Op_relacionalContext.class,0);
		}
		public Exp_relacionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exp_relacional; }
	}

	public final Exp_relacionalContext exp_relacional() throws RecognitionException {
		Exp_relacionalContext _localctx = new Exp_relacionalContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_exp_relacional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(488);
			exp_aritmetica();
			setState(492);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 69269232549888L) != 0)) {
				{
				setState(489);
				op_relacional();
				setState(490);
				exp_aritmetica();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Op_relacionalContext extends ParserRuleContext {
		public TerminalNode IGUAL() { return getToken(T2AlgumaLexerParser.IGUAL, 0); }
		public TerminalNode DIFERENTE() { return getToken(T2AlgumaLexerParser.DIFERENTE, 0); }
		public TerminalNode MAIORIGUAL() { return getToken(T2AlgumaLexerParser.MAIORIGUAL, 0); }
		public TerminalNode MENORIGUAL() { return getToken(T2AlgumaLexerParser.MENORIGUAL, 0); }
		public TerminalNode MAIOR() { return getToken(T2AlgumaLexerParser.MAIOR, 0); }
		public TerminalNode MENOR() { return getToken(T2AlgumaLexerParser.MENOR, 0); }
		public Op_relacionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_relacional; }
	}

	public final Op_relacionalContext op_relacional() throws RecognitionException {
		Op_relacionalContext _localctx = new Op_relacionalContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_op_relacional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(494);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 69269232549888L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressaoContext extends ParserRuleContext {
		public List<Termo_logicoContext> termo_logico() {
			return getRuleContexts(Termo_logicoContext.class);
		}
		public Termo_logicoContext termo_logico(int i) {
			return getRuleContext(Termo_logicoContext.class,i);
		}
		public List<Op_logico_1Context> op_logico_1() {
			return getRuleContexts(Op_logico_1Context.class);
		}
		public Op_logico_1Context op_logico_1(int i) {
			return getRuleContext(Op_logico_1Context.class,i);
		}
		public ExpressaoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressao; }
	}

	public final ExpressaoContext expressao() throws RecognitionException {
		ExpressaoContext _localctx = new ExpressaoContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_expressao);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(496);
			termo_logico();
			setState(502);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OU) {
				{
				{
				setState(497);
				op_logico_1();
				setState(498);
				termo_logico();
				}
				}
				setState(504);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Termo_logicoContext extends ParserRuleContext {
		public List<Fator_logicoContext> fator_logico() {
			return getRuleContexts(Fator_logicoContext.class);
		}
		public Fator_logicoContext fator_logico(int i) {
			return getRuleContext(Fator_logicoContext.class,i);
		}
		public List<Op_logico_2Context> op_logico_2() {
			return getRuleContexts(Op_logico_2Context.class);
		}
		public Op_logico_2Context op_logico_2(int i) {
			return getRuleContext(Op_logico_2Context.class,i);
		}
		public Termo_logicoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_termo_logico; }
	}

	public final Termo_logicoContext termo_logico() throws RecognitionException {
		Termo_logicoContext _localctx = new Termo_logicoContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_termo_logico);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(505);
			fator_logico();
			setState(511);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==E) {
				{
				{
				setState(506);
				op_logico_2();
				setState(507);
				fator_logico();
				}
				}
				setState(513);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Fator_logicoContext extends ParserRuleContext {
		public Parcela_logicaContext parcela_logica() {
			return getRuleContext(Parcela_logicaContext.class,0);
		}
		public TerminalNode NAO() { return getToken(T2AlgumaLexerParser.NAO, 0); }
		public Fator_logicoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fator_logico; }
	}

	public final Fator_logicoContext fator_logico() throws RecognitionException {
		Fator_logicoContext _localctx = new Fator_logicoContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_fator_logico);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(515);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NAO) {
				{
				setState(514);
				match(NAO);
				}
			}

			setState(517);
			parcela_logica();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Parcela_logicaContext extends ParserRuleContext {
		public TerminalNode VERDADEIRO() { return getToken(T2AlgumaLexerParser.VERDADEIRO, 0); }
		public TerminalNode FALSO() { return getToken(T2AlgumaLexerParser.FALSO, 0); }
		public Exp_relacionalContext exp_relacional() {
			return getRuleContext(Exp_relacionalContext.class,0);
		}
		public Parcela_logicaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parcela_logica; }
	}

	public final Parcela_logicaContext parcela_logica() throws RecognitionException {
		Parcela_logicaContext _localctx = new Parcela_logicaContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_parcela_logica);
		int _la;
		try {
			setState(521);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VERDADEIRO:
			case FALSO:
				enterOuterAlt(_localctx, 1);
				{
				setState(519);
				_la = _input.LA(1);
				if ( !(_la==VERDADEIRO || _la==FALSO) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case ABREPAR:
			case SUBTRACAO:
			case PONTEIRO:
			case ENDERECO:
			case NUM_INT:
			case NUM_REAL:
			case IDENT:
			case CADEIA:
				enterOuterAlt(_localctx, 2);
				{
				setState(520);
				exp_relacional();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Op_logico_1Context extends ParserRuleContext {
		public TerminalNode OU() { return getToken(T2AlgumaLexerParser.OU, 0); }
		public Op_logico_1Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_logico_1; }
	}

	public final Op_logico_1Context op_logico_1() throws RecognitionException {
		Op_logico_1Context _localctx = new Op_logico_1Context(_ctx, getState());
		enterRule(_localctx, 90, RULE_op_logico_1);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(523);
			match(OU);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Op_logico_2Context extends ParserRuleContext {
		public TerminalNode E() { return getToken(T2AlgumaLexerParser.E, 0); }
		public Op_logico_2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_logico_2; }
	}

	public final Op_logico_2Context op_logico_2() throws RecognitionException {
		Op_logico_2Context _localctx = new Op_logico_2Context(_ctx, getState());
		enterRule(_localctx, 92, RULE_op_logico_2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(525);
			match(E);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001E\u0210\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007+\u0002,\u0007,\u0002"+
		"-\u0007-\u0002.\u0007.\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0001\u0001\u0001\u0003\u0001f\b\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0003\u0002u\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0005\u0003"+
		"z\b\u0003\n\u0003\f\u0003}\t\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004\u0085\b\u0004\n\u0004"+
		"\f\u0004\u0088\t\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0005\u0004\u008e\b\u0004\n\u0004\f\u0004\u0091\t\u0004\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0005\u0005\u0097\b\u0005\n\u0005\f\u0005"+
		"\u009a\t\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0005\u0006"+
		"\u00a0\b\u0006\n\u0006\f\u0006\u00a3\t\u0006\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0005\u0007\u00a9\b\u0007\n\u0007\f\u0007\u00ac\t\u0007"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\u000b\u0003\u000b"+
		"\u00b5\b\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u00b9\b\u000b\u0001"+
		"\f\u0003\f\u00bc\b\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0005"+
		"\f\u00c4\b\f\n\f\f\f\u00c7\t\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f"+
		"\u0001\f\u0001\f\u0001\f\u0003\f\u00d1\b\f\u0001\r\u0001\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0003\u000e\u00d8\b\u000e\u0001\u000f\u0001\u000f"+
		"\u0003\u000f\u00dc\b\u000f\u0001\u0010\u0001\u0010\u0001\u0011\u0003\u0011"+
		"\u00e1\b\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u00e5\b\u0011\u0001"+
		"\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0005\u0013\u00eb\b\u0013\n"+
		"\u0013\f\u0013\u00ee\t\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0003"+
		"\u0014\u00f3\b\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u00f8"+
		"\b\u0014\n\u0014\f\u0014\u00fb\t\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0005\u0015\u0103\b\u0015\n\u0015"+
		"\f\u0015\u0106\t\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0003\u0016\u010c\b\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u0110\b"+
		"\u0016\n\u0016\f\u0016\u0113\t\u0016\u0001\u0016\u0005\u0016\u0116\b\u0016"+
		"\n\u0016\f\u0016\u0119\t\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0003\u0016\u0120\b\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0005\u0016\u0126\b\u0016\n\u0016\f\u0016\u0129\t\u0016"+
		"\u0001\u0016\u0005\u0016\u012c\b\u0016\n\u0016\f\u0016\u012f\t\u0016\u0001"+
		"\u0016\u0001\u0016\u0003\u0016\u0133\b\u0016\u0001\u0017\u0005\u0017\u0136"+
		"\b\u0017\n\u0017\f\u0017\u0139\t\u0017\u0001\u0017\u0005\u0017\u013c\b"+
		"\u0017\n\u0017\f\u0017\u013f\t\u0017\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0003\u0018\u014b\b\u0018\u0001\u0019\u0001\u0019\u0001\u0019"+
		"\u0003\u0019\u0150\b\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0003\u0019"+
		"\u0155\b\u0019\u0001\u0019\u0005\u0019\u0158\b\u0019\n\u0019\f\u0019\u015b"+
		"\t\u0019\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0005\u001a\u0164\b\u001a\n\u001a\f\u001a\u0167\t\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0005\u001b\u016f\b\u001b\n\u001b\f\u001b\u0172\t\u001b\u0001\u001b\u0001"+
		"\u001b\u0005\u001b\u0176\b\u001b\n\u001b\f\u001b\u0179\t\u001b\u0003\u001b"+
		"\u017b\b\u001b\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0001\u001c\u0001\u001c\u0001\u001c\u0005\u001c\u0185\b\u001c\n\u001c"+
		"\f\u001c\u0188\t\u001c\u0003\u001c\u018a\b\u001c\u0001\u001c\u0001\u001c"+
		"\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0005\u001d\u0196\b\u001d\n\u001d\f\u001d\u0199"+
		"\t\u001d\u0001\u001d\u0001\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0005\u001e\u01a1\b\u001e\n\u001e\f\u001e\u01a4\t\u001e\u0001\u001e"+
		"\u0001\u001e\u0001\u001f\u0001\u001f\u0005\u001f\u01aa\b\u001f\n\u001f"+
		"\f\u001f\u01ad\t\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001 \u0003"+
		" \u01b3\b \u0001 \u0001 \u0001 \u0001 \u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0005!\u01be\b!\n!\f!\u01c1\t!\u0001!\u0001!\u0001\"\u0001\"\u0001\""+
		"\u0001#\u0005#\u01c9\b#\n#\f#\u01cc\t#\u0001$\u0001$\u0001$\u0005$\u01d1"+
		"\b$\n$\f$\u01d4\t$\u0001%\u0001%\u0001%\u0005%\u01d9\b%\n%\f%\u01dc\t"+
		"%\u0001&\u0003&\u01df\b&\u0001&\u0001&\u0001&\u0003&\u01e4\b&\u0001&\u0003"+
		"&\u01e7\b&\u0001\'\u0001\'\u0001\'\u0001\'\u0003\'\u01ed\b\'\u0001(\u0001"+
		"(\u0001)\u0001)\u0001)\u0001)\u0005)\u01f5\b)\n)\f)\u01f8\t)\u0001*\u0001"+
		"*\u0001*\u0001*\u0005*\u01fe\b*\n*\f*\u0201\t*\u0001+\u0003+\u0204\b+"+
		"\u0001+\u0001+\u0001,\u0001,\u0003,\u020a\b,\u0001-\u0001-\u0001.\u0001"+
		".\u0001.\u0000\u0000/\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012"+
		"\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJLNPRTVXZ\\"+
		"\u0000\u0006\u0001\u000078\u0002\u00005599\u0001\u0000\u0006\t\u0003\u0000"+
		"\n\u000b>?AA\u0001\u0000(-\u0001\u0000\n\u000b\u0224\u0000^\u0001\u0000"+
		"\u0000\u0000\u0002e\u0001\u0000\u0000\u0000\u0004t\u0001\u0000\u0000\u0000"+
		"\u0006v\u0001\u0000\u0000\u0000\b\u0081\u0001\u0000\u0000\u0000\n\u0092"+
		"\u0001\u0000\u0000\u0000\f\u009b\u0001\u0000\u0000\u0000\u000e\u00a4\u0001"+
		"\u0000\u0000\u0000\u0010\u00ad\u0001\u0000\u0000\u0000\u0012\u00af\u0001"+
		"\u0000\u0000\u0000\u0014\u00b1\u0001\u0000\u0000\u0000\u0016\u00b8\u0001"+
		"\u0000\u0000\u0000\u0018\u00d0\u0001\u0000\u0000\u0000\u001a\u00d2\u0001"+
		"\u0000\u0000\u0000\u001c\u00d7\u0001\u0000\u0000\u0000\u001e\u00db\u0001"+
		"\u0000\u0000\u0000 \u00dd\u0001\u0000\u0000\u0000\"\u00e0\u0001\u0000"+
		"\u0000\u0000$\u00e6\u0001\u0000\u0000\u0000&\u00e8\u0001\u0000\u0000\u0000"+
		"(\u00f2\u0001\u0000\u0000\u0000*\u00ff\u0001\u0000\u0000\u0000,\u0132"+
		"\u0001\u0000\u0000\u0000.\u0137\u0001\u0000\u0000\u00000\u014a\u0001\u0000"+
		"\u0000\u00002\u014c\u0001\u0000\u0000\u00004\u015e\u0001\u0000\u0000\u0000"+
		"6\u016a\u0001\u0000\u0000\u00008\u017e\u0001\u0000\u0000\u0000:\u018d"+
		"\u0001\u0000\u0000\u0000<\u019c\u0001\u0000\u0000\u0000>\u01a7\u0001\u0000"+
		"\u0000\u0000@\u01b2\u0001\u0000\u0000\u0000B\u01b8\u0001\u0000\u0000\u0000"+
		"D\u01c4\u0001\u0000\u0000\u0000F\u01ca\u0001\u0000\u0000\u0000H\u01cd"+
		"\u0001\u0000\u0000\u0000J\u01d5\u0001\u0000\u0000\u0000L\u01de\u0001\u0000"+
		"\u0000\u0000N\u01e8\u0001\u0000\u0000\u0000P\u01ee\u0001\u0000\u0000\u0000"+
		"R\u01f0\u0001\u0000\u0000\u0000T\u01f9\u0001\u0000\u0000\u0000V\u0203"+
		"\u0001\u0000\u0000\u0000X\u0209\u0001\u0000\u0000\u0000Z\u020b\u0001\u0000"+
		"\u0000\u0000\\\u020d\u0001\u0000\u0000\u0000^_\u0003\u0002\u0001\u0000"+
		"_`\u0005\u0002\u0000\u0000`a\u0003.\u0017\u0000ab\u0005&\u0000\u0000b"+
		"\u0001\u0001\u0000\u0000\u0000cf\u0003\u0004\u0002\u0000df\u0003,\u0016"+
		"\u0000ec\u0001\u0000\u0000\u0000ed\u0001\u0000\u0000\u0000f\u0003\u0001"+
		"\u0000\u0000\u0000gh\u0005\u0003\u0000\u0000hu\u0003\u0006\u0003\u0000"+
		"ij\u0005\u0005\u0000\u0000jk\u0005@\u0000\u0000kl\u0005.\u0000\u0000l"+
		"m\u0003 \u0010\u0000mn\u0005,\u0000\u0000no\u0003$\u0012\u0000ou\u0001"+
		"\u0000\u0000\u0000pq\u0005\u001c\u0000\u0000qr\u0005@\u0000\u0000rs\u0005"+
		".\u0000\u0000su\u0003&\u0013\u0000tg\u0001\u0000\u0000\u0000ti\u0001\u0000"+
		"\u0000\u0000tp\u0001\u0000\u0000\u0000u\u0005\u0001\u0000\u0000\u0000"+
		"v{\u0003\b\u0004\u0000wx\u00053\u0000\u0000xz\u0003\b\u0004\u0000yw\u0001"+
		"\u0000\u0000\u0000z}\u0001\u0000\u0000\u0000{y\u0001\u0000\u0000\u0000"+
		"{|\u0001\u0000\u0000\u0000|~\u0001\u0000\u0000\u0000}{\u0001\u0000\u0000"+
		"\u0000~\u007f\u0005.\u0000\u0000\u007f\u0080\u0003\u001e\u000f\u0000\u0080"+
		"\u0007\u0001\u0000\u0000\u0000\u0081\u0086\u0005@\u0000\u0000\u0082\u0083"+
		"\u0005=\u0000\u0000\u0083\u0085\u0005@\u0000\u0000\u0084\u0082\u0001\u0000"+
		"\u0000\u0000\u0085\u0088\u0001\u0000\u0000\u0000\u0086\u0084\u0001\u0000"+
		"\u0000\u0000\u0086\u0087\u0001\u0000\u0000\u0000\u0087\u008f\u0001\u0000"+
		"\u0000\u0000\u0088\u0086\u0001\u0000\u0000\u0000\u0089\u008a\u00051\u0000"+
		"\u0000\u008a\u008b\u0003\n\u0005\u0000\u008b\u008c\u00052\u0000\u0000"+
		"\u008c\u008e\u0001\u0000\u0000\u0000\u008d\u0089\u0001\u0000\u0000\u0000"+
		"\u008e\u0091\u0001\u0000\u0000\u0000\u008f\u008d\u0001\u0000\u0000\u0000"+
		"\u008f\u0090\u0001\u0000\u0000\u0000\u0090\t\u0001\u0000\u0000\u0000\u0091"+
		"\u008f\u0001\u0000\u0000\u0000\u0092\u0098\u0003\f\u0006\u0000\u0093\u0094"+
		"\u0003\u0010\b\u0000\u0094\u0095\u0003\f\u0006\u0000\u0095\u0097\u0001"+
		"\u0000\u0000\u0000\u0096\u0093\u0001\u0000\u0000\u0000\u0097\u009a\u0001"+
		"\u0000\u0000\u0000\u0098\u0096\u0001\u0000\u0000\u0000\u0098\u0099\u0001"+
		"\u0000\u0000\u0000\u0099\u000b\u0001\u0000\u0000\u0000\u009a\u0098\u0001"+
		"\u0000\u0000\u0000\u009b\u00a1\u0003\u000e\u0007\u0000\u009c\u009d\u0003"+
		"\u0012\t\u0000\u009d\u009e\u0003\u000e\u0007\u0000\u009e\u00a0\u0001\u0000"+
		"\u0000\u0000\u009f\u009c\u0001\u0000\u0000\u0000\u00a0\u00a3\u0001\u0000"+
		"\u0000\u0000\u00a1\u009f\u0001\u0000\u0000\u0000\u00a1\u00a2\u0001\u0000"+
		"\u0000\u0000\u00a2\r\u0001\u0000\u0000\u0000\u00a3\u00a1\u0001\u0000\u0000"+
		"\u0000\u00a4\u00aa\u0003\u0016\u000b\u0000\u00a5\u00a6\u0003\u0014\n\u0000"+
		"\u00a6\u00a7\u0003\u0016\u000b\u0000\u00a7\u00a9\u0001\u0000\u0000\u0000"+
		"\u00a8\u00a5\u0001\u0000\u0000\u0000\u00a9\u00ac\u0001\u0000\u0000\u0000"+
		"\u00aa\u00a8\u0001\u0000\u0000\u0000\u00aa\u00ab\u0001\u0000\u0000\u0000"+
		"\u00ab\u000f\u0001\u0000\u0000\u0000\u00ac\u00aa\u0001\u0000\u0000\u0000"+
		"\u00ad\u00ae\u0007\u0000\u0000\u0000\u00ae\u0011\u0001\u0000\u0000\u0000"+
		"\u00af\u00b0\u0007\u0001\u0000\u0000\u00b0\u0013\u0001\u0000\u0000\u0000"+
		"\u00b1\u00b2\u00056\u0000\u0000\u00b2\u0015\u0001\u0000\u0000\u0000\u00b3"+
		"\u00b5\u0003\u001a\r\u0000\u00b4\u00b3\u0001\u0000\u0000\u0000\u00b4\u00b5"+
		"\u0001\u0000\u0000\u0000\u00b5\u00b6\u0001\u0000\u0000\u0000\u00b6\u00b9"+
		"\u0003\u0018\f\u0000\u00b7\u00b9\u0003\u001c\u000e\u0000\u00b8\u00b4\u0001"+
		"\u0000\u0000\u0000\u00b8\u00b7\u0001\u0000\u0000\u0000\u00b9\u0017\u0001"+
		"\u0000\u0000\u0000\u00ba\u00bc\u0005;\u0000\u0000\u00bb\u00ba\u0001\u0000"+
		"\u0000\u0000\u00bb\u00bc\u0001\u0000\u0000\u0000\u00bc\u00bd\u0001\u0000"+
		"\u0000\u0000\u00bd\u00d1\u0003\b\u0004\u0000\u00be\u00bf\u0005@\u0000"+
		"\u0000\u00bf\u00c0\u0005/\u0000\u0000\u00c0\u00c5\u0003R)\u0000\u00c1"+
		"\u00c2\u00053\u0000\u0000\u00c2\u00c4\u0003R)\u0000\u00c3\u00c1\u0001"+
		"\u0000\u0000\u0000\u00c4\u00c7\u0001\u0000\u0000\u0000\u00c5\u00c3\u0001"+
		"\u0000\u0000\u0000\u00c5\u00c6\u0001\u0000\u0000\u0000\u00c6\u00c8\u0001"+
		"\u0000\u0000\u0000\u00c7\u00c5\u0001\u0000\u0000\u0000\u00c8\u00c9\u0005"+
		"0\u0000\u0000\u00c9\u00d1\u0001\u0000\u0000\u0000\u00ca\u00d1\u0005>\u0000"+
		"\u0000\u00cb\u00d1\u0005?\u0000\u0000\u00cc\u00cd\u0005/\u0000\u0000\u00cd"+
		"\u00ce\u0003R)\u0000\u00ce\u00cf\u00050\u0000\u0000\u00cf\u00d1\u0001"+
		"\u0000\u0000\u0000\u00d0\u00bb\u0001\u0000\u0000\u0000\u00d0\u00be\u0001"+
		"\u0000\u0000\u0000\u00d0\u00ca\u0001\u0000\u0000\u0000\u00d0\u00cb\u0001"+
		"\u0000\u0000\u0000\u00d0\u00cc\u0001\u0000\u0000\u0000\u00d1\u0019\u0001"+
		"\u0000\u0000\u0000\u00d2\u00d3\u00058\u0000\u0000\u00d3\u001b\u0001\u0000"+
		"\u0000\u0000\u00d4\u00d5\u0005<\u0000\u0000\u00d5\u00d8\u0003\b\u0004"+
		"\u0000\u00d6\u00d8\u0005A\u0000\u0000\u00d7\u00d4\u0001\u0000\u0000\u0000"+
		"\u00d7\u00d6\u0001\u0000\u0000\u0000\u00d8\u001d\u0001\u0000\u0000\u0000"+
		"\u00d9\u00dc\u0003&\u0013\u0000\u00da\u00dc\u0003\"\u0011\u0000\u00db"+
		"\u00d9\u0001\u0000\u0000\u0000\u00db\u00da\u0001\u0000\u0000\u0000\u00dc"+
		"\u001f\u0001\u0000\u0000\u0000\u00dd\u00de\u0007\u0002\u0000\u0000\u00de"+
		"!\u0001\u0000\u0000\u0000\u00df\u00e1\u0005;\u0000\u0000\u00e0\u00df\u0001"+
		"\u0000\u0000\u0000\u00e0\u00e1\u0001\u0000\u0000\u0000\u00e1\u00e4\u0001"+
		"\u0000\u0000\u0000\u00e2\u00e5\u0003 \u0010\u0000\u00e3\u00e5\u0005@\u0000"+
		"\u0000\u00e4\u00e2\u0001\u0000\u0000\u0000\u00e4\u00e3\u0001\u0000\u0000"+
		"\u0000\u00e5#\u0001\u0000\u0000\u0000\u00e6\u00e7\u0007\u0003\u0000\u0000"+
		"\u00e7%\u0001\u0000\u0000\u0000\u00e8\u00ec\u0005\u001d\u0000\u0000\u00e9"+
		"\u00eb\u0003\u0006\u0003\u0000\u00ea\u00e9\u0001\u0000\u0000\u0000\u00eb"+
		"\u00ee\u0001\u0000\u0000\u0000\u00ec\u00ea\u0001\u0000\u0000\u0000\u00ec"+
		"\u00ed\u0001\u0000\u0000\u0000\u00ed\u00ef\u0001\u0000\u0000\u0000\u00ee"+
		"\u00ec\u0001\u0000\u0000\u0000\u00ef\u00f0\u0005\u001e\u0000\u0000\u00f0"+
		"\'\u0001\u0000\u0000\u0000\u00f1\u00f3\u0005\u0004\u0000\u0000\u00f2\u00f1"+
		"\u0001\u0000\u0000\u0000\u00f2\u00f3\u0001\u0000\u0000\u0000\u00f3\u00f4"+
		"\u0001\u0000\u0000\u0000\u00f4\u00f9\u0003\b\u0004\u0000\u00f5\u00f6\u0005"+
		"3\u0000\u0000\u00f6\u00f8\u0003\b\u0004\u0000\u00f7\u00f5\u0001\u0000"+
		"\u0000\u0000\u00f8\u00fb\u0001\u0000\u0000\u0000\u00f9\u00f7\u0001\u0000"+
		"\u0000\u0000\u00f9\u00fa\u0001\u0000\u0000\u0000\u00fa\u00fc\u0001\u0000"+
		"\u0000\u0000\u00fb\u00f9\u0001\u0000\u0000\u0000\u00fc\u00fd\u0005.\u0000"+
		"\u0000\u00fd\u00fe\u0003\"\u0011\u0000\u00fe)\u0001\u0000\u0000\u0000"+
		"\u00ff\u0104\u0003(\u0014\u0000\u0100\u0101\u00053\u0000\u0000\u0101\u0103"+
		"\u0003(\u0014\u0000\u0102\u0100\u0001\u0000\u0000\u0000\u0103\u0106\u0001"+
		"\u0000\u0000\u0000\u0104\u0102\u0001\u0000\u0000\u0000\u0104\u0105\u0001"+
		"\u0000\u0000\u0000\u0105+\u0001\u0000\u0000\u0000\u0106\u0104\u0001\u0000"+
		"\u0000\u0000\u0107\u0108\u0005\u001f\u0000\u0000\u0108\u0109\u0005@\u0000"+
		"\u0000\u0109\u010b\u0005/\u0000\u0000\u010a\u010c\u0003*\u0015\u0000\u010b"+
		"\u010a\u0001\u0000\u0000\u0000\u010b\u010c\u0001\u0000\u0000\u0000\u010c"+
		"\u010d\u0001\u0000\u0000\u0000\u010d\u0111\u00050\u0000\u0000\u010e\u0110"+
		"\u0003\u0004\u0002\u0000\u010f\u010e\u0001\u0000\u0000\u0000\u0110\u0113"+
		"\u0001\u0000\u0000\u0000\u0111\u010f\u0001\u0000\u0000\u0000\u0111\u0112"+
		"\u0001\u0000\u0000\u0000\u0112\u0117\u0001\u0000\u0000\u0000\u0113\u0111"+
		"\u0001\u0000\u0000\u0000\u0114\u0116\u00030\u0018\u0000\u0115\u0114\u0001"+
		"\u0000\u0000\u0000\u0116\u0119\u0001\u0000\u0000\u0000\u0117\u0115\u0001"+
		"\u0000\u0000\u0000\u0117\u0118\u0001\u0000\u0000\u0000\u0118\u011a\u0001"+
		"\u0000\u0000\u0000\u0119\u0117\u0001\u0000\u0000\u0000\u011a\u0133\u0005"+
		" \u0000\u0000\u011b\u011c\u0005!\u0000\u0000\u011c\u011d\u0005@\u0000"+
		"\u0000\u011d\u011f\u0005/\u0000\u0000\u011e\u0120\u0003*\u0015\u0000\u011f"+
		"\u011e\u0001\u0000\u0000\u0000\u011f\u0120\u0001\u0000\u0000\u0000\u0120"+
		"\u0121\u0001\u0000\u0000\u0000\u0121\u0122\u00050\u0000\u0000\u0122\u0123"+
		"\u0005.\u0000\u0000\u0123\u0127\u0003\"\u0011\u0000\u0124\u0126\u0003"+
		"\u0004\u0002\u0000\u0125\u0124\u0001\u0000\u0000\u0000\u0126\u0129\u0001"+
		"\u0000\u0000\u0000\u0127\u0125\u0001\u0000\u0000\u0000\u0127\u0128\u0001"+
		"\u0000\u0000\u0000\u0128\u012d\u0001\u0000\u0000\u0000\u0129\u0127\u0001"+
		"\u0000\u0000\u0000\u012a\u012c\u00030\u0018\u0000\u012b\u012a\u0001\u0000"+
		"\u0000\u0000\u012c\u012f\u0001\u0000\u0000\u0000\u012d\u012b\u0001\u0000"+
		"\u0000\u0000\u012d\u012e\u0001\u0000\u0000\u0000\u012e\u0130\u0001\u0000"+
		"\u0000\u0000\u012f\u012d\u0001\u0000\u0000\u0000\u0130\u0131\u0005\"\u0000"+
		"\u0000\u0131\u0133\u0001\u0000\u0000\u0000\u0132\u0107\u0001\u0000\u0000"+
		"\u0000\u0132\u011b\u0001\u0000\u0000\u0000\u0133-\u0001\u0000\u0000\u0000"+
		"\u0134\u0136\u0003\u0004\u0002\u0000\u0135\u0134\u0001\u0000\u0000\u0000"+
		"\u0136\u0139\u0001\u0000\u0000\u0000\u0137\u0135\u0001\u0000\u0000\u0000"+
		"\u0137\u0138\u0001\u0000\u0000\u0000\u0138\u013d\u0001\u0000\u0000\u0000"+
		"\u0139\u0137\u0001\u0000\u0000\u0000\u013a\u013c\u00030\u0018\u0000\u013b"+
		"\u013a\u0001\u0000\u0000\u0000\u013c\u013f\u0001\u0000\u0000\u0000\u013d"+
		"\u013b\u0001\u0000\u0000\u0000\u013d\u013e\u0001\u0000\u0000\u0000\u013e"+
		"/\u0001\u0000\u0000\u0000\u013f\u013d\u0001\u0000\u0000\u0000\u0140\u014b"+
		"\u00032\u0019\u0000\u0141\u014b\u00034\u001a\u0000\u0142\u014b\u00036"+
		"\u001b\u0000\u0143\u014b\u00038\u001c\u0000\u0144\u014b\u0003:\u001d\u0000"+
		"\u0145\u014b\u0003<\u001e\u0000\u0146\u014b\u0003>\u001f\u0000\u0147\u014b"+
		"\u0003@ \u0000\u0148\u014b\u0003B!\u0000\u0149\u014b\u0003D\"\u0000\u014a"+
		"\u0140\u0001\u0000\u0000\u0000\u014a\u0141\u0001\u0000\u0000\u0000\u014a"+
		"\u0142\u0001\u0000\u0000\u0000\u014a\u0143\u0001\u0000\u0000\u0000\u014a"+
		"\u0144\u0001\u0000\u0000\u0000\u014a\u0145\u0001\u0000\u0000\u0000\u014a"+
		"\u0146\u0001\u0000\u0000\u0000\u014a\u0147\u0001\u0000\u0000\u0000\u014a"+
		"\u0148\u0001\u0000\u0000\u0000\u014a\u0149\u0001\u0000\u0000\u0000\u014b"+
		"1\u0001\u0000\u0000\u0000\u014c\u014d\u0005$\u0000\u0000\u014d\u014f\u0005"+
		"/\u0000\u0000\u014e\u0150\u0005;\u0000\u0000\u014f\u014e\u0001\u0000\u0000"+
		"\u0000\u014f\u0150\u0001\u0000\u0000\u0000\u0150\u0151\u0001\u0000\u0000"+
		"\u0000\u0151\u0159\u0003\b\u0004\u0000\u0152\u0154\u00053\u0000\u0000"+
		"\u0153\u0155\u0005;\u0000\u0000\u0154\u0153\u0001\u0000\u0000\u0000\u0154"+
		"\u0155\u0001\u0000\u0000\u0000\u0155\u0156\u0001\u0000\u0000\u0000\u0156"+
		"\u0158\u0003\b\u0004\u0000\u0157\u0152\u0001\u0000\u0000\u0000\u0158\u015b"+
		"\u0001\u0000\u0000\u0000\u0159\u0157\u0001\u0000\u0000\u0000\u0159\u015a"+
		"\u0001\u0000\u0000\u0000\u015a\u015c\u0001\u0000\u0000\u0000\u015b\u0159"+
		"\u0001\u0000\u0000\u0000\u015c\u015d\u00050\u0000\u0000\u015d3\u0001\u0000"+
		"\u0000\u0000\u015e\u015f\u0005%\u0000\u0000\u015f\u0160\u0005/\u0000\u0000"+
		"\u0160\u0165\u0003R)\u0000\u0161\u0162\u00053\u0000\u0000\u0162\u0164"+
		"\u0003R)\u0000\u0163\u0161\u0001\u0000\u0000\u0000\u0164\u0167\u0001\u0000"+
		"\u0000\u0000\u0165\u0163\u0001\u0000\u0000\u0000\u0165\u0166\u0001\u0000"+
		"\u0000\u0000\u0166\u0168\u0001\u0000\u0000\u0000\u0167\u0165\u0001\u0000"+
		"\u0000\u0000\u0168\u0169\u00050\u0000\u0000\u01695\u0001\u0000\u0000\u0000"+
		"\u016a\u016b\u0005\u000f\u0000\u0000\u016b\u016c\u0003R)\u0000\u016c\u0170"+
		"\u0005\u0011\u0000\u0000\u016d\u016f\u00030\u0018\u0000\u016e\u016d\u0001"+
		"\u0000\u0000\u0000\u016f\u0172\u0001\u0000\u0000\u0000\u0170\u016e\u0001"+
		"\u0000\u0000\u0000\u0170\u0171\u0001\u0000\u0000\u0000\u0171\u017a\u0001"+
		"\u0000\u0000\u0000\u0172\u0170\u0001\u0000\u0000\u0000\u0173\u0177\u0005"+
		"\u0012\u0000\u0000\u0174\u0176\u00030\u0018\u0000\u0175\u0174\u0001\u0000"+
		"\u0000\u0000\u0176\u0179\u0001\u0000\u0000\u0000\u0177\u0175\u0001\u0000"+
		"\u0000\u0000\u0177\u0178\u0001\u0000\u0000\u0000\u0178\u017b\u0001\u0000"+
		"\u0000\u0000\u0179\u0177\u0001\u0000\u0000\u0000\u017a\u0173\u0001\u0000"+
		"\u0000\u0000\u017a\u017b\u0001\u0000\u0000\u0000\u017b\u017c\u0001\u0000"+
		"\u0000\u0000\u017c\u017d\u0005\u0010\u0000\u0000\u017d7\u0001\u0000\u0000"+
		"\u0000\u017e\u017f\u0005\u0013\u0000\u0000\u017f\u0180\u0003\n\u0005\u0000"+
		"\u0180\u0181\u0005\u0014\u0000\u0000\u0181\u0189\u0003F#\u0000\u0182\u0186"+
		"\u0005\u0012\u0000\u0000\u0183\u0185\u00030\u0018\u0000\u0184\u0183\u0001"+
		"\u0000\u0000\u0000\u0185\u0188\u0001\u0000\u0000\u0000\u0186\u0184\u0001"+
		"\u0000\u0000\u0000\u0186\u0187\u0001\u0000\u0000\u0000\u0187\u018a\u0001"+
		"\u0000\u0000\u0000\u0188\u0186\u0001\u0000\u0000\u0000\u0189\u0182\u0001"+
		"\u0000\u0000\u0000\u0189\u018a\u0001\u0000\u0000\u0000\u018a\u018b\u0001"+
		"\u0000\u0000\u0000\u018b\u018c\u0005\u0015\u0000\u0000\u018c9\u0001\u0000"+
		"\u0000\u0000\u018d\u018e\u0005\u0016\u0000\u0000\u018e\u018f\u0005@\u0000"+
		"\u0000\u018f\u0190\u0005:\u0000\u0000\u0190\u0191\u0003\n\u0005\u0000"+
		"\u0191\u0192\u0005\u0018\u0000\u0000\u0192\u0193\u0003\n\u0005\u0000\u0193"+
		"\u0197\u0005\u0019\u0000\u0000\u0194\u0196\u00030\u0018\u0000\u0195\u0194"+
		"\u0001\u0000\u0000\u0000\u0196\u0199\u0001\u0000\u0000\u0000\u0197\u0195"+
		"\u0001\u0000\u0000\u0000\u0197\u0198\u0001\u0000\u0000\u0000\u0198\u019a"+
		"\u0001\u0000\u0000\u0000\u0199\u0197\u0001\u0000\u0000\u0000\u019a\u019b"+
		"\u0005\u0017\u0000\u0000\u019b;\u0001\u0000\u0000\u0000\u019c\u019d\u0005"+
		"\u001a\u0000\u0000\u019d\u019e\u0003R)\u0000\u019e\u01a2\u0005\u0019\u0000"+
		"\u0000\u019f\u01a1\u00030\u0018\u0000\u01a0\u019f\u0001\u0000\u0000\u0000"+
		"\u01a1\u01a4\u0001\u0000\u0000\u0000\u01a2\u01a0\u0001\u0000\u0000\u0000"+
		"\u01a2\u01a3\u0001\u0000\u0000\u0000\u01a3\u01a5\u0001\u0000\u0000\u0000"+
		"\u01a4\u01a2\u0001\u0000\u0000\u0000\u01a5\u01a6\u0005\u001b\u0000\u0000"+
		"\u01a6=\u0001\u0000\u0000\u0000\u01a7\u01ab\u0005\u0019\u0000\u0000\u01a8"+
		"\u01aa\u00030\u0018\u0000\u01a9\u01a8\u0001\u0000\u0000\u0000\u01aa\u01ad"+
		"\u0001\u0000\u0000\u0000\u01ab\u01a9\u0001\u0000\u0000\u0000\u01ab\u01ac"+
		"\u0001\u0000\u0000\u0000\u01ac\u01ae\u0001\u0000\u0000\u0000\u01ad\u01ab"+
		"\u0001\u0000\u0000\u0000\u01ae\u01af\u0005\u0018\u0000\u0000\u01af\u01b0"+
		"\u0003R)\u0000\u01b0?\u0001\u0000\u0000\u0000\u01b1\u01b3\u0005;\u0000"+
		"\u0000\u01b2\u01b1\u0001\u0000\u0000\u0000\u01b2\u01b3\u0001\u0000\u0000"+
		"\u0000\u01b3\u01b4\u0001\u0000\u0000\u0000\u01b4\u01b5\u0003\b\u0004\u0000"+
		"\u01b5\u01b6\u0005:\u0000\u0000\u01b6\u01b7\u0003R)\u0000\u01b7A\u0001"+
		"\u0000\u0000\u0000\u01b8\u01b9\u0005@\u0000\u0000\u01b9\u01ba\u0005/\u0000"+
		"\u0000\u01ba\u01bf\u0003R)\u0000\u01bb\u01bc\u00053\u0000\u0000\u01bc"+
		"\u01be\u0003R)\u0000\u01bd\u01bb\u0001\u0000\u0000\u0000\u01be\u01c1\u0001"+
		"\u0000\u0000\u0000\u01bf\u01bd\u0001\u0000\u0000\u0000\u01bf\u01c0\u0001"+
		"\u0000\u0000\u0000\u01c0\u01c2\u0001\u0000\u0000\u0000\u01c1\u01bf\u0001"+
		"\u0000\u0000\u0000\u01c2\u01c3\u00050\u0000\u0000\u01c3C\u0001\u0000\u0000"+
		"\u0000\u01c4\u01c5\u0005#\u0000\u0000\u01c5\u01c6\u0003R)\u0000\u01c6"+
		"E\u0001\u0000\u0000\u0000\u01c7\u01c9\u0003H$\u0000\u01c8\u01c7\u0001"+
		"\u0000\u0000\u0000\u01c9\u01cc\u0001\u0000\u0000\u0000\u01ca\u01c8\u0001"+
		"\u0000\u0000\u0000\u01ca\u01cb\u0001\u0000\u0000\u0000\u01cbG\u0001\u0000"+
		"\u0000\u0000\u01cc\u01ca\u0001\u0000\u0000\u0000\u01cd\u01ce\u0003J%\u0000"+
		"\u01ce\u01d2\u0005.\u0000\u0000\u01cf\u01d1\u00030\u0018\u0000\u01d0\u01cf"+
		"\u0001\u0000\u0000\u0000\u01d1\u01d4\u0001\u0000\u0000\u0000\u01d2\u01d0"+
		"\u0001\u0000\u0000\u0000\u01d2\u01d3\u0001\u0000\u0000\u0000\u01d3I\u0001"+
		"\u0000\u0000\u0000\u01d4\u01d2\u0001\u0000\u0000\u0000\u01d5\u01da\u0003"+
		"L&\u0000\u01d6\u01d7\u00053\u0000\u0000\u01d7\u01d9\u0003L&\u0000\u01d8"+
		"\u01d6\u0001\u0000\u0000\u0000\u01d9\u01dc\u0001\u0000\u0000\u0000\u01da"+
		"\u01d8\u0001\u0000\u0000\u0000\u01da\u01db\u0001\u0000\u0000\u0000\u01db"+
		"K\u0001\u0000\u0000\u0000\u01dc\u01da\u0001\u0000\u0000\u0000\u01dd\u01df"+
		"\u0003\u001a\r\u0000\u01de\u01dd\u0001\u0000\u0000\u0000\u01de\u01df\u0001"+
		"\u0000\u0000\u0000\u01df\u01e0\u0001\u0000\u0000\u0000\u01e0\u01e6\u0005"+
		">\u0000\u0000\u01e1\u01e3\u0005\'\u0000\u0000\u01e2\u01e4\u0003\u001a"+
		"\r\u0000\u01e3\u01e2\u0001\u0000\u0000\u0000\u01e3\u01e4\u0001\u0000\u0000"+
		"\u0000\u01e4\u01e5\u0001\u0000\u0000\u0000\u01e5\u01e7\u0005>\u0000\u0000"+
		"\u01e6\u01e1\u0001\u0000\u0000\u0000\u01e6\u01e7\u0001\u0000\u0000\u0000"+
		"\u01e7M\u0001\u0000\u0000\u0000\u01e8\u01ec\u0003\n\u0005\u0000\u01e9"+
		"\u01ea\u0003P(\u0000\u01ea\u01eb\u0003\n\u0005\u0000\u01eb\u01ed\u0001"+
		"\u0000\u0000\u0000\u01ec\u01e9\u0001\u0000\u0000\u0000\u01ec\u01ed\u0001"+
		"\u0000\u0000\u0000\u01edO\u0001\u0000\u0000\u0000\u01ee\u01ef\u0007\u0004"+
		"\u0000\u0000\u01efQ\u0001\u0000\u0000\u0000\u01f0\u01f6\u0003T*\u0000"+
		"\u01f1\u01f2\u0003Z-\u0000\u01f2\u01f3\u0003T*\u0000\u01f3\u01f5\u0001"+
		"\u0000\u0000\u0000\u01f4\u01f1\u0001\u0000\u0000\u0000\u01f5\u01f8\u0001"+
		"\u0000\u0000\u0000\u01f6\u01f4\u0001\u0000\u0000\u0000\u01f6\u01f7\u0001"+
		"\u0000\u0000\u0000\u01f7S\u0001\u0000\u0000\u0000\u01f8\u01f6\u0001\u0000"+
		"\u0000\u0000\u01f9\u01ff\u0003V+\u0000\u01fa\u01fb\u0003\\.\u0000\u01fb"+
		"\u01fc\u0003V+\u0000\u01fc\u01fe\u0001\u0000\u0000\u0000\u01fd\u01fa\u0001"+
		"\u0000\u0000\u0000\u01fe\u0201\u0001\u0000\u0000\u0000\u01ff\u01fd\u0001"+
		"\u0000\u0000\u0000\u01ff\u0200\u0001\u0000\u0000\u0000\u0200U\u0001\u0000"+
		"\u0000\u0000\u0201\u01ff\u0001\u0000\u0000\u0000\u0202\u0204\u0005\u000e"+
		"\u0000\u0000\u0203\u0202\u0001\u0000\u0000\u0000\u0203\u0204\u0001\u0000"+
		"\u0000\u0000\u0204\u0205\u0001\u0000\u0000\u0000\u0205\u0206\u0003X,\u0000"+
		"\u0206W\u0001\u0000\u0000\u0000\u0207\u020a\u0007\u0005\u0000\u0000\u0208"+
		"\u020a\u0003N\'\u0000\u0209\u0207\u0001\u0000\u0000\u0000\u0209\u0208"+
		"\u0001\u0000\u0000\u0000\u020aY\u0001\u0000\u0000\u0000\u020b\u020c\u0005"+
		"\r\u0000\u0000\u020c[\u0001\u0000\u0000\u0000\u020d\u020e\u0005\f\u0000"+
		"\u0000\u020e]\u0001\u0000\u0000\u00008et{\u0086\u008f\u0098\u00a1\u00aa"+
		"\u00b4\u00b8\u00bb\u00c5\u00d0\u00d7\u00db\u00e0\u00e4\u00ec\u00f2\u00f9"+
		"\u0104\u010b\u0111\u0117\u011f\u0127\u012d\u0132\u0137\u013d\u014a\u014f"+
		"\u0154\u0159\u0165\u0170\u0177\u017a\u0186\u0189\u0197\u01a2\u01ab\u01b2"+
		"\u01bf\u01ca\u01d2\u01da\u01de\u01e3\u01e6\u01ec\u01f6\u01ff\u0203\u0209";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}