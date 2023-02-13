package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipo_endereco := enderecos.Endereco("Rua Manoel Eufrásio")
	fmt.Println("Tipo do endereço:",tipo_endereco)
}