package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("grep", "a")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
}
