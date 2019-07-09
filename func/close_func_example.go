package main

import "fmt"

func main() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
}

func newCounter() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}
