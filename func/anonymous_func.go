package main

import (
	"fmt"
	"strings"
)

func toupper(s string) string {
	return strings.Map(func(r rune) rune {
		return r - 32
	}, s)
}

func main() {
	fmt.Println(toupper("hello"))
}
