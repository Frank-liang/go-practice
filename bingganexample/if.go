package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "abc123"
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

	if n, err := strconv.Atoi(s); err == nil {
		fmt.Println(n)
	}

	if s == "1" {

	} else if s == "2" {

	} else if s == "3" {

	}

	switch s {
	case "1":
		fmt.Println("s=", "1")
	case "2":
		fmt.Println("s=", "2")

	default:
		fmt.Println("s=", "default", s)
	}
}
