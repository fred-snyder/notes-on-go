# Working with UTF-8 runes

@TODO: update this file: perhaps include in 04-basic-types

```go
package main

import "fmt"

func main(){
    var ch1 int = '\u0041'
    var ch2 int = '\u03B2'
    var ch3 int = '\U00101234'
    fmt.Printf("%d - %d - %d\n", ch1, ch2, ch3)  // integer
    fmt.Printf("%c - %c - %c\n", ch1, ch2, ch3)  // character
    fmt.Printf("%X - %X - %X\n", ch1, ch2, ch3)  // UTF-8 bytes
    fmt.Printf("%U - %U - %U", ch1, ch2, ch3)    // UTF-8 code point
}
```

```
The byte data type represents ASCII characters while the rune data type represents a more broader set of Unicode characters that are encoded in UTF-8 format.
```
