package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	pid := "288"

	cmd := exec.Command("ps", "-p", pid, "-o", "pid,time,%cpu,%mem,rss")

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	data := printOutput(output)

	dataFormatted := strings.Split(data, "\n")

	fmt.Println(dataFormatted)
}

func printOutput(outs []byte) string {
	result := ""
	if len(outs) > 0 {
		result += string(outs)
	}
	return result
}
