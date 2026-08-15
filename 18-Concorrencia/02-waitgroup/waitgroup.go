package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 3!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 4!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(text string) {
	for j := 0; j < 5; j++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
