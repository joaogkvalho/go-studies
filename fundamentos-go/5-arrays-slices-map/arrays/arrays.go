package main

import "fmt"

func main() {
	//Arrays dentro do Go são estruturas homogeneas (mesmo tipo) e fixas (mesmo tamanho)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 5.5, 3.3, 8.5
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}
