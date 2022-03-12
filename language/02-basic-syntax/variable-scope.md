# Scope

In Go variable scoping is done using blocks.
A block is a sequence of declarations and statements within matching curly brackets.

- Universe block. That's all Go source code.
- Package block. All the source code in a particular package.
- File block. All the source code in a single file
- Function/if/for/switch statement blocks

## Lexical scoping

Lexical scope: Inner scope has access to the Outer scope.

Defined inside is "transitive"

> Lexical Scoping defines how variable names are resolved in nested functions: inner functions contain the scope of parent functions even if the parent function has returned.
> The last part: "even if the parent function has returned" is called Closure.
https://stackoverflow.com/questions/1047454/what-is-lexical-scope
