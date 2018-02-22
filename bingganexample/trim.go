package main

import "fmt"

type StringStruct struct {
	pointer *byte
	length  int
}

func main() {
	s := "hello\n"
	s1 := s[:len(s)-1]

	s1 = s
	s1.length -= 1
	fmt.Println(s)
}
