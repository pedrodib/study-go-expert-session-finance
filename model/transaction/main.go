package transaction

import "time"

// Transaction Struct
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Transactions is an Array of Transactions
type Transactions []Transaction
