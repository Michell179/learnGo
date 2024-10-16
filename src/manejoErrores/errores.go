package main

import (
	"errors"
	"fmt"
	"strconv"
	"syscall"
)

func main() {

	str := "123"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Número:", num)

	// Manejamos Nuestros errores personalizados.
	fmt.Println("\nError personalizado.")
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Resultado:", resultado)

}

func dividir(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		// Creamos nuestro propio mensaje personalizado con errors.New(string)
		// Se puede usar fmt.ErrorF(string) // formatearlo y retornarlo
		return syscall.INFINITE, errors.New("Infinity")
	}
	return dividendo / divisor, nil
}
