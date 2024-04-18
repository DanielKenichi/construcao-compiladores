package errortokens

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Verificando se o token eh um erro
func IsTokenAError(tokenName string) bool {
	return strings.Contains(tokenName, "ERRO")
}

// Escrita de mensagem de erro
func WriteError(tokenName string, token antlr.Token) string {
	switch tokenName {
	case "ERRO":
		return fmt.Sprintf("Linha %s: %s - simbolo nao identificado\n", strconv.Itoa(token.GetLine()), token.GetText())
	case "ERRO_CADEIA_ABERTA":
		return fmt.Sprintf("Linha %s: cadeia literal nao fechada\n", strconv.Itoa(token.GetLine()))
	case "ERRO_COMENTARIO_ABERTO":
		return fmt.Sprintf("Linha %s: comentario nao fechado\n", strconv.Itoa(token.GetLine()))
	default:
		return fmt.Sprintf("Linha %s: %s - erro nao identificado\n", strconv.Itoa(token.GetLine()), token.GetText())
	}
}
