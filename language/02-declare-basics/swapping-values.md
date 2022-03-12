# Swap values

```go
/*
 * You can swap values of two variables by using a single statement
 */

func swap() {
	var (
		speed     = 120
		prevSpeed = 80
	)

	prevSpeed = speed
	speed = prevSpeed
	fmt.Println(speed, prevSpeed) // prints 120 120

	speed = 120
	prevSpeed = 80
	// swap
	prevSpeed, speed = speed, prevSpeed
	fmt.Println(speed, prevSpeed) // prints 80 120
}
```