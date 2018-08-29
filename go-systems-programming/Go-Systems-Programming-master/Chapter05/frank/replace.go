package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 0))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}
