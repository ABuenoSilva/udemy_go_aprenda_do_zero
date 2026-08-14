package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Administração", "PUC"}
	fmt.Println(e1)
	e1.nome = "Pedro"
	fmt.Println(e1)
	fmt.Println(p1)
}
