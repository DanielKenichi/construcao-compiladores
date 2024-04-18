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
		COMENTARIO=1, WS=2, ALGORITMO=3;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"COMENTARIO", "WS", "ALGORITMO"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, "'algoritmo'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "COMENTARIO", "WS", "ALGORITMO"
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
		case 1:
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
		"\u0004\u0000\u0003 \u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0001\u0000\u0001\u0000\u0005\u0000"+
		"\n\b\u0000\n\u0000\f\u0000\r\t\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u000b\u0000\u0003\u0001\u0001"+
		"\u0003\u0002\u0005\u0003\u0001\u0000\u0002\u0001\u0000\n\n\u0003\u0000"+
		"\t\n\r\r   \u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000"+
		"\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0001\u0007\u0001\u0000"+
		"\u0000\u0000\u0003\u0013\u0001\u0000\u0000\u0000\u0005\u0016\u0001\u0000"+
		"\u0000\u0000\u0007\u000b\u0005{\u0000\u0000\b\n\b\u0000\u0000\u0000\t"+
		"\b\u0001\u0000\u0000\u0000\n\r\u0001\u0000\u0000\u0000\u000b\f\u0001\u0000"+
		"\u0000\u0000\u000b\t\u0001\u0000\u0000\u0000\f\u000e\u0001\u0000\u0000"+
		"\u0000\r\u000b\u0001\u0000\u0000\u0000\u000e\u000f\u0005 \u0000\u0000"+
		"\u000f\u0010\u0005}\u0000\u0000\u0010\u0011\u0001\u0000\u0000\u0000\u0011"+
		"\u0012\u0006\u0000\u0000\u0000\u0012\u0002\u0001\u0000\u0000\u0000\u0013"+
		"\u0014\u0007\u0001\u0000\u0000\u0014\u0015\u0006\u0001\u0001\u0000\u0015"+
		"\u0004\u0001\u0000\u0000\u0000\u0016\u0017\u0005a\u0000\u0000\u0017\u0018"+
		"\u0005l\u0000\u0000\u0018\u0019\u0005g\u0000\u0000\u0019\u001a\u0005o"+
		"\u0000\u0000\u001a\u001b\u0005r\u0000\u0000\u001b\u001c\u0005i\u0000\u0000"+
		"\u001c\u001d\u0005t\u0000\u0000\u001d\u001e\u0005m\u0000\u0000\u001e\u001f"+
		"\u0005o\u0000\u0000\u001f\u0006\u0001\u0000\u0000\u0000\u0002\u0000\u000b"+
		"\u0002\u0001\u0000\u0000\u0001\u0001\u0001";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}