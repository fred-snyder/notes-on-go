# Go get

Since Go 1.18, `go get` will no longer build packages; it will only be used to add, update, or remove dependencies in go.mod. Specifically, go get will always act as if the -d flag were enabled.

```bash
# download the source code of an external package
# aka download dependencies
go help get
go get github.com/adonovan/gopl.io/tree/master/ch1/helloworld # fetch, build, install

go get github.com/gobuffalo/flect # downloads to: $GOPATH/pkg/mod/github.com/gobuffalo

# get a specific version
go get rsc.io/sampler@v1.3.1

# only update go.mod and download source code needed to build packages.
# do not build or install packages
# this is the standard behaviour in Go 1.18 and higher
# so in Go 1.18 the -d flag is redundant
go get -d github.com/gobuffalo/flect
```

```
# https://go.dev/doc/go-get-install-deprecation
go get github.com/go-training/helloworld
go get: installing executables with 'go get' in module mode is deprecated.
	To adjust and download dependencies of the current module, use 'go get -d'.
	To install using requirements of the current module, use 'go install'.
	To install ignoring the current module, use 'go install' with a version,
	like 'go install example.com/cmd@latest'.
```
