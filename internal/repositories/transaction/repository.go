package transaction

import (
	"context"
	"example/go/internal/models"
)

type IRepository interface {
	Create(ctx context.Context, model *models.TransactionModel) (*models.TransactionModel, error)
	GetAll(ctx context.Context) (*[]models.TransactionModel, error)
}
