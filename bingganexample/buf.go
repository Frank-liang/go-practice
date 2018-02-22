package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReaderSize(f, 10)
	line, _ := r.ReadString('\n')
	fmt.Print(line)

	io.Copy(os.Stdout, f)
}
