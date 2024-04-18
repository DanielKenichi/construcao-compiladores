// Generated from /home/lucky/UFSCAR/Compiladores/construcao-compiladores/T1/antlr4/br/ufscar/dc/compiladores/t1/lexico/T1AlgumaLexer.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class T1AlgumaLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		COMENTARIO=1, ALGORITMO=2, DECLARE=3, VAR=4, CONSTANTE=5, LITERAL=6, INTEIRO=7, 
		REAL=8, LOGICO=9, VERDADEIRO=10, FALSO=11, E=12, OR=13, NAO=14, SE=15, 
		FIM_SE=16, ENTAO=17, SENAO=18, CASO=19, SEJA=20, FIM_CASO=21, PARA=22, 
		FIM_PARA=23, ATE=24, FACA=25, ENQUANTO=26, FIM_ENQUANTO=27, TIPO=28, REGISTRO=29, 
		FIM_REGISTRO=30, PROCEDIMENTO=31, FIM_PROCEDIMENTO=32, FUNCAO=33, FIM_FUNCAO=34, 
		RETORNE=35, LEIA=36, ESCREVA=37, FIM_ALGORITMO=38, INTERVALO=39, MENOR=40, 
		MENORIGUAL=41, MAIOR=42, MAIORIGUAL=43, IGUAL=44, DIFERENTE=45, DOISPONTOS=46, 
		ABREPAR=47, FECHAPAR=48, ABRECHAVE=49, FECHACHAVE=50, VIRGULA=51, ASPAS=52, 
		DIVISAO=53, RESTO=54, SOMA=55, SUBTRACAO=56, MULTIPLICACAO=57, ATRIBUICAO=58, 
		PONTEIRO=59, ENDERECO=60, PONTO=61, NUM_INT=62, NUM_REAL=63, IDENT=64, 
		CADEIA=65, ERRO_COMENTARIO_ABERTO=66, ERRO_CADEIA_ABERTA=67, WS=68, ERRO=69;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"COMENTARIO", "ALGORITMO", "DECLARE", "VAR", "CONSTANTE", "LITERAL", 
			"INTEIRO", "REAL", "LOGICO", "VERDADEIRO", "FALSO", "E", "OR", "NAO", 
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
			"INTEIRO", "REAL", "LOGICO", "VERDADEIRO", "FALSO", "E", "OR", "NAO", 
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


	public T1AlgumaLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "T1AlgumaLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	@Override
	public void action(RuleContext _localctx, int ruleIndex, int actionIndex) {
		switch (ruleIndex) {
		case 0:
			COMENTARIO_action((RuleContext)_localctx, actionIndex);
			break;
		case 67:
			WS_action((RuleContext)_localctx, actionIndex);
			break;
		}
	}
	private void COMENTARIO_action(RuleContext _localctx, int actionIndex) {
		switch (actionIndex) {
		case 0:
			skip();
			break;
		}
	}
	private void WS_action(RuleContext _localctx, int actionIndex) {
		switch (actionIndex) {
		case 1:
			skip();
			break;
		}
	}

	public static final String _serializedATN =
		"\u0004\u0000E\u021b\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002"+
		"\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002"+
		"\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002"+
		"\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002"+
		"\u0018\u0007\u0018\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002"+
		"\u001b\u0007\u001b\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002"+
		"\u001e\u0007\u001e\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007"+
		"!\u0002\"\u0007\"\u0002#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007"+
		"&\u0002\'\u0007\'\u0002(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007"+
		"+\u0002,\u0007,\u0002-\u0007-\u0002.\u0007.\u0002/\u0007/\u00020\u0007"+
		"0\u00021\u00071\u00022\u00072\u00023\u00073\u00024\u00074\u00025\u0007"+
		"5\u00026\u00076\u00027\u00077\u00028\u00078\u00029\u00079\u0002:\u0007"+
		":\u0002;\u0007;\u0002<\u0007<\u0002=\u0007=\u0002>\u0007>\u0002?\u0007"+
		"?\u0002@\u0007@\u0002A\u0007A\u0002B\u0007B\u0002C\u0007C\u0002D\u0007"+
		"D\u0001\u0000\u0001\u0000\u0005\u0000\u008e\b\u0000\n\u0000\f\u0000\u0091"+
		"\t\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0001"+
		"\u001b\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001"+
		"\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001"+
		"\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001!\u0001!"+
		"\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001"+
		"%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001&\u0001&\u0001&\u0001\'\u0001"+
		"\'\u0001(\u0001(\u0001(\u0001)\u0001)\u0001*\u0001*\u0001*\u0001+\u0001"+
		"+\u0001,\u0001,\u0001,\u0001-\u0001-\u0001.\u0001.\u0001/\u0001/\u0001"+
		"0\u00010\u00011\u00011\u00012\u00012\u00013\u00013\u00014\u00014\u0001"+
		"5\u00015\u00016\u00016\u00017\u00017\u00018\u00018\u00019\u00019\u0001"+
		"9\u0001:\u0001:\u0001;\u0001;\u0001<\u0001<\u0001=\u0004=\u01e4\b=\u000b"+
		"=\f=\u01e5\u0001>\u0004>\u01e9\b>\u000b>\f>\u01ea\u0001>\u0001>\u0004"+
		">\u01ef\b>\u000b>\f>\u01f0\u0003>\u01f3\b>\u0001?\u0001?\u0005?\u01f7"+
		"\b?\n?\f?\u01fa\t?\u0001@\u0001@\u0005@\u01fe\b@\n@\f@\u0201\t@\u0001"+
		"@\u0001@\u0001A\u0001A\u0005A\u0207\bA\nA\fA\u020a\tA\u0001A\u0001A\u0001"+
		"B\u0001B\u0005B\u0210\bB\nB\fB\u0213\tB\u0001B\u0001B\u0001C\u0001C\u0001"+
		"C\u0001D\u0001D\u0004\u008f\u01ff\u0208\u0211\u0000E\u0001\u0001\u0003"+
		"\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011"+
		"\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010"+
		"!\u0011#\u0012%\u0013\'\u0014)\u0015+\u0016-\u0017/\u00181\u00193\u001a"+
		"5\u001b7\u001c9\u001d;\u001e=\u001f? A!C\"E#G$I%K&M\'O(Q)S*U+W,Y-[.]/"+
		"_0a1c2e3g4i5k6m7o8q9s:u;w<y={>}?\u007f@\u0081A\u0083B\u0085C\u0087D\u0089"+
		"E\u0001\u0000\u0006\u0001\u0000\n\n\u0002\u0000AZaz\u0004\u000009AZ__"+
		"az\u0002\u0000\n\n}}\u0002\u0000\n\n\"\"\u0003\u0000\t\n\r\r  \u0223\u0000"+
		"\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000"+
		"\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000"+
		"\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r"+
		"\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011"+
		"\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015"+
		"\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019"+
		"\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d"+
		"\u0001\u0000\u0000\u0000\u0000\u001f\u0001\u0000\u0000\u0000\u0000!\u0001"+
		"\u0000\u0000\u0000\u0000#\u0001\u0000\u0000\u0000\u0000%\u0001\u0000\u0000"+
		"\u0000\u0000\'\u0001\u0000\u0000\u0000\u0000)\u0001\u0000\u0000\u0000"+
		"\u0000+\u0001\u0000\u0000\u0000\u0000-\u0001\u0000\u0000\u0000\u0000/"+
		"\u0001\u0000\u0000\u0000\u00001\u0001\u0000\u0000\u0000\u00003\u0001\u0000"+
		"\u0000\u0000\u00005\u0001\u0000\u0000\u0000\u00007\u0001\u0000\u0000\u0000"+
		"\u00009\u0001\u0000\u0000\u0000\u0000;\u0001\u0000\u0000\u0000\u0000="+
		"\u0001\u0000\u0000\u0000\u0000?\u0001\u0000\u0000\u0000\u0000A\u0001\u0000"+
		"\u0000\u0000\u0000C\u0001\u0000\u0000\u0000\u0000E\u0001\u0000\u0000\u0000"+
		"\u0000G\u0001\u0000\u0000\u0000\u0000I\u0001\u0000\u0000\u0000\u0000K"+
		"\u0001\u0000\u0000\u0000\u0000M\u0001\u0000\u0000\u0000\u0000O\u0001\u0000"+
		"\u0000\u0000\u0000Q\u0001\u0000\u0000\u0000\u0000S\u0001\u0000\u0000\u0000"+
		"\u0000U\u0001\u0000\u0000\u0000\u0000W\u0001\u0000\u0000\u0000\u0000Y"+
		"\u0001\u0000\u0000\u0000\u0000[\u0001\u0000\u0000\u0000\u0000]\u0001\u0000"+
		"\u0000\u0000\u0000_\u0001\u0000\u0000\u0000\u0000a\u0001\u0000\u0000\u0000"+
		"\u0000c\u0001\u0000\u0000\u0000\u0000e\u0001\u0000\u0000\u0000\u0000g"+
		"\u0001\u0000\u0000\u0000\u0000i\u0001\u0000\u0000\u0000\u0000k\u0001\u0000"+
		"\u0000\u0000\u0000m\u0001\u0000\u0000\u0000\u0000o\u0001\u0000\u0000\u0000"+
		"\u0000q\u0001\u0000\u0000\u0000\u0000s\u0001\u0000\u0000\u0000\u0000u"+
		"\u0001\u0000\u0000\u0000\u0000w\u0001\u0000\u0000\u0000\u0000y\u0001\u0000"+
		"\u0000\u0000\u0000{\u0001\u0000\u0000\u0000\u0000}\u0001\u0000\u0000\u0000"+
		"\u0000\u007f\u0001\u0000\u0000\u0000\u0000\u0081\u0001\u0000\u0000\u0000"+
		"\u0000\u0083\u0001\u0000\u0000\u0000\u0000\u0085\u0001\u0000\u0000\u0000"+
		"\u0000\u0087\u0001\u0000\u0000\u0000\u0000\u0089\u0001\u0000\u0000\u0000"+
		"\u0001\u008b\u0001\u0000\u0000\u0000\u0003\u0095\u0001\u0000\u0000\u0000"+
		"\u0005\u009f\u0001\u0000\u0000\u0000\u0007\u00a7\u0001\u0000\u0000\u0000"+
		"\t\u00ab\u0001\u0000\u0000\u0000\u000b\u00b5\u0001\u0000\u0000\u0000\r"+
		"\u00bd\u0001\u0000\u0000\u0000\u000f\u00c5\u0001\u0000\u0000\u0000\u0011"+
		"\u00ca\u0001\u0000\u0000\u0000\u0013\u00d1\u0001\u0000\u0000\u0000\u0015"+
		"\u00dc\u0001\u0000\u0000\u0000\u0017\u00e2\u0001\u0000\u0000\u0000\u0019"+
		"\u00e4\u0001\u0000\u0000\u0000\u001b\u00e7\u0001\u0000\u0000\u0000\u001d"+
		"\u00eb\u0001\u0000\u0000\u0000\u001f\u00ee\u0001\u0000\u0000\u0000!\u00f5"+
		"\u0001\u0000\u0000\u0000#\u00fb\u0001\u0000\u0000\u0000%\u0101\u0001\u0000"+
		"\u0000\u0000\'\u0106\u0001\u0000\u0000\u0000)\u010b\u0001\u0000\u0000"+
		"\u0000+\u0114\u0001\u0000\u0000\u0000-\u0119\u0001\u0000\u0000\u0000/"+
		"\u0122\u0001\u0000\u0000\u00001\u0126\u0001\u0000\u0000\u00003\u012b\u0001"+
		"\u0000\u0000\u00005\u0134\u0001\u0000\u0000\u00007\u0141\u0001\u0000\u0000"+
		"\u00009\u0146\u0001\u0000\u0000\u0000;\u014f\u0001\u0000\u0000\u0000="+
		"\u015c\u0001\u0000\u0000\u0000?\u0169\u0001\u0000\u0000\u0000A\u017a\u0001"+
		"\u0000\u0000\u0000C\u0181\u0001\u0000\u0000\u0000E\u018c\u0001\u0000\u0000"+
		"\u0000G\u0194\u0001\u0000\u0000\u0000I\u0199\u0001\u0000\u0000\u0000K"+
		"\u01a1\u0001\u0000\u0000\u0000M\u01af\u0001\u0000\u0000\u0000O\u01b2\u0001"+
		"\u0000\u0000\u0000Q\u01b4\u0001\u0000\u0000\u0000S\u01b7\u0001\u0000\u0000"+
		"\u0000U\u01b9\u0001\u0000\u0000\u0000W\u01bc\u0001\u0000\u0000\u0000Y"+
		"\u01be\u0001\u0000\u0000\u0000[\u01c1\u0001\u0000\u0000\u0000]\u01c3\u0001"+
		"\u0000\u0000\u0000_\u01c5\u0001\u0000\u0000\u0000a\u01c7\u0001\u0000\u0000"+
		"\u0000c\u01c9\u0001\u0000\u0000\u0000e\u01cb\u0001\u0000\u0000\u0000g"+
		"\u01cd\u0001\u0000\u0000\u0000i\u01cf\u0001\u0000\u0000\u0000k\u01d1\u0001"+
		"\u0000\u0000\u0000m\u01d3\u0001\u0000\u0000\u0000o\u01d5\u0001\u0000\u0000"+
		"\u0000q\u01d7\u0001\u0000\u0000\u0000s\u01d9\u0001\u0000\u0000\u0000u"+
		"\u01dc\u0001\u0000\u0000\u0000w\u01de\u0001\u0000\u0000\u0000y\u01e0\u0001"+
		"\u0000\u0000\u0000{\u01e3\u0001\u0000\u0000\u0000}\u01e8\u0001\u0000\u0000"+
		"\u0000\u007f\u01f4\u0001\u0000\u0000\u0000\u0081\u01fb\u0001\u0000\u0000"+
		"\u0000\u0083\u0204\u0001\u0000\u0000\u0000\u0085\u020d\u0001\u0000\u0000"+
		"\u0000\u0087\u0216\u0001\u0000\u0000\u0000\u0089\u0219\u0001\u0000\u0000"+
		"\u0000\u008b\u008f\u0005{\u0000\u0000\u008c\u008e\b\u0000\u0000\u0000"+
		"\u008d\u008c\u0001\u0000\u0000\u0000\u008e\u0091\u0001\u0000\u0000\u0000"+
		"\u008f\u0090\u0001\u0000\u0000\u0000\u008f\u008d\u0001\u0000\u0000\u0000"+
		"\u0090\u0092\u0001\u0000\u0000\u0000\u0091\u008f\u0001\u0000\u0000\u0000"+
		"\u0092\u0093\u0005}\u0000\u0000\u0093\u0094\u0006\u0000\u0000\u0000\u0094"+
		"\u0002\u0001\u0000\u0000\u0000\u0095\u0096\u0005a\u0000\u0000\u0096\u0097"+
		"\u0005l\u0000\u0000\u0097\u0098\u0005g\u0000\u0000\u0098\u0099\u0005o"+
		"\u0000\u0000\u0099\u009a\u0005r\u0000\u0000\u009a\u009b\u0005i\u0000\u0000"+
		"\u009b\u009c\u0005t\u0000\u0000\u009c\u009d\u0005m\u0000\u0000\u009d\u009e"+
		"\u0005o\u0000\u0000\u009e\u0004\u0001\u0000\u0000\u0000\u009f\u00a0\u0005"+
		"d\u0000\u0000\u00a0\u00a1\u0005e\u0000\u0000\u00a1\u00a2\u0005c\u0000"+
		"\u0000\u00a2\u00a3\u0005l\u0000\u0000\u00a3\u00a4\u0005a\u0000\u0000\u00a4"+
		"\u00a5\u0005r\u0000\u0000\u00a5\u00a6\u0005e\u0000\u0000\u00a6\u0006\u0001"+
		"\u0000\u0000\u0000\u00a7\u00a8\u0005v\u0000\u0000\u00a8\u00a9\u0005a\u0000"+
		"\u0000\u00a9\u00aa\u0005r\u0000\u0000\u00aa\b\u0001\u0000\u0000\u0000"+
		"\u00ab\u00ac\u0005c\u0000\u0000\u00ac\u00ad\u0005o\u0000\u0000\u00ad\u00ae"+
		"\u0005n\u0000\u0000\u00ae\u00af\u0005s\u0000\u0000\u00af\u00b0\u0005t"+
		"\u0000\u0000\u00b0\u00b1\u0005a\u0000\u0000\u00b1\u00b2\u0005n\u0000\u0000"+
		"\u00b2\u00b3\u0005t\u0000\u0000\u00b3\u00b4\u0005e\u0000\u0000\u00b4\n"+
		"\u0001\u0000\u0000\u0000\u00b5\u00b6\u0005l\u0000\u0000\u00b6\u00b7\u0005"+
		"i\u0000\u0000\u00b7\u00b8\u0005t\u0000\u0000\u00b8\u00b9\u0005e\u0000"+
		"\u0000\u00b9\u00ba\u0005r\u0000\u0000\u00ba\u00bb\u0005a\u0000\u0000\u00bb"+
		"\u00bc\u0005l\u0000\u0000\u00bc\f\u0001\u0000\u0000\u0000\u00bd\u00be"+
		"\u0005i\u0000\u0000\u00be\u00bf\u0005n\u0000\u0000\u00bf\u00c0\u0005t"+
		"\u0000\u0000\u00c0\u00c1\u0005e\u0000\u0000\u00c1\u00c2\u0005i\u0000\u0000"+
		"\u00c2\u00c3\u0005r\u0000\u0000\u00c3\u00c4\u0005o\u0000\u0000\u00c4\u000e"+
		"\u0001\u0000\u0000\u0000\u00c5\u00c6\u0005r\u0000\u0000\u00c6\u00c7\u0005"+
		"e\u0000\u0000\u00c7\u00c8\u0005a\u0000\u0000\u00c8\u00c9\u0005l\u0000"+
		"\u0000\u00c9\u0010\u0001\u0000\u0000\u0000\u00ca\u00cb\u0005l\u0000\u0000"+
		"\u00cb\u00cc\u0005o\u0000\u0000\u00cc\u00cd\u0005g\u0000\u0000\u00cd\u00ce"+
		"\u0005i\u0000\u0000\u00ce\u00cf\u0005c\u0000\u0000\u00cf\u00d0\u0005o"+
		"\u0000\u0000\u00d0\u0012\u0001\u0000\u0000\u0000\u00d1\u00d2\u0005v\u0000"+
		"\u0000\u00d2\u00d3\u0005e\u0000\u0000\u00d3\u00d4\u0005r\u0000\u0000\u00d4"+
		"\u00d5\u0005d\u0000\u0000\u00d5\u00d6\u0005a\u0000\u0000\u00d6\u00d7\u0005"+
		"d\u0000\u0000\u00d7\u00d8\u0005e\u0000\u0000\u00d8\u00d9\u0005i\u0000"+
		"\u0000\u00d9\u00da\u0005r\u0000\u0000\u00da\u00db\u0005o\u0000\u0000\u00db"+
		"\u0014\u0001\u0000\u0000\u0000\u00dc\u00dd\u0005f\u0000\u0000\u00dd\u00de"+
		"\u0005a\u0000\u0000\u00de\u00df\u0005l\u0000\u0000\u00df\u00e0\u0005s"+
		"\u0000\u0000\u00e0\u00e1\u0005o\u0000\u0000\u00e1\u0016\u0001\u0000\u0000"+
		"\u0000\u00e2\u00e3\u0005e\u0000\u0000\u00e3\u0018\u0001\u0000\u0000\u0000"+
		"\u00e4\u00e5\u0005o\u0000\u0000\u00e5\u00e6\u0005u\u0000\u0000\u00e6\u001a"+
		"\u0001\u0000\u0000\u0000\u00e7\u00e8\u0005n\u0000\u0000\u00e8\u00e9\u0005"+
		"a\u0000\u0000\u00e9\u00ea\u0005o\u0000\u0000\u00ea\u001c\u0001\u0000\u0000"+
		"\u0000\u00eb\u00ec\u0005s\u0000\u0000\u00ec\u00ed\u0005e\u0000\u0000\u00ed"+
		"\u001e\u0001\u0000\u0000\u0000\u00ee\u00ef\u0005f\u0000\u0000\u00ef\u00f0"+
		"\u0005i\u0000\u0000\u00f0\u00f1\u0005m\u0000\u0000\u00f1\u00f2\u0005_"+
		"\u0000\u0000\u00f2\u00f3\u0005s\u0000\u0000\u00f3\u00f4\u0005e\u0000\u0000"+
		"\u00f4 \u0001\u0000\u0000\u0000\u00f5\u00f6\u0005e\u0000\u0000\u00f6\u00f7"+
		"\u0005n\u0000\u0000\u00f7\u00f8\u0005t\u0000\u0000\u00f8\u00f9\u0005a"+
		"\u0000\u0000\u00f9\u00fa\u0005o\u0000\u0000\u00fa\"\u0001\u0000\u0000"+
		"\u0000\u00fb\u00fc\u0005s\u0000\u0000\u00fc\u00fd\u0005e\u0000\u0000\u00fd"+
		"\u00fe\u0005n\u0000\u0000\u00fe\u00ff\u0005a\u0000\u0000\u00ff\u0100\u0005"+
		"o\u0000\u0000\u0100$\u0001\u0000\u0000\u0000\u0101\u0102\u0005c\u0000"+
		"\u0000\u0102\u0103\u0005a\u0000\u0000\u0103\u0104\u0005s\u0000\u0000\u0104"+
		"\u0105\u0005o\u0000\u0000\u0105&\u0001\u0000\u0000\u0000\u0106\u0107\u0005"+
		"s\u0000\u0000\u0107\u0108\u0005e\u0000\u0000\u0108\u0109\u0005j\u0000"+
		"\u0000\u0109\u010a\u0005a\u0000\u0000\u010a(\u0001\u0000\u0000\u0000\u010b"+
		"\u010c\u0005f\u0000\u0000\u010c\u010d\u0005i\u0000\u0000\u010d\u010e\u0005"+
		"m\u0000\u0000\u010e\u010f\u0005_\u0000\u0000\u010f\u0110\u0005c\u0000"+
		"\u0000\u0110\u0111\u0005a\u0000\u0000\u0111\u0112\u0005s\u0000\u0000\u0112"+
		"\u0113\u0005o\u0000\u0000\u0113*\u0001\u0000\u0000\u0000\u0114\u0115\u0005"+
		"p\u0000\u0000\u0115\u0116\u0005a\u0000\u0000\u0116\u0117\u0005r\u0000"+
		"\u0000\u0117\u0118\u0005a\u0000\u0000\u0118,\u0001\u0000\u0000\u0000\u0119"+
		"\u011a\u0005f\u0000\u0000\u011a\u011b\u0005i\u0000\u0000\u011b\u011c\u0005"+
		"m\u0000\u0000\u011c\u011d\u0005_\u0000\u0000\u011d\u011e\u0005p\u0000"+
		"\u0000\u011e\u011f\u0005a\u0000\u0000\u011f\u0120\u0005r\u0000\u0000\u0120"+
		"\u0121\u0005a\u0000\u0000\u0121.\u0001\u0000\u0000\u0000\u0122\u0123\u0005"+
		"a\u0000\u0000\u0123\u0124\u0005t\u0000\u0000\u0124\u0125\u0005e\u0000"+
		"\u0000\u01250\u0001\u0000\u0000\u0000\u0126\u0127\u0005f\u0000\u0000\u0127"+
		"\u0128\u0005a\u0000\u0000\u0128\u0129\u0005c\u0000\u0000\u0129\u012a\u0005"+
		"a\u0000\u0000\u012a2\u0001\u0000\u0000\u0000\u012b\u012c\u0005e\u0000"+
		"\u0000\u012c\u012d\u0005n\u0000\u0000\u012d\u012e\u0005q\u0000\u0000\u012e"+
		"\u012f\u0005u\u0000\u0000\u012f\u0130\u0005a\u0000\u0000\u0130\u0131\u0005"+
		"n\u0000\u0000\u0131\u0132\u0005t\u0000\u0000\u0132\u0133\u0005o\u0000"+
		"\u0000\u01334\u0001\u0000\u0000\u0000\u0134\u0135\u0005f\u0000\u0000\u0135"+
		"\u0136\u0005i\u0000\u0000\u0136\u0137\u0005m\u0000\u0000\u0137\u0138\u0005"+
		"_\u0000\u0000\u0138\u0139\u0005e\u0000\u0000\u0139\u013a\u0005n\u0000"+
		"\u0000\u013a\u013b\u0005q\u0000\u0000\u013b\u013c\u0005u\u0000\u0000\u013c"+
		"\u013d\u0005a\u0000\u0000\u013d\u013e\u0005n\u0000\u0000\u013e\u013f\u0005"+
		"t\u0000\u0000\u013f\u0140\u0005o\u0000\u0000\u01406\u0001\u0000\u0000"+
		"\u0000\u0141\u0142\u0005t\u0000\u0000\u0142\u0143\u0005i\u0000\u0000\u0143"+
		"\u0144\u0005p\u0000\u0000\u0144\u0145\u0005o\u0000\u0000\u01458\u0001"+
		"\u0000\u0000\u0000\u0146\u0147\u0005r\u0000\u0000\u0147\u0148\u0005e\u0000"+
		"\u0000\u0148\u0149\u0005g\u0000\u0000\u0149\u014a\u0005i\u0000\u0000\u014a"+
		"\u014b\u0005s\u0000\u0000\u014b\u014c\u0005t\u0000\u0000\u014c\u014d\u0005"+
		"r\u0000\u0000\u014d\u014e\u0005o\u0000\u0000\u014e:\u0001\u0000\u0000"+
		"\u0000\u014f\u0150\u0005f\u0000\u0000\u0150\u0151\u0005i\u0000\u0000\u0151"+
		"\u0152\u0005m\u0000\u0000\u0152\u0153\u0005_\u0000\u0000\u0153\u0154\u0005"+
		"r\u0000\u0000\u0154\u0155\u0005e\u0000\u0000\u0155\u0156\u0005g\u0000"+
		"\u0000\u0156\u0157\u0005i\u0000\u0000\u0157\u0158\u0005s\u0000\u0000\u0158"+
		"\u0159\u0005t\u0000\u0000\u0159\u015a\u0005r\u0000\u0000\u015a\u015b\u0005"+
		"o\u0000\u0000\u015b<\u0001\u0000\u0000\u0000\u015c\u015d\u0005p\u0000"+
		"\u0000\u015d\u015e\u0005r\u0000\u0000\u015e\u015f\u0005o\u0000\u0000\u015f"+
		"\u0160\u0005c\u0000\u0000\u0160\u0161\u0005e\u0000\u0000\u0161\u0162\u0005"+
		"d\u0000\u0000\u0162\u0163\u0005i\u0000\u0000\u0163\u0164\u0005m\u0000"+
		"\u0000\u0164\u0165\u0005e\u0000\u0000\u0165\u0166\u0005n\u0000\u0000\u0166"+
		"\u0167\u0005t\u0000\u0000\u0167\u0168\u0005o\u0000\u0000\u0168>\u0001"+
		"\u0000\u0000\u0000\u0169\u016a\u0005f\u0000\u0000\u016a\u016b\u0005i\u0000"+
		"\u0000\u016b\u016c\u0005m\u0000\u0000\u016c\u016d\u0005_\u0000\u0000\u016d"+
		"\u016e\u0005p\u0000\u0000\u016e\u016f\u0005r\u0000\u0000\u016f\u0170\u0005"+
		"o\u0000\u0000\u0170\u0171\u0005c\u0000\u0000\u0171\u0172\u0005e\u0000"+
		"\u0000\u0172\u0173\u0005d\u0000\u0000\u0173\u0174\u0005i\u0000\u0000\u0174"+
		"\u0175\u0005m\u0000\u0000\u0175\u0176\u0005e\u0000\u0000\u0176\u0177\u0005"+
		"n\u0000\u0000\u0177\u0178\u0005t\u0000\u0000\u0178\u0179\u0005o\u0000"+
		"\u0000\u0179@\u0001\u0000\u0000\u0000\u017a\u017b\u0005f\u0000\u0000\u017b"+
		"\u017c\u0005u\u0000\u0000\u017c\u017d\u0005n\u0000\u0000\u017d\u017e\u0005"+
		"c\u0000\u0000\u017e\u017f\u0005a\u0000\u0000\u017f\u0180\u0005o\u0000"+
		"\u0000\u0180B\u0001\u0000\u0000\u0000\u0181\u0182\u0005f\u0000\u0000\u0182"+
		"\u0183\u0005i\u0000\u0000\u0183\u0184\u0005m\u0000\u0000\u0184\u0185\u0005"+
		"_\u0000\u0000\u0185\u0186\u0005f\u0000\u0000\u0186\u0187\u0005u\u0000"+
		"\u0000\u0187\u0188\u0005n\u0000\u0000\u0188\u0189\u0005c\u0000\u0000\u0189"+
		"\u018a\u0005a\u0000\u0000\u018a\u018b\u0005o\u0000\u0000\u018bD\u0001"+
		"\u0000\u0000\u0000\u018c\u018d\u0005r\u0000\u0000\u018d\u018e\u0005e\u0000"+
		"\u0000\u018e\u018f\u0005t\u0000\u0000\u018f\u0190\u0005o\u0000\u0000\u0190"+
		"\u0191\u0005r\u0000\u0000\u0191\u0192\u0005n\u0000\u0000\u0192\u0193\u0005"+
		"e\u0000\u0000\u0193F\u0001\u0000\u0000\u0000\u0194\u0195\u0005l\u0000"+
		"\u0000\u0195\u0196\u0005e\u0000\u0000\u0196\u0197\u0005i\u0000\u0000\u0197"+
		"\u0198\u0005a\u0000\u0000\u0198H\u0001\u0000\u0000\u0000\u0199\u019a\u0005"+
		"e\u0000\u0000\u019a\u019b\u0005s\u0000\u0000\u019b\u019c\u0005c\u0000"+
		"\u0000\u019c\u019d\u0005r\u0000\u0000\u019d\u019e\u0005e\u0000\u0000\u019e"+
		"\u019f\u0005v\u0000\u0000\u019f\u01a0\u0005a\u0000\u0000\u01a0J\u0001"+
		"\u0000\u0000\u0000\u01a1\u01a2\u0005f\u0000\u0000\u01a2\u01a3\u0005i\u0000"+
		"\u0000\u01a3\u01a4\u0005m\u0000\u0000\u01a4\u01a5\u0005_\u0000\u0000\u01a5"+
		"\u01a6\u0005a\u0000\u0000\u01a6\u01a7\u0005l\u0000\u0000\u01a7\u01a8\u0005"+
		"g\u0000\u0000\u01a8\u01a9\u0005o\u0000\u0000\u01a9\u01aa\u0005r\u0000"+
		"\u0000\u01aa\u01ab\u0005i\u0000\u0000\u01ab\u01ac\u0005t\u0000\u0000\u01ac"+
		"\u01ad\u0005m\u0000\u0000\u01ad\u01ae\u0005o\u0000\u0000\u01aeL\u0001"+
		"\u0000\u0000\u0000\u01af\u01b0\u0005.\u0000\u0000\u01b0\u01b1\u0005.\u0000"+
		"\u0000\u01b1N\u0001\u0000\u0000\u0000\u01b2\u01b3\u0005<\u0000\u0000\u01b3"+
		"P\u0001\u0000\u0000\u0000\u01b4\u01b5\u0005<\u0000\u0000\u01b5\u01b6\u0005"+
		"=\u0000\u0000\u01b6R\u0001\u0000\u0000\u0000\u01b7\u01b8\u0005>\u0000"+
		"\u0000\u01b8T\u0001\u0000\u0000\u0000\u01b9\u01ba\u0005>\u0000\u0000\u01ba"+
		"\u01bb\u0005=\u0000\u0000\u01bbV\u0001\u0000\u0000\u0000\u01bc\u01bd\u0005"+
		"=\u0000\u0000\u01bdX\u0001\u0000\u0000\u0000\u01be\u01bf\u0005<\u0000"+
		"\u0000\u01bf\u01c0\u0005>\u0000\u0000\u01c0Z\u0001\u0000\u0000\u0000\u01c1"+
		"\u01c2\u0005:\u0000\u0000\u01c2\\\u0001\u0000\u0000\u0000\u01c3\u01c4"+
		"\u0005(\u0000\u0000\u01c4^\u0001\u0000\u0000\u0000\u01c5\u01c6\u0005)"+
		"\u0000\u0000\u01c6`\u0001\u0000\u0000\u0000\u01c7\u01c8\u0005[\u0000\u0000"+
		"\u01c8b\u0001\u0000\u0000\u0000\u01c9\u01ca\u0005]\u0000\u0000\u01cad"+
		"\u0001\u0000\u0000\u0000\u01cb\u01cc\u0005,\u0000\u0000\u01ccf\u0001\u0000"+
		"\u0000\u0000\u01cd\u01ce\u0005\"\u0000\u0000\u01ceh\u0001\u0000\u0000"+
		"\u0000\u01cf\u01d0\u0005/\u0000\u0000\u01d0j\u0001\u0000\u0000\u0000\u01d1"+
		"\u01d2\u0005%\u0000\u0000\u01d2l\u0001\u0000\u0000\u0000\u01d3\u01d4\u0005"+
		"+\u0000\u0000\u01d4n\u0001\u0000\u0000\u0000\u01d5\u01d6\u0005-\u0000"+
		"\u0000\u01d6p\u0001\u0000\u0000\u0000\u01d7\u01d8\u0005*\u0000\u0000\u01d8"+
		"r\u0001\u0000\u0000\u0000\u01d9\u01da\u0005<\u0000\u0000\u01da\u01db\u0005"+
		"-\u0000\u0000\u01dbt\u0001\u0000\u0000\u0000\u01dc\u01dd\u0005^\u0000"+
		"\u0000\u01ddv\u0001\u0000\u0000\u0000\u01de\u01df\u0005&\u0000\u0000\u01df"+
		"x\u0001\u0000\u0000\u0000\u01e0\u01e1\u0005.\u0000\u0000\u01e1z\u0001"+
		"\u0000\u0000\u0000\u01e2\u01e4\u000209\u0000\u01e3\u01e2\u0001\u0000\u0000"+
		"\u0000\u01e4\u01e5\u0001\u0000\u0000\u0000\u01e5\u01e3\u0001\u0000\u0000"+
		"\u0000\u01e5\u01e6\u0001\u0000\u0000\u0000\u01e6|\u0001\u0000\u0000\u0000"+
		"\u01e7\u01e9\u000209\u0000\u01e8\u01e7\u0001\u0000\u0000\u0000\u01e9\u01ea"+
		"\u0001\u0000\u0000\u0000\u01ea\u01e8\u0001\u0000\u0000\u0000\u01ea\u01eb"+
		"\u0001\u0000\u0000\u0000\u01eb\u01f2\u0001\u0000\u0000\u0000\u01ec\u01ee"+
		"\u0005.\u0000\u0000\u01ed\u01ef\u000209\u0000\u01ee\u01ed\u0001\u0000"+
		"\u0000\u0000\u01ef\u01f0\u0001\u0000\u0000\u0000\u01f0\u01ee\u0001\u0000"+
		"\u0000\u0000\u01f0\u01f1\u0001\u0000\u0000\u0000\u01f1\u01f3\u0001\u0000"+
		"\u0000\u0000\u01f2\u01ec\u0001\u0000\u0000\u0000\u01f2\u01f3\u0001\u0000"+
		"\u0000\u0000\u01f3~\u0001\u0000\u0000\u0000\u01f4\u01f8\u0007\u0001\u0000"+
		"\u0000\u01f5\u01f7\u0007\u0002\u0000\u0000\u01f6\u01f5\u0001\u0000\u0000"+
		"\u0000\u01f7\u01fa\u0001\u0000\u0000\u0000\u01f8\u01f6\u0001\u0000\u0000"+
		"\u0000\u01f8\u01f9\u0001\u0000\u0000\u0000\u01f9\u0080\u0001\u0000\u0000"+
		"\u0000\u01fa\u01f8\u0001\u0000\u0000\u0000\u01fb\u01ff\u0005\"\u0000\u0000"+
		"\u01fc\u01fe\b\u0000\u0000\u0000\u01fd\u01fc\u0001\u0000\u0000\u0000\u01fe"+
		"\u0201\u0001\u0000\u0000\u0000\u01ff\u0200\u0001\u0000\u0000\u0000\u01ff"+
		"\u01fd\u0001\u0000\u0000\u0000\u0200\u0202\u0001\u0000\u0000\u0000\u0201"+
		"\u01ff\u0001\u0000\u0000\u0000\u0202\u0203\u0005\"\u0000\u0000\u0203\u0082"+
		"\u0001\u0000\u0000\u0000\u0204\u0208\u0005{\u0000\u0000\u0205\u0207\b"+
		"\u0003\u0000\u0000\u0206\u0205\u0001\u0000\u0000\u0000\u0207\u020a\u0001"+
		"\u0000\u0000\u0000\u0208\u0209\u0001\u0000\u0000\u0000\u0208\u0206\u0001"+
		"\u0000\u0000\u0000\u0209\u020b\u0001\u0000\u0000\u0000\u020a\u0208\u0001"+
		"\u0000\u0000\u0000\u020b\u020c\u0005\n\u0000\u0000\u020c\u0084\u0001\u0000"+
		"\u0000\u0000\u020d\u0211\u0005\"\u0000\u0000\u020e\u0210\b\u0004\u0000"+
		"\u0000\u020f\u020e\u0001\u0000\u0000\u0000\u0210\u0213\u0001\u0000\u0000"+
		"\u0000\u0211\u0212\u0001\u0000\u0000\u0000\u0211\u020f\u0001\u0000\u0000"+
		"\u0000\u0212\u0214\u0001\u0000\u0000\u0000\u0213\u0211\u0001\u0000\u0000"+
		"\u0000\u0214\u0215\u0005\n\u0000\u0000\u0215\u0086\u0001\u0000\u0000\u0000"+
		"\u0216\u0217\u0007\u0005\u0000\u0000\u0217\u0218\u0006C\u0001\u0000\u0218"+
		"\u0088\u0001\u0000\u0000\u0000\u0219\u021a\t\u0000\u0000\u0000\u021a\u008a"+
		"\u0001\u0000\u0000\u0000\n\u0000\u008f\u01e5\u01ea\u01f0\u01f2\u01f8\u01ff"+
		"\u0208\u0211\u0002\u0001\u0000\u0000\u0001C\u0001";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}