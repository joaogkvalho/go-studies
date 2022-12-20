package main

import "fmt"

// A palavra defer dentro do Go serve basicamente para atrasar a
// execução de um função ou o retorno dela
func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo...")
		return numero
	}
}

func main() {
	fmt.Println(obterValorAprovado(6000))
}
