package main

import (
	"time"
	"fmt"
)

func main() {

	canal1, canal2 := make(chan string), make(chan string)

	// A cada meio segundo envia a string para o canal1
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	// A cada 2 segundos envia a string para o canal2
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()	

	// Select impede que o canal2 trave o canal1 por conta do delay
	// de 2 segundos e permite que a concorrência aconteça sem perda de tempo
	for {
		select {
		case mensagemCanal1 := <- canal1:
			fmt.Println(mensagemCanal1	)

		case mensagemCanal2 := <- canal2:
			fmt.Println(mensagemCanal2)
		}
	}

}