package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//1 此方法试用于流式处理
	//buf := new(bytes.Buffer)
	//cmd := exec.Command("ls", "-l")
	//cmd.Stdout = buf
	//cmd.Run()
	//fmt.Printf("%s", buf.String())

	//2 把命令输出写入文件
	//f, _ := os.Create("ls.out")
	//defer f.Close()
	//buf := f
	//cmd := exec.Command("ls", "-l")
	//cmd.Stdout = buf
	//cmd.Run()

	//3 字符串处理，字符串输出
	cmd := exec.Command("ls", "-l")
	out, _ := cmd.Output() //此方法只能捕获标准输出
	//cmd.CombinedOutput()   此方法捕获标准输出和标准错误输出
	fmt.Printf("%s\n", out)
}
