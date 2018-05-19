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

func reverseList(p *Node) *Node {
	var pre *Node
	for p != nil {
		// stash p's next node
		next := p.Next
		//The present node's next node is pre
		p.Next = pre
		// update pre , make pre as a present node
		pre = p
		// continue recurive, update p , make p as previous's next node
		p = next
	}
	return pre
}
