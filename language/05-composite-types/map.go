package main

import "fmt"

/*
A map is Go's implementation of a hashtable

https://junminlee3.medium.com/hash-tables-animations-that-will-make-you-understand-how-they-work-d1bcc850ba71

key-value pairs
needs a unique key (example: "passport number": "name")
the type of the keys and values must always be the same
*/

func maps() {
	fmt.Println("========= maps")

	// declare a map variable
	var someMap map[string]int

	// create the actual map
	var someOtherMap = make(map[string]int)
	someOtherMap = map[string]int{
		"key": 1,
	}

	_ = someOtherMap // get rid of unused variable warning

	// map literal // declare and assign values
	personMap := map[string]int{
		"Fred": 456,
		"Maus": 457,
	}

	// add a new key/value pair
	personMap["Thijs"] = 5668

	// delete a key/value pair
	delete(personMap, "Fred")

	// return a boolean if key is present in map
	val, p := personMap["Thijs"]
	fmt.Println("Thijs in personMap:", p, "Value:", val)

	fmt.Println(someMap)
	fmt.Println(personMap)
	fmt.Println("Length:", len(personMap)) // lenght of the map
	fmt.Println(personMap["Maus"])

	// iterate over a map
	for key, val := range personMap {
		fmt.Println("For-loop:", key, val)
	}
}

func moreMaps() {
	fmt.Println("=============== moreMaps")

	// map literal
	emails := map[string]string{
		"Fred":  "fred@gmail.com",
		"Marco": "marco@gmail.com",
	}
	fmt.Println(emails)

	// make a map
	var countries = make(map[string]int)
	_ = countries

	// make a map
	currencies := make(map[string]int)
	_ = currencies

	// add new key/value pairs
	// TODO: question: can you say instead "add properties" to this map?
	currencies["Germany"] = 1
	currencies["France"] = 2

	fmt.Println(currencies)
	fmt.Println(currencies["Germany"])

	// delete a kv-pair
	currencies["Delete"] = 100 // add
	delete(currencies, "Delete")
	fmt.Println(currencies)
}
