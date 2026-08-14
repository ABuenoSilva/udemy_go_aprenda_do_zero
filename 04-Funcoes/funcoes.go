package main

import "fmt"

// Função é um tipo em Go

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Retornos múltiplos
func calculosMatematicos(n1, n2 int32) (int32, int32) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 5)
	fmt.Println(soma)

	f := func(txt string) {
		fmt.Println(txt)
	}

	f("Texto da função 1")

	soma1, sub1 := calculosMatematicos(10, 20)
	fmt.Println(soma1, sub1)

	// Retirando um dos retornos que não quero tratar
	soma2, _ := calculosMatematicos(20, 15)
	fmt.Println(soma2)
}
