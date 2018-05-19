package main

import "fmt"

//go build *.go, put those two go file together

type Student struct {
	Id   int
	Name string
}

type Node struct {
	Val  Student
	Next *Node //try Node
}

//func reverseList(p *Node) *Node {
//	var pre *Node
//	for p != nil {
//		// stash p's next node
//		next := p.Next
//		//The present node's next node is pre
//		p.Next = pre
//		// update pre , make pre as a present node
//		pre = p
//		// continue recurive, update p , make p as previous's next node
//		p = next
//	}
//	return pre
//}

func main() {
	node_a := Node{Val: Student{Id: 1, Name: "student_a"}}
	node_b := Node{Val: Student{Id: 2, Name: "student_b"}}
	node_c := Node{Val: Student{Id: 3, Name: "student_c"}}

	node_a.Next = &node_b
	node_a.Next.Next = &node_c
	p := &node_a
	//if p != nil {
	//	node := *p
	//	fmt.Println(node)
	//	p = node.Next
	//}
	printList(p)
	fmt.Println()
	printList(reverseList(p))
}
