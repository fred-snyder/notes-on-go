# Cross compiling

Build executable for Linux on Windows

[https://stackoverflow.com/questions/20829155/how-to-cross-compile-from-windows-to-linux](https://stackoverflow.com/questions/20829155/how-to-cross-compile-from-windows-to-linux)

```bash
# Create an OSX executable:
GOOS=darwin GOARCH=386 go build

# Create a Windows executable:
GOOS=windows GOARCH=386 go build

# Create a Linux executable
GOOS=linux GOARCH=arm GOARM=7 go build

# On Windows CMD or PowerShell
cmd /c "set GOOS=darwin GOARCH=386 && go build" # TODO: test this

# Target Webassembly
GOOS=js GOARCH=wasm go build -o fileName.wasm
```
