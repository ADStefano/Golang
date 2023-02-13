package main

import "fmt"

type pessoa struct{

	nome string
	sobrenome string
	idade uint8

}

type aluno struct{

	pessoa
	faculdade string
	curso string
	periodo uint8

}

func main()  {
	
	pessoa1 := pessoa{"Ângelo", "Di Stefano", 23}
	fmt.Println(pessoa1)

	aluno1 := aluno{pessoa1,"PUC PR", "Análise e desenvolvimento de sistemas", 2}
	fmt.Println(aluno1)

}