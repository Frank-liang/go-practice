package main

import (
	//"bufio"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Id   int
	Name string
}

var studMap = make(map[string]Student)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var cmd string
		var name string
		var newname string
		var filename string
		var id int
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "add":
			fmt.Sscan(line, &cmd, &id, &name)
			add(id, name)
		case "list":
			fmt.Sscan(line, &cmd)
			list()
		case "del":
			fmt.Sscan(line, &cmd, &name)
			del(name)
		case "update":
			fmt.Sscan(line, &cmd, &id, &name, &newname)
			update(id, name, newname)
		case "save":
			fmt.Sscan(line, &cmd, &filename)
			save(filename)
		case "load":
			fmt.Sscan(line, &cmd, &filename)
			load(filename)

		default:
			usage()

		}
	}
}

func add(id int, name string) {
	if _, ok := studMap[name]; ok {
		fmt.Printf("Duplicate name: %s\n", name)
		return
	}
	student := Student{Id: id, Name: name}
	studMap[name] = student
}

func list() {
	for _, stu := range studMap {
		fmt.Printf("%d\t%s\n ", stu.Id, stu.Name)
	}
}

func del(name string) {
	if _, ok := studMap[name]; !ok {
		fmt.Printf("This %s doesn't exist, please input the right name!\n", name)
	} else {
		delete(studMap, name)
	}
}

func update(id int, name string, newname string) {
	if _, ok := studMap[name]; !ok {
		fmt.Printf("This %s doesn't exist, please input the right name!\n", name)
	} else {
		studMap[name] = Student{
			Id:   id,
			Name: newname,
		}

	}
}

func save(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf, err := json.Marshal(studMap)
	if err != nil {
		panic(err)
	}
	f.Write(buf)
}

func load(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(buf, &studMap)
}

func usage() {}
