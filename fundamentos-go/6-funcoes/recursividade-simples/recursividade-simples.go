package main

import "fmt"

// Usando valores sem sinal o código fica muito mais simples
// pois não existe fatorial de número com sinal, e por causa disso
// o próprio Go impossibilita a passagem de um número negativo para
// a função
func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
