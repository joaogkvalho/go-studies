package main

import "fmt"

//função que recebe 2 valores booleans e retorna 3 valores booleans
func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    //operador E
	comprarTv32 := trab1 != trab2    //ou exclusivo
	comprarSorvete := trab1 || trab2 // operador OU

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(false, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
