package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var contactos []Contact
	err := cargarContactsDesdeArchivo(&contactos)
	if err != nil {
		fmt.Println("Error al cargar contactos:", err)
	}

	// Crear instancia de bufio (Lector de datos)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n### Gestor de Contactos ###\n",
			"1. Agregar un Contacto\n",
			"2. Mostrar contactos\n",
			"3. Salir")

		var option int
		fmt.Print("\nIngresa una opción: ")
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion:", err)
			return
		}

		switch option {
		case 1:
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone: ")
			c.Phone, _ = reader.ReadString('\n')
			contactos = append(contactos, c)
			err = guardarContactoEnArchivo(contactos)
			if err != nil {
				fmt.Println("Error al guardar el contacto:", err)
				return
			}
		case 2:
			fmt.Println("\n####### Contactos ########")
			for i, contacto := range contactos {
				fmt.Printf("%d. Nombre: %s Email: %s Phone: %s\n",
					i+1, contacto.Name, contacto.Email, contacto.Phone)
			}
			fmt.Println("##########################")
		case 3:
			return
		default:
			fmt.Println("Opcion Incorrecta.")
		}

	}

}

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func guardarContactoEnArchivo(contacts []Contact) error {
	file, err := os.Create("contacto.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Creamos un encoder de Json
	encoder := json.NewEncoder(file)
	// Le enviamos el slice de contacts y se crea un encode en el formato Json
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func cargarContactsDesdeArchivo(contacts *[]Contact) error {
	file, err := os.Open("contacto.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}
