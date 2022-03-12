# Blank identifier

> "eat data and send it to a void"

> "keep the compiler happy"

You can't have unused variables in your code. The `_` (underscore) in Golang is known as the Blank Identifier. https://www.geeksforgeeks.org/what-is-blank-identifierunderscore-in-golang/

```go
func hello() {
	var hello = "hello"
	_ = hello // underscore notation

	// Go doesn't accept unused variables
	// using the blank identifier you can assign a variable to a void
	// you can use this for example to demonstrate some code examples
}
```
