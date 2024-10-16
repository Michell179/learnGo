package main

import "fmt"

func main() {
	divir(100, 10)
	divir(100, 25)
	divir(100, 0)
	divir(100, 4)
}

func divir(dividendo, divisor int) {
	// Creamos una funcion anonima
	// Pero tambien posponemos la ejecucion
	// y recuperamos cualquier panico generado
	// Con la palabra recover

	/*
		Nota: El uso de panic y recover() se recomienda para casos excepcionales
		o para rutinas de limpieza, cierre de recursos antes de recuperar el aplicativo.
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	validarCero(divisor)
	fmt.Println(dividendo / divisor)
}

func validarCero(divisor int) {
	if divisor == 0 {
		// Generamos un panico en nuestro aplicativo
		// El lenguaje lo puede generar de manera automatica.
		// Pero se reocimienda usar para casos excepcionales.
		// Como indices fuera de rango, y situaciones excepcionales.
		panic("No se puede dividir por cero")
	}

}
