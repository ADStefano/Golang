package main

import "fmt"

// A função cria uma cópia da variável, não alterando a original
func inverterSinal(numero int) int {
	
	return numero * -1
}

// Altera a variável no endereço de memória, basicamente muda a variável original
func inverterSinalComPonteiro(numero *int)  {
	*numero = *numero * -1
}

func main()  {
	
	numero := 20
	numeroInvertido := inverterSinal(numero)

	fmt.Println("Número 20 original:",numero)
	fmt.Println("Número 20 invertido:",numeroInvertido)
	fmt.Println("Número 20 original continua igual:",numero)

	novoNumero := 40

	fmt.Println("Número 40 original:",novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println("Novo número 40:",novoNumero)

	
}