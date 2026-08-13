package main

import "fmt"

func main() {
	// Declaração com tipo explícito
	var variavel1 string = "Variável 1"
	var (
		variavel3 string = "3"
		variavel4 string = "4"
	)

	// Declaração com tipo implícito
	variavel2 := "Variável 2"
	variavel5, variavel6 := "5", "6"

	// Constante
	const constante1 = "Constante1"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)
	fmt.Println(constante1)

	// Troca de valores de variaveis
	variavel1, variavel2 = variavel2, variavel1
	fmt.Println(variavel1)
	fmt.Println(variavel2)
}
