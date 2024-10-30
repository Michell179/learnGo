package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	myBook2 := book.NewBook("How to Learn Java", "Michell", 260)
	//myBook2.ToString()

	myTextBook := book.NewTextBook("Maths", "Teacher Morgan", 314, "I.E Cali", "2")
	//myTextBook.ToString()

	fmt.Println("Book")
	book.Print(myBook2)
	book.Print(myTextBook)

	fmt.Println("Animals")

	perro := animal.Perro{Nombre: "Haru"}
	gato := animal.Gato{Nombre: "Maven"}

	animal.ReproducirSonido(&perro)
	animal.ReproducirSonido(&gato)

	animales := []animal.Animal{
		&animal.Gato{Nombre: "Maven"},
		&animal.Perro{Nombre: "Haru"},
		&animal.Gato{Nombre: "Hyhal"},
		&animal.Perro{Nombre: "Magnus"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}

}
