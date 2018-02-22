package main

import (
	"bufio"
	"log"
	"net"

	"github.com/Shopify/sarama"
)

// 建立listen socket
// 接收新连接
// 从连接按行读取数据
// 反序列化成common.Metric

func handle(conn net.Conn, ch chan<- *sarama.ProducerMessage) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Print(err)
		}
		if len(line) == 0 {
			continue
		}
		line = line[:len(line)-1]
		/*
			var metric common.Metric
			err = json.Unmarshal([]byte(line), &metric)
			metric := new(common.Metric)
			err = json.Unmarshal([]byte(line), metric)
			if err != nil {
				log.Print(err)
				continue
			}
			log.Print(metric)
		*/
		message := &sarama.ProducerMessage{
			Topic: "falcon",
			Key:   nil,
			Value: sarama.StringEncoder(line),
		}
		ch <- message
	}
}

func main() {
	l, err := net.Listen("tcp", ":6001")
	if err != nil {
		log.Fatal(err)
	}
	producer, err := sarama.NewAsyncProducer([]string{"59.110.12.72:9092"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	ch := producer.Input()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handle(conn, ch)
	}
}
