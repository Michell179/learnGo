package main

import "fmt"

func main() {
	// Bucle infinito
	i := 10
	for {
		fmt.Println("Hola", i)
		i--
		if i < 0 {
			// Hacemos uso de la palabra break, para detener el ciclo.
			break
		}
	}

	i = 10
	fmt.Println(i)
	// Bucle con condicion
	for i <= 15 {
		fmt.Println("Mundo", i)
		i++
	}

	// Bucle con inicializacion, condicion y actualizacion
	for i := 0; i <= 10; i++ {
		if i == 8 {
			// Usando la palabra continue podemos saltar a la siguiente iteracion
			// y de esta manera no imrpimir el valor actual de i en este caso es 8.
			continue
		}
		fmt.Println("Clasico", i)
	}

}
