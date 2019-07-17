package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	num int64 = 604800
)

func walkFunction(path string, info os.FileInfo, err error) error {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	mtime := fileInfo.ModTime().Unix()
	//	sumD := now.Sub(mtime)
	//fmt.Println(now, mtime)

	mode := fileInfo.Mode()
	if mode.IsDir() {
		if now-mtime < num {
			writePathToFile(path, "/tmp/file.txt")
		}
	}
	return nil
}

func writePathToFile(path string, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("create file error: %v\n", err)
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	lineStr := fmt.Sprintf("%s", path)
	fmt.Fprintln(w, lineStr)
	return w.Flush()
}

func main() {
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	Path := flags[0]
	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
