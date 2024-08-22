package generator

import (
	"strconv"
	"strings"

	parser "github.com/DanielKenichi/construcao-compiladores/T5/antlr4/br/ufscar/dc/compiladores/t5/parser/Alguma"
	"github.com/DanielKenichi/construcao-compiladores/T5/go/br/ufscar/dc/compiladores/t5/scope"
	"github.com/DanielKenichi/construcao-compiladores/T5/go/br/ufscar/dc/compiladores/t5/symboltable"
)

/*
	Implementacao do nosso visitor para analise semântica.
	O visitor vai acumulando os erros em um slice de strings
*/

type AlgumaGenerator struct {
	parser.BaseT5AlgumaVisitor
	Scopes *scope.Scope
}

func New() *AlgumaGenerator {
	return &AlgumaGenerator{
		Scopes: scope.New(),
	}
}

func (g *AlgumaGenerator) VisitPrograma(ctx parser.IProgramaContext) []string {
	var programaResult = make([]string, 0)

	result := []string{"#include <stdio.h>\n", "#include <stdlib.h>\n", "\n"}
	programaResult = append(programaResult, result...)

	result = g.VisitDeclaracoes(ctx.Declaracoes())
	programaResult = append(programaResult, result...)
	programaResult = append(programaResult, "\n")

	result = g.VisitCorpo(ctx.Corpo())
	programaResult = append(programaResult, result...)

	return programaResult

}

func (g *AlgumaGenerator) VisitDeclaracoes(ctx parser.IDeclaracoesContext) []string {
	g.Scopes.NewScope()

	declaracoesResult := make([]string, 0)

	result := g.VisitDeclaracoes_variaveis(ctx.AllDeclaracoes_variaveis())
	declaracoesResult = append(declaracoesResult, result...)

	// result = g.VisitDeclaracoes_funcoes(ctx.AllDeclaracoes_funcoes())
	// declaracoesResult = append(declaracoesResult, result...)

	return declaracoesResult
}

func (g *AlgumaGenerator) VisitCorpo(ctx parser.ICorpoContext) []string {
	g.Scopes.NewScope()
	corpoResult := make([]string, 0)

	result := []string{"int main() {\n"}
	corpoResult = append(corpoResult, result...)

	result = g.VisitDeclaracoes_variaveis(ctx.AllDeclaracoes_variaveis())
	corpoResult = append(corpoResult, result...)

	result = g.VisitCmd(ctx.AllCmd())
	corpoResult = append(corpoResult, result...)

	result = []string{"\treturn 0;\n}\n"}
	corpoResult = append(corpoResult, result...)

	return corpoResult
}

func (g *AlgumaGenerator) VisitCmd(ctxs []parser.ICmdContext) []string {
	cmdResult := make([]string, 0)
	result := make([]string, 0)

	for _, ctx := range ctxs {
		result = []string{"\t"}
		if ctx.CmdLeia() != nil {
			result = append(result, g.VisitCmdLeia(ctx.CmdLeia())...)
			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdEscreva() != nil {
			result = append(result, g.VisitCmdEscreva(ctx.CmdEscreva())...)
			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdSe() != nil {
			result = append(result, g.VisitCmdSe(ctx.CmdSe())...)
			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdCaso() != nil {
			result = append(result, g.VisitCmdCaso(ctx.CmdCaso())...)
			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdPara() != nil {
			result = append(result, g.VisitCmdPara(ctx.CmdPara())...)
			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdEnquanto() != nil {
			result = append(result, g.VisitCmdEnquanto(ctx.CmdEnquanto())...)
			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdFaca() != nil {
			result = append(result, g.VisitCmdFaca(ctx.CmdFaca())...)
			cmdResult = append(cmdResult, result...)

			continue
		}

		if ctx.CmdAtribuicao() != nil {
			result = append(result, g.VisitCmdAtribuicao(ctx.CmdAtribuicao())...)
			cmdResult = append(cmdResult, result...)

			continue
		}

		// if ctx.CmdChamada() != nil {
		// 	result := g.VisitCmdChamada(ctx.CmdChamada())
		// 	cmdResult = append(cmdResult, result...)

		// 	continue
		// }

		// if ctx.CmdRetorne() != nil {
		// 	result := g.VisitCmdRetorne(ctx.CmdRetorne())
		// 	cmdResult = append(cmdResult, result...)

		// 	continue
		// }
	}

	return cmdResult
}

func (g *AlgumaGenerator) VisitCmdLeia(ctx parser.ICmdLeiaContext) []string {
	cmdLeiaResult := make([]string, 0)
	result := make([]string, 0)
	readFunction := ""

	for _, ident := range ctx.AllIdentificador() {
		tipo := g.GetIdentifierType(ident)
		if tipo == symboltable.LITERAL {
			readFunction = "gets("
			result = []string{readFunction, ident.GetText()}
		} else {
			readFunction = "scanf("
			operator := "\"" + g.MapTypeToOperatorC(tipo) + "\"" + ", &"
			result = []string{readFunction, operator, ident.GetText()}
		}
		cmdLeiaResult = append(cmdLeiaResult, result...)
	}

	result = []string{");\n"}
	cmdLeiaResult = append(cmdLeiaResult, result...)

	return cmdLeiaResult
}

func (g *AlgumaGenerator) VisitCmdEscreva(ctx parser.ICmdEscrevaContext) []string {
	cmdEscrevaResult := make([]string, 0)
	result := make([]string, 0)

	for _, expressao := range ctx.AllExpressao() {
		expressionType := g.GetExpressaoType(expressao)
		operator := g.MapTypeToOperatorC(expressionType)
		result = []string{"printf(\""}

		if strings.Contains(expressao.GetText(), "\"") {
			result = append(result, strings.ReplaceAll(expressao.GetText(), "\"", ""))
			result = append(result, "\"")
		} else {
			result = append(result, operator)
			result = append(result, "\", ")
			result = append(result, g.VisitExpressao(expressao)...)
		}

		result = append(result, ");\n")
		cmdEscrevaResult = append(cmdEscrevaResult, result...)
	}

	return cmdEscrevaResult
}

func (g *AlgumaGenerator) VisitCmdSe(ctx parser.ICmdSeContext) []string {
	cmdSeResult := make([]string, 0)

	result := []string{"if ("}
	cmdSeResult = append(cmdSeResult, result...)

	result = g.VisitExpressao(ctx.Expressao())
	cmdSeResult = append(cmdSeResult, result...)

	result = []string{") {\n"}
	cmdSeResult = append(cmdSeResult, result...)

	result = g.VisitCmd(ctx.GetSeCmds())
	cmdSeResult = append(cmdSeResult, result...)

	result = []string{"\t}\n"}
	cmdSeResult = append(cmdSeResult, result...)

	if ctx.SENAO() != nil {
		result = []string{"\telse {\n"}
		cmdSeResult = append(cmdSeResult, result...)

		result = g.VisitCmd(ctx.GetSenaoCmds())
		cmdSeResult = append(cmdSeResult, result...)

		result = []string{"\t}\n"}
		cmdSeResult = append(cmdSeResult, result...)
	}

	return cmdSeResult
}

func (g *AlgumaGenerator) VisitCmdCaso(ctx parser.ICmdCasoContext) []string {
	cmdCasoResult := make([]string, 0)

	result := []string{"switch ("}
	cmdCasoResult = append(cmdCasoResult, result...)

	result = g.VisitExp_aritmetica(ctx.Exp_aritmetica())
	cmdCasoResult = append(cmdCasoResult, result...)

	result = []string{") {\n"}
	cmdCasoResult = append(cmdCasoResult, result...)

	// casos
	result = g.VisitSelecao(ctx.Selecao())
	cmdCasoResult = append(cmdCasoResult, result...)

	// senao
	if ctx.SENAO() != nil {
		result = []string{"\tdefault:\n"}
		cmdCasoResult = append(cmdCasoResult, result...)

		cmdCasoResult = append(cmdCasoResult, "\t")
		result = g.VisitCmd(ctx.AllCmd())
		cmdCasoResult = append(cmdCasoResult, result...)
	}

	result = []string{"\t}\n"}
	cmdCasoResult = append(cmdCasoResult, result...)

	return cmdCasoResult
}

func (g *AlgumaGenerator) VisitCmdPara(ctx parser.ICmdParaContext) []string {
	cmdParaResult := make([]string, 0)
	result := make([]string, 0)

	result = append(result, "for (")

	result = append(result, ctx.IDENT().GetText())
	result = append(result, " = ")
	result = append(result, g.VisitExp_aritmetica(ctx.GetExp1())...)
	result = append(result, "; ")
	result = append(result, ctx.IDENT().GetText())
	result = append(result, " <= ")
	result = append(result, g.VisitExp_aritmetica(ctx.GetExp2())...)
	result = append(result, "; ")
	result = append(result, ctx.IDENT().GetText())
	result = append(result, "++) {\n")
	cmdParaResult = append(cmdParaResult, result...)

	result = []string{"\t"}
	result = append(result, g.VisitCmd(ctx.AllCmd())...)

	result = append(result, "\t}\n")

	cmdParaResult = append(cmdParaResult, result...)
	return cmdParaResult
}

func (g *AlgumaGenerator) VisitCmdEnquanto(ctx parser.ICmdEnquantoContext) []string {
	cmdEnquantoResult := make([]string, 0)
	result := make([]string, 0)

	result = append(result, "while (")
	result = append(result, g.VisitExpressao(ctx.Expressao())...)
	result = append(result, ") {\n")

	result = append(result, "\t")
	result = append(result, g.VisitCmd(ctx.AllCmd())...)
	result = append(result, "\t}\n")

	cmdEnquantoResult = append(cmdEnquantoResult, result...)
	return cmdEnquantoResult
}

func (g *AlgumaGenerator) VisitCmdFaca(ctx parser.ICmdFacaContext) []string {
	cmdFacaResult := make([]string, 0)
	result := make([]string, 0)

	result = append(result, "do {\n")
	result = append(result, "\t")

	result = append(result, g.VisitCmd(ctx.AllCmd())...)
	result = append(result, "} while (")

	result = append(result, g.VisitExpressao(ctx.Expressao())...)

	result = append(result, ");\n")

	cmdFacaResult = append(cmdFacaResult, result...)

	return cmdFacaResult
}

func (g *AlgumaGenerator) VisitCmdAtribuicao(ctx parser.ICmdAtribuicaoContext) []string {
	cmdAtibuicaoResult := make([]string, 0)
	result := make([]string, 0)

	result = append(result, ctx.Identificador().GetText(), " = ")
	cmdAtibuicaoResult = append(cmdAtibuicaoResult, result...)

	result = g.VisitExpressao(ctx.Expressao())
	cmdAtibuicaoResult = append(cmdAtibuicaoResult, result...)

	// result = g.VerifyVarAttribution(ctx)
	// cmdAtibuicaoResult = append(cmdAtibuicaoResult, result...)

	cmdAtibuicaoResult = append(cmdAtibuicaoResult, ";\n")
	return cmdAtibuicaoResult
}

// func (g *AlgumaGenerator) VisitCmdChamada(ctx parser.ICmdChamadaContext) []string {
// 	cmdChamadaResult := make([]string, 0)

// 	ident := ctx.IDENT()

// 	result := g.VerifyFuncCall(ident, ctx.AllExpressao())

// 	cmdChamadaResult = append(cmdChamadaResult, result...)

// 	return cmdChamadaResult
// }

// func (g *AlgumaGenerator) VisitCmdRetorne(ctx parser.ICmdRetorneContext) []string {
// 	cmdChamadaResult := make([]string, 0)

// 	return cmdChamadaResult
// }

func (g *AlgumaGenerator) VisitSelecao(ctx parser.ISelecaoContext) []string {
	selecaoResult := make([]string, 0)

	for _, itemSelecao := range ctx.AllItem_selecao() {
		result := g.VisitItem_selecao(itemSelecao)
		selecaoResult = append(selecaoResult, result...)
	}

	return selecaoResult
}

func (g *AlgumaGenerator) VisitItem_selecao(ctx parser.IItem_selecaoContext) []string {
	itemSelecaoResult := make([]string, 0)

	result := g.VisitConstantes(ctx.Constantes())
	itemSelecaoResult = append(itemSelecaoResult, result...)

	result = g.VisitCmd(ctx.AllCmd())
	itemSelecaoResult = append(itemSelecaoResult, result...)

	result = []string{"\tbreak;\n"}
	itemSelecaoResult = append(itemSelecaoResult, result...)

	return itemSelecaoResult
}

func (g *AlgumaGenerator) VisitConstantes(ctx parser.IConstantesContext) []string {
	constantesResult := make([]string, 0)

	result := g.VisitNumero_intervalo(ctx.Numero_intervalo(0))

	constantesResult = append(constantesResult, result...)

	return constantesResult
}

func (g *AlgumaGenerator) VisitNumero_intervalo(ctx parser.INumero_intervaloContext) []string {
	numeroIntervaloResult := make([]string, 0)
	startIdx, endIdx := 0, 0
	step := 1

	// if ctx.Op_unario(0) != nil {
	// 	result := []string{ctx.Op_unario(0).GetText()}
	// 	numeroIntervaloResult = append(numeroIntervaloResult, result...)
	// }

	if ctx.NUM_INT(0) != nil {
		startIdx, _ = strconv.Atoi(ctx.NUM_INT(0).GetText())
		// numeroIntervaloResult = append(numeroIntervaloResult, result...)
	}

	if ctx.INTERVALO() != nil {
		// if ctx.Op_unario(1) != nil {
		// 	result := []string{ctx.Op_unario(1).GetText()}
		// 	numeroIntervaloResult = append(numeroIntervaloResult, result...)
		// }

		endIdx, _ = strconv.Atoi(ctx.NUM_INT(1).GetText())

		for i := startIdx; i <= endIdx; i += step {
			result := []string{"\tcase ", strconv.Itoa(i), ":\n"}
			numeroIntervaloResult = append(numeroIntervaloResult, result...)
		}
	}
	if endIdx == 0 {
		result := []string{"\tcase ", strconv.Itoa(startIdx), ":\n"}
		numeroIntervaloResult = append(numeroIntervaloResult, result...)
	}

	return numeroIntervaloResult
}

func (g *AlgumaGenerator) VisitDeclaracoes_variaveis(ctxs []parser.IDeclaracoes_variaveisContext) []string {
	variaveisResult := make([]string, 0)

	for _, ctx := range ctxs {
		if ctx.DECLARE() != nil {
			result := []string{"\t"}
			result = append(result, g.VisitVariavel(ctx.Variavel())...)
			variaveisResult = append(variaveisResult, result...)

			g.AddVarToSymbolTable(ctx.Variavel(), g.Scopes.CurrentScope())

			continue
		}

		if ctx.CONSTANTE() != nil {
			result := []string{"#define ", ctx.IDENT().GetText(), " "}

			// result = append(result, g.VisitTipo_basico(ctx.Tipo_basico())...)
			// variaveisResult = append(variaveisResult, result...)

			// result = g.VisitValor_constante(ctx.Valor_constante())
			result = append(result, ctx.Valor_constante().GetText(), "\n")
			variaveisResult = append(variaveisResult, result...)

			g.AddConstToSymbolTable(ctx.IDENT(), ctx.Tipo_basico())

			continue
		}

		// if ctx.TIPO() != nil {
		// 	result := g.VisitRegistro(ctx.Registro())
		// 	variaveisResult = append(variaveisResult, result...)

		// 	result = g.AddRegTypeToSymbolTable(ctx.IDENT(), ctx.Registro())
		// 	variaveisResult = append(variaveisResult, result...)

		// 	continue
		// }
	}

	return variaveisResult
}

// func (g *AlgumaGenerator) VisitDeclaracoes_funcoes(ctxs []parser.IDeclaracoes_funcoesContext) []string {
// 	//TODO conferir tipo de retorno com o escopo

// 	funcoesResult := make([]string, 0)
// 	for _, ctx := range ctxs {
// 		result := g.AddFuncToSymbolTable(ctx)
// 		funcoesResult = append(funcoesResult, result...)

// 		symbol := g.GetFuncVarSymbol(ctx.IDENT().GetText())

// 		if symbol.SymbolType == symboltable.PROCEDIMENTO {
// 			for _, cmd := range ctx.AllCmd() {
// 				if cmd.CmdRetorne() != nil {
// 					funcoesResult = append(funcoesResult,
// 						semanticError(cmd.CmdRetorne().RETORNE().GetSymbol(), "comando retorne nao permitido nesse escopo"))
// 				}
// 			}
// 		}

// 	}

// 	return funcoesResult
// }

func (g *AlgumaGenerator) VisitVariavel(ctx parser.IVariavelContext) []string {
	variavelResult := make([]string, 0)

	result := g.VisitTipo(ctx.Tipo())
	variavelResult = append(variavelResult, result...)
	strSize := "[80]"

	result = g.VisitIdentificador(ctx.Identificador(0))
	variavelResult = append(variavelResult, result...)

	aux := 0
	for ctx.VIRGULA(aux) != nil {
		aux++
		result = []string{", "}
		result = append(result, g.VisitIdentificador(ctx.Identificador(aux))...)
		if ctx.Tipo().Tipo_variavel().Tipo_basico().LITERAL() != nil {
			result = append(result, strSize)
		}
		variavelResult = append(variavelResult, result...)
	}

	if ctx.Tipo().Tipo_variavel().Tipo_basico().LITERAL() != nil {
		result = []string{"[80]"}
		variavelResult = append(variavelResult, result...)
	}

	result = []string{";\n"}
	variavelResult = append(variavelResult, result...)

	return variavelResult
}

func (g *AlgumaGenerator) VisitExpressao(ctx parser.IExpressaoContext) []string {
	expressaoResult := make([]string, 0)

	result := g.VerifyExpression(ctx)
	expressaoResult = append(expressaoResult, result...)

	return expressaoResult
}

func (g *AlgumaGenerator) VisitTipo_basico(ctx parser.ITipo_basicoContext) []string {
	tipoBasicoResult := make([]string, 0)

	result := g.MapTypeToTypeC(ctx) + " "
	tipoBasicoResult = append(tipoBasicoResult, result)

	return tipoBasicoResult
}

// func (g *AlgumaGenerator) VisitValor_constante(ctx parser.IValor_constanteContext) []string {
// 	valorConstanteResult := make([]string, 0)

// 	return valorConstanteResult
// }

// func (g *AlgumaGenerator) VisitRegistro(ctx parser.IRegistroContext) []string {
// 	registroResult := make([]string, 0)

// 	return registroResult
// }

func (g *AlgumaGenerator) VisitIdentificador(ctx parser.IIdentificadorContext) []string {
	identificadorResult := make([]string, 0)

	result := ctx.IDENT(0).GetText()
	identificadorResult = append(identificadorResult, result)

	// aux := 0

	// for ctx.ABRECHAVE(aux) != nil {
	// 	result := g.VisitExp_aritmetica(ctx.Exp_aritmetica(aux))
	// 	identificadorResult = append(identificadorResult, result...)
	// 	aux++
	// }

	return identificadorResult
}

func (g *AlgumaGenerator) VisitTipo(ctx parser.ITipoContext) []string {
	tipoResult := make([]string, 0)

	// if ctx.Registro() != nil {
	// 	result := g.VisitRegistro(ctx.Registro())
	// 	tipoResult = append(tipoResult, result...)
	// }

	if ctx.Tipo_variavel() != nil {
		result := g.VisitTipo_variavel(ctx.Tipo_variavel())
		tipoResult = append(tipoResult, result...)
	}

	return tipoResult
}

func (g *AlgumaGenerator) VisitTipo_variavel(ctx parser.ITipo_variavelContext) []string {
	tipoVariavelResult := make([]string, 0)
	result := make([]string, 0)

	if ctx.Tipo_basico() != nil {
		result = append(result, g.VisitTipo_basico(ctx.Tipo_basico())...)
		tipoVariavelResult = append(tipoVariavelResult, result...)
	}

	if ctx.PONTEIRO() != nil {
		tipoVariavelResult = append(tipoVariavelResult, "*")
	}

	return tipoVariavelResult
}

func (g *AlgumaGenerator) VisitExp_aritmetica(ctx parser.IExp_aritmeticaContext) []string {
	expResult := make([]string, 0)

	result := g.VisitTermo(ctx.Termo(0))
	expResult = append(expResult, result...)

	aux := 0
	for ctx.Op1(aux) != nil {
		result = g.VisitOp1(ctx.Op1(aux))
		expResult = append(expResult, result...)

		aux++

		result = g.VisitTermo(ctx.Termo(aux))
		expResult = append(expResult, result...)
	}

	return expResult
}

func (g *AlgumaGenerator) VisitTermo(ctx parser.ITermoContext) []string {
	termoResult := make([]string, 0)

	result := g.VisitFator(ctx.Fator(0))
	termoResult = append(termoResult, result...)

	aux := 0
	for ctx.Op2(aux) != nil {
		result = g.VisitOp2(ctx.Op2(aux))
		termoResult = append(termoResult, result...)

		aux++

		result = g.VisitFator(ctx.Fator(aux))
		termoResult = append(termoResult, result...)
	}

	return termoResult
}

func (g *AlgumaGenerator) VisitOp1(ctx parser.IOp1Context) []string {
	op1Result := make([]string, 0)

	return op1Result
}

func (g *AlgumaGenerator) VisitFator(ctx parser.IFatorContext) []string {
	fatorResult := make([]string, 0)

	result := g.VerifyParcela(ctx.Parcela(0))
	fatorResult = append(fatorResult, result...)

	aux := 0
	for ctx.Op3(aux) != nil {
		result = []string{ctx.Op3(aux).GetText()}
		fatorResult = append(fatorResult, result...)

		aux++

		result = g.VerifyParcela(ctx.Parcela(aux))
		fatorResult = append(fatorResult, result...)
	}

	return fatorResult
}

func (g *AlgumaGenerator) VisitOp2(ctx parser.IOp2Context) []string {
	op2Result := make([]string, 0)

	return op2Result
}
