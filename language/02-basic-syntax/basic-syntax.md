# Basic syntax

https://golang.org/ref/spec#Identifiers

- Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

## Keywords and pre-declared identifiers

https://golang.org/ref/spec#Keywords
https://golang.org/ref/spec#Predeclared_identifiers

- Some sequences of characters cannot be used as identifiers because they are used to define program logic or they declare type/constants/etc.
- They are reserved KEYwords

## Declare variable

```go
// declare with the var keyword
// you can declare variables outside a function with the var keyword
// in other words with the var keyword you can declare variables 
// on "package level scope"
var y = 43

// declare variable with identifier `z` of type `int`
// and automatically assign the value 0 
var z int
```

## Zero value

```go
// Zero value of string is ""
// Zero value of int is 0
```
