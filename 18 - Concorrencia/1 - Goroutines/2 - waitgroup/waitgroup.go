package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func Waitgroup(){
	// Concorrência != Paralelismo
	// Concorrência roda os programas quase que ao mesmo tempo, é como se eles revezassem para rodar
	// Paralelismo roda os programas ao mesmo tempo
	// go antes da função indica que deve ser utilizado a concorrência
	
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // Número de concorrencias

	go func(){
		escrever("Olá Mundo!")
		waitGroup.Done() // Indica que o programa acabou e remove 1 concorrencia do Add
	}()

	go func(){
		escrever("Programando em go!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // Espera os programas rodarem para terminar/ o Add ficar com 0


}

func escrever(texto string){
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}