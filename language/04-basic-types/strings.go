package main

import "fmt"

func strings() {
	// Raw string literals are character sequences between backquotes/backticks
	var a string = `This is a "test"`
	var b string = `This is
	"another test"
	and
	another
	`

	fmt.Println(a, b)
}
