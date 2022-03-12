package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Day())  // 30
	fmt.Println(t.Date()) // 2020 December 30

	// use predefined constants in Format()
	fmt.Println(time.RFC822)            // is always 02 Jan 06 15:04 MST
	fmt.Println(t.Format(time.RFC822))  // Format() takes a predefined format as an argument
	fmt.Println(time.Kitchen)           // is always 3:04PM
	fmt.Println(10 * time.Second)       // 10 seconds
	fmt.Println(t.Format(time.Kitchen)) //
	fmt.Println(t.Format("02-01-2006")) // the 2nd of jan 2006 is always the date you pass format
}

// https://pkg.go.dev/time#Time.Format
