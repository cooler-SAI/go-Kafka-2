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

	msg1 := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("hello kafka from Go"),
	}
	partition, offset, err := producer.SendMessage(msg1)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	fmt.Printf("Message sent to Partition: %d, Offset: %d\n", partition, offset)

	msg2 := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("Sending message to consumer..."),
	}
	partition2, offset, err := producer.SendMessage(msg2)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	fmt.Printf("Message sent to Partition: %d, Offset: %d\n", partition2, offset)
}
