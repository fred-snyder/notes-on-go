package main

import "fmt"

// if you don't specify the length of the array
// the compiler creates the underlying array

func slices() {
	var grades = make([]int, 3)
	fmt.Println(grades)
	fmt.Printf("%T"+"\n", grades) // []int
	// [0 0 0]

	// re-assign a new slice
	grades = []int{1, 2, 3, 4, 5}
	fmt.Println(grades)
	fmt.Printf("%T"+"\n", grades) // []int // [1 2 3 4 5]
}

func nameSlice() {
	nameSlice := []string{"Hans", "John", "James"}
	fmt.Println(nameSlice)

	// print length of array
	fmt.Println(len(nameSlice))
}

func moreSlices() {
	slice1 := []int{2, 4, 6, 7, 9}

	slice2 := make([]int, 10)     // length/capacity
	slice3 := make([]int, 10, 15) // length and capacity

	slice1 = append(slice1, 12)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// the capacity of a slice
	fmt.Println(cap(slice3))
	// Slice: the maximum length the slice can reach when resliced;
}

// a slice points to the underlaying array
// so if you create a: `slice := array[2:3]`
// and then change slice[1] to a new value
// this actually changes array[2]

// in VSCode type slice1.append to autocomplete the code
// into an append() and assignment construction
