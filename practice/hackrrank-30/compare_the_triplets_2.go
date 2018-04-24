package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(a []string) {
	a1 := a[:3]
	a2 := a[3:6]
	score := []int{0, 0}
	for i := 0; i < 3; i++ {
		i1, _ := strconv.Atoi(a1[i])
		i2, _ := strconv.Atoi(a2[i])
		if i1 > i2 {
			score[0]++
		}
		if i2 > i1 {
			score[1]++
		}
	}
	//fmt.Printf("%d %d\n", score[0], score[1])
	fmt.Printf("%d\n", score)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	items := []string{}
	for i := 0; i < 2; i++ {
		txt, _, _ := reader.ReadLine()
		d := strings.Split(string(txt), " ")
		items = append(items, d...)
	}
	solve(items)
}
