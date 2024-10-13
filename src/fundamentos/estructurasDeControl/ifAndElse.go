package main

import (
	"fmt"
	"time"
)

func main() {
	// Usando estructuras de control if, else if y else.
	// Validamos la hora

	//hora := time.Now().Hour()

	//if hora < 12 {
	//	fmt.Println("Mañana!")
	//} else if hora < 18 {
	//	fmt.Println("Tarde!")
	//} else {
	//	fmt.Println("Noche!")
	//}

	// Inicializar variables dentro del if
	if hora := time.Now().Hour(); hora < 12 {
		fmt.Println("Mañana!")
	} else if hora < 18 {
		fmt.Println("Tarde!")
	} else {
		fmt.Println("Noche!")
	}

}
