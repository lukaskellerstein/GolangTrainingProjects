This sample showed simple app which needs File and method with CPU leaks.

# Run benchmark

```Shell
go test -bench=.
```

# Run benchmark with Memory consumption

```Shell
go test -bench=. -benchmem
```

Ex. 

Value | Description
--- | ---
2000000000 | How many attempts benchmark does
0.28 ns/op |  How long it takes on 1 operation in average 
0 B/op | How many memory it consumption on 1 operation in average
0 allocs/op | How many memory allocation does on 1 operation