package main

import (
	"strings"
	"testing"
)

func Benchmark_ReadAndReplace(b *testing.B) {

	for index := 0; index < b.N; index++ {

		ReadAndReplace(strings.NewReader("asdf|asdf;asdf:"))

	}

}
