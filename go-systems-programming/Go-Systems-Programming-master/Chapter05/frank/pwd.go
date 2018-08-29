package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguements := os.Args

	pwd, err := os.Getwd()
	if err == nil {
		fmt.Println(pwd)
	} else {
		fmt.Println("Error: ", err)
	}

	if len(arguements) == 1 {
		return
	}

	if arguements[1] != "-P" {
		return
	}
	fileinfo, err := os.Lstat(pwd)
	fmt.Println(fileinfo.Mode)
	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, err := filepath.EvalSymlinks(pwd)
		if err == nil {
			fmt.Println(realpath)
		}
	}
}
