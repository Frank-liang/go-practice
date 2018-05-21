package main

import "fmt"

type Bird struct {
	Age  int
	Name string
}

func passV(b *Bird) {
	b.Age++
	b.Name = "Great" + b.Name
	fmt.Printf("传入修改后的Bird:\t %+v, \t内存地址：%p\n", b, &b)
}

func main() {
	parrot := &Bird{Age: 1, Name: "Blue"}
	fmt.Printf("原始的Bird:\t\t %+v, \t\t内存地址：%p\n", parrot, &parrot)
	passV(parrot)
	fmt.Printf("调用后原始的Bird:\t %+v, \t\t内存地址：%p\n", parrot, &parrot)
}

//可以看到在函数passP中，参数p是一个指向Bird的指针，传递参数给它的时候会创建指针的副本(0xc420074010)，只不过指针0xc420074000和0xc420074010都指向内存地址0xc420076000。 函数内对*T的改变显然会影响原始的对象，因为它是对同一个对象的操作。
