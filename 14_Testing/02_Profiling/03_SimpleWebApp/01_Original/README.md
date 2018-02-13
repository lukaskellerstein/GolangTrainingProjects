This sample showed simple web app with RAM memory leaks.

# Run RAM Profiling

```Shell
go test -bench=. -memprofile=prof.mem
```

# Analyze RAM Profiling

```Shell
go tool pprof --alloc_space 01_Original.test prof.mem
```

- `top20`
- `top --cum`
- `web`

Methods 

- `list ReadAndReplace`