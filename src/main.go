package main

import "fmt"

func main() {
	//Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	//Println
	fmt.Println(helloMessage, worldMessage) //Imprime con un salto de linea automatico
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos \n", nombre, cursos) //Imprime según el formato
	fmt.Printf("%v tiene más de %v cursos \n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos) //Retorna la un string como variable usando el mismo metodo de Printf
	fmt.Println(message)

	//Tipos de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T", cursos)
}
