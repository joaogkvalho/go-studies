package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Garvalho": 7645.99,
			"Gerson Silva":     5476.56,
		},
		"J": {
			"João Carvalho": 8067.56,
		},
		"P": {
			"Paulo Coelho": 2354.78,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
