package main

import "fmt"

func main() {
	x := 1
	y := 2

	//o Go apenas os operadores unarios pós-fixados
	x++ //x += 1 OU x = x + 1
	fmt.Println(x)

	y-- //y -= 1 OU y = y - 1
	fmt.Println(y)
}
