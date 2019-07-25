package main

import "fmt"

//A type may have a method set associated with it. The method set of an interface type is its interface. The method set of any other type T consists of all methods declared with receiver type T. The method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T (that is, it also contains the method set of T). Further rules apply to structs containing embedded fields, as described in the section on struct types. Any other type has an empty method set. In a method set, each method must have a unique non-blank method name.     The method set of a type determines the interfaces that the type implements and the methods that can be called using a receiver of that type.
type Person struct {
	Name string
	Age  int
}

func (p *Person) Grow() {
	p.Age++
}
func (p Person) DoesNotGrow() {
	p.Age++
}
func main() {
	p := Person{"JY", 10}
	p.Grow()
	fmt.Println(p.Age)
	ptr := &p
	ptr.DoesNotGrow()
	fmt.Println(p.Age)
}
