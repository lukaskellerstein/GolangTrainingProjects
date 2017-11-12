package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// Create an *exec.Cmd
	cmd := exec.Command("ps", "-e")

	// Combine stdout and stderr
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output) // => go version go1.3 darwin/amd64

}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	result = ""
	if len(outs) > 0 {
		result += string(outs)
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
