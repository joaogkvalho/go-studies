package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	// Forma padrão de criação de uma variavel do tipo struct
	var produto1 produto
	produto1 = produto{
		nome:     "Borracha",
		preco:    2.50,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	// Forma reduzida de criação de uma variavel do tipo struct
	produto2 := produto{"Notebook", 2050.99, 0.10}
	fmt.Println(produto2, produto2.precoComDesconto())
}
