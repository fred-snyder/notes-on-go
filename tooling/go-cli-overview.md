# The Go cli-tool

```bash
# check Go version
go version

# show commandline options
go help
go help <command>
go help env
go help test

# compile and run (but does not build an executable)
go run

# Run the "hello world" Go program
go run main.go

# run all files in the folder
# go run *.go may work better if there are non .go files in the folder
go run .

# go run verbose
go run -x sourceFile.go

# MORE COMMANDS 

# compiles .go files
# do not run
# save to an executable
go build

# build and compile an executable (but does not run automatically)
go build
go build -o fileName # output to fileName

# compiles and executes .go files
# do run
# do not save to an executable
go run
go run *.go # does not work on Windows?
go run file1.go file2.go file3.go # alternatively

# format all the code in each file in the current directory
go fmt
go fmt github.com/... # Format with a wildcard. Formats all files in the github.com folder.
go fmt ../... # or, one folder up and all folders down

# show environment variables
go env
go env GOPATH # only show the gopath envvar

# print Go source code reference documentation
go doc -src fmt
go doc -src fmt Println # for a specific function
go doc runtime version # print information # TODO: Not sure about the function-name? Just guess!

# run any test files associated with the current project
go test
```
