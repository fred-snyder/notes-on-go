package main

import (
	"fmt"
	"os"
)

func main() {
	binary := os.Args[0] // [0] is the program executable

	var name string // define variable name

	if len(os.Args) == 2 {
		name = os.Args[1]
	} else {
		name = "World"
	}

	fmt.Println("Path to executable:", binary)
	fmt.Println("Hello", name)
}
