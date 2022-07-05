package main

import (
	"fmt"
	"math"
)

func areaTrapecio(a, b, h float64) float64 {
	area := ((a + b) / 2) * h
	return area
}
func areaCirculo(r float64) float64 {
	area := math.Pi * r * r
	return area
}
func separador() {
	fmt.Println("-----------------------------------------------------------------------")
}

func main() {
	var a, b, h float64 = 14, 20, 5
	var r float64 = 4
	aT := areaTrapecio(a, b, h)
	aC := areaCirculo(r)
	fmt.Println("Medidas del trapecio:  a, b, h = (", a, ",", b, ",", h, ") cm")
	fmt.Println("Área del trapecio:", aT, "cm^2")
	separador()
	fmt.Println("Radio del circulo", r, "cm")
	fmt.Println("Área del ciculo:", aC, "cm^2")

}
