package main

import (
	"fmt"
	"runtime"
)

func main() {

	/*FORMA 1*/
	//os := runtime.GOOS
	//
	//switch os {
	//case "windows":
	//	fmt.Println("Windows")
	//case "linux":
	//	fmt.Println("Linux")
	//default:
	//	fmt.Println("MAC OS")
	//}

	/*FORMA 2
	de usar los Switchs al igual que ifs podemos declarar variables en su seccion de evaluacion*/

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Darwin")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s\n", os)
	}

}
