package main

import "fmt"

func assignmentOps() {
	area := 10

	// the same
	area = area * 2
	area *= 2  // multiply and assign
	area /= 2  // divide and assign
	area -= 10 // subtract and assign
	area += 10 // add and assign
	area %= 2  // remainder

	// ! area += 10.2
	// ^-------this doesn't work because area is defined as an int
	// ! area += int(10.2)
	// ^-------this doesn't work because type conversions don't work on numeric literals

	n := 10.6            // first store the value // truncation towards zero
	area = area + int(n) // then convert to the type
	area += 10.0         // this does work because 10.0 will be converted to 10

	fmt.Println("The area is: ", area)
}
