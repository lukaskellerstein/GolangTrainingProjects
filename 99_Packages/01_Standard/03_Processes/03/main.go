package main

import (
	"fmt"
	"strconv"

	"github.com/struCoder/pidusage"
)

func main() {

	sysInfo, err := pidusage.GetStat(5240)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("CPU : " + strconv.FormatFloat(sysInfo.CPU, 'E', 1, 64))
	fmt.Println("Memory : " + strconv.FormatFloat(sysInfo.Memory, 'E', 1, 64))

	fmt.Println(sysInfo)
}
