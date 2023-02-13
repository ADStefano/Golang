package main

import "fmt"

func generico(in interface{}){
	fmt.Println(in)
}

func main()  {
	
	generico("String")
	generico(10)
	generico(true)
	generico(10.100)
}