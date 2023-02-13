package main

import "fmt"

// Defer = adiar
func funcao1(){
	fmt.Println("Executando a função 1")
}

func funcao2(){
	fmt.Println("Executando a função 2")
}

func calculoDeMedia(n1 , n2 float32) bool{

	defer fmt.Println("Média calculada")
	fmt.Println("Verificando a média")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}else {
		return false
	}
}

func main(){

	defer funcao1()
	funcao2()
	calculoDeMedia(7,8)
}