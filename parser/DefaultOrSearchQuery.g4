grammar DefaultOrSearchQuery;

andExpression
   :orExpression
   |andExpression WHITE_SPACE+ AND WHITE_SPACE+ orExpression
   ;

orExpression
   :notExpression 
   |orExpression WHITE_SPACE (OR WHITE_SPACE)? notExpression
   |KEYWORD_CHARACTER+ PHRASE
   |PHRASE KEYWORD_CHARACTER+
   ;

notExpression
   :phrase
   |NOT WHITE_SPACE+ phrase
   ;

phrase
   :keyword
   |PHRASE
   ;

keyword
   :KEYWORD_CHARACTER+
   |DOUBLE_QUOTE+
   ;

AND
   :'AND'
   ;
OR
   :'OR'
   ;
NOT
   :'NOT'
   ;

WHITE_SPACE
   :[ \t\u3000]+
   ;

KEYWORD_CHARACTER
   :~[" \n\r\t\u3000]
   ;

DOUBLE_QUOTE
   :'"'
   ;

PHRASE
   :DOUBLE_QUOTE .*? DOUBLE_QUOTE
   ;

NL
   :[\n\r]+ -> skip
   ;
