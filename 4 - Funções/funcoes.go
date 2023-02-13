package main

import "fmt"

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int16) (int16, int16) {
	soma:= n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main()  {
	
	soma:= somar(10,20)
	fmt.Println(soma)

	resultadoSoma, _ := calculos(20,30)
	fmt.Println(resultadoSoma)

	resultSoma, _ := calculos(10,11)
	fmt.Println(resultSoma)


	var funcao = func (txt string) string {
		fmt.Println(txt)
		return txt
	}

	f := funcao("texto da função")
	fmt.Println(f)
}