package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	building := "Amsterdam Rijks Museum"

	// len function return the number of bytes in a string value
	// this may not be the same as the number of characters

	// unicode characters are represented in runes ("...a Unicode code point, or in Go language, a rune...")
	// and can occupy multiple runes in memory
	fmt.Println(len(building))

	// the RuneCountInstring package returns the number of runes?
	// so the actual number of characters
	fmt.Println(utf8.RuneCountInString(building))
}

// unicode characters can take 1-4 bytes/runes
