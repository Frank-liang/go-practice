package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// 给定一个url打印url和url的status
// www.baidu.com 200 OK
func printUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(url, resp.Status)
}

func work(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range ch {
		printUrl(url)
	}

	// 等价于
	for {
		url, ok := <-ch
		if !ok {
			break
		}
		printUrl(url)
	}
}

// 1. 只要不close可以永远发送数据和接受数据
// 2. 如果channel里面没有数据，接收方会阻塞
// 3. 如果没有人正在等待channel的数据，发送方会阻塞
// 4. 从一个close的channle取数据永远不会阻塞，同时获取的是默认值

// 主协程启动一个work协程，同时传递一个channel
// 主协程向channel里面发送一个url
// work协程从channel里面获取url,之后调用printUrl打印url

// 启动3个协程
// 主协程向channel里面发送多个url，发送完毕之后关闭channel
// work协程从channel里面获取url,之后调用printUrl打印url
// work协程不停重复第三条，直到channel关闭

// 创建一个WaitGroup
// 调用Add
// 调用Wait等待work协程结束
func main() {
	ch := make(chan string)
	wg := new(sync.WaitGroup)
	//wg.Add(3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work(ch, wg)
	}

	for i := 0; i < 5; i++ {
		url := "http://www.baidu.com"
		ch <- url
	}
	close(ch)

	wg.Wait()
}
