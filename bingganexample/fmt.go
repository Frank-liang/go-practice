package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	var cmd string
	var name string
	var id int
	var line string
	f := bufio.NewReader(os.Stdin)

	//var students map[string]Student
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Printf("binggan 01\njack 02\n")
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			fmt.Printf("add done.\n")
		}
	}
}
