# Construcao-compiladores
Repositorio contendo os trabalhos realizados para a disciplina de construção de compiladores

# Membros do grupo
* Daniel Kenichi Tiago Tateishi - RA: 790837
* Rodrigo Coffani - RA: 800345

# Getting Started

Para começar, instale o [Go](https://go.dev/doc/install)

Apos isso, rode o comando abaixo para instalar as dependencias necessarias 

```bash
    go mod tidy
```

Para compilar o programa, no diretorio ```./T1/go/br/ufscar/dc/compiladores/t1/lexico``` execute:

```bash
    go build -o <output-dir>
```
testes
Para rodar o programa rode:

```bash
    ./<binary_name> <input_file> <output_file>
```

Para rodar os testes, no diretorio do ```corretor``` execute:

>Note: No momento, para o T1 o script de teste espera que exista o binario do compilador com o nome t1-lexico dentro do diretorio do corretor

```bash
    ./run_tests.sh
```
