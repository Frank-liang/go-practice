package main

import (
	"fmt"
	"sort"
)

// GoLang sort  /home/local/go/src/sort/sort.go
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return len(p[i]) < len(p[j]) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	s := []string{"Go", "B", "Gop", "Alpha"}
	sort.Sort(StringSlice(s))
	fmt.Println(s)
}
