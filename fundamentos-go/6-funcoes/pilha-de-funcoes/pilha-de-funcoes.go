package main

import "runtime/debug"

//O conceito de pilha diz basicamente a ordem em que uma serie
//de funções será executada, sempre colocando uma função na pilha,
//e no fim, quando a função termina completamente retirando-a da pilha
func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
