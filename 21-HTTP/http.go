package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Olá Mundo!"))
	if err != nil {
		log.Println("Erro ao escrever resposta:", err)
	}
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Página de usuários!"))
	if err != nil {
		log.Println("Erro ao escrever resposta:", err)
	}
}

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	fmt.Println("Servidor executando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
