package main

import (
	"example/go/config"
	"example/go/internal/adapters/api"
	"example/go/internal/adapters/db"
	"example/go/internal/adapters/stream"

	"example/go/internal/repositories/transaction"
	"example/go/internal/services/transactionsvc"
	"example/go/resources"

	"github.com/go-playground/validator/v10"
)

func main() {
	log, closer := resources.NewLogger()
	defer closer()

	log.Info("start load configurations")
	configs := config.LoadConfig(log)

	log.Info("start httpserver")
	httpServer := api.NewHttpServer(log, configs.Server)

	validate := validator.New()

	log.Info("start connecting to database")
	conn := db.NewDatabaseConnection(log, configs.Database)

	log.Info("start initalize repo")
	transactionRepo := transaction.NewDatabaseRepository(log, conn)

	log.Info("start initalize service")
	transactionSvc := transactionsvc.NewDefaultService(log, transactionRepo)

	log.Info("start initalize controller")
	api.NewTransactionController(httpServer, validate, transactionSvc)

	log.Info("start initalize kafka")
	stream.InitializeKafka(configs.Kafka)

	log.Info("start server")
	httpServer.Start()
}
