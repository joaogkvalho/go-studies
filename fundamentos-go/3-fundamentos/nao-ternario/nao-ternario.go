package main

import "fmt"

//Na linguagem Go não temos o operador ternário,
//a alternativa será usar a estrutura IF.
func obterResultado(nota float64) string {
	// nota > 7 ? "Aprovado" : "Reprovado"
	if nota > 7 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

func main() {
	fmt.Println(obterResultado(6.9))
}
