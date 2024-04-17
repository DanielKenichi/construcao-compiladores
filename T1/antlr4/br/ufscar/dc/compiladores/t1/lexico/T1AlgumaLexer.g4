lexer grammar T1AlgumaLexer;


COMENTARIO       : '{' ~('\n')*?' }' {skip();};

WS               : ( ' ' | '\t' | '\r' | '\n' ) {skip();};
// Palavras chaves.