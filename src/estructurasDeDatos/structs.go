package main

import "fmt"

// structs

type Humano struct {
	nombre string
	edad   int
	correo string
}

func main() {
	/*Forma 1 para definir valores de un Struct*/
	var p1 Humano
	p1.nombre = "Michell"
	p1.edad = 24
	p1.correo = "correo@correo.com"
	fmt.Println("p1:", p1)

	/*Forma 2 para definir valores del Struct*/
	var p2 = Humano{
		nombre: "Hellen",
		edad:   30,
		correo: "correo2@correo.com",
	}
	fmt.Println("p2:", p2)

	/*Forma 3 para definir valores en orden del struct*/
	var p3 = Humano{
		"Haru",
		3,
		"correo3@correo.com",
	}
	fmt.Println("p3:", p3)
}
