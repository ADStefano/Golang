package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main(){
	chave := make([]byte, 64)
	if _, erro := rand.Read(chave); erro != nil{
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}