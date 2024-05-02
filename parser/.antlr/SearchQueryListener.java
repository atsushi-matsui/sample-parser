// Generated from /Users/matsui-atsushi/IdeaProject/sample-parser/parser/SearchQuery.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SearchQueryParser}.
 */
public interface SearchQueryListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SearchQueryParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(SearchQueryParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SearchQueryParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(SearchQueryParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SearchQueryParser#term}.
	 * @param ctx the parse tree
	 */
	void enterTerm(SearchQueryParser.TermContext ctx);
	/**
	 * Exit a parse tree produced by {@link SearchQueryParser#term}.
	 * @param ctx the parse tree
	 */
	void exitTerm(SearchQueryParser.TermContext ctx);
	/**
	 * Enter a parse tree produced by {@link SearchQueryParser#factor}.
	 * @param ctx the parse tree
	 */
	void enterFactor(SearchQueryParser.FactorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SearchQueryParser#factor}.
	 * @param ctx the parse tree
	 */
	void exitFactor(SearchQueryParser.FactorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SearchQueryParser#keywords}.
	 * @param ctx the parse tree
	 */
	void enterKeywords(SearchQueryParser.KeywordsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SearchQueryParser#keywords}.
	 * @param ctx the parse tree
	 */
	void exitKeywords(SearchQueryParser.KeywordsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SearchQueryParser#alphabets}.
	 * @param ctx the parse tree
	 */
	void enterAlphabets(SearchQueryParser.AlphabetsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SearchQueryParser#alphabets}.
	 * @param ctx the parse tree
	 */
	void exitAlphabets(SearchQueryParser.AlphabetsContext ctx);
}