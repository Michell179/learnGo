package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x

	// P contendrá el puntero de memoria de x
	// Imprimimos el puntero a la memoria de x
	fmt.Println(&x)
	// Imprimimos el valor de x
	fmt.Println(x)
	// Imprimimos el puntero de memoria de p que corresponde al de x
	fmt.Println(p)
	// Imprimimos el valor del puntero p siendo este el valor asignado a x
	fmt.Println(*p)

	var a int = 5
	fmt.Println(a)
	// Enviamos la referencia de la memoria.
	editar(&a)
	// Este valor se modificará ya que editar está recibiendo la referencia de memoria.
	fmt.Println(a)

}

// Esta funcion recibe un puntero de memoria, por lo que este modificará
// el valor existente dentro de este espacio en memoria
// Para el ejemplo se pasa el puntero de a siendo este espacio el que se modificará.
func editar(a *int) {
	*a = 20
}
