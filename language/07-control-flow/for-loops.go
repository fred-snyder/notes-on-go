package main

import "fmt"

func forLoops() {
	fmt.Println("======= forLoops")
	for i := 0; i < 5; i++ {
		fmt.Println("Count", i)
	}

	// create some slice
	someSlice := []int{1, 4, 6, 3, 234, 57, 54}

	// use the range function to iterate over an array or slice
	for i, val := range someSlice {
		fmt.Println("range", i, val)
	}
}

func moreForLoops() {
	fmt.Println("======= moreForLoops")
	// slightly different syntax
	x := 8
	for x < 10 {
		fmt.Println(x)
		x++
	}
}
