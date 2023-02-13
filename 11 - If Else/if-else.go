package main

import "fmt"

func main()  {

	fmt.Println("Estruturas de controle")
	
	numero := 10

	if numero > 10{
		fmt.Println("O número é maior que 10")
	} else if numero < 0 {
		fmt.Println("O número é menor que 0")
	} else {
		fmt.Println("O número está entre 0 e 10")
	}

	// If init, permite declarar uma variável no if, a variável só existe no escopo do if
	if outronumero := numero; outronumero == 10 {
		fmt.Println("O número é igual a 10")
	} else {
		fmt.Println("O número é diferente de 10")
	}
}