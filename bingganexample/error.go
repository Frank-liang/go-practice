package main

import (
	"errors"
	"fmt"
)

func main() {
	var e error
	e = errors.New("an error")
	fmt.Println(e.Error())
	var cmd string
	e = fmt.Errorf("bad command:%s", cmd)
	fmt.Println(e)
}
