package main

import "fmt"

func main() {
	/*Operadores Logicos
	! Not
	&& AND
	|| OR
	> Mayor que
	< Menor que
	>= Mayor igual que
	<= Menor igual que
	*/

	x := 5
	y := 5
	z := 10
	// Expresion asignada a una variable usando operadores logicos y aritmeticos.
	resultado := (((x-y)*2)/3) > z && x != y
	fmt.Println(resultado)

	resultado = x >= y && z != y
	fmt.Println(resultado)
}
