package main

import (
	"fmt"
	"sync"
	"time"
)

type Accout struct {
	flag  sync.Mutex
	money int
}

func (a *Accout) DoPrepare() {
	time.Sleep(time.Second)
}

func (a *Accout) GetGongZi(n int) {
	a.money += n
}

func (a *Accout) GiveWife(n int) {
	a.flag.Lock()
	defer a.flag.Unlock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
}

func (a *Accout) Buy(n int) {
	a.flag.Lock()
	defer a.flag.Unlock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
}

func (a *Accout) Left() int {
	return a.money
}

func main() {
	var account Accout
	account.GetGongZi(10)
	// 定义channel
	ch := make(chan int)
	go func() {
		account.GiveWife(6)
		// 向channel发送数据
		ch <- 0
	}()
	go func() {
		account.Buy(5)
		// 向channel发送数据
		ch <- 0
	}()

	deadline := time.After(time.Millisecond * 100)
	// 从channel接收数据，如果收到两条数据, 等待结束
	for i := 0; i < 2; i++ {
		select {
		case <-ch:
		case <-deadline:
			fmt.Println("deadline reach")
			return
		}
	}
	fmt.Println(account.Left())
}
