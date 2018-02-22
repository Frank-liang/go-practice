package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	if true {
		i := 4
		fmt.Println(i)
	}
	fmt.Println(s)
}
