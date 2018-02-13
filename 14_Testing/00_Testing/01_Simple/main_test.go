package mylib

import "testing"

func Test_Fibonacci(t *testing.T) {

	cases := []struct{ a, result uint64 }{
		{1, 1},
		{2, 1},
		{3, 2},
		{40, 102334155},
	}

	for _, c := range cases {
		result := Fibonacci(c.a)

		if result != c.result {
			t.Log("result should be "+string(c.result)+", but got ", result)
			t.Fail()
		}
	}

}
