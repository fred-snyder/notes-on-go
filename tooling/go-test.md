# How to do testing with Go

```bash
# go test will download the package/dependencies and the indirect dependencies
# and then test the go application
go test

# as a result
# Go test now creates a new file `go.sum`
```

> go.sum (for "checksum") contains the cryptographic hashes of the specific module versions

> If you import a module, and go.mod is not initialized > go test will not work
