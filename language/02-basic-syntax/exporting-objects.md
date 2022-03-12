# Exporting functions, types, etc

When an identifier starts with an uppercase letter (e.g. AddNumbers), then the "object" with this identifier is visible in code outside the package (and thus available to "importers" of the package). 

It works in a similar way to `export`-ing in for example Javascript.

Identifiers that start with a lowercase letter are not visible outside the package, but they are visible and usable in the whole package.
