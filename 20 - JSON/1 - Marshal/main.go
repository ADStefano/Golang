package main

import (
	"bytes"
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

	cao := cachorro{"Biscoito","Golden Retriever", 6}
	cao2 := map[string]string{
		"nome":"Bolacha",
		"Raca":"Vira-Lata",
		"Idade":"4",}

	cachorroEmJSON, erro := json.Marshal(cao)
	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println("Biscoito")
	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	cachorro2EmJSON, erro := json.Marshal(cao2)
	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println("Bolacha")
	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}