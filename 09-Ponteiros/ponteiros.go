package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	ponteiro1 := &variavel1
	var ponteiro2 *int = &variavel2

	fmt.Println(variavel1, variavel2, ponteiro1, *ponteiro1, ponteiro2, *ponteiro2)

	variavel1++
	variavel2 += 7

	fmt.Println(variavel1, variavel2, ponteiro1, *ponteiro1, ponteiro2, *ponteiro2)
}
