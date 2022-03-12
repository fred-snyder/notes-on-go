package main

import "fmt"

// these functions have package scope
// and are available everywhere in the main package
// no need to export the function
func printTest() {
	fmt.Println("This is a test.")
}

func printError() {
	fmt.Println("Error!!")
}
