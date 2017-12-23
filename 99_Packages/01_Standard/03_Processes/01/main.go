package main

import (
	"fmt"
)

func main() {
	p, err := NewProcess(5240)
	if err != nil {
		fmt.Println(err)
	}
	pi, err := p.ProcInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pi)
}
