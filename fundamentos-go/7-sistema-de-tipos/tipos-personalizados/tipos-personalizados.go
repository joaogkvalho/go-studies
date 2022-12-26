package main

import "fmt"

type nota float64

// usando um tipo personalizado como receiver para adicionar um metodo para nota
func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaPraConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 9.99) {
		return "B"
	} else if n.entre(5.0, 7.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaPraConceito(8.0))
	fmt.Println(notaPraConceito(5.5))
	fmt.Println(notaPraConceito(2.1))
}
