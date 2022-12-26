package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos = pseudo-herança
	turboLigado bool
}

func main() {
	// Usando essa tecnica podemos gerar uma pseudo-herança de
	// estruturas, mas q na verdade é uma composição
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrai %s está com o turbo ligado? %v", f.nome, f.turboLigado)
}
