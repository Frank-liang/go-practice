package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]

	des, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.create: ", err)
		os.Exit(1)
	}
	defer des.Close()
	fmt.Fprintf(des, "[%s]: ", filename)
	fmt.Fprintf(des, "Using fmt.print in %s\n", filename)
}
