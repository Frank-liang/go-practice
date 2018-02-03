package main

import "sync"

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
	close(ch)
}

func main2() { //加锁控制
	urls := []string{} // 1000个url
	//以10的并发来抓
	wg := new(sync.WaitGroup)
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ch)
		}()
	}
	for _, url := range urls {
		ch <- url
	}
	close(ch)
	wg.Wait()

	//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//	go func() {
	//		time.Sleep(time.Second)
	//		cancel()
	//	}()

	//  go Handleredis(ctx)
	//	t := NewTaskRunner(10)
	//forloop:
	//	for _, url := range url {
	//		url := url
	//		select {
	//			case <-ctx.Done():
	//				break forloop
	//			default:
	//		}
	//		t.Put(func() {
	//			req, _ := http.NewRequest("GET", url, nil)
	//			req = req.WithContext(ctx)
	//			http.DefaultClient.Do(req)
	//		})
	//	}
	//	t.Wait()  这一段代码是为超时写的, 10s钟，即使channel任务没有完成，也会结束,为所有的协程设置deadline

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancle()
	//dialer := new(net.Dialer)
	//conn, err := dialer.DialContext(ctx, "tcp", "")
	//deadline, _ :=ctx.Deadline()
	//conn.SetReadDeadline(deadline)
	//conn.Read()  这段代码设置的是 dail 的deadline

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
