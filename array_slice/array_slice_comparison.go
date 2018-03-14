package main

import (
	"fmt"
)

//The slice type is an abstraction built on top of Go's array type
//Array their most common purpose in Go is to hold storage for a slice

func main() {
	//典型数组初始化赋值
	var a [3]int = [3]int{1, 2, 3}
	var ap = &a   //取数组指针赋值给ap
	ap[1] = 8     //对ap指向的数组进行修改
	(*ap)[2] = 88 //完全等价于上一行

	//初始化一个隐式申明长度的数组
	b := [...]int{1, 2, 3}
	//初始化一个slice切片
	c := []int{1, 2, 3}
	//规定下标初始化切片
	d := []int{0: 1, 10: 11, 12: 111}
	d[1] = 123
	e := d //切片赋值，浅拷贝
	e[1] = 456
	var f [3]int
	f = a //数组赋值，深拷贝
	f[1] = 888
	copy(b[:], c) //把切片赋值给数组的方式
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", ap, ap)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	e[2] = 999
	fmt.Printf("%v,%p, %T\n", d, d, d)
	fmt.Printf("%v,%p, %T\n", e, e, e)
	fmt.Printf("%v,%p, %T\n", f, &f, f)
	//数组比较是深比较
	fmt.Println(b == a)
	b = a
	fmt.Println(b == a)
	//切片只能和nil比较
	fmt.Println(d == e)
}
