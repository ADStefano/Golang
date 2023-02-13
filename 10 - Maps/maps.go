package main

import "fmt"

func main(){

	fmt.Println("Maps")

	usuario := map[string]string{

		"nome" : "João",
		"sobrenome" : "Silvão",
	}

	fmt.Println("Usuário:",usuario)

	// Para acessar apenas 1 valor do map
	fmt.Println("Valor específico de um maps:",usuario["nome"])

	usuario2 := map[string] map[string] string {

		"faculdade": {
			"nome" : "PUCPR",
		},

		"endereco": {
			"logradouro" : "Rua dos bobos",
			"campus" : "Curitiba",
		},
	}

	fmt.Println("Usuário 2:",usuario2)

	// Deletando uma chave 
	delete(usuario2, "faculdade")
	fmt.Println("Resultado após o delete:",usuario2)

	// Adicionando chaves
	usuario2["cursos"] = map[string]string{
		"curso" : "Análise e desenvolvimento de sistemas",
		"curso2" : "Nanotecnologia",
	}

	usuario["signo"] = "leão"

	fmt.Println("Usuário depois de adicionar uma chave nova:",usuario)
	fmt.Println("Usuário 2 depois de adicionar uma chave nova:",usuario2)
	

}	