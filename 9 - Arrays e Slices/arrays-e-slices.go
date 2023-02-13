package main

import (
	"fmt"
	"reflect"
)

func main()  {
	
	fmt.Println("Arrays e Slices")
	
	// Método explícito
	var array1[5] string
	array1[0] = "Posição 1"
	array1[1] = "Posição 2"
	array1[2] = "Posição 3"
	array1[3] = "Posição 4"
	array1[4] = "Posição 5"

	fmt.Println("Array com strings:",array1)

	// Método implícito
	array2 :=[5] int {1,2,3,4,5}

	fmt.Println("Array com int:",array2)

	// Método menos utilizado
	array3 := [...] bool {true, false, true, false}

	fmt.Println("Array booleano:",array3)

	// Slice é um array sem limite de itens
	slice := []int {10, 11, 12, 13, 14, 15}
	fmt.Println("Slice:",slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))

	slice = append(slice, 16)

	fmt.Println("Slice:",slice)

	slice2 := array1[1:3]

	fmt.Println("Slice 2:",slice2)

	array1[1] = "Posição alterada"
	fmt.Println("Slice 2:",slice2)

	// Array interno
	fmt.Println("Array interno")

	slice3 := make([]float32, 5,6)
	fmt.Println("Slice 3:",slice3)
	fmt.Println("Tamanho:",len(slice3)) // length
	fmt.Println("Capacidade:",cap(slice3)) // capacity

	slice3 = append(slice3, 10)
	fmt.Println("Slice 3:",slice3)
	fmt.Println("Tamanho:",len(slice3)) 
	fmt.Println("Capacidade:",cap(slice3))

	slice3 = append(slice3, 11)
	fmt.Println("Slice 3:",slice3)
	fmt.Println("Tamanho:",len(slice3)) 
	fmt.Println("Capacidade:",cap(slice3)) 

	slice4 := make([]float32, 5)
	fmt.Println("Slice 4:",slice4)
	fmt.Println("Tamanho:",len(slice4)) 
	fmt.Println("Capacidade:",cap(slice4))

	slice4 = append(slice4, 1)
	fmt.Println("Slice 4:",slice4)
	fmt.Println("Tamanho:",len(slice4)) 
	fmt.Println("Capacidade:",cap(slice4))

	
}