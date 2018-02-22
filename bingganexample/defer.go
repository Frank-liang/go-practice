package main

import (
	"log"
	"os"
)

func print() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

func main() {
	print()
}
