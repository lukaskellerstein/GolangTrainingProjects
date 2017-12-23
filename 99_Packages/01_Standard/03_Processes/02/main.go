package main

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func main() {

	// pid := flag.Int("p", 5240, "pid")
	// p, _ := process.NewProcess(int32(*pid))
	// myproc := NewProcess(p)

	// cpupencent := myproc.CPUPercent()
	// fmt.Println(cpupencent)

	// pi, err := p.ProcInfo()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(pi)

	p2, _ := process.NewProcess(9362)

	val, _ := p2.Status()
	fmt.Println(val)
}
