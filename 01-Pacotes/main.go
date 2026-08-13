package main

import (
	"fmt"

	"go_aprenda_do_zero/01-Pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Teste pacotes - main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("abueno.silva@gmail.com")
	fmt.Println(erro)
}
