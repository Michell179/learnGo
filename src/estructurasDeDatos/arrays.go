package main

import "fmt"

func main() {

	// Array
	// Funciona de manera que el primer indicice es 0 y almacena 200
	// el ultimo indice es 4 y almacena 40
	// tiene 4 elementos y su indice mayor es 3.
	var array = [...]int{200, 300, 20, 40}
	array[0] = 1000
	array[1] = 2000

	/*Forma 1 de recorrer el array/matriz*/
	for i := 0; i < len(array); i++ {
		fmt.Println("array:", array[i])
	}

	for index, value := range array {
		fmt.Println("index:", index, "value:", value)
	}

	// Matriz o array BiDimensional
	var matriz = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matriz)

	for i := range matriz {
		fmt.Println(matriz[i])
	}

}
