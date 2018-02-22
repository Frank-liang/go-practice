package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

type Student struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func (s Student) String() string {
	return fmt.Sprintf("name:%s,id:%d", s.Name, s.Id)
}

func marshal(x interface{}) string {
	t := reflect.TypeOf(x).Elem()
	v := reflect.ValueOf(x).Elem()
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "{\n")
	for i := 0; i < t.NumField(); i++ {
		fieldt := t.Field(i)
		fieldv := v.Field(i)
		jsonkey := fieldt.Tag.Get("json")
		if jsonkey == "" {
			jsonkey = fieldt.Name
		}

		var jsonvalue string
		switch fieldt.Type.Kind() {
		case reflect.Int:
			jsonvalue = strconv.Itoa(int(fieldv.Int()))
		case reflect.String:
			jsonvalue = `"` + fieldv.String() + `"`
		}
		fmt.Fprintf(buf, "  \"%s\":%s,\n", jsonkey, jsonvalue)
	}
	fmt.Fprintf(buf, "}\n")
	return buf.String()
}

func print(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t.Kind())
	t = t.Elem()
	fmt.Println(t.Name())
	fmt.Println(t.PkgPath())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field)
		fmt.Println("jsonkey:", field.Tag.Get("json"))
	}

	field, _ := t.FieldByName("Name")
	fmt.Printf("%#v\n", field)

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println(method.Name)
	}

	v := reflect.ValueOf(x).Elem()
	vfiled := v.FieldByName("Name")
	fmt.Println(vfiled.String())

	method := v.MethodByName("String")
	ret := method.Call(nil)
	fmt.Println(ret[0].String())
}

func main() {
	s := &Student{
		Name: "binggan",
		Id:   1,
	}
	print(s)

	fmt.Println(marshal(s))

	s1 := []string{"hello", "world"}
	s2 := s1
	fmt.Println(reflect.DeepEqual(s1, s2))
}
