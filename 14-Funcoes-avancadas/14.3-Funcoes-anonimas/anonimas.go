package main

import "fmt"

func main() {
	func() {
		fmt.Println("Olá mundo!")
	}()
	func(texto string) {
		fmt.Println(texto)
	}("Parâmetro")
	retorno := func(numero int) int {
		ret := numero * 2
		return ret
	}(10)
	fmt.Println(retorno)
}
