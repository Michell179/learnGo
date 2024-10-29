package main

import (
	"fmt"
	"github.com/michell179/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Michell", "Hellen", "Haru"}
	messages, err := greetings.Hellos(names)

	message, err := greetings.Hello("Michell")

	if err != nil {
		log.Fatal(err)
	}

	// Mensaje individual
	fmt.Println(message)

	fmt.Println("\nIterando sobre Mensajes:")

	// Multiples mensajes
	for _, message := range messages {
		fmt.Println(message)
	}
}
