package main

import "fmt"

// take user input
// multiple by second input
func main() {
	// store value of first integer input
	var first int

	// store value of second integer input
	var second int
	// result

	fmt.Println("First integer: ")
	fmt.Scanln(&first) // we use the & operator to find the address of a variable

	fmt.Println("Second integer: ")
	fmt.Scanln(&second)

	fmt.Println("Result of the multiplication: ")
	fmt.Println(first * second)
}
