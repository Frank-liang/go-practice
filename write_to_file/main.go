package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("hello\n")
	fmt.Fprintf(file, "%d x %d = %d\n", 3, 4, 3*4)
}
