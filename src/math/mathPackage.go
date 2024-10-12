package main

import (
	"fmt"
	"math"
)

func main() {

	// Documentacion oficial: pkg.go.dev/math

	// Funcion Pow()
	var potencia float64
	var raizCuadrada float64
	var raizCubica float64
	var cinco float64 = 2
	var diez float64 = 7

	potencia = math.Pow(cinco, diez)
	raizCuadrada = math.Sqrt(9)
	raizCubica = math.Cbrt(9)

	fmt.Println("Pow():", potencia)
	fmt.Println("Sqrt():", raizCuadrada)
	fmt.Println("Cbrt():", raizCubica)
}
