package main

import "fmt"

/*
 * The shift operator shifts the left operand by the shift count specified by the right operand,
 * which must be non-negative. If the shift count is negative at run time, a run-time panic occurs.
 * The shift operators implement arithmetic shifts if the left operand is a signed integer and
 * logical shifts if it is an unsigned integer. There is no upper limit on the shift count.
 * Shifts behave as if the left operand is shifted n times by 1 for a shift count of n.
 * As a result, x << 1 is the same as x*2 and x >> 1 is the same as x/2 but truncated towards negative infinity.
 */

func shiftOperator() {
	// TODO: what is the correct explanation for the << operator
	fmt.Println(1 << 1) // times 2 * 1 	// 2
	fmt.Println(1 << 2) // times 2 * 2 	// 4
	fmt.Println(1 << 3) // times 2 * 3 	// 8
	fmt.Println(1 << 4) // times 2 * 4 	// 16
}
