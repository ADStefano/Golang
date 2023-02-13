package main

import "fmt"

func closure() func() int{

	i := 0

	return func() int {

		i ++
		return i
	}
}

func main()  {

	sequencia := closure()
	fmt.Println(sequencia()) // 1
	fmt.Println(sequencia()) // 2
	fmt.Println(sequencia()) // 3

	fmt.Println("--------")

	sequencia2 := closure()
	fmt.Println(sequencia2()) // 1
	fmt.Println(sequencia2()) // 2
	fmt.Println(sequencia2()) // 3
	
	
}