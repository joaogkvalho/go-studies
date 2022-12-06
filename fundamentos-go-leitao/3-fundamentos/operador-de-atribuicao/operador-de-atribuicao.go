package main

import "fmt"

func main() {
	//atribuição simples
	var b byte = 3
	fmt.Println(b)

	//atribuição reduzida com inferencia de tipo
	i := 5 //i = 5
	fmt.Println(i)

	//atribuição aditiva
	i += 3 //i = i + 5
	fmt.Println(i)

	//atribuição subtrativa
	i -= 3 //i = i - 3
	fmt.Println(i)

	//atribuição divisiva
	i /= 3 //i = i / 3
	fmt.Println(i)

	//atribuição multiplicativa
	i *= 3 //i = i * 3
	fmt.Println(i)

	//atribuição modulativa
	i %= 3 //i = i % 3
	fmt.Println(i)

	//atribuição de 2 valores
	x, y := 1, 2
	fmt.Println(x, y)

	//troca de valores usando atribuição
	x, y = y, x
	fmt.Println(x, y)
}
