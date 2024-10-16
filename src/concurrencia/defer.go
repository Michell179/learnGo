package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		Se agrega a una pila de ejecución
		Siendo que el primero en ejecutarse seria el 3
		Pila -> primero en llegar, ultimo en salir.
		defer fmt.Println(3)
		defer fmt.Println(2)
		defer fmt.Println(1)
	*/

	file, err := os.Create("defer.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Esto se ejecutará antes que termine la función principal main()
	defer file.Close()

	_, err = file.Write([]byte("Hola, Soy un Archivo"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
