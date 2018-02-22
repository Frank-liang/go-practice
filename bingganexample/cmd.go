package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	out, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	f := bufio.NewReader(out)
	for {
		line, err := f.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
	cmd.Wait()
}
