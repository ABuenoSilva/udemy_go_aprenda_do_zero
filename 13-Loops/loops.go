package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	i := 0
	fmt.Println("Var i")
	for i < 5 {
		time.Sleep(time.Second)
		i++
		fmt.Println(i)
	}
	fmt.Println("Fora do loop " + strconv.Itoa(i))

	fmt.Println("Var j")
	for j := 0; j < 10; j += 2 {
		time.Sleep(time.Second)
		fmt.Println(j)
	}

	fmt.Println("Range")
	nomes := [3]string{"Bueno", "Zanza", "Felipe"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("String")

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra, string(letra))
	}

	fmt.Println("Map")

	usuario := map[string]string{
		"nome":      "Alexandre",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
