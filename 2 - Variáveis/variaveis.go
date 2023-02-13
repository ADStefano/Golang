package main

import "fmt"

func main()  {
	
	// Método explicito de declarar uma variável
	var variavel1 string = "Variável 1"

	// Método implícito de declarar uma variável/ Inferência de tipo
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Declarando diversas váriaveis de uma vez com o método explícito
	var(
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	// Declarando diversas variáveis de uma vez com o método explicito/inferência
	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	// Trocando o valor de variáveis sem utilizar uma variável auxiliar
	variavel5, variavel6 = "Variável 6", "Variável 5"

	fmt.Println(variavel5)
	fmt.Println(variavel6)
}