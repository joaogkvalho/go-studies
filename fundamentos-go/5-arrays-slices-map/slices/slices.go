package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //criação de um array
	s1 := []int{1, 2, 3}  //criação de um slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// Slice nao é um array! Slice define um pedaço de um array
	a2 := [5]int{1, 2, 3, 4, 5}

	//s2 é um pedaço do array a2 que vai do indice 1 até o 3 mas sem incluir o 3
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	// s3 é um novo slice mas que aponta para o mesmo array a2
	s3 := a2[:2]
	fmt.Println(a2, s3)

	// voce pode imaginar um slice como uma estrutura com tamanho e que aponta
	// para um elemento de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
