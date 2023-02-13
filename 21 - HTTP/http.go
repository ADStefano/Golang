package main

import (
	"log"
	"net/http"
)

func paginaInicial(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Olá Mundo"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}



func main(){
	/*	HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB
		CLIENTE(FAZ A REQUISIÇÃO[REQUEST]) - SERVIDOR(PROCESSA A REQUISIÇÃO E ENVIA RESPOSTA[RESPONSE])
		ROTAS SÃO UMA MANEIRA DE IDENTIFICAR O TIPO DE MENSAGEM E O TIPO DE PROCESSAMENTO DA MENSAGEM 
		URI - Identificador do Recurso
		MÉTODODS - POST, GET, PUT, DELETE
	*/ 

	http.HandleFunc("/", paginaInicial)

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}