package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

func Dump(v interface{}) {
	data, _ := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(data))
}

func CPUCount() int {
	sysCpuTimes, _ := cpu.CPUTimes(true)
	return len(sysCpuTimes)
}

var numCpus = CPUCount()

func Timer() (time uint64) {
	fd, _ := os.Open("/proc/stat")
	defer fd.Close()
	bf := bufio.NewReader(fd)
	line, _, _ := bf.ReadLine()
	fields := strings.Fields(string(line))
	for idx, val := range fields {
		if idx != 0 {
			v, _ := strconv.ParseUint(val, 10, 64)
			time += v
		}
	}
	return time
}

type Process struct {
	proc          *process.Process
	lastProcTimes uint64
	lastSysTimes  uint64
}

func NewProcess(proc *process.Process) *Process {
	return &Process{
		proc: proc,
	}
}

func (p *Process) procTimer() uint64 {
	filename := fmt.Sprintf("/proc/%d/stat", p.proc.Pid)
	data, _ := ioutil.ReadFile(filename)
	rest := string(data[strings.Index(string(data), ")")+2:])
	fields := strings.Fields(rest)
	ut, _ := strconv.ParseUint(fields[11], 10, 64)
	st, _ := strconv.ParseUint(fields[12], 10, 64)
	return ut + st
}

func (p *Process) CPUPercent() float32 {
	st1 := p.lastSysTimes
	pt1 := p.lastProcTimes

	st2 := Timer()
	pt2 := p.procTimer()

	if st1 == 0 {
		p.lastSysTimes = st2
		p.lastProcTimes = pt2
		return 0.0
	}

	deltaProc := pt2 - pt1
	deltaTime := st2 - st1

	p.lastSysTimes = st2
	p.lastProcTimes = pt2
	if deltaTime == 0 {
		return 0.0
	}
	return float32(deltaProc) / float32(deltaTime) * 100.0 * float32(numCpus)
}

// func main() {
// 	log.Println("ncpu:", numCpus)
// 	pid := flag.Int("p", 6233, "pid")
// 	p, _ := process.NewProcess(int32(*pid))
// 	myproc := NewProcess(p)
// 	for i := 0; i < 3; i++ {
// 		cpupencent := myproc.CPUPercent()
// 		log.Println("cpu percent:", cpupencent)
// 		time.Sleep(time.Millisecond * 1000)
// 	}
// }
