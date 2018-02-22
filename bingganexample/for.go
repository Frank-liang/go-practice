package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	i := 5
	for i < 7 {
		fmt.Println(i)
		i = i + 1
	}

	i = 7
	for {
		if i > 10 {
			break
		}
		i = i + 1
	}

	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	s := "hello中文"
	for i, arg := range s {
		fmt.Printf("%d %c\n", i, arg)
	}
}
