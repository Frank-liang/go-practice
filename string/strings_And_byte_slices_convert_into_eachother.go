package main

import "fmt"

func main() {
	var s string = "this is a string"
	fmt.Println(s)
	var b []byte
	b = []byte(s)
	fmt.Println(b)
	for i := range b {
		fmt.Println(string(b[i]))
	}
	s = string(b)
	fmt.Println(s)
}
