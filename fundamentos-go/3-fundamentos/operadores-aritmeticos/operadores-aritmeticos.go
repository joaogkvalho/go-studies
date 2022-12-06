package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	//operações básicas
	fmt.Println("Soma ==>", a+b)
	fmt.Println("Subtração ==>", a-b)
	fmt.Println("Divisão ==>", a/b)
	fmt.Println("Multiplicação ==>", a*b)
	fmt.Println("Módulo ==>", a%b) //resto da divisão

	//operações bit a bit
	fmt.Println("E ==>", a&b)   //11 & 10 = 10
	fmt.Println("Ou ==>", a|b)  //11 | 10 = 11
	fmt.Println("Xor ==>", a^b) //11 ^ 10 = 01

	//operações usando o método Math
	c := 3.0
	d := 2.0
	fmt.Println("Maior numero ==>", math.Max(c, d))
	fmt.Println("Menor numero ==>", math.Min(c, d))
	fmt.Println("Potenciação ==>", math.Pow(c, d))
}
