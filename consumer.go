package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func main() {
	brokerList := []string{"localhost:9092"}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %s\n", err)
	}
	defer func(consumer sarama.Consumer) {
		err := consumer.Close()
		if err != nil {
			log.Fatalf("Error closing consumer: %s\n", err)
		}
	}(consumer)

	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetOldest)

	if err != nil {
		log.Fatalf("Error starting partition consumer: %s\n", err)
	}
	defer func(partitionConsumer sarama.PartitionConsumer) {
		err := partitionConsumer.Close()
		if err != nil {
			log.Fatalf("Error closing partition consumer: %s\n", err)
		}
	}(partitionConsumer)

	fmt.Println("Waiting for messages...")
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Message received: %s\n", string(msg.Value))
	}

}
