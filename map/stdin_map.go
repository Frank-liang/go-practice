package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	seen := make(map[string]bool)
	reader := bufio.NewReader(os.Stdin)
	//scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		//scanner.Scan()
		//line := scanner.Text()
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
}
