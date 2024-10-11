package main

import "fmt"

const Pi = 3.1415926

const (
	X = 100
	Y = 0b1010 // binario
	Z = 0o12   // Octal
	W = 0xFF   // Hexadecimal
)

// Usando iota, se incrementa de 1 las variables debajo de la asignacion.
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println(Pi)
	fmt.Println(X, Y, Z, W)
	fmt.Println(Viernes)
}
