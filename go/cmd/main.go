package main

import (
	"example/go/config"
	"example/go/internal/adapters/api"
	"example/go/resources"
	"github.com/go-playground/validator/v10"
)

func main() {
	log, closer := resources.NewLogger()
	defer closer()

	configs := config.LoadConfig(log)

	httpServer := api.NewHttpServer(log, configs.Server)

	validate := validator.New()

	api.NewTransactionController(httpServer, validate)

	//routes.TransactionRoute(router)

	//log.Infof("before produce")
	//
	//kafka2.Producee()
	//
	//log.Infof("after produce")
	//
	//kafka2.Consumee()
	//
	//log.Infof("after consume")

	httpServer.Start()
}
