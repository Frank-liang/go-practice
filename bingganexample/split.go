package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	count := make(map[string]int)
	words := strings.Fields(string(content))
	for _, word := range words {
		if word存在于count {
			count里面对应的word的value加1
		} else {
			置count里面对应的word的value为初值1
		}
	}

	for word, cnt := range count {
		fmt.Println(word, cnt)
	}
}
