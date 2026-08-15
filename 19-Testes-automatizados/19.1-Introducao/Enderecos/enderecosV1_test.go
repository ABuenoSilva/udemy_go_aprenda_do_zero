// Teste unitário
package enderecos

import (
	"fmt"
	"slices"
	"testing"
)

// Sempre começando com Test e seguindo o nome da função a ser testada (o nome da função não é necessário ser igual mas por padrão é usado)
func TestTipoDeEnderecoV1(t *testing.T) {
	enderecoParaTeste := []string{"Avenida Paulista", "Rua dos Marcianos", "Rodovia dos Bandeirantes", "Estrada dos Alvarengas"}
	tipoDeEnderecoEsperado := []string{"Avenida", "Rua", "Rodovia", "Estrada"}

	for _, endereco := range enderecoParaTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(endereco)
		if !slices.Contains(tipoDeEnderecoEsperado, tipoDeEnderecoRecebido) {
			fmt.Println(tipoDeEnderecoRecebido)
			t.Error("Tipo recebido é diferente do esperado")
		}
	}
}
