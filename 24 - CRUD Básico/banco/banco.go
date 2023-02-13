package banco

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

// Função que cria a conexão com o banco de dados
func Conn() (*sql.DB, error){
	conn_string := "golang:golang@/teste?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", conn_string)
	if erro != nil{
		return nil, erro
	}

	if erro := db.Ping(); erro != nil{
		return nil, erro
	}

	return db, nil
}