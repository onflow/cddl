grammar CDDL;

cddl : S? (rule S?)+ S? EOF;

rule : typeRule
     | groupRule
     ;

typeRule : ID genericParam? S? assignType S? type ;

groupRule : ID genericParam? S? assignGroup S? groupEntry ;

assignType : '=' | '/=' ;
assignGroup : '=' | '//=' ;

genericParam : '<' S? ID S? (',' S? ID S? )* '>' ;
genericArg : '<' S? type1 S? (',' S? type1 S? )* '>' ;

type : type1 (S? '/' S? type1)* ;

type1 : type2 (S? (RANGEOP | CTLOP) S? type2)? ;
/* space may be needed before the operator if type2 ends in a name */

type2 : VALUE
      | ID genericArg?
      | '(' S? type S? ')'
      | '{' S? group S? '}'
      | '[' S? group S? ']'
      | '~' S? ID genericArg?
      | '&' S? '(' S? group S? ')'
      | '&' S? ID genericArg?
      | TAG1 '(' S? type S? ')'
      | TAG2              /* major/ai */
      | '#'               /* any */
      ;

TAG1 : '#' '6' ('.' UINT)? ;

TAG2 : '#' DIGIT ('.' UINT)? ;

RANGEOP : '...'
        | '..'
        ;

CTLOP : '.' ID ;

group : groupChoice (S? '//' S? groupChoice)* ;

groupChoice : (groupEntry optComma)* ;

groupEntry : (OCCUR S?)? (memberKey S?)? type
            | (OCCUR S?)? ID genericArg?  /* preempted by above */
            | (OCCUR S?)? '(' S? group S? ')'
            ;

memberKey : type1 S? ('^' S?)? '=>'
          | ID S? ':'
          | VALUE S? ':'
          ;

optComma : S? ','? S? ;

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