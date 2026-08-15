package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import implícito pq não é esse arquivo que vai usar o driver sql - no caso é o pacote sql no sql.Open
)

func main() {
	// Formato: "usuario:senha@tcp(host:porta)/nome_do_banco"
	stringConexao := "devuser:devpassword@tcp(localhost:3306)/debook?charset=utf8&parseTime=true&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal("Erro ao pingar o banco de dados:", erro)
	}
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	linhas, erro := db.Query("Select * from usuarios")
	if erro != nil {
		log.Fatal("Erro ao executar a query:", erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)
}
