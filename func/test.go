package main

import "fmt"

var DoStuff func() = func() {
	fmt.Println("DoStuff func")
}

func RegFunc() { fmt.Println("reg func") }

func main() {
	DoStuff()
	DoStuff()
	DoStuff = RegFunc
	DoStuff()
}
