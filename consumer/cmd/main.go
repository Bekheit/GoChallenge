package main

import (
	"consumer/config"
	"consumer/internal/adapters/db"
	"consumer/internal/repositories/transaction"
	"consumer/resources"
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type CreatePld struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

func main() {
	log, closer := resources.NewLogger()
	defer closer()

	configs := config.LoadConfig(log)

	conn := db.NewDatabaseConnection(log, configs.Database)

	transactionRepo := transaction.NewDatabaseRepository(log, conn)

	c, err1 := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err1 != nil {
		panic(err1)
	}

	c.SubscribeTopics([]string{"myTopic"}, nil)

	for {
		var tr transaction.Model
		msg, err2 := c.ReadMessage(-1)

		if err2 == nil {

			if err3 := json.Unmarshal(msg.Value, &tr); err3 != nil {
				fmt.Printf("Unmarshal error %v\n", err3)
			} else {
				result, err4 := transactionRepo.Create(context.Background(), &tr)
				if err4 != nil {
					fmt.Printf("can not create a transaction, %v\n", err4)
				} else {
					fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
					fmt.Printf("transaction created %v\n", result)
				}
			}
		} else {

			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err2, msg)
		}
	}

	c.Close()
}
