package main

import (
	"fmt"
)

// TIP: In a possible reduce variable in a package scope.

// Package scope level
var z = 52

// Variable typed
var variableTyped int = 23
var variableTyped2 int

// practice for default values in variables.

func main() {
	name := "Michell"
	lastName := "Arias"

	// Save Response and possible errors.
	response, err := fmt.Println("Hello World! ", name, lastName)

	// Print response and error.
	_, _ = fmt.Println(response, err)

	// Short Declarative Operator
	//... interface{} all data types implements this.
	// Statement :=
	x := 42
	// Expression
	y := 2 + 4
	fmt.Println(x, y)
	numberZ()

	// Zero default value int variableTyped2
	variableTyped2 = 2

	fmt.Println("VariablTyped", variableTyped)
	fmt.Println("VariablTyped2", variableTyped2)

	// Default values for data types
	// false for booleans
	// 0 for integers
	// 0.0 for floats
	// "" empty for strings
	// nil for:
	//        - pointers
	//        - functions
	//        - interfaces
	//        - slices
	//        - channels
	//        - maps

	// practice

}

func numberZ() {
	fmt.Println("Print Z", z)
}
