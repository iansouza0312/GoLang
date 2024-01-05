package main

import (
	"errors"
	"fmt"
)

func main() {
	var firstNumber int16 = 100
	fmt.Println(firstNumber)

	var secondNumber uint32 = 219
	fmt.Println(secondNumber)

	// alias
	// int32 = rune
	var thirdNumber rune = 53264
	fmt.Println(thirdNumber)

	var fourthNumber byte = 123
	fmt.Println(fourthNumber)

	var realNumber float32 = 234.45
	fmt.Println(realNumber)

	var str string = "SOMERANDOMTEXTWITHPHRASE"
	fmt.Println(str)

	char := 'I'
	fmt.Println(char)

	var text int16
	fmt.Println(text)

	// default = FALSE
	var conditional bool = true
	fmt.Println(conditional)

	var erro error
	fmt.Println(erro)

	var newError error = errors.New("Internal Server Error ...") 
	fmt.Println(newError)
}