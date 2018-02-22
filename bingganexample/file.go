package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 裸读取，很少使用
	buf := make([]byte, 10)
	n, err := f.Read(buf)
	buf[:n]

	// 加上buffer的读取，很高效
	r := bufio.NewReader(f)
	r.Read(buf)

	// 按行读取，按分隔符读取
	r1 := bufio.NewScanner(f)

	// 小文件一次性读取
	ioutil.ReadFile("a.txt")
	ioutil.ReadAll(f)

	// 操作类文件的神器
	io.Copy

}
