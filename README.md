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


e depois rode o script

```bash
    ./run_tests.sh <tn> (substitua "n" pelo trabalho a ser corrigido) 
```
