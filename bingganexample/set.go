package main

import "fmt"

func main() {
	set := make(map[string]struct{})
	set["a"] = struct{}{}
	set["a"] = struct{}{}
	if _, ok := set["b"]; ok {
		fmt.Println("ok")
	} else {
		fmt.Println("!ok")
	}
}
