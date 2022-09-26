package api

import (
	"encoding/json"
	"example/go/internal/services/transactionsvc"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type TransactionController struct {
	log      *zap.SugaredLogger
	validate *validator.Validate
}

func NewTransactionController(server *HttpServer, validator *validator.Validate) {
	c := &TransactionController{
		log:      server.Logger,
		validate: validator,
	}

	server.Router.Group(func(r chi.Router) {
		r.Post("/transactions", c.handleCreateTransaction)
		//r.Get("/transactions", c.handleGetAllTransactions)
	})
}

func (c *TransactionController) handleCreateTransaction(writer http.ResponseWriter, req *http.Request) {
	var payload transactionsvc.CreatePld

	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		c.log.Errorf("Malformed body")
		RenderError(req.Context(), writer, err)
		return
	}

	//if err := c.validate.Struct(payload); err != nil {
	//	c.log.Errorf("payload not valid %v", err)
	//	RenderError(req.Context(), writer, err)
	//	return
	//}

	js, er := json.Marshal(payload)

	if er != nil {
		log.Fatal("encode error:", er)
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
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

	// Produce messages to topic (asynchronously)
	topic := "myTopic"

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          js,
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)

	//res, err := c.transactionSvc.Create(req.Context(), &payload)
	//
	//if err != nil {
	//	RenderError(req.Context(), writer, err)
	//	return
	//}

	RenderJSON(req.Context(), writer, http.StatusCreated, payload)
}

//func (c *TransactionController) handleGetAllTransactions(writer http.ResponseWriter, req *http.Request) {
//	res, err := c.transactionSvc.GetAll(req.Context())
//
//	if err != nil {
//		return
//	}
//
//	RenderJSON(req.Context(), writer, http.StatusCreated, res)
//}
