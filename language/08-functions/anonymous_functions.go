package main

import "fmt"

func anonFunc() {

	func() int {
		return 2
	}() // define and call an anonymous function
}

func foo() {
	// anonymous function
	var bar = func() {
		fmt.Println("foobar")
	}

	bar() // call bar
}
