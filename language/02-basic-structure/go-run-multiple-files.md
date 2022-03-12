## Multiple files

If you structure go code over multiple files make sure you they are part of the same package.

Say you have two files:
- main.go
- functions.go

If you want to run this program:
`go run *.go` or `go run main.go functions.go`

> On Windows `go run *.go` doesn't work. @TODO: does it work in the Windows Powershell?
