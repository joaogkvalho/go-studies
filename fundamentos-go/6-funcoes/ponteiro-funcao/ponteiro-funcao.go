package main

import "fmt"

func inc1(n int) {
	n++ // n = n + 1
}

// A função inc2 recebe como parametro um ponteiro de inteiro
func inc2(n *int) {
	*n++ // atribuindo +1 para o valor no qual o ponteiro aponta
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	inc2(&n) // pegando por referencia o valor da variavel n
	fmt.Println(n)
}
