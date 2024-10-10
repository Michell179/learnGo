package main

import (
	"fmt"
)

func main() {

	// Variables
	// Tipos de datos.
	var number int
	number = 23
	fmt.Println(number)

	//String
	var lastName string
	var pais string = "colombia"
	var colegio string = "boyaca"
	lastName = "Arias"
	localidad := "cualquiera"

	// Inferencia de datos
	name := "Michell Arias"
	fmt.Println("Hola!", name, lastName)

	// Asignacion de varias variables
	colegio, pais = "Nuevo Boyaca", "Colombia"
	fmt.Println("Hola para todos en" + pais + " del colegio " + colegio + " localidad " + localidad)

	// Tipado fuerte
	// No acepta el cambio de tipo de dato
	var year int16 = 2024
	var edad int8 = 24
	// edad = 1000 -> Esto no es aceptado por Go.
	fmt.Println("Bienvenido al año ", year)
	fmt.Println("Edad:", edad)

}
