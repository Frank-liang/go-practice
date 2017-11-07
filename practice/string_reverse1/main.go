package main

import "fmt"

//用两个slice实现的函数
//byte 注重于raw date  是unit8
//rune 注重于unicode 是 int32
func reverse(s string) string {
	s1 := []rune(s)
	s2 := make([]rune, len(s1))
	for i := 0; i < len(s1); i++ {
		s2[i] = s1[len(s1)-1-i]
	}
	return string(s2)
}
func main() {
	s := "hello world! 你好"
	fmt.Println(reverse(s))

}
