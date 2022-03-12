package version

import (
	"fmt"
	"runtime"
)

func Version() {
	// print Go version
	fmt.Println(runtime.Version())
}
