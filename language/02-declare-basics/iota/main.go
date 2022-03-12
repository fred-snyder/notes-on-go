package main

import "fmt"

const (
	NETHERLANDS = iota
	GERMANY     = iota
	FRANCE      = iota
)

// The first use of iota gives 0.
// Whenever iota is used again on a new line,
// its value is incremented by 1;

// A new const block or declaration initializes iota back to 0.
const (
	POLAND = iota
	ITALY
)

// You can give enumeration a type name.

type gun int

const (
	DESERT_EAGLE gun = iota + 3 // start at 3
	MAGNUM
)

func main() {
	fmt.Println(NETHERLANDS, GERMANY, FRANCE)
	fmt.Println(POLAND, ITALY) // restarts at 0
	fmt.Printf("Type is: %T\nValue is: %d\n", DESERT_EAGLE, DESERT_EAGLE)
	fmt.Printf("Type is: %T\nValue is: %d\n", MAGNUM, MAGNUM)
}
