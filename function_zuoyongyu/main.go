package main

import "fmt"

var block = "package"

func main() {
	//	block := "function"
	{
		//block := "inner"
		fmt.Printf("the block is %s.\n", block)
	}
	fmt.Printf("the block is %s.\n", block)
}
