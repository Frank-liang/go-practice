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
	time.Sleep(time.Millisecond)
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
	// 定义waitgroup
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		account.GiveWife(6)
		wg.Done()
	}()
	go func() {
		account.Buy(5)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(account.Left())
}
