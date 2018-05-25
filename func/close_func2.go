package main

import "fmt"

func ExFunc(n int) func() {
	sum := n
	a := func() { // 把匿名函数作为值赋给变量a (Go 不允许函数嵌套, 然而你可以利用匿名函数实现函数嵌套)
		fmt.Println(sum + 1) // 调用本函数外的变量
	} // 这里没有()匿名函数不会马上执行
	return a
	//  或者直接 return 匿名函数
	//  return func() { //直接在返回处的匿名函数
	//    fmt.Println(sum + 1)
	//  }
}

func main() {
	myFunc := ExFunc(10)
	myFunc() // 这里输出11

	myAnotherFunc := ExFunc(20)
	myAnotherFunc() // 这里输出21

	myFunc()        // 这里输出11
	myAnotherFunc() // 这里输出21
}
