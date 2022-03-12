package main

import "fmt"

var x int = 2

func sw() {
	caseSyntaxOne()
	caseSyntaxTwo()
}

func caseSyntaxOne() {
	switch x { // x is called the tag
	case 1:
		fmt.Println("Case nr:1")
	case 2:
		fmt.Println("Case nr:2")
	default:
		fmt.Println("No case")
	}
}

func caseSyntaxTwo() string { // don't forget the return type
	switch {
	case x < 3:
		fmt.Println("Case nr:1")
		return "1" //
	case x == 2:
		fmt.Println("Case nr:2")
		return "2"
	default:
		fmt.Println("No case")
		return "def"
	}
}
