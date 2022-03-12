package main

import "fmt"

func incdec() {
	var n int

	n = n + 1
	fmt.Println(n)

	// increase n by 1
	n += 1
	fmt.Println(n)

	// incdec statement
	n++

	// dec by 1
	n -= 1

	// incdec statement
	n--
	fmt.Println(n)

	// illegal syntax: n = n + n--
	// in Go: incdec statement // not incdec operator
}
