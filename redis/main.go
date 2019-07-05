package main

import (
	"log"
	"time"
)

type valueEx struct {
	Name  string
	Email string
}

func main() {
	redisClient := initialize()
	key1 := "sampleKey"
	value1 := &valueEx{Name: "someName", Email: "someemail@abc.com"}
	err := redisClient.setKey(key1, value1, time.Minute*1)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}

	value2 := &valueEx{}
	err = redisClient.getKey(key1, value2)
	if err != nil {
		log.Fatal("Error: %v", err.Error())
	}

	log.Printf("Name: %s", value2.Name)
	log.Printf("Email: %s", value2.Email)
}
