package main

import (
	"fmt"
	"log"
	"net/http"
	"crud/servidor"
	"github.com/gorilla/mux"
)

func main(){

	router := mux.NewRouter()

	router.HandleFunc("/usuario", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuario", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuario/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Escutando porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}