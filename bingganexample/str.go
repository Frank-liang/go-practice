package main

import "fmt"

func main() {
	s := "a"
	fmt.Println(&s)
	s = s + "b"
	fmt.Println(&s)
}
