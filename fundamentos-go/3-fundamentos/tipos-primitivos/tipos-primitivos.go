package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//NUMEROS INTEIROS (Int)
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(50))

	//sem sinal (só positivos)... uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	//representa um mapeamentos da tabela unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O rune é do tipo", reflect.TypeOf(i2))

	//NUMEROS REAIS (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))

	//podemos assim tbm inferir um tipo para uma variavel
	var y = float32(29.99)
	fmt.Println("O tipo de y é", reflect.TypeOf(y))

	//sempre quando o valor é inferido o tipo da variavel é float64
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//BOOLEANS
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))

	//operação de negação
	fmt.Println(!bo)

	//STRINGS
	s1 := "Olá meu nome é João"
	fmt.Println(s1 + "!")

	//retorna a quantidade de caracters presentes na string
	fmt.Println("O tamanho da string é", len(s1))

	//string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	João
	`
	fmt.Println("O tamanho da string é", len(s2))
}
