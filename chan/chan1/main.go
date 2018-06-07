package main

import (
	"fmt"
	"sync"
	"time"
)

func sortList(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println(i)
	ch <- i
}

func main() {
	var ch = make(chan int)
	var quit = make(chan bool)
	var wg sync.WaitGroup
	go func() {
		for {
			select {
			case c := <-ch:
				fmt.Println(c)
			case <-quit:
				fmt.Println("bye")
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	s := [5]int{1, 5, 2, 7, 4}
	for _, i := range s {
		wg.Add(1)
		go sortList(i, ch, &wg)
	}
	wg.Wait()
	close(ch)
	quit <- true
}
