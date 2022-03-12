package main

import (
	"fmt"
	"os"
	"strconv"
)

/*

The os.Args variable hold the command-line arguments, starting with the program name.
Args[0] // the path to the executable
Args[1] // first command line argument
Args[2] // second command line argument

*/

func main() {
	// os.Args is slice of strings
	var args []string = os.Args

	fmt.Printf("%v\n", args) // print the value of os.Args
	fmt.Println(len(args))   // print the length of os.Args

	// store the (second) commandline arg
	arg := os.Args[1]

	// os.Args stores the values as strings
	// convert the values to float with strconv
	// ParseFLoat returns the converted value as a float
	// and it returns an error
	feet, err := strconv.ParseFloat(arg, 32)

	// ! you can also discard the error value with the blank identifier
	// feet, _ := strconv.ParseFloat(args, 32)

	fmt.Println(feet)
	fmt.Println("Error message: ", err)
	fmt.Printf("%T\n", feet)

	// another strconv example:
	// strconv.ParseFloat("4.6", 64) // 3.5 // float64
}
