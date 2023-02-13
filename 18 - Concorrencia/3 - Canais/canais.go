package main

import (
	"fmt"
	"time"
)

func main(){
	canal := make(chan string)

	go escrever("Olá Mundo", canal)

	fmt.Println("Depois da função escrever ser executada!")

	for menssagem := range canal{
		fmt.Println(menssagem)
	}

	fmt.Println("Fim da execução!")

}

func escrever(texto string, canal chan string){
	for i := 0; i < 5; i++{
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)

}