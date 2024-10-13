package main

import (
	"fmt"
	"math"
)

func main() {
	/*Ejercicio:
	Calcule e imprima el área y el perímetro del triángulo
	Crear un programa que solicite al usuario que ingrese los lados
	de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.
	Imprimir:
			Hipotenusa
			Area
			Perimetro
	*/

	/*Solución*/

	const precision = 2
	var base float64
	var altura float64

	fmt.Println("Calculos de un Triangulo.")

	fmt.Println("Ingresa la base:")
	fmt.Scanln(&base)
	fmt.Println("Ingresa la altura:")
	fmt.Scanln(&altura)

	hipotenusa := calcularHipotenusa(base, altura)
	area := calcularArea(base, altura)
	perimetro := calcularPerimetro(base, altura, hipotenusa)

	fmt.Printf("Hipotenusa: %.*f \n", precision, hipotenusa)
	fmt.Printf("Area: %.*f \n", precision, area)
	fmt.Printf("Perimetro: %.*f \n", precision, perimetro)

}

func calcularHipotenusa(x, y float64) float64 {
	var exponente float64 = 2
	a, b := math.Pow(x, exponente), math.Pow(y, exponente)
	return math.Sqrt(a + b)
}

func calcularArea(x, y float64) float64 {
	return (x * y) / 2
}

func calcularPerimetro(x, y, z float64) float64 {
	return x + y + z
}
