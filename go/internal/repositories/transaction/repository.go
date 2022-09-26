package transaction

import (
	"context"
	"example/go/resources/models"
)

type IRepository interface {
	Create(ctx context.Context, model *models.TransactionModel) (*models.TransactionModel, error)
	GetAll(ctx context.Context) (*[]models.TransactionModel, error)
}
