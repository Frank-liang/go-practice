package main

import "fmt"

func main() {
	counter := newCounter()
	fmt.Println(newCounter())
	fmt.Println(newCounter())
}

func newCounter() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}
