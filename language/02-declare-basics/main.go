package main // compile this package to an executable

import "fmt" // declare the function named main

// someVar := "Test characters" // not allowed
// ! short declaration doesn't work at package scope
// ! because every statement at package scope has to start with a keyword

// You can declare variables in multiple different ways:
func main() {
	// declare and assign
	var card string = "Ace of spades"

	/*
	 * Use the short declaration operator whenever possible
	 */

	// Shorthand for declaring and initializing a variable.
	// Go automatically infers the data type
	anotherCard := "Ace of diamonds"

	fmt.Println(card)
	fmt.Println(anotherCard)

	// example of a raw multiline string
	s := `This is 
	a multi line 
	string`

	fmt.Println(s)

	// more examples of declarations
	declareStrings()
	declareBooleans()
	declareNumerics()

	// examples of constants
	declareConstants()
}
