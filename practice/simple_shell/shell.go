package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[root@%s]$ ", host)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print(prompt)
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		args := strings.Fields(line)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
