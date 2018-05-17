package main

import "fmt"
import "unsafe"

func main() {
	//Default int array value is 0 , Default string array value is null
	// var a [3]int equal   a :=[3]int{0,0,0}
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	var a1 [3]string
	fmt.Println(a1[0])
	fmt.Println(a1[len(a1)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for i := range a {
		fmt.Printf("%d %d\n", i, a[i])
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	//Initialize array
	var b [3]int = [3]int{1, 4, 3}
	fmt.Printf("b: %v, len: %d, size: %v\n", b, len(b), unsafe.Sizeof(b))

	c := [...]int{1, 2, 3}
	c[2] = 0
	fmt.Printf("c: %v, len: %d, size: %v\n", c, len(c), unsafe.Sizeof(c))
	//use index, just like d[1] =2 d[3] =3
	d := [4]int{1: 2, 3: 3}
	for k, v := range d {
		fmt.Println(k, v)
	}

	//array 内存连续,8byte
	for i := 0; i < 3; i++ {
		fmt.Printf("b[%d]:%d, &b[%d]: %p, &b: %p\n", i, b[i], i, &b[i], &b)
	}
}
