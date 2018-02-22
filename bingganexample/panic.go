package main

import "fmt"

func print() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	var p *int
	fmt.Println(*p)
}
func main() {
	print()
	panic("不想执行下去了")

	var i = 3
	var slice [3]int
	fmt.Println(slice[i])
}
