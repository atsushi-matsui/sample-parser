// Generated from /Users/atsushimatsui/IdeaProjects/sample-parser/parser/DefaultOrSearchQuery.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class DefaultOrSearchQueryParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		AND=1, OR=2, NOT=3, WHITE_SPACE=4, KEYWORD_CHARACTER=5, DOUBLE_QUOTE=6, 
		NL=7;
	public static final int
		RULE_andExpression = 0, RULE_orExpression = 1, RULE_notExpression = 2, 
		RULE_phrase = 3, RULE_keyword = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"andExpression", "orExpression", "notExpression", "phrase", "keyword"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'AND'", "'OR'", "'NOT'", null, null, "'\"'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "AND", "OR", "NOT", "WHITE_SPACE", "KEYWORD_CHARACTER", "DOUBLE_QUOTE", 
			"NL"
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
	public String getGrammarFileName() { return "DefaultOrSearchQuery.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public DefaultOrSearchQueryParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AndExpressionContext extends ParserRuleContext {
		public OrExpressionContext orExpression() {
			return getRuleContext(OrExpressionContext.class,0);
		}
		public AndExpressionContext andExpression() {
			return getRuleContext(AndExpressionContext.class,0);
		}
		public TerminalNode AND() { return getToken(DefaultOrSearchQueryParser.AND, 0); }
		public List<TerminalNode> WHITE_SPACE() { return getTokens(DefaultOrSearchQueryParser.WHITE_SPACE); }
		public TerminalNode WHITE_SPACE(int i) {
			return getToken(DefaultOrSearchQueryParser.WHITE_SPACE, i);
		}
		public AndExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_andExpression; }
	}

	public final AndExpressionContext andExpression() throws RecognitionException {
		return andExpression(0);
	}

	private AndExpressionContext andExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		AndExpressionContext _localctx = new AndExpressionContext(_ctx, _parentState);
		AndExpressionContext _prevctx = _localctx;
		int _startState = 0;
		enterRecursionRule(_localctx, 0, RULE_andExpression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(11);
			orExpression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(28);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new AndExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_andExpression);
					setState(13);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(15); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(14);
						match(WHITE_SPACE);
						}
						}
						setState(17); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==WHITE_SPACE );
					setState(19);
					match(AND);
					setState(21); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(20);
						match(WHITE_SPACE);
						}
						}
						setState(23); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==WHITE_SPACE );
					setState(25);
					orExpression(0);
					}
					} 
				}
				setState(30);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OrExpressionContext extends ParserRuleContext {
		public NotExpressionContext notExpression() {
			return getRuleContext(NotExpressionContext.class,0);
		}
		public OrExpressionContext orExpression() {
			return getRuleContext(OrExpressionContext.class,0);
		}
		public List<TerminalNode> WHITE_SPACE() { return getTokens(DefaultOrSearchQueryParser.WHITE_SPACE); }
		public TerminalNode WHITE_SPACE(int i) {
			return getToken(DefaultOrSearchQueryParser.WHITE_SPACE, i);
		}
		public TerminalNode OR() { return getToken(DefaultOrSearchQueryParser.OR, 0); }
		public OrExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_orExpression; }
	}

	public final OrExpressionContext orExpression() throws RecognitionException {
		return orExpression(0);
	}

	private OrExpressionContext orExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		OrExpressionContext _localctx = new OrExpressionContext(_ctx, _parentState);
		OrExpressionContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_orExpression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(32);
			notExpression();
			}
			_ctx.stop = _input.LT(-1);
			setState(43);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new OrExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_orExpression);
					setState(34);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(35);
					match(WHITE_SPACE);
					setState(38);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==OR) {
						{
						setState(36);
						match(OR);
						setState(37);
						match(WHITE_SPACE);
						}
					}

					setState(40);
					notExpression();
					}
					} 
				}
				setState(45);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NotExpressionContext extends ParserRuleContext {
		public PhraseContext phrase() {
			return getRuleContext(PhraseContext.class,0);
		}
		public TerminalNode NOT() { return getToken(DefaultOrSearchQueryParser.NOT, 0); }
		public List<TerminalNode> WHITE_SPACE() { return getTokens(DefaultOrSearchQueryParser.WHITE_SPACE); }
		public TerminalNode WHITE_SPACE(int i) {
			return getToken(DefaultOrSearchQueryParser.WHITE_SPACE, i);
		}
		public NotExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_notExpression; }
	}

	public final NotExpressionContext notExpression() throws RecognitionException {
		NotExpressionContext _localctx = new NotExpressionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_notExpression);
		int _la;
		try {
			setState(54);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_CHARACTER:
			case DOUBLE_QUOTE:
				enterOuterAlt(_localctx, 1);
				{
				setState(46);
				phrase();
				}
				break;
			case NOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(47);
				match(NOT);
				setState(49); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(48);
					match(WHITE_SPACE);
					}
					}
					setState(51); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==WHITE_SPACE );
				setState(53);
				phrase();
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
	public static class PhraseContext extends ParserRuleContext {
		public KeywordContext keyword() {
			return getRuleContext(KeywordContext.class,0);
		}
		public List<TerminalNode> DOUBLE_QUOTE() { return getTokens(DefaultOrSearchQueryParser.DOUBLE_QUOTE); }
		public TerminalNode DOUBLE_QUOTE(int i) {
			return getToken(DefaultOrSearchQueryParser.DOUBLE_QUOTE, i);
		}
		public PhraseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_phrase; }
	}

	public final PhraseContext phrase() throws RecognitionException {
		PhraseContext _localctx = new PhraseContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_phrase);
		try {
			int _alt;
			setState(65);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KEYWORD_CHARACTER:
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				keyword();
				}
				break;
			case DOUBLE_QUOTE:
				enterOuterAlt(_localctx, 2);
				{
				setState(57);
				match(DOUBLE_QUOTE);
				setState(61);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
				while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1+1 ) {
						{
						{
						setState(58);
						matchWildcard();
						}
						} 
					}
					setState(63);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
				}
				setState(64);
				match(DOUBLE_QUOTE);
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
	public static class KeywordContext extends ParserRuleContext {
		public List<TerminalNode> KEYWORD_CHARACTER() { return getTokens(DefaultOrSearchQueryParser.KEYWORD_CHARACTER); }
		public TerminalNode KEYWORD_CHARACTER(int i) {
			return getToken(DefaultOrSearchQueryParser.KEYWORD_CHARACTER, i);
		}
		public KeywordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyword; }
	}

	public final KeywordContext keyword() throws RecognitionException {
		KeywordContext _localctx = new KeywordContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_keyword);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(68); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(67);
					match(KEYWORD_CHARACTER);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(70); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 0:
			return andExpression_sempred((AndExpressionContext)_localctx, predIndex);
		case 1:
			return orExpression_sempred((OrExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean andExpression_sempred(AndExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean orExpression_sempred(OrExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u0007I\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0004\u0000\u0010"+
		"\b\u0000\u000b\u0000\f\u0000\u0011\u0001\u0000\u0001\u0000\u0004\u0000"+
		"\u0016\b\u0000\u000b\u0000\f\u0000\u0017\u0001\u0000\u0005\u0000\u001b"+
		"\b\u0000\n\u0000\f\u0000\u001e\t\u0000\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001\'\b\u0001"+
		"\u0001\u0001\u0005\u0001*\b\u0001\n\u0001\f\u0001-\t\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0004\u00022\b\u0002\u000b\u0002\f\u00023\u0001"+
		"\u0002\u0003\u00027\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0005"+
		"\u0003<\b\u0003\n\u0003\f\u0003?\t\u0003\u0001\u0003\u0003\u0003B\b\u0003"+
		"\u0001\u0004\u0004\u0004E\b\u0004\u000b\u0004\f\u0004F\u0001\u0004\u0001"+
		"=\u0002\u0000\u0002\u0005\u0000\u0002\u0004\u0006\b\u0000\u0000M\u0000"+
		"\n\u0001\u0000\u0000\u0000\u0002\u001f\u0001\u0000\u0000\u0000\u00046"+
		"\u0001\u0000\u0000\u0000\u0006A\u0001\u0000\u0000\u0000\bD\u0001\u0000"+
		"\u0000\u0000\n\u000b\u0006\u0000\uffff\uffff\u0000\u000b\f\u0003\u0002"+
		"\u0001\u0000\f\u001c\u0001\u0000\u0000\u0000\r\u000f\n\u0001\u0000\u0000"+
		"\u000e\u0010\u0005\u0004\u0000\u0000\u000f\u000e\u0001\u0000\u0000\u0000"+
		"\u0010\u0011\u0001\u0000\u0000\u0000\u0011\u000f\u0001\u0000\u0000\u0000"+
		"\u0011\u0012\u0001\u0000\u0000\u0000\u0012\u0013\u0001\u0000\u0000\u0000"+
		"\u0013\u0015\u0005\u0001\u0000\u0000\u0014\u0016\u0005\u0004\u0000\u0000"+
		"\u0015\u0014\u0001\u0000\u0000\u0000\u0016\u0017\u0001\u0000\u0000\u0000"+
		"\u0017\u0015\u0001\u0000\u0000\u0000\u0017\u0018\u0001\u0000\u0000\u0000"+
		"\u0018\u0019\u0001\u0000\u0000\u0000\u0019\u001b\u0003\u0002\u0001\u0000"+
		"\u001a\r\u0001\u0000\u0000\u0000\u001b\u001e\u0001\u0000\u0000\u0000\u001c"+
		"\u001a\u0001\u0000\u0000\u0000\u001c\u001d\u0001\u0000\u0000\u0000\u001d"+
		"\u0001\u0001\u0000\u0000\u0000\u001e\u001c\u0001\u0000\u0000\u0000\u001f"+
		" \u0006\u0001\uffff\uffff\u0000 !\u0003\u0004\u0002\u0000!+\u0001\u0000"+
		"\u0000\u0000\"#\n\u0001\u0000\u0000#&\u0005\u0004\u0000\u0000$%\u0005"+
		"\u0002\u0000\u0000%\'\u0005\u0004\u0000\u0000&$\u0001\u0000\u0000\u0000"+
		"&\'\u0001\u0000\u0000\u0000\'(\u0001\u0000\u0000\u0000(*\u0003\u0004\u0002"+
		"\u0000)\"\u0001\u0000\u0000\u0000*-\u0001\u0000\u0000\u0000+)\u0001\u0000"+
		"\u0000\u0000+,\u0001\u0000\u0000\u0000,\u0003\u0001\u0000\u0000\u0000"+
		"-+\u0001\u0000\u0000\u0000.7\u0003\u0006\u0003\u0000/1\u0005\u0003\u0000"+
		"\u000002\u0005\u0004\u0000\u000010\u0001\u0000\u0000\u000023\u0001\u0000"+
		"\u0000\u000031\u0001\u0000\u0000\u000034\u0001\u0000\u0000\u000045\u0001"+
		"\u0000\u0000\u000057\u0003\u0006\u0003\u00006.\u0001\u0000\u0000\u0000"+
		"6/\u0001\u0000\u0000\u00007\u0005\u0001\u0000\u0000\u00008B\u0003\b\u0004"+
		"\u00009=\u0005\u0006\u0000\u0000:<\t\u0000\u0000\u0000;:\u0001\u0000\u0000"+
		"\u0000<?\u0001\u0000\u0000\u0000=>\u0001\u0000\u0000\u0000=;\u0001\u0000"+
		"\u0000\u0000>@\u0001\u0000\u0000\u0000?=\u0001\u0000\u0000\u0000@B\u0005"+
		"\u0006\u0000\u0000A8\u0001\u0000\u0000\u0000A9\u0001\u0000\u0000\u0000"+
		"B\u0007\u0001\u0000\u0000\u0000CE\u0005\u0005\u0000\u0000DC\u0001\u0000"+
		"\u0000\u0000EF\u0001\u0000\u0000\u0000FD\u0001\u0000\u0000\u0000FG\u0001"+
		"\u0000\u0000\u0000G\t\u0001\u0000\u0000\u0000\n\u0011\u0017\u001c&+36"+
		"=AF";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}