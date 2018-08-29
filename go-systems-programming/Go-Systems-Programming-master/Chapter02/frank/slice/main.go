package main

import "fmt"

func main() {
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	fmt.Println(threeD[0][0][1])
}
