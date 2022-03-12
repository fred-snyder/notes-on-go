package main

import "fmt"

// the struct that receives the methods
type Person struct {
	name string
	age  int

	// also valid:
	// name, city, zipcode, country string
}

// methods (value receiver)
func (p Person) returnName() string {
	return p.name
}

// methods (value receiver)
// TODO: question: this age function is perhaps not very useful
// because it doesn't change the underlying struct
// TODO: test the code in the main function
func (p Person) returnAge() int {
	age := p.age + 1
	return age
}

// methods (pointer receiver)
func (p *Person) increaseAge() {
	p.age++
}

func main() {
	fred := Person{name: "Fred", age: 36}
	fmt.Println(fred)

	fred.returnAge()
	fmt.Println(fred)

	age := fred.returnAge()
	fmt.Println(age)

	fred.increaseAge()
	fmt.Println(fred)

	fred.returnName()
}
