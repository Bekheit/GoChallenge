package stream

import (
	"consumer/config"
	"consumer/internal/repositories/transaction"
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var configMap *kafka.ConfigMap

func InitializeKafka(kafkaConf config.KafkaConfigurations) {
	configMap = &kafka.ConfigMap{
		"bootstrap.servers": kafkaConf.Broker,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}
}

func Consume(kafkaConf config.KafkaConfigurations, repo transaction.IRepository) {
	c, err := kafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{kafkaConf.Topic}, nil)

	for {
		var tr transaction.Model
		msg, err := c.ReadMessage(-1)

		if err == nil {

			if err := json.Unmarshal(msg.Value, &tr); err != nil {
				fmt.Printf("Unmarshal error %v\n", err)
			} else {
				result, err := repo.Create(context.Background(), &tr)
				if err != nil {
					fmt.Printf("can not create a transaction, %v\n", err)
				} else {
					fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
					fmt.Printf("transaction created %v\n", result)
				}
			}
		} else {

			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
