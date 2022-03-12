package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Int()   // generates a random number
	b := rand.Intn(8) // generates a random number in [0, n)
	fmt.Printf("a is: %d\n", a)
	fmt.Printf("b is: %d\n", b)
}
