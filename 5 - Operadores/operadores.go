package main

import "fmt"

func main()  {
	
	// Aritiméticos ||

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 4/2
	multiplicacao := 1*2
	resto := 1 % 2

	fmt.Println(soma,subtracao,divisao,multiplicacao,resto)

	// Atribuição

	var string string = "texto"
	texto := "texto 2"

	fmt.Println(string,texto)

	// Relacionais

	fmt.Println("Menor que < \nMaior que >\nIgual ==\nDiferente !=\nMaior ou igual >=\nMenor ou igual <=")

	// Lógicos

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(! false)

	// Unários

	numero := 10

	numero++
	fmt.Println(numero)

	numero += 1
	fmt.Println(numero)
	
	numero--
	fmt.Println(numero)

	numero -= 1
	fmt.Println(numero)

	numero /=2
	fmt.Println(numero)

	numero *= 2
	fmt.Println(numero)

}