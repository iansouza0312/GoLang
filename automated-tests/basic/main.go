package main

import (
	"basic/adress"
	"fmt"
)

func main() {
	fmt.Println("Basic Tests - Automated Tests with Go")

	adressType := adress.AdressType("Avenue something")
	fmt.Println(adressType)

	adressType2 := adress.AdressType("Highway new orleans")
	fmt.Println(adressType2)
}
