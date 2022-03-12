package main

import "fmt"

// infer declaration
var url = "http://someaddress.com" // warning: unused variable at package scope
// every statement at package scope should start with a keyword
// so short declaration operator is not allowed

func declareStrings() {
	// blank identifier
	_ = url // fix unused variable warning from the code above

	// if you don't know the initial value
	// use the long form variable declaration
	// ! advice: only if you don't know the initial value
	var endpoint string

	// multi-line multiple declaration
	var (
		speed        int
		velocity     float32
		acceleration int

		distance int
		gas      int
	)

	// single-line multiple declaration
	// in other words: parallel declarations
	var brand, model string = "Toyota", "Prius"

	// type inference
	var shape, engine = "Station", "V8" // inferred as strings

	// multiple short declaration
	// with different types
	age, gender := 32, "Male"

	// re-declaration
	// ! only valid syntax because one of the variables is a new variable
	var year int
	year, testable := 34, true

	// print values
	fmt.Println(endpoint)
	fmt.Println(speed, velocity, acceleration)
	fmt.Println(distance, gas)
	fmt.Println(brand, model)
	fmt.Println(shape, engine)
	fmt.Println(age, gender)
	fmt.Println(year, testable)
}
