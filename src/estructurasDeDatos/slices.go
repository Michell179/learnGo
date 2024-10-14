package main

import "fmt"

func main() {
	// Los slices son estructuras de datos dinamicos.
	// Podemos eliminar y agregar nuevos elementos

	semana := []string{
		"domingo",
		"lunes",
		"martes",
		"miercoles",
		"jueves",
		"viernes",
		"sabado"}
	fmt.Println(semana)

	// Creamos un nuevo slice basado en el anterior
	// El primer indice es inclusivo y el segundo (5) es no inclusivo
	// por lo tanto vamos a tener los elementos con indice 1,2,3,4
	diasHabiles := semana[1:5]
	fmt.Println("DiasHabiles-1", diasHabiles)

	// Agregamos un nuevo elemento
	diasHabiles = append(diasHabiles, semana[5])
	fmt.Println("DiasHabiles-2", diasHabiles)

	fmt.Println("len:", len(diasHabiles))
	fmt.Println("cap:", cap(diasHabiles))

	// Eliminamos un elemento haciendo uso de append.
	// Pasando dos slices podemos determinar los indices a usar.
	// Vamos a guardar domingo, viernes y sabado del slice semana
	// Podriamos usar el mismo slice sin problema.
	diasHabiles = append(semana[:1], semana[5:]...)
	fmt.Println("DiasHabiles-3", diasHabiles)

	// En este caso crearemos un slice con los dias
	// Domingo y Sabado
	diasHabiles = append(diasHabiles[:1], diasHabiles[2:]...)
	fmt.Println("DiasHabiles-4", diasHabiles)

	/*Usando la funcion make()*/
	// Se usa para crear slices vacios y determinar el espacio y capacidad
	// Esta capacidad podemos aumentarla con la funcion append()
	// Longitud inicial 3 y la capacidad la definimos con 6
	fmt.Println("Funcion make()")
	nombres := make([]string, 3, 6)
	nombres[0] = "Michell"
	fmt.Println(nombres)
	fmt.Println("len-nombres:", len(nombres))
	fmt.Println("cap-nombres:", cap(nombres))

	/*Usando la funcion copy()*/
	// Esta funcion retorna la cantidad de lementos copiados.
	fmt.Println("\nFuncion copy()")
	sliceEstatico := []int{1, 2, 3, 4, 5}
	sliceVacio := make([]int, 5)
	fmt.Println("Antes: sliceEstatico", sliceEstatico)
	fmt.Println("Antes: sliceVacio", sliceVacio)
	// Primer argumento es el slice que recibirá los datos copiados.
	// Segundo argumento es el slice que vamos a copiar
	elementosCopiados := copy(sliceVacio, sliceEstatico)

	fmt.Println("Elementos Copiados:", elementosCopiados)
	fmt.Println("después: SliceEstatico:", sliceEstatico)
	fmt.Println("después: SliceVacio:", sliceVacio)
}
