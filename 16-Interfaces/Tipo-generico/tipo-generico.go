package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(12)
	generica("Bueno")
	generica(false)
	generica(12.8)
	generica([3]int{1, 2, 3})
}
