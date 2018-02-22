package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	s := "/51reboot/golang-01-homework/blob/master/lesson7/CHECK.md"
	//s1 := "\\51reboot\\golang-01-homework\\blob\\master\\lesson7\\CHECK.md"
	dir := filepath.Dir(s)
	name := filepath.Base(s)
	fullname := filepath.Join(dir, name)
	fmt.Println(dir)
	fmt.Println(name)
	fmt.Println(fullname)
}
