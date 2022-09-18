package transactionsvc

import "context"

type IServive interface {
	Create(ctx context.Context, payload *CreatePld) (*CreateRes, error)
	GetAll(ctx context.Context) (*GetAllRes, error)
}
