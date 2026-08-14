package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	x, y := calculosMatematicos(20, 10)
	fmt.Println(x, y)
	fmt.Println(calculosMatematicos(1, 2))
}
