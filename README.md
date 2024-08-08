# Construcao-compiladores
Repositorio contendo os trabalhos realizados para a disciplina de construção de compiladores

# Membros do grupo
* Daniel Kenichi Tiago Tateishi - RA: 790837
* Rodrigo Pavão Coffani Nunes - RA: 800345

# Getting Started

Para começar, instale o [Go](https://go.dev/doc/install)

Apos isso, rode o comando abaixo para instalar as dependencias necessarias 

```bash
    go mod tidy
```

Para compilar o programa, no diretorio contendo a main, execute:

```bash
    go build -o <output-dir>
```
Para rodar o programa rode:

```bash
    ./<binary_name> <input_file> <output_file>
```

Para validação, no diretrio do ```corretor``` execute:

*T1

```go build -o t1 ../T1/go/br/ufscar/dc/compiladores/t1/lexico/*.go```

*T2

```go build -o t2 ../T2/go/br/ufscar/dc/compiladores/t2/parser/*.go```

*T3

```go build -o t3 ../T3/go/br/ufscar/dc/compiladores/t3/*.go```

*T4

```go build -o t4 ../T4/go/br/ufscar/dc/compiladores/t4/*.go```


e depois rode o script

```bash
    ./run_tests.sh <tn> (substitua "n" pelo trabalho a ser corrigido) 
```

Para gerar o código do ANTLR, execute, a partir do diretório onde se encontra o arquivo .g4 (dentro da pasta antlr):

*T1

```antlr4 -Dlanguage=Go -o parser T1AlgumaLexer.g4```

*T2

```antlr4 -Dlanguage=Go -o parser T2AlgumaLexer.g4```

*T3

```antlr4 -Dlanguage=Go -visitor -o Alguma T3Alguma.g4```

*T4

```antlr4 -Dlanguage=Go -visitor -o Alguma T4Alguma.g4```