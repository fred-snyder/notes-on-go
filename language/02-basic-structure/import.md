# Import statements

```go
// Give my current package access to code from package "foo"
import "foo"
```

You can import standard library packages or external packages.

Standard library packages:
https://golang.org/pkg


## Import package with same the name

```go
package main

import (
	// when you want to import a package with a duplicate name
	// add a keyword before the package
    // the package fmt will be imported as someName
	"fmt"
	someName "fmt"
)

func duplicate() {
	fmt.Println()
	someName.Println()
}
```
