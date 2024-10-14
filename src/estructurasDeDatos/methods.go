package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	p := Persona{
		"Michell",
		24,
		"correo@correo.com",
	}
	// Accedemos a saludo mediante p (struct persona)
	p.saludo()
}

// Es una función de la estructura (Humano)
// No se puede acceder directamente.
// Debemmos acceder por medio de p.saludo() en el ejemplo de arriba.
func (p *Persona) saludo() {
	fmt.Println("Hola, mi nombre es", p.nombre)
}
