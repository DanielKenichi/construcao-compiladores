package errortokens

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func IsTokenAError(tokenName string) bool {
	return strings.Contains(tokenName, "ERRO")
}

func WriteError(tokenName string, token antlr.Token) string {
	switch tokenName {
	case "ERRO":
		return fmt.Sprintf("Linha %s: %s - simbolo nao identificado\n", strconv.Itoa(token.GetLine()), token.GetText())
	case "ERRO_CADEIA_ABERTA":
		return fmt.Sprintf("Linha %s: cadeia literal nao fechada\n", strconv.Itoa(token.GetLine()))
	case "ERRO_COMENTARIO_ABERTO":
		return fmt.Sprintf("Linha %s: comentario nao fechado\n", strconv.Itoa(token.GetLine()))
	case "ERRO_SINTATICO":
		return fmt.Sprintf("Linha %s: erro sintatico proximo a %s\n", strconv.Itoa(token.GetLine()), token.GetText())
	default:
		return fmt.Sprintf("Linha %s: %s - erro nao identificado\n", strconv.Itoa(token.GetLine()), token.GetText())
	}
}
