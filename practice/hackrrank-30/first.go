package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanln("%s", &input)
	var sL []string
	for _, ch := range input {
		str := string(ch)
		sL = append(sL, str)
	}
	fmt.Println("Hello, World.")
	fmt.Println(strings.Join(sL, ""))
}
