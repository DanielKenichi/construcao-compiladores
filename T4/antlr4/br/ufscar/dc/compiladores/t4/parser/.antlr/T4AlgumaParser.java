// Generated from /home/lucky/UFSCAR/Compiladores/construcao-compiladores/T4/antlr4/br/ufscar/dc/compiladores/t4/parser/T4Alguma.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class T4AlgumaParser extends Parser {
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
	public String getGrammarFileName() { return "T4Alguma.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public T4AlgumaParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramaContext extends ParserRuleContext {
		public DeclaracoesContext declaracoes() {
			return getRuleContext(DeclaracoesContext.class,0);
		}
		public TerminalNode ALGORITMO() { return getToken(T4AlgumaParser.ALGORITMO, 0); }
		public CorpoContext corpo() {
			return getRuleContext(CorpoContext.class,0);
		}
		public TerminalNode FIM_ALGORITMO() { return getToken(T4AlgumaParser.FIM_ALGORITMO, 0); }
		public ProgramaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_programa; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterPrograma(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitPrograma(this);
		}
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
		public List<Declaracoes_variaveisContext> declaracoes_variaveis() {
			return getRuleContexts(Declaracoes_variaveisContext.class);
		}
		public Declaracoes_variaveisContext declaracoes_variaveis(int i) {
			return getRuleContext(Declaracoes_variaveisContext.class,i);
		}
		public List<Declaracoes_funcoesContext> declaracoes_funcoes() {
			return getRuleContexts(Declaracoes_funcoesContext.class);
		}
		public Declaracoes_funcoesContext declaracoes_funcoes(int i) {
			return getRuleContext(Declaracoes_funcoesContext.class,i);
		}
		public DeclaracoesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracoes; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterDeclaracoes(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitDeclaracoes(this);
		}
	}

	public final DeclaracoesContext declaracoes() throws RecognitionException {
		DeclaracoesContext _localctx = new DeclaracoesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_declaracoes);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(103);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 11005853736L) != 0)) {
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
				setState(105);
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
	public static class Declaracoes_variaveisContext extends ParserRuleContext {
		public TerminalNode DECLARE() { return getToken(T4AlgumaParser.DECLARE, 0); }
		public VariavelContext variavel() {
			return getRuleContext(VariavelContext.class,0);
		}
		public TerminalNode CONSTANTE() { return getToken(T4AlgumaParser.CONSTANTE, 0); }
		public TerminalNode IDENT() { return getToken(T4AlgumaParser.IDENT, 0); }
		public TerminalNode DOISPONTOS() { return getToken(T4AlgumaParser.DOISPONTOS, 0); }
		public Tipo_basicoContext tipo_basico() {
			return getRuleContext(Tipo_basicoContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(T4AlgumaParser.IGUAL, 0); }
		public Valor_constanteContext valor_constante() {
			return getRuleContext(Valor_constanteContext.class,0);
		}
		public TerminalNode TIPO() { return getToken(T4AlgumaParser.TIPO, 0); }
		public RegistroContext registro() {
			return getRuleContext(RegistroContext.class,0);
		}
		public Declaracoes_variaveisContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracoes_variaveis; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterDeclaracoes_variaveis(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitDeclaracoes_variaveis(this);
		}
	}

	public final Declaracoes_variaveisContext declaracoes_variaveis() throws RecognitionException {
		Declaracoes_variaveisContext _localctx = new Declaracoes_variaveisContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_declaracoes_variaveis);
		try {
			setState(119);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DECLARE:
				enterOuterAlt(_localctx, 1);
				{
				setState(106);
				match(DECLARE);
				setState(107);
				variavel();
				}
				break;
			case CONSTANTE:
				enterOuterAlt(_localctx, 2);
				{
				setState(108);
				match(CONSTANTE);
				setState(109);
				match(IDENT);
				setState(110);
				match(DOISPONTOS);
				setState(111);
				tipo_basico();
				setState(112);
				match(IGUAL);
				setState(113);
				valor_constante();
				}
				break;
			case TIPO:
				enterOuterAlt(_localctx, 3);
				{
				setState(115);
				match(TIPO);
				setState(116);
				match(IDENT);
				setState(117);
				match(DOISPONTOS);
				setState(118);
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
		public TerminalNode DOISPONTOS() { return getToken(T4AlgumaParser.DOISPONTOS, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public List<TerminalNode> VIRGULA() { return getTokens(T4AlgumaParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T4AlgumaParser.VIRGULA, i);
		}
		public VariavelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variavel; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterVariavel(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitVariavel(this);
		}
	}

	public final VariavelContext variavel() throws RecognitionException {
		VariavelContext _localctx = new VariavelContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_variavel);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(121);
			identificador();
			setState(126);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(122);
				match(VIRGULA);
				setState(123);
				identificador();
				}
				}
				setState(128);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(129);
			match(DOISPONTOS);
			setState(130);
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
		public List<TerminalNode> IDENT() { return getTokens(T4AlgumaParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(T4AlgumaParser.IDENT, i);
		}
		public List<TerminalNode> PONTO() { return getTokens(T4AlgumaParser.PONTO); }
		public TerminalNode PONTO(int i) {
			return getToken(T4AlgumaParser.PONTO, i);
		}
		public List<TerminalNode> ABRECHAVE() { return getTokens(T4AlgumaParser.ABRECHAVE); }
		public TerminalNode ABRECHAVE(int i) {
			return getToken(T4AlgumaParser.ABRECHAVE, i);
		}
		public List<Exp_aritmeticaContext> exp_aritmetica() {
			return getRuleContexts(Exp_aritmeticaContext.class);
		}
		public Exp_aritmeticaContext exp_aritmetica(int i) {
			return getRuleContext(Exp_aritmeticaContext.class,i);
		}
		public List<TerminalNode> FECHACHAVE() { return getTokens(T4AlgumaParser.FECHACHAVE); }
		public TerminalNode FECHACHAVE(int i) {
			return getToken(T4AlgumaParser.FECHACHAVE, i);
		}
		public IdentificadorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identificador; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterIdentificador(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitIdentificador(this);
		}
	}

	public final IdentificadorContext identificador() throws RecognitionException {
		IdentificadorContext _localctx = new IdentificadorContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_identificador);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			match(IDENT);
			setState(137);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PONTO) {
				{
				{
				setState(133);
				match(PONTO);
				setState(134);
				match(IDENT);
				}
				}
				setState(139);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(146);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ABRECHAVE) {
				{
				{
				setState(140);
				match(ABRECHAVE);
				setState(141);
				exp_aritmetica();
				setState(142);
				match(FECHACHAVE);
				}
				}
				setState(148);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterExp_aritmetica(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitExp_aritmetica(this);
		}
	}

	public final Exp_aritmeticaContext exp_aritmetica() throws RecognitionException {
		Exp_aritmeticaContext _localctx = new Exp_aritmeticaContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_exp_aritmetica);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			termo();
			setState(155);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(150);
					op1();
					setState(151);
					termo();
					}
					} 
				}
				setState(157);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterTermo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitTermo(this);
		}
	}

	public final TermoContext termo() throws RecognitionException {
		TermoContext _localctx = new TermoContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_termo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(158);
			fator();
			setState(164);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DIVISAO || _la==MULTIPLICACAO) {
				{
				{
				setState(159);
				op2();
				setState(160);
				fator();
				}
				}
				setState(166);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterFator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitFator(this);
		}
	}

	public final FatorContext fator() throws RecognitionException {
		FatorContext _localctx = new FatorContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_fator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			parcela();
			setState(173);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==RESTO) {
				{
				{
				setState(168);
				op3();
				setState(169);
				parcela();
				}
				}
				setState(175);
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
		public TerminalNode SOMA() { return getToken(T4AlgumaParser.SOMA, 0); }
		public TerminalNode SUBTRACAO() { return getToken(T4AlgumaParser.SUBTRACAO, 0); }
		public Op1Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op1; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterOp1(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitOp1(this);
		}
	}

	public final Op1Context op1() throws RecognitionException {
		Op1Context _localctx = new Op1Context(_ctx, getState());
		enterRule(_localctx, 16, RULE_op1);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(176);
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
		public TerminalNode MULTIPLICACAO() { return getToken(T4AlgumaParser.MULTIPLICACAO, 0); }
		public TerminalNode DIVISAO() { return getToken(T4AlgumaParser.DIVISAO, 0); }
		public Op2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op2; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterOp2(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitOp2(this);
		}
	}

	public final Op2Context op2() throws RecognitionException {
		Op2Context _localctx = new Op2Context(_ctx, getState());
		enterRule(_localctx, 18, RULE_op2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
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
		public TerminalNode RESTO() { return getToken(T4AlgumaParser.RESTO, 0); }
		public Op3Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op3; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterOp3(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitOp3(this);
		}
	}

	public final Op3Context op3() throws RecognitionException {
		Op3Context _localctx = new Op3Context(_ctx, getState());
		enterRule(_localctx, 20, RULE_op3);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(180);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterParcela(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitParcela(this);
		}
	}

	public final ParcelaContext parcela() throws RecognitionException {
		ParcelaContext _localctx = new ParcelaContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_parcela);
		int _la;
		try {
			setState(187);
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
				setState(183);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SUBTRACAO) {
					{
					setState(182);
					op_unario();
					}
				}

				setState(185);
				parcela_unario();
				}
				break;
			case ENDERECO:
			case CADEIA:
				enterOuterAlt(_localctx, 2);
				{
				setState(186);
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
		public TerminalNode PONTEIRO() { return getToken(T4AlgumaParser.PONTEIRO, 0); }
		public TerminalNode IDENT() { return getToken(T4AlgumaParser.IDENT, 0); }
		public TerminalNode ABREPAR() { return getToken(T4AlgumaParser.ABREPAR, 0); }
		public List<ExpressaoContext> expressao() {
			return getRuleContexts(ExpressaoContext.class);
		}
		public ExpressaoContext expressao(int i) {
			return getRuleContext(ExpressaoContext.class,i);
		}
		public TerminalNode FECHAPAR() { return getToken(T4AlgumaParser.FECHAPAR, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(T4AlgumaParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T4AlgumaParser.VIRGULA, i);
		}
		public TerminalNode NUM_INT() { return getToken(T4AlgumaParser.NUM_INT, 0); }
		public TerminalNode NUM_REAL() { return getToken(T4AlgumaParser.NUM_REAL, 0); }
		public Parcela_unarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parcela_unario; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterParcela_unario(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitParcela_unario(this);
		}
	}

	public final Parcela_unarioContext parcela_unario() throws RecognitionException {
		Parcela_unarioContext _localctx = new Parcela_unarioContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_parcela_unario);
		int _la;
		try {
			setState(211);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(190);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PONTEIRO) {
					{
					setState(189);
					match(PONTEIRO);
					}
				}

				setState(192);
				identificador();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(193);
				match(IDENT);
				setState(194);
				match(ABREPAR);
				setState(195);
				expressao();
				setState(200);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==VIRGULA) {
					{
					{
					setState(196);
					match(VIRGULA);
					setState(197);
					expressao();
					}
					}
					setState(202);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(203);
				match(FECHAPAR);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(205);
				match(NUM_INT);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(206);
				match(NUM_REAL);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(207);
				match(ABREPAR);
				setState(208);
				expressao();
				setState(209);
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
		public TerminalNode SUBTRACAO() { return getToken(T4AlgumaParser.SUBTRACAO, 0); }
		public Op_unarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_unario; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterOp_unario(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitOp_unario(this);
		}
	}

	public final Op_unarioContext op_unario() throws RecognitionException {
		Op_unarioContext _localctx = new Op_unarioContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_op_unario);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(213);
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
		public TerminalNode ENDERECO() { return getToken(T4AlgumaParser.ENDERECO, 0); }
		public IdentificadorContext identificador() {
			return getRuleContext(IdentificadorContext.class,0);
		}
		public TerminalNode CADEIA() { return getToken(T4AlgumaParser.CADEIA, 0); }
		public Parcela_nao_unarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parcela_nao_unario; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterParcela_nao_unario(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitParcela_nao_unario(this);
		}
	}

	public final Parcela_nao_unarioContext parcela_nao_unario() throws RecognitionException {
		Parcela_nao_unarioContext _localctx = new Parcela_nao_unarioContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_parcela_nao_unario);
		try {
			setState(218);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENDERECO:
				enterOuterAlt(_localctx, 1);
				{
				setState(215);
				match(ENDERECO);
				setState(216);
				identificador();
				}
				break;
			case CADEIA:
				enterOuterAlt(_localctx, 2);
				{
				setState(217);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterTipo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitTipo(this);
		}
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_tipo);
		try {
			setState(222);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case REGISTRO:
				enterOuterAlt(_localctx, 1);
				{
				setState(220);
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
				setState(221);
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
		public TerminalNode LITERAL() { return getToken(T4AlgumaParser.LITERAL, 0); }
		public TerminalNode INTEIRO() { return getToken(T4AlgumaParser.INTEIRO, 0); }
		public TerminalNode REAL() { return getToken(T4AlgumaParser.REAL, 0); }
		public TerminalNode LOGICO() { return getToken(T4AlgumaParser.LOGICO, 0); }
		public Tipo_basicoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_basico; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterTipo_basico(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitTipo_basico(this);
		}
	}

	public final Tipo_basicoContext tipo_basico() throws RecognitionException {
		Tipo_basicoContext _localctx = new Tipo_basicoContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_tipo_basico);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(224);
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
		public TerminalNode IDENT() { return getToken(T4AlgumaParser.IDENT, 0); }
		public TerminalNode PONTEIRO() { return getToken(T4AlgumaParser.PONTEIRO, 0); }
		public Tipo_variavelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_variavel; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterTipo_variavel(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitTipo_variavel(this);
		}
	}

	public final Tipo_variavelContext tipo_variavel() throws RecognitionException {
		Tipo_variavelContext _localctx = new Tipo_variavelContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_tipo_variavel);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(227);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PONTEIRO) {
				{
				setState(226);
				match(PONTEIRO);
				}
			}

			setState(231);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LITERAL:
			case INTEIRO:
			case REAL:
			case LOGICO:
				{
				setState(229);
				tipo_basico();
				}
				break;
			case IDENT:
				{
				setState(230);
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
		public TerminalNode CADEIA() { return getToken(T4AlgumaParser.CADEIA, 0); }
		public TerminalNode NUM_INT() { return getToken(T4AlgumaParser.NUM_INT, 0); }
		public TerminalNode NUM_REAL() { return getToken(T4AlgumaParser.NUM_REAL, 0); }
		public TerminalNode VERDADEIRO() { return getToken(T4AlgumaParser.VERDADEIRO, 0); }
		public TerminalNode FALSO() { return getToken(T4AlgumaParser.FALSO, 0); }
		public Valor_constanteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valor_constante; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterValor_constante(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitValor_constante(this);
		}
	}

	public final Valor_constanteContext valor_constante() throws RecognitionException {
		Valor_constanteContext _localctx = new Valor_constanteContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_valor_constante);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(233);
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
		public TerminalNode REGISTRO() { return getToken(T4AlgumaParser.REGISTRO, 0); }
		public TerminalNode FIM_REGISTRO() { return getToken(T4AlgumaParser.FIM_REGISTRO, 0); }
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterRegistro(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitRegistro(this);
		}
	}

	public final RegistroContext registro() throws RecognitionException {
		RegistroContext _localctx = new RegistroContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_registro);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			match(REGISTRO);
			setState(239);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENT) {
				{
				{
				setState(236);
				variavel();
				}
				}
				setState(241);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(242);
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
		public TerminalNode DOISPONTOS() { return getToken(T4AlgumaParser.DOISPONTOS, 0); }
		public Tipo_variavelContext tipo_variavel() {
			return getRuleContext(Tipo_variavelContext.class,0);
		}
		public TerminalNode VAR() { return getToken(T4AlgumaParser.VAR, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(T4AlgumaParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T4AlgumaParser.VIRGULA, i);
		}
		public ParametroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametro; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterParametro(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitParametro(this);
		}
	}

	public final ParametroContext parametro() throws RecognitionException {
		ParametroContext _localctx = new ParametroContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_parametro);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(245);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==VAR) {
				{
				setState(244);
				match(VAR);
				}
			}

			setState(247);
			identificador();
			setState(252);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(248);
				match(VIRGULA);
				setState(249);
				identificador();
				}
				}
				setState(254);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(255);
			match(DOISPONTOS);
			setState(256);
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
		public List<TerminalNode> VIRGULA() { return getTokens(T4AlgumaParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T4AlgumaParser.VIRGULA, i);
		}
		public ParametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametros; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterParametros(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitParametros(this);
		}
	}

	public final ParametrosContext parametros() throws RecognitionException {
		ParametrosContext _localctx = new ParametrosContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(258);
			parametro();
			setState(263);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(259);
				match(VIRGULA);
				setState(260);
				parametro();
				}
				}
				setState(265);
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
		public TerminalNode PROCEDIMENTO() { return getToken(T4AlgumaParser.PROCEDIMENTO, 0); }
		public TerminalNode IDENT() { return getToken(T4AlgumaParser.IDENT, 0); }
		public TerminalNode ABREPAR() { return getToken(T4AlgumaParser.ABREPAR, 0); }
		public TerminalNode FECHAPAR() { return getToken(T4AlgumaParser.FECHAPAR, 0); }
		public TerminalNode FIM_PROCEDIMENTO() { return getToken(T4AlgumaParser.FIM_PROCEDIMENTO, 0); }
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
		public TerminalNode FUNCAO() { return getToken(T4AlgumaParser.FUNCAO, 0); }
		public TerminalNode DOISPONTOS() { return getToken(T4AlgumaParser.DOISPONTOS, 0); }
		public Tipo_variavelContext tipo_variavel() {
			return getRuleContext(Tipo_variavelContext.class,0);
		}
		public TerminalNode FIM_FUNCAO() { return getToken(T4AlgumaParser.FIM_FUNCAO, 0); }
		public Declaracoes_funcoesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracoes_funcoes; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterDeclaracoes_funcoes(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitDeclaracoes_funcoes(this);
		}
	}

	public final Declaracoes_funcoesContext declaracoes_funcoes() throws RecognitionException {
		Declaracoes_funcoesContext _localctx = new Declaracoes_funcoesContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_declaracoes_funcoes);
		int _la;
		try {
			setState(309);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PROCEDIMENTO:
				enterOuterAlt(_localctx, 1);
				{
				setState(266);
				match(PROCEDIMENTO);
				setState(267);
				match(IDENT);
				setState(268);
				match(ABREPAR);
				setState(270);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==VAR || _la==IDENT) {
					{
					setState(269);
					parametros();
					}
				}

				setState(272);
				match(FECHAPAR);
				setState(276);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 268435496L) != 0)) {
					{
					{
					setState(273);
					declaracoes_variaveis();
					}
					}
					setState(278);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(282);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
					{
					{
					setState(279);
					cmd();
					}
					}
					setState(284);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(285);
				match(FIM_PROCEDIMENTO);
				}
				break;
			case FUNCAO:
				enterOuterAlt(_localctx, 2);
				{
				setState(286);
				match(FUNCAO);
				setState(287);
				match(IDENT);
				setState(288);
				match(ABREPAR);
				setState(290);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==VAR || _la==IDENT) {
					{
					setState(289);
					parametros();
					}
				}

				setState(292);
				match(FECHAPAR);
				setState(293);
				match(DOISPONTOS);
				setState(294);
				tipo_variavel();
				setState(298);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 268435496L) != 0)) {
					{
					{
					setState(295);
					declaracoes_variaveis();
					}
					}
					setState(300);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(304);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
					{
					{
					setState(301);
					cmd();
					}
					}
					setState(306);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(307);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCorpo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCorpo(this);
		}
	}

	public final CorpoContext corpo() throws RecognitionException {
		CorpoContext _localctx = new CorpoContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_corpo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(314);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 268435496L) != 0)) {
				{
				{
				setState(311);
				declaracoes_variaveis();
				}
				}
				setState(316);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(320);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(317);
				cmd();
				}
				}
				setState(322);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmd(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmd(this);
		}
	}

	public final CmdContext cmd() throws RecognitionException {
		CmdContext _localctx = new CmdContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_cmd);
		try {
			setState(333);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(323);
				cmdLeia();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(324);
				cmdEscreva();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(325);
				cmdSe();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(326);
				cmdCaso();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(327);
				cmdPara();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(328);
				cmdEnquanto();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(329);
				cmdFaca();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(330);
				cmdAtribuicao();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(331);
				cmdChamada();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(332);
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
		public TerminalNode LEIA() { return getToken(T4AlgumaParser.LEIA, 0); }
		public TerminalNode ABREPAR() { return getToken(T4AlgumaParser.ABREPAR, 0); }
		public List<IdentificadorContext> identificador() {
			return getRuleContexts(IdentificadorContext.class);
		}
		public IdentificadorContext identificador(int i) {
			return getRuleContext(IdentificadorContext.class,i);
		}
		public TerminalNode FECHAPAR() { return getToken(T4AlgumaParser.FECHAPAR, 0); }
		public List<TerminalNode> PONTEIRO() { return getTokens(T4AlgumaParser.PONTEIRO); }
		public TerminalNode PONTEIRO(int i) {
			return getToken(T4AlgumaParser.PONTEIRO, i);
		}
		public List<TerminalNode> VIRGULA() { return getTokens(T4AlgumaParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T4AlgumaParser.VIRGULA, i);
		}
		public CmdLeiaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdLeia; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmdLeia(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmdLeia(this);
		}
	}

	public final CmdLeiaContext cmdLeia() throws RecognitionException {
		CmdLeiaContext _localctx = new CmdLeiaContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_cmdLeia);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(335);
			match(LEIA);
			setState(336);
			match(ABREPAR);
			setState(338);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PONTEIRO) {
				{
				setState(337);
				match(PONTEIRO);
				}
			}

			setState(340);
			identificador();
			setState(348);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(341);
				match(VIRGULA);
				setState(343);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PONTEIRO) {
					{
					setState(342);
					match(PONTEIRO);
					}
				}

				setState(345);
				identificador();
				}
				}
				setState(350);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(351);
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
		public TerminalNode ESCREVA() { return getToken(T4AlgumaParser.ESCREVA, 0); }
		public TerminalNode ABREPAR() { return getToken(T4AlgumaParser.ABREPAR, 0); }
		public List<ExpressaoContext> expressao() {
			return getRuleContexts(ExpressaoContext.class);
		}
		public ExpressaoContext expressao(int i) {
			return getRuleContext(ExpressaoContext.class,i);
		}
		public TerminalNode FECHAPAR() { return getToken(T4AlgumaParser.FECHAPAR, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(T4AlgumaParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T4AlgumaParser.VIRGULA, i);
		}
		public CmdEscrevaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdEscreva; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmdEscreva(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmdEscreva(this);
		}
	}

	public final CmdEscrevaContext cmdEscreva() throws RecognitionException {
		CmdEscrevaContext _localctx = new CmdEscrevaContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_cmdEscreva);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(353);
			match(ESCREVA);
			setState(354);
			match(ABREPAR);
			setState(355);
			expressao();
			setState(360);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(356);
				match(VIRGULA);
				setState(357);
				expressao();
				}
				}
				setState(362);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(363);
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
		public CmdContext cmd;
		public List<CmdContext> seCmds = new ArrayList<CmdContext>();
		public List<CmdContext> senaoCmds = new ArrayList<CmdContext>();
		public TerminalNode SE() { return getToken(T4AlgumaParser.SE, 0); }
		public ExpressaoContext expressao() {
			return getRuleContext(ExpressaoContext.class,0);
		}
		public TerminalNode ENTAO() { return getToken(T4AlgumaParser.ENTAO, 0); }
		public TerminalNode FIM_SE() { return getToken(T4AlgumaParser.FIM_SE, 0); }
		public TerminalNode SENAO() { return getToken(T4AlgumaParser.SENAO, 0); }
		public List<CmdContext> cmd() {
			return getRuleContexts(CmdContext.class);
		}
		public CmdContext cmd(int i) {
			return getRuleContext(CmdContext.class,i);
		}
		public CmdSeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdSe; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmdSe(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmdSe(this);
		}
	}

	public final CmdSeContext cmdSe() throws RecognitionException {
		CmdSeContext _localctx = new CmdSeContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_cmdSe);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(365);
			match(SE);
			setState(366);
			expressao();
			setState(367);
			match(ENTAO);
			setState(371);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(368);
				((CmdSeContext)_localctx).cmd = cmd();
				((CmdSeContext)_localctx).seCmds.add(((CmdSeContext)_localctx).cmd);
				}
				}
				setState(373);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(381);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SENAO) {
				{
				setState(374);
				match(SENAO);
				setState(378);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
					{
					{
					setState(375);
					((CmdSeContext)_localctx).cmd = cmd();
					((CmdSeContext)_localctx).senaoCmds.add(((CmdSeContext)_localctx).cmd);
					}
					}
					setState(380);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(383);
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
		public TerminalNode CASO() { return getToken(T4AlgumaParser.CASO, 0); }
		public Exp_aritmeticaContext exp_aritmetica() {
			return getRuleContext(Exp_aritmeticaContext.class,0);
		}
		public TerminalNode SEJA() { return getToken(T4AlgumaParser.SEJA, 0); }
		public SelecaoContext selecao() {
			return getRuleContext(SelecaoContext.class,0);
		}
		public TerminalNode FIM_CASO() { return getToken(T4AlgumaParser.FIM_CASO, 0); }
		public TerminalNode SENAO() { return getToken(T4AlgumaParser.SENAO, 0); }
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmdCaso(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmdCaso(this);
		}
	}

	public final CmdCasoContext cmdCaso() throws RecognitionException {
		CmdCasoContext _localctx = new CmdCasoContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_cmdCaso);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(385);
			match(CASO);
			setState(386);
			exp_aritmetica();
			setState(387);
			match(SEJA);
			setState(388);
			selecao();
			setState(396);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SENAO) {
				{
				setState(389);
				match(SENAO);
				setState(393);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
					{
					{
					setState(390);
					cmd();
					}
					}
					setState(395);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(398);
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
		public Exp_aritmeticaContext exp1;
		public Exp_aritmeticaContext exp2;
		public TerminalNode PARA() { return getToken(T4AlgumaParser.PARA, 0); }
		public TerminalNode IDENT() { return getToken(T4AlgumaParser.IDENT, 0); }
		public TerminalNode ATRIBUICAO() { return getToken(T4AlgumaParser.ATRIBUICAO, 0); }
		public TerminalNode ATE() { return getToken(T4AlgumaParser.ATE, 0); }
		public TerminalNode FACA() { return getToken(T4AlgumaParser.FACA, 0); }
		public TerminalNode FIM_PARA() { return getToken(T4AlgumaParser.FIM_PARA, 0); }
		public List<Exp_aritmeticaContext> exp_aritmetica() {
			return getRuleContexts(Exp_aritmeticaContext.class);
		}
		public Exp_aritmeticaContext exp_aritmetica(int i) {
			return getRuleContext(Exp_aritmeticaContext.class,i);
		}
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmdPara(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmdPara(this);
		}
	}

	public final CmdParaContext cmdPara() throws RecognitionException {
		CmdParaContext _localctx = new CmdParaContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_cmdPara);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(400);
			match(PARA);
			setState(401);
			match(IDENT);
			setState(402);
			match(ATRIBUICAO);
			setState(403);
			((CmdParaContext)_localctx).exp1 = exp_aritmetica();
			setState(404);
			match(ATE);
			setState(405);
			((CmdParaContext)_localctx).exp2 = exp_aritmetica();
			setState(406);
			match(FACA);
			setState(410);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(407);
				cmd();
				}
				}
				setState(412);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(413);
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
		public TerminalNode ENQUANTO() { return getToken(T4AlgumaParser.ENQUANTO, 0); }
		public ExpressaoContext expressao() {
			return getRuleContext(ExpressaoContext.class,0);
		}
		public TerminalNode FACA() { return getToken(T4AlgumaParser.FACA, 0); }
		public TerminalNode FIM_ENQUANTO() { return getToken(T4AlgumaParser.FIM_ENQUANTO, 0); }
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmdEnquanto(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmdEnquanto(this);
		}
	}

	public final CmdEnquantoContext cmdEnquanto() throws RecognitionException {
		CmdEnquantoContext _localctx = new CmdEnquantoContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_cmdEnquanto);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(415);
			match(ENQUANTO);
			setState(416);
			expressao();
			setState(417);
			match(FACA);
			setState(421);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(418);
				cmd();
				}
				}
				setState(423);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(424);
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
		public TerminalNode FACA() { return getToken(T4AlgumaParser.FACA, 0); }
		public TerminalNode ATE() { return getToken(T4AlgumaParser.ATE, 0); }
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmdFaca(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmdFaca(this);
		}
	}

	public final CmdFacaContext cmdFaca() throws RecognitionException {
		CmdFacaContext _localctx = new CmdFacaContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_cmdFaca);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(426);
			match(FACA);
			setState(430);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(427);
				cmd();
				}
				}
				setState(432);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(433);
			match(ATE);
			setState(434);
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
		public TerminalNode ATRIBUICAO() { return getToken(T4AlgumaParser.ATRIBUICAO, 0); }
		public ExpressaoContext expressao() {
			return getRuleContext(ExpressaoContext.class,0);
		}
		public TerminalNode PONTEIRO() { return getToken(T4AlgumaParser.PONTEIRO, 0); }
		public CmdAtribuicaoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdAtribuicao; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmdAtribuicao(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmdAtribuicao(this);
		}
	}

	public final CmdAtribuicaoContext cmdAtribuicao() throws RecognitionException {
		CmdAtribuicaoContext _localctx = new CmdAtribuicaoContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_cmdAtribuicao);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(437);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PONTEIRO) {
				{
				setState(436);
				match(PONTEIRO);
				}
			}

			setState(439);
			identificador();
			setState(440);
			match(ATRIBUICAO);
			setState(441);
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
		public TerminalNode IDENT() { return getToken(T4AlgumaParser.IDENT, 0); }
		public TerminalNode ABREPAR() { return getToken(T4AlgumaParser.ABREPAR, 0); }
		public List<ExpressaoContext> expressao() {
			return getRuleContexts(ExpressaoContext.class);
		}
		public ExpressaoContext expressao(int i) {
			return getRuleContext(ExpressaoContext.class,i);
		}
		public TerminalNode FECHAPAR() { return getToken(T4AlgumaParser.FECHAPAR, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(T4AlgumaParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T4AlgumaParser.VIRGULA, i);
		}
		public CmdChamadaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdChamada; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmdChamada(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmdChamada(this);
		}
	}

	public final CmdChamadaContext cmdChamada() throws RecognitionException {
		CmdChamadaContext _localctx = new CmdChamadaContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_cmdChamada);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(443);
			match(IDENT);
			setState(444);
			match(ABREPAR);
			setState(445);
			expressao();
			setState(450);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(446);
				match(VIRGULA);
				setState(447);
				expressao();
				}
				}
				setState(452);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(453);
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
		public TerminalNode RETORNE() { return getToken(T4AlgumaParser.RETORNE, 0); }
		public ExpressaoContext expressao() {
			return getRuleContext(ExpressaoContext.class,0);
		}
		public CmdRetorneContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cmdRetorne; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterCmdRetorne(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitCmdRetorne(this);
		}
	}

	public final CmdRetorneContext cmdRetorne() throws RecognitionException {
		CmdRetorneContext _localctx = new CmdRetorneContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_cmdRetorne);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(455);
			match(RETORNE);
			setState(456);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterSelecao(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitSelecao(this);
		}
	}

	public final SelecaoContext selecao() throws RecognitionException {
		SelecaoContext _localctx = new SelecaoContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_selecao);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(461);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SUBTRACAO || _la==NUM_INT) {
				{
				{
				setState(458);
				item_selecao();
				}
				}
				setState(463);
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
		public TerminalNode DOISPONTOS() { return getToken(T4AlgumaParser.DOISPONTOS, 0); }
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterItem_selecao(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitItem_selecao(this);
		}
	}

	public final Item_selecaoContext item_selecao() throws RecognitionException {
		Item_selecaoContext _localctx = new Item_selecaoContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_item_selecao);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(464);
			constantes();
			setState(465);
			match(DOISPONTOS);
			setState(469);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 580542146808977L) != 0)) {
				{
				{
				setState(466);
				cmd();
				}
				}
				setState(471);
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
		public List<TerminalNode> VIRGULA() { return getTokens(T4AlgumaParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(T4AlgumaParser.VIRGULA, i);
		}
		public ConstantesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constantes; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterConstantes(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitConstantes(this);
		}
	}

	public final ConstantesContext constantes() throws RecognitionException {
		ConstantesContext _localctx = new ConstantesContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_constantes);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(472);
			numero_intervalo();
			setState(477);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(473);
				match(VIRGULA);
				setState(474);
				numero_intervalo();
				}
				}
				setState(479);
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
		public List<TerminalNode> NUM_INT() { return getTokens(T4AlgumaParser.NUM_INT); }
		public TerminalNode NUM_INT(int i) {
			return getToken(T4AlgumaParser.NUM_INT, i);
		}
		public List<Op_unarioContext> op_unario() {
			return getRuleContexts(Op_unarioContext.class);
		}
		public Op_unarioContext op_unario(int i) {
			return getRuleContext(Op_unarioContext.class,i);
		}
		public TerminalNode INTERVALO() { return getToken(T4AlgumaParser.INTERVALO, 0); }
		public Numero_intervaloContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numero_intervalo; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterNumero_intervalo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitNumero_intervalo(this);
		}
	}

	public final Numero_intervaloContext numero_intervalo() throws RecognitionException {
		Numero_intervaloContext _localctx = new Numero_intervaloContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_numero_intervalo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(481);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SUBTRACAO) {
				{
				setState(480);
				op_unario();
				}
			}

			setState(483);
			match(NUM_INT);
			setState(489);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERVALO) {
				{
				setState(484);
				match(INTERVALO);
				setState(486);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SUBTRACAO) {
					{
					setState(485);
					op_unario();
					}
				}

				setState(488);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterExp_relacional(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitExp_relacional(this);
		}
	}

	public final Exp_relacionalContext exp_relacional() throws RecognitionException {
		Exp_relacionalContext _localctx = new Exp_relacionalContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_exp_relacional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(491);
			exp_aritmetica();
			setState(495);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 69269232549888L) != 0)) {
				{
				setState(492);
				op_relacional();
				setState(493);
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
		public TerminalNode IGUAL() { return getToken(T4AlgumaParser.IGUAL, 0); }
		public TerminalNode DIFERENTE() { return getToken(T4AlgumaParser.DIFERENTE, 0); }
		public TerminalNode MAIORIGUAL() { return getToken(T4AlgumaParser.MAIORIGUAL, 0); }
		public TerminalNode MENORIGUAL() { return getToken(T4AlgumaParser.MENORIGUAL, 0); }
		public TerminalNode MAIOR() { return getToken(T4AlgumaParser.MAIOR, 0); }
		public TerminalNode MENOR() { return getToken(T4AlgumaParser.MENOR, 0); }
		public Op_relacionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_relacional; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterOp_relacional(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitOp_relacional(this);
		}
	}

	public final Op_relacionalContext op_relacional() throws RecognitionException {
		Op_relacionalContext _localctx = new Op_relacionalContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_op_relacional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(497);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterExpressao(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitExpressao(this);
		}
	}

	public final ExpressaoContext expressao() throws RecognitionException {
		ExpressaoContext _localctx = new ExpressaoContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_expressao);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(499);
			termo_logico();
			setState(505);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OU) {
				{
				{
				setState(500);
				op_logico_1();
				setState(501);
				termo_logico();
				}
				}
				setState(507);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterTermo_logico(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitTermo_logico(this);
		}
	}

	public final Termo_logicoContext termo_logico() throws RecognitionException {
		Termo_logicoContext _localctx = new Termo_logicoContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_termo_logico);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(508);
			fator_logico();
			setState(514);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==E) {
				{
				{
				setState(509);
				op_logico_2();
				setState(510);
				fator_logico();
				}
				}
				setState(516);
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
		public TerminalNode NAO() { return getToken(T4AlgumaParser.NAO, 0); }
		public Fator_logicoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fator_logico; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterFator_logico(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitFator_logico(this);
		}
	}

	public final Fator_logicoContext fator_logico() throws RecognitionException {
		Fator_logicoContext _localctx = new Fator_logicoContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_fator_logico);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(518);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NAO) {
				{
				setState(517);
				match(NAO);
				}
			}

			setState(520);
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
		public TerminalNode VERDADEIRO() { return getToken(T4AlgumaParser.VERDADEIRO, 0); }
		public TerminalNode FALSO() { return getToken(T4AlgumaParser.FALSO, 0); }
		public Exp_relacionalContext exp_relacional() {
			return getRuleContext(Exp_relacionalContext.class,0);
		}
		public Parcela_logicaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parcela_logica; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterParcela_logica(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitParcela_logica(this);
		}
	}

	public final Parcela_logicaContext parcela_logica() throws RecognitionException {
		Parcela_logicaContext _localctx = new Parcela_logicaContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_parcela_logica);
		int _la;
		try {
			setState(524);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VERDADEIRO:
			case FALSO:
				enterOuterAlt(_localctx, 1);
				{
				setState(522);
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
				setState(523);
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
		public TerminalNode OU() { return getToken(T4AlgumaParser.OU, 0); }
		public Op_logico_1Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_logico_1; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterOp_logico_1(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitOp_logico_1(this);
		}
	}

	public final Op_logico_1Context op_logico_1() throws RecognitionException {
		Op_logico_1Context _localctx = new Op_logico_1Context(_ctx, getState());
		enterRule(_localctx, 90, RULE_op_logico_1);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(526);
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
		public TerminalNode E() { return getToken(T4AlgumaParser.E, 0); }
		public Op_logico_2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op_logico_2; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).enterOp_logico_2(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof T4AlgumaListener ) ((T4AlgumaListener)listener).exitOp_logico_2(this);
		}
	}

	public final Op_logico_2Context op_logico_2() throws RecognitionException {
		Op_logico_2Context _localctx = new Op_logico_2Context(_ctx, getState());
		enterRule(_localctx, 92, RULE_op_logico_2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(528);
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
		"\u0004\u0001E\u0213\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"\u0001\u0000\u0001\u0001\u0001\u0001\u0005\u0001f\b\u0001\n\u0001\f\u0001"+
		"i\t\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0003\u0002x\b\u0002\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0005\u0003}\b\u0003\n\u0003\f\u0003\u0080\t\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0005"+
		"\u0004\u0088\b\u0004\n\u0004\f\u0004\u008b\t\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0005\u0004\u0091\b\u0004\n\u0004\f\u0004\u0094"+
		"\t\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005\u009a"+
		"\b\u0005\n\u0005\f\u0005\u009d\t\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0005\u0006\u00a3\b\u0006\n\u0006\f\u0006\u00a6\t\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007\u00ac\b\u0007\n"+
		"\u0007\f\u0007\u00af\t\u0007\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001"+
		"\n\u0001\u000b\u0003\u000b\u00b8\b\u000b\u0001\u000b\u0001\u000b\u0003"+
		"\u000b\u00bc\b\u000b\u0001\f\u0003\f\u00bf\b\f\u0001\f\u0001\f\u0001\f"+
		"\u0001\f\u0001\f\u0001\f\u0005\f\u00c7\b\f\n\f\f\f\u00ca\t\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u00d4\b\f\u0001"+
		"\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u00db\b\u000e"+
		"\u0001\u000f\u0001\u000f\u0003\u000f\u00df\b\u000f\u0001\u0010\u0001\u0010"+
		"\u0001\u0011\u0003\u0011\u00e4\b\u0011\u0001\u0011\u0001\u0011\u0003\u0011"+
		"\u00e8\b\u0011\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0005\u0013"+
		"\u00ee\b\u0013\n\u0013\f\u0013\u00f1\t\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0014\u0003\u0014\u00f6\b\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0005"+
		"\u0014\u00fb\b\u0014\n\u0014\f\u0014\u00fe\t\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0005\u0015\u0106\b\u0015"+
		"\n\u0015\f\u0015\u0109\t\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0003\u0016\u010f\b\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u0113"+
		"\b\u0016\n\u0016\f\u0016\u0116\t\u0016\u0001\u0016\u0005\u0016\u0119\b"+
		"\u0016\n\u0016\f\u0016\u011c\t\u0016\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0003\u0016\u0123\b\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0005\u0016\u0129\b\u0016\n\u0016\f\u0016\u012c"+
		"\t\u0016\u0001\u0016\u0005\u0016\u012f\b\u0016\n\u0016\f\u0016\u0132\t"+
		"\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u0136\b\u0016\u0001\u0017\u0005"+
		"\u0017\u0139\b\u0017\n\u0017\f\u0017\u013c\t\u0017\u0001\u0017\u0005\u0017"+
		"\u013f\b\u0017\n\u0017\f\u0017\u0142\t\u0017\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0003\u0018\u014e\b\u0018\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0003\u0019\u0153\b\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0003"+
		"\u0019\u0158\b\u0019\u0001\u0019\u0005\u0019\u015b\b\u0019\n\u0019\f\u0019"+
		"\u015e\t\u0019\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0005\u001a\u0167\b\u001a\n\u001a\f\u001a\u016a"+
		"\t\u001a\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0001\u001b\u0001"+
		"\u001b\u0005\u001b\u0172\b\u001b\n\u001b\f\u001b\u0175\t\u001b\u0001\u001b"+
		"\u0001\u001b\u0005\u001b\u0179\b\u001b\n\u001b\f\u001b\u017c\t\u001b\u0003"+
		"\u001b\u017e\b\u001b\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0005\u001c\u0188\b\u001c\n"+
		"\u001c\f\u001c\u018b\t\u001c\u0003\u001c\u018d\b\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0001\u001d\u0005\u001d\u0199\b\u001d\n\u001d\f\u001d"+
		"\u019c\t\u001d\u0001\u001d\u0001\u001d\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0005\u001e\u01a4\b\u001e\n\u001e\f\u001e\u01a7\t\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0005\u001f\u01ad\b\u001f\n"+
		"\u001f\f\u001f\u01b0\t\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		" \u0003 \u01b6\b \u0001 \u0001 \u0001 \u0001 \u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0005!\u01c1\b!\n!\f!\u01c4\t!\u0001!\u0001!\u0001\"\u0001\""+
		"\u0001\"\u0001#\u0005#\u01cc\b#\n#\f#\u01cf\t#\u0001$\u0001$\u0001$\u0005"+
		"$\u01d4\b$\n$\f$\u01d7\t$\u0001%\u0001%\u0001%\u0005%\u01dc\b%\n%\f%\u01df"+
		"\t%\u0001&\u0003&\u01e2\b&\u0001&\u0001&\u0001&\u0003&\u01e7\b&\u0001"+
		"&\u0003&\u01ea\b&\u0001\'\u0001\'\u0001\'\u0001\'\u0003\'\u01f0\b\'\u0001"+
		"(\u0001(\u0001)\u0001)\u0001)\u0001)\u0005)\u01f8\b)\n)\f)\u01fb\t)\u0001"+
		"*\u0001*\u0001*\u0001*\u0005*\u0201\b*\n*\f*\u0204\t*\u0001+\u0003+\u0207"+
		"\b+\u0001+\u0001+\u0001,\u0001,\u0003,\u020d\b,\u0001-\u0001-\u0001.\u0001"+
		".\u0001.\u0000\u0000/\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012"+
		"\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJLNPRTVXZ\\"+
		"\u0000\u0006\u0001\u000078\u0002\u00005599\u0001\u0000\u0006\t\u0003\u0000"+
		"\n\u000b>?AA\u0001\u0000(-\u0001\u0000\n\u000b\u0228\u0000^\u0001\u0000"+
		"\u0000\u0000\u0002g\u0001\u0000\u0000\u0000\u0004w\u0001\u0000\u0000\u0000"+
		"\u0006y\u0001\u0000\u0000\u0000\b\u0084\u0001\u0000\u0000\u0000\n\u0095"+
		"\u0001\u0000\u0000\u0000\f\u009e\u0001\u0000\u0000\u0000\u000e\u00a7\u0001"+
		"\u0000\u0000\u0000\u0010\u00b0\u0001\u0000\u0000\u0000\u0012\u00b2\u0001"+
		"\u0000\u0000\u0000\u0014\u00b4\u0001\u0000\u0000\u0000\u0016\u00bb\u0001"+
		"\u0000\u0000\u0000\u0018\u00d3\u0001\u0000\u0000\u0000\u001a\u00d5\u0001"+
		"\u0000\u0000\u0000\u001c\u00da\u0001\u0000\u0000\u0000\u001e\u00de\u0001"+
		"\u0000\u0000\u0000 \u00e0\u0001\u0000\u0000\u0000\"\u00e3\u0001\u0000"+
		"\u0000\u0000$\u00e9\u0001\u0000\u0000\u0000&\u00eb\u0001\u0000\u0000\u0000"+
		"(\u00f5\u0001\u0000\u0000\u0000*\u0102\u0001\u0000\u0000\u0000,\u0135"+
		"\u0001\u0000\u0000\u0000.\u013a\u0001\u0000\u0000\u00000\u014d\u0001\u0000"+
		"\u0000\u00002\u014f\u0001\u0000\u0000\u00004\u0161\u0001\u0000\u0000\u0000"+
		"6\u016d\u0001\u0000\u0000\u00008\u0181\u0001\u0000\u0000\u0000:\u0190"+
		"\u0001\u0000\u0000\u0000<\u019f\u0001\u0000\u0000\u0000>\u01aa\u0001\u0000"+
		"\u0000\u0000@\u01b5\u0001\u0000\u0000\u0000B\u01bb\u0001\u0000\u0000\u0000"+
		"D\u01c7\u0001\u0000\u0000\u0000F\u01cd\u0001\u0000\u0000\u0000H\u01d0"+
		"\u0001\u0000\u0000\u0000J\u01d8\u0001\u0000\u0000\u0000L\u01e1\u0001\u0000"+
		"\u0000\u0000N\u01eb\u0001\u0000\u0000\u0000P\u01f1\u0001\u0000\u0000\u0000"+
		"R\u01f3\u0001\u0000\u0000\u0000T\u01fc\u0001\u0000\u0000\u0000V\u0206"+
		"\u0001\u0000\u0000\u0000X\u020c\u0001\u0000\u0000\u0000Z\u020e\u0001\u0000"+
		"\u0000\u0000\\\u0210\u0001\u0000\u0000\u0000^_\u0003\u0002\u0001\u0000"+
		"_`\u0005\u0002\u0000\u0000`a\u0003.\u0017\u0000ab\u0005&\u0000\u0000b"+
		"\u0001\u0001\u0000\u0000\u0000cf\u0003\u0004\u0002\u0000df\u0003,\u0016"+
		"\u0000ec\u0001\u0000\u0000\u0000ed\u0001\u0000\u0000\u0000fi\u0001\u0000"+
		"\u0000\u0000ge\u0001\u0000\u0000\u0000gh\u0001\u0000\u0000\u0000h\u0003"+
		"\u0001\u0000\u0000\u0000ig\u0001\u0000\u0000\u0000jk\u0005\u0003\u0000"+
		"\u0000kx\u0003\u0006\u0003\u0000lm\u0005\u0005\u0000\u0000mn\u0005@\u0000"+
		"\u0000no\u0005.\u0000\u0000op\u0003 \u0010\u0000pq\u0005,\u0000\u0000"+
		"qr\u0003$\u0012\u0000rx\u0001\u0000\u0000\u0000st\u0005\u001c\u0000\u0000"+
		"tu\u0005@\u0000\u0000uv\u0005.\u0000\u0000vx\u0003&\u0013\u0000wj\u0001"+
		"\u0000\u0000\u0000wl\u0001\u0000\u0000\u0000ws\u0001\u0000\u0000\u0000"+
		"x\u0005\u0001\u0000\u0000\u0000y~\u0003\b\u0004\u0000z{\u00053\u0000\u0000"+
		"{}\u0003\b\u0004\u0000|z\u0001\u0000\u0000\u0000}\u0080\u0001\u0000\u0000"+
		"\u0000~|\u0001\u0000\u0000\u0000~\u007f\u0001\u0000\u0000\u0000\u007f"+
		"\u0081\u0001\u0000\u0000\u0000\u0080~\u0001\u0000\u0000\u0000\u0081\u0082"+
		"\u0005.\u0000\u0000\u0082\u0083\u0003\u001e\u000f\u0000\u0083\u0007\u0001"+
		"\u0000\u0000\u0000\u0084\u0089\u0005@\u0000\u0000\u0085\u0086\u0005=\u0000"+
		"\u0000\u0086\u0088\u0005@\u0000\u0000\u0087\u0085\u0001\u0000\u0000\u0000"+
		"\u0088\u008b\u0001\u0000\u0000\u0000\u0089\u0087\u0001\u0000\u0000\u0000"+
		"\u0089\u008a\u0001\u0000\u0000\u0000\u008a\u0092\u0001\u0000\u0000\u0000"+
		"\u008b\u0089\u0001\u0000\u0000\u0000\u008c\u008d\u00051\u0000\u0000\u008d"+
		"\u008e\u0003\n\u0005\u0000\u008e\u008f\u00052\u0000\u0000\u008f\u0091"+
		"\u0001\u0000\u0000\u0000\u0090\u008c\u0001\u0000\u0000\u0000\u0091\u0094"+
		"\u0001\u0000\u0000\u0000\u0092\u0090\u0001\u0000\u0000\u0000\u0092\u0093"+
		"\u0001\u0000\u0000\u0000\u0093\t\u0001\u0000\u0000\u0000\u0094\u0092\u0001"+
		"\u0000\u0000\u0000\u0095\u009b\u0003\f\u0006\u0000\u0096\u0097\u0003\u0010"+
		"\b\u0000\u0097\u0098\u0003\f\u0006\u0000\u0098\u009a\u0001\u0000\u0000"+
		"\u0000\u0099\u0096\u0001\u0000\u0000\u0000\u009a\u009d\u0001\u0000\u0000"+
		"\u0000\u009b\u0099\u0001\u0000\u0000\u0000\u009b\u009c\u0001\u0000\u0000"+
		"\u0000\u009c\u000b\u0001\u0000\u0000\u0000\u009d\u009b\u0001\u0000\u0000"+
		"\u0000\u009e\u00a4\u0003\u000e\u0007\u0000\u009f\u00a0\u0003\u0012\t\u0000"+
		"\u00a0\u00a1\u0003\u000e\u0007\u0000\u00a1\u00a3\u0001\u0000\u0000\u0000"+
		"\u00a2\u009f\u0001\u0000\u0000\u0000\u00a3\u00a6\u0001\u0000\u0000\u0000"+
		"\u00a4\u00a2\u0001\u0000\u0000\u0000\u00a4\u00a5\u0001\u0000\u0000\u0000"+
		"\u00a5\r\u0001\u0000\u0000\u0000\u00a6\u00a4\u0001\u0000\u0000\u0000\u00a7"+
		"\u00ad\u0003\u0016\u000b\u0000\u00a8\u00a9\u0003\u0014\n\u0000\u00a9\u00aa"+
		"\u0003\u0016\u000b\u0000\u00aa\u00ac\u0001\u0000\u0000\u0000\u00ab\u00a8"+
		"\u0001\u0000\u0000\u0000\u00ac\u00af\u0001\u0000\u0000\u0000\u00ad\u00ab"+
		"\u0001\u0000\u0000\u0000\u00ad\u00ae\u0001\u0000\u0000\u0000\u00ae\u000f"+
		"\u0001\u0000\u0000\u0000\u00af\u00ad\u0001\u0000\u0000\u0000\u00b0\u00b1"+
		"\u0007\u0000\u0000\u0000\u00b1\u0011\u0001\u0000\u0000\u0000\u00b2\u00b3"+
		"\u0007\u0001\u0000\u0000\u00b3\u0013\u0001\u0000\u0000\u0000\u00b4\u00b5"+
		"\u00056\u0000\u0000\u00b5\u0015\u0001\u0000\u0000\u0000\u00b6\u00b8\u0003"+
		"\u001a\r\u0000\u00b7\u00b6\u0001\u0000\u0000\u0000\u00b7\u00b8\u0001\u0000"+
		"\u0000\u0000\u00b8\u00b9\u0001\u0000\u0000\u0000\u00b9\u00bc\u0003\u0018"+
		"\f\u0000\u00ba\u00bc\u0003\u001c\u000e\u0000\u00bb\u00b7\u0001\u0000\u0000"+
		"\u0000\u00bb\u00ba\u0001\u0000\u0000\u0000\u00bc\u0017\u0001\u0000\u0000"+
		"\u0000\u00bd\u00bf\u0005;\u0000\u0000\u00be\u00bd\u0001\u0000\u0000\u0000"+
		"\u00be\u00bf\u0001\u0000\u0000\u0000\u00bf\u00c0\u0001\u0000\u0000\u0000"+
		"\u00c0\u00d4\u0003\b\u0004\u0000\u00c1\u00c2\u0005@\u0000\u0000\u00c2"+
		"\u00c3\u0005/\u0000\u0000\u00c3\u00c8\u0003R)\u0000\u00c4\u00c5\u0005"+
		"3\u0000\u0000\u00c5\u00c7\u0003R)\u0000\u00c6\u00c4\u0001\u0000\u0000"+
		"\u0000\u00c7\u00ca\u0001\u0000\u0000\u0000\u00c8\u00c6\u0001\u0000\u0000"+
		"\u0000\u00c8\u00c9\u0001\u0000\u0000\u0000\u00c9\u00cb\u0001\u0000\u0000"+
		"\u0000\u00ca\u00c8\u0001\u0000\u0000\u0000\u00cb\u00cc\u00050\u0000\u0000"+
		"\u00cc\u00d4\u0001\u0000\u0000\u0000\u00cd\u00d4\u0005>\u0000\u0000\u00ce"+
		"\u00d4\u0005?\u0000\u0000\u00cf\u00d0\u0005/\u0000\u0000\u00d0\u00d1\u0003"+
		"R)\u0000\u00d1\u00d2\u00050\u0000\u0000\u00d2\u00d4\u0001\u0000\u0000"+
		"\u0000\u00d3\u00be\u0001\u0000\u0000\u0000\u00d3\u00c1\u0001\u0000\u0000"+
		"\u0000\u00d3\u00cd\u0001\u0000\u0000\u0000\u00d3\u00ce\u0001\u0000\u0000"+
		"\u0000\u00d3\u00cf\u0001\u0000\u0000\u0000\u00d4\u0019\u0001\u0000\u0000"+
		"\u0000\u00d5\u00d6\u00058\u0000\u0000\u00d6\u001b\u0001\u0000\u0000\u0000"+
		"\u00d7\u00d8\u0005<\u0000\u0000\u00d8\u00db\u0003\b\u0004\u0000\u00d9"+
		"\u00db\u0005A\u0000\u0000\u00da\u00d7\u0001\u0000\u0000\u0000\u00da\u00d9"+
		"\u0001\u0000\u0000\u0000\u00db\u001d\u0001\u0000\u0000\u0000\u00dc\u00df"+
		"\u0003&\u0013\u0000\u00dd\u00df\u0003\"\u0011\u0000\u00de\u00dc\u0001"+
		"\u0000\u0000\u0000\u00de\u00dd\u0001\u0000\u0000\u0000\u00df\u001f\u0001"+
		"\u0000\u0000\u0000\u00e0\u00e1\u0007\u0002\u0000\u0000\u00e1!\u0001\u0000"+
		"\u0000\u0000\u00e2\u00e4\u0005;\u0000\u0000\u00e3\u00e2\u0001\u0000\u0000"+
		"\u0000\u00e3\u00e4\u0001\u0000\u0000\u0000\u00e4\u00e7\u0001\u0000\u0000"+
		"\u0000\u00e5\u00e8\u0003 \u0010\u0000\u00e6\u00e8\u0005@\u0000\u0000\u00e7"+
		"\u00e5\u0001\u0000\u0000\u0000\u00e7\u00e6\u0001\u0000\u0000\u0000\u00e8"+
		"#\u0001\u0000\u0000\u0000\u00e9\u00ea\u0007\u0003\u0000\u0000\u00ea%\u0001"+
		"\u0000\u0000\u0000\u00eb\u00ef\u0005\u001d\u0000\u0000\u00ec\u00ee\u0003"+
		"\u0006\u0003\u0000\u00ed\u00ec\u0001\u0000\u0000\u0000\u00ee\u00f1\u0001"+
		"\u0000\u0000\u0000\u00ef\u00ed\u0001\u0000\u0000\u0000\u00ef\u00f0\u0001"+
		"\u0000\u0000\u0000\u00f0\u00f2\u0001\u0000\u0000\u0000\u00f1\u00ef\u0001"+
		"\u0000\u0000\u0000\u00f2\u00f3\u0005\u001e\u0000\u0000\u00f3\'\u0001\u0000"+
		"\u0000\u0000\u00f4\u00f6\u0005\u0004\u0000\u0000\u00f5\u00f4\u0001\u0000"+
		"\u0000\u0000\u00f5\u00f6\u0001\u0000\u0000\u0000\u00f6\u00f7\u0001\u0000"+
		"\u0000\u0000\u00f7\u00fc\u0003\b\u0004\u0000\u00f8\u00f9\u00053\u0000"+
		"\u0000\u00f9\u00fb\u0003\b\u0004\u0000\u00fa\u00f8\u0001\u0000\u0000\u0000"+
		"\u00fb\u00fe\u0001\u0000\u0000\u0000\u00fc\u00fa\u0001\u0000\u0000\u0000"+
		"\u00fc\u00fd\u0001\u0000\u0000\u0000\u00fd\u00ff\u0001\u0000\u0000\u0000"+
		"\u00fe\u00fc\u0001\u0000\u0000\u0000\u00ff\u0100\u0005.\u0000\u0000\u0100"+
		"\u0101\u0003\"\u0011\u0000\u0101)\u0001\u0000\u0000\u0000\u0102\u0107"+
		"\u0003(\u0014\u0000\u0103\u0104\u00053\u0000\u0000\u0104\u0106\u0003("+
		"\u0014\u0000\u0105\u0103\u0001\u0000\u0000\u0000\u0106\u0109\u0001\u0000"+
		"\u0000\u0000\u0107\u0105\u0001\u0000\u0000\u0000\u0107\u0108\u0001\u0000"+
		"\u0000\u0000\u0108+\u0001\u0000\u0000\u0000\u0109\u0107\u0001\u0000\u0000"+
		"\u0000\u010a\u010b\u0005\u001f\u0000\u0000\u010b\u010c\u0005@\u0000\u0000"+
		"\u010c\u010e\u0005/\u0000\u0000\u010d\u010f\u0003*\u0015\u0000\u010e\u010d"+
		"\u0001\u0000\u0000\u0000\u010e\u010f\u0001\u0000\u0000\u0000\u010f\u0110"+
		"\u0001\u0000\u0000\u0000\u0110\u0114\u00050\u0000\u0000\u0111\u0113\u0003"+
		"\u0004\u0002\u0000\u0112\u0111\u0001\u0000\u0000\u0000\u0113\u0116\u0001"+
		"\u0000\u0000\u0000\u0114\u0112\u0001\u0000\u0000\u0000\u0114\u0115\u0001"+
		"\u0000\u0000\u0000\u0115\u011a\u0001\u0000\u0000\u0000\u0116\u0114\u0001"+
		"\u0000\u0000\u0000\u0117\u0119\u00030\u0018\u0000\u0118\u0117\u0001\u0000"+
		"\u0000\u0000\u0119\u011c\u0001\u0000\u0000\u0000\u011a\u0118\u0001\u0000"+
		"\u0000\u0000\u011a\u011b\u0001\u0000\u0000\u0000\u011b\u011d\u0001\u0000"+
		"\u0000\u0000\u011c\u011a\u0001\u0000\u0000\u0000\u011d\u0136\u0005 \u0000"+
		"\u0000\u011e\u011f\u0005!\u0000\u0000\u011f\u0120\u0005@\u0000\u0000\u0120"+
		"\u0122\u0005/\u0000\u0000\u0121\u0123\u0003*\u0015\u0000\u0122\u0121\u0001"+
		"\u0000\u0000\u0000\u0122\u0123\u0001\u0000\u0000\u0000\u0123\u0124\u0001"+
		"\u0000\u0000\u0000\u0124\u0125\u00050\u0000\u0000\u0125\u0126\u0005.\u0000"+
		"\u0000\u0126\u012a\u0003\"\u0011\u0000\u0127\u0129\u0003\u0004\u0002\u0000"+
		"\u0128\u0127\u0001\u0000\u0000\u0000\u0129\u012c\u0001\u0000\u0000\u0000"+
		"\u012a\u0128\u0001\u0000\u0000\u0000\u012a\u012b\u0001\u0000\u0000\u0000"+
		"\u012b\u0130\u0001\u0000\u0000\u0000\u012c\u012a\u0001\u0000\u0000\u0000"+
		"\u012d\u012f\u00030\u0018\u0000\u012e\u012d\u0001\u0000\u0000\u0000\u012f"+
		"\u0132\u0001\u0000\u0000\u0000\u0130\u012e\u0001\u0000\u0000\u0000\u0130"+
		"\u0131\u0001\u0000\u0000\u0000\u0131\u0133\u0001\u0000\u0000\u0000\u0132"+
		"\u0130\u0001\u0000\u0000\u0000\u0133\u0134\u0005\"\u0000\u0000\u0134\u0136"+
		"\u0001\u0000\u0000\u0000\u0135\u010a\u0001\u0000\u0000\u0000\u0135\u011e"+
		"\u0001\u0000\u0000\u0000\u0136-\u0001\u0000\u0000\u0000\u0137\u0139\u0003"+
		"\u0004\u0002\u0000\u0138\u0137\u0001\u0000\u0000\u0000\u0139\u013c\u0001"+
		"\u0000\u0000\u0000\u013a\u0138\u0001\u0000\u0000\u0000\u013a\u013b\u0001"+
		"\u0000\u0000\u0000\u013b\u0140\u0001\u0000\u0000\u0000\u013c\u013a\u0001"+
		"\u0000\u0000\u0000\u013d\u013f\u00030\u0018\u0000\u013e\u013d\u0001\u0000"+
		"\u0000\u0000\u013f\u0142\u0001\u0000\u0000\u0000\u0140\u013e\u0001\u0000"+
		"\u0000\u0000\u0140\u0141\u0001\u0000\u0000\u0000\u0141/\u0001\u0000\u0000"+
		"\u0000\u0142\u0140\u0001\u0000\u0000\u0000\u0143\u014e\u00032\u0019\u0000"+
		"\u0144\u014e\u00034\u001a\u0000\u0145\u014e\u00036\u001b\u0000\u0146\u014e"+
		"\u00038\u001c\u0000\u0147\u014e\u0003:\u001d\u0000\u0148\u014e\u0003<"+
		"\u001e\u0000\u0149\u014e\u0003>\u001f\u0000\u014a\u014e\u0003@ \u0000"+
		"\u014b\u014e\u0003B!\u0000\u014c\u014e\u0003D\"\u0000\u014d\u0143\u0001"+
		"\u0000\u0000\u0000\u014d\u0144\u0001\u0000\u0000\u0000\u014d\u0145\u0001"+
		"\u0000\u0000\u0000\u014d\u0146\u0001\u0000\u0000\u0000\u014d\u0147\u0001"+
		"\u0000\u0000\u0000\u014d\u0148\u0001\u0000\u0000\u0000\u014d\u0149\u0001"+
		"\u0000\u0000\u0000\u014d\u014a\u0001\u0000\u0000\u0000\u014d\u014b\u0001"+
		"\u0000\u0000\u0000\u014d\u014c\u0001\u0000\u0000\u0000\u014e1\u0001\u0000"+
		"\u0000\u0000\u014f\u0150\u0005$\u0000\u0000\u0150\u0152\u0005/\u0000\u0000"+
		"\u0151\u0153\u0005;\u0000\u0000\u0152\u0151\u0001\u0000\u0000\u0000\u0152"+
		"\u0153\u0001\u0000\u0000\u0000\u0153\u0154\u0001\u0000\u0000\u0000\u0154"+
		"\u015c\u0003\b\u0004\u0000\u0155\u0157\u00053\u0000\u0000\u0156\u0158"+
		"\u0005;\u0000\u0000\u0157\u0156\u0001\u0000\u0000\u0000\u0157\u0158\u0001"+
		"\u0000\u0000\u0000\u0158\u0159\u0001\u0000\u0000\u0000\u0159\u015b\u0003"+
		"\b\u0004\u0000\u015a\u0155\u0001\u0000\u0000\u0000\u015b\u015e\u0001\u0000"+
		"\u0000\u0000\u015c\u015a\u0001\u0000\u0000\u0000\u015c\u015d\u0001\u0000"+
		"\u0000\u0000\u015d\u015f\u0001\u0000\u0000\u0000\u015e\u015c\u0001\u0000"+
		"\u0000\u0000\u015f\u0160\u00050\u0000\u0000\u01603\u0001\u0000\u0000\u0000"+
		"\u0161\u0162\u0005%\u0000\u0000\u0162\u0163\u0005/\u0000\u0000\u0163\u0168"+
		"\u0003R)\u0000\u0164\u0165\u00053\u0000\u0000\u0165\u0167\u0003R)\u0000"+
		"\u0166\u0164\u0001\u0000\u0000\u0000\u0167\u016a\u0001\u0000\u0000\u0000"+
		"\u0168\u0166\u0001\u0000\u0000\u0000\u0168\u0169\u0001\u0000\u0000\u0000"+
		"\u0169\u016b\u0001\u0000\u0000\u0000\u016a\u0168\u0001\u0000\u0000\u0000"+
		"\u016b\u016c\u00050\u0000\u0000\u016c5\u0001\u0000\u0000\u0000\u016d\u016e"+
		"\u0005\u000f\u0000\u0000\u016e\u016f\u0003R)\u0000\u016f\u0173\u0005\u0011"+
		"\u0000\u0000\u0170\u0172\u00030\u0018\u0000\u0171\u0170\u0001\u0000\u0000"+
		"\u0000\u0172\u0175\u0001\u0000\u0000\u0000\u0173\u0171\u0001\u0000\u0000"+
		"\u0000\u0173\u0174\u0001\u0000\u0000\u0000\u0174\u017d\u0001\u0000\u0000"+
		"\u0000\u0175\u0173\u0001\u0000\u0000\u0000\u0176\u017a\u0005\u0012\u0000"+
		"\u0000\u0177\u0179\u00030\u0018\u0000\u0178\u0177\u0001\u0000\u0000\u0000"+
		"\u0179\u017c\u0001\u0000\u0000\u0000\u017a\u0178\u0001\u0000\u0000\u0000"+
		"\u017a\u017b\u0001\u0000\u0000\u0000\u017b\u017e\u0001\u0000\u0000\u0000"+
		"\u017c\u017a\u0001\u0000\u0000\u0000\u017d\u0176\u0001\u0000\u0000\u0000"+
		"\u017d\u017e\u0001\u0000\u0000\u0000\u017e\u017f\u0001\u0000\u0000\u0000"+
		"\u017f\u0180\u0005\u0010\u0000\u0000\u01807\u0001\u0000\u0000\u0000\u0181"+
		"\u0182\u0005\u0013\u0000\u0000\u0182\u0183\u0003\n\u0005\u0000\u0183\u0184"+
		"\u0005\u0014\u0000\u0000\u0184\u018c\u0003F#\u0000\u0185\u0189\u0005\u0012"+
		"\u0000\u0000\u0186\u0188\u00030\u0018\u0000\u0187\u0186\u0001\u0000\u0000"+
		"\u0000\u0188\u018b\u0001\u0000\u0000\u0000\u0189\u0187\u0001\u0000\u0000"+
		"\u0000\u0189\u018a\u0001\u0000\u0000\u0000\u018a\u018d\u0001\u0000\u0000"+
		"\u0000\u018b\u0189\u0001\u0000\u0000\u0000\u018c\u0185\u0001\u0000\u0000"+
		"\u0000\u018c\u018d\u0001\u0000\u0000\u0000\u018d\u018e\u0001\u0000\u0000"+
		"\u0000\u018e\u018f\u0005\u0015\u0000\u0000\u018f9\u0001\u0000\u0000\u0000"+
		"\u0190\u0191\u0005\u0016\u0000\u0000\u0191\u0192\u0005@\u0000\u0000\u0192"+
		"\u0193\u0005:\u0000\u0000\u0193\u0194\u0003\n\u0005\u0000\u0194\u0195"+
		"\u0005\u0018\u0000\u0000\u0195\u0196\u0003\n\u0005\u0000\u0196\u019a\u0005"+
		"\u0019\u0000\u0000\u0197\u0199\u00030\u0018\u0000\u0198\u0197\u0001\u0000"+
		"\u0000\u0000\u0199\u019c\u0001\u0000\u0000\u0000\u019a\u0198\u0001\u0000"+
		"\u0000\u0000\u019a\u019b\u0001\u0000\u0000\u0000\u019b\u019d\u0001\u0000"+
		"\u0000\u0000\u019c\u019a\u0001\u0000\u0000\u0000\u019d\u019e\u0005\u0017"+
		"\u0000\u0000\u019e;\u0001\u0000\u0000\u0000\u019f\u01a0\u0005\u001a\u0000"+
		"\u0000\u01a0\u01a1\u0003R)\u0000\u01a1\u01a5\u0005\u0019\u0000\u0000\u01a2"+
		"\u01a4\u00030\u0018\u0000\u01a3\u01a2\u0001\u0000\u0000\u0000\u01a4\u01a7"+
		"\u0001\u0000\u0000\u0000\u01a5\u01a3\u0001\u0000\u0000\u0000\u01a5\u01a6"+
		"\u0001\u0000\u0000\u0000\u01a6\u01a8\u0001\u0000\u0000\u0000\u01a7\u01a5"+
		"\u0001\u0000\u0000\u0000\u01a8\u01a9\u0005\u001b\u0000\u0000\u01a9=\u0001"+
		"\u0000\u0000\u0000\u01aa\u01ae\u0005\u0019\u0000\u0000\u01ab\u01ad\u0003"+
		"0\u0018\u0000\u01ac\u01ab\u0001\u0000\u0000\u0000\u01ad\u01b0\u0001\u0000"+
		"\u0000\u0000\u01ae\u01ac\u0001\u0000\u0000\u0000\u01ae\u01af\u0001\u0000"+
		"\u0000\u0000\u01af\u01b1\u0001\u0000\u0000\u0000\u01b0\u01ae\u0001\u0000"+
		"\u0000\u0000\u01b1\u01b2\u0005\u0018\u0000\u0000\u01b2\u01b3\u0003R)\u0000"+
		"\u01b3?\u0001\u0000\u0000\u0000\u01b4\u01b6\u0005;\u0000\u0000\u01b5\u01b4"+
		"\u0001\u0000\u0000\u0000\u01b5\u01b6\u0001\u0000\u0000\u0000\u01b6\u01b7"+
		"\u0001\u0000\u0000\u0000\u01b7\u01b8\u0003\b\u0004\u0000\u01b8\u01b9\u0005"+
		":\u0000\u0000\u01b9\u01ba\u0003R)\u0000\u01baA\u0001\u0000\u0000\u0000"+
		"\u01bb\u01bc\u0005@\u0000\u0000\u01bc\u01bd\u0005/\u0000\u0000\u01bd\u01c2"+
		"\u0003R)\u0000\u01be\u01bf\u00053\u0000\u0000\u01bf\u01c1\u0003R)\u0000"+
		"\u01c0\u01be\u0001\u0000\u0000\u0000\u01c1\u01c4\u0001\u0000\u0000\u0000"+
		"\u01c2\u01c0\u0001\u0000\u0000\u0000\u01c2\u01c3\u0001\u0000\u0000\u0000"+
		"\u01c3\u01c5\u0001\u0000\u0000\u0000\u01c4\u01c2\u0001\u0000\u0000\u0000"+
		"\u01c5\u01c6\u00050\u0000\u0000\u01c6C\u0001\u0000\u0000\u0000\u01c7\u01c8"+
		"\u0005#\u0000\u0000\u01c8\u01c9\u0003R)\u0000\u01c9E\u0001\u0000\u0000"+
		"\u0000\u01ca\u01cc\u0003H$\u0000\u01cb\u01ca\u0001\u0000\u0000\u0000\u01cc"+
		"\u01cf\u0001\u0000\u0000\u0000\u01cd\u01cb\u0001\u0000\u0000\u0000\u01cd"+
		"\u01ce\u0001\u0000\u0000\u0000\u01ceG\u0001\u0000\u0000\u0000\u01cf\u01cd"+
		"\u0001\u0000\u0000\u0000\u01d0\u01d1\u0003J%\u0000\u01d1\u01d5\u0005."+
		"\u0000\u0000\u01d2\u01d4\u00030\u0018\u0000\u01d3\u01d2\u0001\u0000\u0000"+
		"\u0000\u01d4\u01d7\u0001\u0000\u0000\u0000\u01d5\u01d3\u0001\u0000\u0000"+
		"\u0000\u01d5\u01d6\u0001\u0000\u0000\u0000\u01d6I\u0001\u0000\u0000\u0000"+
		"\u01d7\u01d5\u0001\u0000\u0000\u0000\u01d8\u01dd\u0003L&\u0000\u01d9\u01da"+
		"\u00053\u0000\u0000\u01da\u01dc\u0003L&\u0000\u01db\u01d9\u0001\u0000"+
		"\u0000\u0000\u01dc\u01df\u0001\u0000\u0000\u0000\u01dd\u01db\u0001\u0000"+
		"\u0000\u0000\u01dd\u01de\u0001\u0000\u0000\u0000\u01deK\u0001\u0000\u0000"+
		"\u0000\u01df\u01dd\u0001\u0000\u0000\u0000\u01e0\u01e2\u0003\u001a\r\u0000"+
		"\u01e1\u01e0\u0001\u0000\u0000\u0000\u01e1\u01e2\u0001\u0000\u0000\u0000"+
		"\u01e2\u01e3\u0001\u0000\u0000\u0000\u01e3\u01e9\u0005>\u0000\u0000\u01e4"+
		"\u01e6\u0005\'\u0000\u0000\u01e5\u01e7\u0003\u001a\r\u0000\u01e6\u01e5"+
		"\u0001\u0000\u0000\u0000\u01e6\u01e7\u0001\u0000\u0000\u0000\u01e7\u01e8"+
		"\u0001\u0000\u0000\u0000\u01e8\u01ea\u0005>\u0000\u0000\u01e9\u01e4\u0001"+
		"\u0000\u0000\u0000\u01e9\u01ea\u0001\u0000\u0000\u0000\u01eaM\u0001\u0000"+
		"\u0000\u0000\u01eb\u01ef\u0003\n\u0005\u0000\u01ec\u01ed\u0003P(\u0000"+
		"\u01ed\u01ee\u0003\n\u0005\u0000\u01ee\u01f0\u0001\u0000\u0000\u0000\u01ef"+
		"\u01ec\u0001\u0000\u0000\u0000\u01ef\u01f0\u0001\u0000\u0000\u0000\u01f0"+
		"O\u0001\u0000\u0000\u0000\u01f1\u01f2\u0007\u0004\u0000\u0000\u01f2Q\u0001"+
		"\u0000\u0000\u0000\u01f3\u01f9\u0003T*\u0000\u01f4\u01f5\u0003Z-\u0000"+
		"\u01f5\u01f6\u0003T*\u0000\u01f6\u01f8\u0001\u0000\u0000\u0000\u01f7\u01f4"+
		"\u0001\u0000\u0000\u0000\u01f8\u01fb\u0001\u0000\u0000\u0000\u01f9\u01f7"+
		"\u0001\u0000\u0000\u0000\u01f9\u01fa\u0001\u0000\u0000\u0000\u01faS\u0001"+
		"\u0000\u0000\u0000\u01fb\u01f9\u0001\u0000\u0000\u0000\u01fc\u0202\u0003"+
		"V+\u0000\u01fd\u01fe\u0003\\.\u0000\u01fe\u01ff\u0003V+\u0000\u01ff\u0201"+
		"\u0001\u0000\u0000\u0000\u0200\u01fd\u0001\u0000\u0000\u0000\u0201\u0204"+
		"\u0001\u0000\u0000\u0000\u0202\u0200\u0001\u0000\u0000\u0000\u0202\u0203"+
		"\u0001\u0000\u0000\u0000\u0203U\u0001\u0000\u0000\u0000\u0204\u0202\u0001"+
		"\u0000\u0000\u0000\u0205\u0207\u0005\u000e\u0000\u0000\u0206\u0205\u0001"+
		"\u0000\u0000\u0000\u0206\u0207\u0001\u0000\u0000\u0000\u0207\u0208\u0001"+
		"\u0000\u0000\u0000\u0208\u0209\u0003X,\u0000\u0209W\u0001\u0000\u0000"+
		"\u0000\u020a\u020d\u0007\u0005\u0000\u0000\u020b\u020d\u0003N\'\u0000"+
		"\u020c\u020a\u0001\u0000\u0000\u0000\u020c\u020b\u0001\u0000\u0000\u0000"+
		"\u020dY\u0001\u0000\u0000\u0000\u020e\u020f\u0005\r\u0000\u0000\u020f"+
		"[\u0001\u0000\u0000\u0000\u0210\u0211\u0005\f\u0000\u0000\u0211]\u0001"+
		"\u0000\u0000\u00009egw~\u0089\u0092\u009b\u00a4\u00ad\u00b7\u00bb\u00be"+
		"\u00c8\u00d3\u00da\u00de\u00e3\u00e7\u00ef\u00f5\u00fc\u0107\u010e\u0114"+
		"\u011a\u0122\u012a\u0130\u0135\u013a\u0140\u014d\u0152\u0157\u015c\u0168"+
		"\u0173\u017a\u017d\u0189\u018c\u019a\u01a5\u01ae\u01b5\u01c2\u01cd\u01d5"+
		"\u01dd\u01e1\u01e6\u01e9\u01ef\u01f9\u0202\u0206\u020c";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}