package main

// - [ ] TODO: Fix this code, create more examples

// Tech and Beyond With Moss - Function Literals and Closures in Golang - Part 1
// https://www.youtube.com/watch?v=EQAM0aWX2ck

// Corey Schafer - Programming Terms: Closures - How to Use Them and Why They Are Useful
// https://www.youtube.com/watch?v=swU3c34d2NQ

// TODO: question: how does a closure work in Javascript?

func closureSumInt() func(int) int {
	sum := 0
	return func(x int) int {
		sum = sum + x
		return sum
	}
}

func closure() {
	v := closureSumInt()
	for i := 0; 0 < 10; i++ {
		v(i)
	}
}

func main() {
	closure()
}
