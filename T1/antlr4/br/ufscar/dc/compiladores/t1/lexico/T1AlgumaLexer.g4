lexer grammar T1AlgumaLexer;

//Regra para ignorar comentario. Detecta uma cadeia iniciada por '{' e finalizada por '}', 
//aceitando qualquer caractere entre '{' e '}' diferente de '\n' 
COMENTARIO: '{' ~('\n')*?' }' {skip();};

//Palavras Chaves.
ALGORITMO: 'algoritmo'; 
DECLARE: 'declare';
LITERAL: 'literal';
INTEIRO: 'inteiro';
LEIA: 'leia';
ESCREVA: 'escreva';
FIM_ALGORITMO: 'fim_algoritmo';

//Delimitador
DOISPONTOS: ':';
ABREPAR: '(';
FECHAPAR: ')';
VIRGULA: ',';

//Regra para identificadores (variaveis). identificadores validos
//começam com uma letra, e podem conter qualquer caracter alfanumérico
IDENT: ('a'..'z'|'A'..'Z')('a'..'z'|'A'..'Z'|'0'..'9')*;

//Regra para detectar cadeias de caracteres. Detecta caracteres diferente
// de '\n' que estejam entre aspas duplas
CADEIA: '"' ~('\n')*? '"';

//Regra para pular whitespaces
WS: ( ' ' | '\t' | '\r' | '\n' ) {skip();};