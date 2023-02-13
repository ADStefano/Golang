package main

import "fmt"

func recuperação(){

	// Recover recupera a execução do programa
	if r := recover(); r != nil {
		fmt.Println("Você fez a recuperação e não está mais reprovado")
	}
}

func calcularMedia(n1 , n2 float32) bool{

	// Sempre usado com o defer
	defer recuperação()
	media := (n1 + n2) / 2

	if media > 6 {
		fmt.Println("Você passou de ano!")
		return true
	}else if media < 6{
		fmt.Println("Você está reprovado")
		return false
	}

	// Panic para o programa e procura por qualquer função que tenha o defer 
	panic("Você está de recuperação")
}

func main()  {
	
	calcularMedia(6,6)
	fmt.Println("Fim de ano")
}