package transactionsvc

import (
	"github.com/google/uuid"
	"time"
)

type CreatePld struct {
	Amount   float32 `json:"amount" validate:"required"`
	Currency string  `json:"currency"`
}

type CreateRes struct {
	ID        uuid.UUID `json:"id"`
	Amount    float32   `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetRes struct {
	ID        uuid.UUID `json:"id"`
	Amount    float32   `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetAllRes struct {
	Transactions []GetRes
}
