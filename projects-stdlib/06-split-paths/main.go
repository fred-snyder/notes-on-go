package main

import (
	"fmt"
	"path"
)

// path.Split
func main() {
	fmt.Println("Split paths")
	dir, file := path.Split("/path/to/some/filename.text")

	// `dir` has a value of "/path/to/some/"
	// `file` has a value of "filename.txt"
	fmt.Println(dir, file)
}
