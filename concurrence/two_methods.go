package main

//并发的两种模式 线程池模型 和 令牌模型
//线程池模型(生产者消费者模型)
func worker(ch chan string) {
	for url := range ch {
		_ = url
	}
}

func main() {
	urls := []string{} // 1000个url
	//以10的并发来抓
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go worker(ch)
	}
	for _, url := range urls {
		ch <- url
	}
}

func main1() { //令牌的方式
	urls := []string{} // 1000个url
	token := make(chan bool, 10)
	for _, url := range urls {
		token <- true //占有令牌
		go func(url string) {
			_ = url
			<-token //释放令牌
		}(url)
	}
}
