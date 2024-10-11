package main

import (
	"fmt"
)

func main() {

	// Declaracion Variables
	// Tipos de datos.
	//var firstName, lastName2 string
	//var age int32
	//firstName = "Haru"
	//lastName2 = "Maven"
	//var number int
	//number = 23

	/*Forma 1*/
	var (
		firstName string = "Haru"
		lastName2 string = "Maven"
		number    int32  = 24
	)
	fmt.Println(firstName, lastName2, number)

	/*Forma 2*/
	var (
		firstName2 = "Haru"
		lastName3  = "Maven"
		number2    = 123
	)
	fmt.Println(firstName2, lastName3, number2)

	/*Forma 3*/
	// Nota: Solo dentro de las funciones
	firstName3, lastName4, number3 := "Haru", "Maven", 123
	fmt.Println(firstName3, lastName4, number3)

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
