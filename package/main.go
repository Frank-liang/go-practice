package main

import (
	"fmt"

	"github.com/Frank-liang/go/package/a"
	"github.com/Frank-liang/go/package/b"
	//	a "github.com/Frank-liang/go/package/a"  alias
	//	b "github.com/Frank-liang/go/package/a"  alias
	//    . "github.com/Frank-liang/go/package/a"  当本地包使用
	//   . "github.com/Frank-liang/go/package/b"
	//   _ "github.com/Frank-liang/go/package/a" 只执行init函数
)

func init() {
	fmt.Println("I am main's init")
}

func main() {
	a.A()
	b.B()
}
