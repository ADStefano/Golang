package main

import "fmt"


// Função variática com 1 parâmetro e com return
func soma(numeros ... int) int{
	total := 0

	for _, numero := range numeros{
		total += numero
	}

	return total
}

// Função variática com mais de um parâmetro e sem return, a função variática tem que ser o último parâmetro
func escrever(texto string, numeros ... int){
	for _, numero := range numeros{
		fmt.Println(texto, numero)
	}
}

func main()  {
	
	soma := soma(1,2,3,4)
	fmt.Println(soma)

	escrever("Olá mundo :)", 1, 2, 3, 4)
	
}