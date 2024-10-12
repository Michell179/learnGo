package main

import (
	"fmt"
	"math"
)

func main() {
	/*int8 int16 int32 int64 => Tipos de enteros*/
	var numeroInteger int // almacena datos dependiento la arquitecutra si es 32 o 64 bits.

	var numeroFlotante32 float32 // almacena datos de punto flotante si es 32 (float32) o 64 (float64).
	var numeroFlotante64 float64 // almacena datos de punto flotante si es 32 (float32) o 64 (float64).

	/*unit8 uint16 uint32 uint64 => Tipos de enteros positivos*/
	var numeroUint uint // almacena solo datos positivos siendo el inicial 0..1...2... etc..

	numeroInteger = math.MaxInt32
	numeroFlotante32 = math.MaxFloat32
	numeroUint = math.MaxUint32

	fmt.Println(numeroInteger, numeroFlotante32, numeroFlotante64, numeroUint)

}
