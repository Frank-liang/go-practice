package main

import "fmt"

type abc struct {
	v int
}

func (a abc) aaaa() { //传入的是值，而不是引用
	fmt.Printf("0:%d\n", a.v)
	a.v = 1
	fmt.Printf("1:%d\n", a.v)
}

func (a *abc) bbbb() { //传入的是引用，而不是值
	fmt.Printf("2:%d\n", a.v)
	a.v = 2
	fmt.Printf("3:%d\n", a.v)
}

func (a *abc) cccc() { //传入的是引用，而不是值
	fmt.Printf("4:%d\n", a.v)
}

func (a abc) dddd() { //传入的是引用，而不是值
	fmt.Printf("5:%d\n", a.v)
}
func main() {
	aobj := abc{} //new(abc);
	aobj.aaaa()
	aobj.bbbb()
	aobj.cccc()
	aobj.dddd()
}

//传值与传指针
//当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上。

//传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。

//Go语言中string，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）

//要访问指针 p 指向的结构体中某个元素 x，不需要显式地使用 * 运算，可以直接 p.x ；
