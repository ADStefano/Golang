package main

import "fmt"

func main(){
	canal := make(chan string, 2) // Tamanho do canal, da erro de deadlock se ultrapassar a capacidade máxima

	canal <- "Olá Mundo"
	canal <- "Programando em Go"

	menssagem := <- canal
	menssagem2 := <- canal
	fmt.Println(menssagem)
	fmt.Println(menssagem2)
}