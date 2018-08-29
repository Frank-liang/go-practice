package main

import (
	"flag"
	"fmt"
)

func main() {
	mO := flag.Bool("o", false, "o")
	mC := flag.Bool("c", false, "c")
	mK := flag.Int("k", 0, "an int")
	flag.Parse()

	fmt.Println("-o:", *mO)
	fmt.Println("-c:", *mC)
	fmt.Println("-k:", *mK)

	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}

}
