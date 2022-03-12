# Literals

A literal is a notation for representing a fixed value in source code. 

(Think of English word literally. If you mean something literally: so in a literal sense or manner.)

Examples:

## A string literal

```go
// Hello world written in quotes is a string literal.
var hello = "Hello world";

// or
fmt.Println("") // empty string
fmt.Println("Hi there")
```


## An integer literal

```go
// the number 30 is an integer literal
int age = 30;
```

## An object literal

```js
// In Javascript you create objects using the object literal notation
var someObj = {
    prop1: 30,
    prop2: {
        cat: yes,
        do: no
    }
}
```
In Go something similar is a struct literal:

```go
type person struct {
	name     string // name property or name field
	hometown string
	zipcode  int
}

pThijs := person{
    name:     "Thijs",
    hometown: "New York",
    zipcode:  556874,
}
```
