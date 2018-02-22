package main

import "fmt"

func main() {
	names := [4]string{
		"a",
		"b",
		"c",
		"d",
	}

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	/*
		                b = names[1:3]
				start := &names[1]

				p := start + 1
				*p = "XXX"
		                等价于
		                b[1] = "XXX"
	*/

	b[1] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	c := a[1:2]
	c[0] = "YYY"

	var p [2]*string
	p[0] = &names[0]
	p[1] = &names[1]
	*p[0] = "AAA"

	/*
		var start *string
		var length int
		//var cap int
		start = &names[1]
		length = 2
		start + 0 => b[0]
		start + 1 => b[1]
	*/

}
