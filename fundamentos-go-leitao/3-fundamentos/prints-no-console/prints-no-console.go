package main

import "fmt"

func main() {
	//Printa no console valores na mesma linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	//No final do valor pula para outra linha
	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	//Retorna o valor em formato de string
	xs := fmt.Sprint(x)

	//Agora é possivel concatenar com outra string
	fmt.Println("O valor de PI é " + xs)

	//Concatena sem a necessidade de fazer a tranformação de tipo
	fmt.Println("O valor de PI é", x)

	//Printa no terminal e ainda ha a possibilidade do valor ser formatado
	fmt.Printf("O valor de PI é %.2f", x)

	a := 2
	b := 1.9999
	c := false
	d := "hello"

	//Print no terminal usando "ganchos", passando os valores depois
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	//Tipos de ganchos
	//%d = Inteiro
	//%f = Float64
	//%t = Bool
	//%s = String
}
