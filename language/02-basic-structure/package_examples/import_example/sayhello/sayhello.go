// example of non main package
// a package is a way to bundle code that shares a function or purpose
// for example: a bunch of functions that can talk to http servers

// by convention: organize packages into unique folders
// all package files should be in the same directory
// all files in the same directory should belong to the same package
package sayhello

// capitalize the first letter of the function to export this function
func Hello(s string) string {
	return "Hello! " + s
}

// TODO: each Go package has it's own scope. TODO: How does this affects your programs?
