package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

func add(args []string) error {
	fmt.Println("call add")
	fmt.Println("args", args)
	/*
		name := args[0]
		id := args[1]
	*/
	// ....
	return nil
}

func list(args []string) error {
	return errors.New("unimplemention")
}

func main() {
	actionmap := map[string]func([]string) error{
		"add":  add,
		"list": list,
	}

	f := bufio.NewReader(os.Stdin)

	//var students map[string]Student
	for {
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		// 去除两端的空格和换行
		line = strings.TrimSpace(line)
		// 按空格分割字符串得到字符串列表
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		// 获取命令和参数列表
		cmd := args[0]
		args = args[1:]

		// 获取命令函数
		actionfunc := actionmap[cmd]
		if actionfunc == nil {
			fmt.Println("bad cmd ", cmd)
			continue
		}
		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}
	}
}
