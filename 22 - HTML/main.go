package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct{
	Nome string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{"Ângelo", "angelo@email.com"}
	templates.ExecuteTemplate(w, "home.html", u)
}

var templates *template.Template

func main(){

	templates = template.Must(templates.ParseGlob("*.html"))

	http.HandleFunc("/home", home) 		

	fmt.Println("Executando localhost...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}