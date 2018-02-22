package main

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	line := "ls | grep file"
	cmds := strings.Split(line, "|")
	s1 := strings.Fields(cmds[0])
	s2 := strings.Fields(cmds[1])

	r, w := io.Pipe()
	cmd1 := exec.Command(s1[0], s1[1:]...)
	cmd2 := exec.Command(s2[0], s2[1:]...)
	cmd1.Stdin = os.Stdin
	cmd1.Stdout = w
	cmd2.Stdin = r
	cmd2.Stdout = os.Stdout
	cmd1.Start()
	cmd2.Start()

	cmd1.Wait()
}
