# Data types

A statically typed language prevents runtime type-bugs because of compile-time type checking.

## Resources

https://go101.org/article/basic-types-and-value-literals.html
https://go.dev/ref/spec#Types
https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html

## Pre-declared types

Pre-declared so a type declared and built-in by Go. Elementary (or primitive) types.

### Numeric types

```plain
int (depends on machine/CPU architecture)
int8
int16
int32 (rune)
int64

uint (on a 64-bit CPU, allocates 64-bits of memory)
uint8 (unsigned integer 8-bits/1-byte)
uint16
uint32
uint64

float32
float64 (default for floats on a 64-bit architecture)
complex64 (two float32)
complex128 (two float64)
```

### Other types

```plain
bool (boolean, takes 1 byte of memory)
string (string, an array of byte values. Or `[]byte` a slice of bytes)
rune (alias for int, architecture dependent)
byte (an alias for uint8)
```
