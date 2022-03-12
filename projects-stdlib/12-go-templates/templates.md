# Templates

https://pkg.go.dev/text/template

Templates are a way to merge data with predefined templates. It's about combining predefined generic data with dynamic unique data.
So that a program can generate unique documents from a data set.

## Standard library packages

`text/template`
- The foundation for Go templates

`html/template`
- Extends the text template package with html specific template parsing
- Safety features: Escaping characters, prevent Cross-site-scripting, etc

file-extension convention: `.gohtml`

## Parsing

The definition of parsing:
> Parsing is syntactic analysis of input code into its component parts. In other words the process of reading and analysing stringdata so that it can be processed in a program.
