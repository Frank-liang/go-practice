package main

import "fmt"

func printList(p *Node) {
	for {
		if p == nil {
			fmt.Println("p is nil.")
			break
		}
		fmt.Println((*p).Val)
		p = p.Next
	}
}
