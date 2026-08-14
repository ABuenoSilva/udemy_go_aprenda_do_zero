package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func main() {
	defer funcao1() // Defer segura até o último momento (final da função atual). Se retornar um valor, sempre é executada imediatamente antes do return
	// Exemplo de uso: encerrar conexão com banco de dados, chama uma vez só com defer e ele sempre vai ser executado no final
	funcao2()
	fmt.Println("Função main")
}
