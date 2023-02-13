package main

import "fmt"

func main()  {
	
	// Função anônima simples
	func(){
		fmt.Println("Olá mundo!")
	}()

	// Função anônima com parâmetros
	func(texto string){
		fmt.Println(texto)
	}("Oi mundo!")

	// Função anônima com return
	retorno := func(texto string) string{
		return fmt.Sprintf("%s mundo!", texto)
	}("Eae")

	fmt.Println(retorno)
}