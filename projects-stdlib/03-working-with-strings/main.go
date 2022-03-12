package main

import "fmt"

func main() {
	var s string

	// string literal // interpreted // use escape sequences for line-breaks and special characters
	s = "Hello world!\nNext line starts here."

	// raw string literal // raw as in not interpreted
	s = `Hello world
	on multiple lines.
	This can go on and on.`

	html := `
	<html>
		<body>
			<h1>Hello world</h1>
		</body>
	</html>
	`

	fmt.Println(s)
	fmt.Println(html)
	fmt.Println("drive:\\path\\to\\folder")
	fmt.Println(`drive:\path\to\folder`)
}
