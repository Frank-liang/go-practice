package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("something", "some"))
	fmt.Println(strings.HasSuffix("something", "thing"))
}

//While there are a lot of useful common functions in the strings package
