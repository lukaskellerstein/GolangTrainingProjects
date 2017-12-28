package main

import (
	"fmt"
)

type TestObject struct {
	a int
	b *int
	// c &int // cant be written this way
}

func (b *TestObject) Print() {
	fmt.Println(b)
	fmt.Println(b.a)
	fmt.Println(b.b)
}

func main() {
	somevalue1 := 1
	somevalue2 := 1
	test1 := TestObject{a: somevalue1, b: &somevalue2}
	test1.Print()
}
