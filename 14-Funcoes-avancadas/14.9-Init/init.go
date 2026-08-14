package main

import "fmt"

var n int

// A função init() sempre é chamada antes de qualquer chamado ao fonte. Pode ter uma por fonte
func init() {
	fmt.Println("Executando init")
	n = 10
}

func main() {
	fmt.Println("Executando main")
	fmt.Println(n)
}
