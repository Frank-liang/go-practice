package main

import "fmt"

func main() {
	fmt.Println("a return: ", a())
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2: ", i)
	}()
	defer func() {
		i++
		fmt.Println("a defer1: ", i)
	}()
	return i
}
