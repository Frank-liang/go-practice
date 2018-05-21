package main

import "fmt"

func min_max(args ...int) (min int, max int) {
	min = args[0]
	max = args[0]
	for i := 0; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
		if args[i] < min {
			min = args[i]
		}
	}
	return
}

func main() {
	fmt.Println(min_max(1, 2, 3, 8))
}
