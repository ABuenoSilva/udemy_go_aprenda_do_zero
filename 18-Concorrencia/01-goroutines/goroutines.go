package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo!") // chamando com go gera uma goroutine
	escrever("Programando em Go!")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
