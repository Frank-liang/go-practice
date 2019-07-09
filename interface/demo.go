package main

import "fmt"

type Person struct {
	name string
	age  int
}

type MyInterface interface {
	Getname() string
	Getage() int
}

func (p Person) Getname() string {
	return p.name
}

func (p Person) Getage() int {
	return p.age
}

func PrintNameAndAge(i MyInterface) {
	fmt.Println(i.Getname(), i.Getage())
}
func main() {

	// First method
	//	var myInterfaceValue MyInterface
	//	var p = Person{}
	//	p.name = "jack"
	//	p.age = 30
	//
	//	myInterfaceValue = p
	//	fmt.Println(myInterfaceValue.Getage(), myInterfaceValue.Getname())

	//Second one
	p := Person{"Alice", 26}
	PrintNameAndAge(p)
}
