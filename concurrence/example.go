package main

//concurrent example
import (
	"fmt"
	"sync"
)

func main() {
	t := NewTaskRunner(10)
	for i := 0; i < 10; i++ {
		i := i //这一步要注意，闭包
		t.Put(func() {
			fmt.Println(i)
		})
	}
	t.Wait()
}

type TaskRunner struct {
	token chan bool
	wg    sync.WaitGroup
}

func NewTaskRunner(concurrent int) *TaskRunner {
	return &TaskRunner{
		token: make(chan bool, concurrent),
	}
}

func (t *TaskRunner) Put(task func()) {
	t.token <- true
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		task()
		<-t.token
	}()
}

func (t *TaskRunner) Wait() {
	t.wg.Wait()
}
