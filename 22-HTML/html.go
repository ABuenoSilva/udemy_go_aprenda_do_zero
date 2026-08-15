package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	// 1. Faz o parse a CADA requisição (mantenha isso durante o desenvolvimento)
	tmpl, err := template.ParseGlob("*.html")
	if err != nil {
		log.Println("Erro ao carregar templates:", err)
		http.Error(w, "Erro interno ao carregar template", http.StatusInternalServerError)
		return
	}

	u := usuario{"Maria", "joao@x.com"}

	// 2. Renderiza o HTML no navegador
	err = tmpl.ExecuteTemplate(w, "home.html", u)
	if err != nil {
		log.Println("Erro ao executar template:", err)
	}
}

func main() {
	http.HandleFunc("/home", home)

	fmt.Println("Servidor executando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
