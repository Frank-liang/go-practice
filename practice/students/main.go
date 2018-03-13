package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

var stu_info = make(map[string]Student)

func list(args []string) error {
	for _, v := range stu_info {
		fmt.Println(v.Name, v.Id)
	}
	return nil
}

func add(args []string) err {
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	if _, ok := stu_info[name]; ok {
		fmt.Print("The name %s you input have exist already\n", name)
	} else {
		stu_info[name] = Student{
			Id:   id,
			Name: name,
		}
	}
	fmt.Printf("Add done.\n")
	return nil
}

func del(args []string) error {
	name := args[0]
	if _, ok := stu_info[name]; !ok {
		fmt.Printf("The name %s you input does not exists, please input the right name\n", name)
	} else {
		delete(stu_info, name)
	}
	return nil
}

func update(args []string) error {
	name := args[0]
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	if _, ok := stu_info[name]; !ok {
		fmt.Printf("The name %s you input does not exists, please input the right name\n", name)
	} else {
		stu_info[name] = Student{
			Id:   id,
			Name: name,
		}
	}
	fmt.Printf("update done.\n")
	return nil
}

func save(args []string) error {
	filename := args[0]
	buf, err := json.Marshal(stu_info)
	if err != nil {
		log.Fatal("marshal err :%s", err)
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprint(f, string(buf))
	fmt.Printf("save one. filename is %s \n", filename)
	return nil

}

func load(args []string) error {
	filename := args[0]
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	j_err := json.Unmarshal(f, &stu_info)
	if j_err != nil {
		log.Fatal("Unmarshal err :%s", j_err)
	}
	fmt.Printf("load done.\n")
	return nil
}

func main() {
	f := bufio.NewReader(os.Stdin)

	actionmap := map[string]func([]string) error{
		"add":    add,
		"list":   list,
		"save":   save,
		"load":   load,
		"delete": del,
		"update": update,
	}

	for {
		fmt.Print(">")
		line, _ := f.ReadString('\n')
		line = strings.TrimSpace(line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		cmd := args[0]
		args = args[1:]
		actionfunc := actionmap[cmd]
		if actionfunc == nil {
			continue
		}

		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}
	}
}
