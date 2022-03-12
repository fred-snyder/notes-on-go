package main

import "fmt"

// printf reference

func main() {
	age, job, country := 100, "developer", "Netherlands"
	fmt.Printf("My age is: %d. And I'm from %[3]v. My job is %[2]v.\n", age, job, country)
}
