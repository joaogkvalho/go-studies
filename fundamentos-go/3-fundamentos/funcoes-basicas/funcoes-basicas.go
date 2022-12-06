package main

import "fmt"

//recebe 2 inteiros como parametro e retorna tambem um inteiro
func somar(a int, b int) int {
	return a + b
}

//recebe um inteiro como parametro mas não retorna nada
func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := somar(5, 5)
	imprimir(resultado)
}
