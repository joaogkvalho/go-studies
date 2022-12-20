package main

import "fmt"

// A função media vai receber um numero variado de parametros do tipo float64
func media(numeros ...float64) float64 {
	total := 0.0
	// Estrutura for ignorando o primero parametro, que seria o indice do array
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.1f", media(5.5, 7.8, 4.6, 10))
}
