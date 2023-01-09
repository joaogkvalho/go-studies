package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para um canal (escrita)
	<-ch    // lendo dados de um canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
