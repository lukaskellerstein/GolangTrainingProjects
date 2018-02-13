
# Run CPU Profiling

```Shell
go test -bench=. -cpuprofile=prof.cpu
```

# Analyze CPU Profiling

```Shell
go tool pprof 02_Better.test prof.cpu
```

- `top20`
- `top --cum`
- `web`

Methods 

- `list Fibonacci`