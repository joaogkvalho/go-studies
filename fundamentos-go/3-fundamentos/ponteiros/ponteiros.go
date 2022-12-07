package main

import "fmt"

func main() {
	i := 1 //criando uma variel I e atribuindo o valor 1

	var p *int = nil //criando um ponteiro que pode armazenar na memória um (int)
	p = &i           //atribuindo o endereço de memoria da variavei I para o ponteiro P

	*p++ //pegando o valor da variavel I e atribuindo ao ponteiro P

	fmt.Println(p, *p, i, &i)
}
