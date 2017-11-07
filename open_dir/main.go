package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	infos, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		if info.IsDir() {
			fmt.Println(info.Name())
		}
	}
	f, _ := os.Open("/proc")
	infos, err = f.Readdir(-1)
}
