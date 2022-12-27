package main

import "fmt"

type esportivo interface {
	turboLigado()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// Poderiamos aqui adicionar mais propriedades ou metodos
}

type bmw7 struct{}

func (b bmw7) turboLigado() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.fazerBaliza()
	b.turboLigado()
}
