package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name string
	id   int
}

func (s *Student) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.id)
}

func main() {
	s := &Student{
		Name: "binggan",
		id:   1,
	}
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}
