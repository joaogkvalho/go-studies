package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	// O conceito de closure diz que a função tem ideia do escopo que foi definida
	// e por conta disso o valor impresso será o valor de x do escopo em que a função
	// foi criada
	imprimiX := closure()
	imprimiX()
}
