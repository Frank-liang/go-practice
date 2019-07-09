package main

import "fmt"

func main() {
	Slice := []int{1, 2, 3, 4, 5}
	slice1 := Slice[0:1]
	slice2 := Slice[2:4]
	fmt.Printf("Slice len %d, cap %d, address %p\n", len(Slice), cap(Slice), &Slice)
	fmt.Printf("slice1 len %d, cap %d, address %p\n", len(slice1), cap(slice1), &slice1)
	fmt.Printf("slice2 len %d, cap %d\n, address %p\n", len(slice2), cap(slice2), &slice2)

	Array := [5]int{6, 7, 8, 9, 10}
	slice3 := Array[0:1]
	slice4 := Array[2:4]
	fmt.Printf("Array len %d, cap %d, address %p\n", len(Array), cap(Array), &Array)
	fmt.Printf("slice3 len %d, cap %d, address %p\n", len(slice3), cap(slice3), &slice3)
	fmt.Printf("slice4 len %d, cap %d, address %p\n", len(slice4), cap(slice4), &slice4)
}
