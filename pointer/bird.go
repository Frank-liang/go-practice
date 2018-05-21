package main

import "fmt"

type Bird struct {
	Age  int
	Name string
}

func passV(b Bird) {
	b.Age++
	b.Name = "Great" + b.Name
	fmt.Printf("传入修改后的Bird:\t %+v, \t内存地址：%p\n", b, &b)
}

func main() {
	parrot := Bird{Age: 1, Name: "Blue"}
	fmt.Printf("原始的Bird:\t\t %+v, \t\t内存地址：%p\n", parrot, &parrot)
	passV(parrot)
	fmt.Printf("调用后原始的Bird:\t %+v, \t\t内存地址：%p\n", parrot, &parrot)
}

//在T类型作为参数的时候，传递的参数parrot会将它的副本(内存地址0xc4200122c0)传递给函数passV,在这个函数内对参数的改变不会影响原始的对象
