package mylib

import "testing"

func Benchmark_Fibonnaci(b *testing.B) {

	for index := 0; index < b.N; index++ {
		Fibonacci(40)
	}

}
