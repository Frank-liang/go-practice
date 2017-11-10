package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Teacher struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Class struct {
	teachers []Teacher `json:"teachers"`
	students []Student `json:"students"`
}

func main() {
	for {
		fmt.Printf("Input Name and Age: ")
		var name string
		var age int
		var s2 Student
		fmt.Scanf("%s%d", &name, &age)
		s1 := Student{
			Id:   1,
			Name: name,
			Age:  age,
		}
		fmt.Println(s1)
		buf, err := json.Marshal(s1)
		if err != nil {
			log.Fatalf("marshal error:%s", err)
		}
		fmt.Println(string(buf))

		err = json.Unmarshal(buf, &s2)
		if err == nil {
			log.Println(s2)
		}
	}
}
