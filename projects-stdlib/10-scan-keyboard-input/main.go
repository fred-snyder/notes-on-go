package main

import "fmt"

// user input is written to a pointer
// returns number of scanned items

func scanInput() {
	var userName string

	fmt.Println("What is your name?")

	// Scan reads keyboard input
	// takes a memory address (in other words: a pointer) as an argument
	fmt.Scan(&userName)
	fmt.Println("Your username is: ", userName)
}

/*

// With Error handling
_, err := fmt.Scan(&userName)
if err != nil {
fmt.Println(err)
}

*/
