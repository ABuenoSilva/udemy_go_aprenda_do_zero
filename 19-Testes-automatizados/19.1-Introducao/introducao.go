package main

import (
	"fmt"

	enderecos "go_aprenda_do_zero/19-Testes-automatizados/19.1-Introducao/Enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia Paulista")
	fmt.Println(tipoEndereco)
}
