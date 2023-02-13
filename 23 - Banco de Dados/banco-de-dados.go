package main

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/go-sql-driver/mysql"
)

func main(){

	Conn := "golang:golang@/teste?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql",Conn)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão está aberta")

	linhas, err := db.Query("SELECT * FROM ada.usuario")

	if err != nil{
		log.Fatal(err)
	}

	defer linhas.Close()
	fmt.Println(linhas)
}