package main

import "fmt"

func ExFunc(n int) func() {
	return func() {
		n++ // 这里对外部变量加 1
		fmt.Println(n)
	}
}

func main() {
	myFunc := ExFunc(10)
	myFunc() // 这里输出 11

	myAnotherFunc := ExFunc(20)
	myAnotherFunc() // 这里输出 21

	myFunc()        // 这里输出 12
	myAnotherFunc() // 这里输出 22
}

//1.内函数对外函数 的变量的修改，是对变量的引用
//2.变量被引用后，它所在的函数结束，这变量也不会马上被烧毁

//闭包函数出现的条件：
//1.被嵌套的函数引用到非本函数的外部变量，而且这外部变量不是“全局变量”;
//2.嵌套的函数被独立了出来(被父函数返回或赋值 变成了独立的个体)，而被引用的变量所在的父函数已结束
