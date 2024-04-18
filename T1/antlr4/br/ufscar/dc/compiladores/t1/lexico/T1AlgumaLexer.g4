lexer grammar T1AlgumaLexer;

//Regra para comentario. Detecta uma cadeia iniciada por { e finalizada por }, 
//aceitando qualquer caractere entre { e } diferente de \n (incluindo nenhum caractere) 
COMENTARIO       : '{' ~('\n')*?' }' {skip();};

//Regra para pular whitespaces
WS               : ( ' ' | '\t' | '\r' | '\n' ) {skip();};

// Palavras chaves.
ALGORITMO: 'algoritmo'; 