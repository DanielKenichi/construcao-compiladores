lexer grammar T1AlgumaLexer;

//Regra para ignorar comentario. Detecta uma cadeia iniciada por '{' e finalizada por '}', 
//aceitando qualquer caractere entre '{' e '}' diferente de '\n' 
COMENTARIO: '{' ~('\n')*? '}' -> skip;

//Palavras Chaves.
ALGORITMO: 'algoritmo';
DECLARE: 'declare';
VAR: 'var';
CONSTANTE: 'constante';
LITERAL: 'literal';
INTEIRO: 'inteiro';
REAL: 'real';
LOGICO: 'logico';
VERDADEIRO: 'verdadeiro';
FALSO: 'falso';
E: 'e';
OR: 'ou';
NAO: 'nao';
SE: 'se';
FIM_SE: 'fim_se';
ENTAO: 'entao';
SENAO: 'senao';
CASO: 'caso';
SEJA: 'seja';
FIM_CASO: 'fim_caso';
PARA: 'para';
FIM_PARA: 'fim_para';
ATE: 'ate';
FACA: 'faca';
ENQUANTO: 'enquanto';
FIM_ENQUANTO: 'fim_enquanto';
TIPO: 'tipo';
REGISTRO: 'registro';
FIM_REGISTRO: 'fim_registro';
PROCEDIMENTO: 'procedimento';
FIM_PROCEDIMENTO : 'fim_procedimento';
FUNCAO: 'funcao';
FIM_FUNCAO: 'fim_funcao';
RETORNE: 'retorne';
LEIA: 'leia';
ESCREVA: 'escreva';
FIM_ALGORITMO: 'fim_algoritmo';

//Intervalo de valores.
INTERVALO: '..';

//Operadores Relacionais.
MENOR: '<';
MENORIGUAL: '<=';
MAIOR: '>';
MAIORIGUAL: '>=';
IGUAL: '=';
DIFERENTE: '<>';

//Simbolos Delimitadores.
DOISPONTOS: ':';
ABREPAR: '(';
FECHAPAR: ')';
ABRECHAVE: '[';
FECHACHAVE: ']';
VIRGULA: ',';
ASPAS: '"';

//Operadores aritméticos.
DIVISAO : '/';
RESTO: '%';
SOMA: '+';
SUBTRACAO: '-';
MULTIPLICACAO: '*';

//Atribuicao
ATRIBUICAO       : '<-';
//Ponteiro
PONTEIRO         : '^';
//Endereco 
ENDERECO         : '&';
//ponto
PONTO            : '.';

//Regra para identificar números inteiros e reais sem sinal.
NUM_INT: ('0'..'9')+;
NUM_REAL: ('0'..'9')+ ('.' ('0'..'9')+)?;

//Regra para identificadores (variaveis). identificadores validos
//começam com uma letra, e podem conter qualquer caracter alfanumérico
//incluindo '_'
IDENT: ('a'..'z'|'A'..'Z')('a'..'z'|'A'..'Z'|'0'..'9'|'_')*;

//Regra para detectar cadeias de caracteres. Detecta caracteres diferente
// de '\n' que estejam entre aspas duplas
CADEIA: '"' ~('\n')*? '"';

//Regra para detectar comentarios abertos
ERRO_COMENTARIO_ABERTO : '{' (~('\n'|'}'))*? '\n';

//Regra para detectar cadeias de caracteres abertas
ERRO_CADEIA_ABERTA : '"' ( ~('\n'|'"') )*? '\n';

//Regra para pular whitespaces
WS: ( ' ' | '\t' | '\r' | '\n' ) -> skip;

//Erro caso nenhuma regra identifique o simbolo.
ERRO: .;