package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Fred"
	fileName := "name.text"

	// this is similar to Println // but only takes 1 argument
	io.WriteString(os.Stdout, name)

	// create a new file in the same folder as the source
	// return a pointer to a file
	f, err := os.Create(fileName)
	if err != nil {
		// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
		log.Fatalln("Error creating the file:", fileName, err)
	}

	// a defer statement defers the execution of a function until the surrounding function returns
	// close the file as soon as the function returns
	defer f.Close()

	// io.Copy(f, name) // name is not a reader
	// convert string name to a reader with string.NewReader()
	io.Copy(f, strings.NewReader(name))
}
