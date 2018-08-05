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

func (h *Http) getHost() string {
	return h.host
}
func (h *Http) getAgent() string {
	return h.agent
}

func (d *data) getName() string {
	return d.name
}

func (d *data) getPassword() string {
	return d.password
}

func main() {
	var h = Http{}
	r := reflect.TypeOf(h)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	fmt.Println(r)
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		fmt.Println(field.Type, field.Name, field.Index, field.Offset)
		if field.Anonymous {
			for x := 0; x < field.Type.NumField(); x++ {
				subField := field.Type.Field(x)
				fmt.Println(subField.Type, subField.Name, subField.Index, subField.Offset)

			}
		}
	}
	userName, ok := r.FieldByName("name")
	if ok {
		fmt.Println(userName.Name, userName.Type)
	}

}
