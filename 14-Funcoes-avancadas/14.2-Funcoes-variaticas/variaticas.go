package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	fmt.Println(texto, numeros)
}

func main() {
	fmt.Println("Soma(1,2)", soma(1, 2))
	fmt.Println("Soma(1,3,5,7)", soma(1, 3, 5, 7))
	escrever("Texto", 1, 4, 65, 7, 3)
}
