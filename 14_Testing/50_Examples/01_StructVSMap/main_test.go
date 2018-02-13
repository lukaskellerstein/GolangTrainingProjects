package structvsmap

import "testing"

func Benchmark_AddToMap(b *testing.B) {
	for index := 0; index < b.N; index++ {
		AddToMap()
	}
}

func Benchmark_AddToStuct(b *testing.B) {
	for index := 0; index < b.N; index++ {
		AddToStruct()
	}
}
