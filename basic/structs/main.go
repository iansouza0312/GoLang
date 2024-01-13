package main

import "fmt"

type user struct {
	name   string
	age    uint8
	adress adress
}

type adress struct {
	street string
	number uint8
}

func main() {
	fmt.Println("Structs - Basic Concepts of Go")

	var u1 user
	u1.name = "Ian Souza"
	u1.age = 20
	fmt.Println(u1)

	fakeAdress := adress{"Imaginary street", 100}

	u2 := user{"Ana", 24, fakeAdress}
	fmt.Println(u2)

	u3 := user{name: "Joao"}
	fmt.Println(u3)
}
