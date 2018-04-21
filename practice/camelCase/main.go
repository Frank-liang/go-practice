package main

import "fmt"
import "strings"

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1
	for _, ch := range input {
		/* First method
		min, max := 'A', 'Z'
		if ch >= min && ch <= max {
			//It is a capital letter
			answer++
		*/
		//Second one
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}

	}
	fmt.Print(answer)
}
