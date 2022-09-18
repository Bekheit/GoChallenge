package transactionsvc

import (
	"context"
	"example/go/internal/models"
	"example/go/internal/repositories/transaction"
	"github.com/dranikpg/dto-mapper"
	"go.uber.org/zap"
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
	var newTransaction = &models.TransactionModel{}
	dto.Map(&newTransaction, payload)
	result, err := s.transactionRepo.Create(ctx, newTransaction)

	if err != nil {
		s.log.Errorf("Fail to create a new transaction")
		return nil, err
	}

	var createResult = &CreateRes{}
	err = dto.Map(&createResult, result)
	if err != nil {
		s.log.Errorf("Failed to map transaction model to createres model")
		return nil, err
	}

	return createResult, nil
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
