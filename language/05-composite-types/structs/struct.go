package main

import "fmt"

// a struct is a composite data type (or, aggregate data type)

type person struct {
	name     string // name property or name field
	hometown string
	zipcode  int
}

// TODO: question: is using a capitalized struct exported?
type SportClub struct {
	name  string
	sport string
	costs int
}

func structs() {
	var person1 person
	person1.name = "Fred"
	person1.hometown = "Berlin"
	person1.zipcode = 4578

	fmt.Println(person1)
	fmt.Println(person1.hometown)

	// initializes fields to zero
	person2 := new(person)
	fmt.Println(person2) // it returns an ampersand?

	// initialize with a struct literal
	person3 := person{
		name:     "Thijs",
		hometown: "New York",
		zipcode:  556874,
	}

	// on one line
	person4 := person{name: "Thijs", hometown: "New York", zipcode: 556874}

	person3.name = "Fred" // change name
	fmt.Println(person3)
	fmt.Println(person4.zipcode)

	// =====================

	// shortform without the property
	clubA := SportClub{"Ajax", "Football", 1200}
	fmt.Println(clubA)

	// another example
	clubR := SportClub{}
	clubR.name = "Dojo"
	clubR.sport = "Boxing"
	clubR.costs = 1000
}
