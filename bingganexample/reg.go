package main

import (
	"fmt"
	"regexp"
)

var (
	reg = regexp.MustCompile("[a-z]+.*+?")
)

func main() {
	ok := reg.MatchString("abc12345")
	fmt.Println(ok)
	fmt.Println(reg.FindString("abc12345"))
}
