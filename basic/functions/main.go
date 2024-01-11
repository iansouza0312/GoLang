package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	result := sum(10, 17)
	fmt.Println(result)
}