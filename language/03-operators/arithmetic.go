package main

import "fmt"

func arithmetic() {
	// int / int = int
	// the expression 9/2 will be evaluated as an int
	var ratio float64 = 9 / 2

	fmt.Println("int/int: ", 9/2)
	fmt.Println("float/int: ", 9.0/2)
	fmt.Println("value of ratio: ", ratio)
	fmt.Printf("type of ratio: %T\n", ratio)

	// float / int = float
	// this expression will result in a float
	ratio = 9.0 / 2
	fmt.Println("value of ratio: ", ratio)
	fmt.Printf("type of ratio: %T\n", ratio)
}

// key take-away:
// an expression with a float operand will evaluate as float
// expressions with int operands will evaluate as int

// ! floats are inaccurate (1.00000000005)
// TODO: check float inaccuracy
