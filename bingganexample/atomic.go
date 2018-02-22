package main

import (
	"fmt"
	"sync"
)

func main() {
	cnt := 0
	for {
		var n int32 = 0
		wg := new(sync.WaitGroup)
		wg.Add(2)
		var lock sync.Mutex
		go func() {
			lock.Lock()
			defer lock.Unlock()
			n = n + 2
			wg.Done()
		}()

		go func() {
			lock.Lock()
			defer lock.Unlock()
			n = n / 2
			wg.Done()
		}()

		wg.Wait()
		cnt++
		if n != 2 && n != 1 {
			fmt.Println(cnt)
			fmt.Println(n)
			panic("bingo")
		}
	}
}
