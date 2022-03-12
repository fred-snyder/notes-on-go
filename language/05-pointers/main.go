package main

import "fmt"

// pointers
func main() {
	ptr := new(int)   // the new function returns a pointer to a variable
	*ptr = 3          // assigns a value to the pointer
	fmt.Println(ptr)  // prints the address in memory
	fmt.Println(*ptr) // prints the value 3

	// print type of ptr var
	fmt.Printf("%T\n", ptr)

	// uncomment: go warns about:
	// &*x will be simplified to x. It will not copy x.
	// fmt.Println(&*ptr) // prints the address in memory

	// example
	fmt.Println("example: assign underlying pointer values")

	a := new(int)                         // new pointer to an int
	fmt.Printf("data type of a: %T\n", a) // data type is *int
	*a = 3                                // store a value in the memory address of the pointer
	fmt.Println("&a: ", &a)               // print the memory address of a
	fmt.Println("a: ", a)                 // TODO: question: wy aren't the memory address the same?
	b := &a                               // assign the memory address of a to variable b
	c := a
	fmt.Println("a: ", a, "b: ", b, "c: ", c)
	fmt.Println("*a: ", *a, "*b: ", *b, "*c: ", *c)
	fmt.Println("&a: ", &a, "&b: ", &b, "&c: ", &c)

	//example
	fmt.Println("example: change values of underlying pointers")

	// Junmin Lee - Golang pointers explained
	// https://www.youtube.com/watch?v=sTFJtxJXkaY

	// - [ ] TODO: Fix this code
}
