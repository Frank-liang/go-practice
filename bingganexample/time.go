package main

import (
	"fmt"
	"time"
)

func main() {
	var n time.Duration
	n = time.Hour
	n = 3*time.Hour + 30*time.Minute
	fmt.Println(int64(n))
	fmt.Println(n.String())
	fmt.Println(n.Seconds())
	fmt.Println(n.Minutes())

	t := time.Now()
	t1 := t.Add(time.Hour)
	fmt.Println(t1.Sub(t))
}
