package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	Nombre     string
	Desc       string
	Completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

func main() {

	// Creamos la instancia del struct ListaTareas
	lista := ListaTareas{}

	// Usando el paquete os y bufio
	// Podemos leer datos de mejor manera.
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Menú de Tareas\n",
			"1. Crear Tarea\n",
			"2. Editar Tarea\n",
			"3. Completar Tarea\n",
			"4. Eliminar Tarea\n",
			"5. Salir")

		fmt.Print("Ingrese el opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el Nombre de la tarea: ")
			t.Nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la Descripcion de la tarea: ")
			t.Desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea creada correctamente!!")
		case 2:
			var t Tarea
			var indice int
			fmt.Print("Ingrese el Indice de la tarea: ")
			fmt.Scanln(&indice)
			fmt.Print("Ingrese el Nombre de la tarea: ")
			t.Nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la Descripcion de la tarea: ")
			t.Desc, _ = leer.ReadString('\n')
			lista.editarTarea(indice, t)
			fmt.Println("Tarea editada correctamente!!")
		case 3:
			var indice int
			fmt.Print("Ingrese el Indice de la tarea a editar: ")
			fmt.Scanln(&indice)
			lista.completarTarea(indice)
			fmt.Printf("Tarea %d completada!!\n", indice)
		case 4:
			var indice int
			fmt.Print("Ingrese el Indice de la tarea a eliminar: ")
			fmt.Scanln(&indice)
			lista.eliminarTarea(indice)
			fmt.Printf("Tarea %d eliminada exitosamente!!\n", indice)
		case 5:
			fmt.Println("Saliendo del Programa...")
			return
		default:
			fmt.Println("Opcion incorrecta")
		}

		// Listamos las tareas
		listarTareas(lista)
	}

}

// Método para crear o agregar tarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// Marcar tarea como completada
func (l *ListaTareas) completarTarea(indice int) {
	l.tareas[indice].Completado = true
}

// Editamos la tarea
func (l *ListaTareas) editarTarea(indice int, t Tarea) {
	l.tareas[indice] = t
}

// Eliminamos la tarea, creando un nuevo slice de tareas sin el indice por argumento.
func (l *ListaTareas) eliminarTarea(indice int) {
	l.tareas = append(l.tareas[:indice], l.tareas[indice+1:]...)
}

func listarTareas(l ListaTareas) {
	// Listar tareas
	fmt.Println("Lista de Tareas")
	fmt.Println("###################################################")
	for i, t := range l.tareas {
		fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.Nombre, t.Desc, t.Completado)
	}
	fmt.Println("###################################################")
}
