# Go modules

Go mod initializes a Go module. It tells the Go compiler "this is the root of my project" and it is used for dependency management. It replaces the `$GOPATH` workflow. (A single path/location where all Go source code was placed.)

- `go.mod` enables dependency tracking for your code.
- `go.sum` (for "checksum") contains the cryptographic hashes of the specific module versions
- `go.mod and go.sum should be checked into Git`


```bash
# enable dependency tracking for your code by creating a go.mod file
# use a domain/URL that you own
# github or a private domain
go mod init <namespace>
go mod init github.com/kaboomshebang/reponame
go mod init kaboomshebang.com/reponame

# removes unused dependencies
go mod tidy
```

> IMPORTANT: `gopls` (VSCode/Editor language tooling) requires a module at the root of your workspace.

## Do you still need the `GOPATH` environment variable?

`GOPATH` still exists. But the "Go workspace" is not really relevant anymore. Instead you can use Go modules.

```bash
# use go env to check the location of $GOPATH
$ go env GOPATH
# for example: /Users/Username/go
```

## Go workspace
- The GOPATH still exists and on Windows points to the go folder in your home directory
- Older versions of Go had a predefined workspace structure for every project and your GOPATH determined the location of that workspace.
- This structure still exist in your default GOPATH folder
    - bin: contains the binary executables.
        - Where `Go install` outputs it's programs
    - src: contains Go source files.
    - pkg: contains Go package archive

## Go modules

>>> After the introduction of Go modules in Go 1.11, you’re no longer required to store your Go code in the Go workspace

In Go version before 1.11 the convention was to store your code in $GOPATH/src/github.com/yourusername/repo. This workflow is modified when Go modules where introduced.

With Go modules the convention is similar in that `go.mod` points to a repo or a URL that is owned by the developer.

## Third-party resources
https://www.honeybadger.io/blog/golang-go-package-management/
https://stackoverflow.com/questions/66356034/what-is-the-difference-between-go-get-command-and-go-mod-download-command

## Official explanations
https://blog.golang.org/using-go-modules
https://golang.org/doc/tutorial/getting-started
    - Point 3. Enable dependency tracking for your code.

https://golang.org/doc/code#Organization
