package main

import "fmt"

func main() {

	a := 40
	b := 100
	// Operadores + - / * % ^(modulo)
	fmt.Println("+:", a+b)
	fmt.Println("-:", a-b)
	fmt.Println("/:", a/b)
	fmt.Println("*:", a*b)
	fmt.Println("%:", a%b)
	//	Operador XOR
	fmt.Println("^:", a^b)
	// Operador AND
	fmt.Println("&:", a&b)
	// Operaor OR
	fmt.Println("|:", a|b)
	// Mover bits de b a la izquierda
	fmt.Println("<<:", a<<b)
	// Mover bits de a a la derecha
	fmt.Println(">>:", a>>b)
	// XOR AND not
	fmt.Println("&^:", a&^b)

	// Operadores de incremento y decremento ++ y --
	b++
	fmt.Println("b++:", b)
	a--
	fmt.Println("b--:", a)

	// a = a + 5
	a += 5
	fmt.Println("a+= 5", a)

}
