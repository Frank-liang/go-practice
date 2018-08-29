package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguements := os.Args
	if len(arguements) != 3 {
		fmt.Printf("Usage: %s message filename \n", filepath.Base(arguements[0]))
		os.Exit(1)
	}

	message := arguements[1]
	filename := arguements[2]

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", message)
}
