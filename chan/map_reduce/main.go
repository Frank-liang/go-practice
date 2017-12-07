package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func mapAction(n int, ch chan int) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	ch <- n * 2
}

func reduceAction(length int, ch chan int) {
	cnt := 0
	sum := 0
	for {
		n := <-ch
		sum += n
		cnt++
		log.Println("cnt", n)
		if cnt == length {
			break
		}
	}
	fmt.Println(sum)
}
func main() {
	rand.Seed(time.Now().Unix()) //随机数的种子，没有种子，随机数的结果是固定的
	list := []int{1, 2, 3, 4, 5, 6}
	ch := make(chan int)
	for _, n := range list {
		go mapAction(n, ch)
	}
	reduceAction(len(list), ch)
}
