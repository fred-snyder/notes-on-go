# Printf reference

```
\\ // escape slash
\n // new line
\" // escape "
```

```go
// type safety
%s // print string
%d // print decimal (int)
%f // print float

%08d // decimal with leading digits
%04d // decimal with leading digits

%.0f // float // zero digits
%.1f // float // one digit
%.2f // float // two digits
%.3f // float // three digits
%.4f // float // four digits

%t // true of false

// no type safety
%v // print value without typesafety

%T // print type of variable
%[1]T // type of first argument
%[2]v // value of second argument
```
