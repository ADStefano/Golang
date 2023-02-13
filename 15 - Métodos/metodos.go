package main

import "fmt"

type user struct {

	nome string
	idade uint8
}

func (u user) mostrarDados() {
	fmt.Printf("Nome:%s\nIdade:%d\n", u.nome, u.idade)	
}

func (u user) maiorDeIdade() bool {
	return u.idade >= 18
}

// Altera o aniversário utilizando o ponteiro como parâmetro
func (u *user) fazerAniversario(){
	u.idade++
}

func main()  {
	
	usuario1 := user{"Roberto Carlos", 30}
	fmt.Println(usuario1)
	usuario1.mostrarDados()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario1.fazerAniversario()
	usuario1.mostrarDados()
}