package main

import "fmt"

func main() {
	// criando um slice de inteiros com tamanho 10
	s := make([]int, 10)

	// atribuindo o valor 5 para o indice 9 do slice s
	s[9] = 5
	fmt.Println(s)

	// reatribuindo a s um slice de tamanho 10 mas que aponta para um array de tamanho 20
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	// atribuindo mais 10 valores ao slice s
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// ao atribuir um valor que ultrapassa o tamanho do array interno
	// o proprio Go dobra o tamanho desse array para suportar a capacidade do slice
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
