package stream

import (
	"encoding/json"
	"example/go/config"
	"example/go/resources/models"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var configMap *kafka.ConfigMap
var message *kafka.Message

func InitializeKafka(kafkaConf config.KafkaConfigurations) {
	log.Printf("%v", kafkaConf)

	configMap = &kafka.ConfigMap{
		"bootstrap.servers": kafkaConf.Broker,
	}
	log.Printf("Inialize kafka: %v\n", configMap)

	message = &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaConf.Topic, Partition: kafka.PartitionAny},
	}
}

func Produce(model models.TransactionModel) {
	log.Printf("%v", configMap)

	p, err := kafka.NewProducer(configMap)

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	defer p.Close()

	js, err := json.Marshal(model)

	if err != nil {
		log.Fatal("Marshal error:", err)
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

	log.Printf("before produce message")
	err = p.Produce(message, nil)
	log.Printf("after produce message")

	if err != nil {
		fmt.Printf("failed to produce message : %v", err)
	}

	p.Flush(60 * 1000)
}
