package main

import "fmt"

func factorial(num uint) uint {
	if num > 1 {
		return num * factorial(num-1)
	} else {
		return 1
	}
}

func main() {
	var N uint
	fmt.Scanln(&N)
	fmt.Printf("%d\n", factorial(N))

}
