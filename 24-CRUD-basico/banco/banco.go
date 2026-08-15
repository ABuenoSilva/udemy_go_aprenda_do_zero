package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	stringConexao := "devuser:devpassword@tcp(localhost:3306)/debook?charset=utf8&parseTime=true&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
