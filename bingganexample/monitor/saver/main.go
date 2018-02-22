package main

import (
	"context"
	"log"
	"time"

	elastic "gopkg.in/olivere/elastic.v5"

	cluster "github.com/bsm/sarama-cluster"
)

func indexName() string {
	date := time.Now().Format("20060102")
	return "falcon-" + date
}

func main() {
	consumer, err := cluster.NewConsumer([]string{"59.110.12.72:9092"},
		"falcon-saver",
		[]string{"falcon"},
		cluster.NewConfig())

	if err != nil {
		log.Fatal(err)
	}

	esclient, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case msg := <-consumer.Messages():
			_, err := esclient.Index().
				Index(indexName()).
				Type("falcon").
				BodyString(string(msg.Value)).
				Do(context.TODO())
			if err != nil {
				log.Print(err)
			}
			log.Print(string(msg.Value))
		case err := <-consumer.Errors():
			log.Print(err)
		}
	}
}
