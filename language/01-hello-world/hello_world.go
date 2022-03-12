// first statement of a .go file is a `package declaration`
// or `package clause`
// every go file is part of a package
// a package can be 1 or more files

// package declaration
package main // compile this package to an executable

import "fmt" // import the fmt standard library package

// package main requires a function main
// declare the function named main
// execution starts in function main
func main() {
	// code block

	fmt.Println("Hello world")
}

/*

Every Go file is organized and structured into three sections
- package declaration
- package imports
- function declarations

*/
