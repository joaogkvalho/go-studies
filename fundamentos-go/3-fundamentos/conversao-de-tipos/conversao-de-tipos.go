package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	//conversão de (int) para (float64)
	fmt.Println(x / float64(y))

	//conversão de (float64) para (int)
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado... conversões diretas de (int) para string
	//resultaram no valor de dentro da tabela UNICODE
	fmt.Println("Teste " + string(97)) //97 = 'a', 'a' = (int32)

	//conversão (int) para (string)
	fmt.Println("Teste " + strconv.Itoa(123))

	//conversão de (string) para (int)
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	//conversão de boolean
	b, _ := strconv.ParseBool("0")
	if b {
		fmt.Println("Verdadeiro")
	}
}
