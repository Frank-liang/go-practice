package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("table.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprintf(file, "%d * %d = %-2d ", i, j, i*j)
		}
	}
}
