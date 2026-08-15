// Teste unitário
package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Sempre começando com Test e seguindo o nome da função a ser testada (o nome da função não é necessário ser igual mas por padrão é usado)
func TestTipoDeEndereco(t *testing.T) {
	t.Parallel() // Permite que a função seja executada em paralelo
	cenariosDeTeste := []cenarioDeTeste{{"Avenida Paulista", "Avenida"}, {"Rua dos Marcianos", "Rua"}, {"RUA dos Marcianos", "Rua"}, {"", "Tipo inválido"}}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("Tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Teste falhou")
	}
}
