package main

import (
	"consumer/config"
	"consumer/internal/adapters/db"
	"consumer/internal/adapters/stream"
	"consumer/internal/repositories/transaction"
	"consumer/resources"
)

func main() {
	log, closer := resources.NewLogger()
	defer closer()

	configs := config.LoadConfig(log)

	conn := db.NewDatabaseConnection(log, configs.Database)

	transactionRepo := transaction.NewDatabaseRepository(log, conn)
	log.Info(transactionRepo)

	stream.InitializeKafka(configs.Kafka)
	stream.Consume(configs.Kafka, transactionRepo)
}
