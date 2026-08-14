package main

import "fmt"

func main() {
	// Aritméticos: + - / * %
	// Não pode utilizar operadores com tipos diferentes, mesmo que sejam int. Exemplo, não posso somar um int8 com um int16

	fmt.Println("Aritméticos")
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	// Atribuição: := =
	fmt.Println("Atribuição")

	var var1 string = "String"
	var2 := "String2"

	fmt.Println(var1, var2)

	// Operadores relacionais
	fmt.Println("Relacionais")

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Operadores lógicos
	fmt.Println("Lógicos")

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Operadores unários
	fmt.Println("Unários")
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 7
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 4
	fmt.Println(numero)
	numero %= 2
	fmt.Println(numero)

	// Operador ternário não existe, tem que usar if/else
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
