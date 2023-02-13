package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct{
	Nome string `json:"Nome"`
	Raca string `json:"Raca"`
	Idade uint  `json:"Idade"`
}

func main(){
	cachorroEmJSON := `{"Nome":"Biscoito","Raca":"Golden Retriever","Idade":6}`
	cachorroEmJSON2 := `{"Nome":"Bolacha","Raca":"Vira-Lata","Idade":4}`

	var cao cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &cao); erro != nil{
		log.Fatal(erro)
	}

	fmt.Println("Biscoito em JSON")
	fmt.Println(cao)

	cao2 := cachorro{}

	if erro := json.Unmarshal([]byte(cachorroEmJSON2), &cao2); erro != nil{
		log.Fatal(erro)
	}

	fmt.Println("Bolacha em JSON")
	fmt.Println(cao2)
}