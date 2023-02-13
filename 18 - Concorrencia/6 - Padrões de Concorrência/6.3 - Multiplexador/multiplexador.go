// Multiplexador junta vários canais em um só

package main

import (
	"fmt"
	"time"
)

func main(){

	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 2; i++{
		fmt.Println(<- canal)
	}

}

func multiplexar(canal_de_entrada_1, canal_de_entrada_2 <- chan string) <- chan string{

	canal_de_saida := make(chan string)

	go func(){
		for {
			select{
			case mensagem := <- canal_de_entrada_1:
				canal_de_saida <- mensagem

			case mensagem := <- canal_de_entrada_2:
				canal_de_saida <- mensagem
			}
		}
	}()

	return canal_de_saida
}

func escrever(texto string) <- chan string{
	canal := make(chan string)

	go func(){
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}	
	}()

	return canal
}