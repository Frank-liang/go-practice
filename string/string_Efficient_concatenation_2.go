package main

//use the strings.Join function if you have all of the strings ahead of time.

import (
	"fmt"
	"strings"
)

func main() {
	var strs []string

	for i := 0; i < 1000; i++ {
		strs = append(strs, randStrings())
	}
	fmt.Println(strings.Join(strs, ""))
}

func randStrings() string {
	return "abc-123-ABC-"
}
