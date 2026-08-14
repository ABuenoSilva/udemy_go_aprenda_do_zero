package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Array

	fmt.Println("ARRAYS")

	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 7
	array1[1] = 10
	array1[3] = -10
	fmt.Println(array1)

	array2 := [5]string{"Um", "Dois", "Três", "Quatro", "Cinco"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4}
	fmt.Println(array3)

	// Slice

	fmt.Println("SLICES")

	slice := []int{1, 2}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 7)
	fmt.Println(slice)

	for x := range slice {
		fmt.Println(slice[x])
	}

	// Criando a partir de um array. Primeira posição é inclusiva e a segunda é exclusiva
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// O slice é como um ponteiro, então usa referência. Mudando o array de origem também muda o slice
	array2[1] = "Mudei índice 1"
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("--------------")
	slice3 := make([]int, 3, 4) // Terceiro parâmetro (capacidade) é opcional, se omitida ela fica igual ao tamanho (segundo parâmetro)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Tamanho (length)
	fmt.Println(cap(slice3)) // Capacidade

	slice3 = append(slice3, 7)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Tamanho (length)
	fmt.Println(cap(slice3)) // Capacidade

	slice3 = append(slice3, 11) // Quando você estoura a capacidade, o Go referencia um novo array interno com o dobro da capacidade estourada (no caso aqui 8)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Tamanho (length)
	fmt.Println(cap(slice3)) // Capacidade
}
