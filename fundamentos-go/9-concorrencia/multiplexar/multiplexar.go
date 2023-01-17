package main

import (
	findtitle "9-module/find-title"
	"fmt"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		findtitle.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		findtitle.Titulo("https://github.com", "https://youtube.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
