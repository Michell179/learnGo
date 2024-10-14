package main

import (
	"fmt"
	"math/rand"
)

const rangoValor int = 100
const limit int = 5

func main() {
	jugar()
}

func generarRandom(x int) int {
	return rand.Intn(x)
}

func jugar() {
	var datoUsuario int
	intentos := 0
	numeroRandom := generarRandom(rangoValor)

	fmt.Println("Bienvenido al Juego del Genio MavenCat")
	fmt.Println("El Juego consiste en que se va a generar un número aleatorio del 0 al 100")
	fmt.Println("el cual deberás adivinar en tan solo 5 intentos, de lo contrario")
	fmt.Println("obedecerás las órdenes del señor Gato MavenCat Supremo.")
	fmt.Println("Que inicien los MAVEN JUEGOS!!")
	for datoUsuario != numeroRandom {
		fmt.Print("\nIngresa un número del 0 al 100: ")
		fmt.Scanln(&datoUsuario)
		intentos++

		if numeroRandom < datoUsuario {
			fmt.Println("El numero aleatorio es menor")
		} else if numeroRandom > datoUsuario {
			fmt.Println("El numero aleatorio es mayor")
		} else {
			fmt.Println("GANASTEE !!!!")
			fmt.Println("Numero ganador:", numeroRandom)
			fmt.Printf("Usaste %d intentos.\n", intentos)
			jugarDeNuevo()
		}

		if limit == intentos {
			fmt.Println("PERDISTEEEE !!! el numero era:", numeroRandom)
			fmt.Printf("Llegaste al limite %d\n", limit)
			fmt.Println("Ahora le debes tú alma al señor Gato Supremo MavenCat!!!")
			break
		}
	}
	jugarDeNuevo()
}

func jugarDeNuevo() {
	var opciones string
	fmt.Println("Quieres jugar nuevamente ? (S/N)")
	fmt.Scanln(&opciones)

	switch opciones {
	case "S":
		jugar()
	case "N":
		fmt.Println("Gracias por Jugar al MavenCat Supremos")
	default:
		fmt.Println("Opcion no valida, por favor intenta de nuevo.")
		jugarDeNuevo()
	}
}
