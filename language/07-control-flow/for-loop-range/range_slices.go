package main

import "fmt"

func rangeSlices() {
	fmt.Println("========== rangeSlices")

	// a slice of ints
	ints := []int{3, 5, 7, 4, 2, 5, 67, 2, 5}

	// loop over with range function
	// gofmt autoremoves the parenthesis if you write range as function call
	// for i, val := range(ints) {
	for i := range ints {
		fmt.Println("index: ", i) // prints the index
	}

	for i, val := range ints {
		fmt.Println("index: ", i, "value: ", val) // prints both index and value
	}

	sum := 0
	// discard the index with the blank identifier
	for _, val := range ints {
		fmt.Println("value: ", val)
		sum += val
	}

	fmt.Println("sum of ids: ", sum)

}
