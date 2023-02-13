package main

import "fmt"

// Struct é algo muito próximo de uma classe
type usuario struct{

	nome string
	idade uint
	endereco endereco
}

type endereco struct{

	logradouro string
	numero uint8
}

func main()  {
	fmt.Println("Exemplo de Structs")

	var user1 usuario
	user1.nome = "Ângelo"
	user1.idade = 23
	fmt.Println(user1)

	enderecoExemplo := endereco{"Rua do Biscoito", 01}

	user2 := usuario{"Biscoito", 5, enderecoExemplo}
	fmt.Println(user2)

	user3 := usuario{nome: "Bolacha"}
	fmt.Println(user3)
}