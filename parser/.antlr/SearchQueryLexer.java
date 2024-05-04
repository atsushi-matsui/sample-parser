// Generated from /Users/atsushimatsui/IdeaProjects/sample-parser/parser/SearchQuery.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class SearchQueryLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, ALPHABETS=3, OR_OPERATOR=4, AND_OPERATOR=5, NOT_OPERATOR=6, 
		WHITE_SPACES=7;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "ALPHABETS", "OR_OPERATOR", "AND_OPERATOR", "NOT_OPERATOR", 
			"WHITE_SPACES", "DIGITS", "HEX", "STRING", "AND_OR", "AND_OR_STRING"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", null, "'OR'", "'AND'", "'NOT'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, "ALPHABETS", "OR_OPERATOR", "AND_OPERATOR", "NOT_OPERATOR", 
			"WHITE_SPACES"
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


	public SearchQueryLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SearchQuery.g4"; }

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

	public static final String _serializedATN =
		"\u0004\u0000\u0007W\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0004\u0006,\b\u0006\u000b\u0006\f\u0006-\u0001\u0006\u0001"+
		"\u0006\u0001\u0007\u0004\u00073\b\u0007\u000b\u0007\f\u00074\u0001\b\u0001"+
		"\b\u0001\b\u0004\b:\b\b\u000b\b\f\b;\u0001\t\u0001\t\u0003\t@\b\t\u0001"+
		"\t\u0001\t\u0005\tD\b\t\n\t\f\tG\t\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0003\nN\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0004"+
		"\u000bT\b\u000b\u000b\u000b\f\u000bU\u0000\u0000\f\u0001\u0001\u0003\u0002"+
		"\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\u0000\u0011"+
		"\u0000\u0013\u0000\u0015\u0000\u0017\u0000\u0001\u0000\u0005\u0003\u0000"+
		"\t\n\r\r  \u0001\u000009\u0003\u000009AFaf\u0004\u000009AZaz~~\u0005\u0000"+
		"++-.09AZazZ\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000"+
		"\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000"+
		"\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000"+
		"\u0000\u0000\r\u0001\u0000\u0000\u0000\u0001\u0019\u0001\u0000\u0000\u0000"+
		"\u0003\u001b\u0001\u0000\u0000\u0000\u0005\u001d\u0001\u0000\u0000\u0000"+
		"\u0007\u001f\u0001\u0000\u0000\u0000\t\"\u0001\u0000\u0000\u0000\u000b"+
		"&\u0001\u0000\u0000\u0000\r+\u0001\u0000\u0000\u0000\u000f2\u0001\u0000"+
		"\u0000\u0000\u00119\u0001\u0000\u0000\u0000\u0013?\u0001\u0000\u0000\u0000"+
		"\u0015M\u0001\u0000\u0000\u0000\u0017S\u0001\u0000\u0000\u0000\u0019\u001a"+
		"\u0005(\u0000\u0000\u001a\u0002\u0001\u0000\u0000\u0000\u001b\u001c\u0005"+
		")\u0000\u0000\u001c\u0004\u0001\u0000\u0000\u0000\u001d\u001e\u0003\u0017"+
		"\u000b\u0000\u001e\u0006\u0001\u0000\u0000\u0000\u001f \u0005O\u0000\u0000"+
		" !\u0005R\u0000\u0000!\b\u0001\u0000\u0000\u0000\"#\u0005A\u0000\u0000"+
		"#$\u0005N\u0000\u0000$%\u0005D\u0000\u0000%\n\u0001\u0000\u0000\u0000"+
		"&\'\u0005N\u0000\u0000\'(\u0005O\u0000\u0000()\u0005T\u0000\u0000)\f\u0001"+
		"\u0000\u0000\u0000*,\u0007\u0000\u0000\u0000+*\u0001\u0000\u0000\u0000"+
		",-\u0001\u0000\u0000\u0000-+\u0001\u0000\u0000\u0000-.\u0001\u0000\u0000"+
		"\u0000./\u0001\u0000\u0000\u0000/0\u0006\u0006\u0000\u00000\u000e\u0001"+
		"\u0000\u0000\u000013\u0007\u0001\u0000\u000021\u0001\u0000\u0000\u0000"+
		"34\u0001\u0000\u0000\u000042\u0001\u0000\u0000\u000045\u0001\u0000\u0000"+
		"\u00005\u0010\u0001\u0000\u0000\u000067\u0005%\u0000\u000078\u0007\u0002"+
		"\u0000\u00008:\u0007\u0002\u0000\u000096\u0001\u0000\u0000\u0000:;\u0001"+
		"\u0000\u0000\u0000;9\u0001\u0000\u0000\u0000;<\u0001\u0000\u0000\u0000"+
		"<\u0012\u0001\u0000\u0000\u0000=@\u0007\u0003\u0000\u0000>@\u0003\u0011"+
		"\b\u0000?=\u0001\u0000\u0000\u0000?>\u0001\u0000\u0000\u0000@E\u0001\u0000"+
		"\u0000\u0000AD\u0007\u0004\u0000\u0000BD\u0003\u0011\b\u0000CA\u0001\u0000"+
		"\u0000\u0000CB\u0001\u0000\u0000\u0000DG\u0001\u0000\u0000\u0000EC\u0001"+
		"\u0000\u0000\u0000EF\u0001\u0000\u0000\u0000F\u0014\u0001\u0000\u0000"+
		"\u0000GE\u0001\u0000\u0000\u0000HI\u0005A\u0000\u0000IJ\u0005N\u0000\u0000"+
		"JN\u0005D\u0000\u0000KL\u0005O\u0000\u0000LN\u0005R\u0000\u0000MH\u0001"+
		"\u0000\u0000\u0000MK\u0001\u0000\u0000\u0000N\u0016\u0001\u0000\u0000"+
		"\u0000OP\u0003\u0015\n\u0000PQ\u0007\u0004\u0000\u0000QT\u0001\u0000\u0000"+
		"\u0000RT\u0003\u0011\b\u0000SO\u0001\u0000\u0000\u0000SR\u0001\u0000\u0000"+
		"\u0000TU\u0001\u0000\u0000\u0000US\u0001\u0000\u0000\u0000UV\u0001\u0000"+
		"\u0000\u0000V\u0018\u0001\u0000\u0000\u0000\n\u0000-4;?CEMSU\u0001\u0006"+
		"\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}