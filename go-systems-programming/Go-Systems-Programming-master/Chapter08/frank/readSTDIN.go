package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := ""
	var f *os.File
	argument := os.Args
	if len(argument) == 1 {
		f = os.Stdin
	} else {
		filename = argument[1]
		fileHander, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening %s: %s", filename, err)
			os.Exit(1)
		}
		f = fileHander
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
