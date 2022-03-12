package main

import (
	"fmt"
	"math"
)

func ints() {
	var height = 255
	height += 1
	fmt.Println(height) // 0 // wraps around the max int value

	fmt.Println(math.MaxInt8)     // max value of a int8
	fmt.Println(math.MinInt8)     // min value of a int8
	fmt.Println(math.MinInt8 - 1) // wraps around to it's max value

	fmt.Println(math.MaxFloat32)
}
