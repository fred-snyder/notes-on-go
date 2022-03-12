package main

import "fmt"

// range function examples
func rangeMaps() {
	fmt.Println("========= rangeMaps")

	zipcodes := map[string]int{
		"Amsterdam": 1000,
		"Rotterdam": 2000,
		"The Hague": 3000,
		"Groningen": 8000,
	}

	for k, v := range zipcodes {
		fmt.Println(k, v)
		fmt.Printf("key: %s, value: %d\n", k, v)
	}

	// only use the key
	for k := range zipcodes {
		fmt.Printf("key: %s\n", k)
	}

	// or discard the key return value
	for _, v := range zipcodes {
		fmt.Printf("value: %d\n", v)
	}
}
