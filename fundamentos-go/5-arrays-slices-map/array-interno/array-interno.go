package main

import "fmt"

func main() {
	// com esse simples teste percebemos que 2 slices
	// diferentes podem apontar para o mesmo array interno
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	fmt.Println(s1, s2)

	s1[0] = 8
	fmt.Println(s1, s2)
}
