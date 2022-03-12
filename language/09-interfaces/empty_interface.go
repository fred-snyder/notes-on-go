package main

import (
	"fmt"
	"html/template"
)

func printSomething() string {
	return "Something"
}

func tplFm() {
	// type FuncMap map[string]interface{}
	fm := template.FuncMap{
		// this interface is empty: an interface with no methods
		// any type has at least 0 methods/receivers: so every type implements the empty interface
		// therefore: you can pass functions
		"print": printSomething,
	}

	fmt.Println(fm)
}
