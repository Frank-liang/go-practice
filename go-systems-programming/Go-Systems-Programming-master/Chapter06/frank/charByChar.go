package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	arguement := os.Args
	if len(arguement) == 1 {
		fmt.Println("Not enough arguement")
		os.Exit(1)
	}
	input := arguement[1]

	buf, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	in := string(buf)
	s := bufio.NewScanner(strings.NewReader(in))
	s.Split(bufio.ScanWords)

	for s.Scan() {
		fmt.Print(s.Text())
	}
}
