package main

import "fmt"

//Função sem parametros e que nao retorna nada
func f1() {
	fmt.Println("F1")
}

//Função que recebe duas strings como parametro
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

//Função que retorna obrigatoriamente uma string
func f3() string {
	return "F3"
}

//Função que recebe duas string declaradas de forma
// reduzida e retorna outra string
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	f3, f4 := f3(), f4("Param1", "Param2")
	fmt.Println(f3)
	fmt.Println(f4)

	f51, f52 := f5()
	fmt.Println("F5:", f51, f52)
}
