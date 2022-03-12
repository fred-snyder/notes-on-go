# Declare your own type

```go
// type typeName dataType
type PersonAge int // this is just an alias for int

// declare and initialize
var AgeFred PersonAge = 99

// another example
type Person struct {
    name    string
    age     int
}

// declare and initialize
p1 := Person{name: "Fred", age: 99}
```
