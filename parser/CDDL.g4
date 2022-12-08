grammar CDDL;

// Implements RFC 8610, Appendix B. ABNF Grammar

cddl : (rules+=rule)+ EOF;

rule : ID genericParam? assignType type           #TypeRule
     | ID genericParam? assignGroup groupEntry    #GroupRule
     ;

assignType : '=' | '/=' ;
assignGroup : '=' | '//=' ;

genericParam : '<' ids+=ID (',' ids+=ID )* '>' ;
genericArg : '<' types+=type1 (',' types+=type1 )* '>' ;

type : types+=type1 ( '/' types+=type1)* ;

type1 : type2 ( (RANGEOP | CTLOP) type2)? ;
/* space may be needed before the operator if type2 ends in a name */

type2 : value=VALUE                  #ValueExpr
      | id=ID arg=genericArg?        #IdExpr
      | '(' type ')'                 #GroupExpr
      | '{' group '}'                #MapExpr
      | '[' group ']'                #ArrayExpr
      | '~' id=ID arg=genericArg?    #UnwrapExpr
      | '&' '(' group ')'            #ChoiceExpr
      | '&' id=ID arg=genericArg?    #ChoiceIDExpr
      | tag=TAG '(' type ')'         #TaggedExpr
      | major=MAJOR                  #MajorExpr
      | '#'                           #AnyExpr
      ;

TAG : '#' '6' ('.' UINT)? ;

MAJOR : '#' DIGIT ('.' UINT)? ;

RANGEOP : '...'
        | '..'
        ;

CTLOP : '.' ID ;

group : groups+=groupChoice ( '//' groups+=groupChoice)* ;

groupChoice : (entry=groupEntry optComma)* ;

groupEntry : OCCUR? memberKey? type                                    #TypeEntry
           | OCCUR? id=ID arg=genericArg?  /* preempted by above */    #NameEntry
           | OCCUR? '(' group ')'                                      #GroupsEntry
           ;

memberKey : type1 cut='^'? '=>'    #TypeMember
          | id=ID ':'              #NameMember
          | value=VALUE ':'        #ValueMember
          ;

optComma : ','? ;

OCCUR : UINT? '*' UINT?
      | '+'
      | '?'
      ;

fragment UINT : DIGIT1 DIGIT*
     | '0x' HEXDIG+
     | '0b' BINDIG+
     | '0'
     ;

VALUE : NUMBER
      | TEXT
      | BYTES
      ;

fragment INT : '-'? UINT ;

/* This is a float if it has fraction or exponent; int otherwise */
fragment NUMBER : HEXFLOAT
       | INT ('.' FRACTION)? ('e' EXPONENT)?
       ;

fragment HEXFLOAT : '-'? '0x' HEXDIG+ ('.' HEXDIG+)? 'p' EXPONENT ;
fragment FRACTION : DIGIT+ ;
fragment EXPONENT : ('+' | '-')? DIGIT+ ;

fragment TEXT : [\u0022] SCHAR* [\u0022] ;
fragment SCHAR : [\u0020-\u0021\u0023-\u005B\u005D-\u007E\u0080-\u{10FFFD}] | SESC ;
fragment SESC : '\\' [\u0020-\u007E\u0080-\u{10FFFD}] ;

fragment BYTES : BSQUAL? [\u0027] BCHAR* [\u0027] ;
fragment BCHAR : [\u0020-\u0026\u0028-\u005B\u005D-\u{10FFFD}] | SESC | CRLF ;
fragment BSQUAL : 'h' | 'b64' ;

ID : EALPHA (('-' | '.')* (EALPHA | DIGIT))* ;

fragment ALPHA : [\u0041-\u005A\u0061-\u007A] ;
fragment EALPHA : ALPHA | '@' | '_' | '$' ;
fragment DIGIT : [\u0030-\u0039] ;
fragment DIGIT1 : [\u0031-\u0039] ;
fragment HEXDIG : DIGIT | 'A' | 'B' | 'C' | 'D' | 'E' | 'F' ;
fragment BINDIG : [\u0030-\u0031] ;

S : WS+ -> skip ;
fragment WS : SP | NL ;
fragment SP : [ \t] ;
fragment NL : COMMENT | CRLF ;
fragment COMMENT : ';' PCHAR* CRLF ;
fragment PCHAR : [\u0020-\u007E\u0080-\u{10FFFD}] ;
fragment CRLF : [\u000A]
              | [\u000D\u000A]
              ;