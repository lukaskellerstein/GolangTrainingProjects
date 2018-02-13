
# Naming of test files

Must be named as `XXX_test.go`, otherwise will not run tests.

# Run the test

`go test`

# Run the Benchmark

`go test -bench=.`

or 

`go test -run=x -bench=.`

## Benchmark with saving into file 

`go test -run=xxx -bench=. | tee benchmark0`

# Run the CPU Profiling

Run it

`go test -run=^$ -bench=. -cpuprofile=cpu.out`

Analyze it

`go tool pprof 04_Optimization.test cpu.out` ad then you will se "interactive" mode, so you can use some commands like this :

- `top20`
- `top --cum`

Show IMAGE - SVG of execution tree

`go tool pprof 04_Optimization.test cpu.out`

then in "interactive" mode write `web`
