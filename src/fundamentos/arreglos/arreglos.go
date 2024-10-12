package main

import "fmt"

func main() {
	/*Arreglos*/
	// Funcionan con espacio estatico.
	var listaPaises = [4]string{"1: Colombia", "2: Brazil", "3: Estados Unidos", "4: Argentina"}

	// Recorriendo arreglo de tipo string.
	fmt.Println("Array Ordenado")
	for i := 0; i < len(listaPaises); i++ {
		fmt.Println(listaPaises[i])
	}

	fmt.Println("\nArray invertido")
	lp := len(listaPaises)
	for i := lp - 1; i >= 0; i-- {
		fmt.Println(listaPaises[i])
	}

	fmt.Println("\nArray por turnos de inicio a extremo")
	k := lp - 1
	for i := 0; i < lp/2; i++ {
		fmt.Println(listaPaises[i])
		fmt.Println(listaPaises[k])
		k--
	}
}
