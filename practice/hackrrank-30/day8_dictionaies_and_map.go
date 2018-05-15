package main

import (
	"fmt"
)

func scanf(name *string) bool {
	readItems, _ := fmt.Scanf("%s", name)
	return readItems > 0
}

func main() {
	nameToNumber := make(map[string]uint)
	var num uint
	fmt.Scanln(&num)
	var name string
	var phone uint

	for num > 0 {
		fmt.Scanf("%s %d", &name, &phone)
		nameToNumber[name] = phone
		num--
	}

	var queryName string

	for scanf(&queryName) {
		phone, ok := nameToNumber[queryName]
		if ok {
			fmt.Printf("%s=%d\n", queryName, phone)
		} else {
			fmt.Printf("Not found\n")
		}
	}
}
