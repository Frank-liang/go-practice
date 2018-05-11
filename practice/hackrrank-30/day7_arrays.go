package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	var a = make([]int, num)
	for j := 0; j < num; j++ {
		fmt.Scan(&a[j])
	}
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Printf("%d", a[i])
	}
	fmt.Printf("\n")
}
