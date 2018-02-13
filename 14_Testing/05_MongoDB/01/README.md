
# Test basics

Test / Benchmark files

Must be named as `XXX_test.go`, otherwise will not run tests.


Test methods

Must be names as `TestXXXX`, otherwise will not run tests.

Benchmarks methods

Must be names as `BenchmarkXXXX`, otherwise will not run benchmarks.


# Benchmark

`go test -bench=.`

BUT, use this (Correlation between time and memory)

`go test -bench . -benchmem`



# Benchmark + CPU Profiling

Run it

`go test -bench=. -cpuprofile=cpu.out`

or BETTER

`go test -bench . -benchmem -cpuprofile cpu.out`


## Analyze Code

Analyze it

`go tool pprof 01.test cpu.out` ad then you will se "interactive" mode, so you can use some commands like this :

- `top10`
- `top --cum`
- `disasm` - disassembled code - checking runtime functions, slices .. etc


## Execution Tree - SVG Image

Show IMAGE - SVG of execution tree

`go tool pprof 01.test cpu.out`

then in "interactive" mode write 

`web`


## Methods

Show lines of methods and his time ambitiousness
	
`go tool pprof 01.test cpu.out`

then in "interactive" mode write 

`list saveToDB_One`

or 

`list saveToDB_Two`

or 

`list saveToDB_Three`

or 

`list saveToDB_Four`






# Benchmark + Memory Profiling

Run it

`go test -bench=. -memprofile=mem.out`

