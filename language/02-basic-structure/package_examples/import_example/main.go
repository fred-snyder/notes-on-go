package main

import (
	"fmt"
	"math" // Vscode auto imports packages for you

	// the go.mod in the workspace root points to
	// github.com/fred-snyder/kbsb-golang-journey
	// so the import for this package should be:
	"github.com/fred-snyder/kbsb-golang-journey/language/02-basic-structure/package_examples/import_example/sayhello"
)

func main() {
	fmt.Println("Test math imports")
	fmt.Println(math.Floor(5.7))
	fmt.Println(math.Ceil(5.3))
	fmt.Println(math.Sqrt(5.3))

	fmt.Println(
		sayhello.Hello("some name"),
	)
}
