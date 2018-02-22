package main

import "fmt"

定义 学生 struct {
	名字 string
	编号 int
}

定义 教室 struct {
	学生们 []学生
}

定义 路径 []点

func main() {
	var 学生1 学生
	学生1.名字 = "binggan"
	fmt.Println(学生1)
}
