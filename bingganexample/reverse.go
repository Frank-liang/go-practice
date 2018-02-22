package main

import "fmt"

// 翻转切片
func reverse(s []int) {

}

func main() {
	s := []int{2, 3, 5, 7, 11}
	reverse(s)
	fmt.Println(s)
	reverse(s[1:4])
	fmt.Println(s)
}
