package main

import "fmt"

func main() {
	x, y := 1, 2

	defer func(a int) {
		fmt.Printf("x:%d,y:%d\n", a, y) // y 为闭包引用
	}(y) // 复制 x 的值

	x += 100
	y += 100
	fmt.Println(x, y)
}
