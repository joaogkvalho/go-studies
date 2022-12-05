package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo infeiro pelo go (float64)

	//forma reduzida de criar variaveis
	area := PI * math.Pow(raio, 2)
	fmt.Println("A area da circunferencia é:", area)

	//agrupamento de variaveis e constantes
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// declarar mais de uma variavel em uma so linha
	var e, f bool = false, true
	fmt.Println(e, f)

	//declarar mais de uma variavel em uma so linha de forma reduzida
	g, h, i := 2, false, "hello" //inferencia de tipos (number, bool, string)
	fmt.Println(g, h, i)
}
