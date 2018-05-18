package main

import "fmt"

func main() {
	s := []string{"tom", "tom", "jack", "jack", "duck"}
	s = removeAdjacentRepeatString(s)
	fmt.Println(s)
}

func removeAdjacentRepeatString(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
