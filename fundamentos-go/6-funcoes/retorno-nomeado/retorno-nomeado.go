package main

import "fmt"

//Declaração de 2 retornos nomeados
func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1

	//Declaração de um retorno limpo
	return
}

func main() {
	//Não é necessário utilizar os mesmos nome dos retornos nomeados
	x, y := trocar(5, 2)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
