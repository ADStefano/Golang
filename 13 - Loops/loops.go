package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("Loops")

	i := 0

	for i < 10{
		time.Sleep(time.Second)
		fmt.Println("Incrementando i:", i)
		i++
	}

	fmt.Println("Valor de i:",i)

	// Método alternativo
	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println("Incrementando j:", j)
	}

	nomes := []string {"Ângelo", "Henrique", "Wildson"}

	for indice, nome := range(nomes){
		fmt.Println("Índice:", indice, "Nome:", nome)
	}

	// Método ignorando uma variável usando underline
	for _, letra := range "PALAVRA"{
		fmt.Println(letra)
	}

	// Transformando o número da tabela ascii em string
	for indice, letra := range "PALAVRA"{
		fmt.Println("Índice:", indice, "Letra:", string(letra))
	}
}