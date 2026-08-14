package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":  {"primeiro": "João", "ultimo": "Pedro"},
		"curso": {"nome": "Engenharia", "campus": "São Paulo", "horario": "noturno"},
	}
	fmt.Println(usuario2)

	delete(usuario2, "curso")
	fmt.Println(usuario2)

	delete(usuario2["nome"], "ultimo")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{"nome": "Câncer"}
	fmt.Println(usuario2)
}
