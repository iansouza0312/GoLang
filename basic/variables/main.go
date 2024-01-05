package main

import "fmt"

func main() {
	var firstVar = "Primeira Variavel"
	fmt.Println(firstVar)

	secondVar := "Segunda Variavel"
	fmt.Println(secondVar)

	var (
		thrirdVar string = "Quarta Variavel"
		fourthVar string = "Quinta Variavel"
	)

	fmt.Println(thrirdVar, fourthVar)

	fivethVar, sixthVar := "Quinta Variavel", "Sexta Variavel"
	fmt.Println(fivethVar, sixthVar)

	const firstConst string = "Nova Constante"
	fmt.Println(firstConst)
}