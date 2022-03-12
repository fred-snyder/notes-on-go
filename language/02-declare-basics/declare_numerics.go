package main

import "fmt"

// define a constant
const pi float32 = 3.1415

func declareNumerics() {
	var age int = 4

	fmt.Println(age)         // print name and age
	fmt.Printf("%T \n", age) // print the type of age
	fmt.Println(pi)          // print a float variable

	floatTest := 4.5 // this will automatically cast to a float64
	fmt.Println(floatTest)
}

// TODO: Is "cast" the correct word?
