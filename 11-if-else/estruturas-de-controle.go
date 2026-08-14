package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// if init limita a variável ao escopo do if
	if outroNumero := -1; outroNumero > 0 {
		fmt.Println("Maior que zero")
	} else if outroNumero == 0 {
		fmt.Println("Igual a zero")
	} else {
		fmt.Println("Menor que zero")
	}
}
