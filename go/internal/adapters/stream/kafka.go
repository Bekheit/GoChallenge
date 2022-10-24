package stream

import (
	"encoding/json"
	"example/go/config"
	"example/go/resources/models"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

var configMap *kafka.ConfigMap
var message *kafka.Message

func InitializeKafka(kafkaConf config.KafkaConfigurations) {
	configMap = &kafka.ConfigMap{
		"bootstrap.servers": kafkaConf.Broker,
	}
	log.Printf("Inialize kafka: %v\n", configMap)

	message = &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaConf.Topic, Partition: kafka.PartitionAny},
	}
}

func Produce(model models.TransactionModel) {
	p, err := kafka.NewProducer(configMap)
	defer p.Close()

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	js, err := json.Marshal(model)

	if err != nil {
		log.Fatal("encode error:", err)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	message.Value = js

	err = p.Produce(message, nil)

	if err != nil {
		fmt.Printf("faled to produce message : %v", err)
	}

	log.Printf("Updated")
	p.Flush(15 * 1000)
}
