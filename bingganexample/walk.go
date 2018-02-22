package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println(path)
		}
		return nil
	})
}
