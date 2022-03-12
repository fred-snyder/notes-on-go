package main

import (
	"fmt"
)

func typeSw() {
	// this type switch works because you pass it an empty interface.
	// an empty interface accepts any type
	whatAmI := func(i interface{}) {
		switch t := i.(type) { // TODO: how can you explain this line?
		// It calls an empty function???
		// or it calls type on the empty interface?
		// perhaps this is a special kind of statement specific to the type switch
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
