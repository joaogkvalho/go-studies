package main

import (
	findtitle "9-module/find-title"
	"fmt"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := findtitle.Titulo(url1)
	c2 := findtitle.Titulo(url2)
	c3 := findtitle.Titulo(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://google.com",
		"https://youtube.com",
		"https://github.com",
	)

	fmt.Println(campeao)
}
