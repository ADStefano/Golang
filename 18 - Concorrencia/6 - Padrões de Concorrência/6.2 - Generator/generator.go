// Padrão Generate para go routine, serve para encapsular uma chamada de go routine
// e retorna um canal de comunicação para ser feita com esa go routine
package main

import (
	"fmt"
	"time"
)

func main(){

	canal := escrever("Olá Mundo!")

	fmt.Println(<-canal)

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