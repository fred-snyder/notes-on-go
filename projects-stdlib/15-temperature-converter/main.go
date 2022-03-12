package main

import "fmt"

// temperature type aliases
type Celsius float32
type Fahrenheit float32

// convert celsius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {

	// type-casting
	return Fahrenheit((t * 9 / 5) + 32)
}

func main() {
	fmt.Println(toFahrenheit(100), "Fahrenheit")
}
