// executable package
// only for creating executables
package main

import (
	"fmt"

	// the go.mod file determines the import path
	// in this case: github.com/fred-snyder/kbsb-golang-journey
	"github.com/fred-snyder/kbsb-golang-journey/language/02-basic-structure/package_examples/library_example/libpack"
)

// only one function main can exist in your main package
func main() {
	fmt.Println("main.go")
	libpack.SomeFunction()

	// functions from files within the same package
	// are directly accesible
	printTest()
	printError()
}
