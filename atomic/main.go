package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	counter int64
	wg      sync.WaitGroup
	myLock  sync.Mutex
	myMap   = make(map[string]int64)
)

//func incCounter(id int) {
//	defer wg.Done()
//	for i := 0; i < 2; i++ {
//		myLock.Lock()
//		value := counter
//		//atomic.AddInt64(&counter, 1)
//		runtime.Gosched()
//		value++
//		counter = value
//		myLock.Unlock()
//	}
//}

func opReadMap() {
	defer wg.Done()
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}

func opWriteMap() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		myLock.Lock()
		myMap[strconv.Itoa(i)] = int64(i)
		myLock.Unlock()
	}
}

func main() {
	//runtime.GOMAXPROCS(1)
	wg.Add(2)
	//	go incCounter(2)
	//	go incCounter(1)
	//	go incCounter(3)
	go opReadMap()
	go opWriteMap()
	wg.Wait()
}
