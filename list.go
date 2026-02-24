package main

import "time"

type Expense struct {
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	CreatedAt   time.Time  `json:"created_at"`
}