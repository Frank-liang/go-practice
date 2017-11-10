package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	infos, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		if info.IsDir() {
			id, _ := strconv.Atoi(info.Name())
			cmd, _ := ioutil.ReadFile("/proc/" + strconv.Itoa(id) + "/cmdline")
			fmt.Printf("%5d  %s\n", id, cmd)
		}
	}
}
