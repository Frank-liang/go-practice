package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num int
	fmt.Scan(&num)
	scanner := bufio.NewReader(os.Stdin)
	for j := 0; j < num; j++ {
		var input string
		input, _ = scanner.ReadString('\n')

		for i := 0; i < len(input); i += 2 {
			if input[i] != '\n' {
				fmt.Printf("%c", input[i])
			}
		}
		fmt.Printf(" ")

		for i := 1; i < len(input); i += 2 {
			if input[i] != '\n' {
				fmt.Printf("%c", input[i])
			}
		}
		fmt.Printf("\n")
	}
}
