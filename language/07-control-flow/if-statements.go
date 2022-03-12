package main

import "fmt"

func ifStatements() {
	fmt.Println("====== ifStatements")
	x := 100
	y := 200

	// if (x < y) {}
	// go fmt auto-removes the parenthesis
	if x < y {
		fmt.Println("X is smaller then Y")
	} else if x > y {
		fmt.Println("X is bigger then Y")
	} else {
		fmt.Println("Something else")
	}
}
