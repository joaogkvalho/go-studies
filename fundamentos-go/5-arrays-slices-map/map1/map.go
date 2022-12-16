package main

import "fmt"

func main() {
	// maps devem ser sempre inicializados
	// var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[80989789797] = "Maria"
	aprovados[85293859835] = "João"
	aprovados[52352353255] = "Gabriel"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// lendo um valor de um map
	fmt.Println(aprovados[85293859835])

	// deletando um valor de um map
	delete(aprovados, 85293859835)
	fmt.Println(aprovados)
}
