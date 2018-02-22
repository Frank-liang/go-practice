package main

import "fmt"

func main() {
	a := [...]int{2, 3, 5, 7, 11, 13}
	s := a[:0]
	fmt.Println(s, len(s), cap(s))
	s = s[:4]
	fmt.Println(s, len(s), cap(s))
	s = s[2:]
	fmt.Println(s, len(s), cap(s))
	var s1 []int
	if s1 == nil {
		fmt.Println("nil")
	}
	fmt.Println(len(s1))
	s1 = a[:0]
	fmt.Println(s1 == nil)

	a = []int{0, 0, 0, 0, 0}
	make([]int, 5)
}
