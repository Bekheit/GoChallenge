package transactionsvc

import (
	"context"
	"encoding/json"
	"example/go/internal/repositories/transaction"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dranikpg/dto-mapper"
	"go.uber.org/zap"
	"log"
)

type DefaultService struct {
	log             *zap.SugaredLogger
	transactionRepo transaction.IRepository
}

func NewDefaultService(log *zap.SugaredLogger, tr transaction.IRepository) *DefaultService {
	return &DefaultService{
		log:             log,
		transactionRepo: tr,
	}
}

func (s *DefaultService) Create(ctx context.Context, payload *CreatePld) (*CreateRes, error) {

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

	//var newTransaction = &models.TransactionModel{}
	//dto.Map(&newTransaction, payload)
	//result, err := s.transactionRepo.Create(ctx, newTransaction)
	//
	//if err != nil {
	//	s.log.Errorf("Fail to create a new transaction")
	//	return nil, err
	//}
	//
	//var createResult = &CreateRes{}
	//err = dto.Map(&createResult, result)
	//if err != nil {
	//	s.log.Errorf("Failed to map transaction model to createres model")
	//	return nil, err
	//}
	//
	return nil, nil
}

func (s *DefaultService) GetAll(ctx context.Context) (*GetAllRes, error) {
	result, err := s.transactionRepo.GetAll(ctx)

	if err != nil {
		s.log.Errorf("Cannot get all transactions")
		return nil, err
	}

	var getAllRes = GetAllRes{}
	err = dto.Map(&getAllRes.Transactions, result)
	if err != nil {
		s.log.Errorf("Failed to map transactions model to getres model")
		return nil, err
	}
	return &getAllRes, nil
}

//func Produce() {
//
//	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
//	if err != nil {
//		panic(err)
//	}
//
//	defer p.Close()
//
//	// Delivery report handler for produced messages
//	go func() {
//		for e := range p.Events() {
//			switch ev := e.(type) {
//			case *kafka.Message:
//				if ev.TopicPartition.Error != nil {
//					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
//				} else {
//					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
//				}
//			}
//		}
//	}()
//
//	// Produce messages to topic (asynchronously)
//	topic := "myTopic"
//
//	p.Produce(&kafka.Message{
//		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
//		Value:          []byte(word),
//	}, nil)
//
//	// Wait for message deliveries before shutting down
//	p.Flush(15 * 1000)
//}
