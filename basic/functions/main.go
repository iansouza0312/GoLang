package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalcs(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	account := sum(10, 17)
	fmt.Println(account)

	var newFunc = func(text string) string {
		fmt.Println(text)
		return text
	}

	result := newFunc("New text function ...")
	fmt.Println(result)

	sumResult, subResult := mathCalcs(17, 13)
	fmt.Println(sumResult, subResult)
}
