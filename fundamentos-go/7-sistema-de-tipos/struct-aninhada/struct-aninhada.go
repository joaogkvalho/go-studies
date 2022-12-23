package main

import "fmt"

type item struct {
	productId int
	qntd      int
	preco     float64
}

type pedido struct {
	userId int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qntd)
	}

	return total
}

func main() {
	pedido := pedido{
		userId: 1,
		itens: []item{
			{productId: 1, qntd: 4, preco: 2.99},
			{productId: 2, qntd: 9, preco: 8.75},
			{productId: 3, qntd: 80, preco: 20.50},
		},
	}

	fmt.Printf("O valor total do pedido é R$%.2f", pedido.valorTotal())
}
