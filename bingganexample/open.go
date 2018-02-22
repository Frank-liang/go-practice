package main

import (
	"os"
)

func main() {
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0644)
	f.Close()
}
