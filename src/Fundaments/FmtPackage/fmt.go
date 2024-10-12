package main

import "fmt"

func main() {

	// Imprime con salto de linea print"ln" line
	fmt.Println("Hello World")

	// Imprime sin salto de linea.
	fmt.Print("Soy un texto sin linea")
	fmt.Print("Soy un texto sin linea y me agrego a la anterior linea.\n")

	// Imprimir texto con formato usamos Printf
	name := "Michell"
	age := 24
	fmt.Printf("Resultado de Printf: "+"Hola, soy %s y tengo %d años de edad.\n", name, age)

	// Obtenemos el string formateado y poder almacenarlo en una variable.
	response := fmt.Sprintf("Hola, soy %s y tengo %d años de edad.", name, age)
	// Una vez obtenido el resultado formateado, pasamos a imprimirlo en la consola
	fmt.Println("Resultado de Sprintf:", response)

	// Ahora obtenemos datos del teclado usando el mismo paquete fmt.
	// Para lo cual usamos Scanln y pasamos por argumento-
	// -la referencia de memoria donde almacenaremos el dato ingresado por teclado.

	var (
		nameScan string
		ageScan  int16
	)

	fmt.Println("Ingresa su nombre: ")
	fmt.Scanln(&nameScan)
	fmt.Println("Ingresa su edad: ")
	fmt.Scanln(&ageScan)

	fmt.Printf("Hola ingresaste el nombre %s y la edad %d.\n", nameScan, ageScan)

	// Podemos imprimir el tipo de dato de una variable
	fmt.Printf("El tipo de nameScan es: %T\n", nameScan)
	fmt.Printf("El tipo de ageScan es: %T\n", ageScan)

}
