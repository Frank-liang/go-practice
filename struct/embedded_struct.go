package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Students struct {
	Person
	studentId int
}

func (s Students) GetStudentId() int {
	return s.studentId
}

func (p Person) Getname() string {
	return p.name
}

func (p Person) Getage() int {
	return p.age
}

func main() {
	s := Students{
		Person: Person{
			name: "Frank",
			age:  30,
		},
		studentId: 2,
	}
	fmt.Println(s.Person.Getname(), s.Person.Getage())
}
