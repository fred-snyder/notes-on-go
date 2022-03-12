package main

import "fmt"

// a function that return data > needs a return data type
func newCard() string {
	return "Five of diamonds"
}

func printCard() {
	fmt.Println("Card")
}

func printNumber() {
	fmt.Println(12)
}

func returnNumber() int {
	return 12
}

func addNumbers(n1 int, n2 int) int {
	return n1 + n2
}
