package main

import "fmt"

//Declaração de uma variavel que contém uma função
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(5, 5))

	//Declaração de uma variavel que contém uma função de forma reduzida
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(10, 5))
}
