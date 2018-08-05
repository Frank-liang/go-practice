package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	name     string
	password string
}

type Http struct {
	host  string
	agent string
	Data
}

func (h *Http) GetHost() string {
	return h.host
}
func (h *Http) GetAgent() string {
	return h.agent
}

func (d *Data) GetName() string {
	return d.name
}

func (d *Data) GetPassword() string {
	return d.password
}

func main() {
	var h Http

	t := reflect.TypeOf(&h)

	//range method
	s := []reflect.Type{t, t.Elem()}
	for _, m := range s {
		fmt.Println(m, m.NumMethod())
		for i := 0; i < m.NumMethod(); i++ {
			fmt.Println(m.Method(i))
		}
	}
}
