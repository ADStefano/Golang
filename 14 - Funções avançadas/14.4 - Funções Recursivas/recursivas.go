package main

import "fmt"

func fibonacci(posicao uint8) uint8 {

	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}

func main(){

	fmt.Println("Função Recursiva")

	fmt.Println("Sequência de Fibonacci:")

	var posicao uint8 = 10 // uint8 pra não correr o risco de stack overflow

	for i := uint8(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}


}