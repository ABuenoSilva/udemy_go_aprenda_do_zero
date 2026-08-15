package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo!", canal)

	/*
		for {
			mensagem, aberto := <-canal
			if !aberto {
				break
			}
			fmt.Println(mensagem)
		}
	*/

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(text string, canal chan string) {
	for j := 0; j < 5; j++ {
		canal <- text
		time.Sleep(time.Second)
	}
	close(canal)
}
