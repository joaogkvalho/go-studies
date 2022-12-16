package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"João Gabriel":   3500.00,
		"Maria José":     4568.98,
		"Roberta Judite": 3556.50,
	}

	funcsESalarios["Lucas Silva"] = 7423.55

	// tentar deletar um valor inexistente de um map nao gera erro dentro do Go
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
