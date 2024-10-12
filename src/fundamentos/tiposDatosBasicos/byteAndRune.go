package main

import "fmt"

func main() {

	// Puede almacenar caracteres y valores entre 0 y 255
	var a byte = 255
	fmt.Println(a)

	// Al acceder al indice de 0 de un string obtenemos su valor en ASCII
	var name string = "Hola"
	fmt.Println(name[0])

	var r rune = -1232
	fmt.Println(r)

}
