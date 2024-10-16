package main

import (
	"log"
	"os"
)

func main() {
	// Centralizar todos los logs en un solo archivo.
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Agregamos la salida de los logs para ser almacenados en el archivo info.log
	log.SetOutput(file)

	// Lo que se ejecute despues de esto agrega el prefijo main() a todos nuestros logs
	log.SetPrefix("main(): ")

	// log.Fatal, terminará con la ejecución del aplicativo, codigo de status 1
	//log.Fatal("Terminando app")

	// log.Panic, termina nuestra app al igual que Fatal, pero con codigo de status 2
	//log.Panic("Soy un panico")

	// Imrpime un registro en consola, con fecha y hora de ejecucion
	log.Println("Soy un registro de log")
	log.Print("Soy un log nuevo.!")

}
