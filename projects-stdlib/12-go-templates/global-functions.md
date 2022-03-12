# Global functions

https://pkg.go.dev/text/template#hdr-Functions

## Function overview
```
- html
- index
- range
- len
- js
- printf
- not / and / or
- eq / lt / gt
```

## if
```go
{{ if and var1 var2 }}
```

## range
```go
// range over slice/array
{{ range $index, $element := data }} // index-element are key-value of the slice

// range over a pipeline
{{ range $index, $element := func1 | func2 | func3 }}
```