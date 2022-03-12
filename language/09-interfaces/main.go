package main

import (
	"fmt"
	"math"
)

// every type that has a method named area that returns a float64
// implements the interface Shape
type shape interface {
	area() float64
}

type circle struct {
	x, y, radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("======= Interfaces")

	c := circle{2, 2, 3}
	r := rectangle{2, 4}

	fmt.Println(c, r)
	fmt.Println(c.area())
	fmt.Println(r.area())
	fmt.Println(getArea(c))
	fmt.Println(getArea(r))

	// TODO: check this example: better expain how it works
	fmt.Println("======== Empty interface")
	tplFm()
}
