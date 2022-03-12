package main

// type-casting
// Go does not allow implicit conversion: Go never converts by itself.

import (
	"fmt"
	"strconv"
)

func main() {
	var acceleration = 2.4 // in m/s^2
	var distance = 30      // in meters
	var speed = 50         // in km/hr

	// convert types
	fmt.Println(float64(distance))

	// float32() to convert to float32
	// int8() to convert to an int8
	// int() to convert to int (for example from int32 to int)
	// etc

	// convert int to string
	fmt.Println(strconv.FormatInt(int64(speed), 10))
	fmt.Println(strconv.Itoa(speed)) // same as above // Itoa > Integer to a (strconv)

	// fmt.Println(string(64))
	// ^------- This actually does not print a string of the integer 64,
	// but instead prints the string associated with the UTF-8 rune of `64`
	// which is the `@` character

	_ = acceleration

	// convert Boolean type to string
	fmt.Println(strconv.FormatBool(true) + strconv.FormatBool(false))

	// convert float to an int
	someFloat := 5.2
	fmt.Println(int(someFloat))
}
