# go install

```bash
################## installing executables with 'go get' in `module mode` is deprecated.

# compiles and installs a package
# install executable in the $GOPATH\bin folder
go install # install to the ~/go/bin folder
go env GOBIN # The directory where 'go install' will install a command.

################## To install an executable while ignoring the current module, use go install with a version suffix like @v1.2.3 or @latest, as below. When used with a version suffix, go install does not read or update the go.mod file in the current directory or a parent directory.

# download, build and install the application
# go install will NOT add the package as a dependency to go.mod
go install github.com/go-training/helloworld@latest
```

## Add to $PATH

```bash
# add $GOPATH/bin to your $PATH
# so that you can always run your installed executables
# edit ~/.profile
```
