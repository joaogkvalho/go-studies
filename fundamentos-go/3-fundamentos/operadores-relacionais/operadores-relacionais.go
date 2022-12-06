package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Operador ==", "Banana" == "Banana")
	fmt.Println("Operador !=", 3 != 2)
	fmt.Println("Operador >", 2 > 3)
	fmt.Println("Operador <", 3 > 2)
	fmt.Println("Operador >=", 3 >= 3)
	fmt.Println("Operador <=", 2 <= 3)

	//comparações de datas
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("Datas", d1 == d2)
	fmt.Println("Datas", d1.Equal(d2))

	//comparações de struct
	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoas", p1 == p2)
}
