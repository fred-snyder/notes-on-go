# Packages

The first line of a Go file starts with a package declaration. This defines which package the file is part of. A package is a bunch of Go files that are part of the same project/app/workspace

> A group of functional similar code which is bundled together. A way to organize or structure code.

> Similar: a library, module, or namespace.

## Two types of packages

Executable
- This package type compiles to an executable file
- Executable packages are always part of the **main** package
    - If you name package main, you get an executable

Reusable
- Every other package names do not compile to an executable
- They are like dependencies
- These packages are like helpers. You can organize your code into separate reusable logic.
    - These packages are meant to built libraries

## More info

A very nice Golang "Hello world!" explainer for new programmers. Explains packages and Go source code structure.
https://junminlee3.medium.com/basic-stuff-most-beginners-miss-learn-how-to-use-golang-with-the-hello-world-code-97eafec58ed1
