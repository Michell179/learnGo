package main

import "fmt"

func main() {
	resp := saludo("Michell")
	fmt.Println(resp)

	sum, mult := calc(4, 5)
	sum2, mult2 := calc2(4, 5)

	fmt.Println("La suma es", sum)
	fmt.Println("La multiplicacion es", mult)

	fmt.Println("La suma es", sum2)
	fmt.Println("La multiplicacion es", mult2)

}

// Funciones que reciben un argumento de tipo string y retornan
// Datos de tipo string
func saludo(nombre string) string {
	return "Saludo " + nombre
}

// Funciones que reciben dos datos de tipo entero, simplificando en el ultimo dato
// La funcion
func calc(a, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

// Podemos retornar variables directamente de x tipo de valor
// y solo asignarles el valor y usando return definimos donde termina la funcion.
func calc2(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}
