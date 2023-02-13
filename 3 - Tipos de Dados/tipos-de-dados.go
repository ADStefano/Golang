package main

import (
	"errors"
	"fmt"
)

func main()  {

	//Números inteiros

	var numero1 int64 = 100000000
	fmt.Println(numero1)
	
	var numero2 uint32 = 100000
	fmt.Println(numero2)
	
	// Alias

	// RUNE = int32
	var numero3 rune = 200000
	fmt.Println(numero3)

	// uint8 = BYTE
	var numero4 byte = 200
	fmt.Println(numero4)

	// Números reais

	var numero5 float32 = 123.45
	fmt.Println(numero5)

	var numero6 float64 = 123456789.000000
	fmt.Println(numero6)

	// Strings

	var texto string = "TEXTO"
	fmt.Println(texto)

	char := 'A'
	fmt.Println(char)

	// Booleano

	var Booleano bool = true
	fmt.Println(Booleano)

	// Erro
	// String de erro sempre em letras minúsculas
	var erro error = errors.New("exemplo de erro")
	fmt.Println(erro)

}