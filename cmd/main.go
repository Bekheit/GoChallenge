package main

import (
	"example/go/config"
	"example/go/internal/adapters/api"
	"example/go/internal/adapters/db"
	"example/go/internal/repositories/transaction"
	"example/go/internal/services/transactionsvc"
	"example/go/resources"
	"github.com/go-playground/validator/v10"
)

func main() {
	log, closer := resources.NewLogger()
	defer closer()

	configs := config.LoadConfig(log)

	httpServer := api.NewHttpServer(log, configs.Server)

	validate := validator.New()

	conn := db.NewDatabaseConnection(log, configs.Database)

	transactionRepo := transaction.NewDatabaseRepository(log, conn)

	transactionSvc := transactionsvc.NewDefaultService(log, transactionRepo)

	api.NewTransactionController(httpServer, validate, transactionSvc)

	//routes.TransactionRoute(router)

	httpServer.Start()
}
