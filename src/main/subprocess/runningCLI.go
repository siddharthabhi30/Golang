package main

import (
	"fmt"
	"os/exec"
)

func runCli() {
	cmd := exec.Command("ls")
	err := cmd.Run()
	fmt.Println(err)

	out, err := exec.Command("pwd").Output()
	if err == nil {
		fmt.Printf("Location of the output is %s\n", out)
	}
}

func main() {
	runCli()
}
