This sample showed simple app which needs File and method with CPU leaks.

# Run CPU Profiling

```Shell
go test -bench=. -cpuprofile=prof.cpu
```

# Analyze CPU Profiling

```Shell
go tool pprof 01_Original.test prof.cpu
```

- `top20`
- `top --cum`
- `web`

Methods 

- `list ReadAndReplace`