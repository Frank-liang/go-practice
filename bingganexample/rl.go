package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.New(">>> ")
	if err != nil {
		panic(err)
	}
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		fmt.Println(line)
	}
}
