package main

import "fmt"

func main() {
	// byte => uint8
	// rune => int32
	s := "golang你好"
	fmt.Println(len(s))
	cnt := 0
	for _, r := range s {
		cnt += 1
		fmt.Printf("%c\n", r)
	}
	ss := []rune("hello")

	fmt.Println("cnt", cnt)
}
