# Syntax

## Comments

Similar to Golang multi line comments
```go
{{/* comments */}}
```

## Dot

```go
// Output the template input
{{ . }}
```

## Composition

?? What does template composition means
Composition vs. Inheritance
https://en.wikipedia.org/wiki/Composition_over_inheritance

## Variables

Use the dollar notation to use variables inside Go templates.
```go
// Declare and assign the variable
{{ $someVariable := 45 }}
{{ $someString := "Hello world" }}
{{ $someData := . }}
```

```go
// Print the variable
{{ $someVariable }}
```

```go
// Print a struct.field
{{.Name}} - {{.Email}}
{{.name}} - {{.email}} // only exported fields??

```

## Pipelines

```go
// pipe the data from func1 to func2
{{ func1 | func2 }}
```

Example: Pipe data into a function that takes an argument
```go
// the output of someMethod becomes the argument of methodThatTakesAnArg
{{ .someMethod | .methodThatTakesAnArg }}
```

Pipeline a function
```Go
// the output of a pipeline can be captured by a function
{{ $varName := func1 | func2 | func3 }}
```