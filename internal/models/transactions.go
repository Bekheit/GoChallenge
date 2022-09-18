package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type TransactionModel struct {
	bun.BaseModel `bun:"table:transactions""`
	ID            uuid.UUID `bun:"id,default:gen_random_uuid()"`
	Amount        float32   `bun:"amount"`
	Currency      string    `bun:"currency"`
	CreatedAt     time.Time `bun:"createdat,default:current_timestamp"`
}
