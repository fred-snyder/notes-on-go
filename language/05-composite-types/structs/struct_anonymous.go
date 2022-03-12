package main

import "fmt"

// struct example without defining the struct type
func structsAnon() {
	data := struct {
		field1 int    // can also be a struct
		field2 string // can also be a map
	}{
		6,
		"some Text",
	}

	// this can be useful when you want to combine data with the goal to pass the data to a function
	// for example // when passing data to a template
	fmt.Println(data)
}
