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
		COMENTARIO=1, ALGORITMO=2, DECLARE=3, LITERAL=4, INTEIRO=5, LEIA=6, ESCREVA=7, 
		FIM_ALGORITMO=8, DOISPONTOS=9, ABREPAR=10, FECHAPAR=11, VIRGULA=12, IDENT=13, 
		CADEIA=14, WS=15;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"COMENTARIO", "ALGORITMO", "DECLARE", "LITERAL", "INTEIRO", "LEIA", "ESCREVA", 
			"FIM_ALGORITMO", "DOISPONTOS", "ABREPAR", "FECHAPAR", "VIRGULA", "IDENT", 
			"CADEIA", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'algoritmo'", "'declare'", "'literal'", "'inteiro'", "'leia'", 
			"'escreva'", "'fim_algoritmo'", "':'", "'('", "')'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "COMENTARIO", "ALGORITMO", "DECLARE", "LITERAL", "INTEIRO", "LEIA", 
			"ESCREVA", "FIM_ALGORITMO", "DOISPONTOS", "ABREPAR", "FECHAPAR", "VIRGULA", 
			"IDENT", "CADEIA", "WS"
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
		case 14:
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
		"\u0004\u0000\u000f\u0083\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0001\u0000\u0001\u0000\u0005\u0000\"\b\u0000\n\u0000\f\u0000%\t\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0005\fs\b\f\n\f\f\fv\t\f\u0001\r\u0001\r\u0005\rz\b"+
		"\r\n\r\f\r}\t\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0002"+
		"#{\u0000\u000f\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005"+
		"\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019"+
		"\r\u001b\u000e\u001d\u000f\u0001\u0000\u0004\u0001\u0000\n\n\u0002\u0000"+
		"AZaz\u0003\u000009AZaz\u0003\u0000\t\n\r\r  \u0085\u0000\u0001\u0001\u0000"+
		"\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000"+
		"\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000"+
		"\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000"+
		"\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000"+
		"\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000\u0000"+
		"\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000\u0000"+
		"\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000\u0000"+
		"\u0001\u001f\u0001\u0000\u0000\u0000\u0003+\u0001\u0000\u0000\u0000\u0005"+
		"5\u0001\u0000\u0000\u0000\u0007=\u0001\u0000\u0000\u0000\tE\u0001\u0000"+
		"\u0000\u0000\u000bM\u0001\u0000\u0000\u0000\rR\u0001\u0000\u0000\u0000"+
		"\u000fZ\u0001\u0000\u0000\u0000\u0011h\u0001\u0000\u0000\u0000\u0013j"+
		"\u0001\u0000\u0000\u0000\u0015l\u0001\u0000\u0000\u0000\u0017n\u0001\u0000"+
		"\u0000\u0000\u0019p\u0001\u0000\u0000\u0000\u001bw\u0001\u0000\u0000\u0000"+
		"\u001d\u0080\u0001\u0000\u0000\u0000\u001f#\u0005{\u0000\u0000 \"\b\u0000"+
		"\u0000\u0000! \u0001\u0000\u0000\u0000\"%\u0001\u0000\u0000\u0000#$\u0001"+
		"\u0000\u0000\u0000#!\u0001\u0000\u0000\u0000$&\u0001\u0000\u0000\u0000"+
		"%#\u0001\u0000\u0000\u0000&\'\u0005 \u0000\u0000\'(\u0005}\u0000\u0000"+
		"()\u0001\u0000\u0000\u0000)*\u0006\u0000\u0000\u0000*\u0002\u0001\u0000"+
		"\u0000\u0000+,\u0005a\u0000\u0000,-\u0005l\u0000\u0000-.\u0005g\u0000"+
		"\u0000./\u0005o\u0000\u0000/0\u0005r\u0000\u000001\u0005i\u0000\u0000"+
		"12\u0005t\u0000\u000023\u0005m\u0000\u000034\u0005o\u0000\u00004\u0004"+
		"\u0001\u0000\u0000\u000056\u0005d\u0000\u000067\u0005e\u0000\u000078\u0005"+
		"c\u0000\u000089\u0005l\u0000\u00009:\u0005a\u0000\u0000:;\u0005r\u0000"+
		"\u0000;<\u0005e\u0000\u0000<\u0006\u0001\u0000\u0000\u0000=>\u0005l\u0000"+
		"\u0000>?\u0005i\u0000\u0000?@\u0005t\u0000\u0000@A\u0005e\u0000\u0000"+
		"AB\u0005r\u0000\u0000BC\u0005a\u0000\u0000CD\u0005l\u0000\u0000D\b\u0001"+
		"\u0000\u0000\u0000EF\u0005i\u0000\u0000FG\u0005n\u0000\u0000GH\u0005t"+
		"\u0000\u0000HI\u0005e\u0000\u0000IJ\u0005i\u0000\u0000JK\u0005r\u0000"+
		"\u0000KL\u0005o\u0000\u0000L\n\u0001\u0000\u0000\u0000MN\u0005l\u0000"+
		"\u0000NO\u0005e\u0000\u0000OP\u0005i\u0000\u0000PQ\u0005a\u0000\u0000"+
		"Q\f\u0001\u0000\u0000\u0000RS\u0005e\u0000\u0000ST\u0005s\u0000\u0000"+
		"TU\u0005c\u0000\u0000UV\u0005r\u0000\u0000VW\u0005e\u0000\u0000WX\u0005"+
		"v\u0000\u0000XY\u0005a\u0000\u0000Y\u000e\u0001\u0000\u0000\u0000Z[\u0005"+
		"f\u0000\u0000[\\\u0005i\u0000\u0000\\]\u0005m\u0000\u0000]^\u0005_\u0000"+
		"\u0000^_\u0005a\u0000\u0000_`\u0005l\u0000\u0000`a\u0005g\u0000\u0000"+
		"ab\u0005o\u0000\u0000bc\u0005r\u0000\u0000cd\u0005i\u0000\u0000de\u0005"+
		"t\u0000\u0000ef\u0005m\u0000\u0000fg\u0005o\u0000\u0000g\u0010\u0001\u0000"+
		"\u0000\u0000hi\u0005:\u0000\u0000i\u0012\u0001\u0000\u0000\u0000jk\u0005"+
		"(\u0000\u0000k\u0014\u0001\u0000\u0000\u0000lm\u0005)\u0000\u0000m\u0016"+
		"\u0001\u0000\u0000\u0000no\u0005,\u0000\u0000o\u0018\u0001\u0000\u0000"+
		"\u0000pt\u0007\u0001\u0000\u0000qs\u0007\u0002\u0000\u0000rq\u0001\u0000"+
		"\u0000\u0000sv\u0001\u0000\u0000\u0000tr\u0001\u0000\u0000\u0000tu\u0001"+
		"\u0000\u0000\u0000u\u001a\u0001\u0000\u0000\u0000vt\u0001\u0000\u0000"+
		"\u0000w{\u0005\"\u0000\u0000xz\b\u0000\u0000\u0000yx\u0001\u0000\u0000"+
		"\u0000z}\u0001\u0000\u0000\u0000{|\u0001\u0000\u0000\u0000{y\u0001\u0000"+
		"\u0000\u0000|~\u0001\u0000\u0000\u0000}{\u0001\u0000\u0000\u0000~\u007f"+
		"\u0005\"\u0000\u0000\u007f\u001c\u0001\u0000\u0000\u0000\u0080\u0081\u0007"+
		"\u0003\u0000\u0000\u0081\u0082\u0006\u000e\u0001\u0000\u0082\u001e\u0001"+
		"\u0000\u0000\u0000\u0004\u0000#t{\u0002\u0001\u0000\u0000\u0001\u000e"+
		"\u0001";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}