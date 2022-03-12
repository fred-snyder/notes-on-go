package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)

	// package functions
	printCard()
	printNumber()
	returnNumber()

	// print the return value
	fmt.Println(addNumbers(1, 2))

	// anonymous functions
	fmt.Println("========== anonymous functions")
	anonFunc()
	foo()
}
