package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func main() {
	brokerList := []string{"localhost:9092"}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalf("Error creating the producer: %v", err)
	}
	defer func(producer sarama.SyncProducer) {
		if err := producer.Close(); err != nil {
			log.Printf("Error closing the producer: %v", err)
		}
	}(producer)

	msg := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("hello kafka from Go"),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	fmt.Printf("Message sent to Partition: %d, Offset: %d\n", partition, offset)
}
