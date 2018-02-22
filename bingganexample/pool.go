package main

import (
	"log"
	"net/http"
	"sync"
)

func work(ch chan string, wg *sync.WaitGroup) {
	for u := range ch {
		resp, err := http.Get(u)
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("%s:%s", u, resp.Status)
		resp.Body.Close()
	}
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(5)
	taskch := make(chan string)
	for i := 0; i < 5; i++ {
		go work(taskch, wg)
	}

	urls := []string{"http://www.baidu.com", "http://www.zhihu.com"}
	for _, url := range urls {
		taskch <- url
	}
	close(taskch)
	wg.Wait()
}
