package transaction

import (
	"context"
	"example/go/internal/models"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type DatabaseRepository struct {
	log *zap.SugaredLogger
	db  *bun.DB
}

func NewDatabaseRepository(log *zap.SugaredLogger, conn *bun.DB) *DatabaseRepository {
	return &DatabaseRepository{
		log: log,
		db:  conn,
	}
}

func (r *DatabaseRepository) Create(ctx context.Context, model *models.TransactionModel) (*models.TransactionModel, error) {
	_, err := r.db.NewInsert().Model(model).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *DatabaseRepository) GetAll(ctx context.Context) (*[]models.TransactionModel, error) {
	var transactions = &[]models.TransactionModel{}
	err := r.db.NewSelect().Model(transactions).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
