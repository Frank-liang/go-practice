package main

import "fmt"

type Students struct {
	Id   int
	Name string
}

type Teachers struct {
	Id   int
	Name string
	Sex  string
}

func main() {
	var s Students
	s.Id = 1
	s.Name = "jack"
	fmt.Println(s)
	s1 := Teachers{
		Id:   2,
		Name: "alice",
		Sex:  "M",
	}
	fmt.Println(s1)
}
