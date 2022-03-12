// library package
// a library package is not executable
// a library is importable
// you can import lib packages into other packages
// library packages are created for reusability
package libpack

import "fmt"

func SomeFunction() {
	fmt.Println("Hi, I'm a function")
}

// the filename doesn't matter when importing a function from a package
// because all functions have package scope
// you can access any function from any file in the same package

// nested functions are not allowed
