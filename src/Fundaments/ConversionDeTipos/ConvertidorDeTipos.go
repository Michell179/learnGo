package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Conversion de tipos de manera explicita haciendo uso de funciones integradas.
	var integer16 int16 = 16
	var integer32 int32 = 32
	var integer64 int64
	var integer int = 100
	var stringNumero string = "100s"

	// Ejemplo con enteros conversion explicita
	fmt.Println("Suma de int16 + int32:", int32(integer16)+integer32)

	// Ejemplo de String a enteros haciendo uso de paquete strconv interno de Go.
	// Esta operacion strconv.Atoi me retorna un int y posible error.
	var stringConver string
	var intConver int
	var err error
	// Convertimos de String a entero
	intConver, err = strconv.Atoi(stringNumero)

	//Convertimos de entero a String
	stringConver = strconv.Itoa(integer)
	fmt.Println("stringConver", stringConver)

	// valida si existe error en la conversion de tipos.
	if err != nil {
		fmt.Println(err)
	}

	// Convertimos los tipos int a int64 de manera explicita.
	integer64 = int64(intConver + integer)
	fmt.Println("Resultado de int + int almacenada en int64:", integer64)

}
