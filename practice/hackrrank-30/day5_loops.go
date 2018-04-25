package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	for i := 1; i < 11; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
}
