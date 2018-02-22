package main

import "fmt"

func sum(args ...int) int {
	n := 0
	for i := 0; i < len(args); i++ {
		n += args[i]
	}
	return n
}

func main() {
	fmt.Println(sum(1, 2, 3))
	s := []int{1, 2, 3}
	fmt.Println(sum(s...))
	/*
		ta(n+1) = 2* a(n) + n
		//2 10
		// 2 6
	*/
}
