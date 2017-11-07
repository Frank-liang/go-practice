package main

import "log"
import "os"

func main() {
	file, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

}
