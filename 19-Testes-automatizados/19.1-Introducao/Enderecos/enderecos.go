package enderecos

import (
	"slices"
	"strings"
	"unicode"
)

// TipoDeEndereco verifica se o endereco tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavra := strings.Split(strings.ToLower(endereco), " ")[0]

	if slices.Contains(tiposValidos, primeiraPalavra) {
		return capitalize(primeiraPalavra)
	}

	return "Tipo inválido"
}

func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
