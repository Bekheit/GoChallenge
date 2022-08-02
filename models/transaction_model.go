package models

import "time"

type Transaction struct {
	ID        string    `json:"id"`
	Amount    float32   `json:"amount"`
	Currency  string    `json:"currency" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}
