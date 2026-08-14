package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	enderecoExemplo := endereco{"Rua 1", 70}
	var usuario1 usuario
	usuario1.idade = 53
	usuario1.nome = "Alexandre"
	usuario1.endereco = enderecoExemplo
	fmt.Println(usuario1)

	usuario2 := usuario{"Alex", 50, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Zanza"}
	fmt.Println(usuario3)

	/* Não está na aula, fiz sem querer com new() e pesquisei - ponteiro para um struct. Mas o mais utilizado atualmente é o &

	usuario4 := new(usuario)
	fmt.Println(usuario4)

	// 1. Usando new() -> cria um ponteiro *usuario (tudo com zero-value)
	usuario4 := new(usuario)

	// 2. Forma idomática moderna -> TAMBÉM cria um ponteiro *usuario!
	usuario5 := &usuario{}

	// 3. E com & você já pode preencher o que quiser se precisar:
	usuario6 := &usuario{nome: "Alexandre"}

	*/
}
