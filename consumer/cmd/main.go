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

	log.Info("start load configurations")
	configs := config.LoadConfig(log)

	log.Info("start connect to database")
	conn := db.NewDatabaseConnection(log, configs.Database)

	log.Info("start initalize repo")
	transactionRepo := transaction.NewDatabaseRepository(log, conn)

	log.Info("start kafka server")
	stream.InitializeKafka(configs.Kafka)

	log.Info("start kafka consumer")
	stream.Consume(configs.Kafka, transactionRepo)
}
