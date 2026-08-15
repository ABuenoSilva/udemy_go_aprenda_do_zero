package main

import "fmt"

func main() {
	// Canais com buffer só dão deadlocks se estourar a capacidade no envio ou no retorno do canal
	canal := make(chan string, 2)
	canal <- "Olá, mundo!"
	canal <- "Segunda msg"

	mensagem := <-canal
	fmt.Println(mensagem)
	msg2 := <-canal
	fmt.Println(msg2)
}
