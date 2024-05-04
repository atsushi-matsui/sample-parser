grammar DefaultOrSearchQuery;

andExpression
   :orExpression
   |andExpression WHITE_SPACE+ AND WHITE_SPACE+ orExpression
   ;

orExpression
   :notExpression 
   |orExpression WHITE_SPACE (OR WHITE_SPACE)? notExpression
   ;

notExpression
   :phrase
   |NOT WHITE_SPACE+ phrase
   ;

phrase
   :keyword
   |DOUBLE_QUOTE .*? DOUBLE_QUOTE
   ;

keyword
   :KEYWORD_CHARACTER+
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

NL
   :[\n\r]+ -> skip
   ;
