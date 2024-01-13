package main

import "fmt"

func main() {
	fmt.Println("Pointers - Basic Concepts of Go")

	var firstVar int = 10
	var secondVar int = firstVar

	fmt.Println(firstVar, secondVar)

	firstVar++
	fmt.Println(firstVar, secondVar)

	// Pointer use the memory adress of some var
	var thirdVar int
	var pointer *int

	thirdVar = 100
	pointer = &thirdVar

	fmt.Println(thirdVar, pointer)
	fmt.Println(thirdVar, *pointer)

	thirdVar = 173
	fmt.Println(thirdVar, pointer)
	fmt.Println(thirdVar, *pointer)
}
