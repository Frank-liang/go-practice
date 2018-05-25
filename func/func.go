package main

import "fmt"

func f(i int) func() int {
	return func() int {
		i++
		fmt.Println(i)
		return i
	}
}

func main() {
	c1 := f(1)
	c2 := f(0)
	c1()
	c2()
}
