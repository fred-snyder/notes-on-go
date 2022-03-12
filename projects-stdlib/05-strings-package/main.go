package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	a := os.Args[1]
	fmt.Println(strings.Repeat("Hello", 5))
	fmt.Println(strings.ToUpper("hello"))

	// small exercise
	rep := strings.Repeat("!", 3)
	fmt.Println(rep + a + rep)
}
