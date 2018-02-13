This sample showed simple web app with RAM memory leaks.

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
1000000 | How many attempts benchmark does
1597 ns/op |  How long it takes on 1 operation in average 
1440 B/op | How many memory it consumption on 1 operation in average
11 allocs/op | How many memory allocation does on 1 operation