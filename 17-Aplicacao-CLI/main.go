package main

import (
	"fmt"
	"log"
	"os"

	"go_aprenda_do_zero/17-Aplicacao-CLI/app"
)

func main() {
	fmt.Println("Iniciando...")
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
