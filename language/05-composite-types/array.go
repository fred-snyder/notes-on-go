package main

import (
	"fmt"
	"reflect"
)

func basicArray() {
	// long form
	var cities [3]string
	cities[0] = "Amsterdam"
	cities[1] = "London"
	cities[2] = "Berlin"
	fmt.Println(cities)

	// short form
	countries := [3]string{"Netherlands", "Germany", "United Kingdom"}
	fmt.Println(countries)
}

func arrays() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	arr1Slice := arr1[2:3]
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{55, 56, 57} // infers size from declaration
	var arr4 [3][3]int           // a 2-dimensional array

	// this is the slice syntax // check out `slice.go`
	slice1 := []int{1, 4, 5, 6}
	_ = slice1 // get rid if unused variable warning

	fmt.Println(arr1)
	fmt.Println(arr1[2:3])                 // slice
	fmt.Println(arr1Slice)                 // slice
	fmt.Println(reflect.TypeOf(arr1Slice)) // TypeOf

	fmt.Println("\nArray functions\n====")
	fmt.Println("Length of an array: ", len(arr1))         // returns the length
	fmt.Println("Capacity of the array: ", cap(arr1[0:1])) // returns the capacity

	fmt.Println(arr2)
	fmt.Println(arr2[2:4]) // slice
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr4[:2]) // slice

	// loop over the items of an array
	for i, val := range arr1 {
		fmt.Printf("index: %v, value: %v \n", i, val)
	}
}
