# Initialization

From Wikipedia
> The assignment of an initial value for a data object or variable.

Assigning an initial value to a variable.

> In Go variables are always initialized with a zero-value

```golang
var age int; // declare variable age
fmt.Println(age); // 0 // Go initializes to zero-value automatically

var isAdult bool = true; // declare and initialize

// declare without type
var brand = "BMW" // Go infers the data-type from the initialized value
```

## Short declaration assignment

> You can only use the short declaration operator from within a function.

```golang
func someFunction() {
    // short declaration statement
    country := "Germany"; // declared and initialized
}
```
