package main

import "fmt"

func helloWorldFunction(message string) {
	fmt.Println(message)
}

func tripleArgumento(a, b int, c string) {
	fmt.Println(a, b, c)
}

func multiplicaPorDos(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	helloWorldFunction("Hola mundo")
	tripleArgumento(1, 2, "hola")

	value := multiplicaPorDos(2)
	fmt.Println("Value:", value)

	_, value2 := doubleReturn(2) // ignorar un valor de una funcion que retorna más valores
	fmt.Println("Value y Value *2:", value2)
}
