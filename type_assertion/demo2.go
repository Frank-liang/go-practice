package main

import (
	"fmt"
)

//var container = []string{"zero", "one", "two"}
var container = "hello"

func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", containerI)
		return
	}
	return

}

func main() {
	//container := map[int]string{0: "zero", 1: "one", 2: "two"}

	elem, err := getElement(container)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n",
		elem, container)

}
