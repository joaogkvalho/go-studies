package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// não é possivel fazer append em arrays
	// array1 = append(array1, 4, 5, 6)

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)

	//usando o metodo copy o slice2 vai receber todos os dados do slice1
	copy(slice2, slice1)
	fmt.Println(slice2)
}
