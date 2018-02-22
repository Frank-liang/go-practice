package main

import (
	"fmt"
)

func print() {
	fmt.Println("hello")
}

func a(n int) int {
	fmt.Println("enter n", n)
	if n <= 1 {
		return 2
	}
	m := 2*a(n-1) + n - 1
	fmt.Println("return from a(n-1):", n)
	return m
}

func main() {
	/*
		通项公式: a(n) = 2*a(n-1) + n -1
		第一项:a(1) = 2
		求:第10项 a(10)
		a(10) = 2 * a(9) + 9 // n = 9
	*/
	fmt.Println(a(10))
}
