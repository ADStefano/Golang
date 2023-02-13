package main

import "fmt"

// Variável global
var numero int

func main(){

	fmt.Println("Função main")
	fmt.Println(numero)
}

func init()  {
	
	// Atribuindo um valor a variável global
	numero = 10
	fmt.Println("Função init roda antes da função main")
}