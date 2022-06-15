package main

import "fmt"

func main() {
	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta

	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicación

	result = x * y
	fmt.Println("Multiplicacion:", result)

	//Division

	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Residuo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

	//Retos, calcular el area de un trapecio y un circulo
}
