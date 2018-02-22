package main

import "fmt"

func main() {
	var x string
	var p *int

	fmt.Println("&x=", &x)
	p = &x
	fmt.Println("p=", p)
	fmt.Println("*p", *p)

}
