package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var classrooms map[string]*ClassRoom
var currentClassRoom *ClassRoom

type Student struct {
	Name string
	Id   int
}

type ClassRoom struct {
	teacher  string
	students map[string]*Student
}

func (c *ClassRoom) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["teacher"] = c.teacher
	m["students"] = c.students
	return json.Marshal(m)
}

func (c *ClassRoom) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, &c.students)
}

func (c *ClassRoom) List() {
	for _, stu := range c.students {
		fmt.Println(stu.Name, stu.Id)
	}
}

func (c *ClassRoom) Add(name string, id int) error {
	c.students[name] = &Student{
		Name: name,
		Id:   id,
	}
	return nil
}

func (c *ClassRoom) Update(name string, id int) error {
	if stu, ok := c.students[name]; ok {
		c.students[name] = &Student{
			Name: name,
			Id:   id,
		}

		c.students[name].Id = id
		stu.Id = id
	} else {
		///
	}
	return nil
}

func save() error {
	buf, err := json.Marshal(classrooms)
	if err != nil {
		return err
	}
	fmt.Println(string(buf))
	return nil
}

func choose(args []string) error {
	name := args[0]
	if classroom, ok := classrooms[name]; ok {
		currentClassRoom = classroom
	} else {
		//
	}
}

func add(args []string) error {
	name := "" //
	id := 0    //
	currentClassRoom.Add(name, id)
}

func main() {
	classrooms = make(map[string]*ClassRoom)

	classroom1 := &ClassRoom{
		students: make(map[string]*Student),
	}
	classroom1.Add("binggan", 1)
	fmt.Println("students of classroom 51reboot")
	classroom1.List()

	classroom2 := &ClassRoom{
		students: make(map[string]*Student),
	}
	classroom2.Add("binggan", 2)
	fmt.Println("students of classroom golang")
	classroom2.List()

	classrooms["51reboot"] = classroom1
	classrooms["golang"] = classroom2

	if err := save(); err != nil {
		log.Fatal(err)
	}
}
