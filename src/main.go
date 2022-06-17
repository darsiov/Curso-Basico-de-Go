package main

import (
	"fmt"
	"math"
)

func main() {
	//Trapecio
	a, b, h := 14, 20, 5

	area := ((a + b) / 2) * h

	fmt.Println("Medidas del trapecio:  a, b, h = (", a, ",", b, ",", h, ") cm")
	fmt.Println("Área del trapecio:", area, "cm^2")

	fmt.Println("-----------------------------------------------------------------------")
	//Circulo
	var r float64 = 4

	areaCirculo := math.Pi * r * r

	fmt.Println("Radio del circulo", r, "cm")
	fmt.Println("Área del ciculo:", areaCirculo, "cm^2")
}
