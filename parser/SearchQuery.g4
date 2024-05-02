grammar SearchQuery;
/**
 * parser
 */
expr:   term | expr OR_OPERATOR term;
term:   factor | term AND_OPERATOR factor;
factor: keywords | NOT_OPERATOR keywords;
keywords: '(' expr ')' | alphabets;
alphabets: ALPHABETS;

/**
 * lexer
 */
ALPHABETS
    : AND_OR_STRING
    ;
OR_OPERATOR : 'OR';
AND_OPERATOR : 'AND';
NOT_OPERATOR : 'NOT';
WHITE_SPACES : [ \t\n\r]+ -> skip ;
fragment DIGITS
    : [0-9]+
    ;
fragment HEX
    : ('%' [a-fA-F0-9] [a-fA-F0-9])+
    ;
fragment STRING
    : ([a-zA-Z~0-9] | HEX) ([a-zA-Z0-9.+-] | HEX)*
    ;
fragment AND_OR: ('AND' | 'OR');
fragment AND_OR_STRING
    : (AND_OR [a-zA-Z0-9.+-] | HEX)+;