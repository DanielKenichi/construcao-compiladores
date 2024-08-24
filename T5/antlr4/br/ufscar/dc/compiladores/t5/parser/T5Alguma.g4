grammar T5Alguma;

// -----------------------------------
// PARTE LÉXICA (T1)
// -----------------------------------

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
OU: 'ou';
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


// -----------------------------------
// PARTE SINTÁTICA (T2)
// -----------------------------------

programa
  : declaracoes ALGORITMO corpo FIM_ALGORITMO
  ;

// Declarações de variáveis e funções
declaracoes
  : (declaracoes_variaveis | declaracoes_funcoes)*
  ;

// Seção das variaveis
declaracoes_variaveis
  : DECLARE variavel
  | CONSTANTE IDENT DOISPONTOS tipo_basico IGUAL valor_constante
  | TIPO IDENT DOISPONTOS registro
  ;
variavel
  : identificador (VIRGULA identificador)* DOISPONTOS tipo
  ;

identificador
  : IDENT (PONTO IDENT)* (ABRECHAVE exp_aritmetica FECHACHAVE)*
  ;
// Regras de definição de operações aritméticas, separadas para definir a ordem de prioridade
exp_aritmetica
  : termo (op1 termo)*
  ;
termo
  : fator (op2 fator)*
  ;
fator
  : parcela (op3 parcela)*
  ;
op1
  : SOMA | SUBTRACAO
  ;
op2
  : MULTIPLICACAO | DIVISAO
  ;
op3
  : RESTO
  ;
// Regras de definição de expressões aritméticas, com variáveis, ou números constantes.
parcela
  : op_unario? parcela_unario | parcela_nao_unario
  ;
parcela_unario
  : PONTEIRO? identificador
	| IDENT ABREPAR expressao (VIRGULA expressao)* FECHAPAR
	| NUM_INT
	| NUM_REAL
	| ABREPAR expressao FECHAPAR
    ;
// Regra para definir se o número da regra do comando "Caso" é negativo.
op_unario
  : SUBTRACAO
  ;

// Regra de definição de recuperação do valor de um endereço da variável.
parcela_nao_unario
  : ENDERECO identificador | CADEIA
  ;

// Regra para variáveis mais complexas (estruturas/ponteiros)
tipo
  : registro
  | tipo_variavel
  ;

// Tipos básicos de variaveis
tipo_basico
  : LITERAL
  | INTEIRO
  | REAL
  | LOGICO
  ;
// Declaração de ponteiro (qualquer tipo de var)
tipo_variavel
  : PONTEIRO? (tipo_basico | IDENT)
  ;

// Valores constantes
valor_constante
  : CADEIA
  | NUM_INT
  | NUM_REAL
  | VERDADEIRO
  | FALSO
  ;

// Declaração de "estrutura de dados"
registro
  : REGISTRO variavel* FIM_REGISTRO
  ;

parametro
  : VAR? identificador (VIRGULA identificador)* DOISPONTOS tipo_variavel
  ;

parametros
  : parametro (VIRGULA parametro)*
  ;

// Regra de declaração de funções e procedimentos.
declaracoes_funcoes
  : PROCEDIMENTO IDENT ABREPAR parametros? FECHAPAR declaracoes_variaveis* cmd* FIM_PROCEDIMENTO
  | FUNCAO IDENT ABREPAR parametros? FECHAPAR DOISPONTOS tipo_variavel declaracoes_variaveis* cmd* FIM_FUNCAO
  ;

// Regra de definição do corpo de uma função ou procedimento.
corpo
  : declaracoes_variaveis* cmd*
  ;

// Regras de definições de comandos da linguagem.
cmd
  : cmdLeia
  | cmdEscreva
  | cmdSe
  | cmdCaso
  | cmdPara
  | cmdEnquanto
  | cmdFaca
  | cmdAtribuicao
  | cmdChamada
  | cmdRetorne
  ;
cmdLeia
    : LEIA ABREPAR PONTEIRO? identificador (VIRGULA PONTEIRO? identificador)* FECHAPAR
    ;
cmdEscreva
    : ESCREVA ABREPAR expressao (VIRGULA expressao)* FECHAPAR
    ;
cmdSe
    : SE expressao ENTAO seCmds+=cmd* (SENAO senaoCmds+=cmd*)? FIM_SE
    ;
cmdCaso
    : CASO exp_aritmetica SEJA selecao (SENAO cmd*)?  FIM_CASO
    ;
cmdPara
    : PARA IDENT ATRIBUICAO exp1=exp_aritmetica ATE exp2=exp_aritmetica FACA cmd* FIM_PARA
    ;
cmdEnquanto
    : ENQUANTO expressao FACA cmd* FIM_ENQUANTO
    ;
cmdFaca
    : FACA cmd* ATE expressao
    ;
cmdAtribuicao
    : PONTEIRO? identificador ATRIBUICAO expressao
    ;
cmdChamada
    : IDENT ABREPAR expressao (VIRGULA expressao)* FECHAPAR
    ;
cmdRetorne
    : RETORNE expressao
    ;

// Regras auxiliares para definição de seleção do valor para o comando "Caso".
selecao
    : item_selecao*
    ;
item_selecao
    : constantes DOISPONTOS cmd*
    ;
constantes
    : numero_intervalo (VIRGULA numero_intervalo)*
    ;
numero_intervalo
    : op_unario? NUM_INT (INTERVALO op_unario? NUM_INT)?
    ;

// Regras de definição de expressões relacionais.
exp_relacional
  : exp_aritmetica (op_relacional exp_aritmetica)?
  ;
op_relacional
  : IGUAL | DIFERENTE | MAIORIGUAL | MENORIGUAL | MAIOR | MENOR
  ;
expressao
  : termo_logico (op_logico_1 termo_logico)*
  ;
termo_logico
  : fator_logico (op_logico_2 fator_logico)*
  ;
fator_logico
  : NAO? parcela_logica
  ;
parcela_logica
  : ( VERDADEIRO | FALSO )
| exp_relacional
  ;
op_logico_1
  : OU
  ;
op_logico_2
  : E
  ;
