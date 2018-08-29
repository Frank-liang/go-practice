package main

import "fmt"

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

func main() {
	min, max := minMax(10, 5)
	fmt.Println(min, max)
}
