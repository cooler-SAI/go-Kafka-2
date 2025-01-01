package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func main() {
	brokerList := []string{"localhost:9092"}
	config := sarama.NewConfig()

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalf("Error creating producer: %s\n", err)
	}
	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {
			log.Fatalf("Error closing producer: %s\n", err)
		}
	}(producer)

	msg := &sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("hello Kafka from Go"),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Error sending message: %s\n", err)

	}
	fmt.Printf("Partition: %d, Offset: %d\n", partition, offset)
}
