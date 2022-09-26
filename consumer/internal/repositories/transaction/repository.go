package transaction

import (
	"context"
)

type IRepository interface {
	Create(ctx context.Context, model *Model) (*Model, error)
	GetAll(ctx context.Context) (*[]Model, error)
}
