package main

import "fmt"

func imprimirResultado(nota float64) {
	//no Go dentro da estrutura If os parametros nao sao envolvidos com parenteses
	if nota >= 7 {
		fmt.Println("Aprovado com a nota: ", nota)
	} else {
		fmt.Println("Reprovado com a nota: ", nota)
	}
}

func main() {
	imprimirResultado(7)
}
