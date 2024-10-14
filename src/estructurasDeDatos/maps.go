package main

import "fmt"

func main() {
	/*
		Estructura de datos donde almacenamos datos
		con Clave y valor
	*/

	colors := map[string]string{
		"rojo":   "#FF0000",
		"azul":   "#0000FF",
		"blanco": "#FFFFFF",
	}
	fmt.Println(colors)

	// Agregamos un nuevo valor al mapa
	// creamos la nueva clave y asignamos el valor correspondiente a la clave.
	colors["negro"] = "#000000"
	fmt.Println(colors)

	// Recuperar un valor de un mapa usando su clave
	color := colors["azul"]
	fmt.Println(color)

	// Se puede recuperar un valor booleano si el dato/elemento existe.
	// true cuando existe, y false cuando no existe, cuando no se recupera el valor
	color, ok := colors["azul"]
	fmt.Println("Dato:", color, "Existe:", ok)

	// Recomendado para poder recuperar un valor de un mapa de manera correcta.

	color, ok = colors["blanco"]

	if ok {
		fmt.Printf("color %s encontrado\n", color)
	} else {
		fmt.Println("No existe el color")
	}

	//Eliminar elemento de un mapa usando delete()
	delete(colors, "negro")
	fmt.Println(colors)

	// Iterar elementos de un mapa.

	fmt.Println("Iterando un mapa...")
	for clave, valor := range colors {
		fmt.Println("clave:", clave, "valor:", valor)
	}
}
