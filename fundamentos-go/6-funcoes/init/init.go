package main

import "fmt"

// A função init é executada automaticamente antes da função main
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
