package main

import "fmt"

//原地翻转的方式实现
func reverse(s string) string {
	s1 := []rune(s)
	for x, y := 0, len(s1)-1; x < y; x, y = x+1, y-1 {
		s1[x], s1[y] = s1[y], s1[x]
	}
	return string(s1)
}
func main() {
	s := "hello!你好"
	fmt.Println(reverse(s))
}
