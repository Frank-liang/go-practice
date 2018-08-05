package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for ch := 'a'; ch < 'a'+26; ch++ {
				fmt.Printf("%c", ch)
			}
			fmt.Printf("\n")
		}
	}()
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for ch := 'A'; ch < 'A'+26; ch++ {
				fmt.Printf("%c", ch)
			}
			fmt.Printf("\n")
		}
	}()
	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("Finished")

}
